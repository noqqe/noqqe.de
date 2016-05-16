---
categories:
- code
- crypto
- development
- osbn
date: 2016-05-16T11:13:02+02:00
tags:
- rvo
- python
- cmddocs
title: rvo
---

Habe ja nun schon oft angedeutet, dass ich an etwas bastle, was langfristig
`cmddocs` ablösen wird. An diesem Stück Software habe ich sehr lange
gebastelt. Fast 5 Monate. Es war aber nicht das Coding, viel mehr das Konzept.
Im Nachhinein betrachtet hätte ich mir im Vorfeld mehr Gedanken machen
sollen und danach mit der Implementierung beginnen.

Vor einer Woche hab ich es auf Github geladen.

Während cmddocs filebasiert ist, speichere ich bei `rvo` alles in einer
MongoDB Instanz. Verschiedene Aspekte haben mich zu diesem Schritt geführt.
Besonders Metadaten. Eine einfache Variante von Tags und Kategorien waren
mit purer Filebasis nur schwer möglich. Außerdem stellte ich fest dass ich bei
cmddocs viel Aufwand betreiben musste, damit alles mit Files und der `git`
Anbindung glatt geht. Außerdem hebe ich nun viel mehr Sachen in `rvo` auf,
als nur mein Wiki. Ich migrierte nach und nach Notizen, Dokumentation,
Journal und Bookmarks nach `rvo`.

Neben viel besseren Möglichkeiten nach Inhalten zu suchen, sind auch
[Verschlüsselung](https://noqqe.de/blog/2016/02/14/how-do-i-salsa/) und
Text-Stats etwas, dass `rvo` kann.

[rvo](https://github.com/noqqe/rvo)

Zusammengefasst würde ich nicht sagen, dass rvo nun der Nachfolger von
cmddocs ist. Es sind unterschiedliche Arten (teils wegen Features, teils
wegen der Benutzung an sich) mit den eigenen Daten umzugehen.

Ich werde cmddocs weiterhin noch
[bugfixen](https://github.com/noqqe/cmddocs/issues?q=is%3Aissue+is%3Aclosed).
Neue Features sind aber nicht mehr zu erwarten.

Feedback gerne willkommen.
