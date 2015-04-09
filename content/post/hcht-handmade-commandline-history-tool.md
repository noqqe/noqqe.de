---
date: 2010-12-01T20:58:12+02:00
type: post
comments: true
title: hcht | Handmade Commandline History Tool
aliases:
- /blog/2010/12/01/hcht-handmade-commandline-history-tool
- /archives/1407
categories:
- Shell
- Development
- Linux
tags:
- bash
- bash-it
- commandline
- handmade
- history
- store
- tool
---

Mit dem Namen habe ich mich natürlich mal wieder selbst übertroffen. NICHT.
Ehe man sich versieht wird aus einem kleinen Fetzen Code ein ausgewachsenes
Plugin für bash-it, in dem der Name des Code-Schnippsels schon so tief
sitzt, das sich der Aufwand nicht mehr lohnt ihn zu ändern. Aber mal weg
von Oberflächlichkeiten hin zum Code.

{{< figure src="/uploads/2010/12/4094723007_34388e39d8.jpg" >}}

Creative Commons by [Iwan Gabovitch](http://www.flickr.com/photos/qubodup/) "Notizen"

Es war Dienstag und ich war auf der Suche nach einem kleinen Tool oder
einer Idee, wie ich am Besten kleine einzeilige "Hacks", gebastelte Regular
Expressions, MicroNotizen, ToDo's, Logfile Schnippsel oder sonstige
Informationen in Textform speichern und aufheben kann.  Nachdem ich nichts
fand fing ich an mir selber so ein kleines Tool zu schreiben. Als Plugin
für [bash-it](http://github.com/revans/bash-it).


## Funktionen

Idee ist einfach. Alles wird in einem zentralen Ordner abgespeichert und
dieser wird mit tollen Features durch hcht befüllt.

###Editor ###

Die Basis sozusagen. Eine Notiz oder ein Kommando per Hand einfügen

```
$ hcht das-ist-eine-Notiz.hch
```

###List-Funktion###

Anzeigen aller abgespeicherten Files

```
$ hcht
```

###Einzeilige Notiz###

Den ganzen Spaß gibts auch einzeilig.

```
$ hcht Hallo, das ist eine kleine Notiz
```

###Pipeable###

Natürlich hat mein [Lesen von Stdin Post](/archives/1402) auch einen Sinn gehabt.

```
$ cat mail.log | hcht maillog

```

###Wiederholbar###

Die bashinterne Funktion ist zum Beispiel auch hilfreich. Angenommen man
hat grade einen total coolen Hack gebastelt und will diesen aufheben:

```
$ find . -iname '*.png' -exec echo '<br><img src="{}">'  ; > gallery.html
$ hcht !!
```

Für mehr und vor allem genauere Beschreibung siehe den Source und das
Readme auf github:

Dabei heraus kam:
[https://github.com/revans/bash-it/blob/master/plugins/hcht.plugin.bash](https://github.com/revans/bash-it/blob/master/plugins/hcht.plugin.bash)
