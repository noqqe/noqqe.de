---
comments:
- author: tux.
  content: "Man k\xF6nnte Weechat auch einfach in einer Emacs-Shell laufen lassen
    und dann per M-w ...\n\nNicht?"
  date: '2015-04-20T22:33:35.941367'
- author: noqqe
  content: what?
  date: '2015-04-21T08:54:56.066003'
- author: flumblum
  content: "Oder man verwendet einfach Alt-l und kopiert sich ganz entspannt den gew\xFCnschten
    Text.\nSiehe hier unter 3.10. https://weechat.org/files/doc/weechat_faq.en.html"
  date: '2015-04-21T10:50:37.792900'
- author: noqqe
  content: "Unter ALT-L passiert bei mir ein @, wegen Macbook. Aber das k\xF6nnte
    ich zumindest mappen."
  date: '2015-04-21T14:58:23.492500'
- author: flumblum
  content: '"/window bare" ist der Befehl den du mappen willst.'
  date: '2015-04-21T17:13:10.265263'
- author: noqqe
  content: 'Unter Mac: alt-cmd und markieren.'
  date: '2015-04-22T08:49:49.091887'
date: '2015-04-20T19:57:54'
tags:
- stats
- paste
- copypasta
- alias
- opensource
- irc
- weechat
title: Weechat Copy & Paste
---

Etwas aus `vim` mit Tagbar oder `mutt` mit Sidebarpatch kopieren ist halt
immer irgendwie kacke. Bei `weechat` gibts das gleiche Problem mit dem
[buffers](https://weechat.org/scripts/source/buffers.pl.html/) Plugin und der `nicklist`.

{{< figure src="/uploads/2015/04/irc.png" >}}

Kreuz und quer Pipes und nicks drin. Nervt. Deshalb hier zwei kleine
Aliases, die ich mir gebastelt hab.

```
/alias hidebars /bar hide nicklist ; /bar hide buffers
/alias showbars /bar show nicklist ; /bar show buffers
```

So ist's für mich am Einfachsten irgendwelches Zeugs aus IRC zu kopieren.

* `/hidebars`
* Copy
* `/showbars`