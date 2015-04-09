---
date: 2011-11-26T14:53:21+02:00
type: post
comments: true
title: Loading | Eine Bash-Progress-Bar
aliases:
- /blog/2011/11/26/loading-eine-bash-progress-bar
- /archives/1819
categories:
- Shell
- Development
- Debian
- Linux
- ubuntuusers
- Web
tags:
- balken
- bar
- bash
- einfach
- fortschrittsbalken
- ladebalken
- laden
- loading bar
- progress
- progressbar
- pv
- shell
- simple
---

Für ein kleines Projekt, an dem ich so nebenher immer etwas schreibe habe
ich eine Art Ladebalken gebraucht. Habe ein paar wirklich coole
Lösungsansätze gefunden, aber es läuft meistens auf Depencies raus (pv
z.B.) oder nicht wirklich mein Anwendungsfall.

{{< figure src="/uploads/2011/11/not-sure-if-loading.png" >}}

Ich hab mir dann kurzerhand was selber gebastelt. Ich gebe zu ich hätte es
auch so gestaltet können das es einfach nur für meinen Use-Case gereicht
hätte, aber das erschien mir unsinnig. Wenn ich mich schon einen halben
Abend hinsetze, dann können ja evtl. auch mehr Menschen was davon haben. So
entstand dann die
[bash-progress-bar](https://github.com/noqqe/bash-progress-bar).

Zu allererst besteht der Ladebalken aus einer while true Schleife. Sollte
die Bar in ein Skript einbaut werden wäre die Bedingung dem Skript
anzupassen. Ob das jetzt ein test -e auf ein File ist das getouched wird
oder eine Art Counter bleibt jedem selbst überlassen.


    $ git clone git://github.com/noqqe/bash-progress-bar.git
    $ cd /bash-progress-bar/
    $ ./loading.sh
    > [            #####       ]


Alle Parameter sind natürlich anpassbar. Ich habe versucht so gut wie alles
anpassbar zu halten. Ich hoffe das ist mir gelungen ;)

```
./loading.sh Groesse Geschwindigkeit Rahmen-Anfang Füllcharacter Rahmen-Ende
./loading.sh 50 0.02 [ "######" ]
```

Ohne irgendwie ein GIF-File zu erstellen kann ich das jetzt leider schlecht
im Blog demonstrieren. Deshalb: ausprobieren :) Mehr Infos auf der Github
Page.

Fragen, Anregungen, Kritik erwünscht!
