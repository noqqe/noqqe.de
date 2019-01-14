---
title: "Strava GPS Tracks analysieren - Teil 2 (jetzt erst Recht)."
date: 2018-05-16T17:07:08+02:00
tags:
- strava
- gpx
- gps
- R
- gpsbabel
---

Vor einer Weile hatte ich bereits mal einen
[Post](/blog/2017/05/03/strava-gps-tracks-analysieren/) über GPX Files mittels
R zu visualisieren. Also wieder über Strava Einstellungen ein Archiv aller
Aktivitäten anfordern, dann auspacken und im Ordner `activities` die GPX
Dateien finden. Abhängigkeiten installieren:

```
$ R
> p <- c("plotKML", "maps", "ggmap")
> install.packages(p)
```

und die leicht aktualisierte Version als
[gist](https://gist.github.com/noqqe/55542812068d02d0996bb03847a20d3f) nach
angepassten Einstellungen ausführen. Der Unterschied ist eigentlich nur das
ich diesmal pro Jahr ein einzelnes File bekomme:

{{< figure src="/uploads/2018/05/strava-2018.png" >}}

Sinn dahinter ist eigentlich, dass ich gerne ein GIF erstellt hätte, welches
dann die Jahre einzeln durchscrollt. Das (mich immer wieder überraschende)
ImageMagick gibt einem dafür `convert` zur Hand.

```
$ ls
.rw-r--r--  1.4M strava-progress-2011.png
.rw-r--r--  1.4M strava-progress-2012.png
.rw-r--r--  1.4M strava-progress-2013.png
.rw-r--r--  1.4M strava-progress-2014.png
.rw-r--r--  1.4M strava-progress-2015.png
.rw-r--r--  1.4M strava-progress-2016.png
.rw-r--r--  1.4M strava-progress-2017.png
.rw-r--r--  1.4M strava-progress-2018.png
$ convert  -delay 70 -loop 100 strava-progress-*.png strava-progress.gif
```

und das Ergebnis schaut ganz gut aus :)

{{< figure src="/uploads/2018/05/strava-progress.gif" >}}

Alles schön und gut, nur dauert die Erstellung der einzelnen Dateien ultra
lange... Fast 40 Minuten muss man wie innerlich tot auf den Bildschirm
starrend Lebenszeit verschwenden. Aber dann erinnerte ich mich an ein Gespräch
*&lt;Werbung&gt;* am letzten [Chaostreff Nürnberg](https://chaostreff-nuernberg.de) *&lt;/Werbung&gt;*.

Mittels [GPS Simplification](http://www.gpsvisualizer.com/tutorials/track_filters.html)  sollte es eigentlich möglich sein
die Rechenzeit zu reduzieren. Wie viele Trackpoints brauche ich wohl um
das Bild halbwegs repräsentativ darstellen zu können? TL;DR: 3%.

Um alle Trackpoints (`<trkpt>`) aller Jahre zu rendern brauche ich **38 Min**.
Wenn ich nun `gpsbabel` benutze um ein GPX File zu verkleinern,

```
gpsbabel -rt -i gpx -f 20180505-123412.gpx -x simplify,count=2000 -o gpx -F smoothed.gpx
```

Spannend dabei ist, dass man nicht einfach Werte in % bei `count=` angeben
kann, sondern man eine **absolute** Zahl an Trackpoints angeben muss. Um meine
Tracks angemessen reduzieren zu können muss ich die Gesamtzahl kennen und dann
einen Prozentwert ausrechnen.

```
$ grep -c "trkpt" 20180505-123412.gpx
> 4000
```

Da ich aber mehr als 300 Aktivitäten von Strava heruntergeladen habe wollte
ich das automatisieren. Zeit für etwas `fish` Shell.

```
for f in *-Ride.gpx
  set -l points (math (grep -c "trkpt" $f) / 30)
  gpsbabel -rt -i gpx -f $f -x simplify,count={$points} -o gpx -F SIMPLIFIED-$f
end
```

Natürlich hab ich etwas herumprobiert mit verschiedenen Verhältnissen, aber
bei der Auflösung in denen ich meine Bilder generiere reichen mir gerade
einmal **~3%** der Trackpoints um nahezu das gleiche Bild zu erzeugen:

{{< figure src="/uploads/2018/05/vergleich.png" >}}

... und die Dauer des Erzeugens sank von 38 auf 3 Minuten.
