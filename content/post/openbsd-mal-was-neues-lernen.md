---
type: post
title: "OpenBSD - Mal was neues lernen"
date: 2012-11-30T12:00:00+02:00
comments: true
categories:
- OpenBSD
- osbn
aliases:
- /blog/2012/11/30/openbsd-mal-was-neues-lernen/
---

Schon seit Längerem liegt unter meinem Sofa noch ein 1HE Dell Server, mit dem
ich nicht so wirklich weiss was ich damit anstellen soll. Habe mich aber jetzt
entschlossen den Server ins Rechenzentrum von [noris](http://noris.de) zu stellen und
damit was Neues zu machen.

[OpenBSD](http://openbsd.org) also. Ich weiss garnicht so genau warum. Irgendwie ein Mix aus Einflüssen, Neugier und
Migrationsbedarf. Nach allem was ich gehört und gelesen habe soll die Figur
[Theo de Raadt](http://en.wikipedia.org/wiki/Theo_de_Raadt)
auch ganz...seltsam...sein. Aber das ist [Linus](http://en.wikipedia.org/wiki/Linus_Torvalds)
auch, also keine Umgewöhnung nötig.

{{< figure src="/uploads/2012/11/taskwarrior.png" >}}

So sieht es übrigens aus wenn ich mit [taskwarrior](http://taskwarrior.org) ein
Projekt abwickle. Ich liebe diese Software, aber dazu vielleicht wann anders
mal.

Dinge die mir aktuell bei OpenBSD Spaß machen:

* Login Classes
* nginx in base/
* tmux in base/
* pf
* Chroots, chroots, chroots!

Dinge an die ich mich erst gewöhnen muss:

* ksh
* Die (sich echt sau komisch verhaltende) Version von Vi
* Das fehlende `free`

Öfters habe ich mich beim hantieren mit OpenBSD dabei erwischt Änderungen
am System zu machen bei denen ich mich gefragt habe "Mache ich das jetzt weil
ich das gut finde, oder weil ich es von Debian so gewohnt bin?". Diese Frage
abzuwägen tat ich mich bei manchen Dingen extrem schwer. Ob nun die fehlende
`bash` oder `tmux` statt `screen`. Ich habe versucht ausschliesslich Software zu
verwenden die bereits im OpenBSD Base enthalten ist. Frei nach dem Motto "Die
wissen schon was sie tun". Und dabei habe ich viel neues gelernt.

Jetzt noch warten bis die Second Edition "Absolute OpenBSD" von
[Michael W. Lucas](http://blather.michaelwlucas.com) herauskommt. Laut seinem
Blog könnte es im Frühling 2013 soweit sein. Oder man liesst ein Buch von 2003, also
die First Edition.

