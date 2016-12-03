---
title: Statistical Prediction
date: 2013-03-31T15:24:11
tags: 
- Programming
- R
---

### Majority Rule

Daten sind: Regen=1, kein Regen=0

Um die Vorhersage für den nächsten Tag zu treffen wird ein Wert "k" als
Anzahl vorheriger Daten bestimmt. Wenn die Zahl der 1en mindestens k/2 ist,
ist die vorhersage für den nächsten Tag auf jedenfall 1, anderenfalls 0

Wenn also 4/2 = 2 (es hat an zwei Tagen geregnet), wird es am 5. Tag wieder regnen.

Um Herauszubekommen wie man "k" am besten setzt, wird einfach durch das
vorherige Dataset durchgegangen und geschaut, bei welchem "k" Wert es die
besten Ergebnisse gab.

