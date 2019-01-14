---
title: MySQL Benutzerverwaltung
date: 2012-03-06T09:07:40
tags:
- Databases
- MySQL
---

## Neuen Benutzer anlegen

    GRANT ALL ON testdb.* TO 'testuser'@'%' IDENTIFIED BY 'supergeheim23;

## Benutzer löschen

    drop user username;

## Passwort neu setzen

    SET PASSWORD FOR 'bob'@'%.loc.gov' = PASSWORD('newpass');

## Root ähnlichen User erzeugen

    GRANT ALL PRIVILEGES ON *.* TO 'root'@'localhost' IDENTIFIED BY PASSWORD 'mypassword' WITH GRANT OPTION

### Passwort fuer alle root user neu setzen

    use mysql;
    update user set password=PASSWORD("XXX") where user='root';

## Aktuelle Berechtigungen anschauen

    SHOW GRANTS for username;

    SHOW GRANTS;

## Update Old MySQL Password Hashes to new ones

~~~
mysql> SET old_passwords = 0;
Query OK, 0 rows affected (0.00 sec)

mysql> SET PASSWORD FOR 'username'@'%' = PASSWORD('XXX');
Query OK, 0 rows affected (0.03 sec)
~~~

## Leere Benutzerfelder und GRANTS for any

~~~
select * from mysql.db;
*************************** 37. row ***************************
                 Host: %
                   Db: test
                 User:
          Select_priv: Y
~~~

Das geht so:

    grant ALL on test.* to ''@'%';

Dann kann man auch die Berechtigungen dafür so anschauen:

~~~
mysql> show grants for '';
+--------------------------------------------+
| Grants for @%                              |
+--------------------------------------------+
| GRANT USAGE ON *.* TO ''@'%'               |
| GRANT ALL PRIVILEGES ON `test`.* TO ''@'%' |
+--------------------------------------------+
2 rows in set (0.00 sec)
~~~

## Wildcard User != Asterisk User

It actually is a user.

~~~
mysql> show grants for '*';
+---------------------------------------------+
| Grants for *@%                              |
+---------------------------------------------+
| GRANT USAGE ON *.* TO '*'@'%'               |
| GRANT ALL PRIVILEGES ON `test`.* TO '*'@'%' |
+---------------------------------------------+
2 rows in set (0.00 sec)

mysql> drop user '*'@'%';
Query OK, 0 rows affected (0.00 sec)

mysql> show grants for '*';
ERROR 1141 (42000): There is no such grant defined for user '*' on host '%'
~~~