---
date: 2010-10-09T17:28:58+02:00
layout: post
slug: git-verschiedene-repos-zusammenfuhren
status: publish
comments: true
title: 'Git | Verschiedene Repos zusammenführen '
alias: /archives/1302
categories:
- Coding
- Debian
- git
- PlanetenBlogger
tags:
- branching
- fetch
- git
- merge
- repos
- vereinen
- zusammenführen
---

In einem Anfall des Ordnungswahns entschloss ich mich im Mai (oder einem deren anderen 11 Monate des letzten Jahres) dazu, mein textbasiertes Wiki, meine Dotfiles und meine Scripte in separaten Git-Repos zu verwalten und zu verteilen. Mit der Zeit merkt man aber, dass das ziemlicher Unsinn war.

Kurz: Ich brauchte eine Lösung, wie ich alle 3 Repositories (ohne History-Verlust) zusammen gemerged bekomme.


    /repos/
    |-- dotfiles.git
    |-- scripts.git
```
-- wiki.git


Hierfür hat [Zrajm C Akfohg](http://zrajm.org/) schon eine [wunderbare schnelle Lösung](http://zrajm.org/ref/git-repo-merging.html) gefunden.
Zuerst ein neues Repo erstellen:

```
mkdir ~/repos/noqqe.git
cd ~/repos/noqqe.git
git init
```


Im Grunde wird jedes Repository in einen separaten Branch gefetched:

```
git fetch ~/repos/dotfiles master:dotfiles
mkdir dotfiles #Subdir fuer dotfiles
```


Um die Files im neuen Branch jetzt in einen Unterordner umzumappen hat Zrajm auch einen wunderbaren Weg gefunden:


    SUBDIR_NAME=dotfiles
    BRANCH_NAME=dotfiles
    git filter-branch --index-filter
        'git ls-files -s |
            sed "s-t-&'"$SUBDIR_NAME"'/-" |
            GIT_INDEX_FILE=$GIT_INDEX_FILE.new git update-index --index-info &&
            mv $GIT_INDEX_FILE.new $GIT_INDEX_FILE
        ' "$BRANCH_NAME"


Wenn alles erfolgreich war, kann man den neuen Branch mit dem Master zusammenführen:

```
git merge dotfiles
```


und das Spiel wie beschrieben von vorne beginnen.
