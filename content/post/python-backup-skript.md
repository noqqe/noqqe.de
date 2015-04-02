---
date: 2008-07-16T14:03:12+02:00
type: post
slug: python-backup-skript
comments: true
title: Python Backup Skript
aliases:
- /archives/111
categories:
- Development
- Linux
tags:
- python
---

Sozusagen das mein erstes Nutzskript in python.

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
