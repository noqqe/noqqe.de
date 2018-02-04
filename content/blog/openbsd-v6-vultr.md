---
title: "OpenBSDs IPv6 on Vultr"
date: 2018-02-04T16:15:05+01:00
tags:
- vultr
- IPv6
- OpenBSD
---

Es ist ein [nie](/blog/2014/10/09/openbsd-ipv6/)
[enden](/blog/2014/12/18/router-advertisments-in-openbsd/)
[wollendes](/blog/2016/06/26/ipv6-fuer-hue-bridge/)
[Thema](/blog/2008/12/17/ipv6/) dieses Blogs zu erklären wie man IPv6
konfiguriert. Auf [Vultr](https://vultr.com) läuft so die eine oder andere
meiner Maschinen und das bisher ohne v6. Die Konfiguration verlief nicht so
reibungslos wie die letzten male, weshalb dieser Post viel mehr Details enthält
und ich nun viel besser verstehe was eigentlich passiert!

In den Settings zeigt einem das Webinterface von Vultr die
v6 Adresse die man für seine VM benutzen kann daneben der Hinweis:

> Gateway: (use router discovery)

Cool, dachte ich. Das kenne ich ja schon. In den Beispielen wird für
OpenBSD aber noch die bereits deaktivierte Option für `hostname.if`
`rtsol` empfohlen.

{{< figure src="/uploads/2018/02/vultr.png" >}}

Das funktioniert schonmal nicht, aber gut. Hab den Support zumindest mal
darauf hingewiesen. Mal sehen ob sie antworten.

Wie man eine statische IP und Router Discovery unter OpenBSD konfiguriert
hier der Vollständigkeit halber nochmal.

```
$ cat /etc/hostname.vio0
dhcp
inet6 2001:19f0:6c01:db:5400:00ff:fe69:d0d4 64
inet6 autoconf
```

Die Konfiguration mittels `/etc/netstart vio0` angewendet. Passiert
nichts. `route -n show` ausgeführt, aber in dem Wust kein Default Gateway
für v6 finden können. Was nun. Wie läuft das eigentlich genau? Darum geht
es hier in diesem Post.

Ich fing etwas an zu kramen und fand heraus, dass `slaacd(8)` die
Konfiguration meines Default Gateways zuständig ist. Ein kurzes `ps aux`
bestätigte auch das dieser Daemon läuft. Den Daemon neu gestartet, `rcctl restart
slaacd`. Weiterhin ließ mich `ping6 heise.de` enttäuscht zurück.

Ich begann die Manpage zu lesen.

> slaacd(8)
>
> It listens for IPv6 router advertisement messages on interfaces
> with the AUTOCONF6 flag set.

Das heisst... das die `int6 autoconf` Option das ich oben in
`/etc/hostname.vio0` setzte nichts anderes tut als...

```
# ifconfig vio0
vio0: flags=208843<UP,BROADCAST,RUNNING,SIMPLEX,MULTICAST,AUTOCONF6> mtu 1500
```

Genau! `ifconfig` setzt einfach nur das Flag am Interface, damit ein
Daemon es später wieder finden kann. Wieder so ein "Aha, geil!" Moment im
Umgang mit OpenBSD. Am Ende der Manpage ist auch das Client Programm
referenziert mit dem ich mit `slaacd` kommunizieren kann: `slaacctl(8)`.

```
$ slaacctl show interface vio0
vio0:
         index:   1 running: yes privacy: yes
        lladdr: 56:00:00:69:8c:bf
         inet6: fe80::5400:ff:fe69:8cbf%vio0
```

Das sieht erstmal nicht so aus als wäre die Router Discovery erfolgreich
gewesen. Nur die poplige `fe80` Adresse. Irgendwas fehlte. Liegts an mir?
Oder vielleicht ist einfach Infrastruktur von Vultr genau so fehlerhaft
wie die Doku? Als ich dann nochmal die Manpage öffnete und "It listens [...]"
laß klingelte nochmal was im Hinterkopf. Firewall! IPv6 nutzt das
[ICMPv6](https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol_version_6)
Protokoll um Neighbor Discovery und Router Advertisements zu verschicken.
Meine `pf` Settings lehnen aber all das erstmal ab.

```
# Allow receiving ipv6 types
icmp6types = "{128, 133, 134, 135, 136, 137}"
pass in quick on $extif inet6 proto ipv6-icmp icmp6-type $icmp6types keep state
```

Aber das lässt sich natürlich ändern. Um via `pf` eine Hand voll ICMPv6
Types zuzulassen, kann man wie oben gezeigt vorgehen. Welche Codes man
genau braucht einfach in Wikipedia/RFC nachlesen.

Aber jetzt...

```
$ pfctl -f /etc/pf.conf
$ sh /etc/netstart vio0
$ rcctl restart slaacd
```

und zack, hab ich ein funktionierendes v6.


```
# slaacctl show interface vio0
vio0:
         index:   1 running: yes privacy: yes
        lladdr: 56:00:00:6e:5a:b9
         inet6: fe80::5400:ff:fe6e:5ab9%vio0
        Router Advertisement from fe80::fc00:ff:fe6e:5ab9%vio0
                received: 2018-02-04 16:11:41; 342s ago
                Cur Hop Limit:  64, M: 0, O: 0, Router Lifetime:  1800s
                Default Router Preference: Medium
                Reachable Time:         0ms, Retrans Timer:         0ms
                prefix: 2001:19f0:6c01:232::/64
                        On-link: 1, Autonomous address-configuration: 1
                        vltime:    2592000, pltime:     604800
                rdns: 2001:19f0:300:1704::6, lifetime: 3600
        Address proposals
                id:    2, state:      CONFIGURED, privacy: y
                vltime:     589027, pltime:      70493, timeout:      68838s
                updated: 2018-02-04 15:50:04; 1640s ago
                2001:19f0:6c01:232:782f:b082:b044:b5ef, 2001:19f0:6c01:232::/64
                id:    1, state:      CONFIGURED, privacy: n
                vltime:    2592000, pltime:     604800, timeout:     604443s
                updated: 2018-02-04 16:11:41; 342s ago
                2001:19f0:6c01:232:5400:ff:fe6e:5ab9, 2001:19f0:6c01:232::/64
        Default router proposals
                id:    3, state:      CONFIGURED
                router: fe80::fc00:ff:fe6e:5ab9%vio0
                router lifetime:       1800
                Preference: Medium
                updated: 2018-02-04 16:11:41; 342s ago, timeout:       1443s
```


Und sogar `ping6 heise.de` tickt fröhlich vor sich hin.

Warum schreibe ich das alles? Es war ein sehr großer Spass für mich zu
"entdecken" wie das eigentlich funktioniert und demonstriert einmal mehr
wie einfach Dinge sein können wenn man das richtige OS benutzt. Etwas das
man sich auch nur leisten kann, wenn
Kernel und Userland vom selben Team entwickelt werden.
