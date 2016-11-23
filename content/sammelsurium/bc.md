---
title: bc
date: 2014-01-07T17:25:07.000000
tags: 
- Software
- bc
---


## pure bc syntax

~~~ { .bc }
scale=2;
r=7.8+0.5-(50 - 234) / 60 / 60);
g=(($(date +%s ) - $begin) / 60 / 60)-0.5;
print "Ausstehend: ";
 if(r<1) print 0;
print r ;
print "\n";
print "Gearbeitet: ";
 if(g<1) print 0;
print g ;
print "\n";
~~~

## Programmatisches Beispiel in Bash

~~~ { .bash }
 begin=$(date +%s -d "$1")
 (
 echo "scale=2;"
 echo "r=7.8+0.5-(($(date +%s ) - $begin) / 60 / 60)"
 echo "g=(($(date +%s ) - $begin) / 60 / 60)-0.5;"
 echo 'print "Ausstehend: "; if(r<1)print 0; print r ; print "\n";'
 echo 'print "Gearbeitet: "; if(g<1) print 0; print g ; print "\n";'
 ) | bc -l

 ## calc finish time
 end=$((begin+29880))
 echo "Ende: $(date +%H:%M:%S -d @${end})"
~~~
