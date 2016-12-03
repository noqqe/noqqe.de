---
aliases:
- /blog/2010/10/19/bash-zre-und-random
- /archives/1321
date: '2010-10-19T18:35:18'
tags:
- development
- git
- zufall
- random
- humans
- zre
- shell
- zombies
- bash
title: Bash | ZRE und $RANDOM
---

Zugegeben, dieser ganze [Zombie-Sh!t](/archives/1314) unterhält mich
länger, als es mir lieb ist. Insbesondere die aktuell stattfindende
Blockschulwoche, trägt zur außergewöhnlich häufigen Nutzung bei.

In den meisten Durchläufen des ZRE beendet es ganz von alleine die
Simulation, weil entweder keine Humans oder Zombies mehr am Leben sind. Auf
Letzteres trifft das natürlich immer irgendwie zu. Aber das ignoriere ich
jetzt einfach mal. Wenn eine der beiden Seiten als Sieger hervor geht eben.

{{< figure src="/uploads/2010/10/243104896_eb10db6e1d.jpg" >}}

In wenigen Fällen® kommt es allerdings vor, das beide Seiten gleichmäßig
stark an Mitgliedern gewinnen. Und zwar mehr als 150.000 ca. Die Chance
diesen Case zu treffen würde ich jetzt auf 1:50 schätzen. Habe ihn heute 3
mal getroffen. Wie oft habe ich wohl ca. gespielt? Traurig.

Wenn dieser Case eintritt, nimmt die Simulation einen etwas schrägen
Verlauf.

  * Wachsende Bevölkerung: Die Anzahl der Mitglieder einer Seite wird mit

``` bash
local size=$(($RANDOM % $humans + 1))
humans=$(($humans + $size))
```

berechnet.

  * Maximale Opfer eines Angriffs: Die Anzahl der $victims ermittelt ZRE folgendermaßen:

``` bash
defenders=$(($RANDOM % $zombies +1))
victims=$(($RANDOM % $defenders + 1))
zombies=$(($zombies - $victims))
```

Die Problemstellung beläuft sich eigentlich auf $RANDOM. $RANDOM erzeugt
einen "zufälligen" Integer wert von 0 bis 32767.

Die Erzeugung der Unterstützung addiert sich immer wieder sofort auf den
Wert der jeweiligen Rasse. Die Anzahl der Opfer allerdings errechnet sich
aus einem Teil der Verteidiger, welcher wiederum nur ein Teil der gesamten
Rasse darstellt.

Klartext: Die Wahrscheinlichkeit, dass (max.) 32767 Mitglieder Opfer werden
ist wesentlich geringer als die Wahrscheinlichkeit 32767 Mitglieder Support
zu bekommen. Das führt zu unverhältnismäßig großem Wachstum im Environment.

Nach einiger Überlegung habe ich beschlossen die Werte der Kämpfe nicht
anzupassen und stattdessen etwas anderes Einzuführen.

**Naturkatastrophen**. Diese sollen automatisch bei Überschreitung eines
Limits (300.000 Einwohner) oder einfach nur sporadisch etwas aufräumen.
Vulkan, Hurricane oder Erdbeben. Alles dabei. Wer genau, wieviele und warum
bei diesem Event ums Leben kommen entscheidet wie immer $RANDOM.