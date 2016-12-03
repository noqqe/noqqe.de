---
date: 2012-01-22T15:09:07+02:00
comments: true
title: MySQL | Table Migration zu InnoDB
aliases:
- /blog/2012/01/22/mysql-table-migration-zu-innodb
- /archives/1850
tags:
- Shell
- Debian
- Linux
- ubuntuusers
- Databases
- acid
- database
- datenbank
- fulltext
- innodb
- Migration
- myisam
- mysql
- mysqldump
- sql
- tables
---

Ich durfte letztens einer schönen Schulung zum Thema MySQL lauschen. Dabei
kam viel herrum. Ein Teil davon hat sich mit verschiedenen Storage Engines
beschäftigt.

Ich habe mich entschieden den Großteil meiner Datenbanken zu InnoDB
umzuwandeln. Welche Vor- und Nachteile das hatt sollte sich jeder vorher
klarmachen. Stichwort: ACID und Fulltext-Search. Natürlich könnte man
einfach mit ALTER arbeiten. Aber ich wollte nicht umbedingt die alten
MyISAM Tables noch im FS liegen haben. Also alle Datenbanken droppen und
den geänderten Dump wieder einspielen, sodass die Tables neu (mit InnoDB)
angelegt werden.

Zuersteinmal eine Liste mit den generieren mit den Datenbanken die man
Bearbeiten möchte. Es ist zu beachten, dass man den mysql Table selbst
nicht auf InnoDB umstellen möchte.

``` bash
mysql -u root -p -e "show databases;" -N --batch | grep -v ^information_schema$ | grep -v ^mysql$
```

Die Liste der Datenbanken wird nachher noch hilfreich sein. Danach will man
wahrscheinlich erstmal alle Dienste beenden, die auf dem MySQL zugreifen
(Apache, Tomcat, whatever). In der my.cnf habe ich dann folgende Optionen
für die InnoDB spezifiziert.

```
defaulot-storage-engine = InnoDB
innodb_buffer_pool_size = 16M
innodb_additional_mem_pool_size = 2M
innodb_log_file_size = 5M
innodb_log_buffer_size = 8M
innodb_flush_log_at_trx_commit = 1
innodb_lock_wait_timeout = 50
```


Anschliessend den Dump erstellen und alle ENGINE=MyISAM durch InnoDB ersetzen:

``` bash
mysqldump -u root -p > all-databases.sql
sed -i -e 's#ENGINE=MyISAM#ENGINE=InnoDB#g' all-databases.sql
```

**Vorsicht. Hier ist mysql als Datenbank mit gedumped!** Mir ist dabei
keine wirklich einfache Zeile eingefallen die mit Suche/Ersetze Spielchen
mysql ausschliesst. Es gibt bei mysqldump die Option "--ignore-table=" aber
auch hier hätte ich jeden mysql Table einzeln nennen müssen. Ich hab die
Datenbank dann einfach per hand aus dem Dump herausgelöscht.

Außerdem sollte man seinen Datenbank Dump nach FULLTEXT durchsuchen, da
dieser von InnoDB nicht unterstützt wird. In meinem Fall hat es nur ein
altes Forum das niemand mehr benutzt betroffen, weshalb ich die Zeile
einfach löschen konnte.

``` bash
grep "FULLTEXT" all-databases.sql
```

Um jetzt alle Datenbanken zu droppen hab ich mir folgende Line gebastelt:

``` bash
for x in $(mysql -u root -phierstehteinpasswort -e "show databases;" -N --batch | grep -v ^information_schema | grep -v ^mysql$) ; do
  mysql -u root -phierstehteinpasswort -e "drop database $x ; " --batch
done
```

Nach der Bearbeitung kann man den Dump mit der neuen Engine für die Tables
wieder einspielen:

``` bash
mysql -u root -p < all-databases.sql
```
