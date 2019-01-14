---
aliases:
- /archives/224
date: '2008-09-10T17:49:18'
slug: mein-erster-apache2
tags:
- web
- linux
- ubuntu
- apache2
title: Mein erster Apache2
---

Hier in aller Kürze und ohne großartige Comments: die apache2.conf

```
ServerRoot "/etc/apache2"
LockFile /var/lock/apache2/accept.lock
PidFile $APACHE_PID_FILE
Timeout 300

KeepAlive On
MaxKeepAliveRequests 100
KeepAliveTimeout 15

StartServers          5
MinSpareServers       5
MaxSpareServers      10
MaxClients          150
MaxRequestsPerChild   0

StartServers          2
MaxClients          150
MinSpareThreads      25
MaxSpareThreads      75
ThreadsPerChild      25
MaxRequestsPerChild   0

# These need to be set in /etc/apache2/envvars
User fbaumann
Group fbaumann
AccessFileName .htaccess

Order allow,deny
Deny from all

DefaultType text/plain
HostnameLookups Off
ErrorLog /var/log/apache2/error.log

LogLevel warn
Include /etc/apache2/mods-enabled/*.load
Include /etc/apache2/mods-enabled/*.conf
Include /etc/apache2/httpd.conf
Include /etc/apache2/ports.conf
LogFormat "%h %l %u %t "%r" %>s %b "%{Referer}i" "%{User-Agent}i"" combined
LogFormat "%h %l %u %t "%r" %>s %b" common
LogFormat "%{Referer}i -> %U" referer
LogFormat "%{User-agent}i" agent
ServerTokens Full
ServerSignature On
Include /etc/apache2/conf.d/
Include /etc/apache2/sites-enabled/
```

Hier mal eben die envvars datei

```
export fbaumann=APACHE_RUN_USER
export fbaumann=APACHE_RUN_GROUP
export APACHE_PID_FILE=/var/run/apache2.pid
```

Im großen und ganzen installiert man nur kurz das deb Paket apache2, nimmt
kurze Anpassungen an den Configs vor und schon ist unter 127.0.0.1 der
Webserver abzurufen. Garkein so großes Kunststück wie ich dachte. Hier
zuerst die index.html (http://127.0.0.1)

Linux is like a Wigwam. Apache inside and no windows.
