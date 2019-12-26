---
title: "OpenBSD Desktop auf #36c3"
date: 2019-12-26T14:55:17+01:00
tags:
- OpenBSD
- cwm
- apm
---

In letzter Zeit spielte ich etwas mit OpenBSD auf dem Desktop herum, was ich
schon jahrelang mal machen wollte, aber irgendwie nie tat.

Mir fiel auf dem Dachboden der [K4CG](https://k4cg.org) ein IBM Thinkpad R32
in die Haende, welches ich dafuer benutzte. Die Hardware hatte so ihre
Probleme. Wackelkontakt auf der Tastatur, Akkulaufzeit ~20min, integriertes
Wifi konnte nur WEP, CMOS Batterie leer, Pentium 4. Allesamt loesbare
Probleme, nur leider nicht wirtschaftlich.

Was aus der Benutzung aber herausfiel, sind ein paar aktualisierte
Screenshots auf in der Wikipedia fuer
[OpenBSD](https://wikipedia.de/wiki/OpenBSD). Die Vorherigen waren von vor 15
Jahre und mit OpenBSD 3.8.

Auf lange Sicht hab ich das Teil aber dann als defekt zurueck auf den
Dachboden gelegt, bis ich es mal\* entsorge hab ich einen gebrauchtes Thinkpad
x240 geschossen.

Fuer den #36c3 habe ich mir vorgenommen etwas mit OpenBSD auf dem Desktop zu
arbeiten.

\*Nie.

## CWM

Als Erstes hab ich ein komplett blankes OpenBSD installiert mit nichts ausser
`cwm`. OpenBSD liefert diesen gleich mit und von daher lag es nahe den
auszuprobieren.

{{< figure src="/uploads/2019/12/cwm.png" >}}

Ich habe kein einziges mal irgendwas Suchmaschinieren
muessen. Ich habe auch kein Cheatsheet anlegen muessen. Warum? Weil die
Manpage so ultra geil ist. Immer wenn ich etwas kurz nicht mehr weiss, schaue
ich kurz in `man cwm` und weiss Bescheid. So geht Dokumentation.

Wie lange ich dabei bleibe wird sich ueber den Congres noch herausstellen.
Ich liebaeugle mindestens genausolange schon mit `herbstluftwm`. Support your
local Windowmanager!

## WIFI

OpenBSD supported jetzt schon seit ein paar Monaten einfache statische
listen  von Wifis. Freut mich das das mittlerweile so einfach ist.

```
join "k4cg-intern" wpakey "xxx"
join "WLAN-590596" wpakey "yyy"
join "WIFIonICE"
dhcp
inet6 autoconf
up
```

Das ist ehrlich gesagt ziemlich cool weil man nicht mehr Zeilen mit `nwid`
aus- und einkommentieren muss.

## Zusammenfassung

Was mache ich nun damit? Erstmal habe ich das ganze Konzept Workstation und
was das fuer mich bedeutet ein wenig ueberdacht. Mehr Unix-Artig. One tool
does one Job. Gute kleine Software benutzen.

Ich hoffe ich werde noch etwas mehr Posts ueber einzelne Details
schreiben, auch wie ich mir die Compose Keys fuer Umlaute in X11 entweder
Merke ode Umkonfiguriere :P

Gesendet von einem OpenBSD im ICE auf dem Weg zum Congress.
