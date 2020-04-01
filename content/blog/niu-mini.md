---
title: "NIU Mini 40%"
date: 2020-04-01T20:14:48+02:00
tags:
- Niu Mini
- Mechanical Keyboard
---

Als meine Nachbarin mir die Tür zur Paketübergabe öffnete, bat sie mich
herein. Meine Verwunderung über diese Einladung sollte sich aber gleich
klären.
Und zwar als sie sagte: "Ihr Paket kommt aus China! Bitte nehmen Sie es
selbst, ich will es nicht anfassen". Ich hob das Paket aus einer Ecke ihres
Flurs auf, in das sie es geschoben hatte.

<!--more-->

Die Lieferung, die meine Nachbarin im Rentenalter freundlicherweise für mich
annahm, kam von
[kbdfans.com](https://kbdfans.com/products/niu-mini-40-diy-kit?_pos=4&_sid=6a0dd94ae&_ss=r).
Ich hatte ein Keyboard Kit zum Selbstbau bestellt.

{{< figure src="/uploads/2020/04/niu.png" caption="Hersteller Foto" >}}

Seit dieser Situation sind ein paar Wochen vergangen. So wurde eines meiner
*Ausgangsbeschränkungsprojekte* - tolles Wort oder? - der Aufbau meines **40% orthogonalen Keyboards**.

## Orthogo-was?

Mit Orthogonal ist die Anordnung der Tasten gemeint, also das Layout. Die 40%
wiederrum beziehen sich auf die Größe der Tastatur.

Warum man das will, haben andere bereits besser erklärt. Eine kleine Linkauswahl:

* [youtube.com/watch?v=AKGXZ1ReU54](https://www.youtube.com/watch?v=AKGXZ1ReU54)
* [Planck Homepage](https://olkb.com/planck)


2016 ging es bei mir Mechanical Keyboards [los](https://noqqe.de/blog/2016/07/01/mechanische-keyboards).
Nach dem [Code](https://codekeyboards.com), ging es weiter zu einem [Vortex Race 3](https://mechanicalkeyboards.com/shop/index.php?l=product_detail&p=3917) für die Arbeit.

40% Orthogonal ist die nächsttiefere Etage im Kaninchenbau. Way down we go.
Switches und Keys hatte ich noch, wenn auch nur Alte.

## Assembling things

Ich würde mich selbst nicht als Hardwarebastler bezeichnen, aber ein bisschen
Löten bekomme ich hin. Ein paar Eindrücke.

{{< figure src="/uploads/2020/04/IMG_5355.JPG" >}}
{{< figure src="/uploads/2020/04/IMG_8054.JPG" >}}
{{< figure src="/uploads/2020/04/IMG_2983.JPG" >}}
{{< figure src="/uploads/2020/04/IMG_6901.JPG" >}}
{{< figure src="/uploads/2020/04/IMG_0360.JPG" >}}

Alles in allem keine Stunde basteln. Viel Zeit davon habe ich damit verbracht
herauszufinden wie der "Stabilizer" eingebaut gehört. Das kleine schwarze
Plastikteil ist zur Unterstützung der Space Taste da.

Ja. Auch ich habe bereits schönere Lötstellen gesehen.

## .hex Dateien auf USB Geräte flashen

Die Herstellerseite gibt einem eine `json` Datei und eine Website mit auf den
Weg. [kbfirmware.com](https://kbfirmware.com/) heisst diese. Man lädt die
Firmware hoch und kann dann seine Tasten in die Firmware der Tastatur
einstellen, wie man möchte. Zurück kommt eine `.hex` Datei, die man mittels
[QMK Toolbox](https://qmk.fm/toolbox/) auf die Tastatur flashen kann.

Dann stieß ich über Twitter auf [config.qmk.fm](https://config.qmk.fm).
Der Konfigurator basiert auf einem anderen `json` Format, spuckt aber ebenso
`.hex` Firmwares aus.

{{< figure src="/uploads/2020/04/config.png" >}}

Beide Websites liefern funktionierende Firmwares, aber letztere ist besser
zu bedienen. Außerdem kann ich mir ein PDF mit meinem Mapping generieren.

## Gewohnheit ist ein Arschloch.

Dieser Blogpost war so ziemlich der erste Text den ich mit orthogonalem
Layout schrieb. Ich kann nur sagen, es ist die Krätze.

{{< figure src="/uploads/2020/04/IMG_5510.JPG" >}}

Einfachste Dinge wie das `Z` tippen sind eine Qual. Wir sprechen hier von um
Millimeter verschobene Knöpfe, die mir das Gefühl vermitteln als müsste ich
nach einem Schlaganfall elementare Dinge neu lernen.

Das ist aber nur der Einstieg. Ich denke ich werde grob 3-5 Tage dran
verschwendet haben ein Layout zu finden, um meine Tasten zu mappen. Wie viele
Modi möchte ich? Schaffe ich es überhaupt körperlich unter macOS einen
Screenshot (`Apfel+Shift+Mo(1)+3`) zu greifen? HABE ICH ÜBERHAUPT ALLE
TASTEN?

An blindes Passwörter tippen ist momentan noch nicht zu denken. An
Programmierung auch nicht. Ein paar Shell Befehle in das schwarze Terminal
eintippern? Vielleicht. Zumindest habe ich es geschafft diesen Eintrag mit
`hugo` zu verfassen.

Es gibt aber bereits einen Plan... sofern die Ausgangsbeschränkungen lange
genug andauern.
