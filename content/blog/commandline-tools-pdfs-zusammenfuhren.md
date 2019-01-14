
---
aliases:
- /blog/2011/04/20/commandline-tools-pdfs-zusammenfuhren
- /archives/1622
comments:
- author: Thomas
  content: "<p>Hey,<br>du kannst bei einem einfachen Zusammenf\xFChren von PDFs das
    cat weglassen. Das ist nur interessant, wenn du nur bestimmte Seiten aus einem
    Dokument zusammenf\xFChren m\xF6chtest.</p><p>Also f\xFCr gesamte Dokumente:<br>pdftk
    Dokument1.pdf Dokument2.pdf Dokument3.pdf output Gesamtdokument.pdf</p><p>F\xFCr
    einzelne Seiten gehst du dann wie folgt vor:<br>pdftk A=Dokument1.pdf B=Dokument2.pdf
    C=Dokument3.pdf cat A2-5 B1-3 A1 C6-10 output Gesamtdokument.pdf</p><p>So kann
    man sich sehr einfach seine PDFs zusammenw\xFCrfeln.</p>"
  date: '2011-04-20T14:51:47'
- author: Mario
  content: '<p>Wenn ein GUI-Programm bevorzugt wird: <a href="http://live.gnome.org/PdfMod"
    rel="nofollow">http://live.gnome.org/PdfMod</a><br><a href="http://www.pdfsam.org/"
    rel="nofollow">http://www.pdfsam.org/</a><br><a href="http://sourceforge.net/projects/pdfshuffler/"
    rel="nofollow">http://sourceforge.net/projects/pdfshuffler/</a></p>'
  date: '2011-04-20T14:54:27'
- author: Niko
  content: "<p>Ich stand vor einiger Zeit auch vor dem Problem, dass ich eine Hand
    voll PDFs zu einem zusammenf\xFChren musste.</p><p>Ich habe damals pdfmod (<a
    href=\"http://live.gnome.org/PdfMod)\" rel=\"nofollow\">http://live.gnome.org/PdfMod)</a>
    verwendet. Wer eine GUI m\xF6chte, der wird vermutlich darauf zur\xFCckgreifen
    wollen.</p>"
  date: '2011-04-20T14:58:12'
- author: sinusQ
  content: <p>Wie auch Niko schon sagte.. Ich verwende meistens PDF MOD. </p><p><a
    href="http://wiki.ubuntuusers.de/PDF?highlight=Pw%20Tbaustell%20Zpdfmod#PDF-Mod"
    rel="nofollow">http://wiki.ubuntuusers.de/PDF?highlight=Pw%20Tbaustell%20Zpdfmod#PDF-Mod</a></p>
  date: '2011-04-20T16:43:01'
- author: Funatiker
  content: <p>Und was genau spricht gegen die Verwendung von pdfjoin?</p>
  date: '2011-04-20T16:57:59'
- author: BigBomber
  content: <p>Und was spricht gegen Ghostscript?</p><p>Es gibt viele Wege nach Rom</p>
  date: '2011-04-20T18:05:10'
- author: noqqe
  content: "<p>Ich glaube so generell habe ich keinerlei abwertige Bermerkungen \xFCber
    andere PDF Tools gepostet</p><p>Sollte der Post den Gedanken suggeriert haben,
    w\xFCrde ich diese Anschludigungen nat\xFCrlich sofort zur\xFCckweisen :) </p><p>Ich
    habe nach einer L\xF6sung f\xFCr mein Problem gesucht und dieses Problem wurde
    (in meinem Fall) von pdftk gel\xF6st (und das auch noch sch\xF6n). </p><p>Sollten
    andere PDF-Tools f\xFCr mich nicht in Betracht gekommen sein, liegt das lediglich
    daran das sie sich meiner Kenntnis entzogen. Aber daf\xFCr wurde ja hier in den
    Kommentaren entsprechend gesorgt. </p><p>@Thomas: Danke f\xFCr den Tipp! Das mit
    den spezifischen Seiten werde ich auf jedenfall noch ausprobieren :)</p>"
  date: '2011-04-20T18:17:42'
- author: des Direktors Gehilfe
  content: "<p>Ich benutze daf\xFCr immer PDF-Shuffler, denn bei sowas komm ich mit
    einer GUI deutlich schneller ans Ziel und PDF-Shuffler kann dar\xFCber hinaus
    auch die Reihenfolge der Seiten neu ordnen, Seiten drehen, einzelne Seiten einfach
    entfernen und man kann die PDFs einfach per Drag&amp;Drop importieren.</p>"
  date: '2011-04-20T18:33:49'
- author: Paradiesstaub
  content: <p>Das funktioniert auch:</p><p>gs -q -sPAPERSIZE=letter -dNOPAUSE -dBATCH
    -sDEVICE=pdfwrite -sOutputFile=out.pdf namePDF1.pdf namePDF2.pdf</p>
  date: '2011-04-20T19:54:13'
- author: Thomas
  content: <p>Versuch' es mal mit PDF Split and Merge. <a href="http://www.pdfsam.org/"
    rel="nofollow">http://www.pdfsam.org/</a><br>Dort hast Du es auch noch grafisch.</p>
  date: '2011-04-20T23:01:03'
- author: Maria Amman
  content: "Hi schaut euch mal http://yumpu.com/de/pdfs-zusammenfuegen an. Das ist
    auch sehr n\xFCtzlich!"
  date: '2015-02-03T15:18:57.196778'
date: '2011-04-20T10:14:50'
tags:
- merge
- split
- pdf
- bash
title: "Commandline Tools | PDFs zusammenf\xFChren"
---

Aufgrund meiner aktuell vorherrschenden Bewerbungsphase wollte ich ein paar
PDF Dateien (hauptsächlich Zertifikate und Zeugnisse) in **eine** PDF Datei
zusammenführen. Auf der Suche nach einem derartigen Programm liefen mir
natürlich allerlei (dem Anschein nach) hübsche Windows Tools für diese
Aufgabe über den Weg. Wie schon oft, fand ich dann aber
[im Unixboard meine Antwort](http://www.unixboard.de/vb3/showthread.php?31512-PDF-Dateien-%28mehrere%29-zusammenf%FChren-in-ein-PDF).

{{< figure src="/uploads/2011/04/Adobe_PDF.png" >}}

Das Commandline Tool
**[pdftk](http://www.pdflabs.com/tools/pdftk-the-pdf-toolkit/)**. Wunderbar
für [Debian](http://packages.debian.org/squeeze/pdftk) und
[Ubuntu](http://packages.ubuntu.com/natty/pdftk) paketiert.

```
sudo aptitude install pdftk
```

Ich muss trotzdem zugeben, dass die Syntax mir etwas ungewöhnlich erschien,
aber nicht unbezwingbar ist :) Nach ein paar Blicken in die Manpage, kam
ich auch da an wo ich wollte.

```
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
```
