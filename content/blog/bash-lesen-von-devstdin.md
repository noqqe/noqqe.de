---
date: 2010-12-01T20:27:18+02:00
type: post
comments: true
title: Bash | Lesen von /dev/stdin
aliases:
- /blog/2010/12/01/bash-lesen-von-devstdin
- /archives/1402
categories:
- Bash
- Development
- git
- Linux
- PlanetenBlogger
tags:
- /dev/stdin
- bash
- lesen
- lesen vom stdin
- read
- read from stdin
- script
- shell
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
