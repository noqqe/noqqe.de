---
comments:
- author: Patrick Kohan
  content: <p>sehr guter Artikel, hatten das mal behandelt in der Vorlesung Betriebssysteme,
    mein Vorschlag ist /dev/random oder /dev/urandom auszulesen ;-)</p>
  date: '2012-12-28T21:47:08'
- author: dAnjou
  content: '<p>Du hast leider nicht nur den alt-Text des XKCD-Comics vergessen, sondern
    auch gleich noch die Lizenz verletzt: <a href="http://xkcd.com/license.html" rel="nofollow">http://xkcd.com/license.html</a></p>'
  date: '2012-12-28T23:13:29'
- author: noqqe
  content: <p>Vielen Dank :) Jupp. Hast ne gute Zeile parat gleich nur Numerische
    Sachen auszugeben?</p>
  date: '2012-12-29T08:03:09'
- author: noqqe
  content: "<p>Ah. Title != Alt Text. Von Dilbert gab es auch einen Comic der sehr
    gut gepasst h\xE4tte, aber da konnt ich \xFCberhaupt nicht durchblicken ob man
    das jetzt darf oder nicht. </p><p>In dem FAQ geht es meistens um \"use in presentations\"
    usw. Ziemlich komisch. </p><p>Habs jetzt mal noch einen ALT Text hinzugef\xFCgt.</p>"
  date: '2012-12-29T08:15:07'
- author: sgolb
  content: '<p>Je nach Bedarf angepasst, so was: od -An -N2 -i /dev/random</p>'
  date: '2012-12-29T11:03:20'
- author: Thomas Gericke
  content: <p>+1</p>
  date: '2012-12-29T11:57:08'
- author: Lorag
  content: <p>$RANDOM liefert doch aber trotzdem keine Zufallszahlen, auch wenn die
    statischen Eigenschaften der Werte in Ordnung sind.</p>
  date: '2012-12-29T13:50:52'
- author: noqqe
  content: <p>Ich weiss nicht was du damit versuchst zu sagen. Genauer?</p>
  date: '2012-12-29T19:10:15'
- author: Lorag
  content: "<p>Zufallszahlen zu erzeugen ist f\xFCr Computer normalerweise sehr schwierig,
    meist bedarf es dazu spezieller Hardware. F\xFCr die meisten Anwendungen behilft
    man sich daher mit Zahlen, deren Generierung auf einem nicht einsehbaren Ausgangsmuster
    beruht und deren gleichm\xE4\xDFige H\xE4ufigkeitsverteilung und niedrige Korrelation
    dem echter Zufallszahlenreihen gleichen. Wenn man sich die Zahl \u03C0 als Zahlenreihe
    denkt und der Computer auf Tastendruck anfangen w\xFCrde durch die Ziffern zu
    gehen, w\xFCrde man, wenn der Computer auf einen zweiten Tastendruck stoppt, eine
    \u201Ezuf\xE4llige\u201C Ziffer zwischen 1 und 9 bekommen. </p><p>Tats\xE4chlich
    ist diese Zahl aber nicht zuf\xE4llig, sondern nur undurchschaubar. Wer den exakten
    Abstand zwischen beiden Tastendr\xFCcken kennt und dazu die Geschwindigkeit, mit
    der der PC rechnet, kann die Zahl bestimmen. Genauso, wie ich das Sample von $RANDOM
    bestimmen kann, wenn ich den seed kenne \u2013 oder in diesem Fall Zeit und PID.
    Deshalb spricht man dabei von Pseudozufallszahlen. </p><p>$RANDOM ist ein deterministischer
    Pseudozufallszahlengenerator. Wenn du RANDOM f\xFCr jeden Durchlauf des Skripts
    den gleichen Seed gibst, bekommst du exakt dieselben Zufallswerte. Das macht auch
    nichts, solange man jedes Mal reseedet und den Generator nur f\xFCr triviale Anwendungen
    ben\xF6tigt. Eine andere L\xF6sung braucht man allerdings, wenn man etwa kryptologisch
    sichere Zufallszahlen haben will. Da stehen meistens auch keine nicht deterministischen
    Generatoren zur Verf\xFCgung, weshalb viele Mathematiker an L\xF6sungen gearbeitet
    haben, die trotzdem kryptologisch sicher sind. </p><p>Unter dem Strich komme ich
    allerdings zu demselben Ergebnis $RANDOM ist f\xFCr das meiste mehr als ausreichend
    und hat vor allem den Vorteil, schnell zu sein.</p>"
  date: '2012-12-30T16:23:20'
date: '2012-12-28T18:15:00'
tags:
- shell
- stats
- zufall
- random
- bash
title: Wie der Zufall so will.
---

Ab und zu stößt man auf etwas und fragt sich "Wie kommt das eigentlich zu
stande?". So ging es mir die Woche mit `$RANDOM`. Die globale Variable der
`bash` liefert einen Wert zufälligen Wert zwischen 0 und 32767 zurück. Dieser
lustige kleine Integer erzeugende Kamerad ist mir jetzt nicht umbedingt neu.
2009 schrieb ich bei Codecocktail mal
[einen Blogpost](http://codecocktail.wordpress.com/2009/02/01/zufallszahlen-mit-der-shell-bash/)
darüber. Anders als damals interessierte mich aber was hinter der Value steht
und wie sie zustande kommt.

### Randomness in Bash

Ich lud mir den [Source](http://ftp.gnu.org/gnu/bash/) der `bash` herunter und
laß erstmal etwas. Sollte man eigentlich öfters mal tun, weil man auch über
andere unbekannte Dinge wie z.B. `$LINENO` stößt. Außerdem interessante
Comments:

> /* From "Random number generators: good ones are hard to find", Park and
> Miller, Communications of the ACM, vol. 31, no. 10, October 1988, p.
> 1195. filtered through FreeBSD */

[Das Paper](http://www.cems.uwe.ac.uk/~irjohnso/coursenotes/ufeen8-15-m/p1192-parkmiller.pdf)
(das nebenbei gesagt älter ist als ich) lässt schon bisschen erahnen wohin die Reise
geht. Außerdem erfährt man, dass der Random Number Generator von Turbo Pascal
wohl nicht so der Wahnsinn ist. Schade eigentlich.

<p><img class="center" src="/uploads/2012/12/random_number.png" title="XKCD - 211" alt="Creative Commons Attribution-NonCommercial 2.5 License http://xkcd.com/221/"></p>

Im Source etwas umgeschaut findet man aber schnell das was man sucht.
Mit etwas Bitwise XOR Spass einen Seed auf Zeit- und PID-Basis erstellt, durch den
minimalen Standard aus dem Paper gejagt und schon wars das eigentlich.

``` c
sbrand (tv.tv_sec ^ tv.tv_usec ^ getpid ());
```

### Taugt das was?

Ich wollts dann schon noch etwas genauer wissen. Mit einem Einzeiler

``` bash
S=100000000 ; time while [ $S -gt 0 ]; do echo $RANDOM >> random.txt ; ((S--)) ; done
```

habe ich dann (innerhalb 90 Minuten) 100 Millionen Zahlen generiert.
Mein Thinkpad war von dem ganzen Gerechne (immerhin 18500 Values
pro Sekunde) allerdings nicht so begeistert:

```
[45106.836033] intel ips: MCP limit exceeded: Avg temp 9513, limit 9000
[45106.836036] intel ips: MCP limit exceeded: Avg power 41641, limit 35000
```

Wahrscheinlich habe ich es etwas übertrieben. Was solls. Auf die [Sample
Size](http://en.wikipedia.org/wiki/Sample_size_determination)
kommts schliesslich an!

Die 540MB große Datei hab ich dann etwas sortiert und in einen Graphen gestopft.
Die Null-Hypothese dabei: Einer der Werte ist statistisch signifikant höher als
alle anderen. Weil für Google Charts 32k Values wohl etwas zu viel ist, ein
Graph mit `gnuplot`

{{< figure src="/uploads/2012/12/random_values_chart.png" >}}

Den vielen Values gedankt sei auch die Unübersichtlichkeit des Graphen.
Deshalb noch ein kleinwenig mehr Statistikpr0n. Im Durchschnitt wird jede
der Zahlen bei 100 Mio. Durchgängen ca. 3050 mal genannt. Wesentlich
interessanter dabei ist aber die Abweichung vom Durchschnitt (siehe auch
[Standard Deviation](http://en.wikipedia.org/wiki/Standard_deviation)).
Bei der Bestimmung derer hilft eine kleine `awk` Zeile.

``` awk
awk '{sum+=$1; sumsq+=$1*$1} END {print sqrt(sumsq/NR - (sum/NR)**2)}' sorted.txt
```

Bei meinen File wich also jeder der Werte durchschnittlich 53.6643 von der
Durchschnittshäufigkeit ab. Wenn man das weiß, hat man schonmal ein
ziemlich gutes Gefühl für das Auftreten der Integers.

Unterm Strich empfinde ich `$RANDOM` als eine mehr als zufriedenstellende
Variante eines Random Number Generators. Vielleicht war aber auch einfach
nur ein guter Tag (wörtlich!) für das zufällige generieren von Zahlen.
