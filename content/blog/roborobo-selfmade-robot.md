---
aliases:
- /blog/2010/04/11/roborobo-selfmade-robot
- /archives/977
comments:
- author: jurkan
  content: "<p>Jetzt gibts daf\xFCr auch ein Arch Linux Package im AUR :D</p><p>Funktioniert
    ganz gut, allerdings kriege ich beim Aufruf von \"sudo roborobo\" immer die Meldung
    \"error\", hervorgerufen in Zeile 68 (auskommentieren =&gt; Error weg), allerdings
    habe ich keine Ahnung wo das her kommt. </p><p>Irgendwelche Ideen?</p>"
  date: '2010-04-11T17:11:18'
- author: jurkan
  content: <p>OK, Error tritt nur auf, wenn ich die Datei /etc/hosts mit dabei habe.
    Deren Rechte sind allerdings -rw-r--r-- root root, also keine Ahnung wo der Error
    herkommt...</p>
  date: '2010-04-11T17:36:18'
- author: noqqe
  content: "<p>Hallo Jurkan!</p><p>vielen Dank f\xFCr deine Kommentare :)<br>Arch
    Linux AUR ? :) musst du mir erkl\xE4ren ;)</p><p>Der Error mh. Mit sudo hab ich
    es ehrlichgesagt noch garnicht ausprobiert. <br>Ich werde es mal durchspielen.</p><p>kann
    man mit dir irgendwie kontakt aufnehmen? :) Jabber oder \xE4hnliches? Mich intressiert
    dein Arch Paket!</p><p>Gr\xFC\xDFe</p>"
  date: '2010-04-11T21:22:46'
- author: jurkan
  content: "<p>jurkan@jabber.org (der Server ist leider etwas instabil, hat mich nie
    gest\xF6rt da ich jabber seltenn benutze)</p><p>alternativ auch ICQ -zensiert-
    </p><p>Bitte mit aussagekr\xE4ftiger Authorisierungsanfrage, sonst wirds geblockt
    ;-)</p>"
  date: '2010-04-11T21:37:24'
- author: noqqe
  content: "<p>Vielen Dank Jurkan nochmal f\xFCr die Arch AUR Portierung. <br><a href=\"http://aur.archlinux.org/packages.php?ID=36366\"
    rel=\"nofollow\">http://aur.archlinux.org/packages.php?ID=36366</a></p><p>Bug ist gefixed.
    Lag an der File-Identifikation. War nicht eindeutig genug. Update ist in Git(-hub).
    Genaugenommen 2 Zeichen.</p>"
  date: '2010-04-12T00:03:29'
- author: Uli
  content: "<p>Gibt es die M\xF6glichkeit das Schreiben ins Logfile nur bei Fehlern
    zu schreiben. Das Toolf f\xFCllt mir das Logfile.</p>"
  date: '2012-10-12T16:46:41'
- author: noqqe
  content: "<p>\xC4hm, also ich hab den source jetzt garnicht mehr im Blick, weil
    ich mittlerweile auch zu git gewechselt bin und damit meine Dateien verwalte.
    </p><p>Wenn dann m\xFCsstest du da selbst Hand anlegen... </p>"
  date: '2012-10-13T11:20:45'
- author: Uli
  content: "<p>Okay, denke das ist ne sache vom cron und hat direkt mit dem script
    nichts zu tun.</p><p>Danke f\xFCr Deine schnelle Antwort</p>"
  date: '2012-10-15T17:37:33'
date: '2010-04-11T08:40:28'
tags:
- development
- roborobo
- daemon
- shell
- ubuntu
- robot
- sicherung
- linux
- deb
- backup
- debian
title: roborobo | selfmade robot
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
