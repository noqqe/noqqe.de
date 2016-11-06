---
date: 2009-09-24T19:59:58+02:00
type: post
comments: true
title: Shell | Vollständiger File-Path
aliases:
- /blog/2009/09/24/shell-vollstandiger-file-path
- /archives/663
categories:
- Development
- Linux
tags:
- file
- path
- filesystem
---

Sonst vergess ich es sowieso wieder:

```
find . -exec ls -d {} ;
```

gibt den vollständigen Datei-Pfad aus.  Sollte ausser mir nochjemand mal
seinen shoutcast trans mit einer Playlist befüllen müssen und aus Gründen
der total Pfusch-Config soetwas brauchen.
