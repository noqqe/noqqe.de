---
date: 2008-12-05T15:37:56+02:00
type: post
slug: unpack-entpackt-sie-alle-marke-eigenbau
status: publish
comments: true
title: unpack - Entpackt sie alle - Marke Eigenbau
alias: /archives/346
categories:
- Coding
- General
- Linux
tags:
- .tar.gz
- archive tool
- debian
- entpacken
- Linux
- server
- shell skript
- tar
- unpack
- unpacking tool
- zip
---

Vor ein paar Tagen erst haben wir uns in der Arbeit über die [Komplexität der tar Befehle unterhalten und die Millionen Archivetypen](http://www.linux-fuer-alle.de/doc_show.php?docid=76) die es gibt. Habe mir dann etwas gedanken gemacht, wie das ganze leichter zu haben wäre und ein kleines Skript geschrieben.

Relativ unspektakulär hab ich es **unpack** getauft :)
Ziel des Skriptes ist es, nach Aufruf und Übergabe(durch Parameter) der zu entpackenden Datei, selbstständig das Archivformat zu erkennen und zu entpacken. Im Endeffekt ist unpack nichts anderes als ein textbasierter Archivmanager (mir ist natürlich klar das ich nicht der erste bin der auf diese glorreiche Idee kommt :P  ). Aber vielleicht gehts dem einen oder anderen genauso wie uns, das wir uns nie die nötigen Parameter merken können.

Im Anhang erstmal v1 des Skripts. Eventuell werd ich mich noch mal drüber machen und diverse neue/alte Formate einfügen sowie die Struktur überarbeiten. Aber erstmal hat es alles was es braucht um zu funktionieren.

"Install - Guide":



	
  * Skript [downloaden](http://zwetschge.org/unpack/) | [Quellcode](http://paste.pocoo.org/show/93956/)



	
  * in /home/user/.bashrc alias eintragen:




> alias unpack = 'sh /pfad/zum/skript'
(/usr/bin/skript empfohlen...)





	
  * Ausführbar machen$ chmod +x /pfad/zum/skript



	
  * Skript ausführen:




> $ unpack /pfad/zum/archiv





	
  * Entpackte Dateien befinden sich im aktuellen Verzeichnis! (Daran werde ich noch arbeiten :) )


Wer  Shell Fanatiker ist oder viel auf Servern herumfuhrwerkt wird hoffe ich Spaß daran haben. Ansonsten für die GUI-Freaks gibts ja immernoch den guten alten Archivmanager :)

Das einzig doofe an der Geschichte ist... jetzt musst ich mich mit den .tar Parametern auseinander setzen ... und würde sie auch so wissen *seufz*

#############################

# Crackpodsmeinung: :P

**crackpod:** das ist wirklich ganz geil
**crackpod:** dann kannst einfach
**crackpod: **für jeden Archivtypen
**crackpod:**** **so ganz lässig
**crackpod:** und leet
**crackpod:** unpack xyz.xy
**noqqe@jabber.ccc.de:** jaa da hats jemand erkannt :P  ^^
# Marc von MBlog hat mir die Ehre erwiesen und das Skript in seinen WeihnachtsKalender eingebaut :)
# [http://www.marcboe.de/adventskalender-9-dezember2/](http://www.marcboe.de/adventskalender-9-dezember2/)



