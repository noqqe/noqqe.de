---
categories:
- opensource
- osbn
- stats
- ubuntuusers
date: 2015-04-20T21:57:54+02:00
tags:
- weechat
- irc
- alias
- copypasta
- paste
title: Weechat Copy & Paste
---

Etwas aus `vim` mit Tagbar oder `mutt` mit Sidebarpatch kopieren ist halt
immer irgendwie kacke. Bei `weechat` gibts das gleiche Problem mit dem
[buffers]() Plugin und der `nicklist`.

{{< figure src="/uploads/2015/04/irc.png" >}}

Kreuz und query Pipes und nicks drin. Nervt. Deshalb hier zwei kleine
Aliases, die ich mir gebastelt hab.

```
/alias hidebars /bar hide nicklist ; /bar hide buffers
/alias showbars /bar show nicklist ; /bar show buffers
```

So ist's für mich einfachsten, Zeug aus IRC zu kopieren.

* `/hidebars`
* Copy
* `/showbars`

