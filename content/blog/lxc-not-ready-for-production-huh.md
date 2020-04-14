---
comments:
- author: Asdf
  content: "<p>\_Das ist EXAKT das was ich brauche :-D Mal ausprobieren...</p>"
  date: '2012-05-14T15:50:53'
- author: noqqe
  content: "<p>\_Sau gut :) </p>"
  date: '2012-05-14T17:36:15'
- author: Silvio Knizek
  content: "<p>Ich wei\xDF nicht\u2026 ich werde erst auf lxc umstellen, wenn es so
    einfach wie vServer geworden ist. </p>"
  date: '2012-05-15T19:07:17'
- author: noqqe
  content: "<p>\_Was meinst du mit vServer genau? Xen, vmware ? </p>"
  date: '2012-05-16T09:45:05'
- author: Anonym
  content: "<p>Schau dir mal bitte deinen Post im Ubuntuusers-Planet an, sieht furchtbar
    aus. W\xE4re gut, wenn du mal schaust, dass du das fixt \u2026</p>"
  date: '2012-05-17T09:50:09'
- author: noqqe
  content: "<p>\_Ich habe an dieser Stelle nichts zu fixen. </p><p>Die XML die Ubuntuusers
    Planeten Software auswertet ist: <a href=\"http://noqqe.de/ubuntuusers.xml\" rel=\"nofollow\">http://noqqe.de/ubuntuusers.xml</a>.
    Die ganzen Formatzeichen sind dort nicht im Code. W\xFCrde die Planetensoftware
    das alles so auswerten w\xFCrde es wie in allen Readern funktionieren. </p><p>Habe
    viele Reader getestet: Liferea, tinytinyrss und Goolge Reader. Bei allen verlief
    der Test problemlos. </p>"
  date: '2012-05-17T11:19:12'
- author: noqqe
  content: "<p>Ich weiss, dass das furchtbar aussieht. Wenn du meinen Feed in allen
    anderen Feedreader abbonnierst sieht aber alles ganz normal aus. </p><p>Das soll
    heissen: Ich kann nichts daf\xFCr. Aus dem Feed kommt valides HTML und damit w\xE4re
    mein teil erledigt. Was das Inyoka von <a href=\"http://ubuntuusers.de\" rel=\"nofollow\">ubuntuusers.de</a>
    damit macht ist f\xFCr mich nicht \"fixbar\". </p><p>Ich habe allerdings geh\xF6rt
    das es dazu schon einen Bug im Bugtracker von Inyoka gibt.\_ </p>"
  date: '2012-05-28T12:46:51'
- author: Foo
  content: <p>Arsch! :)</p>
  date: '2012-05-30T22:32:13'
- author: Thorsten Reiser
  content: <p>Die Installer aus der 0.8 sind genial aber die Probleme mit der cgroup
    failed to rename cgroup /sys/fs/cgroup// sind immer noch da ..... wirklich schade</p>
  date: '2013-03-18T00:07:11'
date: '2012-05-14T13:30:00'
tags:
- development
- bridge
- lan
- container
- thinkpad
- vm
- wireless
- virtualisierung
- lxc
- linux
- wlan
- eth
- debian
title: LXC. Not ready yet for production use, huh?
---

Von [LinuX Containern](http://lxc.sourceforge.net) hab ich zum ersten Mal auf den
Linux Tagen im [Vortrag](http://chemnitzer.linux-tage.de/2012/vortraege/1035)
von [Erkan Yanar](http://linsenraum.de/erkules) gehört.

Kurz gesagt ist LXC größtenteils eine Sammlung von Skripten welche in
soetwas wie einem `chroot` einen neuen init-Prozess spawnen. Das bietet mit
cgroups in eigenen Namespaces auch eine Art dynamische Ressourcen Zuweisung.

{{< figure src="/uploads/2012/05/xzibit-virtualisierung.jpg" >}}

Dafür gibts jetzt verschiedene Use-Cases. Bei mir ist das sowas wie eine Art
Sandboxing für Spielereien auf meinem Thinkpad.

> I said before that in my opinion LXC is not ready yet for production use, and I
> maintain that opinion today. I would also rephrase it in something that might
> make it easier to understand what I think: I would never trust root on a
> container to somebody I wouldn’t trust root with on the host.

Das Zitat triffts eigentlich recht gut auf den Punkt. Besonders gefährlich ist
das gemeinsam genutzte /dev und /proc Verzeichnis, wobei es unter Umständen auch
einfach möglich ist die gesamte Maschine herunterzufahren. Deshalb würde ich
davon abraten virtuelle Maschinen an Dritte abzugeben.

## LXC: Debian Stable oder Testing?

Trotz allem reizt mich das Thema. Installiert wird das wie so ziemlich alles
über

``` bash
apt-get install lxc bridge-utils debootstrap
```

Allerdings gibt es bei testing und auch bei stable Nachteile.

```
$ apt-cache policy lxc
lxc:
  Installiert: 0.7.2-1
  Kandidat:    0.7.2-1
  Versionstabelle:
     0.8.0~rc1-4 0
        500 ftp://ftp.de.debian.org/debian/ testing/main amd64 Packages
 *** 0.7.2-1 0
        990 http://ftp.uni-erlangen.de/debian/ squeeze/main amd64 Packages
        100 /var/lib/dpkg/status
```

Testing: Im Moment scheinen die cgroups kaputt. Dafür konnte ich keine Lösung
finden da ich auch noch keine Erfahrung mit cgroups habe. Schade. Verison 0.8 hätte mich gereizt.

```
$ lxc-start --name vm0
lxc-start: No such file or directory - failed to rename cgroup /sys/fs/cgroup//14051->/sys/fs/cgroup//vm0/14051
lxc-start: failed to spawn 'vm0'
lxc-start: No such file or directory - failed to remove cgroup '/sys/fs/cgroup//vm0/14051'
```

Stable: Unter Stable ist kein Template für Debian Squeeze verfügbar. Da
heissts dann Lenny Template [umbauen](http://jtrancas.wordpress.com/2011/02/10/debian-squeeze-lxc-template/).
Anleitung ist recht einfach zu verstehen und funktioniert.

## Networking über WLAN

Das mit dem normalen Bridiging über WLAN Interfaces haut erfahrungsgemäß nicht
hin. Deshalb hab ich mir [hier](http://s3hh.wordpress.com/2011/05/17/lxc-containers-on-a-host-with-wireless/)
etwas helfen lassen. Meine Config sieht und in etwa so aus:

```
auto lxcbr0
iface lxcbr0 inet static
  address 10.10.0.1
  netmask 255.255.255.0
  pre-up brctl addbr lxcbr0
  post-up /usr/local/bin/lxcbr0-up
  post-down brctl delbr lxcbr0
```

```
$ cat /usr/local/bin/lxcbr0-up
#!/bin/sh
# This is the address we assigned to our bridge in /etc/network/interfaces
braddr=10.10.0.1
# ip address range for containers
brrange=10.10.0.10,10.10.0.100
iptables -A FORWARD -i lxcbr0 -s ${braddr}/24 -m conntrack --ctstate NEW -j ACCEPT
iptables -A FORWARD -m conntrack --ctstate ESTABLISHED,RELATED -j ACCEPT
iptables -A POSTROUTING -t nat -j MASQUERADE
dnsmasq --bind-interfaces --conf-file= --listen-address $braddr --except-interface lo --dhcp-range $brrange --dhcp-lease-max=253 --dhcp-no-override
```

Wenn man das ohne DHCP möchte geht das auch einfach so - ohne Umkonfiguration :)

## Mass-LXC

Ich möchte meine Linux Container nun immer in einem `screen` starten. Ein Start
sieht daher ungefähr immer so aus:

```
screen -d -m -S vm0 lxc-start -n vm0 -f /var/lib/lxc/vm0/config
```

Aus erst einem Alias wurden dann Mehrere und dann ein Wrapper Script
mit dem ich Infos und Start/Stops für alle VMs auslesen/anweisen kann.
Ich muss zugeben es ist etwas ausgeartet.
Das Script ist wie immer auf [Github](https://gist.github.com/2693967) zu haben. Unter Umständen kann
nochjemand etwas damit anfangen.

```
$ mlxc info
Screens:
No Sockets found in /var/run/screen/S-root.

LXC Status:
'vm0' is STOPPED
'vm1' is STOPPED
'vm2-mysqld' is STOPPED
'vm3-postgresql' is STOPPED
'vm4-matomat' is STOPPED

$ mlxc all
$ mlxc info
Screens:
There are screens on:
  8682.vm0  (14.05.2012 15:29:47) (Detached)
  8686.vm1  (14.05.2012 15:29:47) (Detached)
  8690.vm2-mysqld (14.05.2012 15:29:47) (Detached)
  8694.vm3-postgresql (14.05.2012 15:29:47) (Detached)
  8698.vm4-matomat  (14.05.2012 15:29:47) (Detached)
5 Sockets in /var/run/screen/S-root.

LXC Status:
'vm0' is RUNNING
'vm1' is RUNNING
'vm2-mysqld' is RUNNING
'vm3-postgresql' is RUNNING
'vm4-matomat' is RUNNING

$ mlxc net
Netstat for vm0:
IP: 10.10.0.10
Aktive Internetverbindungen (Server und stehende Verbindungen)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name
tcp        0      0 0.0.0.0:22              0.0.0.0:*               LISTEN      10815/sshd
tcp        0      0 127.0.0.1:25            0.0.0.0:*               LISTEN      10728/master
[...]
```

Auf Ideen, Anmerkungen, Kritik und Beleidigungen freu ich mich natürlich immer.
