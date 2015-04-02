---
type: post
title: "The sleepy sorting algorithm"
date: 2012-11-25T14:53:00+02:00
comments: true
categories:
- osbn
- Shell
tags:
- algorithm
- Bash
- 4chan
---

Gerade aus Grunden&reg; herum gegoogelt und das gefunden:

> Man, am I a genius. Check out this sorting algorithm I just invented.

``` bash
#!/bin/bash
function f() {
    sleep "$1"
    echo "$1"
}
while [ -n "$1" ]
do
    f "$1" &
    shift
done
```

Das fand ich so gut, dass ich das jetzt posten musste. Prozess- und Performanceoptimierung
ist natürlich was anderes.

`./sleepsort.bash 5 3 6 3 6 3 1 4 7`

Bzw. der Thread an sich ist auch lesenswert :P

