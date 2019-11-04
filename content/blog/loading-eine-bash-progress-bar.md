---
aliases:
- /blog/2011/11/26/loading-eine-bash-progress-bar
- /archives/1819
comments:
- author: no bo dy
  content: <p>Wieso komme ich nur immer auf solche Sachen? Aber:</p><p>./<a href="http://loading.sh"
    rel="nofollow">loading.sh</a> 5 0.05 "    8" = D</p>
  date: '2011-11-26T21:52:50'
- author: nuit
  content: <p>gibt es bereits :)</p><p>zumindest in arch gibt es im community repository
    ein bash script namens "bar".<br><a href="http://pastebin.com/GzRgiWzF" rel="nofollow">http://pastebin.com/GzRgiWzF</a></p><p>aber
    nette idee und schoen ausgefuehrt :)</p>
  date: '2011-11-28T08:35:11'
- author: noqqe
  content: '<p>@no bo dy: ;D na klar. </p><p>@nuitmh ich kann das script von arch
    irgendwie nicht benutzen. Bekomme Syntax Errors.</p>'
  date: '2011-11-29T00:52:23'
date: '2011-11-26T12:53:21'
tags:
- development
- web
- shell
- bar
- balken
- fortschrittsbalken
- pv
- ladebalken
- simple
- einfach
- loading bar
- laden
- linux
- progress
- progressbar
- debian
- bash
title: Loading | Eine Bash-Progress-Bar
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
