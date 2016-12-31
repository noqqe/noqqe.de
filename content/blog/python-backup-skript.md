---
aliases:
- /archives/111
comments:
- author: basti
  content: "<p>Der Post hier ist ja schon etwas \xE4lter, aber meinen Senf gebe ich
    trotzdem gern dazu ab:</p><p>Mit einem Backup-Script habe ich selbst angefangen,
    Python zu lernen, bzw. bin eh gerade dabei. Genauso wie du hab ich das hier[0]
    festgehalten. </p><p>Dein Paste-Link ist \xFCbrigens tot. Warum hast du den auch
    ausgelagert...?!<br>[0]<a href=\"http://zufallsheld.de/2013/09/29/python-backup-script-with-rsync/\"
    rel=\"nofollow\">http://zufallsheld.de/2013/09/...</a></p>"
  date: '2013-10-10T19:20:29'
- author: noqqe
  content: "<p>Hi, mittlerweile w\xFCrde ich zu python Backup Skript nicht mehr raten.
    Siehe auch <a href=\"http://noqqe.de/blog/2009/03/18/backup-bash-vs-python/\"
    rel=\"nofollow\">http://noqqe.de/blog/2009/03/1...</a></p>"
  date: '2013-10-11T06:51:33'
date: '2008-07-16T12:03:12'
slug: python-backup-skript
tags:
- development
- python
- linux
title: Python Backup Skript
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