---
aliases:
- /archives/1846
comments:
- author: Ole Tange
  content: '<p>GNU Parallel takes around 2 ms per job. 0.5 ms is spent setting up
    atomic output, which you do not use. You can disable atomic output with ''-u''.</p><p>But
    your general observation still holds: If your jobs are so short that slowing down
    each job by 2 ms is a considerable amount, then GNU Parallel is probably not the
    right tool for you.</p><p>In this case you could group the md5sum in batches of
    100 and run these scripts in parallel:</p><p>    parallel "echo ''echo {1}{2}|
    md5sum &amp;&gt; /dev/null '' &gt;&gt;/tmp/md5.{2}" :::: &lt;(seq 100) &lt;(seq
    100)<br>    ls md5.* | time parallel bash</p><p>Another situation where you may
    not see a speedup is if the jobs compete for disk I/O. It is often slower to read
    two big files from the same disk simultaneously than to read them sequentially.
    To help this you can have a look at &amp;#039sem&amp;#039 which is part of GNU
    Parallel.</p>'
  date: '2012-01-09T18:16:36'
- author: noqqe
  content: <p>Hi Ole,</p><p>wow. Grouping the md5 Sums seems to be a nice way to solve
    this "problem". <br>I think I will do a retry with -u option in my posted commands
    to see if it saves ressources :)</p><p>I had never thought about the disk I/O
    situation you mentioned. Seems to make sense. </p><p>Thank you!</p>
  date: '2012-01-10T13:53:59'
date: '2012-01-08T12:47:14'
slug: gnu-parallel
tags:
- web
- shell
- seq
- parallel
- gzip
- test
- linux
- gnu parallel
- md5
title: GNU Parallel
---

Vom [GNU Parallel Projekt](http://www.gnu.org/software/parallel/) habe ich
vor einiger Zeit in der Arbeit so am Rande etwas mitbekommen. Nachdem ich
mir die gute [Dokumentation](http://www.gnu.org/software/parallel/man.html)
etwas angeschaut habe, hab ich Lust bekommen das mal selbst auszuprobieren.

{{< figure src="/uploads/2012/01/logo-gray+black300.png" >}}

Ich dachte es wäre eine gute Idee einfach ein paar md5 Summen zu bilden.

```
$ time seq 1 10000 | parallel 'echo {}| md5sum &> /dev/null '
real  0m20.102s
user  0m35.082s
sys 0m24.918s
```

Nun. Ich bilde nicht so oft 10.000 md5 Summen. War das jetzt viel? Oder
wenig? Um einen Vergleichswert zu haben sollte ich wohl auch mal nachsehen,
wie das ohne Parallel so aussieht.

```
$ time for x in $(seq 1 10000); do echo $x | md5sum &> /dev/null; done
real  0m13.504s
user  0m2.368s
sys 0m3.948s
```

Ziemlich seltsam. Obwohl ich 10.000 md5 Summen gebildet habe war die
sequenzielle Methode schneller als die Parallele. Zumindest dachte ich zu
dem Punkt noch das es seltsam ist. Aber an was lag das. Ich hab mir dann
überlegt ob ich nicht vielleicht doch noch eine andere Aufgabe als md5
Summenbildung abbilden sollte. Ich entschied mich dazu 1000 mal eine
100.000 Zeichen lange Zeichenkette durch gzip zu schubsen.

```
$ time seq 1 1000 | parallel 'cat /dev/urandom | head -c 100000 | gzip &> /dev/null'
real  0m7.845s
user  0m4.064s
sys 0m20.485s
```

7 Sekunden. Sieht eigentlich ganz nett aus. Und in der Schleife
sequenziell?

```
$ time for x in $(seq 1 1000); do cat /dev/urandom | head -c 100000 | gzip &> /dev/null; done
real  0m31.869s
user  0m8.301s
sys 0m33.658s
```

Okay. Jetzt weiss ich, das GNU Parallel eher was für (rechen-)intensivere
Aufgaben ist als für viele kleine Prozesse. Anscheinend braucht das Parsing
des zusätzlichen Binaries doch etwas zu lange um einen Prozess zu ordnen
der sowieso nach sehr kurzer Zeit wieder beendet ist. Alles in allem
gefällt mir GNU Parallel aber sehr gut wenn man weiss für was man es
einsetzen muss :)
