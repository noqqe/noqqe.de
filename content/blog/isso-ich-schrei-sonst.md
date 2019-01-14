---
aliases:
- /blog/2013/11/10/isso-ich-schrei-sonst/
comments:
- author: noqqe
  content: First!
  date: '2013-11-10T15:58:09.970656'
- author: tux.
  content: '"Privacy!" Mhm, ja, soso. Das Flattr-Script in deiner Sidebar kommt dir
    nicht komisch vor?

    Ach, wen frag'' ich das eigentlich...'
  date: '2013-11-10T16:24:11.743617'
- author: noqqe
  content: Daran arbeite ich noch ;)
  date: '2013-11-10T16:26:11.616002'
- author: Anonymous
  content: test
  date: '2013-11-19T08:54:36.251564'
- author: DeBaer
  content: 'Wot! '
  date: '2013-12-06T14:06:00.234745'
- author: Clemens
  content: 'Apropos Privacy: <shameless-plug>https://neverpanic.de/blog/2014/03/19/downloading-google-web-fonts-for-local-hosting/</shameless-plug>'
  date: '2014-03-20T01:19:11.514685'
- author: Clemens
  content: Da sollten eigentlich noch &lt;shameless-plug&gt; &lt;/shameless-plug&gt;-Tags
    drum, aber die hat der Formatter gefressen, Preview gibts nicht und beim editieren
    hab ich (trotz angezeigtem Edit-Button) jeweils einen 403 bekommen.
  date: '2014-03-20T19:06:59.580997'
- author: "u\u01DDq"
  content: "\u0287s\u01DD\u0287"
  date: '2014-08-13T15:51:08.249447'
date: '2013-11-10T13:50:00'
tags:
- isso
- octopress
- kommentare
- comments
- blog
- disqus
title: Isso - Ich schrei sonst
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
