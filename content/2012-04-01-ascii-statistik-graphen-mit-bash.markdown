---
layout: post
title: "ASCII Statistik Graphen mit Bash"
date: 2012-04-01T11:11:00+02:00
comments: true
categories:
- ubuntuusers
- Bash
- Web
- Coding
tags:
- ascii
- statistics
- statistik
- bash
- code
- graph
- chart
- plot
- terminal
---

[Vor einem Jahr](/blog/2011/04/14/statistical-statistiken-visualisieren-im-terminal/)
hab ich mich mit einem ASCII Balkendiagramm beschäftigt. Als es dann fertig war
kams unter dem Namen statistical auf Github:
[http://github.com/noqqe/statistical](http://github.com/noqqe/statistical)

{% img center /uploads/2012/04/statistics-fail.jpg %}

Hab mir überlegt wie schwierig das wohl ist, dass ganze Teil nochmal so
umzubauen, dass vertikale Balken entstehen. Horizontal gesehen ist das ja
relativ easy. Etwas skalieren hier und da und dann einfach die Anzahl der
Zeichen in einer Zeile ausgeben.

Im v-barplot ist da schon etwas mehr Logik nötig. Aber es ging dann trotzdem. Heraus kam dabei folgendes:

    ./v-barplot foo:24 bar:90 alice:76 flo:32 blu:79 fa:12 bob:230 kika:121

{% img center /uploads/2012/04/v-barplot.png %}

Beide Skripte (v-barplot & h-barplot) sind jetzt im statistical Repo. Anregungen
& Kritik immer gerne.
