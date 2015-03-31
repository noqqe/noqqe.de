---
type: post
title: "Thinkpad X201"
date: 2012-04-15T11:16:00+02:00
comments: true
categories:
- Web
- Hardware
- Linux
tags:
- thinkpad
- x201
- linux
- debian
- stable
---

Ich erstand bei [ralf-scharbert.de](http://ralf-scharbert.de) ein Thinkpad X201
(refurbished) für einen m.E. guten Preis.

{{< figure src="/uploads/2012/04/thinkpadx201-400px.png" >}}

Alles ist toll, Debian installiert, Funktionstasten gehen fast alle
Out-of-the-Box. Nur das wlan musste kurz nachgebaut werden, weil non-free
Firmware:

    $ apt-get install firmware-iwlwifi

Wenn ich mich jetzt noch dran gewöhnen könnte den roten Nippel zu nutzen statt
das kleine Touchpad ist alles gut. Zumindest wurde mir das Nahe gelegt ;)

{{< figure src="/uploads/2012/04/thinkpadx201-400px-2.png" >}}

Um das zu Windows 7 zu sichern (sollte ich es mal brauchen) hab ich außerdem
mal die Windows Image Sicherung ausprobiert. Konkret auf ein CIFS Share im
Netzwerk die .VHD Files überspielt und aus dem Internet den "Windows
Systemrettungsdatenträger" geholt :). Habe nicht probiert ob das wirklich
funktioniert, aber es klingt zumindest so als könne man sich darauf verlassen.

Hurra.

