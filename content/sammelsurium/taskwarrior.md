---
title: taskwarrior - Recurring Tasks
date: 2013-11-04T08:46:53
tags: 
- Software
- taskwarrior
---

Die Kunst liegt eigentlich nur dabei, relative Datums zu wählen.
Wurde aber alles in späteren Versionen nochmal überarbeitet.

Jede Woche am Donnerstag öffnen, bis due: Freitag

    task add pro:work pri:M due:friday recur:weekly wait:thursday Stunden eintragen

Jeden Monat am ersten öffnen, bis 2. fällig:

    task add pro:work pri:H due:2nd recur:monthly wait:1st Stundenabschluss machen

Jedes halbe Jahr Updates

    task add pro:updates pri:H due:20th recur:semiannual wait:25th Gitolite
