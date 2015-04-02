---
date: 2009-12-17T11:43:13+02:00
type: post
slug: wordpress-rss-und-planet
comments: true
title: Wordpress | RSS und Planet
aliases:
- /archives/805
categories:
- Linux
- PHP
tags:
- blog
- feed
- PHP
- planet
- RSS
- rss2
- wordpress
---

![rss](/uploads/2009/12/rss.gif)Um bestimmte Beiträge des Blogs im Planeten erscheinen zulassen, hab ich wie [bereits beschrieben](/?p=752) einen Feed einer bestimmten Kategorie erstellt.
**/?feed=rss&cat=ID_364**

Nach kurzer Rücksprache mit [Ritze](http://ubuntuusers.de/user/Ritze/) vom Ubuntuusers-Team ergab sich aber das mit dem generierten Feed was nicht stimmt. **<published>** und **<updated>** Tags würden fehlen und somit liess sich der Feed nicht in den Planeten einbinden. Ich begann meinen Feed mit anderen Blogfeeds des Planeten zu vergleichen. Diese rief ich immer mit **http://blog.de/?feed=rss **auf. Mysteriöserweise erkannte ich nie Unterschiede.

Am End war trotzdem nur eine einzige Ziffer ausschlaggebend.
**/?feed=rss&cat=ID_364
/?feed=rss2&cat=ID_364**
Wer erkennt den Unterschied?
Kopf -> Tisch.
