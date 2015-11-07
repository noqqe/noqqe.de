---
categories:
- administration
- bsd
- openbsd
- opensource
- osbn
- planetenblogger
- ubuntuusers
date: 2015-11-07T21:59:22+01:00
tags:
- null
title: OpenBSD Upgrade Guidelines
---

OpenBSD ist eine der am Besten dokumentierten Distributionen überhaupt.
Nichts, aber auch garnichts, führt an der offiziellen Upgrade Dokumentation
vorbei, da diese jedes mal auch ein wenig anders ist.

{{< figure src="/uploads/2015/11/upgrade.jpg" >}}

Trotz allem Pflege ich so meine eigene kleine Anleitung in `cmddocs` um die
vier Kisten unter meiner "Obhut" unfallfrei durch das Upgrade zu führen.
Vor allem da nicht alle Aspekte in der offiziellen Anleitung ausgeführt
werden. Seit 6 Releases mach ich das jetzt so.

### Pre-Upgrade

* Upgrade Manual auf http://www.openbsd.org/faq/upgrade5X.html genauestens durchlesen
* Downtime im Monitoring eintragen
* root Login direkt über sshd sicherstellen
* Sicherstellen dass root eine /bin/ksh zugewiesen hat
* Verifizieren dass die letzten Backups liefen

### Upgrade

* Von einem Host dediziert als root user einloggen, pur, nicht über sudo
* Sicherstellen ob der Mirror, den man benutzt schon gesynct ist.
* Base Sets und Kernel und Signaturen ziehen.

```
VER=5.8
ARCH=$(uname -m)
wget -r --no-parent -A.tgz http://openbsd.cs.fau.de/pub/OpenBSD/$VER/$ARCH/
cd openbsd.cs.fau.de/pub/OpenBSD/$VER/$ARCH/
wget http://openbsd.cs.fau.de/pub/OpenBSD/$VER/$ARCH/bsd
wget http://openbsd.cs.fau.de/pub/OpenBSD/$VER/$ARCH/bsd.rd
wget http://openbsd.cs.fau.de/pub/OpenBSD/$VER/$ARCH/bsd.mp
wget http://openbsd.cs.fau.de/pub/OpenBSD/$VER/$ARCH/bsd.sp
wget http://openbsd.cs.fau.de/pub/OpenBSD/$VER/$ARCH/SHA256
wget http://openbsd.cs.fau.de/pub/OpenBSD/$VER/$ARCH/SHA256.sig
```

* Schritten aus der Anleitung zum Einspielen auf openbsd.org folgen
* Reboot
* Auf Package- / Base-Changes im Tutorial achten.

### Packages und Ports

* Packages: PKG_PATH in ksh/bash aktualisieren und updaten

		export PKG_PATH=http://openbsd.cs.fau.de/pub/OpenBSD/$(uname -r)/packages/$(uname -m)/
		pkg_add -u

* CVS src Tree updaten

		cd /usr
		cvs -qd anoncvs@openbsd.cs.fau.de:/cvs get -rOPENBSD_5_X -P src

* CVS ports Tree updaten

		cd /usr
		cvs -qd anoncvs@openbsd.cs.fau.de:/cvs get -rOPENBSD_5_X -P ports

### Errata

Da oftmals Errata publik werden nachdem das eigentliche Release schon im
freeze ist, am besten gleich nach dem Upgrade schauen, ob schon Errata
verfügbar sind. Ich benutze hierfür Binpatching via
[openup](http://www.mtier.org/index.php/solutions/apps/openup/)


    PKG_PATH_MAIN=http://openbsd.cs.fau.de/pub/OpenBSD/5.X/packages/amd64/
    ./openup

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
