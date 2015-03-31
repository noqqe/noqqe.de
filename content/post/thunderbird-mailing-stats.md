---
layout: post
title: "Thunderbird Mailing Stats"
date: 2012-08-25T11:25:00+02:00
comments: true
categories:
- Web
- Statistik
- Thunderbird
- Stats
- Debian
- Bash
- Shell
- ubuntuusers
---

Ich habe schon öfter nach einem Thunderbird Plugin gesucht, dass mir ein paar
kleine Statistiken aus meiner Mailbox bastelt. Zum Beispiel das
[tbstats](http://tbstats.sourceforge.net/) Plugin welches aber leider (wie so ziemlich alles
auf Sourceforge...) tot ist. Ich wurde durch [diesen Lifehacker Post](http://www.lifehacker.com.au/2012/03/what-lessons-could-you-learn-if-you-had-analytics-for-your-life/)
vor kurzem nochmal motiviert mir soetwas doch zu bauen, wenn es schon keine
fertige Lösung für Thunderbird dafür gibt.

## Inbox

Wie viele Mails bekommt man eigentlich so. Dieses Monat hat sichs extrem viel
angefühlt. Aber aufs Gefühl verlassen? Wo kommen wir denn da hin.

Thunderbird hält seine Daten in

> /home/&lt;user&gt;/.thunderbird/&lt;generisch&gt;.default/ImapMail/&lt;imapserver&gt;/Inbox

Ab hier ist dann ein klein bisschen Shell Zauberei nötig, damit die Daten in
gewünschter Form herausspringen:

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
Zeitstempels. Das musste
ich auch beim Wechsel von 1und1 zu Uberspace feststellen. Den Output muss man dann
nur noch zwischen ein Google Charts Template pasten und dann kann das ungefähr so aussehen:

<script type="text/javascript" src="https://www.google.com/jsapi"></script>
<script type="text/javascript">
google.load("visualization", "1", {packages:["corechart"]});
google.setOnLoadCallback(drawChart);
function drawChart() {
var data = new google.visualization.DataTable();
data.addColumn('string', 'Month');
data.addColumn('number', 'Mails');
data.addRows([
['2012-01',  93],
['2012-02',  106],
['2012-03',  129],
['2012-04',  120],
['2012-05',  226],
['2012-06',  291],
['2012-07',  204],
['2012-08',  118],
]);

var options = {
height: 500,
title: 'Mails per Month' ,
};

var chart = new google.visualization.ColumnChart(document.getElementById('colchart6_div'));
chart.draw(data, options);
}
</script>
<div id="colchart6_div"></div>

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

für jede dieser Inbox Verzeichnisse hab ich bis jetzt eine CSV Datei erstellt,
die ich nun weiter verwenden kann.


Eine modifizierte Version der Inbox Line baut mir die Daten in das folgende Format:

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

<script type="text/javascript">
google.load("visualization", "1", {packages:["corechart"]});
google.setOnLoadCallback(drawChart);
function drawChart() {
var data = google.visualization.arrayToDataTable([
['Month',  '2008',  '2009',  '2010',  '2011',  '2012'],
['01',  0,  0,  155,  105,  93],
['02',  0,  0,  102,  89,  106],
['03',  0,  0,  50,  104,  129],
['04',  0,  13,  126,  75,  120],
['05',  0,  30,  130,  146,  226],
['06',  0,  56,  102,  117,  291],
['07',  0,  48,  86,  99,  204],
['08',  0,  49,  131,  125,  118],
['09',  0,  135,  80,  166,  0],
['10',  2,  136,  154,  126,  0],
['11',  2,  157,  119,  123,  0],
['12',  4,  111,  95,  104,  0],
]);
var options = {
height: 500,
title: 'Mails per Year' ,
};

var chart = new google.visualization.LineChart(document.getElementById('chart3_div'));
chart.draw(data, options);
}
</script>
<div id="chart3_div"></div>

## Spam

Das Thema Spam ist gerade aktuell für mich interessant, da ich meine [Provider gewechselt](/blog/2012/08/17/ich-wechselte-zu-uberspace-dot-de/)
habe.

Soetwas auszuwerten erfordert natürlich etwas... Motiviation und Separation.
Motiviation weil man Spam nicht einfach löscht sondern in einen Ordner
verschiebt, den man nachher auswerten kann und Separation (gibts das Wort
eigentlich?) weil man klar unterscheiden muss was Spam ist und was nicht.

Klar die Werbemail von Paypal/Ebay  ist Spam. Aber Eigenverschuldeter. Ich
unterscheide also zwischen Mails von denke zu wissen woher Sie kommen und
Mails von Pharma- und Kasino Unternehmen.

Mit einer ähnlichen Zeile aber anderer Mailbox sieht das dann so aus:

<script type="text/javascript">
google.load("visualization", "1", {packages:["corechart"]});
google.setOnLoadCallback(drawChart);
function drawChart() {
var data = google.visualization.arrayToDataTable([
['Month',  'Spam'],
['2009-05',  23],
['2009-06',  35],
['2009-07',  24],
['2009-08',  29],
['2009-09',  37],
['2009-10',  17],
['2009-11',  1],
['2009-12',  12],
['2010-01',  4],
['2010-02',  5],
['2010-03',  12],
['2010-04',  21],
['2010-05',  2],
['2010-06',  9],
['2010-07',  12],
['2010-08',  20],
['2010-09',  30],
['2010-10',  38],
['2010-11',  34],
['2010-12',  44],
['2011-01',  25],
['2011-02',  5],
['2011-03',  3],
['2011-04',  35],
['2011-05',  41],
['2011-06',  49],
['2011-07',  29],
['2011-08',  29],
['2011-09',  20],
['2011-10',  17],
['2011-11',  20],
['2011-12',  34],
['2012-01',  21],
['2012-02',  30],
['2012-03',  32],
['2012-04',  20],
['2012-05',  55],
['2012-06',  59],
['2012-07',  38],
['2012-08',  33],
]);
var options = {
height: 500,
title: 'Recieved Spam',
};

var chart = new
google.visualization.LineChart(document.getElementById('chart2_div'));
chart.draw(data,
options);
}
</script>
<div id="chart2_div"></div>

Natürlich wieder alles nur Kosmetik und 1. Welt Probleme. Aber was will man
machen...
