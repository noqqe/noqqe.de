---
layout: post
title: "Traffic Analysis in pf"
date: 2015-03-17T13:02:00+02:00
comments: true
categories:
- OpenBSD
- PF
- Firewall
- Network
- Netflow
- label
- pflog
- logging
- analysis
---

Für [devnull-as-a-serivce](http://devnull-as-a-service.com) hoste ich auf Port 9
über den `inetd` einen Discard Service. Aus historischen Gründen. Backwards
compatibility. Letztens hatte ich bereits über ein paar Leute
[geschrieben](/blog/2014/11/04/hartes-discard-protokoll/), die sehr lange
Connections zu Port 9 offen hielten. Über Monate. Aber wie viel Traffic haben
diese Connections eigentlich verbraucht. Und überhaupt, welche Services
verbrauchen auf meiner Kiste wie viel Traffic?

Seit ich `Tor` Relays auf meinen Maschinen betreibe habe ich darauf nochmal ein
extra Augenmerk. Einfach um abwägen zu können wie viel Bandwith ich Tor zuweisen
kann und wo ich Limits setzen muss. Zum Vergleich, Rechnung meines Providers vor
und nach Tor Installation:

{{< figure src="/uploads/2015/03/invoice.png" >}}

`pf` bietet viele verschiedene Möglichkeiten Traffic zu
analysieren. Drei davon habe ich mir mal angesehen.

### pflog

Das erste worauf man stößt ist `pflog`. Das Pseudo-Interface, von dem per
default eines unter /dev/pflog0 existiert aber beliebig viele pflogX nachgeladen
werden können.

{{< figure src="/uploads/2015/03/logging.jpg" >}}

Logging wird einfach in die spezielle Regel eingebettet.

```
pass in log (all, to pflog0) on $extif proto tcp from any to any port 9001
```

das sympathische daran ist, dass sowohl über Regeln als auch über den
eigentlichen `tcpdump` der zum Auslesen des Logdevices benutzt wird, noch
gefiltert werden kann.

```
$ tcpdump -n -e -ttt -i pflog0
tcpdump: WARNING: snaplen raised from 116 to 160
tcpdump: listening on pflog0, link-type PFLOG
Mar 16 09:38:10.163701 rule 38/(match) pass in on em0: 166.84.7.148.55675 > 185.34.0.188.9001: P 629511846:629512411(565) ack 110187399 win 605 <nop,nop,timestamp 131188075 3701794136>
```

Über mehrere Logdevices uns vordefinierte `tcpdump` Calls kann daraus ein
ziemlich advancedes Setup entstehen.

### pflow

`pf` ist außerdem im Stande Netflow Daten zu generieren. Das hört sich ziemlich
overengineert an für eine einzelne Kiste. NetFlow wollt ich aber auch schon
immer mal ausprobieren.

{{< figure src="/uploads/2015/03/netflow.jpg" >}}

Kurz gesagt braucht man nur einen NetFlow Sensor und einen NetFlow Collector
Zuerst installiert man einen Collector Dienst. Am einfachsten und
minimalistischsten ist hier `flowd`, der von Damien Miller geschrieben wurde.
Kann eigentlich nur gut sein. Ein paar Restrictions und das Ziel für das Logfile
konfiguriert und fertig.

```
$ pkg_add flowd
$ cat /etc/flowd.conf
logfile "/var/log/flowd"
listen on 127.0.0.1:3001
flow source 127.0.0.1
store ALL
discard all
accept agent 127.0.0.1
$ flowd -gf /etc/flowd.conf
```

Der Sensor ist im Grunde ein weiteres Pseudo Interface, dass mit `ifconfig`
eingerichtet wird. Die Parameter des Interfaces bestimmen Quelle und Ziel der
Netflow Daten.

```
$ ifconfig pflow0 flowsrc 127.0.0.1 flowdst 127.0.0.1:3001
$ ifconfig pflow0
pflow0: flags=20041<UP,RUNNING,NOINET6> mtu 1492
        priority: 0
        pflow: sender: 127.0.0.1 receiver: 127.0.0.1:3001 version: 5
        groups: pflow
```

Aber ohne Infos aus `pf` hilft auch das schönste NetFlow Interface nichts. Wie
bei `pflog` werden entweder bestimmte Rules via Parameter dazu bewogen die State
Informationen an das pflow0 Interface zu schicken oder eben gleich alle States
per default.

```
pass in on $extif proto tcp from any to any port 9001 keep state (pflow)
# or
set state-defaults pflow
```

Um die Daten dann auszulesen, lässt man das das mitgebrachte Tool `flowd-reader`
auf das flowd Logfile los.

```
$ flowd-reader /var/log/flowd
FLOW recv_time 2015-03-16T10:28:07.883062 proto 6 tcpflags 00 tos 00 agent [127.0.0.1] src [188.226.189.53]:60608 dst [185.34.0.188]:9001 packets 12 octets 2862
FLOW recv_time 2015-03-16T10:28:07.883062 proto 6 tcpflags 00 tos 00 agent [127.0.0.1] src [185.34.0.188]:9001 dst [188.226.189.53]:60608 packets 8 octets 2727
FLOW recv_time 2015-03-16T10:28:45.982868 proto 6 tcpflags 00 tos 00 agent [127.0.0.1] src [93.115.94.243]:10361 dst [185.34.0.188]:9001 packets 10 octets 2647
FLOW recv_time 2015-03-16T10:28:45.982868 proto 6 tcpflags 00 tos 00 agent [127.0.0.1] src [185.34.0.188]:9001 dst [93.115.94.243]:10361 packets 8 octets 2631
```

Hier können mit Filter usw. spezifischer Traffic im definierten Output Format
ausgelesen werden. Ziemlich nice.

### pf labels

Die letzte (und mir sympathischste) Variante sind `label`s.

{{< figure src="/uploads/2015/03/labels.png" >}}

Hinter einer bestimmten Rule in der `pf.conf` kann entsprechendes Label
angebracht werden.

```
pass in on $extif proto tcp from any to any port 9001 label tor
```

Nun werden Bytes(in/out), Packets(in/out), Statechanges usw. die auf dieser Regel
matchen erfasst. Reihenfolge am Besten in der Manpage nachlesen.

```
$ pfctl -sl
ssh 22115 247697 74839158 130093 9043197 117604 65795961 1467
tor 17813 1329213 988230585 603981 291521453 725232 696709132 9923
monitoring 17813 54035 11840438 27936 5725257 26099 6115181 3145
discard 17813 4 200 2 120 2 80 2
```

Da die meisten Rulesets mit Tables für die Ports versehen sind, ist die Lösung
mit dem statischen String als Label aber ungeeignet. Dafür gibts auch eine
Lösung, da im Labelstring auch ein Set aus Variablen verwendet werden kann.

{% blockquote %}
The following macros can be used in labels:
$dstaddr     The destination IP address.
$dstport     The destination port specification.
$if          The interface.
$nr          The rule number.
$proto       The protocol name.
$srcaddr     The source IP address.
$srcport     The source port specification.
{% endblockquote %}

Die `pf.conf` könnte dann wie folgt aussehen.

```
# Port tables
tcpin = "{ ssh, smtp, domain, www, https, ftp, ftp-data }"
udpin = "{ 9, domain }"
# Incoming rule
pass in on $extif proto tcp from any to any port $tcpin label "tcp:in:$dstport"
pass in on $extif proto udp from any to any port $udpin label "udp:in:$dstport"
```

und der dynamisch generierte Output:

```
tcp:out:80 151 2099 1121133 1105 996878 994 124255 6
tcp:out:443 151 17524 14295556 9411 13062701 8113 1232855 35
```

### Fazit

Was mir tatsächlich fehlt ist etwas das dynamische Auswerten von Source IP und
verbrauchtem Traffic. Was `pf` zwar beherrscht ist einer bestimmten Regel eine
Bandwith (über `queues`) zu Garantieren/Drosseln aber tatsächlich konsumierte
Bandbreite limitieren funktioniert nicht so ohne weiteres.

Dazu müsste man sich schon etwas eigenes Basteln, das Traffic zählt und dann
darauf reagiert. Erschwerend kommt hinzu, dass sich bei tcpdump kein
tatsächliches Outputformat definieren lässt. Zum Beispiel nur SRC IP und Bytes.
Wenn man danach im Netz sucht, existieren zwar allerhand wilde sed und awk
konstrukte, aber nichts auf das ich mich wirklich verlassen möchte.

Also gibts das erstmal nicht.
