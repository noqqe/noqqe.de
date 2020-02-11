---
title: "Abandoned Tuples - Teil 4"
date: 2020-02-11T12:48:21+01:00
tags:
- Fotos
- Python
- Pillow
---

Nachdem ich jetzt zufällige Bilder generiert, Helligkeiten, ISO, Kontrast und
Temperatur Filter erstellt und vor allem verstanden habe, ist es Zeit etwas
Tooling herum zu bauen.

# Tooling

Dazu habe ich [argparse](https://docs.python.org/3/library/argparse.html)
verwendet. Via Argparse kann ich nun die einzelnen Filter in Reihe schalten
und per Flag die Werte aus den Commandline Parametern abrufen.

```
./abandoned-tuples \
  --contrast -20 \
  --iso 30 \
  --brightness -40 \
  --flatten 50 \
  --temperature +20 \
  cat.png test123.png

./abandoned-tuples -c 30 -i 30 -f 50 -b -30 -t +20 verkehrsmuseum_47964855892_o.jpg  out3.png
```

Das Tool ist auf [Github unter Abandoned-Tuples](https://github.com/noqqe/abandoned-tuples)

{{< figure src="/uploads/2020/02/out3.png" >}}

Dann habe ich noch etwas Meta Informationen eingebaut. Speicherverbauch und Dauer der
Laufzeit.

```
./abandoned-tuples -c 30 kreta_48223742362_o.jpg  out2.png
Time:  14.00934251
Memory: 1.0 GG
```

1GB ist schon ganz schön viel Memory, aber klar ich muss ja auch das ganze
Bild einlesen. Aber ich war bisher auch nie wirklich auf Effizienz getrimmt
was die Funktionen angeht. Anstatt eine neue Liste aufzubauen `npixels`
(siehe Vorherige Code Beispiele) habe ich nun die aktuelle Liste bearbeitet
und mittels `enumerate` die originale Liste aktualisiert.

```
for i, pixel in enumerate(pixels):
    l = list(pixel)
    r = random.randint(0, dec)
    l[0] = l[0] + r
    l[1] = l[1] + r
    l[2] = l[2] + r
    pixels[i] = tuple(l)
```

Und siehe da, Speicherverbauch halbiert. Wirklich schneller (zeitlich
gesehen) habe ich es nicht bekommen. Zwischenzeitlich habe ich versucht
anstatt der Tuple->List Conversion direkt das Tuple zu bearbeiten (bzw zu
konkatinieren) aber das war noch langsamer als die List Conversion.

```
./abandoned-tuples -c 30 kreta_48223742362_o.jpg  out2.png
Time:  13.686766427
Memory: 0.56 GB
```

Mit dem Commandline Tool das ich jetzt mit `abandoned-tuples` habe, war ich
sehr happy. Ich konnte nun mit immer den gleichen Werten (meine Art von
Filtern) meinen Fotos einen bestimmen Look verleihen. Yay.

## Filter finden.

Nachdem ich jetzt ein programmierbares Interface habe, kann ich einfach
Bilder in allen möglichen Werten generieren lassen. Quasi explorativ neue
Instagram Filter finden. Hier mal 200 Varianten
eines Bildes welches ich als Argument übergeben kann.

```
for i in range(0,200):
    # read file
    x, y, pixels = get_image(sys.argv[1])

    # apply filters
    pixels = contrast(pixels, random.randrange(-128, 128))
    pixels = iso(pixels, random.randrange(0, 64))
    pixels = flatten(pixels, random.randrange(0, 64))
    pixels = brightness(pixels, random.randrange(-128, 128))
    pixels = temperature(pixels, random.randrange(-128, 128))

    # write file with all changes to disk
    new_image(x, y, "random/x-{}{}.png".format(sys.argv[1],i), pixels)
```

Das hat wegen den extremen Werten fast immer ein bisschen was von Andy Warhol

{{< figure src="/uploads/2020/02/random2.png" >}}
{{< figure src="/uploads/2020/02/random3.png" >}}

Ein paar glückliche Zufälle (wörtlich) wollte ich aber trotzdem mal teilen.

{{< figure src="/uploads/2020/02/arthund.png" >}}
{{< figure src="/uploads/2020/02/artkadse.png" >}}
{{< figure src="/uploads/2020/02/artmountain.png" >}}

Aber alles in allem zu wenig zu gebrauchen. So ist das jetzt
Datenmüll den man auch gleich wieder löschen kann.

Mit etwas Erfahrung weiss man ungefähr was einem Bild jetzt gut tut und was
davon nicht. Und so hab ich ein paar wenige Bilder mit meiner eigenen
Software bearbeitet. Ein Beispiel hier aus Kreta:

{{< figure src="/uploads/2020/02/out2.png" >}}

## Ende

Und somit habe ich nun erreicht was ich wollte. Mein eigenes kleines Tool zur
Bildbearbeitung geschrieben. In der Analog Fotografie gibt es oft das
Argument, das es gut ist, den kompletten Prozess unter der eigenen
Kontrolle zu haben. Wirklich zu wissen was passiert und jede Kleinigkeit
selbst steuern können, von Aufnahme bis zum Bild an der Wand.

Ähnlich ist das hier auch, es fühlt sich gut an Ergebnisse solcher
Bearbeitung selbst zu erzielen. Zu wissen was abgeht und warum das Bild so
ist wie es ist.

Das Projekt Bildbearbeitung selbst machen, hört jetzt auf und ich habe dabei
sehr viel gelernt und viele Aha Momente gehabt. Kann ich nur jedem empfehlen.

Wo ich aufhöre, fängt natürlich auch professionelle Foto Software an.
Farbkurvenbearbeitung und alles weitere, wie auch schon in Teil
3 hervorgehoben.

Happy Pixel Editing
