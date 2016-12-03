---
aliases:
- /archives/894
- /blog/2010/02/20/apt-aktualisierungen-anzeigen-lassen
date: '2010-02-20T16:27:38'
tags:
- development
- kernel
- search
- aktualisierungen
- anzeigen
- show
- list
- update
- apt
- upgrade
- updates
- aptitude
- linux
- apt-get
- debian
- ubuntu
title: Apt | Aktualisierungen anzeigen lassen
---

Für ein kleines Script, welches ich auf der Arbeit verwende, habe ich
versucht einen Weg zu finden, zu prüfen ob Kernel-Updates verfügbar sind. Im
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


Ich fand allerdings die erste Möglichkeit am Einleuchtensten.  Die Ausgabe
ist zwar nicht zwingend Script geeignet, aber das lässt sich ja ändern:

    SUPDATEKERNEL=$(ssh root@$host "apt-get --just-print upgrade | grep linux | uniq | wc -l")

In der Variable $SUPDATEKERNEL steht logischerweise die Anzahl der
Verfügbaren Updates die "linux" enthalten. Linux fand ich persönlich jetzt
am einfachsten zur Identifikation von Kernel-Updates. Jemand bessere
Vorschläge?