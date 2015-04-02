---
date: 2010-09-05T20:58:33+02:00
type: post
slug: bash-watch-n-mal-anders
comments: true
title: Bash | watch -n mal anders
aliases:
- /archives/1230
categories:
- Coding
- Linux
---

```
watch -n 1 "arp -a | grep 192.168.1.12"
```


[watch](http://linux.about.com/library/cmd/blcmdl1_watch.htm) hat die blöde Angewohnheit, den aktuellen Screen immer zu leeren. Gerade beim oberen Beispiel ist das von Nachteil (IP-Konflikt nachvollziehen), da der Vergleichswert wegfällt. Quick&Dirty die Lösung:


    while true; do arp -a | grep 192.168.1.12 ; sleep 0.2 ; done
    _________/ |  ____/  ________________/   _______/  ____/
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
