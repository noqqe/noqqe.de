---
type: post
title: "Paying 5 bucks a month for stupid statistics?"
date: 2012-07-02T21:00:00+02:00
comments: true
categories:
- planetenblogger
- Debian
- Shell
- Web
- Stats
tags:
- Sport
- runkeeper
- gnuplot
- barchart
- linechart
- plot
- bars
- stats
---

Seit ich angefangen habe Ausflüge durch Wälder und anderem grobem Gelände mit
meinem Mountainbike zu machen ist [Runkeeper](http://runkeeper.com) mein Begleiter.

Mit der iPhone App kann ich allerhand Daten tracken wie

* Durchschnittliche Geschwindigkeit
* Höhenmeter
* Kalorien (wenn man deren Berechnungen glauben mag)
* Karten
* Länge der Fahrt

Natürlich will auch Runkeeper irgendwie Geld verdienen und
somit gibt es was Accounts angeht auch noch einen "Elite Pro Wahnsinns" Account
welcher irgendwie 5-8 Euro im Monat kostet und mir als Ausgleich bessere
Statistiken und sogenannte "Fitness Reports" liefert.

Ehrlichgesagt war ich davon nicht sehr angetan so viel Geld für ein paar extra
Skripte zu bezahlen. Deshalb hab ich mich entschlossen meine eigenen Graphen mit
[GnuPlot](http://gnuplot.org) zu malen und irgendwie zu einer Website zu machen.

## runkeeper-statistics

Im Grunde liefert mir Runkeeper alle Rohdaten die ich brauche. Sie weigern sich
nur mir die Statistiken daraus zu bauen die gerne hätte. Ich habe mir dann ein
Shell Script gebaut, welches gnuplot mit Daten füttert, die in einem CSV File
spezifiziert sind.

{{< figure src="/uploads/2012/07/calories-activity.png" >}}

Zuerst schreibe ich alle meine Daten in die activity.dat. Ziemlich im CSV Stil,
aber mir fiel nicht blöderes ein ohne SQL sprechen zu müssen.

``` csv
# Add your data here to generate awesome graphs ;)
# ID, Date,      Distance, Duration, Pace, Speed, Burned, Climb
  1,  2012-06-23, 14.07,    0.54,     3.54, 15.38, 458,    156
  2,  2012-06-26, 16.28,    1.09,     4.17, 14.03, 582,    292
  3,  2012-06-28, 17.65,    1.13,     4.11, 14.36, 618,    242
  4,  2012-06-30, 27.28,    1.47,     3.56, 15.24, 876,    379
```

Danach baut mir das Shellscript die GnuPlots und die HTML Site

```
$ ./runkeeper-statistics
```

Das alles gibts auf [github](http://github.com/noqqe/runkeeper-statistics/) und
die "Demo" hab ich mal hochgeladen und mit zufälligen Daten befüllt:<b>[Demo](/uploads/2012/07/runkeeper-statistics/html/index.html)</b>
