---
aliases:
- /blog/2010/12/01/bash-lesen-von-devstdin
- /archives/1402
comments:
- author: vain
  content: "<p>Hey,</p><p>du kannst auch explizit das Terminal nochmal \xF6ffnen mit
    /dev/tty. Kann zum Beispiel so aussehen:</p><p>(read -p 'Gib ein: ' -r line &lt;
    /dev/tty; echo \"$line\")  hurz<br>cat  hurz</p><p>Cheers!</p><p>PS.: War dein
    Blog nicht mal im UbuntuUsers-Planeten? Oder kommt da nur nichts mehr an?</p>"
  date: '2010-12-11T23:57:02'
- author: vain
  content: "<p>Huch, da ging irgendwie die H\xE4lfte verloren. Ich wei\xDF mir grade
    nicht anders zu helfen, also:</p><p><a href=\"https://gist.github.com/737572\"
    rel=\"nofollow\">https://gist.github.com/737572</a></p><p>;)</p>"
  date: '2010-12-12T00:00:02'
- author: noqqe
  content: "<p>Oh das ist auch eine sehr gute Idee :) </p><p>Ja mein Blog ist auch
    immernoch im Ubuntuusers Planeten. Aber ich finde im Moment nicht die n\xF6tige
    Relevanz in meinen Posts um diese dort zu posten :)</p>"
  date: '2010-12-12T14:05:36'
date: '2010-12-01T18:27:18'
tags:
- development
- git
- script
- read
- /dev/stdin
- shell
- linux
- lesen
- bash
title: Bash | Lesen von /dev/stdin
---

Gestern hat mich eine Idee für ein Skript beschäftigt. Im Detail wollte ich
die Syntax:

```
$ cat any_file.txt | script.sh
```

abbilden. Lesen von File Descriptor 0. Also Standard-Input. Mich
beschäftigte das Thema aber länger, als ich es erwartet hatte. Auf Sinn und
Unsinn des Skripts möchte zumindest in diesem Blogpost nicht weiter
eingehen ;)

## Versuche

Nach kurzem bemühen einer Suchmaschine fand ich heraus, dass /dev/stdin das
magische Schlüsselwort der Sache ist. Wie also von Standard Input lesen?

```
cat /dev/stdin
cat /dev/stdin > /tmp/foobar
```

War in etwa mein erster Ansatz. Mich wunderte extrem, dass der übergebene
Inhalt zwar ausgegeben wird, aber nicht in dem File steht. Erkenntnis
daraus: File Descriptor 0 lässt sich anscheinend nur einmal auslesen.
Zweiter Versuch:

```
INPUT=$(cat /dev/stdin)
```

Schliesslich wollte ich den Input auch im Script weiter verarbeiten. Erwies
sich aber auch als ungünstig, da Zeilenumbrüche in Variablen irgendwie
verschluckt werden.

## Lösung

Anstatt hier weiterhin mit erfolglosen Lösungsansätzen herum zu schmeissen:

```
cat < /dev/stdin >> /tmp/foobar
```

Sieht unlogisch aus, funktioniert aber besser als alle Anderen.

## /dev/stdin und read

Ich habe es nicht geschafft, nach dem einlesen von /dev/stdin eine
bedienbare

```
read irgendeinevariable
```

Operation durchzuführen. Alle Eingaben werden automatisch mit dem Inhalt
von /dev/stdin gefüllt. Ich konnte dafür leider noch keine Lösung finden.