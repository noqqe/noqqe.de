---
title: "Downwards - Wikipedia auf der Kommandozeile"
date: 2019-12-14T14:13:12+01:00
tags:
- wikipedia
- man
- commandline
---

Manpages haben ein Format
([roff](https://en.wikipedia.org/wiki/Roff_(computer_program)), welches ich
mir noch nie angeschaut habe. Beim Einlesen habe
ich dann gemerkt das Roff eigentlich ein generelles Textsatzprogramm ist. Im
speziellen bei OpenBSD hat sich das Format aber dann mit Fokus auf Manpages
zu [mdoc](https://en.wikipedia.org/wiki/Mandoc) weiterentwickelt.

Ingo Schwarze hat hier für OpenBSD eine ganzen Eimer voll Arbeit geleistet
und neben komplettem Umbau des Toolings um `man`, `mdoc` und `mandoc` auch
weite Teile von Content (zum Beispiel für LibreSSL) neu geschrieben.

Um damit etwas mehr spielen zu können, habe ich ein
[Programm](https://github.com/noqqe/downwards) geschrieben,
welches einen Wikipedia Artikel herunterlädt, den Mediawiki Text parsed und
dann ein `mdoc` file generiert.

``` bash
$ pip3 install downwards
$ downwards -l en 'Blood Quran'
$ downwards --help
Usage: downwards [OPTIONS] ARTICLE

  downwards lets you read a wikipedia page on command line as a manpage.

Options:
  -l, --language TEXT  Language for wikipedia
  -s, --stdout         Print to stdout
  --help               Show this message and exit.
```

Aussehen tut es dann so.

{{< figure src="/uploads/2019/12/blood.png" >}}

Den Teil der den Text parst, will ich mir allerdings nochmal anschauen. Ein
bisschen zu viele Informationen gehen mir im Parser verloren, sodass ich die
nötigen Flags in `mdoc` nicht setzen kann. Aber alles in allem funktioniert
das Teil und ich hab ehrlich schon ein paar Einträge in meinem Terminal
gelesen.

`mdoc` halte ich für etwas verwirrend aber nett um mal damit zu arbeiten.
Es ist halt kein Markdown und zustätzliche Formattierungsmöglichkeiten kosten
eben mehr Komplexität beim Handling.

