---
categories:
- bash
- development
- osbn
- ubuntuusers
- web
date: 2015-04-06T18:16:25+02:00
tags:
- bookmarks
- pinboard
- bm
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

