---
date: '2014-04-20T08:20:00'
tags:
- development
- web
- imagewall
- python
- soupio
- k4cg
- wall
- parasoup
- images
title: nichtparasoup
---

In der [k4cg](http://k4cg.org) laufen auf dem Beamer meistens irgendwelche
Bilder von [soup.io](http://soup.io) durch. Bisher mittels dem (fremdgehosteten)
[soupcache](https://github.com/exi/soupcache).

Da diese in letzter Zeit oft keine Bilder auswirft und wir das Stück Software
aufgrund der Codebasis auch nicht selber hosten können/wollen, hab ich letzte
Woche etwas eigenes geschrieben. [nichtparasoup](https://github.com/k4cg/nichtparasoup)

{{< figure src="/uploads/2014/04/nichtparasoup.png" >}}

Ums kurz zu machen, ein kleines Python Skript ersetzen Nodejs, MongoDB,
Crawler-Shell-Skripte und viel zu viel JS-Code. Beim JavaScript-Part in der neuen Lösung
geht großer Dank an Jan. Mit Webkram kenn ich mich überhaupt nicht aus.

Code und Anleitung auf Github: [github.com/k4cg/nichtparasoup](https://github.com/k4cg/nichtparasoup).
Wer möchte kanns auch gerne ausprobieren und Rückmeldung geben.
