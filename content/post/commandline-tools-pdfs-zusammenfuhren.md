---
date: 2011-04-20T12:14:50+02:00
type: post
slug: commandline-tools-pdfs-zusammenfuhren
comments: true
title: Commandline Tools | PDFs zusammenführen
aliases:
- /archives/1622
categories:
- Bash
- Development
- Debian
- Blog
- ubuntuusers
- Web
tags:
- combine
- kombinieren
- merge
- pdf
- pdfs kombinieren
- split
- zusammenführen
---

Aufgrund meiner aktuell vorherrschenden Bewerbungsphase wollte ich ein paar PDF Dateien (hauptsächlich Zertifikate und Zeugnisse) in **eine** PDF Datei zusammenführen. Auf der Suche nach einem derartigen Programm liefen mir natürlich allerlei (dem anschein nach) hübsche Windows Tools für diese Aufgabe über den Weg. Wie schon oft, fand ich dann aber [im Unixboard meine Antwort](http://www.unixboard.de/vb3/showthread.php?31512-PDF-Dateien-%28mehrere%29-zusammenf%FChren-in-ein-PDF).

{{< figure src="/uploads/2011/04/Adobe_PDF.png" >}}

Das Commandline Tool **[pdftk](http://www.pdflabs.com/tools/pdftk-the-pdf-toolkit/)**. Wunderbar für [Debian](http://packages.debian.org/squeeze/pdftk) und [Ubuntu](http://packages.ubuntu.com/natty/pdftk) paketiert.

```
$ sudo aptitude install pdftk
```


Ich muss trotzdem zugeben, dass die Syntax mir etwas ungewöhnlich erschien, aber nicht unbezwingbar ist :) Nach ein paar Blicken in die Manpage, kam ich auch da an wo ich wollte.


    $ pdftk novell-cert.pdf ripe.pdf lpic-1.pdf cisco-ccna1.pdf cat output Zertifikate.pdf
      ___/ ________________________________________________/  _/ ____________________/
        |                            |                           |             |
        |                            |                           |             - Ausgabe an
        |                            |                           |                Zertifikate.pdf
        |                            |                           |
        |                            |                           - Ausgabe der PDFs
        |                            |
        |                            - Angabe aller zu kombinierenden PDFs
        |
        - Programmaufruf
