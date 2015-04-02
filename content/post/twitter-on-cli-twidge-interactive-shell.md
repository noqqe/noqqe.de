---
date: 2010-10-05T14:47:10+02:00
type: post
slug: twitter-on-cli-twidge-interactive-shell
comments: true
title: Twitter on CLI | Twidge Interactive Shell
aliases:
- /archives/1285
categories:
- Development
- Debian
- General
- Linux
- PlanetenBlogger
- Web
tags:
- alias
- cli
- cli-twitter
- customized
- follow
- interactive
- john goerzen
- shell
- twidge-sh
- twidge-shell
- twitter
- twitter-shell
---

Seit längerem benutze ich [John Goerzen](http://www.complete.org/JohnGoerzen)'s CLI Twitter Client [Twidge](http://wiki.github.com/jgoerzen/twidge/). Weniger zum Posten als zum Lesen, aber dennoch kann ich dieses Stück Software nicht mehr weg denken. Ich finde Twidge erntet allgemein zu wenig Beachtung. Wie dem auch sei, sieht eine durchschnittliche Nutzung von Twidge wie folgt aus:

```
$ twidge setup
$ twidge lsrecent
<timpritlove> Verkannt, ignoriert, nahezu vergessen - aber sinnvoll, nützlich und wertvoll: das Semikolon!
[...]
$ twidge update "Ich benutze gerade Twidge"
```


Bei intensiver Nutzung nervt die Syntax allerdings etwas. Deshalb hab ich mich hingesetzt und eine Twidge-Shell gebaut.

Die angepasste Shell "[twidge-sh](http://github.com/noqqe/twidge/blob/master/twidge-sh)" kann einfach über Aufruf ./twidge-sh gestartet werden.

```
$ wget http://github.com/noqqe/twidge/raw/master/twidge-sh
$ chmod +x twidge-sh
$ ./twidge-sh
noqqe@twitter> lsrecent
<rivva> Skype jetzt für Android verfügbar – Skype Blog auf
Deutsch http://rivva.de/QajD
[...]
noqqe@twitter> update "Ich benutze gerade Twidge-Shell"
noqqe@twitter> lsdm
noqqe@twitter> follow technicallife
```


Die Twidge-Shell hat weiterhin folgende Features:



	
  * Der Prompt wird aus der .twidgerc generiert (username@service>)



	
  * Alle Standart Kommandos  von Twidge werden automatisch komplettiert und sind benutzbar (z.B. user@service> lsrecent)



	
  * Alle in der .twidgerc definierten Aliase werden übernommen und sind ebenfalls verwendbar. (z.B. user@service> rls)



	
  * Die Twidge-Shell funktioniert weiterhin als ganz normale Shell mit allen Zusätzen und Funktionen.


Ich habe mich dabei an Ryan Tomayko's git-sh orientiert, der ähnliches für Git gebaut hat. Was übrigens auch wirklich gut funktioniert.

[Mein Twidge-Fork auf Github (mit Twidge-sh)](http://github.com/noqqe/twidge)
[Twidge von John Goerzen](http://github.com/jgoerzen/twidge/)
[Twidge zum Download für Debian](http://packages.debian.org/search?keywords=twidge)
[Twidge zum Download für Ubuntu](http://packages.ubuntu.com/de/karmic/twidge)
[Twidge-Shell Script zum Download](http://github.com/noqqe/twidge/raw/master/twidge-sh)
