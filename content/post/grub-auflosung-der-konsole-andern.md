---
date: 2010-07-14T19:17:18+02:00
type: post
slug: grub-auflosung-der-konsole-andern
comments: true
title: Grub | Auflösung der Konsole ändern
aliases:
- /archives/1098
categories:
- Development
- Debian
- Linux
- Ubuntu
tags:
- ändern
- auflösung
- boot
- change
- grub
- grub 1
- grub 2
- hex
- resolution
- table
---

Die Auflösung der (ich nenne es mal so) Boot-Konsole ist nicht gerade die Höchste. Gerade im Recovery Mode oder bei anderem stört (mich persönlich) das immer etwas.
Lösung gefunden und damit ich es nicht vergesse, nun hier:

[George Notaras](http://www.g-loaded.eu/2005/09/30/change-the-console-resolution/) hat für **grub 1 und 2(!)** folgende wunderschöne Tabelle gebastelt.


         | 640x480  800x600  1024x768 1280x1024
    ----+-------------------------------------
    256 |  0x301   0x303    0x305    0x307
    32k |  0x310   0x313    0x316    0x319
    64k |  0x311   0x314    0x317    0x31A
    16M |  0x312   0x315    0x318    0x31B


Dieser Hex-Wert muss als zusätzlicher Parameter in **/boot/grub/menu.lst** an den Kernel angehängt werden. Folgendermaßen kann sowas aussehen:
```
kernel /boot/vmlinuz-2.6.26-2-686 root=/dev/ida/c0d0p1 ro vga=0x318 quiet
```


Aktuell für Grub 2: **/boot/grub/grub.cfg**. Sieht bisschen anders aus, funktioniert aber genauso:
```
linux   /boot/vmlinuz-2.6.32-22-generic root=UUID=92892dbf-af24-4dbd-b2a4-8debdbb08981 ro  vga=0x318 quiet splash
```


Sollte es eventuell noch eine schönere Möglichkeit für Grub 2 geben (was ich mir durchaus vorstellen kann) bitte ich diese doch kurz zu kommentieren :)
