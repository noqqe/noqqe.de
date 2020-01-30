---
title: "Abandoned Tuples - Teil 2"
date: 2020-01-29T17:48:21+01:00
tags:
- Fotos
- Python
- Pillow
---

Nachdem wir im [ersten Teil](/blog/2020/01/29/abandoned-tuples/) ein paar
schöne zufällige Bilder generiert haben, geht es nun daran was man mit
bestehenden Bildern so machen kann.

Die Basis ist auch dieses mal wieder die Liste mit Tupeln, die ich von
`Image.getdata()` bekomme.

```
[ (255,255,255), (0,0,0), ... ]
```

Im folgenden habe ich mir zwei Helper-Funktionen geschrieben, um ein
existierendes Bild einzulesen und eine um ein Bild mit Veränderten Werten
schreiben zu können.

```
# Bild auslesen
def get_image(inf):
    x, y = Image.open(inf).size
    pixels = list(Image.open(inf).getdata())
    return x, y, pixels

# Verändertes Bild abspeichern
def new_image(x, y, out, data):
    img = Image.new('RGB', (x, y))
    img.putdata(data)
    img.save(out)
```

## Fotos abdunkeln/aufhellen

Was passiert zum Beispiel wenn ich die Daten eines Bildes einlese, überall
`20` Punkte abziehe und das Bild abspeichere? Ich hatte davon keine Ahnung
etwas herumgespielt.

```
# alt
[ (255,255,255), (145,77,83), ... ]

# neu
[ (235,235,235), (115,47,63), ... ]
```

Was passiert ist, dass ich damit die **Helligkeit** des Fotos verändere!

Ich habe mir also eine kleine Funktion gebaut, welches über alle Tuples
iteriert und einen konfigurierbaren Wert abzieht.

```
def darken(inf, outf, dec):
    x, y, pixels = get_image(inf)

    npixels = []
    for triples in pixels:
        l = list(triples)
        l[0] = l[0] - dec
        l[1] = l[1] - dec
        l[2] = l[2] - dec

        npixels.append(tuple(l))

    new_image(x, y, outf, npixels)

darken("cat.png", "dark.png", 20)
```

Zum Verdeutlichen habe ich mir mal eine kleine süße Katze von
[unsplash](https://unsplash.com) runtergeladen.

{{< figure src="/uploads/2020/01/darkcat.png" >}}

Yeah. Am Ende ist es einfach genau das, was Programme tun wenn man am
Helligkeits Slider herumspielt?! Wenn ich die Werte addiere, wird das Bild
dementsprechend heller.

{{< figure src="/uploads/2020/01/allcats.png" >}}

## Fotos kälter/wärmer aussehen lassen

Einen weiteren gängigen Effekt den man aus der Foto Bearbeitung kennt ist
Fotos wärmer bzw kühler zu machen.

RGB besteht ja wie der Name schon sagt aus 3 Farben, Rot, Grün, Blau.

```
  R  G  B
(145,77,83)
```

Das heisst, um ein Bild kälter wirken zu lassen (und nicht Putin drauf sein
soll) kann ziehe ich bei Rot etwas ab und gebe  es bei Blau hinzu. So bekommt
ein Bild einen Blaustich und es wirkt kalt.

```
[ (230, 155, 30), ... ]
   -20,   0, +20
[ (210, 155, 50), ... ]
```

Um das zu tun nehme ich ehrlich gesagt fast die gleiche Funktion zu Hand und
ändere sie nur etwas. Zum Beispiel für Abkühlung

```
def cooling(inf, outf, dec):
    x, y, pixels = get_image(inf)

    npixels = []
    for triples in pixels:
        l = list(triples)
        l[0] = l[0] - dec
        l[2] = l[2] + dec


        npixels.append(tuple(l))

    new_image(x, y, outf, npixels)

cooling("cat2.png", "cooling.png", 20)
```

{{< figure src="/uploads/2020/01/coldcats.png" >}}


## Flatten

## ISO

## Randomisieren

... So. Jetzt ist der Schritt zum eigenen Instagram Filter eigentlich nicht
mehr weit. Nur wie finde ich jetzt coole Muster?
