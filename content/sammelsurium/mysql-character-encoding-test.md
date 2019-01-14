---
title: MySQL Character Encoding Test
date: 2012-02-22T13:41:55
tags: 
- Databases
- MySQL
---

Wichtig ist: Wie kamen die Daten in die DB. Wenn mit Latin, kamen die ANSI
Chars rein und die UTF8 Kommands werden doppelt Encoded. Dadurch entstehen
die A1/4 Zeichen. Wenn die Daten aber mit LATIN1 aus dem Table geholt
werden encoded die LOCALES des Systems die Chars:

Mit UTF8 (wie man glauben möchte der richtige)

~~~
mysql --default-character-set=UTF8 -e "select id, query from data where query like '%nchen%' limit 1; "
+----+----------------------------------------------------------+
| id | query                                                    |
+----+----------------------------------------------------------+
|  7 | 81539 MÃ¼nchen, Deutschland                              |
+----+----------------------------------------------------------+
~~~

Mit Latin1

~~~
mysql --default-character-set=LATIN1 -e "select id, query from data where query like '%nchen%' limit 1; "
+----+-------------------------------------------------------+
| id | query                                                 |
+----+-------------------------------------------------------+
|  7 | 81539 München, Deutschland                            |
+----+-------------------------------------------------------+
~~~

## Charset Configs #

### Latin1 ##

~~~
[client]
default-character-set = latin1
[mysqld]
default-character-set = latin1
character-set-server = latin1
collation-server = latin1_general_ci
init_connect = 'SET collation_connection = latin1_general_ci'
init_connect = 'SET NAMES latin1'
[mysqldump]
default-character-set = latin1
[mysqlimport]
default-character-set = latin1
[mysql]
default-character-set = latin1
~~~

### UTF8 ##

~~~
[client]
default-character-set = utf8
[mysqld]
default-character-set = utf8
character-set-server = utf8
collation-server = utf8_general_ci
init_connect = 'SET collation_connection = utf8_general_ci'
init_connect = 'SET NAMES utf8'
[mysqldump]
default-character-set = utf8
[mysqlimport]
default-character-set = utf8
[mysql]
default-character-set = utf8
~~~

## Richtig Dumpen #

~~~
mysqldump --default-character-set="UTF8" db > db_utf8.sql
mysql --default-character-set="UTF8" db < db_utf8.sql
~~~

## Links #

http://www.gerd-riesselmann.de/softwareentwicklung/php-und-utf-8-eine-anleitung-teil-1-mysql

http://dev.mysql.com/doc/refman/5.1/de/charset-connection.html