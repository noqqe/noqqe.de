---
aliases:
- /blog/2011/08/23/postgresql-1000-und-1-query
- /archives/1742
comments:
- author: Knorkebrot
  content: "<p>Ich bin kein PGSQL-Experte und das, was ich hier erz\xE4hle, habe ich
    gerade eben erfragt:<br>psql geht bei einem Verbindungsaufbau deutlich mehr Securitymechanismen
    durch, was ihn viel teurer macht als bei MySQL. Ausserdem reserviert Postgre beim
    anlegen neuer Datens\xE4tze auch direkt den maximalen Speicher f\xFCr dieses Feld,
    um nicht bei \xC4nderungen evtl. nachbessern zu m\xFCssen w\xE4hrend MySQL zu
    optimieren versucht, psql muss also mehr schreiben, aber mysql l\xE4uft Gefahr
    rumkorrigieren zu m\xFCssen.</p><p>Interessant f\xE4nde ich deswegen eher, wie
    schnell es w\xE4re, wenn du eine Datei mit deinen 1000 SQL-Statements hast du
    die in einem Rutsch in psql und mysql reinschiebst. W\xFCrde zumindest eventuelle
    Differenzen beim Verbindungsaufbau vermeiden.</p><p>M\xDFG</p>"
  date: '2011-08-30T16:41:44'
- author: noqqe
  content: "<p>@Knorkebrot: Wow, nice. Okay hab ich wieder was gelernt. Bei wem haste
    das erfragt? </p><p>Naja _aktuell_ ist das mit dem 1 connection per action leider
    mein Use-Case. Ein Event in ZRE passiert und zugleich wird es in die DB geschrieben.
    Vielleicht sollte ich die aber auch zu gr\xF6\xDFeren Files zusammenmergen und
    dann in einem Rutsch reinschreiben... mh.</p>"
  date: '2011-09-03T00:53:36'
- author: Knorkebrot
  content: "<p>Von meinem (damaligen) Ausbilder habe ich das.<br>Ich bin nicht unbedingt
    sehr fit im Shellscripting, kann man vielleicht die for-Schleife in Klammern setzen
    und deren Output in psql pipen? Dann h\xE4tte man das Problem auch nicht mehr.
    Habe leider gerade nichts zum testen da (bzw. meine neue Maschine baut gerade
    ;) )</p>"
  date: '2011-09-04T12:49:04'
date: '2011-08-23T16:39:15'
tags:
- development
- 1000 queries
- postgresql
- for
- 1001 query
- sequel
- mysql
- sql
- databases
- query
- debian
- bash
title: PostgreSQL | 1000 und 1 Query
---

Zur Zeit spiele und bastle ich nebenher mit PostgreSQL rum. Überlege ob ich
mal eine alternative DB für das
[Zombie-Revolution-Environment](http://zombies.n0q.org) an den Start
bringe...

{{< figure src="/uploads/2011/08/5384_c4d1.jpeg" >}}

Für meinen Use-Case scheint das allerdings nur begrenzt von Nutzen zu sein.
Ich mache vielleicht etwas falsch, aber wenn ich 1000 Queries in MySQL
kippe, dauert nur einen Bruchteil so lange wie in postgreSQL. Um das zu
veranschaulichen:

```
$ time for x in $(seq 1 1000) ; do mysql -u root -ppw -e "insert into foobar.foo values ($x, now());" ; done
real    0m7.349s
user    0m0.060s
sys     0m0.380s
```

```
$ time for x in $(seq 1 1000) ; do psql --quiet -d foobar -c "insert into foobar values ($x, now());" ; done
real    1m28.363s
user    0m37.450s
sys     0m13.020s
```

Kann mir jemand sagen woran das liegt? Ich kann mir nur schwer vorstellen
das PostgreSQL so hinterher hinkt.