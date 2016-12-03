---
date: 2011-08-23T18:39:15+02:00
comments: true
title: PostgreSQL | 1000 und 1 Query
aliases:
- /blog/2011/08/23/postgresql-1000-und-1-query
- /archives/1742
tags:
- Bash
- Development
- Debian
- Databases
- 1000 queries
- 1001 query
- bash
- for
- mysql
- postgresql
- query
- sequel
- sql
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
