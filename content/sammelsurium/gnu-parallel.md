---
title: gnu-parallel
date: 2012-01-06T12:24:56.000000
tags: 
- Software
- gnu
- parallel
---


~~~
## time seq 1 10000 | parallel 'echo {}| md5sum &> /dev/null '

real	0m20.102s
user	0m35.082s
sys	0m24.918s

## time for x in $(seq 1 10000); do echo $x | md5sum &> /dev/null; done

real	0m13.504s
user	0m2.368s
sys	0m3.948s
~~~

Das dreht sich aber schnell sobald die Aufgaben groesser werden:

~~~
## time seq 1 1000 | parallel 'cat /dev/urandom | head -c 100000 | gzip &> /dev/null'

real	0m7.845s
user	0m4.064s
sys	0m20.485s

## time for x in $(seq 1 1000); do cat /dev/urandom | head -c 100000 | gzip &> /dev/null; done

real	0m31.869s
user	0m8.301s
sys	0m33.658s
~~~

multiple Greps auf ein File

    cat uniqueHouseholds.txt | time ./bin/parallel -j150% 'grep ^{}, TR_HOUSEHOLD.csv' > lolo.test

Ein grep auf ein riesen File in verschiedenen Blöcken

    cat bigfile | parallel --pipe --block 2M grep foo
