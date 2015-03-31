---
date: '2011-04-29 10:30:52'
layout: post
slug: namensliste-in-mysql-datenbank-einspeisen
status: publish
comments: true
title: Namensliste in MySQL Datenbank einspeisen
alias: /archives/1645
categories:
- Bash
- Coding
- Debian
- PlanetenBlogger
- Shell-Zauberei
- SQL
- Web
tags:
- bash
- mysql
- namen
- namensliste
- requesttracker
- shell
- Shell-Zauberei
- sql
- txt
---

**Code**
```
for x in $(mysql --batch -u root --password=passw0rd -e "USE rtdb; SELECT DISTINCT id FROM Users;" | grep -v ^id); do mysql -u root --password=passw0rd -e "USE rtdb; UPDATE Users SET Name="$(sed -n $(($RANDOM % $(cat Names | wc -l) +1 ))p Names)" WHERE id="$x" ;" ; done
```


**Hintergrund**
Wie auch schon beim letzten mal dreht es sich wieder um die Anonymisierung der [RequestTracker](http://bestpractical.com/rt/) Datenbank für die ich zur Zeit an einem [Statistik Tool ](http://github.com/noqqe/RequestTracker-Stats)arbeite. Diesmal will ich aber nicht die EmailAdressen ändern, sondern die Namen der Benutzer. Da ich diese nicht so einfach generieren lassen kann, habe ich mir aus dem Interweb eine [Liste mit Namen](http://www.ta7.de/txt/listen/list0013.htm) besorgt und mit diesen Namen die eingetragenen überschrieben. Jetzt kann ich endlich den Post über das Statistik Tool schreiben und mit Beispielen versehen :)

**Funktion**
Das Ganze läuft wie folgt ab: Für jede ID die ich mittels Datenbank-Verbindung in die For-Schleife einbette, setze ich einen UPDATE Befehl ab, der die Tabelle "Users" und das Feld "Name" aktualisiert. Der Aktualisierungsvorgang passiert aber generisch. Das heisst ich setze den Namen des Users auf einen zufällig ausgewählten neuen Namen aus der Datei "Names". Das habe ich mit sed -n p FILE gelöst, was sicher auch schöner geht, aber für meine Zwecke hat es ausgereicht.

