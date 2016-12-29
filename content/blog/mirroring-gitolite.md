---
comments:
- author: spion
  content: <p>das ganze mit logger gibts wo? ;)</p>
  date: '2013-02-17T21:38:48'
- author: Ranlvor
  content: "<p>soweit ich wei\xDF kannst du die Schleife, die \xFCber alle Branches
    iteriert weglassen und daf\xFCr git fetch origin benutzen, das sollte alle Branches
    laden und mit Namen nach dem Muster origin/&lt;name&gt; verf\xFCgbar machen. Wenn
    du wirklich alle Branches direkt im Hauptnamespace haben willst kannst du auch
    einmal mit git fetch origin alle Branches laden und mit git checkout $rbranch;
    git merge origin/$rbranch die merges durchf\xFChren, dann musst du f\xFCr jedes
    Repository nur eine SSH Verbindung und nicht eine pro Branch aufbauen.&lt;/name&gt;</p>"
  date: '2013-02-18T15:12:25'
date: '2013-02-17T18:14:00'
tags:
- git
- slave
- administration
- devops
- gitolite
- shell
- mirror
- debian
title: Mirroring gitolite
---

[gitolite](https://github.com/sitaramc/gitolite) hat einen eingebauten
Mechanismus sich selbst zu replizieren.  Finde aber unschön dass ich dafür
einen zweiten gitolite betreiben und neue Repos auf beiden Systemen anlegen
muss.

Wie also das Backup realisieren. `rsync` für git ist in an so vielen
Stellen eine Verschwendung von Ressourcen, Zeit und Funktionalität. Macht
einfach keinen Sinn.

Wer sich schonmal mit

    ssh git@git.example.com

auf seinen Server verbunden hat weiss aber, dass bei vorhandenem Public Key
alle für diesen User verfügbaren Repositories ausgespuckt werden.
Mit diesem Feature kann man sich den Slave einfach selber bauen.

``` bash
#!/bin/bash

MASTER="git@git.example.com"
BASEDIR="/data/gitslave/"

for repo in $(ssh $MASTER | grep '^ R' | awk '{print $3}') ; do
  if [ -d $BASEDIR/$repo ]; then
    cd $BASEDIR/$repo
    for rbranch in $(git branch -a | grep -v '^*\ master\|HEAD\ -' | awk -F/ '{print $3}'); do
      git checkout $rbranch
      git pull origin $rbranch
    done
    cd -
  else
    git clone $MASTER:$repo $BASEDIR/$repo
  fi
done  &>/dev/null
```

Das Skript clont eigenständig alle Remote Branches und neue Repositories.
Relativ pragmatischer Ansatz.  Der Übersichtlichkeit halber, habe ich die
`logger` Commands jetzt weggelassen.

Wems gefällt darfs behalten.