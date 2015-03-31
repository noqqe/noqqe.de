---
layout: post
title: "Mirroring gitolite"
date: 2013-02-17T20:14:00+02:00
comments: true
categories:
- git
- gitolite
- debian
- slave
- mirror
- ubuntuusers
- osbn
---

[gitolite](https://github.com/sitaramc/gitolite) hat einen eingebauten Mechanismus sich selbst zu replizieren.
Finde aber unschön dass ich dafür einen zweiten gitolite betreiben und neue Repos
auf beiden Systemen anlegen muss.

Wie also das Backup realisieren. `rsync` für git ist in an so vielen Stellen
eine Verschwendung von Ressourcen, Zeit und Funktionalität. Macht einfach keinen
Sinn.

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

Das Skript clont eigenständig alle Remote Branches und neue Repositories. Relativ pragmatischer Ansatz.
Der Übersichtlichkeit halber, habe ich die `logger` Commands jetzt weggelassen.

Wems gefällt darfs behalten.
