---
date: 2010-03-31T20:12:03+02:00
comments: true
title: Apache2 | vhost-Logs und logrotate
aliases:
- /archives/946
- /blog/2010/03/31/apache2-vhost-logs-und-logrotate
categories:
- Development
- Linux
- ubuntuusers
tags:
- apache
- apache2
- logging
- logrotate
- logs
- vhost
- vhosts
- virtual host
---

![169px-ASF-logo.svg](/uploads/2010/03/169px-ASF-logo.svg.png)

Nachdem mein Apache immer mehr und mehr virtuelle Hosts / eingerichtete
Subdomains bekommt, welche alle in das /var/log/apache2/access.log
loggen, fand ich es an der Zeit etwas mehr Struktur rein zu bringen.
Generell werden alle bestehenden vhosts in /etc/apache2/sites-available/*
deklariert.  Darin befinden sich allerlei Deklarationen wie der Ort des
Directory und ähnliches. Kurzer Head-Auszug:

```
ServerAdmin webmaster@zwetschge.org
ServerName noqqe.de
ServerAlias www.noqqe.de
DocumentRoot /var/www/blog/
```

Unter anderem auch die Anweisung für logging:

```
LogLevel warn
CustomLog /var/log/apache2/access.log combined
```

Alle meine vhosts haben den selben Eintrag und loggen demnach auch alle in
das selbe File. Dieser Umstand wird relativ schnell zum Problem. Ständiges
greppen und ähnliches an der Tagesordnung. Apache ist aber eben auch in der
Lage für jeden vhost eine extra Logfile zu generieren.

```
LogLevel warn
ErrorLog /var/log/apache2/noqqe.de-error.log
CustomLog /var/log/apache2/noqqe.de-access.log common
```


Mit der Syntax `[$domain]-[access|error].log` ist es im
Apache Logdir schön auseinander zu halten auf welchen vhost welche
Anfragen gingen. Erleichtert die Suche ungemein. Auch Aufgliederung in
apache2/error/* und apache2/access/* wäre denkbar. Aber für mich gerade
oversized.

Nächster Punkt auf der Liste: logrotate

Ein Logfile wird _rotiert_. Diese Rotation übernimmt das Programm
Logrotate. Jeder kennt es:

```
-rw-r----- 1 root adm 360K 29. Mär 00:25 access.log.1
-rw-r----- 1 root adm 193K 31. Jan 00:25 access.log.10.gz
-rw-r----- 1 root adm 192K 24. Jan 00:25 access.log.11.gz
```

Ich möchte auch gerne meine vhost-Logfiles rotieren lassen. Wie in jedem
(guten) Daemon gibt es auch hier ein Verzeichnis logrotate.d/*. Darin
befinden sich manuell angepasste config-Files die z.B. bei Upgrades nicht
überschrieben werden. Der folgende Eintrag lässt alle Dateien die mit .log
enden und sich im Verzeichnis /var/log/apache2/ befinden wöchentlich bis zu
10 Wochen rotieren und komprimieren:

    /var/log/apache2/*.log {
            weekly
            missingok
            rotate 10
            compress
            delaycompress
            notifempty
            create 640 root adm
            sharedscripts
            postrotate
                    if [ -f "`. /etc/apache2/envvars ; echo ${APACHE_PID_FILE:-/var/run/apache2.pid}`" ]; then
                            /etc/init.d/apache2 reload > /dev/null
                    fi
            endscript
    }
