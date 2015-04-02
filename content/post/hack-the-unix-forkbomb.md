---
date: 2009-11-30T22:24:46+02:00
type: post
slug: hack-the-unix-forkbomb
comments: true
title: Hack | The Unix Forkbomb
aliases:
- /archives/740
categories:
- Development
- Linux
- PlanetenBlogger
tags:
- exploit
- fork
- forkbomb
- hack
- Linux
- unix
---

Eine der einfachsten Varianten ein unixoides Betriebssystem abzuschiessen ist mir
heute über den Weg gelaufen. Wurde 2002 von Jaromill verfasst und lautet wie folgt:

```
x(){ x|x& };x
```

Im Endeffekt wird die Funktion "x" definiert und darin zweimal aufgerufen. Somit
entstehen Prozesse, ich weiss garnicht wie viele ungefähr, vielleicht 1000? vielleicht
25000? Wie hoch ist wohl die Anzahl der Prozesse die ein BSD/Linux aushält? Naja
egal ich schweife ab. Genauso wie das System wenn man diesen Code-Schnippsel ausführt.

ps:Aus Gründen der Formatierung habe ich ":" aus der Orginalfassung durch "x" ersetzt.
Find ich persöhnlich schöner. Und mein code-block in Wordpress mag mich heute irgendwie
nicht.
