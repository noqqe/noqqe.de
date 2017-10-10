---
aliases:
- /blog/2012/05/28/apache-2-dot-4-und-mpm-event
- /blog/2012/05/28/apache-2-dot-4-und-mpm-event/
comments:
- author: Higgi
  content: "<p>Ich glaube du meinst bei 600 gleichen Requests oder irre ich mich da
    ?\_</p><p>\"Warum MPM Prefork bei 1000 gleichzeitigen Connections immer so \u201Cabhaut\u201D
    kann ich nicht sagen\"</p>"
  date: '2012-05-28T17:15:45'
- author: noqqe
  content: "<p>Teils... das der Event bei 600 abhaut ist auch komisch... aber geschrieben
    hab ich eigentlich \xFCber den Prefork bei 1000. :)</p>"
  date: '2012-05-28T17:26:35'
- author: Higgi
  content: <p>ups ;) Wer lesen kann ist stark im Vorteil ;) Ich habs ja sogar zitiert.</p>
  date: '2012-05-29T05:07:49'
- author: noqqe
  content: <p>hehe :)</p>
  date: '2012-05-29T05:28:06'
- author: StackLinux
  content: <p>Beautiful Thanks! I'm on 2.4 event.c and flying.</p>
  date: '2013-07-31T06:25:33'
date: '2012-05-28T12:41:00'
tags:
- web
- eventbased
- worker
- apachebench
- benchmark
- event mpm
- mpm
- administration
- apache2
- prefork
- apache
- ab
- eventbasiert
- mpm_event
- debian
title: Apache 2.4 und mpm_event
---

[Apache 2.4](http://httpd.apache.org/docs/2.4/) ist seit einiger Zeit in
[Debian Experimental](http://www.debian.org/releases/experimental/)
paktiert. Yeah!

Grade das eventbasierte MPM
[mpm_event](http://httpd.apache.org/docs/2.4/mod/event.html), dass in 2.4
stable wird hat mich daran interessiert. Zeit mal wieder Graphen zu bauen
und Benchmarks anzustoßen.

## Das Setup

Getestet werden:

* MPM prefork Apache2 (2.2, Debian Stable)
* MPM worker Apache2 (2.2, Debian Stable)
* MPM event Apache2 (2.4, Debian Experimental)

Um möglichst genaue Daten zu bekommen hab ich keine Maschine gewählt auf der
ich per Datenkabel erst hin muss um so leitungsverschuldete Messfehler zu
vermeiden. Für die Tests musste jeweils ein LinuXContainer auf meinem Thinkpad
herhalten.

Eckdaten des Tests:

* Debian Minimal Installation
* 4x Prozessoren (i5)
* 4GB RAM
* Netzwerk über eine lokale Ethernet Bridge
* Auslieferung der Apache2 Default HTML Page
* Konfigurationen der MPMs alle Paket default behalten
* Alle Benchmarks sind mit `ab` durchgeführt. Dieses Apache Benchmark Tool wird
mit dem Paket `apache2-utils` ausgeliefert.

## Max Requests

Ein typischer Apache Bench mit `ab` sieht ungefähr so aus:

```
$ ab -c 200 -n 2500 http://host.example.org/
```

* `-c` gibt die Anzahl der gleichzeitigen Verbindungen an
* `-n` Die Anzahl der Connections insgesamt

Jedem Apache habe ich nun stufenweise Connections hin geschossen
und mir die Dauer jedes Benchmarks weggegreppt:

``` bash
C=0
while [ $C -lt 100000 ]; do
  C=$((C + 10000))
  echo -n "$C:"
  ab -c 500 -n $C http://10.10.0.16/ 2>/dev/null | grep "Time taken for tests:" | awk '{print $5}'
done
```

Das hab ich für jeden Host einzeln durchgeführt und anschliessend alles durch
Gnuplot gejagt. Dabei kam einmal mehr zum Vorschein, dass man Prefork einfach
nicht haben möchte :)

{{< figure src="/uploads/2012/05/max-requests.png" >}}

## Max Concurrency

Natürlich haben die verschiedenen Modelle verschiedene Eigenschaften die
sie für bestimmte Aufgaben tauglicher macht als andere. Deshalb hab ich auch noch
von 100-1000 gleichzeitigen Requests geprüft. Der Einzeiler ist dafür nur
geringfügig modifiziert:

``` bash
C=0
while [ $C -lt 1000 ]; do
  C=$((C + 100))
  echo -n "$C:"
  ab -c $C -n $C http://10.10.0.16/ 2>/dev/null | grep "Time taken for tests:" | awk '{print $5}'
done
```

Das Ergebnis von MPM Event kann sich sehen lassen, wie ich finde.

{{< figure src="/uploads/2012/05/concurrency.png" >}}

Warum MPM Prefork bei 1000 gleichzeitigen Connections immer so "abhaut" kann ich
nicht sagen. Habe mehrmals versucht mit immer dem gleichen Ergebnis.

## Fazit

Die Daten zu deuten ist jedem selber überlassen. Was hier auch überhaupt nicht
zur Sprache kam ist die Administrierbarkeit der Module oder Stabilität. Ob man
Prefork mag weil es nativ mit mod_php kann oder man sich wegen der Performance
mpm_worker mit fcgid antut... Geschmäcker gehen hier auseinander aber allein
wegen der Verträglichkeit des mpm_event gegenüber
[Slowloris](http://de.wikipedia.org/wiki/Slowloris) Attacken sollten man sich
den "neuen" mpm_event jedenfalls mal ansehen. Performacetechnisch sieht es
jedenfalls nicht schlecht aus ;)

Wer möchte kann sich die [Rohdaten](https://gist.github.com/2764231) der Tests
bei Github abholen.

Den Source zu den Plots gibts auf Nachfrage ebenfalls.
