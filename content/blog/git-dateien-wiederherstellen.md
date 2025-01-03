---
aliases:
- /blog/2010/08/18/git-dateien-wiederherstellen
- /archives/1201
date: '2010-08-18T19:53:13'
tags:
- development
- git
title: 'Git | Dateien wiederherstellen'
---

Versehentlich gelöschte Dateien recovern. Schön wenn einem dann klar wird,
warum man ein VCS benutzt. In der Annahme das wahrscheinlich mehr als genug
(genug == >3) Blogs oder HowTo's diese Thematik bereits behandeln, erstelle
ich trotzdem kurz einen Post, wie sich Dateien mit Git zurückholen lassen.

### Datei ging gerade eben verloren

aus aktuellem HEAD wiederherstellen:

``` bash
git checkout HEAD -- verloren.txt
```

### Datei hat vor gewisser Zeit einmal existiert

Aus vorher gegangenem Commit, Branch oder Tag. Feststellen, wo das File
noch existiert haben könnte:

```
$ git log --oneline
5aadc10 formatierte Ausgabe
88e22fb Aufräumaktion
746f92c bugfix #1234
ee8a1da initial commit
```

Vorletzter Commit 88e22fb betitelt mit "Aufräumaktion" lässt stark darauf
schließen, dass hier etwas verloren gegangen sein könnte. Checkout lässt
sich eigentlich mit so ziemlich allem füttern, was ein Object ist und
anhand eines SHA1 Hashwertes identifizieren lässt. Gewählt wird der Commit
_vor_ der Aufräumaktion.

```
git checkout 746f92c -- verloren.txt
```

Nachdem die Datei nun wieder im aktuellen Working-Directory liegt:
via git-add hinzufügen

```
git add verloren.txt
```

und Commit absetzen.

``` bash
git commit -a -m "Ich werde ab jetzt besser aufpassen"
```
