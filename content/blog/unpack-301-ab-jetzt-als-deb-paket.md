---
aliases:
- /archives/422
comments:
- author: Thomas
  content: "<p>Nicht schlecht Herr Specht, so ein .deb-Paket macht schon ganz sch\xF6n
    was her finde ich =)</p><p>Ich kann nur nicht ganz nachvollziehen, warum das erst
    nach der Installation \xFCber das .dep-Paket mit dem \"Tabben\" funktioniert,
    bei mir geht das auch, wenn ich einfach ein Script in einen Ordner der PATH-Variable
    lege oder einen alias anlege...</p>"
  date: '2009-01-11T17:02:35'
- author: noqqe
  content: "<p>Okay, danke f\xFCrs Lob erstmal ;)</p><p>Mh, also ... ich hatte das
    Skript vorher via --autoinstall installiert... und konnte nichts vervollst\xE4ndigen.
    Aber okay, warscheinlich kam da wieder meine Unf\xE4higkeit ins Spiel :) . Jetzt
    gehts aufjedenfall :)</p>"
  date: '2009-01-11T17:16:45'
date: '2009-01-10T17:43:22'
slug: unpack-301-ab-jetzt-als-deb-paket
tags:
- development
- shell
- debian paket
- unpacking tool
- linux
- deb
- unpack
- debian
- bash
title: unpack-3.0.1 ab jetzt als .deb Paket
---

Als ich heute morgen nochmal über [syncN](http://zwetschge.org/syncN/)
drüber gesehen habe, wollte ich den --autoinstall mal überarbeiten. Aber
irgendwie hatte ich nicht das Gefühl als würde das jemals was
problemfreies. Deshalb hab ich mich hingesetzt, das Install-Skript
rausgenommen und angefangen ein .deb Paket zu erstellen. Ein bisschen
Gebastel mit der DateiStruktur und gekonfiguriere via dpkg und schon kam
das:

[http://zwetschge.org/unpack/unpack_3.0.1/](http://zwetschge.org/unpack/unpack_3.0.1/)

dabei heraus.

Wesentliche Änderungen:


* .deb Paket und dadurch weder Probleme mit Installation/Anpassung ans
  System oder nicht vorhandenen Abhängigkeiten wie tar, unrar und bzip2.
  Diese werden jetzt automatisch mit installiert falls nicht vorhanden.
* man - Page hinzugefügt (man unpack)
* Code wesentlich übersichtlicher
* Es können jetzt .deb Archive entpackt werden.
* (War mir persöhnlich wichtig) Durch .deb Installation ist jetzt
  Autovervollständigung via TAB möglich ( unp->TAB=unpack)

Übrigens **umbedingt** vor Installtion die alten BASHRC aliase
auskommentieren oder am besten rauslöschen! Ich hoffe es gefällt. Viel Spaß
damit!