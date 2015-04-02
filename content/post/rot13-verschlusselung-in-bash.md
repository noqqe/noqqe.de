---
date: 2011-08-13T14:03:20+02:00
type: post
slug: rot13-verschlusselung-in-bash
comments: true
title: ROT13 Verschlüsselung in Bash
aliases:
- /archives/1732
categories:
- Bash
- Coding
- Debian
- git
- Linux
- ubuntuusers
tags:
- /dev/radio
- bash
- CCC
- decryption
- encryption
- rot
- rot13
- shell
- very simple encryption
---

Seit ich meinen neuen Arbeitsweg antrete und täglich ca eine Stunde im Zug verbringe, höre ich umso mehr Podcasts. Besonders gut finde ich [http://ulm.ccc.de/dev/radio/](http://ulm.ccc.de/dev/radio/). Es ging unter anderem um Algorithmen und Kryptologie im Allgemeinen. Unter anderem eben auch die ROT13 Verschlüsselung.

{{< figure src="/uploads/2011/08/475px-ROT13_table_with_example.svg_.png" >}}

 Ich habe mir überlegt wie schwierig es wohl sein kann, diesen in Bash zu implementieren. Nach kurzem googlen findet man immer wieder eine kleine wirklich ausgefuchste, trickreiche Zeile:

```
$ tr A-Za-z N-ZA-Mn-za-m
```


Irgendwie war mir das aber zu langweilig, mir die Arbeit von einem vorgefertigten Binary erledigen zu lassen. Ich wollte es in echtem Bash selbst schreiben. Was ich dann auch getan habe.



Usage:

```
$ ./rot13 "hello world"
uryyb jbeyq
```


```
$ echo "hello world" | ./rot13
uryyb jbeyq
```


```
# Verschlüsseln und entschlüsseln hintereinander.
$ echo "hello world" | ./rot13 | ./rot13
hello world
```


```
# Zeichentabelle anzeigen
 $ ./rot13 -t
```


Sollten Zeichen vorkommen, die sicht nicht im Table befinden, werden diese automatisch erkannt und bleiben unverschlüsselt.

```
$ ./rot13 "Ich wollte nur [...] und dann ist das Universum explodiert."
Ipu jbyygr ahe [...] haq qnaa vfg qnf Uavirefhz rkcybqvreg.
```


Genau genommen ist es sogar mehr als nur ROT13. Mit wachsendem Zeichen-Table wächst auch automatisch die Verschiebung der Stellen mit. Es ist also möglich einen eigens definierten Table mit Zeichen anzulegen und diesen zu nutzen. Einzige Bedingung: Es muss eine gerade Anzahl von Zeichen sein.

Das Skript ist [auf Github](https://gist.github.com/noqqe/1143762/) zu finden.

