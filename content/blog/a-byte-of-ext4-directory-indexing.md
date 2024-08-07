---
aliases:
- /blog/2012/03/08/a-byte-of-ext4-directory-indexing/
comments:
- author: Ice Polar
  content: "<p>\_Mir pers\xF6hnlich gefallen solche Hintergrundinformationen ganz
    besonders, denn sie zeigen auf, was unter der Oberfl\xE4che alles passiert. Die
    Systeme heutzutage sind derart komplex, so dass sich keiner mehr wirklich im Klaren
    dar\xFCber ist, worauf er sich da unter Umst\xE4nden verl\xE4sst... Und so ein
    Filesystem ist doch immer wieder eine grundlegende Funktion, die man einfach mal
    so benutzt also ob das die einfachste Sache der Welt w\xE4re.</p>"
  date: '2012-03-09T19:45:48'
- author: noqqe
  content: "<p>\_Geb ich dir Recht. War auch interessant zum Debuggen. Hat Spa\xDF
    gemacht. </p>"
  date: '2012-03-09T20:09:29'
- author: sebix
  content: "<p>\_debugfs ist als root auszuf\xFChren nicht als normaler User</p>"
  date: '2012-03-12T20:14:40'
- author: noqqe
  content: "<p>\_Ich nehme an du machst das an dem Prefix $ fest, dass ich als Identifier
    f\xFCr Shell Commands verwende, # f\xFCr Kommentare. </p><p>Das ist jedenfalls
    ein Trugschluss :) </p><p></p>"
  date: '2012-03-12T20:17:59'
- author: sebix
  content: <p>Du hast vollkommen recht ;)</p>
  date: '2012-03-13T12:34:27'
- author: Ice Polar
  content: "<p>Ob jemand einmal etwas zu den Unterschieden zwischen ext4 und dem neuen
    MS REFS ( \"Resilient File System\") sagen kann?\_</p>"
  date: '2012-03-14T21:59:02'
- author: Datenbank archINFORM
  content: "<p>Hi Florian,</p><p>Hatte gerade das gleiche Problem beim zwischencache
    eines webservers und war auch erstmal ziemlich ratlos, da noch gen\xFCgend Inodes
    und Plattenplatz zur Verf\xFCgung standen. <br>\"Man schafft ein Feature,<br>dass
    dann einspringt wenn es viele Files werden (bei meinen Tests ab 300.000<br>Files),
    welches dann sp\xE4ter wegen diesen vielen Files zu fehlern f\xFChrt.\" gut auf
    den Punkt gebracht! ;)<br>Mich w\xFCrde ja jetzt ziemlich interessieren, wie sich
    denn nun die maximale Anzahl Files pro Directory bei gegebener Blockgr\xF6\xDFe,
    etc. berechnen l\xE4sst. Wirklich ein ziemliches Loch in der offiz. Doku. Diese
    Info ist ja bei entsprechenden Anforderungen schon bei der Einrichtung der Partition
    (Blockgr\xF6\xDFe/Inodeanzahl) essentiell.</p><p>Gru\xDF, Sascha\_ \_ </p>"
  date: '2012-03-20T08:11:41'
- author: noqqe
  content: "<p>\_Jap. Also in den Test-Cases habe ich immer ne max Limit Count von
    508 gefunden. Wie man das ext4 formatieren muss um da mehr zu erreichen w\xFCrde
    mich auch interessieren... <br>Wahrscheinilch ists aber statisch..</p>"
  date: '2012-03-22T18:10:37'
- author: Daniii
  content: Toller Artikel!!! hat mir sehr geholfen... haben uns dagegen entschieden
    den index einzusetzen :-)
  date: '2014-07-31T08:57:42.091639'
- author: Anonymous
  content: "Danke f\xFCr den Blogbeitrag! Der war ebenfalls einer der Gr\xFCnde, wodurch
    ich auf das Feature von ext4 verzichtet habe."
  date: '2015-10-14T12:06:16.745945'
date: '2012-03-08T17:38:00'
tags:
- index
- ext4
- dir_index
- centos
- indexing
- administration
- filesystems
- dir index
- ext
- full
- linux
- debian
title: Das Ext4 Directory Indexing
---

Mit ext4 kam unter anderem ein Feature hinzu, welches sich Directory Indexing
nennt. Es ist dazu gedacht in Verzeichnissen mit vielen Dateien eine Art Map
anzulegen an welchen Inodes welche Files liegen und kommt erst ab Fileanzahl X
pro Directory automatisch dazu. Ziel der Entwickler war natürlich Performance Gewinn.

Allerdings kauft man mit diesem Feature auch eine zusätzliche Limitierung mit
ein, bei der ich mich schwer tat wirklich (aufgrund in meiner Wahrnehmung
mangelnder Berichte und Doku) das Problem zu debuggen.

Die Limitierung bei vielen Files sind in den meisten Fällen erstmal die Inodes
der Platte. Ist die Platte aber ausreichend groß (oder die Inodes ausreichend
klein) kommt hier das Directory Indexing ins Spiel:

```
EXT4-fs warning (device /dev/sdd): ext4_dx_add_entry: Directory index full!
```

`rsync`s bzw. Schreibende Zugriffe in das Verzeichnis brechen ab und im `dmesg`
findet man obige Meldung. Problem dabei: In dem Directory Index sind so viele
Files gelistet, dass in dem (wahrscheinlich FS designtechnisch limitierten,
korrigiert mich bitte) Hashmap Block kein Platz mehr ist. Es können schlicht
und ergreifend keine neuen Einträge hinzugefügt werden.

## Was jetzt?

Längeres wälzen von Dokumentation lässt sich nicht vermeiden.

> A linear array of directory entries isn't great for performance, so a new
> feature was added to ext3 to provide a faster (but peculiar) balanced tree keyed
> off a hash of the directory entry name. If the EXT4_INDEX_FL (0x1000) flag is
> set in the inode, this directory uses a hashed btree (htree) to organize and
> find directory entries.

Oder auch andere schaurige Threads aus Mailinglisten:

[http://www.mail-archive.com/cwelug@googlegroups.com/msg01937.html](http://www.mail-archive.com/cwelug@googlegroups.com/msg01937.html)

Mit `debugfs` lassen sich die Informationen des Filesystems
auszulesen, die einen interessieren. Basic Problematik ist einfach, welches
Verzeichnis ist betroffen, wie konnte das passieren und wie bekomme ich es wieder heile.

``` bash
$ cd /var/
$ debugfs
debugfs> open /dev/sdd1
debugfs> cd log/
debugfs> htree .
[...]
Number of Entries (count): 508
Number of Entries (limit): 508
[...]
```

htree listet die Indexes in dem Ordner und Informationen zum hashed B-Tree. Die
Limit und Count Values sprechen denke ich für sich.

## Okay. Problem gefunden. Wie beheb ich es?

Man baut den Directory Index neu auf, wobei der alte Dir Index prinzipiell
behalten wird und der Neue an der letzten Stelle auf den weiteren Index
referenziert.

``` bash
fsck.ext4 -yfD /dev/sdd1
```

Dazu ist ein unmounten zwingend erforderlich, was bei produktiven Systemen
unschön ist. Oder man löscht das Verzeichnis.

Wie sich später herausgestellt hat war die Ursache davon ein nicht korrekt
funktionierendes Logrotate welches ca 8 Mio. Files produziert hat. Das Löschen
hat auch nur ca. 10 Tage gedauert. Nur so am Rande.

## Irritiert.

Zum Abschluss muss ich sagen bin ich etwas.. irritiert. Man schafft ein
Feature, dass dann einspringt wenn es viele Files werden (bei meinen Tests ab
300.000 Files), welches dann später wegen diesen vielen Files zu fehlern führt.
In dem Bereich dazwischen ist das eventuell wirklich nett zu haben, weil
Perfomance. Aber dass ich den "Referenzakt" nicht anstossen kann während das FS
gemountet ist finde ich dann speziell für den HA Betrieb von Servern ...
fragwürdig.

Klar es ist in keinem Fall eine gute Idee mehr als 2 Mio. files in *ein*
Verzeichnis zu legen, aber hey.

Wirklich gestört hat mich eher das mangelhaft Dokumentierte Vorgehen dafür.
Seitens Entwickler, seitens Community. Aber deswegen hab ich ja jetzt diesen
Post verfasst.
