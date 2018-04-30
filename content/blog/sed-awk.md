---
title: "sed & awk"
date: 2018-04-30T18:00:28+02:00
tags:
- sed
- aws
- code
- gedanken
---

Letztens stolperte ich über dieses Buch. Es ist von 1992 und beschreibt die
Funktionalität von `sed` und `awk`. Beim durchblättern merkte ich: Das meiste
in diesem Buch funktioniert noch.

{{< figure src="/uploads/2018/04/sedawk.jpg" >}}

Es funktioniert einfach. Beides Tools die bis heute von Admins und
Unix-Menschen benutzt werden und wahrscheinlich in tausenden Skripten Einzug
gefunden haben. Man merkt schon, irgendwie saß ich da und hab mich beim
herumsinnieren über dieses Buch erwischt.

Ich habe es nicht komplett gelesen aber alleine die Tatsache das Entwickler
ein Programm aus der Senke gehoben haben welches 25 Jahre eine stabile,
kompatible Sprache vorweisen kann hat mich ziemlich beeindruckt. Auch `C` und
andere Software aus dieser Zeit tun noch genauso wie früher. Klar war nicht
alles am ersten Tag toll, aber nimm heutzutage mal ein 5 Jahre altes Buch über
`Docker` oder `Puppet` oder `MongoDB` oder `JavaScript` in die Hand. Damit
kannst du bestenfalls noch deinen Bildschirm in eine angenehmere Position
bringen.

Auf der anderen Seite las ich letztens einen Toot, der sagte "We've lost the
war on package managent".

> I think we've lost the battle of package management. I don't see how we can
> dig ourselves out of this hole. Rust, Node, Go, Python, Elixir, Ruby...
> Everything rolled their own package management, virtual environments, the
> ability to pin/lock to specific versions of code. All of this is wholly
> incompatible with traditional package management.
>
> The battle is over. The next generation of programmers reinvented the wheel
> and it's full of razorblades. I think we should focus our resources on
> containing this mess with jails/containers. There's no other viable
> solution. It's our grim reality now.
>
> by [feld](https://bikeshed.party/notice/543744)

Was sich daraus ergibt sind eigentlich 2 Welten in einem Tech Stack der sich
immer weiter voneinander entfernt. Die gut abgehangnen Coreutils, C Code
der von OS zu OS weiter maintained wird usw. und auf der anderen Seite neumodische
Softwareprojekte, die es teilweise nichtmal schaffen aktuelle
Dokumentation und Schnittstellen zu definieren (sofern überhaupt vorhanden)
geschweige denn halbwegs kompatibel zu halten.

Hier, meine 2 Pfennig.
