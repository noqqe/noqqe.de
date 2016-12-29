---
comments:
- author: Matthias
  content: "Sieht gut aus und webkit2png gibts auf Linux eh nicht.  \nBeim mir kam
    eine Fehlermeldung beim adden: \"ggrep\" kannte mein System nicht und auch die
    Websuche ergab nichts. Ich habe einfach ein \"grep\" daraus gemacht und jetzt
    l\xE4ufts.  \nEine *rm* Funktion w\xE4re noch klasse, weil vermutlich irgendwann
    das h\xE4ndische L\xF6schen recht aufwendig wird. Und vielleicht eine Option,
    um den gefundenen Link in die *Zwischenablage* zu schieben.\n\nAuf alle F\xE4lle
    ein dickes **Dankesch\xF6n**."
  date: '2015-04-08T18:29:05.769749'
- author: noqqe
  content: "Ah siehst, das mit Webkit2png wusste ich garnicht. \n\nMit dem \"ggrep\"...
    Das hat mit OpenBSD zu tun, dass ich benutze. F\xFCr das Title Fetching brauche
    ich das \"Gnu Grep\", da das mitgelieferte grep unter BSD ein paar Optionen die
    ich brauche nicht mitbringt. \n\nBzgl. \"rm\": wie w\xE4re das am Besten umzusetzen?
    So was gel\xF6scht werden sollte? Einen eindeutigen Indikator gibts ja nicht."
  date: '2015-04-09T09:15:09.063398'
- author: Anonymous
  content: 'Offtopic:

    Dein identicon: http://www.instructables.com/files/orig/FM5/CVAV/GNOD5VX7/FM5CVAVGNOD5VX7.png'
  date: '2015-04-09T17:33:39.965263'
- author: matthias
  content: "Ja, dazu m\xFCsste jeder Eintrag einen eindeutigen Schl\xFCssel bekommen
    anhand dessen er identifiziert werden kann. \n\nIch habe sogar noch weiter herum
    gesponnen: wenns au\xDFerdem noch git-Funktionalit\xE4t g\xE4be, k\xF6nnte auf
    einfach Wei\xDFe gepusht und gepult werden. Dadurch w\xE4re ein einfaches Verteilen
    auf verschiedene Instanzen m\xF6glich oder man sammelt sogar kollaborativ.\n\n*Die
    W\xFCnsche gehen nicht aus.*"
  date: '2015-04-09T17:45:52.773150'
- author: noqqe
  content: hahaha :D
  date: '2015-04-10T10:55:37.399750'
- author: Matthias
  content: "Ich habe das zweite jetzt mal mit *rsync* realisiert: https://github.com/mtthff/bm/tree/clipboard
    \nFunktioniert ganz rund. Per cronjob mache ich mir t\xE4glich ein Backup der
    Datei auf dem Host, falls ich mich mal vertue. Was h\xE4ltst du davon?"
  date: '2015-04-14T13:22:48.544812'
- author: noqqe
  content: "Huh, sieht doch gut aus. Ich mach auch Backups, aber ganz einfach \n\n`cp
    ~/.bookmarks ~/Backups/bookmarks.$(date '+\\%F')` \n\n\nSo Conflicts werden halt
    bei dir \xFCberschrieben :) Musst halt vorsichtig sein :=)"
  date: '2015-04-14T13:41:33.884715'
date: '2015-04-06T16:16:25'
tags:
- development
- web
- pinboard
- bm
- bookmarks
- bash
title: Bookmarks
---

Seit 2013 benutze ich Pinboard. Letztens flog auf GitHub allerdings ein
[Commandline Bookmark Manager](https://github.com/cym13/bookmark) vorbei.
Den Gedanken fand ich eigentlich total toll. Die Syntax fand ich komisch,
Ausgabe sah strange aus. Aber es kam auch noch erschwerend hinzu, dass es
nicht mal gebaut werden konnte. Also hab ich weiter gesucht und einen
ziemlich schönen gefunden.

{{< figure src="/uploads/2015/04/fork.jpg" >}}

[bm](https://github.com/tj/bm) ist in Bash geschrieben, sah schön aus,
einfach gestrickt. Eigentlich super. Aber auch hier liess die Bedienung
etwas zu wünschen übrig. Aber da das Zeug ja Opensource ist, fork it baby.

### bm Fork

Was anfangs nur ein "ja ich passe mir das Teil ein klein bisschen an" war,
wurde irgendwie zum Komplettumbau.

{{< figure src="/uploads/2015/04/bm.jpg" >}}

* Der ganze Dropbox Folder Kack ist weg
* HTML Preview im Browser mit webkit2png ist weg
* Title wird automatisch mit curl beim adden eines Links hinzugefügt.
* Syntax zum Bedienen stark umgeschrieben
* Clean Funktion entfernt
* Datum (wann wurde der Link geadded) wird automatisch hinzugefügt.
* Farblich / Formatsmäßig umstrukturiert.

Das Repo: [github.com/noqqe/bm](https://github.com/noqqe/bm)

Eine Lizenz würde ich dem ganzen auch gerne hinzufügen, aber da der
original Autor noch keine Lizenz hinzugefügt hat, muss ich damit erstmal
[warten](https://github.com/tj/bm/issues/14).

### Migration von Pinboard

Jetzt musste ich nur noch alle &gt;1000 Bookmarks von Pinboard umziehen.
Pinboard bietet einen `json` Export der eigenen Bookmarks mit allen Meta
Informationen an.

Diesen hab ich mir per einfacher Download Funktion lokal gespeichert und
mittels dieses Python Schnipsels in das "neue" bm Format umkonvertiert.

``` python
import json

with open('dump.json') as dataf:
  data = json.load(dataf)
    for x in data:
      print x["href"]+"|"+x["tags"]+"|"+x["time"]+"|"+x["description"]
```

Und wenn es beim Umleiten des `STDOUT` mit Python wegen des Encodings nicht
klappt, `export PYTHONIOENCODING=utf-8` benutzen. Saugeil. Damit hatte ich
immer Probleme.

Ich weiss von ein paar Leuten, dass Sie den RSS Feed meines Pinboard
Profils lesen. Für den Moment wird es dort keine neuen Dinge geben.
Eventuell bastle ich mich in `Python` ein kleines Skript dass meine
Bookmarks regelmäßig in RSS giesst. Da schau ich aber erstmal.