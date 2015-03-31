---
date: 2008-09-10T19:49:18+02:00
type: post
slug: mein-erster-apache2
status: publish
comments: true
title: Mein erster Apache2
alias: /archives/224
categories:
- Linux
tags:
- apache2
- Linux
- ubuntu
- web
---

**Hier in aller Kürze und ohne großartige Comments: die apache2.conf**

```
# Based upon the NCSA server configuration files originally by Rob McCool.
# This is the main Apache server configuration file.  It contains the
# configuration directives that give the server its instructions.
# See http://httpd.apache.org/docs/2.2/ for detailed information about
# the directives.
# Do NOT simply read the instructions in here without understanding
# what they do.  They're here only as hints or reminders.  If you are unsure
# consult the online docs. You have been warned.
```


```
 ServerRoot "/etc/apache2"
LockFile /var/lock/apache2/accept.lock
PidFile $APACHE_PID_FILE
Timeout 300
```


```
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
```


```
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
```


```
 LogLevel warn
Include /etc/apache2/mods-enabled/*.load
Include /etc/apache2/mods-enabled/*.conf
Include /etc/apache2/httpd.conf
Include /etc/apache2/ports.conf
LogFormat "%h %l %u %t "%r" %>s %b "%{Referer}i" "%{User-Agent}i"" combined
LogFormat "%h %l %u %t "%r" %>s %b" common
```


```
 LogFormat "%{Referer}i -> %U" referer
LogFormat "%{User-agent}i" agent
ServerTokens Full
ServerSignature On
Include /etc/apache2/conf.d/
Include /etc/apache2/sites-enabled/

```

**Hier mal eben die envvars datei**
```

# envvars - default environment variables for apache2ctl
# Since there is no sane way to get the parsed apache2 config in scripts, some
# settings are defined via environment variables and the
```


```
n used in apache2ctl,
# /etc/init.d/apache2, /etc/logrotate.d/apache2, etc.
export fbaumann=APACHE_RUN_USER
export fbaumann=APACHE_RUN_GROUP
export APACHE_PID_FILE=/var/run/apache2.pid
```


Im großen und ganzen installiert man nur kurz das deb Paket apache2, nimmt kurze Anpassungen an den Configs vor und schon ist unter 127.0.0.1 der Webserver abzurufen. Garkein so großes Kunststück wie ich dachte. Hier zuerst die index.html (http://127.0.0.1)

![](http://farm4.static.flickr.com/3043/2844596855_ec40db47d7.jpg?v=0)

und dann ein Unterordner (http://127.0.0.1/test)

![](http://farm4.static.flickr.com/3022/2844596767_440587c2a3.jpg?v=0)

Linux is like a Wigwam. Apache inside and no windows.

In diese Sinne,

Habe die Ehre.
