---
date: 2016-11-11T16:53:04+01:00
title: Moderne Websites
tags:
- blog
- js
- css
---

Websites sind ja seit Kurzem im Durchschnitt größer als
[Doom](https://mobiforge.com/research-analysis/the-web-is-doom).
Furchtbarer Trend. Mich störte schon länger so einiges an meinen Blog.
Deswegen gabs jetzt mal Veränderung. Irgendwie bisschen Bloated und ich hab
das CSS selbst nicht mehr verstanden.

Die Webfonts mussten weichen. Weniger HTML, weniger CSS, weniger JS mit
Highlightning-Foo.

Hier mal der Vergleich

{{< figure src="/uploads/2016/11/old.png">}}

{{< figure src="/uploads/2016/11/new.png">}}

Man kann auch ungenutzte CSS Klassen wunderbar mit [uncss](https://github.com/giakki/uncss) ausbauen.

```
uncss http://localhost:1313/blog/2016/10/31/magic-the-archivist/ themes/noqqe-v3/static/css/styles.css > newcss.css
```

Webdevelopment liegt mir nicht. Sieht man bestimmt.
Mir gefällt die neue Schlichtheit persönlich extrem gut. Außerdem gibts
jetzt mal http/2 durch `nginx`

```
server {
  listen       443 ssl http2;
  listen       [2a00:d1e0:1000:d00:216:3eff:fe07:bfad]:443 ssl http2;
  server_name  noqqe.de gtn6uc5wbeavcda3.onion;
  [...]
}
```

Auch die Kommentare via [isso](https://github.com/posativ/isso) sind damit
gestorben. Da ich aber auch weiterhin gerne Hasskommentare^WAnmerkungen
bekomme, könnt ihr mir Mails an `felix-bloginput (at) noqqe.de` schreiben.
Vermerke den Input dann gerne.
