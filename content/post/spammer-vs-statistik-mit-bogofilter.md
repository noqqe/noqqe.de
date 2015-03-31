---
type: post
title: "Spammer vs. Statistik mit bogofilter"
date: 2013-10-26T17:40:00+02:00
comments: true
categories:
- Spam
- Bogofilter
- Statistik
- Stats
- ubuntuusers
- osbn
- Bayesian
- Filter
- maildrop
- email
- mail
- bogo
---

Als ich [das letzte Mal](/blog/2012/12/08/spam/) über Spam
schrieb wurde mir [bogofilter](http://bogofilter.sourceforge.net/)
empfohlen. Die Software setzt die Idee [Paul Grahams](http://paulgraham.com/) des [Better Bayesian Filtering](http://paulgraham.com/better.html) um. Im Endeffekt geht es
um statistische Auswertung des Contents.

{{< figure src="/uploads/2013/10/warumnur.gif" >}}

## Setup

Praktisch wird `bogofilter` trainiert. Was ist Ham, was Spam.

```
$ bogofilter -s -B /home/noqqe/Maildir/.Spam/
$ bogofilter -n -B /home/noqqe/Maildir/INBOX
```

Einmal angelernt ensteht eine wordlist.db im BerkeleyDB Format in der die
erlernten Wörter mit Good/Bad Scores abgelegt werden.  In meinem Fall
habe ich Anbindung Serverseitig über maildrop, dass ich sowieso schon
verwende, realisiert. Ja maildrop und ja die 90er sind vorbei.  Sorry.

{% codeblock "integration in maildrop" %}

xfilter "bogofilter -u -e -p -R -c /home/noqqe/.bogofilter.cf"

# Filter with bogofilter Spam or Unsure
if ( /^X-Bogosity: Spam, tests=bogofilter/:h ) {
  to "$VUSERMAILDIR/.$SPAMDIR/"
}
```


## Spam, aber warum?

Persönlich am Herzen liegt mir immer das "warum?". Wie hart verbose
sich `bogofilter` sich bei der Ausgabe des "warum" verhält lässt sich grob
in 3 Stufen in `-v` bis `-vvv` einstellen. Zusätzlich dazu beherrscht
`bogofilter` aber auch noch den Parameter `-R`, der in den Mailheader
für die Programmiersprache `R` kompatiblen Output generiert.

```
$ cat 927470317.21490.txt | bogofilter -R
X-Bogosity: Spam, tests=bogofilter, spamicity=1.000000, version=1.2.4
                                  n        pgood     pbad      fw     U
"rcvd:HELO"                      2283687 0.991322  0.216304  0.179115 -
"button"                          39285  0.013171  0.004399  0.250394 -
"spam"                            41359  0.010381  0.005241  0.335478 -
"from:Vigara"                      2110  0.000000  0.000360  0.999996 +
"IMPORTANT!"                       2189  0.000000  0.000373  0.999996 +
"subj:Cilais"                      8492  0.000000  0.001449  0.999999 +
"subj:day"                        25423  0.000000  0.004336  1.000000 +
"NEW!"                            27717  0.000000  0.004728  1.000000 +
[...]
"subj:Vigara"                     27837  0.000000  0.004748  1.000000 +
"Vigara"                          34204  0.000000  0.005834  1.000000 +
"head:IPS"                        35711  0.000000  0.006091  1.000000 +
"Cilais"                          37712  0.000000  0.006433  1.000000 +
"Levtira"                         37987  0.000000  0.006480  1.000000 +
"Gifts"                           39115  0.000000  0.006672  1.000000 +
N_P_Q_S_s_x_md                       42  0.000000  1.000000  1.000000
```

Am Anfang wirkt das alles etwas verwirrend, macht aber Sinn. Die Anzahl
des Vorkommens in der wordlist.db, das Rating in wie vielen Mails das
"Wort" gut war, in wie vielen schlecht, ein gewichteter Index aus good
und bad, dann noch + oder -. Also ob der Wert zum Rating subtrahiert
oder addiert wird.

Ich schrob deshalb "Wort" weil Wort im Kontext von `bogofilter` eine
bestimmte Bedeutung hat. Die meisten wortbasierten Tools separieren
Wörter einfach nach bestimmten Limitern wie Space, "-" oder
Sonstiges. Herr Graham hat aber im oben verlinkten Artikel auch
Vorschläge für bestimmte Eigenschaften eines bayesian Filters im Bereich
Spam-Mails vorgebracht. So werden beispielsweise Words in Header-Fields wie
Subject, From, To das entsprechende Feld vorangestellt. Beispielsweise
`subj:Cialis` (siehe oben). Späßchen wie diese sind es, die die hohe
Trefferquote von bogofilter ausmachen. \*hust\* und die Trainingdaten \*hust\*.

## Trainingsets

So viel Spam auch in mein Postfach trudelt, mehr Daten sind mehr Daten
und mehr Daten sind besser. Ich bin mir garnicht mehr sicher wie genau
ich auf den Link zu [untroubled.org/spam](http://untroubled.org/spam)
gekommen bin.

Wie dort Spam gesammelt wird ist einfach erklärt: Catch-All. Der über
die Jahre zusammengekommene Spam wird in `7z` Archiven zur Verfügung
gestellt. Diesen habe ich mir dann heruntergeladen. Herzliches Dankeschön
an dieser Stelle btw.

```

$ mkdir archives ; cd archives
$ wget http://untroubled.org/spam/{1998..2012}.7z
$ wget http://untroubled.org/spam/2013-{01..12}.7z

$ for x in $(echo *.7z); do
>  mkdir ../${x/.7z} ;
>  7z e $x -o../${x/.7z}
> done
```

Dabei aber aufpassen. Ich habe die Schleife benutzt und für das
auspacken etwas über 2,5 Tage(!) gebraucht. Problem war aber
nichtmal Disk IO sondern einfach nur die CPU Auslastung. Mehr Cores
haben macht hier mehr Spass.  Würde ich die Files nochmal auspacken
müssen, würde ich wohl eine Variante a la `make -j 4` oder gleich
[GNU Parallel](/blog/2012/01/08/gnu-parallel/) wählen.

{{< figure src="/uploads/2013/10/wakeup.gif" >}}

Nach dem Training der ausgepackten Daten ist meine wordlist.db von 12MB
auf eine größe von 1,4GB gewachsen. Ordentlich. Natürlich gehört
nach dem ganzen Spam auch noch eure normale Inbox als Ham angelernt.
Nicht vergessen, sonst blöd.

Fahre mit dieser wordlist.db jetzt seit ca. 2 Wochen und meine Ergebnisse
sind gefühlt besser, aber für eine wirklich verlässliche Auswertung
ist es noch etwas zu früh.

Wer sich mit der Methodik (auch als nicht Mathematiker) mal vertraut
machen möchte kann ich auch
[Probabilistic-Programming-and-Bayesian-Methods-for-Hackers](http://camdavidsonpilon.github.io/Probabilistic-Programming-and-Bayesian-Methods-for-Hackers/)
herzlichst empfehlen.
