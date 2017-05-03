---
date: 2017-05-03T17:56:34+02:00
tags:
- Strava
- GPX
- R
- Map
title: Strava GPS Tracks analysieren in R
---

In der Recherche zu dem Thema, ist mir aufgefallen das ich scheinbar einen
Hang zu selbstgemachten Graphen habe. Siehe
[Paying 5 bucks a month for stupid statistics?](https://noqqe.de/blog/2012/07/02/paying-5-bucks-a-month-for-stupid-statistics/) oder
[Photos auf Google Maps](https://noqqe.de/blog/2015/08/09/photos-auf-google-maps/).

Wie auch immer. Auch dieses mal bietet mir mein Tracking Tool
[Strava](https://strava.com) nicht die Möglichkeit all meine gefahrenen
Strecken auf einer Karte zu visualisieren. Problem: Das hätte ich gerne.

Nach einigem Wälzen im Netz nach einer Art HTML/JS Library die mir das
interaktiv im Browser macht, fand ich dann einen Blogpost der mich
ebenfalls
interessierte. [Visualizing gps data in R](http://dangoldin.com/2014/02/05/visualizing-gps-data-in-r/). `R`, mh.
Schon lange nicht mehr angefasst!

Ich meldete mich bei Strava an und forderte unter
[Settings](https://www.strava.com/settings/profile) meine GPS Tracks zum
Download an.

Danach spielte ich etwas mit dem Script herum und heraus kam das.

{{< figure src="/uploads/2017/05/viz1.png" >}}

Das war schonmal eine tolle Sache. Vor allem weil in `R` die Visualisierung
so schön war. Voller Stolz hier ein jiff^Wgif des Vorgangs. Toll ist, dass
die Tracks auch noch chronologisch aufgezeichnet werden.

{{< figure src="/uploads/2017/05/viz2.gif" >}}

Das fand ich zwar jetzt alles extrem cool, aber irgendwie wollte ich noch
ein bisschen mehr mit Karte im Hintergrund. Ich wollte das man auch die
Karte sieht! Ich fand einen weiteren Blogpost zu R und GPX:
[Running paths in Amsterdam](https://www.visualcinnamon.com/2014/03/running-paths-in-amsterdam-step-2.html)

Kurzer hand hab ich beide Varianten irgendwie zusammengeworfen und heraus
kam dabei folgendes.

{{< figure src="/uploads/2017/05/viz3.png" >}}

Wen das Skript interessiert, das gibts
[hier](https://gist.github.com/noqqe/5a780a2132b781a01bb0debed7765ab9). Im
Grunde hat das ganze eigentlich auch garnichts mit Strava zu tun, da man
wahrscheinlich beliebige `gpx` Dateien reinschmeißen kann.
