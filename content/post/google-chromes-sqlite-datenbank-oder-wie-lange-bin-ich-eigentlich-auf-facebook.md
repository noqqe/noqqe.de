---
type: post
title: "Google Chromes SQLite Datenbank oder &quot;Wie viel Zeit verschwende ich eigentlich auf Facebook?&quot;"
date: 2012-08-26T12:30:00+02:00
comments: true
categories:
- ubuntuusers
- Development
- Debian
- Bash
- Shell
tags:
- Google
- Chrome
- Browser
- Browsing
---

Im O'reilly Blog habe ich letztens einen [interessanten Ausschnitt](http://community.oreilly.de/blog/2012/08/10/auszug-aus-computer-forensik-hacks-teil-3/) aus
dem Buch Computer Forensik Hacks gefunden. Wie man hier in letzter Zeit merken
kann, begeistern mich solche Themen, also ausprobiert und etwas rumgespielt.

## Welche Sites besuche ich am häufigsten?

In der SQLite Datenbank unter

> /home/&lt;user&gt;/.config/google-chrome/Default/History

speichert Chrome seine History Daten. Einen der Queries die ich als erstes probiert habe, war herauszufinden welche
Seiten ich eigentlich am meisten Besuche

``` sql
SELECT urls.title, urls.visit_count
  FROM urls
  ORDER BY urls.visit_count DESC
  LIMIT 20;
```

Der Query sieht eigentlich relativ übersichtlich aus und ebenso die Auswertung

```
Facebook|1120
Google Reader|902
SPIEGEL ONLINE - Nachrichten|385
Soup Freunde|372
noqqe.de|337
Wiki - Front Page|330
GitHub|280
[...]
```

## Wie oft Suche ich auf Google?

Für Searches pflegt Chrome hier in einer eigenen Table. Simpler
Inner Join, alles klar.

``` sql
SELECT SUM(keyword_search_terms.keyword_id)
  FROM keyword_search_terms, urls, visits
  WHERE urls.id = keyword_search_terms.url_id
  AND visits.url = urls.id;
```

13154 Google Searches. Mich würde aber noch interessieren wie sich das verteilt
pro Monat oder so vielleicht.

``` bash
DEST=/home/noqqe/.config/google-chrome/Default/History
for x in {1..12}; do
    if [ ${#x} -eq 1 ]; then x="0$x"; fi
    echo -n "2012-$x - "
    SEARCHES=$(sqlite3 -list $DEST "SELECT SUM(keyword_search_terms.keyword_id) from keyword_search_terms, urls, visits WHERE urls.id = keyword_search_terms.url_id AND \
        visits.url = urls.id AND datetime(visits.visit_time/1000000-11644473600,'unixepoch', 'localtime') LIKE \"$(date +%Y-)$x%\" ;" )
    if [ -z $SEARCHES ]; then SEARCHES=0 ; fi
    echo $SEARCHES
done

2012-01 - 0
2012-02 - 0
2012-03 - 0
2012-04 - 0
2012-05 - 2504
2012-06 - 5717
2012-07 - 3351
2012-08 - 1582
2012-09 - 0
2012-10 - 0
2012-11 - 0
2012-12 - 0
```

So sieht das schon besser aus :) Zu dem seltsam aussehenden Date String gleich mehr.

## Wie lange surfe ich eigentlich so im Monat?

``` sql
SELECT SUM((strftime('%s',datetime(visits.visit_duration/1000000-11644473600,'unixepoch', 'localtime')) - strftime('%s','1601-01-01 00:00:00')))
  FROM urls, visits
  WHERE urls.id = visits.url
  AND datetime(visits.visit_time/1000000-11644473600,'unixepoch', 'localtime') LIKE '2012-06%';
```

Seien wir ehrlich, das sieht schlimm aus. Warum? Weil Chrome seine Zeitstempel
nicht im Default Epoch abspeichert sondern in Mikrosekunden seit
dem 1.1.1601, 00:00:00. Das ist ziemlich doof und aufwändig.

``` bash
_hms() {
 local S=${1}
 ((h=S/3600))
 ((m=S%3600/60))
 ((s=S%60))
 printf "%d-%d-%d\n" $h $m $s
}

for x in {1..12}; do
    if [ ${#x} -eq 1 ]; then x="0$x"; fi
    echo -n "2012-$x "
    SECONDS=$(sqlite3 -list $DEST "SELECT
    SUM((strftime('%s',datetime(visits.visit_duration/1000000-11644473600,'unixepoch','localtime')) \
    - strftime('%s','1601-01-01 00:00:00'))) from urls, visits WHERE urls.id = visits.url \
      AND datetime(visits.visit_time/1000000-11644473600,'unixepoch', 'localtime') LIKE \"$(date +%Y-)$x%\" ")
    _hms $SECONDS
done

2012-01 0-0-0
2012-02 0-0-0
2012-03 0-0-0
2012-04 0-0-0
2012-05 1111-0-0
2012-06 4415-0-0
2012-07 6344-26-39
2012-08 13182-33-34
2012-09 0-0-0
2012-10 0-0-0
2012-11 0-0-0
2012-12 0-0-0
```

Mh. 13.182 Stunden im August. Das scheint mir etwas übertrieben. Bei der
Auswertung muss man aber verstehen, dass es sich hier bei um die Zeit handelt,
in der der jeweilige Tab geöffnet ist. Soll konkret heissen: Mein Monat hat
keine 13.182 Stunden :P

Wenn ich mit offenen 100 Tabs eine Stunde rumchrome' bekomme ich schnell +100
Stunden. Die Zeit der tatsächlich angesehenen Tabs (das was man hier eigentlich
will) gibt die SQLite Datenbank von Chrome nicht her. Schade...

Trotzdem halte ich die Kennzahl für aussagefähig.

## Wie viel Zeit verbringe ich eigentlich auf Facbook?

Aber nun zur Eingangs erwähnten Frage. Ein bisschen kombinierte Queries von oben
hier und da und schon kommt das raus was man möchte:

``` bash
for x in {1..12}; do
    if [ ${#x} -eq 1 ]; then x="0$x"; fi
    echo -n "2012-$x - "
    sqlite3 -list $DEST "SELECT COUNT(urls.id) FROM urls, visits WHERE urls.id = visits.url AND urls.url LIKE '%facebook.com/%' AND \
    datetime(visits.visit_time/1000000-11644473600,'unixepoch', 'localtime') LIKE \"$(date +%Y-)$x%\" ;"
done

2012-01 - 0
2012-02 - 0
2012-03 - 0
2012-04 - 0
2012-05 - 130
2012-06 - 208
2012-07 - 275
2012-08 - 487
2012-09 - 0
2012-10 - 0
2012-11 - 0
2012-12 - 0
```

Jetzt kenn ich meine Visits..aber wie siehts mit der Dauer aus?

``` bash
for x in {1..12}; do
    if [ ${#x} -eq 1 ]; then x="0$x"; fi
    echo -n "2012-$x - "
    SECONDS=$(sqlite3 -list $DEST "SELECT SUM((strftime('%s',datetime(visits.visit_duration/1000000-11644473600,'unixepoch', 'localtime')) - strftime('%s','1601-01-01 00:00:00'))) \
        from urls, visits WHERE urls.id = visits.url AND datetime(visits.visit_time/1000000-11644473600,'unixepoch', 'localtime') LIKE \"$(date +%Y-)$x%\" AND urls.url LIKE '%facebook.com/%' ")
    if [ -z $SECONDS ]; then SECONDS=0 ; fi
    _hms $SECONDS
done

2012-01 - 0-0-0
2012-02 - 0-0-0
2012-03 - 0-0-0
2012-04 - 0-0-0
2012-05 - 130-0-0
2012-06 - 208-0-0
2012-07 - 275-52-43
2012-08 - 529-17-21
2012-09 - 0-0-0
2012-10 - 0-0-0
2012-11 - 0-0-0
2012-12 - 0-0-0
```

Das hab ich auch noch für 9gag, Google Reader und andere Seiten auf denen ich
so abspacke durchgespielt. Aber am Ende ist es eigentlich egal, ob man sich
Social Media Foo reinzieht oder sich auf einer kommerzialisierten Meme Plattform
rumtreibt. Am Ende ist's Zeitverschwendung...
