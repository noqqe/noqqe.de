---
comments:
- author: tux.
  content: OpenBSD ist keine Distribution.
  date: '2015-11-07T22:20:42.061276'
- author: Anonymous
  content: Pinguine hier mal den den Kopf zu machen, wenn der Teufel an der Arbeit
    ist.
  date: '2015-11-07T23:38:19.068166'
- author: noqqe
  content: Die auf der ... Berkley Software Distribution basiert? Irgendwie doch.
  date: '2015-11-07T23:39:55.567497'
- author: tux.
  content: "Ich nutze BSD schon 'n bisschen l\xE4nger, Troll."
  date: '2015-11-08T00:08:58.760291'
- author: tux.
  content: "Nein, OpenBSD ist l\xE4ngst ein eigenst\xE4ndiges System. Bzw. wenn OpenBSD
    eine Distribution ist, ist Windows auch eine. (1992 musste s\xE4mtlicher Original-BSD-Code
    ja ersetzt werden.)"
  date: '2015-11-08T00:09:43.690936'
date: '2015-11-07T20:59:22'
tags:
- openbsd
- administration
- opensource
- bsd
title: OpenBSD Upgrade Guidelines
---

OpenBSD ist eine der am Besten dokumentierten Betriebssysteme überhaupt.
Nichts, aber auch garnichts, führt an der offiziellen Upgrade Dokumentation
vorbei, da diese jedes mal auch ein wenig anders ist.

{{< figure src="/uploads/2015/11/upgrade.jpg" >}}

Trotz allem Pflege ich so meine eigene kleine Anleitung in `cmddocs` um die
vier Kisten unter meiner "Obhut" unfallfrei durch das Upgrade zu führen.
Vor allem da nicht alle Aspekte in der offiziellen Anleitung ausgeführt
werden. Seit 6 Releases mach ich das jetzt so.

### Pre-Upgrade

* Upgrade Manual auf [http://www.openbsd.org/faq/upgrade5X.html](http://www.openbsd.org/faq/upgrade5X.html) genauestens durchlesen
* Downtime im Monitoring eintragen
* root Login direkt über sshd sicherstellen
* Sicherstellen dass root eine /bin/ksh zugewiesen hat
* Verifizieren dass die letzten Backups liefen

### Upgrade

* Von einem Host dediziert als root user einloggen, pur, nicht über sudo
* Sicherstellen ob der Mirror, den man benutzt schon gesynct ist.
* Base Sets und Kernel und Signaturen ziehen.

``` bash
VER=6.3
ARCH=$(uname -m)
wget -r --no-parent -A.tgz https://ftp.uni-erlangen.de/openbsd/$VER/$ARCH/
cd ftp.uni-erlangen.de/openbsd/$VER/$ARCH/
wget https://ftp.uni-erlangen.de/openbsd/$VER/$ARCH/bsd
wget https://ftp.uni-erlangen.de/openbsd/$VER/$ARCH/bsd.rd
wget https://ftp.uni-erlangen.de/openbsd/$VER/$ARCH/bsd.mp
wget https://ftp.uni-erlangen.de/openbsd/$VER/$ARCH/bsd.sp
wget https://ftp.uni-erlangen.de/openbsd/$VER/$ARCH/SHA256
wget https://ftp.uni-erlangen.de/openbsd/$VER/$ARCH/SHA256.sig
```

* Schritten aus der Anleitung zum Einspielen auf openbsd.org folgen
* Reboot
* Auf Package- / Base-Changes im Tutorial achten.

### Packages und Ports

* Packages: PKG_PATH in ksh/bash aktualisieren und updaten

``` bash
export PKG_PATH=https://ftp.uni-erlangen.de/openbsd/$(uname -r)/packages/$(uname -m)/
pkg_add -u
```

* CVS src Tree updaten

``` bash
cd /usr
cvs -qd anoncvs@ftp.hostserver.de:/cvs get -rOPENBSD_6_X -P src
```

* CVS ports Tree updaten

``` bash
cd /usr
cvs -qd anoncvs@ftp.hostserver.de:/cvs get -rOPENBSD_6_X -P ports
```

### Errata

Da oftmals Errata publik werden nachdem das eigentliche Release schon im
freeze ist, am besten gleich nach dem Upgrade schauen, ob schon Errata
verfügbar sind. Seit 6.0 gibt es hierfür `syspatch`.

    syspatch

### 3rd Party Software

Sicher treffen diese Punkte nicht auf jeden zu, aber um ein Bild zu
verschaffen um was es geht:

* `pip` Packages neu installieren
* `gem` Packages neu installieren
* `go` Software neu bauen
* Selbst kompilierte Software neu bauen (taskwarrior, tarsnap)

### Post-Upgrade

* sshd root login deaktivieren
* Default Shell auf Bash zurückstellen (falls nötig)

Eventuell ist das für Neulinge oder Andere nützlich. Mir hilft es jedes mal
wieder. Und wenn es nur darum geht mir sicher zu sein, nichts vergessen zu
haben.
