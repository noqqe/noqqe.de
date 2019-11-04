---
title: ImageMagick
date: 2016-05-07T22:57:15
tags:
- Software
- image
- imagemagick
---

Identifizieren von Bildmaßen

    identify -format "%[fx:h] %[fx:w] %d/%f\n" {}
    400 500 ./2016/01/snape.jpg
    351 500 ./2016/02/crypto.jpg
    225 373 ./2016/02/hash.jpg
    376 500 ./2016/03/iphone.jpg
    386 500 ./2016/03/piepie.jpg
    400 400 ./2016/03/versioning.jpg

Resize eines Bildes

    convert crypto.jpg -resize 50% crypto.png

GIF aus mehreren JPGs erstellen. Delay sind Ticks pro Sekunde. 70 ist ein
annehmbares Maß.

    convert -resize 768x576 -delay 60 -loop 1 in*.JPG out.gif
    convert -resize 20% -delay 60 -loop 1 in*.JPG out.gif
