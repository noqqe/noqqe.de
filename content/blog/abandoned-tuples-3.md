---
title: "Abandoned Tuples - Teil 3"
date: 2020-02-02T17:48:21+01:00
tags:
- Fotos
- Python
- Pillow
---

Nach [Teil 1](/blog/2020/01/29/abandoned-tuples/) und [Teil 2](/blog/2020/01/29/abandoned-tuples-2/)
wollte ich noch ein paar weitere Effekte basteln die man aus der Fotografie
kennt.

## Flatten

Eine Sache die viele gängige Filter benutzten ist so eine Art
"flattening". Was beim Fotografieren unter Dynamic Range eigentlich
erstrebenswert ist, wird bei Instagram Filtern für viele Brauntöne wie Kaffee und Ähnlichem
bewusst weggemischt.

Wie kann man hier vorgehen?

Im Zweifel nimmt man sich einfach wieder jeden Wert im RGB Tuple und wenn der
Wert kleiner ist (dunkler) als zb. 50, setzt man in auf 50.

```
# alt
(48, 150, 30)
  < , > , <

# neu
(50, 150, 50)
```

So werden dunkle Stellen zu einem eher grauen Einheitsbrei und viele Details
gehen verloren, was einen durchaus schönen Effekt haben kann. Vor allem bei
Holz/Kaffee kann das oft gut aussehen.

```
def flatten(inf, outf, fn):
    x, y, pixels = get_image(inf)

    npixels = []
    for triples in pixels:
        l = list(triples)

        if l[0] < fn:
            l[0] = fn

        if l[1] < fn:
            l[1] = fn

        if l[2] < fn:
            l[2] = fn

        npixels.append(tuple(l))

    new_image(x, y, outf, npixels)

flatten("coffee.png", out, 50)
```

{{< figure src="/uploads/2020/01/coffee.png" >}}

Das Ergebnis sieht eigentlich ganz gut aus. Dafür das ich dafür eigentlich
(wie bei jedem anderen Effekt bisher) wiedermal nur ein paar Integer
inkrementiert habe.

## ISO

Ebenfalls viele Filter gibt es zum Thema "Analog". Filme haben eine feste
ISO. Oder Körnung mit anderen Worten. Wie könnte man das am Einfachsten
herbei führen.

Körnung bedeutet ja im Endeffekt ein bisschen rauschen. Also Punkte die
eigentlich sehr ähnlich bis gleich sind, unterscheiden sich auf der Aufnahme
voneinander. So wirkt Himmel nicht mehr blau sondern eben mit einem rauschen
darin.

Was ist Rauschen eigentlich? Farbverläufe entstehen durch die Anreihung
möglichst gleichartiger Pixel nebeneinander. Will man diesen Verlauf
aufbrechen kann man zum Beispiel einen `random` Wert auf jedes Pixel
dazurechnen.

```
(48, 150, 30), (49, 151, 31)
 +10,+10,+10    +1, +1, +1
```

{{< figure src="/uploads/2020/01/iso2.png" >}}

Wichtig ist, auf jeden der Ints von RGB die gleiche Random Zahl
aufzuschlagen, damit die Ursprungsfarbe möglichst gleich bleibt. Ansonsten
würde es ein "buntes Rauschen" werden. Und so aussehen:

{{< figure src="/uploads/2020/01/iso4.png" >}}

Und das sieht einfach nicht schön aus.

```
def iso(inf, outf):
    x, y, pixels = get_image(inf)

    npixels = []
    for triples in pixels:
        l = list(triples)
        r = random.randint(0,90)
        l[0] = l[0] + r
        l[1] = l[1] + r
        l[2] = l[2] + r

        npixels.append(tuple(l))

    new_image(x, y, outf, npixels)
```

Fertig wäre ein Teil unseres Hipster Analog Filters...

## Kontrast

Kontraste sind nicht ganz so einfach zu berechnen, da ich zu einem bestimmten
Pixel wissen muss wie ich die Farbe "intensiviere". Kontrast erhöhen bedeutet
ja im Endeffekt das die Farbe des Pixels kräftiger wird.

Dazu habe ich mir die Formel eines schlauen [Menschen geklaut](https://www.dfstudios.co.uk/articles/programming/image-programming-algorithms/image-processing-algorithms-part-5-contrast-adjustment/).

Das heisst jeder RGB Wert eines Pixels wird mit einem Faktor
multipliziert und verändert. Die Verarbeitung läuft in 2 Schritten. Faktor
berechnen und Faktor auf normalisierten Wert anwenden.

```
# Kontrast +90
[103, 148, 205]
[76, 169, 287]

# Beispiel für Blau
((103 - 128) * 2.073) + 128 = 76
```

In diesem Fall wird das Blau einfach etwas "leuchtender"

{{< figure src="/uploads/2020/01/contrast.png" >}}

Das habe ich nun in meine Funktion eingebaut

```
def contrast(inf, outf, contrast=90):
    x, y, pixels = get_image(inf)

    factor = (259 * (contrast + 255)) / (255 * (259 - contrast))

    npixels = []
    for triples in pixels:
        l = list(triples)

        l[0] = int(factor * (l[0] - 128) + 128)
        l[1] = int(factor * (l[1] - 128) + 128)
        l[2] = int(factor * (l[2] - 128) + 128)

        npixels.append(tuple(l))

    new_image(x, y, outf, npixels)

contrast("treppen.png", "inccontrast.png", 60)
contrast("treppen.png", "redcontrast.png", -60)
```

Angewendet sieht das ganze dann so aus:

{{< figure src="/uploads/2020/01/contrastcomp.png" >}}

## Fazit

Das schöne ist, das ich wirklich immer nur jedes einzelne Pixel betrachtet
habe und Modifikationen an eben diesem mache. Ich habe nie berücksichtigt wo
es liegt im Bild, welche Nachbarn es unter Umständen hat. Welche Farbe es
ungefähr sein könnte oder Ähnliches. Das alles ist völlig egal wenn man den
RGB Farbwert hernimmt und diesen nach einem bestimmten Muster verändert.

Nachdem ich jetzt ungefähr alles an Grundwerkzeug habe, kann ich anfangen an
meinem eigenen Instagram Filter zu bauen. Im nächsten Teil. Hatte ich
eigentlich schon erwähnt das ich Instagram garnicht nutze? :P
