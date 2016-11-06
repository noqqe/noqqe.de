---
title: "Isso - Ich schrei sonst"
date: 2013-11-10T15:50:00+02:00
comments: true
categories:
- osbn
- Blog
tags:
- isso
- disqus
- kommentare
- comments
- octopress
aliases:
- /blog/2013/11/10/isso-ich-schrei-sonst/
---

{{< figure src="/uploads/2013/11/machines.jpg" >}}

[posativ](http://blog.posativ.org) hat neben seinem in Python geschriebenen
Static Site Generator [acrylamid](http://github.com/posativ/acrylamid) seit
kurzem auch sein [eigenes Kommentarsystem](https://posativ.org/isso/) wiederbelebt. [Isso - Ich schrei sonst](https://github.com/posativ/isso).
Die Alternative zu [Disqus](http://disus.com) die auch hier im Blog bisher zum
Einsatz kam setzt auf Python auf, ist einfach einzubauen und sogar alle alten
Disqus Kommentare lassen sich importieren. Für die User bedeutet das konkret
kein Traffic mehr zu Gravatar, Disqus und keine lustige Registrierungen nötig. Privacy!

Abweichend von der Anleitung die auf Debian/Ubuntu Systeme abziehlt ist die
Installation dank guter Python Packages unter [OpenBSD](http://openbsd.org)
easy.

``` bash
$ sudo pkg_add py-pip
$ sudo pip install isso
```

Der Rest läuft dann genauso, wie unter anderen Distributionen auch. Zusätzlich
habe ich noch ein [kleines Daemon Skript](https://gist.github.com/noqqe/7397719) für OpenBSD geschrieben, damit das neue
Kommentarsystem auch schön automatisch startet.

Viel Spaß beim Ausprobieren von Isso!
