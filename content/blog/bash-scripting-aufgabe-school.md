---
date: 2010-05-22T09:09:35+02:00
comments: true
title: Bash | Scripting-Aufgabe @ School
aliases:
- /blog/2010/05/22/bash-scripting-aufgabe-school
- /archives/1011
categories:
- Development
- Linux
tags:
- Schule
- Aufgabe
- Bash
- Rechner
---

Wie [erwähnt](/?p=1005), gab es gestern den Test in Anwendungsentwicklung.

> _Aufgabenstellung: Erstelle ein Skript "rechner", welches 2 Parameter
> aufnimmt und über Abfrage innerhalb des Skriptes den Modus auswählbar
> macht. Folgende Modi und Funktionen müssen enthalten sein:_

> _Modus 1: Multiplikation
> Usage: ./rechner 1  2
> Ausgabe: Das Ergebnis von 1 * 2 = 2_
>
> _Modus 2: Counter
> Usage: ./rechner 2 8
> Ausgabe : 2-3-4-5-6-7-8_

Formatierung muss genau der Vorgabe entsprechen. Außerdem darf der erste
Parameter nicht größer als der Zweite sein und falsche Modi-Angabe mit
Fehler gemeldet werden.

Ich muss für meinen Teil ganz ehrlich sagen, dass das wohl ein noch viel
unsinnigeres Programm ist. Wenn ich auch glaube, dass es einfach nur um die
Erfüllung bestimmter Umstände geht. Der Rest ist anscheinend ziemlich egal.

Gelöst habe ich das ganze wie folgt:

``` bash
#!/bin/bash
i=$1
if [ $1 -gt $2 ]; then echo "Falscheingabe" ;  exit 200; fi

echo "Multiplikation = 1"
echo "Einserschritte = 2"
echo -n "Modus eingeben:" ; read modus

if [ $modus -eq 1 ] ; then
        echo -n "Die Multiplikation aus $1 * $2 = "; expr $1 * $2
elif [ $modus -eq 2 ] ; then
        while [ $i -lt $2 ] ; do
        echo -n  "$i-"
        i=$[$i+1]
        done
        echo "$2"
else
        echo "Falscheingabe"
fi
```

Und wie von[@nicohofmann](http://twitter.com/nicohofmann) [prophezeit](/?p=1005): Die 1 ;)
