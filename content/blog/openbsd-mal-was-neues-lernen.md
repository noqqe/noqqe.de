---
aliases:
- /blog/2012/11/30/openbsd-mal-was-neues-lernen/
comments:
- author: Flo
  content: "<p>Finde ich sehr gut, dass du dich da einarbeitest - w\xE4re nett, wenn
    du uns da fortlaufend dran teilhaben k\xF6nntest, interessiert mich n\xE4mlich
    auch brennend :)</p><p>&gt;Das fehlende free</p><p>Gibts da eigentlich einen guten
    Grund, dass das fehlt? Nervte mich bisher bei so quasi allen BSD-Versionen.</p><p>+
    was f\xFCr Vorteile hast du jetzt gegen\xFCber Debian? Klar, soll sicherer sein,
    ich teste auch immer gern etwas neues, aber ist da noch mehr?</p>"
  date: '2012-11-30T20:31:01'
- author: noqqe
  content: "<p>Hi Flo :) </p><p>Wenn es was Interessantes gibt bzw ich mal ein Res\xFCmee
    ziehe nach ein paar Wochen - klar gerne :)</p><p>Wegen free hab ich auch noch
    etwas herumgegoogelt, aber es sieht so aus als schreiben die meisten iwelche kleinen
    Wrapper Skripte oder \xE4hnliches. Warum genau weiss ich auch nicht. </p><p>Ich
    bin mir nicht ganz sicher ob es jetzt wirklich einen objektiven Vorteil gibt.
    Im Moment gef\xE4llt es mir einfach ein System zu haben das sich so \"monolithisch\"
    anf\xFChlt. In base/ ist bisher fast alles was ich brauche schon drin. Alles ist
    Security technisch schon per default hart eingerichtet (automatisch Chroots f\xFCr
    nginx zb. ) Und generell tickt das OS einfach anders als ich es gewohnt bin und
    das macht Spass :)</p>"
  date: '2012-12-01T09:57:18'
- author: zralb
  content: <p>&gt; Die (sich echt sau komisch verhaltende) Version von Vi</p><p>So
    fuehlt sich nun mal ein echtes vi an ;)</p><p>vi ist einfach nicht vim, auch wenn
    es bei vielen Linux-Distributionen so aussehen mag.<br>Das erkennt man u.A. auch
    an den arg unterschiedlichen Groessen der Binaries.</p>
  date: '2012-12-03T10:37:13'
- author: zralb
  content: <p>Das monolithische war (und ist auch immernoch) der Grund, warum ich
    das auf meinen Servern benutze.</p><p>Auf dem Laptop benutze ich dennoch Debian,
    denn unter uns gesagt, taugt OpenBSD zum browsen und Flash-Videos abspielen einfach
    nicht (nicht nur, weil es gar keinen anstaendigen Flash-Player gibt ;-)). Es fuehlt
    sich auf dem Laptop einfach super langsam an, aber die Entwicklugnsschwerpunkte
    liegen eben auch wo anders.</p><p>The right tool for the right job. Vieles ist
    anders, aber auch das hat meistens seine (guten) Gruende.</p>
  date: '2012-12-03T10:45:11'
- author: noqqe
  content: "<p>Das hab ich auch gedacht und dann mal \xFCber den Ports Tree vim nachinstallisiert.
    Aber irgendwie...trotzdem anders :) Naja. Wird schon werden!</p>"
  date: '2012-12-03T19:17:59'
- author: noqqe
  content: "<p>:) Jepp. Wie gesagt - mir gef\xE4llts. Aufm Desktop kann ich das noch
    \xFCberhaupt nicht einsch\xE4tzen wie sich das wohl anf\xFChlt. Hab noch kein
    X oder Desktop Environment gestartet. </p><p>Der letzte Teil triffts eigentlich
    auf den Kopf :)</p>"
  date: '2012-12-03T19:19:32'
- author: Martin
  content: "<p>Das ist unter FreeBSD auch so. Die Entscheidung ist allerdings nachvollziehbar
    (ebenso weshalb tcsh default ist (ist das bei OpenBSD \xFCberhaupt tcsh?)): kleine
    base Umgebung, wo ein kaputter Port das System (als root) nicht unbedienbar macht.</p><p>Vim
    und bash zus\xE4tzlich noch in base zu betreuen ist sicherlich auch ein gro\xDFer
    Aufwand gegen\xFCber einem eher geringen Nutzen (da normale User durchaus Vim
    und bash nutzen k\xF6nnen).</p>"
  date: '2012-12-03T21:44:33'
date: '2012-11-30T10:00:00'
tags:
- theo
- openbsd
- noris
- linux
- bsd
title: OpenBSD - Mal was neues lernen
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