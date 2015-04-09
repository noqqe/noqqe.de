---
date: 2010-04-11T10:40:28+02:00
type: post
comments: true
title: roborobo | selfmade robot
aliases:
- /blog/2010/04/11/roborobo-selfmade-robot
- /archives/977
categories:
- Development
- Linux
- PlanetenBlogger
- Shell
tags:
- backup
- daemon
- deb
- debian
- roborobo
- robot
- sicherung
- ubuntu
---

![Robot-icon](/uploads/2010/04/Robot-icon.png) Auf dem Weg durch die
Filesysteme meiner Rechner/Laptops/Server kam ich immer wieder in
Situationen, in denen ich gerne Files (die mir wichtig waren) an einer
bestimmten Stelle aufheben wollte. Ich hab über die Monate bzw. fast schon
Jahre hin immer andere Systeme, Praktiken und Plätze entwickelt in denen
ich diese Configs und ähnliches ablege. Über kurz oder lang ist aber jede
dieser Methoden zu aufwändig oder zu unstrukturiert. Wenn ich
Konfigurationsdateien von Daemons editierte, kopierte ich vorher die alten
Files an eine bestimmte Stelle. Ziemlich Standart. Der Vorgang ist an sich
ziemlich mühsam. Wirr liegen irgendwo irgendwelche Files rum.

Vor ein paar Wochen habe ich dann angefangen mir ein kleines Helferlein zu
coden. Anfangs war dieses Helferlein nur für mich gedacht und demnach
relativ speziell. Ich nannte ihn "roborobo".

Was er tut ? Ich gebe meinem Helfer einfach das File "in die Hand". Alles
andere erledigt er.

File hinzufügen
```
$ roborobo /etc/postfix/main.cf`
```

Alle bekannten Files updaten:
```
$ roborobo
```

Er nimmt das File an, ordnet es ein und sichert es in seinem Verzeichnis
mit dem kompletten Verzeichnispfad nach Baumstruktur-Art. Außerdem prüft
roborobo jetzt jede Stunde anhand der sha1sum ob sich in dem File seit der
letzten Prüfung etwas getan hat. Falls Veränderungen da sind, wird das File
mit neuem Datum wieder abgespeichert. Das sieht ungefähr so aus:

    .roborobo/
    |-- etc
    |   |-- hosts
    |   |   |-- hosts-20100409-1348
    |   |   `-- hosts-20100409-1651
    |   `-- network
    |       `-- interfaces
    |           |-- interfaces-20100409-1654
    |           `-- interfaces-20100409-1655


Ich brauche mich somit um _nichts_ mehr kümmern. Gebe dem "kleinen" die
Files die mir wichtig sind und er passt darauf auf. Fühlt sich irgendwie an
wie ein Backup-Daemon ;)

Jedenfalls, habe ich roborobo jetzt für den Einsatz auf jedem beliebigen
System umgebaut und in ein Debianpaket gebastelt. Dokumentation erstellt,
Config-Dateien ausreichend selbsterklärend gestaltet usw.

Wer sich dafür interessiert oder mal testen mag:

[roborobo-Debianpaket bei Github](http://github.com/noqqe/roborobo/downloads)

[roborobo-Projekt auf Github](http://github.com/noqqe/roborobo/)

Content:

```
    roborobo
    |-- DEBIAN
    |   `-- control
    |-- etc
    |   |-- cron.d
    |   |   `-- roborobo
    |   `-- roborobo
    |       |-- roborobo.conf
    |       `-- roborobo.path
    -- usr
        |-- bin
        |   `-- roborobo
        -- share
            |-- doc
            |   `-- roborobo
            |       |-- changelog
            |       `-- copyright
            -- man
            -- man1
                `-- roborobo.1
```

Das ganze wie alles was ich tue, unter GPLv3.

Update 2014: Ich bin mir nicht sicher ob man das wirklich noch nutzen möchte.
Viel eher würde ich zu [etckeeper](http://joeyh.name/code/etckeeper/) raten.
