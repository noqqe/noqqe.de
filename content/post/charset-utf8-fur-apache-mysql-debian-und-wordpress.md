---
date: 2011-02-24T11:22:13+02:00
type: post
slug: charset-utf8-fur-apache-mysql-debian-und-wordpress
status: publish
comments: true
title: 'Charset | UTF8 für Apache, PHP, MySQL, Debian und Wordpress '
aliases:
- /archives/1487
categories:
- Bash
- Coding
- Debian
- General
- Linux
- PlanetenBlogger
- Ubuntu
- Web
tags:
- apache2
- charset
- debian
- mysql
- PHP
- sql
- utf8
- wordpress
---

Nachdem ich die Migration meines Blogs auf meinen neues Stück Blech größtenteils abgeschlossen hatte, wurde ich wieder an den Charset Wirr-Warr von IT-Systemen erinnert. Um meinem Blog seine Umlaute wieder zu beschaffen habe ich folgende Änderungen an verschiedenen Stellen eingespielt. Vorzugsweise immer in den entsprechenden conf.d/ Verzeichnissen, da die Änderungen evtl. beim nächsten Upgrade überschrieben werden könnten.



# �


**Apache2 Charset**
```
vim /etc/apache2/conf.d/charset
AddDefaultCharset UTF-8
```


**PHP5 Charset**
```
$ vim /etc/php5/apache2/conf.d/charset.ini
[PHP]
default_charset = "utf-8"
[mbstring]
mbstring.language = utf-8
mbstring.internal_encoding = utf-8
mbstring.http_input = utf-8
mbstring.http_output = utf-8
```


**MySQL Charset**
```
$ vim /etc/mysql/conf.d/character.cnf
[client]
default-character-set = utf8
[mysqld]
default-character-set = utf8
character-set-server = utf8
collation-server= utf8_general_ci
init_connect = ‘SET collation_connection = utf8_general_ci’
init_connect = ‘SET NAMES utf8′
[mysqldump]
default-character-set = utf8
[mysqlimport]
default-character-set = utf8
[mysql]
default-character-set = utf8
```


**Debian Wordpress Config**
```
$ vim /etc/wordpress/config-blog.url.php
define('DB_CHARSET', 'utf8');
define('DB_COLLATE', '');
define('WPLANG', 'de_DE.UTF-8');
```


**Debian Locales**
```
$ dpkg-reconfigure locales
```


Sollte ich es mal wieder brauchen, les ich hier nach.
