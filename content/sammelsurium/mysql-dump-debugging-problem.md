---
title: MySQL Dump Debugging Problem
date: 2013-07-08T09:57:51.000000
tags: 
- Databases
- MySQL
---


Beim Dump tritt Fehler auf

~~~
mysqldump: Couldn't execute 'show events': Cannot proceed because system tables used by
Event Scheduler were found damaged at server start (1577)
~~~

Genauere Inspektion:

~~~
mysql> use information_schema;
Reading table information for completion of table and column names
You can turn off this feature to get a quicker startup with -A

Database changed
mysql> SELECT /*!40001 SQL_NO_CACHE */ * FROM EVENTS;
ERROR 1577 (HY000): Cannot proceed because system tables used by Event Scheduler were found damaged at server start
~~~

## Lösung

    mysql_upgrade -u root -h localhost -p --verbose

und die Datenbank restarten. Danach ist wieder alles fresh. Auch wenn das eigentliche Upgraden dann so aussieht als wäre nichts passiert:

~~~
$ mysql_upgrade --verbose
Looking for 'mysql' as: mysql
Looking for 'mysqlcheck' as: mysqlcheck
This installation of MySQL is already upgraded to 5.5.31, use --force if you still need to run mysql_upgrade
~~~

Ist aber trotzdem gut danach

~~~
Database changed
mysql> SELECT /*!40001 SQL_NO_CACHE */ * FROM EVENTS;
Empty set (0.00 sec)
~~~


