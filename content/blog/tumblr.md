---
title: "Tumblr, echt jetzt?"
date: 2020-07-16T15:32:53+02:00
tags:
- Tumblr
- soup.io
---

Im März 2019 habe ich begonnen Tumblr zu benutzen. Ich wollte eine
Plattform auf der ich Bilder, Textschnippsel und Gifs die ich
im Internet finde sammeln kann.

Das hatte ich früher(tm) schon auf soup.io mittels eines JavaScript
Bookmarklets. Das soup.io jetzt abgeschaltet wurde nehme ich mal zum Anlass
ein paar Gedanken hier zu lassen.

<!--more-->

Tumblr ist jetzt kein Nachfolger oder soetwas für soup.io.
Überhaupt habe ich soup.io seit Jahren nicht benutzt und der Unterschied ist
auch klar. Auf Tumblr konsumiere ich nicht, ich poste nur und sammle für mich.

## Warum Tumblr?

Tumblr eignet sich zum Sammeln von Content jeglicher Art super!

{{< figure src="/uploads/2020/07/tumup.png" >}}

Ich mag den Gedanken das ich Schnipsel und allen möglichen Kram der mir
gerade vor die Füße fällt an einem Ort leicht zugänglich zusammensammeln
kann. Über die Archive Page kann ich auch leicht nach Zeit und Tags filtern

{{< figure src="/uploads/2020/07/tumarchive.png" >}}

... und warum nicht selfhosted? Ich bin ein grosser Fan von
selfhosted Services. Während eine Webseite die Bilder anzeigt leicht zu bauen
wäre, stellt sich die Anbindung eher als Problem heraus. Wenn das Posten
kompliziert und langwierig ist, mache ich es nicht.


Zum Beispiel möchte ich von meinem iPhone mittels des "Share" Buttons Medien hochladen
\- dazu brauche ich eine native iOS App. Tumblr hat diese App.

Ich will im Browser möglichst ohne großes Herumtun Bilder posten
können (wie früher mit dem Bookmarklet). Für Firefox gibt es ein Tumblr
Plugin.

{{< figure src="/uploads/2020/07/browser.png" >}}

Außerdem gibt auch noch gute Libraries und total angenehme [API Limits](https://www.tumblr.com/docs/en/api/v2#rate-limits), die es mir
zusätzlich auf alle anderen Wege erlaubt Content zu posten.

Und dann spare ich mir auch noch die Kosten und den Aufwand das alles selbst
zu machen. Da ist Tumblr, auch wenn es sich nach 2009 anfühlt, der perfekte
Dienst.

Das alles kann ich privat nicht leisten. Und es hat eine peinlich lange Zeit
gedauert um das zu verstehen. Und nochmal länger, dass es bereits einen
Service gibt der genau das für mich tut.

## Posten von der Commandline

Es gibt dann doch immer mal wieder Anwendungsfälle von denen ich "schnell mal
was von der Commandline" posten will. Dafür gibt es die oben beschriebene
[Library](https://github.com/tumblr/pytumblr)

```python
#!/usr/bin/env python3

import os
import datetime
import sys
import glob
import pytumblr

# api keys
client = pytumblr.TumblrRestClient('x', 'x', 'x', 'x')

for x in glob.glob(sys.argv[1]):
    print("Uploading %s: " % x, end=''),

    r = client.create_photo("hoheitsakt",
      state="published",
      data=x)
    try:
        print(r["display_text"])
    except:
        print(r['meta']['status'], r["response"]["errors"][0])
```

## Soup.io Content

Und wo wir dann schon dabei sind: Ich habe meine ~400 Posts von 2011 aus der
sterben Suppe gerettet. Mithilfe des
[Scrapers](https://github.com/rixx/ripsoup) von
[@rixx](https://twitter.com/rixxtr).

Ein bisschen mit Extra-Daten angereichert (Datum ~2011 und getagged mit
'soup.io') und die Posts waren migriert...

```python3
client.create_photo("hoheitsakt",
  state="published",
  data=x,
  tags=['soup.io'],
  date=datetime.datetime(2011,4,30,3,20,6))
```

Tada.

{{< figure src="/uploads/2020/07/soupmig.png" >}}
