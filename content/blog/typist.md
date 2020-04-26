---
title: "Typist - Tippen lernen"
date: 2020-04-26T11:02:22+02:00
tags:
- Go
- Keyboard
---

Ein 40% Keyboard ist eine größere Umstellung als gedacht. Da hilft nur Übung.
Aber wie?

<!--more-->

## GNU Typist

Um mit meiner [neuen Tastatur](/blog/2020/04/01/niu-mini/) klarzukommen, habe ich eine Zeit lang
[gnu-typist](https://www.gnu.org/software/gtypist/) benutzt. Anfangs ging das
auch ganz gut, aber herkömmliches "Schreiben" ist dann doch schnell erlernt.


{{< figure src="/uploads/2020/04/gtypist.png" caption="gnu-typist screenshot" >}}

Die verschiedenen Lektionen sind gut und machen auch fast etwas
Spass. Vor allem wenn man merkt das man besser wird. Aber...

## Typist Eigenimplemntierung

Dann fand ich den Tipptrainer GNU Typist etwas zu unflexibel für meine
Bedürfnisse. Texte schreiben war in akzeptabler Fehlerrate möglich geworden.

Mein Problem das ich zu lösen versuche ist ein Anderes. Sonderzeichen, Shell
und Programmieren. Viele aufeinander folgende `*`, `(`, `>` und `|`. Ich
dachte kurz darüber nach einen Importer für .bash_history in GNU Typist zu
schreiben. Dann aber habe ich mich gefragt wie schwierig es wohl ist selbst
einen Tipptrainer zu bauen.

Und wenn schon, dann gleich in `Go` was ich sowieso gerade lerne. COVID-19
h00ray. Nach einiger Zeit kam dann dabei
[noqqe/typist](https://github.com/noqqe/typist) heraus.

{{< figure src="/uploads/2020/04/typist.png" caption="noqqe/typist screenshot" >}}

Es wird eine "Challenge" vorgegeben und man muss sie Zeichen für Zeichen
nachtippen. Am Ende wird die Fehleranzahl berechnet und die Zeit gemessen,
die die Challenge benötigt hat.

Ich habe dann mehrere Challenge Files geschrieben die in ungefähr so
aussehen:

``` yaml
description: "Smileys are hard on new keyboards. Lets try some."
challenges:
  - :D
  - :)
  - :/
  - :(
  - xD
  - ^^
  - :'(
  - -.-
  - /o\
  - \o/
```

Auch ein schöner Anwendungsfall sind Smileys zu trainieren. Aber im Endeffekt
benutze ich dieses Set an Trainingsdateien:

```
 acronyms.yml
 bash.yml
 fish.yml
 geekthings.yml
 intro-1.yml
 intro-2.yml
 intro-3.yml
 smileys.yml
```

Das Datenformat ist in `yaml` und relativ einfach zu schreiben.

Ich muss zugeben ich habe bisher wahrscheinlich mehr Zeit mit dem Schreiben
des Programms als mit dem eigentlichen Tipptraining verbracht. Aber auch das
`Go` Programm selbst hab ich mit der NIU 40% geschrieben. Und tippen ist
schliesslich auch tippen, oder? ODER?

## Please press F11

Ich habe relativ schnell gemerkt, das ich anfange Wahllos herumzudrücken,
wenn ich ein bestimmtes Sonderzeichen nicht finde. Außerdem kam es in der
Eingewöhnungsphase oft dazu, dass ich Tasten doch nocheinmal anders gemapped habe als ich
ursprünglich dachte. Das war ein Problem da sich die gerade gebildete Muscle
Memory doch wieder falsch war.

Unter [config.qmk.fm](https://config.qmk.fm) kann man
sich ein PDF generieren, welches die eigene Belegung widerspiegelt. Das hilft
um nachsehen zu können, wo welcher Key in der aktuell benutzten Firmware
hingemappt wurde.

Dieses PDF habe ich als Desktop Hintergrund eingestellt. Ganz einfach weil
ich unter macOS mit `F11` diesen anzeigen lassen kann. Somit ist mein
Cheatsheet immer zu Hand wenn ich es gerade brauche. Immer.

{{< figure src="/uploads/2020/04/cheat.gif" >}}

So wurde ich nach einiger Zeit sicherer in Sonderzeichen. In Verbindung mit
dem Tipptrainer kann eigentlich nichts mehr schief gehen.
