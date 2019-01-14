---
aliases:
- /blog/2011/02/24/charset-utf8-fur-apache-mysql-debian-und-wordpress
- /archives/1487
comments:
- author: savier
  content: "<p>Charsets sind wirklich so eine Sache f\xFCr sich :) Man kann mit den
    kleinsten Problemen den gr\xF6\xDFten Spa\xDF haben.</p>"
  date: '2011-02-24T19:43:39'
- author: Christian
  content: "<p>Bin grade auf deinen Blog gesto\xDFen, nachdem ich es selbst verzweifelt
    probiert habe. Ich habe auch noch keine funktionierende Sache, glaube aber, dass
    du trotzdem einen Fehler hier drin hast. Kannst es ja vlt mal \xFCberpr\xFCfen:
    Meine Datei war die \"/etc/mysql/conf.d/mysqld_anything_utf8.cnf\" ich habe die
    Datei-Endung .cnf gew\xE4hlt, weil in der letzten Zeilen von /etc/mysql/my.cnf
    steht, dass alle Dateien in /etc/mysql/conf.d/ die Dateiendung cnf haben m\xFCssen,
    sonst werden sie ignoriert. Ich vermute, daher wird deine Datei irgnoriert und
    du bekommst nicht den Fehler, den ich bekomme. Ich habe n\xE4mlich in meiner Datei
    /etc/mysql/conf.d/mysqld_anything_utf8.cnf stehen:<br>[mysqld]<br>character_set_database=utf8<br>character_set_server=utf8<br>Und
    der mysql-Server verweigert dann mit \"/etc/init.d/mysql restart\" das starten,
    bis ich die Datei /etc/mysql/conf.d/mysqld_anything_utf8.cnf\_ wieder gel\xF6scht
    habe.<br>Deine Blog mag also funktionieren, aber das liegt vermutlich an was anderem.<br>LG
    Chris</p><p>PS: Warum zum Teufel k\xF6nnen die nicht standardm\xE4\xDFig \xFCberall
    UTF8 nehmen....!?!?!<br></p>"
  date: '2012-06-12T15:39:46'
- author: Christian
  content: "<p>\_Ok, hier meine L\xF6sung, die bei mir mit MySQL 5.5 unter Ubuntu12.04
    nun endlich funktioniert. Meine Datei ist /etc/mysql/conf.d/mysqld_anything_utf8.cnf<br>Die
    Datei hat den Inhalt<br>_________<br>[mysqld]<br>character_set_server=utf8<br>_________<br>Die
    Zeile character_set_database=utf8 durfte schlichtweg nicht enthalten sein, da
    nach dem Neustart des MySQL-Service mittels \"sudo service mysql restart\" character_set_database
    automatisch mit auf utf8 ge\xE4ndert wurde. Wichtig ist aber, dass die Datei die
    Endung .cnf aufwei\xDFt, sonst wird sie ignoriert.<br>Danach kontrolliert man
    noch sch\xF6n alles mit \"mysql -u root -p\" Mysql fragt einen dann nat\xFCrlich
    nach dem Passwort. Eingeben und dann ist man in der MySQL-Console. Dort gibt man
    ein \"SHOW VARIABLES LIKE 'char%';\" (nat\xFCrlich ohne die Anf\xFChrungszeichen,
    die ich gesetzt habe). Nun sollte alles (au\xDFer nat\xFCrlich character_set_filesystem)
    nicht mehr auf latin1, sondern auf utf8 stehen. Vorher war das ja nicht der Fall,
    da stand dort latin1. Die Collations stimmen nun auch. (show variables like 'collation%';)<br>Ich
    werde noch ein Programm schreiben, welches \xFCberpr\xFCft ob auch die Connection
    auf utf8 gesetzt ist, wenn ich per Java und MysqlConnectorJ darauf zugreife.</p>"
  date: '2012-06-12T16:02:38'
- author: noqqe
  content: "<p>Oh, ja da hast du nat\xFCrlich vollkommen recht! Die Dateiendung ist
    falsch! </p><p>Ich \xE4nder das :)</p><p>Vielen Dank!</p>"
  date: '2012-06-12T17:48:53'
date: '2011-02-24T09:22:13'
tags:
- shell
- mysql
- utf8
- charset
- administration
- apache2
- wordpress
- sql
- linux
- php
- debian
title: "Charset | UTF8 f\xFCr Apache, PHP, MySQL, Debian und Wordpress "
---

Nachdem ich die Migration meines Blogs auf meinen neues Stück Blech
größtenteils abgeschlossen hatte, wurde ich wieder an den Charset Wirr-Warr
von IT-Systemen erinnert. Um meinem Blog seine Umlaute wieder zu beschaffen
habe ich folgende Änderungen an verschiedenen Stellen eingespielt.
Vorzugsweise immer in den entsprechenden conf.d/ Verzeichnissen, da die
Änderungen evtl. beim nächsten Upgrade überschrieben werden könnten.

# �

### Apache2 Charset

```
vim /etc/apache2/conf.d/charset
AddDefaultCharset UTF-8
```

### PHP5 Charset

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

### MySQL Charset

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

### Debian Wordpress Config

```
$ vim /etc/wordpress/config-blog.url.php
define('DB_CHARSET', 'utf8');
define('DB_COLLATE', '');
define('WPLANG', 'de_DE.UTF-8');
```

### Debian Locales

```
dpkg-reconfigure locales
```

Sollte ich es mal wieder brauchen, lese ich hier nach.
