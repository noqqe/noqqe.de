---
date: 2011-04-29T11:55:33+02:00
layout: post
slug: rt-requesttracker-stats
status: publish
comments: true
title: RT | RequestTracker-Stats
alias: /archives/1649
categories:
- Bash
- Coding
- Debian
- git
- PlanetenBlogger
- SQL
- Ubuntu
- Web
tags:
- Addon
- BestPractical
- modul
- plugin
- requesttracker
- rt
- Statistik
- tool
---

Oft  angesprochen und trotzdem bisher nicht die Zeit gefunden drüber zu bloggen: [RequestTracker-Stats](http://github.com/noqqe/RequestTracker-Stats). Vor kurzem habe ich mir das [Balkendiagramm-Shellskript](/archives/1611) [statistical](http://github.com/noqqe/statistical) gebastelt. Nachdem es so gut funktionierte hatte, hatte ich mir überlegt, was ich damit jetzt anfangen könnte. Ich brauchte einen großem Umfang an Datenmengen, den ich visualisieren konnte (abgesehen von zufällig erzeugten Daten). Am Besten noch etwas, dass Sinn macht :)

{{< figure src="/uploads/2011/04/3927_ede8_550.jpeg" >}}

At Work war "Ticket-Squashing" immer wieder ein gutes Stichwort in unserem Ticketsystem. Wir benutzen den [RequestTracker](http://bestpractical.com/rt/) von [BestPractical](http://bestpractical.com) und ich hatte mir überlegt ein kleines Skript zu basteln, welches die Anzahl der erledigten Tickets pro User aus der MySQL Datenbank ausliesst und dann im Key:Value Format an statistical übergibt. Das hat auch ganz gut funktioniert.



    Resolved ticket statistic for this month (April)
    ---------------------------------------------------
    Tracy		|##################### (22)
    Amelie		|##################### (22)
    Kiri		|############ (13)
    Tersina		|########## (11)
    Birgitta	|######### (10)
    Justine		|######### (10)
    Frank		|######## (9)
    Betteann	|####### (8)
    Cyndy		|# (2)
    Kaleena		|# (2)
    Kiah		|# (2)
    Roxy		| (1)
    Estella		| (1)
    Marj		| (1)



Allerdings haben sich dann im Laufe des Tages immer mehr (ich nenne es mal statistische-) Anwendungsmöglichkeiten ergeben. Zum Beispiel die Anzahl der erstellten Tickets pro Benutzer:



    Most active creators for this month (April)
    ---------------------------------------------------
    Christel@company.com	|############ (13)
    Sydelle@company.com	|########### (12)
    Birgitta@company.com	|######### (10)
    Ainsley@company.com	|######## (9)
    Halette@company.com	|##### (6)
    Martguerita@comp	|##### (6)
    Tracy@company.com	|#### (5)
    care@company.com	|#### (5)
    fooo@company.com	|### (4)
    Christyna@company.com	|## (3)
    Ethel@company.com	|## (3)
    [...]



Oder die Anzahl der Tickets pro Kategorie:



    Queues for this month (April)
    ---------------------------------------------------
    General		|##################### (22)
    Web		|##################### (22)
    Management	|#################### (21)
    WebContent	|################### (20)
    IT-Interal	|################ (17)
    Categ		|############### (16)
    Access		|############### (16)
    E-Mail-Service	|############## (15)
    SWAN		|########### (12)
    Domain-Service	|########## (11)
    Junk		|########## (11)
    DSL		|## (3)
    Other		|# (2)
    Hotspot-Service	|# (2)
    Buchhaltung	| (1)



Um nur ein paar Beispiele zu nennen. Leider waren es zu diesem Zeitpunkt noch etliche separate Skripte, was mir eigentlich nicht gefiel. Darum habe ich es in ein modular aufgebautes Statistik Umgebungstool umgewandelt. Module sind (de-)aktivierbar und lassen sich leicht in das Rahmenprogramm einfügen. Letztendlich gibt es jetzt einen ganzen Satz von Modulen der unter [Github](http://github.com/noqqe/RequestTracker-Stats) zur Verfügung steht.

```
git clone git://github.com/noqqe/RequestTracker-Stats.git
```


Sollte außer uns noch jemand RequestTracker Stats verwenden und Interesse daran haben, ist er herzlich eingeladen die Stats Umgebung zu benutzen ggf. auch Module hinzuzufügen oder zu verbessern :) Eine (ich hoffe doch) ausreichende Anleitung zur Benutzung befindet sich im README des Github Repos.



