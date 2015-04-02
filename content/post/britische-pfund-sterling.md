---
type: post
title: "Britische Pfund Sterling"
date: 2013-06-17T20:45:00+02:00
comments: true
categories:
- ubuntuusers
- osbn
- Development
- Stats
tags:
- Pfund
- Euro
- Kurs
- R
- Statistik
- CSV
- XML
- EZB
---

Im Juli werde ich wohl etwas Sightseeing auf den britischen Inseln betreiben.
Es ist anscheinend eine gute Idee etwas lokales Bargeld zu besitzen.
Ich fragte Siri ob ich morgen Euro gegen Pfund tauschen solle.

{{< figure src="/uploads/2013/06/kurs.png" >}}

Wechselkurse sind Dinge für die ich mich noch nie interessiert hab. Wann
geh ich zur Bank. Wann ists teuer, wann nicht?  Eigentlich wollte ich das
schon ewig gemacht haben da man munkelte, der Kurs wäre gerade gut.
Irgendwie hab ich es aber nicht geschafft.

Die EZB bietet auf Ihrer (maximal fragwürdig aussehenden)
[Website](http://www.ecb.int/stats/exchange/eurofxref/html/eurofxref-graph-gbp.en.html)
[XML Files](http://www.ecb.int/stats/exchange/eurofxref/html/gbp.xml)
mit Rohdaten an. OpenData, alter! XML liess
sich in dem Fall leicht zu CSV konvertieren

    sed -e 's#.*TIME_PERIOD=\(.*\)\ OBS_VALUE=\"\(.*\)"\ OBS_STATUS.*#\1 \2#' gbp.xml > gbp.csv

Alles in `R` geworfen. CSV importiert, Tag dazugerechnet.

``` r
> f <- read.csv(file="gbp.csv", header=F, as.is=T, sep=" ")
> f$Tag <- format(as.Date(f$Datum), format="%A")
> f
[...]
3679 2013-05-14 0.84815   Tuesday
3680 2013-05-15 0.84640 Wednesday
3681 2013-05-16 0.84550  Thursday
3682 2013-05-17 0.84475    Friday
3683 2013-05-20 0.84560    Monday
3684 2013-05-21 0.84910   Tuesday
3685 2013-05-22 0.85570 Wednesday
3686 2013-05-23 0.85515  Thursday
[...]
```

Die Daten spiegeln übrigens (was man den Metadaten des XMLs entnehmen kann) dem
täglichen Tauschwert um 2:15 pm (C.E.T.) wider. Dass die Daten immer nachmittags
aufgezeichnet wurden ist ein bisschen schade, vielleicht wäre es morgens
billiger gewesen? Kann man aber nicht helfen.

``` r
> kursmean <- NULL
> days <- unique(f$Tag)
> for (n in days) kursmean <- c(kursmean,with(f, mean(Kurs[ Tag == n ])))
> names(kursmean) <- days
> kursmean
   Monday   Tuesday Wednesday  Thursday    Friday
   0.7306814 0.7303953 0.7303909 0.7294906 0.7290670
```

Somit scheint der Montag ein guter Tag zu sein um Pfund zu tauschen. Auch der aktuelle
Kurs ist mit 0.84 irgendwo im 3. und 4. Quarter. Annehmbar. Nach der selben
Methodik kann man auch beim Monat vorgehen. Leider bringt mir das herzlich
wenig, da März schon vorbei ist. Gna.
