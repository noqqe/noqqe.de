---
title: "Abandoned Tuples - Teil 3"
date: 2020-01-29T17:48:21+01:00
tags:
- Fotos
- Python
- Pillow
---

Nach [Teil 1] und [Teil 2] will ich noch ein paar weitere Effekte basteln die man aus
der Fotografie kennt.

## Flatten

Eine Sache die viele gängige Filter benutzten ist so eine Art
"flattening". Was beim Fotografieren unter Dynamic Range eigentlich
erstrebenswert ist, wird bei Instagram Filtern für Coffeehouse und ähnlichem
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


```

{{< figure src="/uploads/2020/01/coffee.png" >}}

Das Ergebnis sieht eigentlich ganz gut aus. Dafür das ich dafür eigentlich
(wie bei jedem anderen Effekt bisher) wiedermal nur ein paar Integer
inkrementiert habe.

## ISO

Ebenfalls viele Filter gibt es zum Thema "Analog". Filme haben eine feste
ISO. Oder Körnung mit anderen Worten. Wie könnte man das am einfachsten
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

... So. Jetzt ist der Schritt zum eigenen Instagram Filter eigentlich nicht
mehr weit. Nur wie finde ich jetzt coole Muster?


## Ende

Das schöne ist, das ich wirklich immer nur jedes einzelne Pixel betrachtet
habe und Modifikationen an eben diesem mache. Ich habe nie berücksichtigt wo
es liegt im Bild, welche Nachbarn es unter umständen hat. Welche Farbe es
ungefähr sein könnte oder ähnliches. Das alles ist völlig egal wenn man den
RGB Farbwert hernimmt und diesen nach einem bestimmten Muster verändert.

Wo ich aufhöre, fängt natürlich auch professionelle Foto Software an. HSL
Kurven bearbeitung und alles das...
