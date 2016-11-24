---
title: MySQL Master Switch
date: 2012-08-01T10:04:25.000000
tags: 
- Databases
- MySQL
---


## 1. Methode
### Readonly setzen

~~~
mysql> set GLOBAL read_only = true;
mysql> INSERT INTO foo VALUES ( 'tata2');
ERROR 1290 (HY000): The MySQL server is running with the-read-only
option so it can not execute this statement
~~~

### Readonly entfernen

    mysql> set GLOBAL read_only = false;

## 2. Methode

### Setzen

    FLUSH TABLES WITH READ LOCK;

### Entfernen

    UNLOCK TABLES;

## 3. Methode

### my.cnf

~~~
read-only
read-only=1
read-only=0
~~~
