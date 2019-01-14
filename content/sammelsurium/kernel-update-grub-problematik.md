---
title: Kernel Update Grub Problematik
date: 2013-05-23T11:35:33
tags: 
- Filesystems
---

Wenn beim Patchday Linux Kernel Updates vorliegen, kann es unter Umständen
zu Problemen kommen, wenn der Grub neu installiert wird. Das Update bricht
mit:

    > Running postinst hook script /usr/sbin/update-grub.
    > Generating grub.cfg ...
    > /usr/sbin/grub-probe: error: cannot find a GRUB drive for /dev/sdac1.  Check
    > your device.map.
    > User postinst hook script [/usr/sbin/update-grub] exited with value 1
    > dpkg: Fehler beim Bearbeiten von linux-image-2.6.32-33-server (--configure):
    > Unterprozess installiertes post-installation-Skript gab den Fehlerwert 1 zurück

ab. Problematik dabei ist dann, dass das Boot-Device möglicherweise auf
/dev/sdac rausrotiert. Hintergrund; Zu viele LUNs als Devices hinterlegt.
dev/sda-/dev/sdz sind vergeben und es geht wieder bei sdaa,sdab,sdac
weiter. Leider kann Grub diese längeren Devicenamen nicht als Bootdevice
nutzen.  Erster Lösungsansatz war die Nummerierung über udev-rules zu
ändern, was allerdings fehlschlägt.  Problembehebung: Ins RZ fahren.

    ls -lah /dev/sdac1
    brw-rw---- 1 root disk 8, 160 2011-09-20 18:54 /dev/sdac
    brw-rw---- 1 root disk 8, 161 2011-09-20 18:54 /dev/sdac1
    brw-rw---- 1 root disk 8, 162 2011-09-20 18:54 /dev/sdac2
    brw-rw---- 1 root disk 8, 165 2011-09-20 18:54 /dev/sdac5

Ein neues Device auf die selben Grenzen erstellen: (Devicename nach
freigewordenem Namen nach LUN Abzug auswählen)

    mknod /dev/sdb b 8 160
    mknod /dev/sdb1 b 8 161

Grub Device Map Updaten

    grub-mkdevicemap

Grub config sicher und neue generieren

    cp /boot/grub/grub.cfg /root/grub.cfg.bak
    grub-mkconfig -o /boot/grub/grub.cfg

Neue Konfiguration auf innere Logik überprüfen ;) und Rebooten.