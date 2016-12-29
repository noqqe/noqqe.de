---
comments:
- author: Anaximander
  content: "<p>B\xF6se Zungen k\xF6nnten jetzt behaupten, dass das Discard Protocol
    beim Support mancher Unternehmen immer noch t\xE4glich genutzt wird. ;-)</p><p>Das
    Echo Protocol auf Port 9 k\xF6nnte man auch noch nennen. Es macht genau das was
    man vermutet. Es sendet einfach das zur\xFCck was man hinschickt.</p>"
  date: '2013-05-07T11:09:29'
- author: noqqe
  content: <p>hehe, like!</p>
  date: '2013-05-07T14:53:12'
- author: Freiheit
  content: "<p>Bei einigen Unternehmen muss es da noch ein mir unbekanntes Protokoll
    geben. Dies senden einen Standardtext auf eine Beschwerde zur\xFCck, welches nichts
    mit dem Problem zu tun hat. ;-)</p>"
  date: '2013-05-08T09:36:24'
date: '2013-05-06T18:11:00'
tags:
- openbsd
- network
- ip
- udp
- tcp
- inetd
- bsd
- xinetd
- protokoll
- openbsd-inetd
- protokolle
title: Ports low as fuck
---

Wenn ich Urlaub habe spiele ich [Dwarf Fortress](http://www.bay12games.com/dwarves/), bastle an
[coffeestats.org](coffeestats.org) und schaue mir Protokolle an, die kein
Arsch mehr kennt.

### Port 9: Discard Protocol

Erster Kandidat. Wow. Ein Protokoll das deine Message nimmt und wegwirft.
Eingeführt wurde es aus Debugging Gründen.  In den meisten Cases möchte ich
wohl eher feststellen, dass das Netz funktioniert. Nicht umgekehrt.
Vielleicht eine Art Null-Hypothese-Protokoll für den TCP/IP Stack.

{{< figure src="/uploads/2013/05/whaaat.gif" >}}

Wie die meisten der nachfolgenden Protokolle ist
[Discard](http://tools.ietf.org/html/rfc863) in (openbsd-)inetd enthalten
und können durch einen einfachen Eintrag aktiviert und zum Spielen benutzt
werden.

    discard   stream  tcp     nowait  root    internal
    discard   dgram   udp     wait    root    internal

### Port 17: QOTD Protocol

Das [Quote-of-the-Day](http://tools.ietf.org/html/rfc865) Protokoll. Total gut.

    $ nc ota.iambic.com  17
    Always listen to experts. They'll tell you what can't be done, and why. Then do it. | Lazarus Long

Es gibt anscheinend eine Hand voll Server die diesen Dienst noch fahren.
Wer braucht da noch einen Zitate Kalender? :)

### Port 19: Character Generator Protocol

Wie bei so vielen RFCs hängt auch hier [Jon Postel](http://en.wikipedia.org/wiki/Jon_Postel)
mit drin. Das [CharGen](http://tools.ietf.org/html/rfc864) Protokoll gibt Zeichenketten zurück.
Interessant ist aber das Verhalten bei TCP und UDP.

{{< figure src="/uploads/2013/05/batman.gif" >}}

Bei TCP schickt einem der Dienst so lange Zeichen bis die Session beendet wird.
Damals muss das ziemlich viel Sinn gemacht haben. Da die Zeichenketten
eindeutig sind konnte man leicht geflippte Bits oder fehlende Pakete ausmachen.

```
telnet localhost 19
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
!"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefgh
"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghi
#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghij
$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijk
%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijkl
&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklm
'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmn
()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmno
)*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmnop
*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmnopq
+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmnopqr
,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmnopqrs
-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmnopqrst
./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmnopqrstu
```

Bei UDP bekommt man eine zufällige Anzahl von Zeichen (zw. 0-512). Zumindest
für [DDOS Attacken](https://isc.sans.edu/diary/A+Chargen-based+DDoS%3F+Chargen+is+still+a+thing%3F/15647) soll es sich ja anbieten.

Konfiguration für inetd

    chargen   stream  tcp     nowait  root    internal
    chargen   dgram   udp     wait    root    internal

### Port 13: Daytime Protocol

Nicht zu verwechseln mit dem Vorgänger von NTP ([Time](http://tools.ietf.org/html/rfc868) Protokoll)
bietet das [Daytime](http://tools.ietf.org/html/rfc867) Protokoll eine für
Debugging und Menschen lesbare Variante eines Timestamps.

    $ nc localhost 13
    > Mon May  6 16:27:33 2013

Die entsprechende inetd Konfiguration

    daytime   stream  tcp     nowait  root    internal
    daytime   dgram   udp     wait    root    internal