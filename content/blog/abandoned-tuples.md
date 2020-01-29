---
title: "Abandoned Tuples - Teil 1"
date: 2020-01-29T16:03:01+01:00
tags:
- Fotos
- Python
- Pillow
---

Es ist ja mittlerweile völlig normal seine Fotos in irgend einer Weise zu
bearbeiten. Auf Twitter gibt's den Zauberstab, Instagram stellt dutzende
Filter zur Verfügung, das eigene mobile Betriebssystem hat allerhand
Bordmittel dabei um Fotos zu bearbeiten.

Seit 1-2 Jahren schieße (witizg, weil es eine Leica gibt die wirklich wie ein
Gewehr
[aussieht](https://petapixel.com/2015/06/09/leica-camera-rifle-prototype-valued-at-over-350000/))
ich auch selbst das eine oder andere Foto (siehe [/photos](/photos/)).

Aber ich wusste über ein digitales Foto nicht viel. Die
Bilddaten werden in irgendeiner Weise reinkodiert sein, ein paar Exif
Metadaten und Komprimierung drüber. Ich fing an mir die Wikipedia Artikel zu
[PNG](https://de.wikipedia.org/wiki/PNG) und [JPG](https://de.wikipedia.org/wiki/JPEG) durchzulesen.

Mh. Um ehrlich zu sein wusste ich danach immer noch nicht so viel über den genauen
Aufbau. Aber ein Pixel so eines Bildes kann mit [sRGB](https://de.wikipedia.org/wiki/SRGB-Farbraum) 8 Bit gefüllt werden.

Also `255,255,255`. Und dann wurde ich ein bisschen neugierig. Wie schwer war es wohl
ein PNG/JPG zu generieren? Zeit für etwas Python.

Ich nahm die Library [Pillow](http://python-pillow.github.io/) zur Hand. Eine
Bibliothek zu benutzen, nimmt einem das das eigentliche Bildformat ab. Das
Format nachbauen war zumindest für den Anfang auch nicht mein Ziel, von daher
nahm ich diese Funktionalität dankend an.

Nach guter, alter
[Stackoverflow Engineering Tradition](https://stackoverflow.com/questions/12062920/how-do-i-create-an-image-in-pil-using-a-list-of-rgb-tuples)
Fand ich heraus das es eigentlich nur eine Liste mit 3er Tupeln braucht um
mit Pillow ein Leeres Bild mit Daten zu befüllen.

```
>>> # Pillow Datenstruktur
>>> [ (255,255,255), (0,0,0) ]
```

Ich hab mir das ein sehr rudimentäres Script gebaut.

```
import random
from PIL import Image

def random_pixel():
    return (random.randint(0,255), random.randint(0,255), random.randint(0,255))

img = Image.new('RGB', (30, 30))

pixels = []

# generate 90 randomly colored pixels :)
for pixel in range(0, 30*30):
    pixels.append(random_pixel())

img.putdata(pixels)
img.save("randcolor.png")

```

Um ein Bild mit 30x30 Auflösung zu generieren brauche ich also 90 Pixel.
Und das Ergebnis fand ich schick (vergrößert für die Sichtbarkeit).

{{< figure src="/uploads/2020/01/randcolor.png" >}}

Ich spielte noch viel herum, und Pillow ist **ultra** nett zu einem was die
Pixeldaten angeht.

* Es schneidet Tupel die zu viel in der Liste sind
* Werte im negativen Bereich werden einfach geglättet
* 1 Dimensionale Liste von Tupeln werden einfach je nach Breite und Höhe ins
  Bild eingefüllt

Im Schnippsel von oben war dann schon recht viel Krepel dabei, weshalb ich das ein kleines
bisschen wiederverwendbarer gestalten wollte.

```
import sys
import random
from PIL import Image

out = sys.argv[1]

def new_image(x, y, out, data):
    img = Image.new('RGB', (x, y))
    img.putdata(data)
    img.save(out)

def random_pixel():
    return (random.randint(0,255), random.randint(0,255), random.randint(0,255))

def random_image(x, y, out):
    pixels = []
    for pixel in range(0, x*y):
        pixels.append(random_pixel())

    new_image(x, y, out, pixels)

random_image(1200, 1200, out)
```

Und dann konnte ich Bilder in beliebiger Größe generieren, ohne viel
herumbasteln zu müssen

{{< figure src="/uploads/2020/01/randcolorbig.png" >}}

Zu meiner Enttäuschung sehen die Zufallsbilder größtenteils gleich aus. Das
wirkliche Lernerlebnis war, dass die Datenstruktur der Bilder so simpel ist.

Eine völlig neue Liste mit Tupeln erzeugen ist ja ganz nett. Aber ich kann
auch die Liste eines existierenden Bildes auslesen....

```
from PIL import Image
pixels = Image.open("foo.png").getdata()
```

... to be continued
