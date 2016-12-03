---
date: '2015-08-09T09:00:49'
tags:
- osx
- shell
- web
- devops
- opensource
- bash
title: Photos auf Google Maps
---

Seit iPhoto mit "Photos" generalüberholt wurde ist die Kartenübersicht mit GPS
Koordinaten sehr [schlecht geworden](http://apple.stackexchange.com/questions/180284/how-to-show-the-map-in-photos).

{{< figure src="/uploads/2015/08/photos.jpg" >}}

Informationen werden beim Zoom verschluckt, die Karte ist entweder winzig klein
oder nur ein Subset aus allen Bildern.

### exiftool

Die wundervolle Software `exiftool` lässt Meta Daten aus Bildern lesen und
schreiben. Mit folgender Zeile werden für alle Bilder die GPS Längen- und
Breitengrade in Dezimalkodierung ausgegeben.

    exiftool -gpslongitude -gpslatitude -n -T /Users/noqqe/Pictures/Photos\ Library.photoslibrary/Masters/ -r

### Google Maps

Bisher hab ich nur die Koordinaten. Wie ich die in Google Maps reinfülle, hab
ich per Stackoverflow
[herausgefunden](http://stackoverflow.com/questions/3059044/google-maps-js-api-v3-simple-multiple-marker-example)

### Koordinaten konvertieren

Ich muss also

    7.42793055555556 43.7407694444444

zu

    ["Bild", 43.7370305555556, 7.42133888888889, 3598],

konvertieren. Wobei es aber zwingend erforderlich ist, dass die letzte ID
eindeutig und von 1 an fortlaufend ist. Diese zwei Zeilen bewerkstelligen den
ganzen Arbeitsablauf.

    $ exiftool -gpslongitude -gpslatitude -n -T /Users/noqqe/Pictures/Photos\ Library.photoslibrary/Masters/ -r  | uniq | awk '{print $2" "$1}' > /tmp/foo

    $ grep -v -- "- -" /tmp/foo | gsed -e 's/\s/, /' -e 's/^/["Bild", /' | gawk '{print $0 ", " FNR "],"}'']'

Danach hab ich einfach alles in das File aus Stackoverflow pastiert und fertig.
`maps.html` im Browser aufgerufen.

{{< figure src="/uploads/2015/08/map.png" >}}


Allgemein gefragt, gibt es eine self-Hosted, Webinterface Anwendung für Fotos
die Gesichtserkennung, ACLs und Kartenvisualisierung kann? Ich will sowas.