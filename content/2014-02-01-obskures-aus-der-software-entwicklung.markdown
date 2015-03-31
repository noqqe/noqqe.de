---
layout: post
title: "Obskures aus der Software Entwicklung"
date: 2014-02-01T17:50:00+02:00
comments: true
categories:
- Software Development
- Software
- Entwicklung
- osbn
- ubuntuusers
- heisenbug
- stonith
- planning poker
- airplane
- unix
- ha
- wikipedia
---

### Heisenbug

{% img right /uploads/2014/02/heisenberg.jpg %}

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

{% blockquote %}
Shoot The Other Node In The Head
{% endblockquote %}

Ein Konzept das bei hochverfügbaren
Setups das Verhalten anderer Nodes im Fehlerfall eines Nachbarn beschreibt. Und das sehr
unmissverständlich.

### It is Easier to Ask for Forgiveness than Permission

Dieses Motto wird folgender guten Dame zugeschrieben.

{% img center /uploads/2014/02/gracehopper.jpg %}

[Grace Hopper](https://en.wikipedia.org/wiki/Grace_Hopper), die auch den Spitznamen "Grandma COBOL" trägt. Exception-Handling:

{% codeblock lang:python %}
try:
    ham = spam.eggs
except AttributeError:
    handle_error()
{% endcodeblock %}

...anstelle von if-Conditions:

{% codeblock lang:python %}
if hasattr(spam, 'eggs'):
    ham = spam.eggs
else:
    handle_error()
{% endcodeblock %}

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

{% blockquote %}
Adding manpower to a late software project makes it later.
{% endblockquote %}

Darum gehts in dem Buch [The Mythical Man-Month](https://en.wikipedia.org/wiki/The_Mythical_Man-Month).

### Planning Poker

Ein Kollege klärte mich diese Woche über [Planning Poker](https://en.wikipedia.org/wiki/Planning_poker) auf.
Ein typischer Kandidat aus dem fast schon esoterischen Umfeld von
[Scrum](https://en.wikipedia.org/wiki/Scrum_\(development\)) und
anderen bewusstseinserweiterden Optimierungsmethoden im Software Development.

{% img center /uploads/2014/02/planningpoker.jpg %}

Kurz zusammengefasst, geht es darum Tasks mit Karten (von 1-100) nach eigenem Ermessen auf ihre
Schwierigkeit zu bewerten. So entstehen Working-Slots und Arbeitsverteilung, da
man sich selbst einschätzen kann. A la "In einer Woche kann ich ca. 70 Punkte
bewältigen". Egal ob das nun 7x 10er Probleme oder 1x 70 sind. Sehr tolles
Prinzip.

### The Airplane Rule

Die [Flugzeug Regel](http://www.catb.org/jargon/html/A/airplane-rule.html)
ist eine Versinnbildlichung für das [KISS Prinzip](https://en.wikipedia.org/wiki/KISS_principle)

{% blockquote %}
A twin-engine airplane has twice as many engine problems as a single-engine airplane
{% endblockquote %}


Quellen:
[Werner Heisenberg](https://en.wikipedia.org/wiki/File:Bundesarchiv_Bild183-R57262,_Werner_Heisenberg.jpg),
[Grace Hopper](https://en.wikipedia.org/wiki/File:Grace_Hopper_and_UNIVAC.jpg),
[Planning Poker Cards](https://en.wikipedia.org/wiki/File:CrispPlanningPokerDeck.jpg),
[List of software development philosophies](https://en.wikipedia.org/wiki/List_of_software_development_philosophies),
[Hackers Dictonary](http://www.catb.org/jargon/) von Eric Raymond
