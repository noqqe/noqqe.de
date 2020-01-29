---
title: "Fotofilter selbst gemacht - Teil 1"
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

Seit 1-2 Jahren schiesse ich auch selbst das eine oder andere Foto (siehe
[/photos](/photos/)).

Aber ich  wusste über ein digitales Foto nicht viel. Die
Bilddaten werden in irgendeiner Weise reinkodiert sein, ein paar Exif
Metadaten und Komprimierung drüber. Ich fing an mir die Wikipedia Artikel zu
[PNG](https://de.wikipedia.org/wiki/PNG) und [JPG](https://de.wikipedia.org/wiki/JPEG) durchzulesen.

Mh. Um ehrlich zu sein wusste ich danach immer noch nicht so viel, über den genauen
Aufbau. Aber eine Kodierung so eines Bildes kann mit
[sRGB](https://de.wikipedia.org/wiki/SRGB-Farbraum) 8 Bit gefüllt werden.
Also `255,255,255`. Und dann wurde ich ein bisschen neugierig. Wie schwer war es wohl
ein PNG/JPG zu generieren?

Ich nahm die Python Library [Pillow](http://python-pillow.github.io/) zur Hand. Die
Library zu benutzen, nimmt einem das das eigentliche Bildformat ab und man
kann einfach die Werte die einem gefallen. Nach guter, alter
[Stackoverflow Engineering Tradition](https://stackoverflow.com/questions/12062920/how-do-i-create-an-image-in-pil-using-a-list-of-rgb-tuples)
Fand ich heraus das es eigentlich nur eine Liste mit 3er Tupeln braucht.

```
>>> # Pillow Datenstruktur
>>> [ (255,255,255), (0,0,0) ]
```

Ich hab mir das ein sehr rudimentäres Script gebaut, welches.

```
import random
from PIL import Image

def random_pixel():
    return (random.randint(0,255), random.randint(0,255), random.randint(0,255))

img = Image.new('RGB', (30, 30))

pixels = []

# generate 90 random colored pixels :)
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

Nunja da war dann schon recht viel Krepel dabei, weshalb ich das ein kleines
bisschen wiederverwendbarer gemacht hab.

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
    """
    generate random pixels for certain dimensions
    will be just colors and stuff
    """

    pixels = []
    print(x*y)
    for pixel in range(0, x*y):
        pixels.append(random_pixel())

    new_image(x, y, out, pixels)

random_image(600, 600, out)
random_image(1200, 1200, out)
```

Und dann konnte ich Bilder in beliebiger Größe generieren, ohne viel
herumbasteln zu müssen

{{< figure src="/uploads/2020/01/randcolorbig.png" >}}



