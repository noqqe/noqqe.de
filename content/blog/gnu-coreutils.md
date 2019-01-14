---
comments:
- author: Anonymous
  content: Is it an ASCII key? If so you might speed up the join(1) solution further
    by first doing LC_ALL=C join ...
  date: '2014-03-25T19:45:25.850672'
- author: raimue
  content: "Fuck Coreutils. ;-)\nMein Vorschlag ist POSIX-kompatibel:\n\n$ grep -f
    <(awk '{ print \"^\"$0 }' idsubset.txt) dataset.csv\n\nDabei liest grep nur ein
    Mal die gro\xDFe Datei ein und pr\xFCft mehrere Patterns gleichzeitig, was mit
    regul\xE4ren Automaten eben sehr effizient geht. Davor muss man ein Mal vor alle
    Suchbegriffe noch das \"^\" setzen.\n\nInsgesamt garantiert schneller als die
    L\xF6sung mit join. Sp\xE4testens, wenn man die Zeit f\xFCr das Sortieren vorher
    mit reinnimmt."
  date: '2014-03-25T19:56:52.598101'
- author: noqqe
  content: Yep, its an ASCII key. Does this really speed the join command up? I will
    give it a try
  date: '2014-03-25T20:13:29.863577'
- author: noqqe
  content: "woah, coole Idee. Denk das probier ich auch mal aus. Mit dem Stdin bin
    ich mir net sicher, vielleicht mal ichs mal mit Bracket im File reingesedet. \n\nDanke!
    :)"
  date: '2014-03-25T20:15:18.452960'
- author: Clemens
  content: Process Substitution ist nicht Teil von POSIX, aber mit mkfifo kann man
    sich da schon drum rum mogeln.
  date: '2014-03-25T20:29:48.625541'
- author: raimue
  content: "Okay, da hast du recht. Dann k\xF6nnte man es auch mit einer tempor\xE4ren
    Datei l\xF6sen.\n\nWieso hat das Isso eigentlich meine ^ verschluckt? Anscheinend
    mag es die nur nicht, wenn sie in Quotes stehen: \"^\""
  date: '2014-03-25T22:32:19.291227'
- author: noqqe
  content: "ich lass das gerade laufen... es sieht nicht gut aus ... 10GB RAM f\xFCr
    den einen grep Prozess, der schon gut 5 Minuten l\xE4uft. "
  date: '2014-03-26T09:19:26.455401'
- author: noqqe
  content: "macht keinen Sinn. Hab nach 50 Minuten abgebrochen. \n\ntime grep -f idsubset.txt
    dataset.txt > result.txt\n\nreal    53m5.481s  \nuser    52m47.950s  \nsys     0m16.190s
    \ \n"
  date: '2014-03-26T10:03:57.187671'
- author: waldner
  content: "Ich schlage awk vor. Es h\xE4ngt vom Datenformat ab, aber awk sollte die
    \nleistungsf\xE4higste und schnellste L\xF6sung sein (und ohne die Dateien zu
    \nsortieren). ZB, wenn die ID das vierte CSV Feld von \"dataset.txt\" ist, \nkann
    man tun:\n\nawk -F, 'NR==FNR{a[$4]; next} $4 in a' idsubset.txt dataset.txt\n\nNat\xFCrlich
    soll man das an das konkrete Datenformat anpassen."
  date: '2014-03-30T21:29:57.230860'
- author: waldner
  content: "Sorry, das muss sein:\n\n\n\n\n\nawk -F, 'NR==FNR{a[$0]; next} $4 in a'
    idsubset.txt dataset.txt\n\ndas setzt voraus, dass jede Zeile von \"idsubset.txt\"
    eine ID ist, und dass IDs im vierten Feld von \"dataset.txt\" gesucht werden m\xFCssen.\n\nMit\n\n\n\nLC_ALL=C
    awk -F, 'NR==FNR{a[$0]; next} $4 in a' idsubset.txt dataset.txt\_\n\n\nsollte
    es noch schneller laufen."
  date: '2014-03-31T16:01:56.729736'
- author: David
  content: "\"Pauschall\xF6sung\"!? *Niemals* \"for ... in $(cat ...)\" f\xFCr das
    Zeilenweise lesen von Dateien benutzen. Daf\xFCr gibts\n\n\_ \_ while read line;
    do [COMMAND]; done < [INPUT_FILE]\n\n\nDas vermeidet das (nebem dem Laden der
    ganzen Datei in den Speicher) die shell glob expansion auf jede einzelne Zeile
    angewendet wird. Nicht dass es in diesem Fall viel hilft, aber for ... in taugt
    daf\xFCr generell nicht."
  date: '2014-05-19T11:27:14.609332'
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
