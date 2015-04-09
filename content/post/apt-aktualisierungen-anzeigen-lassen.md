---
date: 2010-02-20T18:27:38+02:00
type: post
comments: true
title: Apt | Aktualisierungen anzeigen lassen
aliases:
- /archives/894
- /blog/2010/02/20/apt-aktualisierungen-anzeigen-lassen
categories:
- Development
- Linux
- ubuntuusers
tags:
- aktualisierungen
- anzeigen
- apt
- apt-get
- aptitude
- debian
- Kernel
- list
- search
- show
- ubuntu
- update
- updates
- upgrade
---

Für ein kleines Script, welches ich auf der Arbeit verwende, habe ich
versucht einen Weg zufinden, zu prüfen ob Kernel-Updates verfügbar sind. Im
Netz und in der Man-Page von apt-get bzw aptitude wurde ich nicht eindeutig
fündig. Nach langem suchen ergaben sich allerdings folgende Möglichkeiten
Updates anzeigen zulassen:

```
apt-get --just-print upgrade
```

```
apt-get -s upgrade
```

```
aptitude search ~U
```


Ich fand allerdings die erste Möglichkeit am einleuchtensten.  Die Ausgabe
ist zwar nicht zwingend Script geeignet, aber das lässt sich ja ändern:

    SUPDATEKERNEL=$(ssh root@$host "apt-get --just-print upgrade | grep linux | uniq | wc -l")

In der Variable $SUPDATEKERNEL steht logischerweise die Anzahl der
Verfügbaren Updates die "linux" enthalten. Linux fand ich persönlich jetzt
am einfachsten zur Identifikation von Kernel-Updates. Jemand bessere
Vorschläge?
