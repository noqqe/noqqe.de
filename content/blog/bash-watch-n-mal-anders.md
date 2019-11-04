---
aliases:
- /blog/2010/09/05/bash-watch-n-mal-anders
- /archives/1230
date: '2010-09-05T18:58:33'
tags:
- development
- linux
title: Bash | watch -n mal anders
---

```
watch -n 1 "arp -a | grep 192.168.1.12"
```

[watch](http://linux.about.com/library/cmd/blcmdl1_watch.htm) hat die blöde
Angewohnheit, den aktuellen Screen immer zu leeren. Gerade beim oberen
Beispiel ist das von Nachteil (IP-Konflikt nachvollziehen), da der
Vergleichswert weg fällt. Quick&Dirty die Lösung:

```
while true; do arp -a | grep 192.168.1.12 ; sleep 0.2 ; done
\________/ |  ____/  ________________/   _______/  ____/
     |      |     |             |               |         |
     |      |     |             |               |         - Schleifenen
     |      |     |             |               |            de
     |      |     |             |               |
     |      |     |             |               - danach 0.2 Sekunden
     |      |     |             |                  pausieren
     |      |     |             |
     |      |     |             - nach bestimmter Adresse suchen
     |      |     |
     |      |     - arp-table ausgeben
     |      |
     |      - fuehre aus
     |
     - bis strg+c
```
