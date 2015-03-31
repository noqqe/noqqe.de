---
layout: post
title: "OpenBSD IPv6"
date: 2014-10-09 19:15
comments: true
categories:
- OpenBSD
- IPv6
- netstart
- hostname
- ip
- ping6
- ubuntuusers
- osbn
---

Was ich ja schon immer komisch fand ist, dass die Dokumentation oder auch uch How-Tos im Netz was das Thema OpenBSD und IPv6
angeht echt ultra duerftig ist.

{% img center /uploads/2014/10/tcpip.jpg %}

Da ich meine Debian VM auf der ich
weechat, mutt, jabber usw laufen habe vor nem Monat ebenfalls auf OpenBSD
umgezogen hab, stellte sich mir das Problem schon wieder.

### IPv6 mit Router Adverisement (rtadv)

Am komfortablesten ist natuerlich einfach das rtadv den Hosting Providers
herzunehmen. In meinem Fall [rootbsd.net](http://rootbsd.net) haben aber nur
kleine Anleitungen fuer FreeBSD.

Um den IP Stack auf IPv6 Advertisements antworten zu lassen muss nur
`/etc/sysctl.conf` editiert werden mit

    net.inet6.ip6.accept_rtadv=1

Einmal rebooten oder per Hand `sysctl net.inet6.ip6.accept_rtadv=1` ausfuehren.

{% codeblock %}
$ ifconfig
em0: flags=8843<UP,BROADCAST,RUNNING,SIMPLEX,MULTICAST> mtu 1500
        lladdr 00:16:3e:2c:4a:41
        priority: 0
        groups: egress
        media: Ethernet autoselect (1000baseT full-duplex)
        status: active
        inet6 fe80::216:3eff:fe2c:4a41%em0 prefixlen 64 scopeid 0x1
        inet 12.34.56.78 netmask 0xffffff80 broadcast 185.34.0.255
        inet6 2a00:d1e0:1000:3100:dead:beef:4a41 prefixlen 64 autoconf pltime 604729 vltime 2591929
        inet6 2a00:d1e0:1000:3100:dead:beef:df24 prefixlen 64 autoconf autoconfprivacy pltime 18185 vltime 537125
{% endcodeblock %}

### IPv6 mit statischer IP

Wenn der Provider kein Router Advertisment anbietet, aber dafür eine IP assigned und auf diese IP
ein eigenens Netz für einen routet.

Konfigurieren der Adressen:

{% codeblock %}
$ vi /etc/hostname.em0
inet 213.95.21.200 255.255.255.0 NONE
inet6 alias 2001:780:3:5::122 64   # Transit IP
inet6 alias 2001:780:132::1 48     # IP aus eigenem Netz
inet6 alias 2001:780:132::2 48     # IP aus eigenem Netz
inet6 alias 2001:780:132::3 48     # IP aus eigenem Netz
{% endcodeblock %}

Konfiguration des Gateways

{% codeblock %}
vi /etc/mygate
213.95.21.1
2001:780:3:5::1
{% endcodeblock %}

Und danach Interface reloaden mit `sh /etc/netstart em0`.
