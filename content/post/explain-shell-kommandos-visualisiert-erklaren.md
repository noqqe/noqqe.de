---
date: 2010-08-29T22:31:32+02:00
type: post
slug: explain-shell-kommandos-visualisiert-erklaren
status: publish
comments: true
title: Explain | Shell-Kommandos visualisiert erklären
alias: /archives/1224
categories:
- Coding
- Ubuntu
- ubuntuusers
tags:
- bash
- debian
- dokumentation
- explain
- find
- png
- python script
- shell
- ubuntu
- vain
---

Neulich bin ich über das [github Profil](http://github.com/vain) von [Peter Hofmann](http://uninformativ.de) gestolpert. Darin befand sich ein Projekt, welches ich sehr interessant fand.

[Explain](http://github.com/vain/explain) versucht Shell Kommandos zu erklären und zu visualisieren. Gerade für Blogs oder andere Dokumentationen finde ich das mehr als sinnvoll. Es erstellt aus einem simpel gestricktem Markdown File eine ASCII-Art ähnliche Erläuterung des Kommandos. Beispielsweise:


    $ ./explain.py command.markdown
    find . -iname '*.png' -exec echo '<br><img src="{}">' ; > gallery.html
    __/ | ___________/  ________/ ___________________/ |  ___________/
      |  |       |             |               |           |        |
      |  |       |             |               |           |        - Ausgeben nach
      |  |       |             |               |           |           gallery.html
      |  |       |             |               |           |
      |  |       |             |               |           - find Syntax Ende.
      |  |       |             |               |
      |  |       |             |               - mit folgendem Inhalt aus.
      |  |       |             |
      |  |       |             - und führe echo
      |  |       |
      |  |       - alle Dateien die mit .png enden
      |  |
      |  - im aktuellen Verzeichnis
      |
      - Finde (via find)


(PlainText: [/uploads/2009/09/015](/uploads/2009/09/015))

Die Syntax des Files das zur Deklaration der Ausgabe dient:


    find . -iname '*.png' -exec echo '<br><img src="{}">' ; > gallery.html
    ---- ! -------------  ---------- --------------------- ! -------------

    Finde (via find)

    im aktuellen Verzeichnis

    alle Dateien die mit .png enden

    und führe echo

    mit folgendem Inhalt aus.

    find Syntax Ende.

    Ausgeben nach gallery.html


Die Trennzeichen  sind via Parameter austauschbar und auch ansonsten tut das kleine Python Script seinen Job hervorragend. Sollte demnächst mal wieder ein Kommando erläutert werden müssen, werde ich definitiv darauf zurückgreifen. Weitere Beispiele auch unter:

[1] [http://www.uninformativ.de/?section=news&ndo=single&newsid=118](http://www.uninformativ.de/?section=news&ndo=single&newsid=118)
[2] [http://github.com/vain/explain](http://github.com/vain/explain)
