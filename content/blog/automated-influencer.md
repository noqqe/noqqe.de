---
title: "Automated Influencer mit Python Pillow"
date: 2018-12-05T11:21:04+01:00
tags:
- instagram
---

Ich hatte überlegt [entbehrlich.es](https://entbehrlich.es) nach Instagram zu
bringen. Instagram also. Bilder bearbeiten.

Was ich wollte war folgendes: Ein Script nimmt einen Artikel und schmeisst
Title und Text auf ein "hübsches" Foto was man dann auf Instagram posten kann.

## Unsplash

Ich bin ja ein bisschen sehr Fan von [Unsplash](https://unsplash.com). Jetzt gibt es eine
[Library](https://github.com/yakupadakli/python-unsplash) die ich zum Glück
einfach benutzen konnte. Lass ich mir von Unsplash einfach ein zufälliges Bild
gegeben (anhand von ein paar keywords) und habe schonmal den ersten Schritt
zum baldigen Social Media Erfolg von Sinnlosem Wissen. Ein Foto.

```python
def search_image(query, title):

    # login
    log('Picking random image with %s as query' % query)
    auth = Auth(client_id, client_secret, redirect_uri, code=code)
    api = Api(auth)

    # search
    photo = api.photo.random(query=query, count=1)[0]
    url = photo.urls.raw

    # download
    fname = title + ".jpg"
    log('Downloading the image to %s' % fname)
    urllib.request.urlretrieve(url, fname)
    return fname

keywords = [ 'library', 'knowledge', 'learning', 'read', 'books' ]
fname = search_image(random.choice(keywords), title)
```

{{< figure src="/uploads/2018/12/unsp.png" >}}

## Text auf Bild

Der schöne Teil (und auch der Grund warum ich diesen Blogpost eigentlich
schreibe) ist nun Text auf das Bild einzufügen.

Ich hatte wegen entbehrlich.es schonmal mit
[Pillow](https://python-pillow.org/) gearbeitet (aber nur um Bilder auf ein
bestimmtes Maß zuzuschneiden), aber die Library scheint eigentlich alles zu
können was ich brauchte.

Zuerst mal den Title, eine einfach Box in der der Artikel Titel steht auf
bestimmten Koordinaten im Bild.

```
def draw_title(image, title):

    # get dimensions of text element
    text_size = title_font.getsize(title)

    # calc and init new image and generate button
    button_size = (text_size[0]+20, text_size[1]+20)
    button_img = Image.new('RGBA', button_size, "gray")
    button_draw = ImageDraw.Draw(button_img)

    # place text on button image
    button_draw.text((10, 10), title, font=title_font)

    # merge button image on unsplash image
    image.paste(button_img, (100, 100))

    return image
```

{{< figure src="/uploads/2018/12/mallon.png" >}}


## Multiline Text auf Bild

Etwas schwieriger wurde es dann schon bei dem Beschreibungstext. Ich muss die
Länge des Textes in Pixel ausrechnen den ich in die Box werfen will,
gleichzeitig muss ich wissen wie gross die Box sein soll um den Text dann für
die Box passend mit Linebreaks zu versehen.

```
def draw_text(image, text):
    log('Drawing description text')

    # Draw text
    w, _ = image.size
    lines = text_wrap(text, text_font, w - 250)
    line_height = text_font.getsize('hg')[1]
    line_length = len(max(lines, key=len))
    log('Placing text box with %s width and %s height' % (str(w-130), str(line_height + 20)))

    button_size = (w - 230, len(lines) * line_height + 20)
    button_img = Image.new('RGBA', button_size, "gray")
    button_draw = ImageDraw.Draw(button_img)
    x = 10
    y = 10
    for line in lines:
        button_draw.text((x, y), line, fill="white", font=text_font)
        y = y + line_height

    image.paste(button_img, (100, 200))

    return image
```

Da sind noch allerhand statische Werte die mir eigentlich nicht gefallen, aber
es ist "gut genug" für meinen Anwendungsfall.

{{< figure src="/uploads/2018/12/Mary-Mallon.jpg" >}}

## Fazit

Ich muss sagen dass ich _keine_ Ahnung hatte bzw. habe was Bildbearbeitung und
solche Dinge angeht. Während des Bastelns war ich aber immer wieder überrascht
wie low-levelig das eigentlich alle ist. Etwas mehr Abstraktion wäre für mich
angenehmer gewesen. Zumindest den Teil mit dem Ausrechnen. Dafür das ich
bisher immer nur CLI Zeug gefrickelt hab bin ich mit dem Ergebnis aber ganz
Happy.

Wen der Code (auch der Wrapper um die eigentlichen Artikel usw) interessiert,
kann sich das auf GitHub anschauen:
[https://gist.github.com/noqqe/5e6ce0e196e8a8722e2c5c8f3855bc3c](https://gist.github.com/noqqe/5e6ce0e196e8a8722e2c5c8f3855bc3c)

Was ich nun mit den ganzen Bildern mache ist tatsächlich etwas fraglich, da
man wohl auf Instagram keine Links posten kann? Da muss ich noch etwas
recherchieren.

{{< figure src="/uploads/2018/12/Blood-Quran.jpg" >}}
