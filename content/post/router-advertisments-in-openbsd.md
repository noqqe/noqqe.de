---
layout: post
title: "Router Advertisments in OpenBSD"
date: 2014-12-18T11:14:00+02:00
comments: true
categories:
- OpenBSD
- IPv6
- Routeradvertisement
- rtadv
- sysctl
- osbn
---
Seit OpenBSD 5.6 hat sich die Konfiguration für IPv6 Router Advertisements
geändert.

{{< figure src="/uploads/2014/12/suprise.gif" >}}

Bisher hat man über `sysctl` ein globales Flag eingeschaltet, das die
automatische Konfiguration von IPv6 Adressen erlaubt. Das Flag enabled die
Option aber automatisch für alle Interfaces. Deshalb wurde es entfernt und durch
ein Flag in `ifconfig` ersetzt.

```
$ sysctl net.inet6.ip6.accept_rtadv=1
sysctl: fourth level name accept_rtadv in net.inet6.ip6.accept_rtadv is invalid
```

Die neue Konfiguration wird nun in `/etc/hostname.em0` angezogen.

```
$ cat hostname.em0
dhcp
inet6 autoconf
```

Ist also pro Interface einzeln konfigurierbar. Bin darüber gestolpert, da mir
das beim (etwas dürftig formulierten) Changelog nicht aufgefallen ist. Nach
etwas Mailinglisten stöbern
[hier](http://marc.info/?l=openbsd-tech&m=140632650926622&w=2) und
[da](http://marc.info/?l=openbsd-tech&m=140632650926622&w=2) konnt ichs aber
finden.

Danach noch ein `sh /etc/netstart em0` und IPv6 geht wieder.
