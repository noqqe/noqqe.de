---
aliases:
- /blog/2009/09/24/shell-vollstandiger-file-path
- /archives/663
comments:
- author: Knorkebrot
  content: "<p>Wo ist da nun der Unterschied zu einem einfachen `find .`? Gibt das
    selbe aus und ist ungemein schneller.<br>M\xDFG</p>"
  date: '2009-12-16T23:11:10'
- author: noqqe
  content: "<p>Gut das du fr\xE4gst. <br>Angenommen ich bin im Homeverzeichnis und
    suche nach etwas mit find . kommt<br>./ordner/datei</p><p>Mit dem Befehl allerdings:<br>/home/user/ordner/datei</p><p>Das
    ist der gravierende Unterschied :)</p>"
  date: '2009-12-16T23:14:35'
- author: Knorkebrot
  content: "<p>Auf welchem System? Nach POSIX.1 zeigt die Option \"-d\" Verzeichnisse
    als w\xE4ren sie Dateien, sprich sie selbst, nicht den Inhalt, wenn man auf sie
    zeigt. Das gilt f\xFCr GNU-ls wie auch f\xFCr BSD-ls.<br>Also auf GNU/Linux, MacOSX,
    wie auch den ganzen BSDs.</p><p>Schneller dennoch w\xE4re `find $PWD`, das gibt
    dann auch den absoluten Pfad ;)</p><p>M\xDFG</p>"
  date: '2009-12-17T00:50:42'
- author: noqqe
  content: <p>Ah - der Tipp mit find $PWD ist gut... sehr gut ;) danke!</p>
  date: '2009-12-17T01:50:14'
date: '2009-09-24T17:59:58'
tags:
- development
- path
- filesystem
- file
- linux
title: "Shell | Vollst\xE4ndiger File-Path"
---

Sonst vergess ich es sowieso wieder:

```
find . -exec ls -d {} ;
```

gibt den vollständigen Datei-Pfad aus.  Sollte ausser mir nochjemand mal
seinen shoutcast trans mit einer Playlist befüllen müssen und aus Gründen
der total Pfusch-Config soetwas brauchen.