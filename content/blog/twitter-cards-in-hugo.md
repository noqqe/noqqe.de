---
comments:
- author: Raim
  content: "Und als n\xE4chstes kommen dann bitte eingebettete Tweets statt Screenshots?
    ;-)"
  date: '2016-01-28T12:14:44.168614'
- author: noqqe
  content: lol. hat was.
  date: '2016-01-30T16:39:38.122553'
date: '2016-01-27T19:12:28'
tags:
- blog
- web
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

<!--more-->

``` html
<!-- Twitter Cards -->
<meta name="twitter:card" content="summary">
<meta name="twitter:site" content="{{ .Site.Params.Twitter }}">
<meta name="twitter:creator" content="{{ .Site.Params.Twitter }}">
<meta name="twitter:title" content="{{ $isHomePage := eq .Title .Site.Title }}{{ .Title }}{{ if eq $isHomePage false }} - {{ .Site.Title }}{{ end }}">
<meta name="twitter:url" content="{{ .Permalink }}">
<meta name="twitter:image" content="{{ .Site.BaseURL }}/images/noqqe_square.png">
<meta name="twitter:description" content="{{if .IsPage}}{{ .Summary }}{{else}}{{.Site.Params.Description}}{{end}}">
```

Die sonstigen Inhalte werden aus der `config.yaml` geholt.

``` yaml
params:
 Description: "No advertising, no support, no bug fixes, payment in advance."
 Author: "Florian Baumann"
 Twitter: "@noqqe"
```
