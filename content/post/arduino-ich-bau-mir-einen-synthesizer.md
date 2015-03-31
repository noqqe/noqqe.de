---
date: 2011-12-21T17:30:58+02:00
type: post
slug: arduino-ich-bau-mir-einen-synthesizer
status: publish
comments: true
title: Arduino | Ich bau mir einen Synthesizer
alias: /archives/1833
categories:
- Coding
- General
- Hardware
- PlanetenBlogger
- Web
tags:
- Arduino
- auduino
- community
- kabel
- lfo
- löten
- low frequenzy synthesizer
- soldering
- synthesizer
- synthi
- verkabeln
- wiring
---

Vor ca. 2 Wochen habe ich auf der [Suppe vom K4CG](http://k4cg.soup.io) ein [Video](http://vimeo.com/2266458) über einen auf Arduino basierten Synthesizer. Die "Firmware" die darauf läuft nennt sich "Auduino".

Auf deren [Projektseite](http://code.google.com/p/tinkerit/wiki/Auduino) habe ich mich dann etwas schlau gelesen und wie auch schon bei dem [Megabitmeter](/archives/1751) über tinkersoup.de meine Teile bestellt. Auf der Projektseite ist die Konstruktion des ganzen finde ich zwar nicht sonderlich gut beschrieben, aber man kommt mit ein bisschen Googeln und Reverse Engineering schon weiter.

Habe dabei aber einen [Arduino Nano](http://arduino.cc/en/Main/ArduinoBoardNano) benutzt, weil mir die Anschlüsse bzw "Architektur" besser gefällt und ich nicht erst ein Breakout Board von Seriell auf USB nachkaufen musste. Entgegen aller Erwartungen musste ich die Firmware dafür nichtmal modifizieren, da auch bei diesem Board ein ATMega328 verbaut ist.

Ich habe mir wegen der einfacheren Anbringung am Nano so eine Art Halterung/Breadboard mitbestellt, in dem ich die Adern mit Schrauben einfacher verbauen konnte.

{{< figure src="/uploads/2011/12/IMG_0604.jpg" >}}

Die Potenziometer (wieder was, das ich gelernt habe) sind in Reihe an den Ground und den 5V Pin geschlossen. Der jeweils mittlere Pin der Drehschalter kommt an die Analog Pins 0 bis 4.

{{< figure src="/uploads/2011/12/IMG_0606.jpg" >}}

Danach kam der (für mich) kniffligere Teil. Der Audio Jack (bzw. Klinke Buchse) hat von Haus aus 5 Pins. Auf der Projekte Seite von Auduino nur Input und Ground. Nach bisschen schlaulesen in Wikis und Foren scheint es, als würden die verschiedenen Revisionen von Klinke andere Features mit sich bringen. für den Mini Synthesizer hätte vollkommen Klinke Mono ausgereicht. Diverse Zusatzfunktionskanäle sind da eigentlich überflüssig aber im [Audio Jack bei TinkerSoup](http://www.tinkersoup.de/product_info.php?products_id=74&osCsid=3c2172e4114e78d30b2788b3cd0d6077) integriert.

Nach etwas Trial and Error Verfahren den weg für Doofe gewählt. Ich hab ehrlichgesagt einfach ein altes Klinke Stecker auf Buchse Kabel aufgeschnitten und mir die Belegung auf der Steckerseite angesehen.

{{< figure src="/uploads/2011/12/IMG_0609.jpg" >}}

Bei 3poligen Klinken Steckern sind die vorderen beiden Kontakte fürs Signal (Links, Rechts) und hinten für Ground. Habe dann die beiden Signaladern auf der Buchsenseite verdrillt und wie vorgesehen in den Digitalen Pin 3 geklemmt. Ground natürlich an seine Stelle.

{{< figure src="/uploads/2011/12/IMG_0611.jpg" >}}

Im Endeffekt wars dann schon fertig. Firmware mit dem Arduino IDE auf den Chip geladen und hat auch schon funktioniert.  Aber weil ich dann ständig die Potenziometer durcheinander gebracht habe, hab ich noch eine alte Plastikbox aus dem Baumarkt meiner Wahl benutzt, die entsprechenden Löcher gebohrt dort das ganze eingebaut.

{{< figure src="/uploads/2011/12/IMG_0616.jpg" >}}

Etwas smoother ;) Noch ein paar kleine Kostproben von einem wirklich unbegabten Synthesizer-Bediener. Beim hören etwas aufpassen, ab und zu ist mir da ein Ton entglitten.

[Auduino Sample Mp4](/uploads/2011/12/record.mp4)
[Auduino Sample WAV 16 bit 22 Kkhz](/uploads/2011/12/record16bit-22kkhz.wav)
