---
title: "Abandoned Tuples - Teil 2"
date: 2020-01-30T17:48:21+01:00
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

<!--more-->


``` python
[ (255,255,255), (0,0,0), ... ]
```

Im Folgenden habe ich mir zwei Helper-Funktionen geschrieben, um ein
existierendes Bild einzulesen und eine um ein Bild mit veränderten Werten
schreiben zu können.

```python
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
`20` Punkte abziehe und das Bild abspeichere? Ich hatte keine Ahnung
und dann etwas herumgespielt, was man so verändern könnte:

``` python
# alt
[ (255,255,255), (145,77,83), ... ]

# neu
[ (235,235,235), (115,47,63), ... ]
```

Was passiert ist, dass ich damit die **Helligkeit** des Fotos verändere!

{{< figure src="/uploads/2020/01/darkcat.png" >}}

Ich habe mir also eine kleine Funktion gebaut, welches über alle Tuples
iteriert und einen konfigurierbaren Wert abzieht.

```python
def darken(inf, outf, dec):
    x, y, pixels = get_image(inf)

    npixels = []
    for pixel in pixels:
        l = list(pixel)
        l[0] = l[0] - dec
        l[1] = l[1] - dec
        l[2] = l[2] - dec

        npixels.append(tuple(l))

    new_image(x, y, outf, npixels)

darken("cat.png", "dark.png", 20)
```

Zum Verdeutlichen habe ich mir mal eine kleine süße Katze von
[unsplash](https://unsplash.com) runtergeladen.

Am Ende ist es einfach genau das, was Programme tun wenn man am
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

Das heisst, um ein Bild kälter wirken zu lassen (ohne das ein Bärenreitender Putin darauf sein
soll) kann ziehe ich bei Rot etwas ab und gebe  es bei Blau hinzu. So bekommt
ein Bild einen Blaustich und es wirkt kalt.

``` python
[ (230, 155, 30), ... ]
   -20,   0, +20
[ (210, 155, 50), ... ]
```

Um das zu tun nehme ich ehrlich gesagt fast die gleiche Funktion wie bei der
Helligkeit zur Hand und ändere sie etwas. Zum Beispiel für Abkühlung

```python
def cooling(inf, outf, dec):
    x, y, pixels = get_image(inf)

    npixels = []
    for pixel in pixels:
        l = list(pixel)
        l[0] = l[0] - dec
        l[2] = l[2] + dec
        npixels.append(tuple(l))

    new_image(x, y, outf, npixels)

cooling("cat2.png", "cooling.png", 20)
```

{{< figure src="/uploads/2020/01/coldcats.png" >}}

Das erste der 3 Bilder wurde eben umgekehrt (rot +20, blau -20) verändert
und wirkt daher wärmer... :)

Next up: Kontrast & Flattening
