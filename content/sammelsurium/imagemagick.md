---
title: ImageMagick
date: 2016-05-07T22:57:15.019000
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
