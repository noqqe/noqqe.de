---
date: 2011-05-22T20:38:54+02:00
type: post
slug: beats-eine-in-ruby-geschriebene-drummachine
status: publish
comments: true
title: Beats | Eine in Ruby geschriebene Drummachine
alias: /archives/1680
categories:
- Bash
- Coding
- General
- PlanetenBlogger
tags:
- bash
- beats
- drummachine
- drums
- random
- ruby
- sound
---

Neulich habe ich voller Begeisterung das [Drummachine](http://beatsdrummachine.com/) Projekt [Beats](https://github.com/jstrait/beats/) entdeckt. Beats stellt so in etwa die Musicbox für Nerds dar. In einem Verzeichnis mit einem Set an *.wav Dateien (wohl der Standard) legt man eine Art Konfigurationsdatei, welche durch Beats interpretiert wird. In einer gewissen Syntax beschreibt man dieses Lied und dessen Abfolge.

{{< figure src="/uploads/2011/05/6459_d5e4.jpeg" >}}

Das Traurige an der Sache ist eigentlich lediglich, dass ich mir bewusst geworden bin, wie wenig Kreativität/Können ich im musikalischen Bereich vorzuweisen habe. Teilweise sind beim Herumexperimentieren zwar Interessante Ergebnisse entstanden, aber nichts wofür ich mich nicht schämen müsste ;)

```
$ beats song1.txt song.wav
$ open song.wav
```


Ich hab dann allerdings kurzer Hand begonnen, den Vorgang zu automatisieren. Erst wollte ich es outsourcen, aber dann habe ich es doch automatisiert. Da die ganze "Ich baue mir ein neues Lied"-Sache nur auf einem einzigen File beruht, dachte ich mir dass sich das mit Sicherheit auch automatisch generieren lässt.

Das Skript das dabei herauskam, taufte ich randombeats. Wie alles was ich in letzter Zeit tue habe ich es natürlich auf github veröffentlicht.

[https://github.com/noqqe/randombeats](https://github.com/noqqe/randombeats)

Benutzung:

Ins jeweilige Verzeichnis mit den Roh-Daten kopieren/wechseln und Skript ausführen.

```
$ git clone http://github.com/noqqe/randombeats
$ cp randombeats/randombeats.bash /path/to/music
$ cd /path/to/music
$ ./randombeats.bash > randomsong.txt
$ beats randomsong.txt
$ open randomsong.wav
```


Raus kommen allerlei sehr Interessante und manchmal auch schöne Rhythmen. Aber manchmal auch akkute Ohrenschmerzen verursachen. Ein paar davon sammle ich mittlerweile in einem separaten Branch im Git-Repo. Diesen kann ich aber nicht uploaden, weil die größe meines Github Accounts auf 300MB beschränkt ist. Aber vielleicht finde ich anderweitig Möglichkeiten die Ergebnisse zu publizieren.

Mittlerweile habe ich  folgenden Befehl in einer Schleife laufen.

```
./randombeats.bash > rnd.txt && beats rnd.txt && open rnd.wav
```


An Feedback, Kritik oder Anmerkungen zu Verbesserungen an den Eckdaten des Skripts bin ich wie immer sehr interessiert :)

Update: Kurze WAV-Files generiert und als Beispiel hochgeladen.

[randombeats example 1](/uploads/2011/05/rnd.mp3)
[randombeats example 2](/uploads/2011/05/wheep.mp3)



