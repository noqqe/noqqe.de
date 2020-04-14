---
title: "Video aus Bildern generieren"
date: 2019-02-08T13:18:41+01:00
tags:
- convert
- ffmpeg
---

Ich hatte an Weihnachten ein paar Bilder gemacht und dabei eine Funktion
meiner Kamera benutzt, welche viele Fotos hintereinander macht.

Beim Durchscrollen sah das ein bisschen aus wie ein Stop-Motion Video, was
ziemlich lustig war. Deshalb wollte ich daraus wirklich ein Video machen.

Zuerst mal alle Bilder etwas verkleinert, 6000x4000 ist da etwas arg groß
gewesen. ImageMagick <3

``` fish
for x in DSC*
  convert  $x -resize 30% -auto-orient medium-$x
end
```

Dann ... `ffmpeg` Magie.

    ffmpeg -framerate 8 -pattern_type glob -i 'medium-DSC*.JPG' -c:v libx264 -r 30 -pix_fmt yuv420p -o out.mp4

Man beachte dabei `-framerate` und `-r`. Diese beiden Schalter konfigurieren
letztlich die Geschwindigkeit des Videos und anzeigedauer jedes einzelnen
Bildes im Video.

Was dann etwas blöd war, ist das hochkant Bilder so abgeschnitten wirkten. Da
hat sich aber schonmal ein findiger Mensch auf
[SuperUser](https://stackoverflow.com/questions/30789367/ffmpeg-how-to-convert-vertical-video-with-black-sides-to-video-169-with-blur)
Gedanken gemacht.

Den Optionshaufen dann mit meinem vermischt und fertig.

``` fish
ffmpeg
  -framerate 8
  -pattern_type glob
  -i 'medium-DSC*.JPG'
  -c:v libx264
  -r 30
  -pix_fmt yuv420p -y
  -lavfi '[0:v]scale=ih*16/9:-1,boxblur=luma_radius=min(h\,w)/20:luma_power=1:chroma_radius=min(cw\,ch)/20:chroma_power=1[bg];[bg][0:v]overlay=(W-w)/2:(H-h)/2,crop=h=iw*4/6'
  out.mp4
```

Ich würde ja gerne ein Beispiel Video poste, aber da sind überall Menschen
drauf.
