---
title: "ICMP for Dummies"
date: 2019-02-22T16:51:09+01:00
tags:
- icmp
- ping
- OpenBSD
---

Letztens habe ich wegen einer kleinen Recherche für ein anderes Projekt mich
mit `ping` in OpenBSD beschäftigt.

Anfang 2015 gab es diesen
[Commit](https://github.com/openbsd/src/commit/08eef1f27acac7f50229bbf7e098d60a720e9b86#diff-ae88590b6e798b8577758800d3fce759)
in OpenBSD base, der Randomness in Ping Pakete einführt. Ping Crypto/Security
was wie wo? Irgendwie wollte ich da noch etwas genauer verstehen um was es
geht.

## ICMP

Zunächst geht es dabei erstmal um ein `ICMP Echo` Paket. Ich hab im
[RFC792](https://tools.ietf.org/html/rfc792) gelernt das jedes ICMP Paket
etwas anders aussieht, vom Aufbau und den Fields einfach her.

{{< figure src="/uploads/2019/02/rfc.png" >}}

Das ICMP Paket in der Abbildung ist aber das was mich jetzt im Detail
interessiert.

## Wireshark

Ich habe dann irgendwie mir etwas Netzwerk Traffic anschauen wollen und habe
dafür Wireshark (bzw. `tshark` benutzt, weil Terminal usw.)

Bei der Gelegenheit fiel mir auch gleich auf, wie lange ich schon kein Packet
Inspection mehr gemacht hab.

Ich fing also mit einem simplen, `tshark icmp` an um mir einfach allen
Traffic auf meinem Laptop anzeigen zu lassen der mit ICMP zu tun hat.

Ich wusste aber noch das ich mir auch einzelne Felder der Pakete anzeigen
lassen kann. Aber wer weiss schon auswendig wie die heissen. Da hilft
`tshark -G fields`  weiter. Es zeigt alle filterbaren Fields an.

```
tshark -T fields -e icmp.data_time -e data ‘icmp’
```

Das Ergebnis sieht dann so aus.

{{< figure src="/uploads/2019/02/macshark.png" >}}

## macOS ping

Mit meinem Paketfilter war ich also schonmal gut unterwegs mir anzuschauen.
Etwas freudig schickte ich also einen `ICMP Echo Request` vom MacBook zu
meinem OpenBSD Server (und nach kurzer `pf` Konfiguration) erhielt ich auch
ein `ICMP Echo Reply`. Freudig schaute ich mir die Daten an.

```
Feb 16, 2019 10:18:10.755628000 UTC     08090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f3031323334353637
```

Das kann ich dann auch irgendwie in Wireshark GUI in kodiert ansehen. Schön
zu sehen hierbei ist, wie die Packet data load einfach aufsteigende Hex Werte
ausgibt und es irgendwann in den `ASCII` Space reinfällt und dann mit ein
paar Sonderzeichen und den Zahlen 0 bis 7 aufhört.

{{< figure src="/uploads/2019/02/icmpreply1.png" >}}

Da wäre auch mit Sicherheit noch mehr gestanden, wenn nicht die Paketlänge
vorher erreicht gewesen wäre.

## Packet Data

Aber Irgendwas ist falsch. Jedes Paket sieht so aus. Kein Anzeichen von
RANDOM, irgendwelchen ChaCha Streams oder Datetime Offsets.

{{< figure src="/uploads/2019/02/mactrace.png" >}}

Bis auf die Uhrzeit ist das Data Field Identisch. **Aber ich habe es ja auch
falsch gemacht.**

## Ping von OpenBSD

Um die Vorteile von den OpenBSD Changes nutzen zu können muss man einen Ping
**verschicken** und nicht empfangen. Logischerweise muss OpenBSD ja auf einen
Echo Request, mit dem selben Content antworten des es empfangen hat. Ein
**ECHO** eben.

Also dieses mal `tshark` auf dem OpenBSD angeworfen, nicht auf dem MacBook.
Meine Server 1 pingt nun Server 2. Und es sieht… anders aus.

```
$ tshark -T fields -e icmp.data_time -e data 'icmp'
4a87cd44ef7cf000af2189efac649e8efa0f1fb12c7e993818191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f3031323334353637
4a87cd44ef7cf000af2189efac649e8efa0f1fb12c7e993818191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f3031323334353637
4a87cd44ef7cf001af2189efac6490ff3370fb9a4f37946318191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f3031323334353637
4a87cd44ef7cf001af2189efac6490ff3370fb9a4f37946318191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f3031323334353637
```

Kein `data_time` Field. Ich hab erst gedacht das Flag heisst vielleicht
anders, aber OpenBSD schickt einfach im `Echo Request` nicht seine lokale
Zeit mit. Warum auch?

Das sieht man nochmal gut wenn man von beiden Betriebssystemen den gleichen
Server pingt und `tshark icmp` anschaut:

{{< figure src="/uploads/2019/02/both.png" >}}

## Randomization

Okay. Daten Fehlen. Aber wo ist jetzt der ChaCha20 Stream? Wenn man also ein
paar Pakete sammelt, sieht der geübte Hex-Lexer (der ich nicht bin)
variierende teile im Data der `Echo Requests`:

macOS:

```
0000   08 09 0a 0b 0c 0d 0e 0f 10 11 12 13 14 15 16 17   ................
0010   18 19 1a 1b 1c 1d 1e 1f 20 21 22 23 24 25 26 27   ........ !"#$%&'
0020   28 29 2a 2b 2c 2d 2e 2f 30 31 32 33 34 35 36 37   ()*+,-./01234567
```

{{< figure src="/uploads/2019/02/packetmac.png" >}}


OpenBSD:

```
0000   12 d5 a4 11 8a 05 63 89 d9 1e 6e 56 be 90 00 60   .Õ¤...c.Ù.nV¾..`
0010   d8 ef 87 a7 1d bb e0 a7 18 19 1a 1b 1c 1d 1e 1f   Øï.§.»à§........
0020   20 21 22 23 24 25 26 27 28 29 2a 2b 2c 2d 2e 2f    !"#$%&'()*+,-./
0030   30 31 32 33 34 35 36 37                           01234567
```

{{< figure src="/uploads/2019/02/packetopenbsd.png" >}}


## Summary

Das war mal wieder ein längerer Blogpost, in der ich mit viel so low-level
Kram herumgetan hab. Low dabei im doppelten Sinne: im Layer und im Anspruch.
Kompliziert ist das eigentlich alles nicht, aber ich hab mal wieder Tools
benutzt und Dinge gelernt. Ob dieser Randomization Part jetzt tatsächlich
dazu beiträgt das ICMP Packets weniger gut zum OS Fingerprinting beitragen
bleibt für mich etwas rätselhaft. Was allerdings gut ist, ist das kein
Timestamp im ICMP Paket drin steht.





