---
date: 2008-03-26T11:46:52+02:00
slug: des-ratsels-losung
comments: true
title: Des Rätsels Lösung?
aliases:
- /archives/8
- /blog/2008/03/26/des-r%C3%A4tsels-l%C3%B6sung/
categories:
- Hardware
- Linux
---

Moin Moin!

Dem Wiki-Lösungsweg ("Grub-Methode 3" aus ubuntuusers.de) bin ich einfach
mal blind gefolgt ;)

Ob dieser Pfad mich mein Ubuntu wieder Booten lässt weis ich noch nicht.

Gut durchdacht und auf mein Problem angepasst sieht dieser mal wie folgt
aus:

```
knoppix@Knoppix:~$ sudo mount --bind /dev /media/sda1/dev
```

Sollte  Systeminformationen übertragen.

```
knoppix@Knoppix:~$ sudo mount -t proc /proc /media/sda1/proc
```

Sollte "proc"s übertragen.

```
knoppix@Knoppix:~$ sudo chroot /media/sda1
```

Wechselt ins die Partition sda1 (Ubuntu-Partiton)

```
root@Knoppix:/# cp /proc/mounts /etc/mtab
```

Übernimmt irgendwelche Einstellungen  ^^

```
root@Knoppix:/# grub-install /dev/sda
```

Installiert Grub auf Ubuntu Partiton Sollte gehen und danke enchbladexp!
Super Helfer aus ubuntuusers.de.
