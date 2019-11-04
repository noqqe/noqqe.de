---
date: 2016-12-11T17:47:30+01:00
tags:
- bookmarks
- pinboard
- R
- instapaper
title: Commandline Bookmarks don't work (for me)
---

**TL;DR**: 2013 zog ich meine Lesezeichen aus dem Browser zu Pinboard.
[2015](/blog/2015/04/06/bookmarks/) von Pinboard auf die Cmd zu
[bm](https://github.com/noqqe/bm) und später [rvo](https://github.com/noqqe/rvo).
Jetzt wieder zurück zu [Pinboard](https://pinboard.in).

Fast 1,5 Jahre sammle ich nun Lesezeichen und Weblinks in der Commandline.
Besser gesagt wurde `rvo` dafür etwas umfunktioniert.

{{< figure src="/uploads/2016/12/wowcommandline.jpg" >}}

Warum hab ich nun meine Lesezeichen wieder zurück nach Pinboard gezogen? Auch nach der ganzen Zeit fällt es mir immernoch schwer mich zu motivieren für etwas wirklich ein Lesezeichen _anzulegen_. Es ist
eben ein Wechsel des Programms. Man verliert den Fokus im Browser um das gefundene
kurz mal ins Terminal zu klopfen. Dadurch machte ich gefühlt viel viel
weniger Bookmarks und erwischte mich häufig dabei zum 3. mal nach dem
selben Scheiss zu googeln^Wduckduckgoen. Da ist es leichter einfach auf
einen Button im Browser zu drücken.

Migriert war alles eigentlich sehr schnell. Das Format eines Pinboard
Backups ist angenehmerweise in `json`. Das konnte ich mittels ein paar
Zeilen Python und `rvo` export zurecht fricklen.

Dann war alles wieder auf Pinboard. Tags, Datum usw. alles schön.

### "Ist das so?"

Stimmt es eigentlich wirklich das ich weniger bookmarke?
Um das herauszufinden hab ich mal wieder die `R` Console angeworfen.

``` R
library("rjson")
j <- fromJSON(file="~/Downloads/pins.json")

# convert list to a char vector
time = c()
for (l in j[1:length(j)]) { time = c(time, l[["time"]]) }

# strip away day and hours
months = substr(time, 0, 7)

# converting a table to a dataframe is a nice way
# of counting occurrences
c = as.data.frame(table(months))

# make line chart
plot(c)
lines(c)
```

Eigentlich sehen die Zahlen garnicht so schlimm aus wie ich dachte. Hatte
wirklich mit viel weniger Bookmarks gerechnet. Aber woher soll man es
wissen, wenn man es nicht nachprüft, oder?

Was auch immer 2013 los war..

{{< figure src="/uploads/2016/12/graph.png" >}}

Die Zahlen beweisen also erstmal nicht das ich signifikant weniger
Bookmarks erstelle. Was ich leider nicht nachmessen kann ist aber wie viele
davon ich **brauche** und wie viel Schmerz es macht diese zu erstellen.
In beiden Fällen muss ich mich auf mein Gefühl verlassen.

Mit ein paar kleinen Änderungen bekomm' ich natürlich auch die
Jahreszahlen. Diese ändern aber auch nicht wirklich etwas an der Tatsache.

``` R
> year = substr(time, 0, 4)
> c = as.data.frame(table(year))
> print(c)
  year Freq
1 2009   17
2 2010   72
3 2011   40
4 2012   78
5 2013  539
6 2014  287
7 2015  227
8 2016  128
```

Zum Glück ist das Jahr gerade zuende und ich muss 2016 nicht extrapolieren
:)

### Read it later

Dafür hatte ich mir [Instapaper](https://instapaper.com) angeschaut. Ein
wirklich schöner Service. Für [entbehrlich.es](https://entbehrlich.es)
hatte ich das Bedürfnis Wikipedia Artikel irgendwo zu archivieren und
Notizen dazu zu machen. Ich nutzte das ca. 4 Monate bis mir wegen des
Pinboard Umbaus klar wurde, dass es eigentlich nur eine andere Lösung für
das gleiche Problem ist. Ich brauche auf allen Geräten Zugang zu einem
Linkverzeichnis. Also hab ich auch meinen Instapaper Account geschlossen
und mit einem [Python Script](https://gist.github.com/noqqe/98acc173dd19a7213b84a8cf15409c4d)
die Inhalte umgezogen.

Frustrierend dagegen war, dass es keine Möglichkeit gibt das Datum eines
Eintrags zu exportieren.

Das schöne (bei Instapaper eingebettete) aufbereiten eines Posts werde ich
mal versuchen mit dem "[Reader View](https://support.apple.com/kb/PH21467?locale=en_US)" aus Safari zu ersetzen
