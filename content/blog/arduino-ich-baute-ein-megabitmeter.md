---
date: 2011-08-27T13:10:02+02:00
comments: true
title: Arduino | Ich baute ein Megabitmeter
aliases:
- /archives/1751
- /blog/2011/08/27/arduino-ich-baute-ein-megabitmeter
tags:
- Hardware
- ubuntuusers
- Web
- Arduino
- ATMega
- controller
- hack
- Hardware
- löten
- Megabit
- Megabitmeter
- Meter
- mikrocontroller
- nsfw
- pritlove
- prozessor
- skytee
- tim
- usb
---

Vor einiger Zeit hab ich über den Podcast von Tim Pritlove (NSFW) von dem
Projekt [MegabitMeter]( http://megabitmeter.de) erfahren. Es hat mich
ehrlich gesagt fasziniert. Der eigentliche (im Namen implizierte) Zweck
zwar weniger, aber generell einfach Werte über ein USB Device darstellen
zu können. Genaue Anwendungsfälle gibts weiter unten.

Zunächst mal hab ich mich ausgiebig mit dem
[HowTo](http://megabitmeter.de/2010/12/megabitmeter-%E2%80%93-bandwidth-meter-diy-kit-howto/)
beschäftigt welches ich (bis auf ein paar Stellen) sehr gut finde. Die
letzten Paar Unstimmigkeiten habe ich dann mit dem wirklich netten und
zuvorkommenden Betreibern des Projekts via Mail abstimmen dürfen :) Das lag
aber auch an meinen fehlenden elektrotechnischen Kenntnissen.

Ich muss dazu sagen, dass ich in der glücklichen Situation war und jemanden
kannte, der in einer Firma tätig ist, die solche Messgeräte anfertigt
wodurch ich keinen Kit kaufen musste. Das ist auch der Grund warum das
Messgerät etwas anders aussieht.

{{< figure src="/uploads/2011/08/IMG_0547.jpg" >}}

{{< figure src="/uploads/2011/08/IMG_0546.jpg" >}}

Ich denke auch das der Sinn des Projekts eher das "do-it-yourself" ist und
der wirtschaftliche Erfolg nur eine positiver Nebeneffekt der Betreiber
ist... (Zumindest hoffe ich, dass ich das nicht falsch interpretiert habe).
Außerdem schwierig war, dass der Kit 2 (incl. dem Arduino) nicht verfügbar
war, da laut Twitter gerade keine Arduinio Nano in rauen Mengen billig
einzukaufen sind. Den Arduino Prozessor habe ich deswegen von
[TinkerSoup.de](http://tinkersoup.de) geordert.

Wenn alles fertig gebastelt ist, kann man dem Gerät einfach via echo Werte
übergeben.

```
$ echo "200" > /dev/ttyUSB0
```

Ein paar Anwendungsbeispiele:

Zufällige Zahlen auf das Gerät projizieren

```
while true; do RND=$(($RANDOM % 99 * 10)); printf "$RNDn" > /dev/ttyUSB0 ;echo $RND; sleep 2 ; done
```

Zombie Kill Meter (In Verbindung zu [zombies.n0q.org](http://zombies.n0q.org))

```
while true ; do mysql -u user -ppw -e "SELECT kills from zombies.zre_kills ORDER BY id DESC LIMIT 1;" | grep -v ^kills ; sleep 3; done" > /dev/ttyUSB0
```

Port 80 Verbindungen des Webservers

```
ssh user@host "while true; do echo $(( $(netstat -tapn | grep -c -e ':80s*') * 100 )) ; sleep 2; done"  > /dev/ttyUSB0
```

An weiteren Snippets bastle ich im Moment noch. Werden eventuell
nachgereicht wenn Sie spruchreif sind. Am Ende noch ein riesen Danke für
die Software und die Projektidee sowie das wunderbare HowTo an das Team von
[megabitmeter.de](http://megabitmeter.de)
