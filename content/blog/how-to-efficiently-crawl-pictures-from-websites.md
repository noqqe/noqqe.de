---
date: '2014-09-17T20:41:00'
tags:
- development
- web
- reddit
- pr0gramm
- pictures
- soupio
- nichtparasoup
- images
- crawl
- crawler
title: How to do simple and efficient image crawling
---

Im Rahmen von [nichtparasoup](http://github.com/k4cg/nichtparasoup) ergab
es sich, dass wir uns Gedanken über crawling machen mussten.

{{< figure src="/uploads/2014/09/papers.gif" >}}

Bei nichtparasoup sollen zufällig Bilder aus dem Netz wiedergegeben werden
um wenn möglich den optimalen Unterhaltungswert für den "Zuschauer"
darzustellen. Dabei galt es vier Probleme zu lösen:

* Möglichst aktuelle Bilder (kein alter Schrott)
* Niemals ohne Content dastehen
* Keine Duplicates
* Möglichst effizientes Crawling

Wie also diese Kriterien am Besten vereinen?  Bei allen Quellen die benutzt
werden gibt es eine Art "paging". Also eine fixe Anzahl von Bildern, die
pro Request aus der Quelle ausgespuckt werden.

### 1. Ein Buch lesen

Wir begannen, die Seiten einfach von Anfang bis Ende durchzuscrollen. Seite
merken. Wenn alle Bilder gesehen wurden, an gemerkter Stelle weitermachen.
Effizient. Aber neuer Content wurde damit komplett ignoriert.

### 2. Yo Dawg, i heard you like HTTP-GETs

Um neuen Content mitzubekommen, haben wir nach dem ersten Versuch, die
gerade durchsuchte Quelle jedes mal wenn wir Bilder brauchten von vorne
durchzusuchen. Alle Bilder wurden gleichzeitig zu zwei Listen hinzugefügt.
Zur ImageMap und zu einer Blacklist.  Die Blacklist hatte den Charme, dass
wir ausschliesslich Bilder in die ImageMap aufnahmen, die noch nicht auf
der Blacklist waren.  Gab es also auf der ersten Seite (neueste Einträge)
nicht genügend Bilder um die ImageMap zu füllen, ging der Crawlvorgang so
lange weiter, bis man auf unbekannte Bilder stößt. Diese befanden sich
jedoch meistens am Ende.

Dem Erfahrenen-Internet-Cralwer wird aufgefallen sein, dass das nicht
wirklich effizient ist.  Umso länger die nichtparasoup-Instanz läuft, umso
größer wird der Gap zwischen dem "neuen" Content und der letzten bekannten
Seite. Ineffizient.

### 3. Lesezeichen

Eine erweiterte Version der zweiten Methodik war es, bei jeden Crawlvorgang
die erste Seite zu Crawlen und danach an der letzten bekannten Stelle im
"Buch-index" zu springen und dort weiter nach unbekannten Bildern zu
suchen.

Aber was ist, wenn die Instanz lange nicht besucht wird? Wenn es auf den
Index-Seiten 2,3,4,5,6 bereits neuen Content gibt den wir noch garnicht
kennen? Wie können wir wissen bis an welche Stelle neuer Content ist?

### 4. Short-Term-Memory-Loss

Die Lesezeichen Methodik hat ganz gut funktioniert, bis auf die Schwäche
des neuen Contents auf den Seiten 2,3,4,5,6.

Nachfolgendes, kann man sich wie einen Rentner ohne Lesezeichen vorstellen.
Er versucht sich die Seite, auf der er im Buch stehengeblieben ist zu
merken. "Wo war ich nochmal stehen geblieben?" ... "Fang ich das Buch halt
nochmal von vorne an."

Um auch neuen Content gecrawlt zu bekommen, wird das Lesezeichen, welches
wir bei jedem Durchlauf pro Quelle setzen, alle 3 Stunden "vergessen".
Alles fängt wieder von vorne an. Wir müssen nicht zu viel Bookkeeping
betreiben und verballern auch nicht unendlich viele Requests. Die
Implementation war relativ einfach.

{{< figure src="/uploads/2014/09/notimpressed.gif" >}}

Ich bin mir ziemlich sicher, dass es hier akademische Beleuchtungen noch
und nöcher gibt, schlaue Menschen bei allen Möglichen Suchmaschinen sich
Köpfe zerbrochen haben und entsprechende Papers veröffentlicht haben. Für
unsere Zwecke eignet sich diese Vorgehensweise aber sehr gut. Zumindest was
die Erfahrungswerte der letzten 3 Monate angeht. Für bessere Vorschläge
sind wir aber sehr gern zu haben.