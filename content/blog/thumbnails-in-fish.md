---
title: "Thumbnails mit fish"
date: 2019-11-27T18:50:53+01:00
tags:
- fish
---

Ich habe mal wieder meine Fotos umgezogen und von Flickr weg transferiert.
Und zwar hier in den Blog direkt unter [photos](/photos)

Dazu muss ich aber ein paar Sachen selbst tun, zum Beispiel Thumbnails
generieren. Dazu hab ich in meine `fish` Shell ein kleines Snippet gebaut was
ich einfach im jeweiligen Albumordner ausführen kann.

```
# Create a thumbnail image for all pictures in current folder
function create_thumbs ()

  # fix file extension names
  for jpeg in *.jpeg
    mv $jpeg (echo $jpeg |sed 's/\.jpeg$/.jpg/')
  end

  # generate thumbnail
  for x in *.jpg
    set -l tn_name (echo $x | sed 's/\.jpg/-thumb.jpg/')
    convert $x -define jpeg:size=300x300 -thumbnail '300x300>' $tn_name
  end
end
```

Ich mag ehrlich gesagt die `fish` Syntax mittlerweile. Hat auch nur so 1-2
Jahre gedauert.
