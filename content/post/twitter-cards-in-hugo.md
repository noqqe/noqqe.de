---
categories:
- blog
- web
date: 2016-01-27T20:12:28+01:00
tags:
- twitter
- html
- meta
title: Twitter Cards in Hugo
---

Twitter Cards sind diese seltsamen kurzen Vorschauen die einem auf Twitter
bei geposteten Links angezeigt werden.

{{< figure src="/uploads/2016/01/twittercard.png" >}}

Bei den meisten Zeilen werden nur statische Dinge aus dem Hugo Context oder
der Konfiguration angezogen. Die einzige interessante Zeile ist wohl der
`twitter:title`. Damit nicht nur einzelne Blogposts sondern auch Seiten wie
Archiv oder Home mit dem richtigen Seitentitel bestückt werden ist die
`if`/`else` Unterscheidung nötig.
~~~ html
<!-- Twitter Cards -->
<meta name="twitter:card" content="summary">
<meta name="twitter:site" content="{{ .Site.Params.Twitter }}">
<meta name="twitter:creator" content="{{ .Site.Params.Twitter }}">
<meta name="twitter:title" content="{{ $isHomePage := eq .Title .Site.Title }}{{ .Title }}{{ if eq $isHomePage false }} - {{ .Site.Title }}{{ end }}">
<meta name="twitter:url" content="{{ .Permalink }}">
<meta name="twitter:image" content="{{ .Site.BaseURL }}/images/noqqe_square.png">
<meta name="twitter:description" content="{{if .IsPage}}{{ .Summary }}{{else}}{{.Site.Params.Description}}{{end}}">
~~~

Die sonstigen Inhalte werden aus der `config.yaml` geholt.

~~~ yaml
params:
 Description: "No advertising, no support, no bug fixes, payment in advance."
 Author: "Florian Baumann"
 Twitter: "@noqqe"
~~~
