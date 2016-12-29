---
comments:
- author: Franck
  content: "<p>Hi Florian,</p><p>danke f\xFCr die Anleitung. Ich nehme an das die
    Berechnung f\xFCr id2entry: <br>( 1 root Page + 1 internal + 29 leaf ) * 16KB
    = 496 KB heissen (16 statt 4KB)</p><p>Viele Gr\xFCsse</p><p>Franck</p>"
  date: '2013-10-22T12:46:44'
- author: noqqe
  content: "<p>Urghs, ja! </p><p>Danke f\xFCr den Hint! Ich update das :)</p>"
  date: '2013-10-23T06:50:33'
date: '2013-03-08T16:40:00'
tags:
- berkeleydb
- bdb
- cache
- administration
- openldap
- ldap
- performance
title: BerkeleyDB Cache Size in OpenLDAP
---

Ehrlich gesagt fühlt man sich als wäre es 1990, wenn man anfängt sich zu
Performance Tuning für OpenLDAP schlau zu lesen.

Das Default Backend `BDB` kann eigentlich Out-of-the-Box benutzt werden. Sobald
ein paar Objects im DIT sind begegnet einem aber schnell die Cache Size.

```
$ slapindex
51252ecd bdb_db_open: warning - no DB_CONFIG file found in directory /usr/local/var/openldap-data: (2).
Expect poor performance for suffix "dc=noqqe,dc=de".
```

### Welche Files sind interessant?

Um die Files der Berkely DB zu finden und auswerten zu können werden extra Tools
benötigt.

``` bash
$ aptitude install db-util
$ db_stat -h /usr/local/var/openldap-data/ -m
```

So werden erstmal alle statistischen Daten zu den einzelnen Datenbankfiles
ausgegeben. Zwei dieser Files kommen bei OpenLDAP eine besondere Aufgabe zu.
`id2entry` und `dn2id`

### dn2id.bdb

Zuerst wird man etwas erschlagen von Werten. Wenn man weiss nach was
man suchen muss, ist aber nur noch wenig Aufwand nötig um die richtigen
Werte rauzusuchen.

``` bash
$ db_stat -h /usr/local/var/openldap-data/ -d dn2id.bdb

Fri Mar  8 09:11:16 2013  Local time
53162 Btree magic number
9 Btree version number
Little-endian Byte order
duplicates, sorted duplicates Flags
2 Minimum keys per-page
4096  Underlying database page size
1007  Overflow key/data size
2 Number of levels in the tree
1149  Number of unique keys in the tree
2281  Number of data items in the tree
3 Number of tree internal pages
11444 Number of bytes free in tree internal pages (6% ff)
29  Number of tree leaf pages
48174 Number of bytes free in tree leaf pages (59% ff)
4 Number of tree duplicate pages
642 Number of bytes free in tree duplicate pages (96% ff)
0 Number of tree overflow pages
0 Number of bytes free in tree overflow pages (0% ff)
0 Number of empty pages
0 Number of pages on the free list
```

Das einzig wichtige dabei ist folgendes Extrakt:

```
4096  Underlying database page size
3 Number of tree internal pages
29  Number of tree leaf pages
```

Die Page Size bei dn2id.bdb ist abhängig vom darunter liegenden Filesystem. Also
ein Stück weit generisch. Anders bei id2entry.

### id2entry.bdb

Dort ist die Page Size nämlich konstant 16kb. Egal welches Filesystem

``` bash
$ db_stat -h /usr/local/var/openldap-data/ -d id2entry.bdb
16384 Underlying database page size
1 Number of tree internal pages
29  Number of tree leaf pages
```

### Berechnung der Cache Size


dn2id: ( `1 root Page` + `3 internal` + `29 leaf` ) * `4KB` = `132 KB`

id2entry: ( `1 root Page` + `1 internal` + `29 leaf` ) * `16KB` = `496 KB`

Insgesamt `628 KB` was `643072 Bytes` entspricht. Da so ein DIT aber ständig
wächst möchte man vielleicht noch zwischen 20 und 50% Puffer einbauen. Ist
aber jedem selbst überlassen.

### Setzen der Cache Size

Variante 1:

Direkt in der von BDB dafür vorgesehenen Datei `DB_CONFIG`

```
vim /usr/local/var/openldap-data/DB_CONFIG
# Cache      GB  Bytes   Anzahl
set_cachesize 0  643072  1
```

Variante 2:

oder in der `/usr/local/etc/openldap/slapd.conf` - Sektion Backend/Database

```
dbconfig set_cachesize 0 643072 1
```

Noch einen kurzen `slapindex` und der Cache ist aktiv.