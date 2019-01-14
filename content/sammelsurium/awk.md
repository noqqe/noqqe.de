---
title: awk
date: 2011-07-26T12:38:13
tags:
- Programming
---

Eine kleine Sammlung von Schnippseln die hilfreich sein könnten.

## Genereller Aufbau

```
awk 'program' input-file1 input-file2 ...
awk -f program-file input-file1 input-file2 ...

rule/pattern { action }
rule/pattern { action }
```

## Beispiele

Zeug das ich eigentlich immer wieder brauche...

Zusammenzählen alle Werte einer Spalte

```
awk '{ SUM += $1} END { print SUM * 48 }' /tmp/fos
```

Einfache Loop

```
awk '{for(i=3;i<=NF;++i)print $i}'
```

Zeige die Anzahl der Chars der längsten Input Zeile

```
awk '{ if (length($0) > max) max = length($0) } END { print max }' data
```

Variablenzuweisung auf der Commandline

```
$ awk -v sq="'" 'BEGIN { print "Here is a single quote <" sq ">" }'
-| Here is a single quote <'>
```

For Loop für multiple Spalten echos

~~~
awk -F, '{ for (i = 1; i <= NF; i++)
         print $i
}'
~~~

## Quotes in awk print

`\x27` heisst das Zauberwort

~~~
echo $x | awk -F, '{print "date.addColumn(\x27" $2 "\x27, \x27" $1 "\x27);" }'
~~~