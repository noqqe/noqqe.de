---
title: MySQL Replication
date: 2012-03-28T15:10:59
tags:
- Databases
- MySQL
---

## Einzelne Schritte in Quick and Dirty

#### Master

* Server ID prüfen

```
server-id = 1
```

* Master Status prüfen

```
show master status;
```

* Mysql Sklave nutzer überprüfen und anlegen

``` sql
mysql> CREATE USER 'repl'@'%.mydomain.com' IDENTIFIED BY 'slavepass';
mysql> GRANT REPLICATION SLAVE ON *.* TO 'repl'@'%.mydomain.com';
```

* Table gesperrt

``` sql
FLUSH TABLES WITH READ LOCK;
mysqldump --all-databases --master-data >dbdump.db
UNLOCK TABLES;
```

#### Slave

* Dump wie gewohnt einspielen

* Master konfigurieren

``` sql
mysql> CHANGE MASTER TO MASTER_HOST='db00.example.com',
MASTER_USER='slave', MASTER_PASSWORD='xxx',
MASTER_LOG_FILE='mysql-bin.000254', MASTER_LOG_POS=106;
```

* Slave starten

```
start slave;
show slave status;
```

* table created und wieder gelöscht.

## Slave und Master Konfiguration komplett löschen

```
STOP SLAVE;
RESET SLAVE;
## Use RESET SLAVE ALL; for MySQL 5.5.16 and later
```

```
STOP MASTER;
RESET MASTER;
## Use RESET MASTER ALL; for MySQL 5.5.16 and later
```

## Was tun wenn

### Loggin enabled

```
Error 'You cannot 'ALTER' a log table if logging is enabled' on query. Default database: 'mysql'. Query: 'ALTER TABLE slow_log
```

Das Problem liegt wohl an den Slow query Log

``` sql
STOP SLAVE;
SET GLOBAL slow_query_log = "OFF";
START SLAVE;
SET GLOBAL slow_query_log = "ON";
```

### alte Binlogs loeschen

Auf master schauen welches logfile aktuell ist:

```
show master status\G
     Master_Log_File: mysql-bin.000687

```

Auf Slave schauen ob er auch auf dem aktuellen ist.

```
show slave status\G
     Master_Log_File: mysql-bin.000687
```

Auf Master die Binlogs mit

``` sql
PURGE BINARY LOGS TO 'mysql-bin.000685';
PURGE BINARY LOGS BEFORE '2008-04-02 22:46:26';
```

löschen
