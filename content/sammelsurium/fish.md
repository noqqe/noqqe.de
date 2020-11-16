---
title: "Fish"
date: 2020-11-16T14:54:06+01:00
tags:
- Software
- Databases
- OS
- AWS
---

Fish Cheatsheet

<!--more-->

## Loops

Ohne do und done

```fish
for x in *.txt ; foo ; end
```

## Redirects

Output

```
To read standard input from a file, write <SOURCE_FILE
To write standard output to a file, write >DESTINATION
To write standard error to a file, write 2>DESTINATION [2]
To append standard output to a file, write >>DESTINATION_FILE
To append standard error to a file, write 2>>DESTINATION_FILE
To redirect both standard output and standard error, write &> all_output.txt

```

Pipe Redirection

```
echo lol | grep lol
make fish 2>| less  # Std Error Redirect
```


## Globale Vars

```
set -x FOO BAR
```

## Search and Replace

Wie bei Bash Variablen Expansion, kann bei Fish auch Strings bearbeitet
werden.

```fish
string replace -a ' ' '_' 'spaces to underscores'
```

```fish
for x in *.HEIC ; convert $x (string replace HEIC jpg $x) ; end
```

https://fishshell.com/docs/current/cmds/string-replace.html?highlight=string
