---
comments:
- author: zogan
  content: "Scrum ist nicht esoterisch, sondern eher best-practice\u2122. Und ganz
    dumm ist die Idee des Planning-Poker nicht, zumindest _wenn_ man aufgrund von
    Projektleitungsbissnessfoo gezwungen ist, Sch\xE4tzungen abzugeben. Erschlie\xDFt
    sich einem aber auch erst, wenn man mal Software in gr\xF6\xDFeren Teams entwickelt."
  date: '2014-02-10T16:46:58.705967'
- author: noqqe
  content: "\"Sehr tolles Prinzip.\" - H\xF6rt sich zwar ironisch gemeint an, wars
    in dem Fall aber garnicht :) \n\nDas meinte ich zur Abwechslung mal ernst ;) Sorry"
  date: '2014-02-10T16:56:17.040586'
- author: zogan
  content: Oh. :-)
  date: '2014-02-15T14:33:46.672552'
date: '2014-02-01T15:50:00'
tags:
- development
- heisenbug
- wikipedia
- entwicklung
- stonith
- unix
- planning poker
- airplane
- ha
- software
title: Obskures aus der Software Entwicklung
---

### Heisenbug

{{< figure src="/uploads/2014/02/heisenberg.jpg" >}}

Aus der aktuellen Situation heraus suggerierte mir "[Heisenbug](https://en.wikipedia.org/wiki/Heisenbug)" eher
einen Methamphetamin kochenden Walter White. Bei näherer Betrachtung geht es aber
doch um den Wissenschafter [Werner Heisenberg](https://en.wikipedia.org/wiki/Werner_Heisenberg).
Abgeleitet von der [Heisenbergschen Unschärferelation](https://de.wikipedia.org/wiki/Heisenbergsche_Unsch%C3%A4rferelation)
besagt Heisenbug, dass sobald Debugging-Methoden ergriffen werden, der Bug im
Programm nicht mehr nachvollziehbar wird.

Auch andere Wissenschaftler bekommen ihre Referenzen. Siehe [Bohrbug](http://www.catb.org/jargon/html/B/Bohr-bug.html),
[Mandelbug](http://www.catb.org/jargon/html/M/mandelbug.html) und [Schrödinbug](http://www.catb.org/jargon/html/S/schroedinbug.html)

### STONITH

[STONITH](https://en.wikipedia.org/wiki/STONITH) hatte ich bereits
[vertwittert](https://twitter.com/noqqe/status/421178550012874752).

> Shoot The Other Node In The Head

Ein Konzept das bei hochverfügbaren
Setups das Verhalten anderer Nodes im Fehlerfall eines Nachbarn beschreibt. Und das sehr
unmissverständlich.

### It is Easier to Ask for Forgiveness than Permission

Dieses Motto wird folgender guten Dame zugeschrieben.

{{< figure src="/uploads/2014/02/gracehopper.jpg" >}}

[Grace Hopper](https://en.wikipedia.org/wiki/Grace_Hopper), die auch den Spitznamen "Grandma COBOL" trägt. Exception-Handling:

``` python
try:
    ham = spam.eggs
except AttributeError:
    handle_error()
```

...anstelle von if-Conditions:

``` python
if hasattr(spam, 'eggs'):
    ham = spam.eggs
else:
    handle_error()
```

Der Coding-Stil floss in allerlei Sprachen ein.
Unter anderem [Python](https://en.wikipedia.org/wiki/Python_syntax_and_semantics#Exceptions).

### Shotgun Debugging

Das sogenannte "[Shotgun Debugging](https://en.wikipedia.org/wiki/Shotgun_debugging)" ist
eine (zurecht) verkannte Methodik mit Bugs fertig zu werden. Statt
strukturellem Vorgehen werden wild mit der Präzision einer Schrotflinte
Einstellungen/Codepassagen geändert ohne deren Auswirkung vorherzusehen.

### The Mythical Man-Month

Wer in der IT sein Geld verdient, wird das kennen. Verzögerungen bei einem Projekt
sind nicht selten. Zusätzliche Entwickler zum richtigen
Zeitpunkt zum Projekt holen dagegen sehr.

> Adding manpower to a late software project makes it later.

Darum gehts in dem Buch [The Mythical Man-Month](https://en.wikipedia.org/wiki/The_Mythical_Man-Month).

### Planning Poker

Ein Kollege klärte mich diese Woche über [Planning Poker](https://en.wikipedia.org/wiki/Planning_poker) auf.
Ein typischer Kandidat aus dem fast schon esoterischen Umfeld von
[Scrum](https://en.wikipedia.org/wiki/Scrum_\(development\)) und
anderen bewusstseinserweiterden Optimierungsmethoden im Software Development.

{{< figure src="/uploads/2014/02/planningpoker.jpg" >}}

Kurz zusammengefasst, geht es darum Tasks mit Karten (von 1-100) nach eigenem Ermessen auf ihre
Schwierigkeit zu bewerten. So entstehen Working-Slots und Arbeitsverteilung, da
man sich selbst einschätzen kann. A la "In einer Woche kann ich ca. 70 Punkte
bewältigen". Egal ob das nun 7x 10er Probleme oder 1x 70 sind. Sehr tolles
Prinzip.

### The Airplane Rule

Die [Flugzeug Regel](http://www.catb.org/jargon/html/A/airplane-rule.html)
ist eine Versinnbildlichung für das [KISS Prinzip](https://en.wikipedia.org/wiki/KISS_principle)

> A twin-engine airplane has twice as many engine problems as a single-engine airplane

Quellen:
[Werner Heisenberg](https://en.wikipedia.org/wiki/File:Bundesarchiv_Bild183-R57262,_Werner_Heisenberg.jpg),
[Grace Hopper](https://en.wikipedia.org/wiki/File:Grace_Hopper_and_UNIVAC.jpg),
[Planning Poker Cards](https://en.wikipedia.org/wiki/File:CrispPlanningPokerDeck.jpg),
[List of software development philosophies](https://en.wikipedia.org/wiki/List_of_software_development_philosophies),
[Hackers Dictonary](http://www.catb.org/jargon/) von Eric Raymond
