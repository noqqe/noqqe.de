---
aliases:
- /blog/2011/04/16/statistical-its-about-internal-functions
- /archives/1611
date: '2011-04-16T14:48:23'
tags:
- development
- web
- balkendiagramm
- balken
- script
- terminal
- blog
- shell
- statistical
- statistic
- bash
title: 'statistical | It''s about internal functions. '
---

Als ich [statistical](http://github.com/noqqe/statistical) auf GitHub
hochgepushed habe, fing ich an mir Gedanken über die Leistungsfähigkeit des
Scripts zu machen. Ich meine es verhielt sich in Anbetracht der Daten (in
meinen Augen) wunderbar. Die Key Länge wird bis zu 4 Tab-Längen mit
skaliert, genauso wie die Values, in Form der Bars. Aber wie verhält sich
es mit größeren Datenmengen?

{{< figure src="/uploads/2011/04/3913_12dd_450.jpeg" >}}

Diesbezüglich wollte ich eine kleine For-Schleife benutzen um mehrere
zufällige Werte zu generieren und in statistical  zu pipen.

```
time for x in $(seq 1 6000); do echo "$x:$RANDOM" ; done | statistical > /dev/null
```

Das Ergebnis war mit 6000 Datensätzen und guten 7 Minuten relativ
ernüchternd. Vor kurzem hat mich dann auch noch
[Vain](http://uninformativ.de) via GitHub auf die Geschwindigkeit von
statistical hingewiesen. In [seinem
Fork](https://github.com/vain/statistical), hat er alle extern spawnenden
Befehle gegen Bash interne Funktionen ausgetauscht. Siehe da:

```
real    0m7.788s
user    0m7.610s
sys     0m0.250s
```

Wahnsinn oder? Durch den Austausch von awk und grep durch interne Bash
Funktionen wird das ganze ernsthaft 23x mal schneller. Vielen Dank an Vain
an dieser Stelle!