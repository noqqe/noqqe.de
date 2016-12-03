---
date: '2014-03-25T17:18:00'
tags:
- shell
- join
- for
- gnu
- admin
- iowait
- coreutils
- flatfile
- time
- grep
- csv
- id
- bash
title: GNU Coreutils
---

Man stelle sich folgendes Szenario vor. Eine große CSV Datei enthält Datensätze.
Eine weitere Datei enthält ~1,5mio IDs die ein Subset der Datensätze darstellen.
Gewünscht ist ein File das alle Datensätze des Subsets enthält.

{{< figure src="/uploads/2014/03/stallman.jpg" >}}

### for-loop grep

Die gewohnte Pauschallösung für derartige Probleme. Ganz im Bash-Admin-Stil

``` bash
$ time for x in $(cat idsubset.txt) ; do
>  grep ^$x dataset.csv
> done > result.csv
```

Nur leider kommen dabei ganze 1,5 Records pro Sekunde heraus, was alles in allem
in über 2 Wochen Rechenzeit endet. `IOwait` enstand dabei nicht.

### GNU parallel

16 Core-Maschine. Einfach härter parallel greppen. [GNU parallel](https://www.gnu.org/software/parallel/)
hatte ich 2012 einmal [ausprobiert](https://noqqe.de/blog/2012/01/08/gnu-parallel/).

``` bash
$ cat idsubset.txt | time parallel 'grep -m 1 ^{} dataset.csv' > result.csv
[...]
Command terminated by signal 2
13165.04user 56967.06system 1:23:04elapsed 1406%CPU (0avgtext+0avgdata 40816maxresident)k
```

Nach knapp 90 Minuten war das gute Stück bei ca. 80% des Files angekommen.
Annehmbar, auch wenn die Cores und der RAM der Kiste damit gut beschäftigt
waren.

### join

Das effizienteste war allerdings `join` aus den [GNU core utilities](https://www.gnu.org/software/coreutils/)

``` bash
$ sort idsubset.txt > sidsubset.txt
$ sort dataset.csv > sdataset.csv
$ time join sidsubset.txt sdataset.csv > result.txt
[...]
real    0m38.965s
user    0m36.290s
sys     0m0.991s
```

Fucking 38 Sekunden. Zwei Dinge sind zu beachten. Sortierung und
Formatierung.

Das Field, das zusammengeführt werden soll _muss_ in beiden Files über den
gleichen Trenner identifizierbar sein. Zurecht-ge-`sed`-et&copy;

Beide Files müssen alphabetisch sortiert sein, nicht numerisch. Das ist im
wesentlichen dem Algorithmus geschuldet der in `join` verbaut ist. Linecounts
anstelle von Fullscans bei jeder Iteration sind der Trick.

BigData Krams? Lolo. Fucking [Coreutils](http://rms.sexy).