---
date: 2011-06-12T13:09:28+02:00
comments: true
title: BitCoin | Debian Daemon Script für poclbm-mod
aliases:
- /blog/2011/06/12/bitcoin-debian-daemon-script-fur-poclbm-mod
- /archives/1689
categories:
- Bash
- Debian
- Ubuntu
- ubuntuusers
- Crypto
tags:
- BitCoin
- debian
- Mining
- ubuntu
---

Man muss Dinge über Hirnschäden von Menschen lesen die
[Nachts guten Gewissens neben einem 4 Grafikarten im SLI-Verbund schlafen](http://www.bitcoinminingaccidents.com/?p=196)
und zugleich die [Bedenken von sinnierenden Typen](http://tav.espians.com/why-bitcoin-will-fail-as-a-currency.html) die
an den Limitierungen von 21 Billionen maximal möglichen BitCoins zweifeln.
Umweltverschmutzung ist natürlich auch ein Thema. Klar. Gerechtfertigter
Weise.

{{< figure src="/uploads/2011/06/Tux-G2_bitcoin.png" >}}

Wenn Ihr mich fragt, schiesst die BitCoin Mining Gesellschaft am Ziel
vorbei. Separate Rechner betreiben schiesst am Ziel vorbei. Hunderte von
Euros für neue Grafikkarten ausgeben um 5 MegaHashes/s mehr rechnen zu
können schiesst am Ziel vorbei. Euro's für BitCoins bezahlen schiesst
sowieso am Ziel vorbei. Man kann BitCoins meines Erachtens auch benutzen
ohne Kopfstände zu machen. Immer wenn der Rechner sowieso gerade läuft, mit
der Hardware die man zur Verfügung hat.

Aus diesem Grund habe ich ein kleines Skript gebastelt. Einen
Start-Stop-Daemon für **/etc/init.d/**. Es startet automatisch wenn mein
Rechner hochfährt und hört auf wenn ich Ihn herunterfahre. Ganz einfach

```
$ wget -O /etc/init.d/bitcoin https://gist.github.com/raw/1007794/bitcoin.sh
$ chmod +x /etc/init.d/bitcoin
$ update-rc.d bitcoin defaults
```

[https://gist.github.com/1007794](https://gist.github.com/1007794)

Ich möchte aber dazu sagen, dass ich dies Funktionstüchtigkeit des Skripts
nicht auf anderen Rechnern/Betriebssystemen getestet habe. Vor Benutzung
also bitte lesen, verstehen ggf. anpassen. Außerdem wird ein Account bei
einem Mining Pool benötigt und der Mining Client an sich
([poclbm-mod](https://en.bitcoin.it/wiki/Poclbm-mod)). Während ich diesen
Post geschrieben habe, hat mein BitCoin Mining Client entspannt auf meiner
Geforce 8600 GS mit nahezu niedlichen 950 KiloHashes/s vor sich hin
gemined.
