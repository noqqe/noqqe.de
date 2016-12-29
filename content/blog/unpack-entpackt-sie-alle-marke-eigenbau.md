---
aliases:
- /archives/346
- /blog/2008/12/05/unpack-entpackt-sie-alle-marke-eigenbau/
comments:
- author: Marc
  content: "<p>Das ist ja mal ein n\xFCtzliches Skript, kann ich gut gebrauchen.<br>Vielleicht
    bringe ich das sogar noch irgendwann in meinem Adventskalender. ;)</p>"
  date: '2008-12-08T12:56:17'
- author: Marc
  content: <p>Schreib mal Marc, statt "MarcBoe". ;)</p>
  date: '2008-12-09T17:37:15'
date: '2008-12-05T13:37:56'
slug: unpack-entpackt-sie-alle-marke-eigenbau
tags:
- development
- shell skript
- zip
- tar
- targz
- entpacken
- unpacking tool
- server
- archive tool
- linux
- unpack
- debian
title: unpack - Entpackt sie alle - Marke Eigenbau
---

Vor ein paar Tagen erst haben wir uns in der Arbeit über die [Komplexität
der tar Befehle unterhalten und die Millionen
Archivetypen](http://www.linux-fuer-alle.de/doc_show.php?docid=76) die es
gibt. Habe mir dann etwas gedanken gemacht, wie das ganze leichter zu haben
wäre und ein kleines Skript geschrieben.

Relativ unspektakulär hab ich es **unpack** getauft :) Ziel des Skriptes
ist es, nach Aufruf und Übergabe(durch Parameter) der zu entpackenden
Datei, selbstständig das Archivformat zu erkennen und zu entpacken. Im
Endeffekt ist unpack nichts anderes als ein textbasierter Archivmanager
(mir ist natürlich klar das ich nicht der erste bin der auf diese
glorreiche Idee kommt :P  ). Aber vielleicht gehts dem einen oder anderen
genauso wie uns, das wir uns nie die nötigen Parameter merken können.

Im Anhang erstmal v1 des Skripts. Eventuell werd ich mich noch mal drüber
machen und diverse neue/alte Formate einfügen sowie die Struktur
überarbeiten. Aber erstmal hat es alles was es braucht um zu funktionieren.

"Install - Guide":

* Skript [downloaden](http://zwetschge.org/unpack/) |
  [Quellcode](http://paste.pocoo.org/show/93956/)
* in /home/user/.bashrc alias eintragen:
  ` alias unpack = 'sh /pfad/zum/skript'`
* Ausführbar machen
  `$ chmod +x /pfad/zum/skript`
* Skript ausführen:
  `$ unpack /pfad/zum/archiv`
* Entpackte Dateien befinden sich im aktuellen Verzeichnis! (Daran werde
  ich noch arbeiten :) )

Wer  Shell Fanatiker ist oder viel auf Servern herumfuhrwerkt wird hoffe
ich Spaß daran haben. Ansonsten für die GUI-Freaks gibts ja immernoch den
guten alten Archivmanager :)

Das einzig doofe an der Geschichte ist... jetzt musst ich mich mit den .tar
Parametern auseinander setzen ... und würde sie auch so wissen *seufz*

> crackpod:** das ist wirklich ganz geil
> crackpod:** dann kannst einfach
> crackpod: **für jeden Archivtypen
> crackpod:**** **so ganz lässig
> crackpod:** und leet
> crackpod:** unpack xyz.xy

und [http://www.marcboe.de/adventskalender-9-dezember2/](http://www.marcboe.de/adventskalender-9-dezember2/)