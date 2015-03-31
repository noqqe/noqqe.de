---
layout: post
title: "LXC. Not ready yet for production use, huh?"
date: 2012-05-14T15:30:00+02:00
comments: true
categories:
- ubuntuusers
- Linux
- Debian
- Web
- Coding
- Virtualisierung
keywords: "LXC, lxc, linux, linux container, container, thinkpad, wireless,
bridge, wlan, lan, eth, vm, virtualisierung"
---

Von [LinuX Containern](http://lxc.sourceforge.net) hab ich zum ersten Mal auf den
Linux Tagen im [Vortrag](http://chemnitzer.linux-tage.de/2012/vortraege/1035)
von [Erkan Yanar](http://linsenraum.de/erkules) gehört.

Kurz gesagt ist LXC größtenteils eine Sammlung von Skripten welche in
soetwas wie einem `chroot` einen neuen init-Prozess spawnen. Das bietet mit
cgroups in eigenen Namespaces auch eine Art dynamische Ressourcen Zuweisung.

{% img center /uploads/2012/05/xzibit-virtualisierung.jpg %}

Dafür gibts jetzt verschiedene Use-Cases. Bei mir ist das sowas wie eine Art
Sandboxing für Spielereien auf meinem Thinkpad.

{% blockquote Diego Elio “Flameeyes” Pettenò, http://blog.flameeyes.eu/2010/06/lxc-and-why-it-s-not-prime-time-yet blog.flameeyes.eu %}
I said before that in my opinion LXC is not ready yet for production use, and I
maintain that opinion today. I would also rephrase it in something that might
make it easier to understand what I think: I would never trust root on a
container to somebody I wouldn’t trust root with on the host.
{% endblockquote %}

Das Zitat triffts eigentlich recht gut auf den Punkt. Besonders gefährlich ist
das gemeinsam genutzte /dev und /proc Verzeichnis, wobei es unter Umständen auch
einfach möglich ist die gesamte Maschine herunterzufahren. Deshalb würde ich
davon abraten virtuelle Maschinen an Dritte abzugeben.

## LXC: Debian Stable oder Testing?

Trotz allem reizt mich das Thema. Installiert wird das wie so ziemlich alles
über

{% codeblock %}
$ apt-get install lxc bridge-utils debootstrap
{% endcodeblock %}

Allerdings gibt es bei testing und auch bei stable Nachteile.

{% codeblock %}
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
{% endcodeblock %}


* Testing: Im Moment scheinen die cgroups kaputt. Dafür konnte ich keine Lösung
finden da ich auch noch keine Erfahrung mit cgroups habe. Schade. Verison 0.8 hätte mich gereizt.

{% codeblock %}
$ lxc-start --name vm0
lxc-start: No such file or directory - failed to rename cgroup /sys/fs/cgroup//14051->/sys/fs/cgroup//vm0/14051
lxc-start: failed to spawn 'vm0'
lxc-start: No such file or directory - failed to remove cgroup '/sys/fs/cgroup//vm0/14051'
{% endcodeblock %}

* Stable: Unter Stable ist kein Template für Debian Squeeze verfügbar. Da
heissts dann Lenny Template [umbauen](http://jtrancas.wordpress.com/2011/02/10/debian-squeeze-lxc-template/).
Anleitung ist recht einfach zu verstehen und funktioniert.

## Networking über WLAN

Das mit dem normalen Bridiging über WLAN Interfaces haut erfahrungsgemäß nicht
hin. Deshalb hab ich mir [hier](http://s3hh.wordpress.com/2011/05/17/lxc-containers-on-a-host-with-wireless/)
etwas helfen lassen. Meine Config sieht und in etwa so aus:

{% codeblock %}
auto lxcbr0
iface lxcbr0 inet static
  address 10.10.0.1
  netmask 255.255.255.0
  pre-up brctl addbr lxcbr0
  post-up /usr/local/bin/lxcbr0-up
  post-down brctl delbr lxcbr0
{% endcodeblock %}


{% codeblock %}
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
{% endcodeblock %}

Wenn man das ohne DHCP möchte geht das auch einfach so - ohne Umkonfiguration :)

## Mass-LXC

Ich möchte meine Linux Container nun immer in einem `screen` starten. Ein Start
sieht daher ungefähr immer so aus:

{% codeblock %}
$ screen -d -m -S vm0 lxc-start -n vm0 -f /var/lib/lxc/vm0/config
{% endcodeblock %}

Aus erst einem Alias wurden dann Mehrere und dann ein Wrapper Script
mit dem ich Infos und Start/Stops für alle VMs auslesen/anweisen kann.
Ich muss zugeben es ist etwas ausgeartet.
Das Script ist wie immer auf [Github](https://gist.github.com/2693967) zu haben. Unter Umständen kann
nochjemand etwas damit anfangen.

{% codeblock %}
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
{% endcodeblock %}

Auf Ideen, Anmerkungen, Kritik und Beleidigungen freu ich mich natürlich immer.

