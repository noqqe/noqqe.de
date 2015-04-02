---
date: 2010-04-24T10:44:11+02:00
type: post
slug: git-workflow-und-branching
comments: true
title: Git | Workflow und Branching
aliases:
- /archives/998
---

![git-logo](/uploads/2010/02/git-logo.png)

Man merkt bestimmt, dass mich [Git](http://scm-git.org) im Moment sehr fasziniert. Es ist eine Mischung aus "Wow, ist das umfangreich" und "Ah Dokumentation ist alles". Mittlerweile habe ich die verschiedensten Projekte in meinen Git-Server ausgelagert. Sogar meine [Tiddlywiki](http://tiddlywiki.com) hab ich aufgelöst und als Git-Repo umgesetzt.

Ich las jedenfalls viel in [www.progit.org](http://www.progit.org). Progit ist übrigens auch ein sehr schönes Projekt. Es stellt eine Dokumentation für alle Sprachen über Git dar. Diese Dokumentation ist auf [Github](http://github.com/progit/progit) für jeden forkbar und jeder kann theoretisch daran mitwirken.

Ich schweife schon wieder ab. In Progit findet man eine [wundervolle Beschreibung](http://progit.org/book/ch3-4.html) wie man in Git einzelne Workflows bzw Branches verwaltet und wie man am effektivsten mit ihnen arbeitet. Angenommen ich habe bereits in einem bestehenden Projekt ca. 5 Commits und möchte aber vom Inital-commit noch einmal anfangen, um in eine andere Richtung weiterzuentwickeln.


                Master
                |
    o--o--o--o--o


Sieht das ca so aus. Es besteht nun die Möglichkeit (und das ist der Grund warum ich diesen Post hier verfasse) eine neue "Entwicklungssparte" aka Branch zu starten und auf Anfang zu setzen.


    Initial     Master
    |           |
    o--o--o--o--o


```
# Neuen Branch erstellen namens inital
git branch initial
```


```
# Branches anzeigen
git branch -a -v
* master                 0bfb896 removed r_error() and added some comments
remotes/origin/HEAD    -> origin/master
remotes/origin/initial d6600f1 First inital commit
remotes/origin/master  0bfb896 removed r_error() and added some comments
```


```
# Branche initial auf spezielle Commitnummer setzen
echo "d6600f10479bb2d0d69aa8086ebe4e3149d4ef76" >> roborobo.git/.git/refs/heads/initial
```


```
# In den neu erstellten Branch wechseln
git checkout initial
```


```
# Via git log kontrollieren
git log
```


```
# neuen Branch zum Server pushen
git push origin initial
```


Weiterhin angenommen ich entwickle in der neuen Spalte "initial" weiter und commite das wiederrum dürfte das bild so aussehen:


    Initial
       |
    o--o        Master
    |           |
    o--o--o--o--o


<3 Git.
