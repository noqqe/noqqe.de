---
date: 2008-07-16T14:03:12+02:00
layout: post
slug: python-backup-skript
status: publish
comments: true
title: Python Backup Skript
alias: /archives/111
categories:
- Coding
- Linux
tags:
- python
---

Sozusagen das mein erstes Nutzskript in python :)

(Ich weiss es ist unsinn das /home Verzeichnis nach /home zu backuppen :) aber ist auf dem EeePC ja nur zu Testzwecken.)

``` python
#!/usr/bin/python

import os

import time

quellen = [ '/home/noqqe' ]

ziel_verzeichnis = ('/home/noqqe/backup/' )

ziel = ziel_verzeichnis + time.strftime( '%d%m%Yum%H:%M' ) + '.zip'
zip_befehl = 'zip -qr %s %s' % (ziel, ' '.join(quellen))

if os.system(zip_befehl) == 0:

print 'Erfolgreiche Sicherung nach', ziel

else:

print 'Sicherung fehlgeschlagen!'
```

Hier nochmal mit Syntax Highlightning und mit _wichtigen_ Einrückungen.

[http://paste.pocoo.org/show/79569/](http://paste.pocoo.org/show/79569/)
