---
categories:
- code
- web
date: 2016-10-31T22:12:14+01:00
tags:
- magic
title: Magic the Gathering - The Archivist
---

[Magic](http://magic.wizards.com) ist ein Sammelkartenspiel. Seit mein Cousin mir
mit 10 ein paar Karten geschenkt hat geht das so. Ich spiele nie wirklich
regelmäßig, noch kenne ich mich besonders gut aus. Ob damals zu
Schulzeiten, während der Ausbildung oder auch gerade aktuell wieder.

{{< figure src="/uploads/2016/10/archivist.png" >}}

Das Fantasy Spiel irgendwo zwischen Schach und Poker macht einfach Spaß.
Auch wenn Fähigkeiten und Spielmechanismen in den letzten Editionen etwas
ausarten und gefühlt immer verrückter werden. Deswegen finde ich die
"alten" Karten irgendwie toller. Wahrscheinlich auch weil ich
von Jahr zu Jahr härter am Nostalgieren bin wenn ich die Karten in Händen
halte.

Jetzt gibt es da dieses Kartenarchiv
[gatherer.wizards.com](https://gatherer.wizards.com) bei dem man sich alle
Karten ansehen kann die es so gibt. Das ist ziemlich toll, weil man sich
die alten Karten alle ansehen kann die man will.

Ich scrollte lange durch die Site und noch eine Edition und noch eine
Edition... Irgendwann nervte mich das aber und ich wollte die Images der
Karten auch haben. Ich bastelte ein kleines Python Skript und lud mir
die Karten herunter.

{{< figure src="/uploads/2016/10/archivist-demo.png" >}}

Das Skript gibts als gist
[hier](https://gist.github.com/noqqe/b0f4b24649f62154bc9f307cd867842e)
So kann ich jetzt meine Suchen im Webinterface definieren und dann die URL
vor das Skript werfen.

Alle Karten von [Urzas Saga](https://en.wikipedia.org/wiki/Urza_block) oder [Merkadische Masken](http://magic.wizards.com/de/game-info/products/card-set-archive/mercadian-masques) runterladen? Yaay!

    mkdir mercadian-masques
    ./archivist.py 'http://gatherer.wizards.com/Pages/Search/Default.aspx?action=advanced&set=+[%22Mercadian%20Masques%22]' mercadian-masques

