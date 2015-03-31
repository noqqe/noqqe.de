---
date: 2009-12-01T11:57:37+02:00
layout: post
slug: wordpress-feed-bestimmter-kategorie-generieren
status: publish
comments: true
title: Wordpress | Feed bestimmter Kategorie generieren
alias: /archives/752
categories:
- Coding
- Linux
- PHP
tags:
- feed
- kategorie
- planet
- RSS
- ubuntu
- ubuntuusers
- ubuntuusersplanet
- wordpress
---

![ubuntuusers-logo](/uploads/2009/12/ubuntuusers-logo.serendipityThumb.png)Überlege zur Zeit, ob ich mich nicht mal anfrage den Blog in den [ubuntuusers.de Planet](http://planet.ubuntuusers.de) aufzunehmen. Ob meine Postings die Qualität des Planeten erreichen können lass ich jetzt mal dahingestellt. Ohnehin kann ich nicht den ganzen Feed in den Planet laufen lassen. Ich bräuchte eine Art extra Output-Lösung.

Aufgrund dessen habe ich überlegt wie ich sowas realisieren könnte. Eine Option die ich anklicke um den Post auch weiter an den Planet zu geben oder ähnliches. Eine Kategorie vergeben die sich weiterschickt. Nach kurzem herumprobieren mit dem RSS-Feed von Wordpress hab ich (während einer sehr langweiligen C++ Programmierstunde in der Schule) herausgefunden das sich der RSS-Feed eine zusätzliche Kategorie-Variable mitübergeben lässt um nur Artikel der definierten Kategorie auszuspucken.

**/?feed=rss&cat=ID_364**

Das ist recht nice, weil ich Postings für den Planeten wahlweise nur in den Kategorien ankreuze und automatisch beim Planeten nur eben diese Artikel ankommen.

Wordpress bewundernd,
Flo
