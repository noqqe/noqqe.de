---
comments:
- author: Anonymous
  content: "Hallo,\_\n\ndas sieht interessant aus! Gibt es das evtl in Form eines
    Addons f\xFCr TB ohne bash/shell?"
  date: '2014-07-08T13:19:43.369275'
date: '2012-08-25T09:25:00'
tags:
- web
- shell
- stats
- thunderbird
- mail
- debian
- bash
- statistik
title: Thunderbird Mailing Stats
---

Ich habe schon öfter nach einem Thunderbird Plugin gesucht, dass mir ein
paar kleine Statistiken aus meiner Mailbox bastelt. Zum Beispiel das
[tbstats](http://tbstats.sourceforge.net/) Plugin welches aber leider (wie
so ziemlich alles auf Sourceforge...) tot ist. Ich wurde durch
[diesen Lifehacker Post](http://www.lifehacker.com.au/2012/03/what-lessons-could-you-learn-if-you-had-analytics-for-your-life/)
vor kurzem nochmal motiviert mir soetwas doch zu bauen, wenn es schon keine
fertige Lösung für Thunderbird dafür gibt.

## Inbox

Wie viele Mails bekommt man eigentlich so. Dieses Monat hat sichs extrem
viel angefühlt. Aber aufs Gefühl verlassen? Wo kommen wir denn da hin.

Thunderbird hält seine Daten in

> /home/&lt;user&gt;/.thunderbird/&lt;generisch&gt;.default/ImapMail/&lt;imapserver&gt;/Inbox

Ab hier ist dann ein klein bisschen Shell Zauberei nötig, damit die Daten
in gewünschter Form herausspringen:

``` bash
MAILDIR="/home/<user>/.thunderbird/<generisch>.default/ImapMail/<imapserver>/Inbox"
PROJECT="../data/Mail/"
$ echo "data.addColumn('string', 'Month');" > $PROJECT/Mails-per-Month-2012.csv
$ echo "data.addColumn('number', 'Mails');" >> $PROJECT/Mails-per-Month-2012.csv
$ echo "data.addRows([" >> $PROJECT/Mails-per-Month-2012.csv
$ egrep "\s*id\ \S*\ for <yourmailaddress>;" $MAILDIR/INBOX | awk '{print $8"-"$7 }' \
    | sed -e 's#Jan#01#' -e 's#Feb#02#' -e 's#Mar#03#' -e 's#Apr#04#' \
    -e 's#May#05#' -e 's#Jun#06#' -e 's#Jul#07#' -e 's#Aug#08#' -e 's#Sep#09#' \
    -e 's#Oct#10#' -e 's#Nov#11#' -e 's#Dec#12#' \
    | sort -n | grep "^20" | uniq -c \
    | awk '{print "\x27"$2"\x27, "$1}' >> $PROJECT/Mails-per-Month-2012.csv

> ['2012-01',  93],
> ['2012-02',  106],
> ['2012-03',  129],
> ['2012-04',  120],
> ['2012-05',  226],
> ['2012-06',  291],
> ['2012-07',  204],
> ['2012-08',  118],
```

Ich muss hier dazu sagen, dass sich der String des egrep von Provider zu
Provider (oder SMTP Daemon) unterscheiden kann, genau wie das Format des
Zeitstempels. Das musste ich auch beim Wechsel von 1und1 zu Uberspace
feststellen. Den Output muss man dann nur noch zwischen ein Google Charts
Template pasten.

Ich gehe hier jetzt nicht näher auf das verskripten der Templates ein, das
sollte nicht allzu schwer sein.

## Archives

Ich für meinen Teil lagere (sobald ein Jahr um ist) meine Inbox in Archives um.
Man kann sich das ungefähr so vorstellen:

```
Archives/
├── 2008
│   └── Inbox
├── 2009
│   └── Inbox
├── 2010
│   └── Inbox
└── 2011
    └── Inbox
```

für jede dieser Inbox Verzeichnisse hab ich bis jetzt eine CSV Datei
erstellt, die ich nun weiter verwenden kann.

Eine modifizierte Version der Inbox Line baut mir die Daten in das folgende
Format:

```
'2010-01', 155
'2010-02', 102
'2010-03', 50
'2010-04', 126
[...]
```

Um diese Werte nun alle in eine Zeile zu bekommen schrob ich
die folgenden Zeilen:

``` bash
echo "['Month',  '2008',  '2009',  '2010',  '2011',  '2012'],"  > $PROJECT/Mails-per-Year.csv
for y in {1..12}; do
    if [ ${#y} -eq 1 ]; then
        y="0$y"
    fi
    echo -n "[ '$y', "
    for x in {2008..2012}; do
        MAILS=$(grep "$x-$y"  $PROJECT/Mails-per-Month-$x.csv | awk '{print $2}')
        if [ -z $MAILS ]; then
            MAILS="0"
        fi
        echo -n "$MAILS, "
    done
    echo "],"
done >> $PROJECT/Mails-per-Year.csv
sed -i 's#,\s*\],\s*$#\ ],#g' $PROJECT/Mails-per-Year.csv
```

Der sed Aufruf am Ende ist ein ziemlicher Hack, aber ehrlichgesagt hätte ich das sonst in
der Loop abfangen müssen obs der letzte Eintrag ist, was umständlich gewesen wäre. Ergebnis:

## Spam

Das Thema Spam ist gerade aktuell für mich interessant, da ich meine
[Provider gewechselt](/blog/2012/08/17/ich-wechselte-zu-uberspace-dot-de/)
habe.

Soetwas auszuwerten erfordert natürlich etwas... Motiviation und
Separation.  Motiviation weil man Spam nicht einfach löscht sondern in
einen Ordner verschiebt, den man nachher auswerten kann und Separation
(gibts das Wort eigentlich?) weil man klar unterscheiden muss was Spam ist
und was nicht.

Klar die Werbemail von Paypal/Ebay  ist Spam. Aber Eigenverschuldeter. Ich
unterscheide also zwischen Mails von denke zu wissen woher Sie kommen und
Mails von Pharma- und Kasino Unternehmen.

Mit einer ähnlichen Zeile aber anderer Mailbox sieht das dann so aus:

Natürlich wieder alles nur Kosmetik und 1. Welt Probleme. Aber was will man
machen...
