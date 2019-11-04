---
comments:
- author: Das wort
  content: "<p>Also jetzt noch ein paar mehr Funktionen und ab ins Debian-Repo *tr\xE4um*</p>"
  date: '2012-04-11T02:10:13'
- author: noqqe
  content: "<p>Hehe :) Bin f\xFCr Pull Requests immer offen :) </p>"
  date: '2012-04-11T05:49:36'
date: '2012-04-01T09:11:00'
tags:
- development
- web
- shell
- code
- graph
- plot
- chart
- terminal
- statistics
- ascii
- bash
- statistik
title: ASCII Statistik Graphen mit Bash
---

[Vor einem Jahr](/blog/2011/04/14/statistical-statistiken-visualisieren-im-terminal/)
hab ich mich mit einem ASCII Balkendiagramm beschäftigt. Als es dann fertig war
kams unter dem Namen statistical auf Github:
[http://github.com/noqqe/statistical](http://github.com/noqqe/statistical)

{{< figure src="/uploads/2012/04/statistics-fail.jpg" >}}

Hab mir überlegt wie schwierig das wohl ist, dass ganze Teil nochmal so
um zubauen, dass vertikale Balken entstehen. Horizontal gesehen ist das ja
relativ easy. Etwas skalieren hier und da und dann einfach die Anzahl der
Zeichen in einer Zeile ausgeben.

Im v-barplot ist da schon etwas mehr Logik nötig. Aber es ging dann trotzdem. Heraus kam dabei folgendes:

    ./v-barplot foo:24 bar:90 alice:76 flo:32 blu:79 fa:12 bob:230 kika:121

{{< figure src="/uploads/2012/04/v-barplot.png" >}}

Beide Skripte (v-barplot & h-barplot) sind jetzt im statistical Repo. Anregungen
& Kritik immer gerne.
