---
comments:
- author: Anonymous
  content: "Coole Sache! Kennst du halbwegs brauchbare Tutorials f\xFCr DF? Oder hast
    du dir alles leidvoll angeeignet? ;)"
  date: '2015-02-08T23:23:31.119230'
- author: noqqe
  content: "Das erste Jahr hab ich eigentlich mit ein paar Internet Howtos und dem
    DF Wiki rumgebracht. Das lief ganz gut. Danach hab ich dann \nhttp://shop.oreilly.com/product/0636920022565.do
    gelesen. Das bringts mehr :) w\xFCrd ich empfehlen!"
  date: '2015-02-09T09:19:11.168417'
date: '2015-02-07T15:16:00'
tags:
- osx
- dwarf
- mac
- homebrew
- dwarffortress
- yosemite
title: Dwarf Fortress unter OS X Yosemite
---

[Dwarf Fortress](http://www.bay12games.com/dwarves/) unter OSX Yosemite zum
Laufen zu bekommen ist garnichtmal so einfach. Bisher hab ich immer direkt
die Version von DF von den Entwicklern heruntergeladen. Diesmal die OS X
Homebrew Version heruntergeladen. Die Probleme sind aber die gleichen.

{{< figure src="/uploads/2015/02/df.png" >}}

### Installation

Viel mehr als die Sources herunterladen und an eine bestimmte Stelle
auspacken, wird bei der Brew Version auch nicht gemacht.

```
brew tap homebrew/games
brew install dwarf-fortress
```

Da Dwarf Fortress auf viel X11 und deren Libraries aufbaut,
[Xquartz](http://xquartz.macosforge.org/landing/) installieren.

### Fehlerbehebung

```
$ ./df
dyld: Library not loaded: /usr/X11R6/lib/libfreetype.6.dylib
  Referenced from:
    /usr/local/Cellar/dwarf-fortress/0.40.23/libexec/libs/SDL_ttf.framework/Versions/A/SDL_ttf
  Reason: no suitable image found.  Did find:
    /usr/local/lib/libfreetype.6.dylib: mach-o, but wrong architecture
```

Der obige Fehler entsteht, da die default Location der Libraries nicht mit
Xquartz kommt. Daher diese umkopieren.

```
sudo mkdir /usr/X11R6
sudo cp -a /usr/X11/* /usr/X11R6/
```

Ein weiteres Problem gibt es allerdings mit Retina Displays. Die
Standardauflösung bzw. der Darstellungsmodus ist dafür nicht gemacht.

Dafür noch den `PRINT_MODE` von `2D` auf `STANDARD` umstellen.

```
$ vim /usr/local/Cellar/dwarf-fortress/0.40.23/libexec/data/init/init.txt
[...]
[PRINT_MODE:STANDARD]
[...]
```

Schon läufts.