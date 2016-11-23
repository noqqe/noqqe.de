---
title: OpenLDAP Backend Tuning
date: 2013-02-23T11:47:14.000000
tags: 
- Software
- OpenLDAP
---


~~~
## ls -lah /usr/local/var/openldap-data
total 1.1M
drwxr-sr-x 2 root staff 4.0K Feb 20 20:59 .
drwxr-sr-x 4 root staff   36 Feb 15 17:43 ..
-rw-r--r-- 1 root staff 2.0K Feb 20 20:59 alock
-rw------- 1 root staff 8.0K Feb 20 20:59 cn.bdb
-rw------- 1 root staff  24K Feb 20 20:59 __db.001
-rw------- 1 root staff 192K Feb 20 20:59 __db.002
-rw------- 1 root staff 264K Feb 20 20:59 __db.003
-rw------- 1 root staff 160K Feb 20 20:59 __db.004
-rw------- 1 root staff 792K Feb 20 20:59 __db.005
-rw------- 1 root staff  32K Feb 20 20:59 __db.006
-rw------- 1 root staff 8.0K Feb 20 20:58 departmentNumber.bdb
-rw------- 1 root staff 8.0K Feb 20 20:32 dn2id.bdb
-rw------- 1 root staff  32K Feb 20 20:59 id2entry.bdb
-rw------- 1 root staff  10M Feb 20 20:59 log.0000000001
-rw------- 1 root staff 8.0K Feb 20 20:59 objectClass.bdb
-rw------- 1 root staff 8.0K Feb 20 20:59 sn.bdb
-rw------- 1 root staff 8.0K Feb 20 20:59 uid.bdb
~~~

Inital und Übersicht

~~~
aptitude install db-util
db_stat -h /usr/local/var/openldap-data/ -m
~~~

* Spezielle Einstellungen für kontretes Backend

#### BerkleyDB

* Übersicht aller Datenbanken:

~~~
db_stat -h /usr/local/var/openldap-data/ -m
~~~

* Konkrete Datenbanken auslesen

~~~
db_stat -h /usr/local/var/openldap-data/ -d dn2id.bdb
db_stat -h /usr/local/var/openldap-data/ -d id2entry.bdb
~~~

* Korrekte Cachesize errechnen

~~~
## db_stat -h /usr/local/var/openldap-data/ -d dn2id.bdb
Fri Mar  8 09:11:16 2013	Local time
53162	Btree magic number
9	Btree version number
Little-endian	Byte order
duplicates, sorted duplicates	Flags
2	Minimum keys per-page
4096	Underlying database page size
1007	Overflow key/data size
2	Number of levels in the tree
1149	Number of unique keys in the tree
2281	Number of data items in the tree
3	Number of tree internal pages
11444	Number of bytes free in tree internal pages (6% ff)
29	Number of tree leaf pages
48174	Number of bytes free in tree leaf pages (59% ff)
4	Number of tree duplicate pages
642	Number of bytes free in tree duplicate pages (96% ff)
0	Number of tree overflow pages
0	Number of bytes free in tree overflow pages (0% ff)
0	Number of empty pages
0	Number of pages on the free list
~~~

Die wichtigen Angaben nochmal in kurz

~~~
3	Number of tree internal pages
29	Number of tree leaf pages
~~~

Blockgröße des Dateisystems: 4KB

Formel:

    ( 1 root Page + 3 internal Pages + 29 leaf Pages) * 4KB Blocksize = 132 KB Cache Size

~~~
root@vm29-ldap:~## db_stat -h /usr/local/var/openldap-data/ -d id2entry.bdb
Fri Mar  8 09:30:37 2013	Local time
53162	Btree magic number
9	Btree version number
Little-endian	Byte order
	Flags
2	Minimum keys per-page
16384	Underlying database page size
4079	Overflow key/data size
2	Number of levels in the tree
573	Number of unique keys in the tree
573	Number of data items in the tree
1	Number of tree internal pages
15728	Number of bytes free in tree internal pages (4% ff)
29	Number of tree leaf pages
45622	Number of bytes free in tree leaf pages (90% ff)
0	Number of tree duplicate pages
0	Number of bytes free in tree duplicate pages (0% ff)
0	Number of tree overflow pages
0	Number of bytes free in tree overflow pages (0% ff)
0	Number of empty pages
0	Number of pages on the free list
~~~

bei der

    ( 1 root Page + 3 internal Pages + 29 leaf Pages) * 4KB Blocksize = 132 KB Cache Size

Einbau

    DB_CONFIG
    set_cachesize 2 524288000 1
    dbconfig set_cachesize 2 524288000 1

Wichtige Datenbanken der BDB

~~~
db_stat -h /usr/local/var/openldap-data/ -d dn2id.bdb
Wed Feb 20 21:10:00 2013        Local time
53162   Btree magic number
9       Btree version number
Little-endian   Byte order
duplicates, sorted duplicates   Flags
2       Minimum keys per-page
4096    Underlying database page size
1007    Overflow key/data size
1       Number of levels in the tree
17      Number of unique keys in the tree
26      Number of data items in the tree
0       Number of tree internal pages
0       Number of bytes free in tree internal pages (0% ff)
1       Number of tree leaf pages
3106    Number of bytes free in tree leaf pages (24% ff)
0       Number of tree duplicate pages
0       Number of bytes free in tree duplicate pages (0% ff)
0       Number of tree overflow pages
0       Number of bytes free in tree overflow pages (0% ff)
0       Number of empty pages
0       Number of pages on the free list

## db_stat -h /usr/local/var/openldap-data/ -d id2entry.bdb
Wed Feb 20 21:09:37 2013        Local time
53162   Btree magic number
9       Btree version number
Little-endian   Byte order
        Flags
2       Minimum keys per-page
16384   Underlying database page size
4079    Overflow key/data size
1       Number of levels in the tree
8       Number of unique keys in the tree
8       Number of data items in the tree
0       Number of tree internal pages
0       Number of bytes free in tree internal pages (0% ff)
1       Number of tree leaf pages
11574   Number of bytes free in tree leaf pages (29% ff)
0       Number of tree duplicate pages
0       Number of bytes free in tree duplicate pages (0% ff)
0       Number of tree overflow pages
0       Number of bytes free in tree overflow pages (0% ff)
0       Number of empty pages
0       Number of pages on the free list
~~~

### Anpassung der Cache Größe

~~~
## slapindex
51252ecd bdb_db_open: warning - no DB_CONFIG file found in directory /usr/local/var/openldap-data: (2).
Expect poor performance for suffix "dc=example,dc=com".
~~~

Config File editieren:

~~~
vim /usr/local/var/openldap-data/DB_CONFIG
set_cachesize   0       3000000       1
~~~

Danach nochmal

~~~
## slapindex
51252e44 bdb_db_open: DB_CONFIG for suffix "dc=example,dc=com" has changed.
51252e44 Performing database recovery to activate new settings.
51252e44 bdb_monitor_db_open: monitoring disabled; configure monitor database to enable
~~~

Hurra! Schönere Indexe
