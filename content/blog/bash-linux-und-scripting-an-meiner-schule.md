---
date: 2010-05-20T20:03:32+02:00
comments: true
title: Bash | Linux und Scripting an meiner Schule
aliases:
- /blog/2010/05/20/bash-linux-und-scripting-an-meiner-schule
- /archives/1005
tags:
- Development
- Linux
- Shell
- bash
- schule
- scripting
- skript
---

Diese Woche war (bzw. ist immernoch, aber wenn ich in der Vergangenheit
spreche fühlt es sich so an als wärs schon vorbei und das ist gut)
Berufsschulwoche. Nun jedenfalls habe ich die grosse Freude dort in
Bash-Scripting unterrichtet zu werden. Ungewöhnlich für Lehrkörper ist
unserer relativ begeisterter Linux-Benutzer und schiebt an allen Enden
etwas Linux Know-How mit ein. Das ist sehr schön weil ich Bash und Linux
sehr mag.

# ****#!/bin/bash****


Morgen wird es einen Test geben, bei dem wir innerhalb 30 Minuten eine
bestimmte Aufgabe scripten sollen und als "Abgabe" sozusagen vorführen
müssen.

Eine Vorbereitung für diesen Test lautete wie folgt:

> Schreibe ein Script, welches durch Angabe von Parametern in verschiedenen
> Zeitintervallen und maximaler Ausführdauer das aktuelle Datum und die
> Uhrzeit ausgibt. Außerdem soll zwischen 2 Modi gewechselt werden können.
> Mit vorangestelltem Text "Uhrzeit ist:" und nur die Uhrzeit.
>
> Usage Beispiel:
> ./zeitausgabe 2 10 modus


Naja. Bevor wieder diverse Trolle in den Kommentaren ihrem Ruf gerecht
werden: Die Klasse ist bis auf einen %-Anteil von ca 5 komplett
Windows-User und haben Bash-Scripting seit ca 7 Schultagen unterrichtet
bekommen. In dieser Dimension wird also auch der Test Morgen ausfallen. Ich
bin gespannt.

Meine Lösung der Aufgabe:

``` bash
#!/bin/bash
interval=$1
gesamtdauer=$2
modus=$3
count=0

if [ "$modus" = "txt" ]; then
   while [ $count -lt $gesamtdauer ]; do
   let count+=$interval
   sleep $interval ; echo "Das ist die Systemzeit: $(date)"
   done
else
   while [ $count -lt $gesamtdauer ]; do
   let count+=$interval
   sleep $interval ; echo $(date)
   done
fi
```
