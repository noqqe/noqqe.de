---
aliases:
- /archives/1732
comments:
- author: Hesmon
  content: "<p>K\xFCrzer:<br><a href=\"http://www.commandlinefu.com/commands/view/1792/rot13-using-the-tr-command\"
    rel=\"nofollow\">http://www.commandlinefu.com/commands/view/1792/rot13-using-the-tr-command</a></p>"
  date: '2011-08-13T20:30:46'
- author: noqqe
  content: <p>Schade, dass du den Post nichtmal gelesen hast.</p>
  date: '2011-08-13T20:45:12'
- author: Frank
  content: <p>Wo kann ich das downloaden was du gebastelt hast?</p>
  date: '2011-08-14T02:22:50'
- author: noqqe
  content: <p>In dem eingebundenen JavaScript von Github gibts einen Knopf mit "View
    raw"</p><p>Hier ist der Link:</p><p><a href="https://gist.github.com/raw/1143762/a88606a269ed45945b2dead9424886a04ef9fe6b/rot13.bash"
    rel="nofollow">https://gist.github.com/raw/1143762/a88606a269ed45945b2dead9424886a04ef9fe6b/rot13.bash</a></p>
  date: '2011-08-14T15:48:54'
- author: cp
  content: "<p>Rot13 ist zwar sch\xF6n einfach, aber man kann nat\xFCrlich sehr leicht
    den Klartext wiederherstellen. Die Vigen\xE8re-Chiffre dagegen ist ohne Wissen
    um den Schl\xFCssel \"von Hand\" nicht leicht entschl\xFCsselbar, aber nicht wirklich
    kompliziert anzuwenden. F\xFCr mich also als Tipp an dich f\xFCr weitere Spielerein
    :)</p>"
  date: '2011-08-15T04:39:34'
- author: noqqe
  content: "<p>@cp: danke f\xFCr den Tipp! Schau ich mir mal an :) aber sieht h\xE4rter
    aus. Um einiges ;)</p>"
  date: '2011-08-16T18:59:06'
date: '2011-08-13T12:03:20'
slug: rot13-verschlusselung-in-bash
tags:
- development
- decryption
- /dev/radio
- shell
- encryption
- rot13
- ccc
- linux
- very simple encryption
- rot
- debian
- bash
title: "ROT13 Verschl\xFCsselung in Bash"
---

Seit ich meinen neuen Arbeitsweg antrete und täglich ca eine Stunde im Zug
verbringe, höre ich umso mehr Podcasts. Besonders gut finde ich
[http://ulm.ccc.de/dev/radio/](http://ulm.ccc.de/dev/radio/). Es ging unter
anderem um Algorithmen und Kryptologie im Allgemeinen. Unter anderem eben
auch die ROT13 Verschlüsselung.

{{< figure src="/uploads/2011/08/475px-ROT13_table_with_example.svg_.png" >}}

Ich habe mir überlegt wie schwierig es wohl sein kann, diesen in Bash zu
implementieren. Nach kurzem googlen findet man immer wieder eine kleine
wirklich ausgefuchste, trickreiche Zeile:

``` bash
tr A-Za-z N-ZA-Mn-za-m
```

Irgendwie war mir das aber zu langweilig, mir die Arbeit von einem
vorgefertigten Binary erledigen zu lassen. Ich wollte es in echtem Bash
selbst schreiben. Was ich dann auch getan habe.

Usage:

``` bash
$ ./rot13 "hello world"
uryyb jbeyq
```

``` bash
$ echo "hello world" | ./rot13
uryyb jbeyq
```

``` bash
# Verschlüsseln und entschlüsseln hintereinander.
$ echo "hello world" | ./rot13 | ./rot13
hello world
```

```
# Zeichentabelle anzeigen
$ ./rot13 -t
```

Sollten Zeichen vorkommen, die sicht nicht im Table befinden, werden diese
automatisch erkannt und bleiben unverschlüsselt.

```
$ ./rot13 "Ich wollte nur [...] und dann ist das Universum explodiert."
Ipu jbyygr ahe [...] haq qnaa vfg qnf Uavirefhz rkcybqvreg.
```

Genau genommen ist es sogar mehr als nur ROT13. Mit wachsendem
Zeichen-Table wächst auch automatisch die Verschiebung der Stellen mit. Es
ist also möglich einen eigens definierten Table mit Zeichen anzulegen und
diesen zu nutzen. Einzige Bedingung: Es muss eine gerade Anzahl von Zeichen
sein.

Das Skript ist [auf Github](https://gist.github.com/noqqe/1143762/) zu finden.
