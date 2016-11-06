---
title: "Kilometer 1000"
date: 2012-10-10T20:55:00+02:00
comments: true
categories:
- osbn
- Web
- Debian
- Shell
tags:
- bash
- Google
- Runkeeper
- GPX
- GPS
---

Die Saison ist eigentlich so gut wie gelaufen und ich hab mein Ziel erreicht.

{{< figure src="/uploads/2012/10/1000km.png" >}}

Auch Runkeeper hat das anscheinend zu würdigen gewusst. Das dieser Dienst immer
mein Begleiter ist habe ich ja schon [hier](/blog/2012/07/02/paying-5-bucks-a-month-for-stupid-statistics/)
erwähnt.

Das Ende der Saison ist auch mal eine Möglichkeit ein Resümee zu ziehen. Kurz
hier ein paar Graphen gebaut und dort ein paar Dinge getan. Nicht
erwähnenswertes innerhalb meines runkeeper-statistics Projekts. Aber irgendwie fehlte was.
Eine Karte oder so. Wo war ich eigentlich überall?
Zum Glück bietet Runkeeper eine Export Funktion für GPX Dateien an.
Das Ziel ist alle Trails in ein GPX Datei zusammen zu
mergen und das dann in eine KML Datei zu converten. KML Dateien kann ich bei
Google Maps einfach per HTTP Pfad in das Suchfeld einkippen und diese werden
gleich interpretiert.

Ich hatte schon kurz gedacht
ich muss nun doch zum Geldbeutel greifen, aber Export ist auch kostenlos verfügbar.
Nur [gut versteckt](http://runkeeper.com/exportDataForm).

Heraus fällt als Zip Datei eine Zusammenfassung aller *.gpx Dateien für jede Tour. Einzeln.
Ich habe mir GPX Dateien zwar vorher noch nie angesehen, aber so schwierig kann
das wohl nicht sein :)

Erster Schritt: Header in eine Textdatei `combined.gpx` einfügen

``` xml
<?xml version="1.0" encoding="UTF-8"?>
<gpx
  version="1.1"
  creator="RunKeeper - http://www.runkeeper.com"
  xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
  xmlns="http://www.topografix.com/GPX/1/1"
  xsi:schemaLocation="http://www.topografix.com/GPX/1/1 http://www.topografix.com/GPX/1/1/gpx.xsd"
  xmlns:gpxtpx="http://www.garmin.com/xmlschemas/TrackPointExtension/v1">
```

Jetzt eine kleine Parsing Party in der ich alles von &lt;trk&gt; bis &lt;/trk&gt; heraus schneide und an die Datei
`combined.gpx` zusammenfüge.

``` bash
sed -ne '/^<trk>/,/<\/trk>/p' *.gpx >> combined.gpx
```

Als letztes den ganzen Spass noch mit `</gpx>` abschliessen. Auf
[http://gpx2kml.com](http://gpx2kml.com) kann man dann alles zu KML konverten
und Google vorschmeissen.

{{< figure src="/uploads/2012/10/maps.png" >}}

Natürlich kann man sich da über den Datenschutz jetzt Gedanken machen, aber
schlussendlich ist es doch egal, ob Google meinen GPX Daten gleich (unverknödelt mit
Accounts) oder später durch den Kauf von Runkeeper bekommt.

Und naja... so kann aus einem Klapperrad für Bahnhofsbedarf und ein bisschen Technik ein Hobby
werden.
