---
aliases:
- /blog/2010/07/14/grub-auflosung-der-konsole-andern
- /blog/2010/07/14/grub--aufl%C3%B6sung-der-konsole-%C3%A4ndern/
- /archives/1098
comments:
- author: savier
  content: "<p>Hey, netter Blogeintrag! Werde das heute umgehend unter meine Xubuntu-VM
    testen :) Dann resized sich das Fenster beim Booten auch nicht st\xE4ndig.</p><p>Ich
    sag Bescheid, ob das geklappt hat.</p>"
  date: '2010-07-15T12:51:17'
- author: Tanja
  content: "<p>Besser ist es, die Kerneloptionen nicht direkt in der Kernel-Zeile
    anzuh\xE4ngen sondern sie in die daf\xFCr vorgesehenen Veriablen einzutragen und
    anschlie\xDFend das update-script auszuf\xFChren. Dann muss man n\xE4mlich nicht
    nach Kernelupdate die Optionen wieder erneut eintragen, sondern es passiert automatisch.</p><p><a
    href=\"http://wiki.ubuntuusers.de/Konsolen-Aufl%C3%B6sung#GRUB-Konfiguration\"
    rel=\"nofollow\">http://wiki.ubuntuusers.de/Konsolen-Aufl%C3%B6sung#GRUB-Konfiguration</a></p>"
  date: '2010-07-15T14:10:06'
- author: Tanja
  content: "<p>Bei grub2 sollte man au\xDFerdem den Paramater \"video=\" statt \"vga=\"
    verwenden.</p><p>Sorry f\xFCr die Fehler.. :-)</p>"
  date: '2010-07-15T14:15:35'
- author: noqqe
  content: "<p>@Tanja: Vielen Dank f\xFCr die Tipps :) Scheint mir sinnvoll zu sein.
    Denke ich werd das auch dementsprechend umbasteln. :)</p>"
  date: '2010-07-15T16:38:21'
date: '2010-07-14T17:17:18'
tags:
- grub 1
- grub
- grub 2
- hex
- linux
- boot
- ubuntu
- table
- resolution
- debian
- change
title: "Grub | Aufl\xF6sung der Konsole \xE4ndern"
---

Die Auflösung der (ich nenne es mal so) Boot-Konsole ist nicht gerade die
Höchste. Gerade im Recovery Mode oder bei anderem stört (mich persönlich)
das immer etwas.  Lösung gefunden und damit ich es nicht vergesse, nun
hier:

[George Notaras](http://www.g-loaded.eu/2005/09/30/change-the-console-resolution/)
hat für **grub 1 und 2(!)** folgende wunderschöne Tabelle gebastelt.

         | 640x480  800x600  1024x768 1280x1024
    ----+-------------------------------------
    256 |  0x301   0x303    0x305    0x307
    32k |  0x310   0x313    0x316    0x319
    64k |  0x311   0x314    0x317    0x31A
    16M |  0x312   0x315    0x318    0x31B

Dieser Hex-Wert muss als zusätzlicher Parameter in **/boot/grub/menu.lst**
an den Kernel angehängt werden. Folgendermaßen kann sowas aussehen:

```
kernel /boot/vmlinuz-2.6.26-2-686 root=/dev/ida/c0d0p1 ro vga=0x318 quiet
```

Aktuell für Grub 2: **/boot/grub/grub.cfg**. Sieht bisschen anders aus,
funktioniert aber genauso:

```
linux   /boot/vmlinuz-2.6.32-22-generic root=UUID=92892dbf-af24-4dbd-b2a4-8debdbb08981 ro  vga=0x318 quiet splash
```

Sollte es eventuell noch eine schönere Möglichkeit für Grub 2 geben (was
ich mir durchaus vorstellen kann) bitte ich diese doch kurz zu kommentieren
:)
