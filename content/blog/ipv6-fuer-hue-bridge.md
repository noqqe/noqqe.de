---
date: '2016-06-26T10:44:36'
tags:
- hue
- openbsd
- network
- phillips
- administration
- hardware
- ipv6
- bsd
title: "IPv6 f\xFCr Phillips Hue Bridge"
---

In der K4CG haben wir seit Längerem ein paar Phillips Hue Lampen die wir
auch über ein [CLI](https://github.com/k4cg/k4cglicht) steuern können.
Wir wollen das auch übers Internet steuerbar haben um die Glühbirnen letztendlich mit
dem Bot [Rezeptionistin](https://github.com/k4cg/rezeptionistin) zu
bedienen.

Das Problem dabei: Wir haben nur eine permanente IP die v6 ist und die Hue
Bridge kann nur IPv4. Deshalb musste irgendwie eine Weiterleitung von einer v6
Adresse zu einer internen v6 Adresse her. Am besten nur Portforwarding.

Am Anfang hab ich's mit `pf` versucht, aber das endet mit dem Error `AF
Mismatch`. Danach wurde es eine Lösung mit einem Netzwerktool, das wohl so
ziemlich alles kann: `socat`.

``` bash
/usr/local/bin/socat TCP6-LISTEN:1337,fork TCP4:10.0.0.17:80
```

Das alles über `supervisord` gestartet, mit `pf` abgesichert und fertig ist die permanente
Portweiterleitung von v4 auf v6.
