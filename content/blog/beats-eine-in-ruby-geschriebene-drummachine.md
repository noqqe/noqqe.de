---
aliases:
- /blog/2011/05/22/beats-eine-in-ruby-geschriebene-drummachine
- /archives/1680
comments:
- author: Jurkan
  content: "<p>Super Sache! Ich hab vor nem halben Jahr mal was vergleichbares (mit
    GUI) gesucht, bin aber nicht f\xFCndig geworden. BEATS entspricht da genau meinen
    Vorstellungen, GUI braucht man bei dem einfachen Format nicht wirklich.</p><p>Das
    Randomsong-Skript dazu ist ganz interessant, aber zum Gl\xFCck reicht meine Kreativit\xE4t
    auch aus, um von Hand was einigerma\xDFen brauchbares zu erzeugen. Oder du setzt
    einfach h\xF6here Ma\xDFst\xE4be an ;-)</p><p>Danke jedenfalls f\xFCr den Tipp.</p>"
  date: '2011-05-23T01:21:01'
- author: noqqe
  content: "<p>Hi Jurkan! Super das es dir gef\xE4llt :) Der Autor von Beats hat mir
    auch neulich Feedback dazu geben und fand es gerade zum Erstellen neuer Ideen
    sehr hilfreich :)</p>"
  date: '2011-05-24T18:02:20'
date: '2011-05-22T18:38:54'
tags:
- development
- sound
- shell
- ruby
- random
- beats
- drums
- drummachine
- bash
title: Beats | Eine in Ruby geschriebene Drummachine
---

Neulich habe ich voller Begeisterung das
[Drummachine](http://beatsdrummachine.com/) Projekt
[Beats](https://github.com/jstrait/beats/) entdeckt. Beats stellt so in
etwa die Musicbox für Nerds dar. In einem Verzeichnis mit einem Set an
\*.wav Dateien (wohl der Standard) legt man eine Art Konfigurationsdatei,
welche durch Beats interpretiert wird. In einer gewissen Syntax beschreibt
man dieses Lied und dessen Abfolge.

{{< figure src="/uploads/2011/05/6459_d5e4.jpeg" >}}

Das Traurige an der Sache ist eigentlich lediglich, dass ich mir bewusst
geworden bin, wie wenig Kreativität/Können ich im musikalischen Bereich
vorzuweisen habe. Teilweise sind beim Herumexperimentieren zwar
Interessante Ergebnisse entstanden, aber nichts wofür ich mich nicht
schämen müsste ;)

```
$ beats song1.txt song.wav
$ open song.wav
```

Ich hab dann allerdings kurzer Hand begonnen, den Vorgang zu
automatisieren. Erst wollte ich es outsourcen, aber dann habe ich es doch
automatisiert. Da die ganze "Ich baue mir ein neues Lied"-Sache nur auf
einem einzigen File beruht, dachte ich mir dass sich das mit Sicherheit
auch automatisch generieren lässt.

Das Skript das dabei herauskam, taufte ich randombeats. Wie alles was ich
in letzter Zeit tue habe ich es natürlich auf github veröffentlicht.

[https://github.com/noqqe/randombeats](https://github.com/noqqe/randombeats)

Benutzung:

Ins jeweilige Verzeichnis mit den Roh-Daten kopieren/wechseln und Skript
ausführen.

```
$ git clone http://github.com/noqqe/randombeats
$ cp randombeats/randombeats.bash /path/to/music
$ cd /path/to/music
$ ./randombeats.bash > randomsong.txt
$ beats randomsong.txt
$ open randomsong.wav
```

Raus kommen allerlei sehr Interessante und manchmal auch schöne Rhythmen.
Aber manchmal auch akute Ohrenschmerzen verursachen. Ein paar davon sammle
ich mittlerweile in einem separaten Branch im Git-Repo. Diesen kann ich
aber nicht uploaden, weil die Größe meines Github Accounts auf 300MB
beschränkt ist. Aber vielleicht finde ich anderweitig Möglichkeiten die
Ergebnisse zu publizieren.

Mittlerweile habe ich folgenden Befehl in einer Schleife laufen.

```
./randombeats.bash > rnd.txt && beats rnd.txt && open rnd.wav
```

An Feedback, Kritik oder Anmerkungen zu Verbesserungen an den Eckdaten des
Skripts bin ich wie immer sehr interessiert :)

Update: Kurze WAV-Files generiert und als Beispiel hochgeladen.

[randombeats example 1](/uploads/2011/05/rnd.mp3)
[randombeats example 2](/uploads/2011/05/wheep.mp3)