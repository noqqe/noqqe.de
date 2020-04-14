---
aliases:
- /blog/2010/10/09/git-verschiedene-repos-zusammenfuhren
- /archives/1302
comments:
- author: Speiderlars
  content: "<p>Hallo, irgend</p><p>wie klappt es bei mir so nicht.<br>Kann es sein
    das ich nach dem mkdir dot... noch etwas machen muss ? Zb denn Branch Master richtig
    init. </p><p>Ich komme an denn Punkt das ich <br>die Branches habe zb master:libs<br>dann
    will ich das Script ausf\xFChren und es kommt zb ein:</p><p>mv: cannot stat `/srv/repos/git/test/.git-rewrite/t/../index.new':
    No such file or directory</p><p>Der sub Ordner hei\xDFt libs und der Branch auch
    libs also habe ich in dem script auch:</p><p>SUBDIR_NAME=libs<br>BRANCH_NAME=libs</p><p>stehen.
    Ich hoffe du verstehst mein problem und kannst mir helfen.</p><p>Gru\xDF Speiderlars</p>"
  date: '2011-10-18T17:27:11'
- author: noqqe
  content: "<p>Mh. Und du bist auch wirklich im richtigen verzeichnis wenn du das
    Skript ausf\xFChrst?</p>"
  date: '2011-10-22T15:09:19'
date: '2010-10-09T15:28:58'
tags:
- development
- branching
- git
title: "Git | Verschiedene Repos zusammenführen"
---

In einem Anfall des Ordnungswahns entschloss ich mich im Mai (oder einem
deren anderen 11 Monate des letzten Jahres) dazu, mein textbasiertes Wiki,
meine Dotfiles und meine Scripte in separaten Git-Repos zu verwalten und zu
verteilen. Mit der Zeit merkt man aber, dass das ziemlicher Unsinn war.

Kurz: Ich brauchte eine Lösung, wie ich alle 3 Repositories (ohne
History-Verlust) zusammen gemerged bekomme.

```
    /repos/
    |-- dotfiles.git
    |-- scripts.git
     -- wiki.git
```

Hierfür hat [Zrajm C Akfohg](http://zrajm.org/) schon eine [wunderbare
schnelle Lösung](http://zrajm.org/ref/git-repo-merging.html) gefunden.
Zuerst ein neues Repo erstellen:

``` bash
mkdir ~/repos/noqqe.git
cd ~/repos/noqqe.git
git init
```

Im Grunde wird jedes Repository in einen separaten Branch gefetched:

``` bash
git fetch ~/repos/dotfiles master:dotfiles
mkdir dotfiles #Subdir fuer dotfiles
```

Um die Files im neuen Branch jetzt in einen Unterordner umzumappen hat
Zrajm auch einen wunderbaren Weg gefunden:

``` bash
SUBDIR_NAME=dotfiles
BRANCH_NAME=dotfiles
git filter-branch --index-filter
    'git ls-files -s |
        sed "s-t-&'"$SUBDIR_NAME"'/-" |
        GIT_INDEX_FILE=$GIT_INDEX_FILE.new git update-index --index-info &&
        mv $GIT_INDEX_FILE.new $GIT_INDEX_FILE
    ' "$BRANCH_NAME"
```

Wenn alles erfolgreich war, kann man den neuen Branch mit dem Master
zusammenführen:

``` bash
git merge dotfiles
```

und das Spiel wie beschrieben von vorne beginnen.
