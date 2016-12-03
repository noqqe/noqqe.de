---
title: "Über gefixte Syntax und Stunden mit RegEx"
date: 2012-11-22T22:27:00+02:00
comments: true
tags:
- Blog
- Web
- Development
- octopress
- osbn
- bash
- regex
- markdown
- html
---

Manchmal sitze ich abends einfach so herum und repariere kaputte/unschöne
Syntax alter (via [exitwp](https://github.com/thomasf/exitwp)
[importierter](/blog/2012/03/05/switched-to-octopress/)) Blogposts unter
Zuhilfenahme regulärer Ausdrücke damit besagte Eintraege nacher genauso
aussehen wie vorher. Das hilft keinem. Nur mir, wenn ich mir die Markdown
Sourcen des Octopress Repos anschaue.

{{< figure src="/uploads/2012/11/bale.jpg" >}}

## Umbrüche

Einen Großteil der Zeit habe ich mir überlegt wie ich es schaffe nach 80
Zeichen einen automatischen Umbruch mit `sed` zu machen ohne dabei ein Wort
zu zerhacken. Lösung:

``` bash
sed -i -e 's#\(.\{80\}\S*\s\)#\1\n#g' *
```

## Ugly syntax is ugly

Folgendes ist eine sehr schlimme Variante von Bildeinbettung in Markdown
(imho). Exitwp ist da aber anscheinend anderer Meinung und hat damals alle
Bilder so eingebunden:

```
Original:
[![CC by Drunken Monkey](/uploads/2010/10/243104896_eb10db6e1d.jpg)](/uploads/2010/10/243104896_eb10db6e1d.jpg)

Wünschenswert:
&#123;&#37; img center /uploads/2010/10/243104896_eb10db6e1d.jpg &#37;&#125;
```

Aber auch das hab ich dann noch hinbekommen.

```
sed -i -e 's#.*\[![^\(]*(\([^\)]*\).*#&#123;&#37; img center \1 &#37;&#125;#g' *
```

## Misc

Zusätzlich noch die Domain zwetschge.org komplett aus den Posts gekratzt,
Alte removed und kaputte Syntax aus exitwp generierten Posts.

```
$ git commit -a -m "fixed every broken markdown img tag with epic sed line"
[master b0b1693] fixed every broken markdown img tag with epic sed line
 78 files changed, 137 insertions(+), 137 deletions(-)
```

Außerdem hab ich die Hälfte des Blogposts schon bei Twitter gespoilert.
Komisch das ich meinen Account da wieder reaktiviert hab.
