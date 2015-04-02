---
type: post
title: "Analog BitCoin Ticker"
date: 2014-03-08T13:39:00+02:00
comments: true
categories:
- osbn
- ubuntuusers
- Hardware
tags:
- bitcoin
- meter
- megabitmeter
- arduino
- currency
- analog
- api
- bitstamp
---

[BitCoin](https://bitcoin.org). Cheffinnen von Trading-Sites werden
[tot in Wohnungen gefunden](http://www.focus.de/finanzen/news/polizei-kein-natuerlicher-tod-chefin-von-bitcoin-boerse-in-wohnung-tot-aufgefunden_id_3667398.html).
Ein [Magic the Gathering Online Exchange](http://mtgox.com) geht [insolvent](http://www.zeit.de/digital/2014-02/bitcoin-mtgox-insolvenz).
[FlexCoin gehackt](http://www.spiegel.de/netzwelt/web/flexcoin-bitcoin-handelsplattform-ausgeraubt-a-957103.html).

Der Kurs zeigt sich trotzdem relativ unbeeindruckt, was mich irgendwie
verblüfft. Egal. Es macht momentan Spass BitCoin zu verfolgen. Gerade die "Suche" nach
[Satoshi Nakamoto](http://www.manager-magazin.de/finanzen/boerse/phantomsuche-um-bitcoin-erfinder-satoshi-nakamoto-a-957484.html).

{{< figure src="/uploads/2014/03/nakamoto.jpg" >}}

2011 bastelte ich einen [Megabitmeter](/blog/2011/08/27/arduino-ich-baute-ein-megabitmeter/) und als er mir letztens wieder
in die Hände fiel, dachte ich mir hey, BitCoins&lt;-&gt;Dollar anzeigen lassen!

{{< figure src="/uploads/2014/03/bitcoinmeter.png" >}}

Im Endeffekt hole ich mir nur mit etwas Python (schlechtes Python, btw) den
aktuellen Kurs vom letzten halbwegs stabilen BitCoin Trader
[Bitstamp](https://bitstamp.net)

``` python
#!/usr/bin/python

import json
import urllib2
import math
import time

# configure
url = "https://www.bitstamp.net/api/ticker/"
analogmeter = "/dev/ttyUSB0"

# runtime loop
while True:

    # parse api json result
    data = json.load(urllib2.urlopen(url))
    value = math.floor(float(data["last"]))
    value = str(int(value)) + '\n' # use ALL the data types

    # write to usb device
    f = open(analogmeter, 'w')
    f.write(value)
    f.close()

    # wait until the next refresh
    time.sleep(60)
```
