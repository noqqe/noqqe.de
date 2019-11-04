---
aliases:
- /archives/1635
date: '2011-04-23T16:53:04'
slug: shell-zauberei-benutzer-emailadressen-in-mysql-datenbank-anonymisieren
tags:
- development
- rt
- shell
- onliner
- databases
- einzeiler
- mysql
- requesttracker
- bash
title: Emailadressen in MySQL Datenbank anonymisieren
---

### Code

```
for x in $(mysql -u root --password=passw0rd --batch -e "use rtdb; select id from Users" | grep -v ^id); do mysql -u root --password=passw0rd -e "use rtdb; UPDATE Users SET EmailAddress="$x@mail.com" WHERE id="$x"; "; done
```

### Hintergrund**

Ich arbeite seit kurzem an einem kleinen Statistik Tool für den
[RequestTracker](http://bestpractical.com/rt/) von Bestpractical. Einem
Ticket-System. Um Beispiele für dieses Tool generieren zu können, brauchte
ich eine manipulierte Datenbank. Datenschutz. EmailAdressen mit anderem
Inhalt überschreiben.

### Funktion

Im Endeffekt ist es nur eine For-Schleife, die alle User ID's aus der
RequestTracker Datenbank (Table: Users) ausliesst und für jeden gefundenen
Eintrag die EmailAdresse auf "ID@mail.com" setzt. So bleiben die Daten
auswertbar, sind aber "anonym".
