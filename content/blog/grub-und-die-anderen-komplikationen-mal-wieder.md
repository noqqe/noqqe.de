---
aliases:
- /archives/6
date: '2008-03-24T18:46:31'
slug: grub-und-die-anderen-komplikationen-mal-wieder
tags:
- hardware
- knoppix
- linux
title: Grub und andere Komplikationen...
---

Dank des genialen Einfalls eines Freundes Windows XP auf dem Linux-PC zu
installieren, um Age of Empires 2 über LAN spielen zu können, spielt im
moment mein Grub Bootloader nicht mehr mit.

Mithilfe einer Knoppix Live CD versuche ich zur Zeit den Grub Bootloader
wieder in den Master Boot Record (MBR) zu bekommen, um mein Ubuntu wieder
booten zu können!

Meine Versuche sahen, unter anderem, wie folgt aus:

```
root@Knoppix:~# grub-install --root-directory=/media/sda1 /dev/sda
rm: Entfernen von >>/media/sda1/boot/grub/stage1<< nicht möglich: Das Dateisystem ist nur lesbar
```

Aus irgendwelchen Grüden stellt es sich als problematisch dar, die Rechte
für die sda1 Partition zu ändern..

```
root@Knoppix:~# grub-install --recheck /media/sda
```

diverse Fehlende/Falsche Parameter, mit denen ich mich noch nicht so
auskenne. Aber das werde ich noch ändern ;)

Nach 3 Stunden Probiererei und mal wieder am Rande des Wahnsinns bleibt mir
nichts anderes als im Ubuntuusers.de Forum nach Hilfe zu fragen.
gna ...
