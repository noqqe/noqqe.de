---
title: "Strava GPS Tracks analysieren - Teil 3 (Die Rache)"
date: 2020-01-12T14:07:08+02:00
tags:
- strava
- gpx
- gps
- python
- gpsbabel
---

Um nach [Teil 1](/blog/2017/05/03/strava-gps-tracks-analysieren/) und
[Teil 2 - jetzt erst Recht](blog/2018/05/16/strava-gps-tracks-analysieren-2/)
jetzt eine Trilogie aus der Blogreihe zu machen, heute nun ein weiterer Teil.

Nach dem immer wiederkehrenden Schmerz [R](https://www.r-project.org/), das
dazu passende Rstudio und alle Libraries auf aktuellen Stand zu kriegen,
reichte es mir dieses Jahr. Ich hatte bereits mehrere Versuche gemacht und es
immer wieder verworfen.

{{< figure src="/uploads/2020/01/i-have-gps.jpg" >}}

## Die dritte Iteration - Ziele

Ich will ein einfaches kleines Python Script welches alljährlich
wiederverwenden kann um meine Tracks dieses Jahr zu analysieren.

Als ich gestern von der ersten Tour nach Hause kam, hab ich mich 2 Stunden
hingesetzt und versucht das einfach ganz MVP mäßig hinzukriegen. 1 GPX File
mittels Stack-Overflow-Engineering auf eine Karte zeichnen.

{{< figure src="/uploads/2020/01/stravamap.png" >}}

Ich wurde mit [Folium](https://github.com/python-visualization/folium) auch
fündig.

## Daten vorbereiten

Ansonsten das übliche Prozedere, bei Strava einen Datenexport anfordern und
entpacken. Witzigerweise ist dieser Export mittlerweile bei "Ja ich möchte
meinen Account wirklich löschen" versteckt und nicht mehr in der Data
Sektion.

Auch der Inhalt des Dumps hat sich stark verändert. Wo vorher ein Unix
Timestamp der Name der Datei mit `.gpx` angehängt war, findet sich
mittlerweile im `activities` Ordner eine Vielzahl an unterschiedlich
Formatierten Dateien.

* gpx (Was wir wollen)
* gpx.gz (Ebenfalls okay)
* tcx (Datenformat für Heimtrainer, also nicht relevant)
* fit (Binärformat das ich noch nicht kannte)
* fit.gz (selbes, in gezippt)

Wir wollen möglichst alles davon als minifiziertes GPX haben.

```
cd activities/

unzip  *.gz

for x in *.fit
  gpsbabel -i garmin_fit -f $x -o gpx -F $x.gpx
end

# Good old minifier friend, aus Teil 2
for f in *.gpx
  set -l points (math (grep -c "trkpt" $f) / 30)
  gpsbabel -rt -i gpx -f $f -x simplify,count={$points} -o gpx -F SIMPLIFIED-$f
end
```

Obacht, `fish` Syntax, kein `bash`.

## Das Tool

Ich hab mein kleines Projekt [stramap](https://github.com/noqqe/stramap)
getauft und auf Github gestellt.

Im Skript kurz die Config anpassen und `./stramap` ausführen.

```
# gpxfiles = [ 'test.gpx', 'test2.gpx' ]
gpxfiles = glob.glob('./activities/*.gpx')

# What years to filter (None == All years)
years = list(range(2012, 2020))
years.append(None)
```

Mit dem Ergebnis bin ich sehr zufrieden, es ist zum ersten mal eine
Interaktive Map mit ALLEN Tracks die man wirklich noch bedienen kann, selbst
wenn man die Daten nicht minified.

```
> ./stramap
[+] Wrote 2014 to ./index.2014.html
[+] Plotting 2015
[+] Wrote 2015 to ./index.2015.html
[+] Plotting 2016
[+] Wrote 2016 to ./index.2016.html
[+] Plotting 2017
[+] Wrote 2017 to ./index.2017.html
[+] Plotting 2018
[+] Wrote 2018 to ./index.2018.html
[+] Plotting None
[+] Wrote None to ./index.None.html
```

{{< figure src="/uploads/2020/01/stramap.png" >}}

Kommentare gerne dazu :)
