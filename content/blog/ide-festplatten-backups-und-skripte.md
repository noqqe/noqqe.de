---
aliases:
- /archives/315
date: '2008-11-05T08:38:29'
slug: ide-festplatten-backups-und-skripte
tags:
- development
- shell
- python
- administration
- zwetschge.org
- cron
- linux
- zwetschge
- backup
title: IDE Festplatten, Backups und Skripte
---

Zwetschge hat gestern eine zusätzliche Festplatte für Backups bekommen die
ich per Python und Cron automatisch jeden Mittwoch um 4:01 Uhr machen
lassen möchte. Nach ein paar Master/Slave Spielereien lief die Platte dann
auch. Wichtige Erkenntnis dabei:

```
hda -> master
hdb -> slave
hdc -> master
hdd -> slave
```

mount /dev/hdd /mount/backup

Einhängen wäre damit schonmal geschafft.

Mein nettes Pythonskript:

``` python
#!/usr/bin/python
import os
import time
quellen = ['/home /etc /var/www ']
ziel_verzeichnis = ('/media/backup/backup/')
ziel = ziel_verzeichnis + time.strftime('%Y%m%d') + '.tar.gz'
zip_befehl = 'tar -czvf %s %s' % (ziel, ' '.join(quellen))
if os.system(zip_befehl) == 0:
  print 'Erfolgreiche Sicherung nach', ziel
else:
  print 'Sicherung fehlgeschlagen!'
```

Wird nun via cronjob (python /etc/backup) ausgeführt, und das jeden Mittwoch um 4:01 Uhr:

```
01   04   * * 3 root python /etc/backup
```
