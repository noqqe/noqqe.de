---
aliases:
- /archives/752
- /blog/2009/12/01/wordpress-feed-bestimmter-kategorie-generieren
date: '2009-12-01T09:57:37'
tags:
- development
- feed
- linux
- planet
- wordpress
- ubuntu
- php
- kategorie
- rss
title: Wordpress | Feed bestimmter Kategorie generieren
---

![ubuntuusers-logo](/uploads/2009/12/ubuntuusers-logo.serendipityThumb.png)

Überlege zur Zeit, ob ich mich nicht mal anfrage den Blog in den
[ubuntuusers.de Planet](http://planet.ubuntuusers.de) aufzunehmen. Ob meine
Postings die Qualität des Planeten erreichen können lass ich jetzt mal
dahingestellt.  Ohnehin kann ich nicht den ganzen Feed in den Planet laufen
lassen. Ich bräuchte eine Art extra Output-Lösung.

Aufgrund dessen habe ich überlegt wie ich sowas realisieren könnte. Eine
Option die ich anklicke um den Post auch weiter an den Planet zu geben oder
ähnliches. Eine Kategorie vergeben die sich weiterschickt. Nach kurzem
herumprobieren mit dem RSS-Feed von Wordpress hab ich (während einer sehr
langweiligen C++ Programmierstunde in der Schule) herausgefunden das sich
der RSS-Feed eine zusätzliche Kategorie-Variable mitübergeben lässt um nur
Artikel der definierten Kategorie auszuspucken.

**/?feed=rss&cat=ID_364**

Das ist recht nice, weil ich Postings für den Planeten wahlweise nur in den
Kategorien ankreuze und automatisch beim Planeten nur eben diese Artikel
ankommen.
