---
title: Bash Time Conversion
date: 2012-08-26T10:12:26.000000
tags: 
- Programming
- Bash
---


Gefunden auf:

http://unstableme.blogspot.de/2009/01/convert-seconds-to-hour-minute-seconds.html


~~~ { .bash }
#!/bin/sh
#Convert seconds to h:m:s format

[ -z ${1} ] && echo "Usage: $(basename $0) <seconds>" && exit||secs=${1}
_hms()
{
 local S=${1}
 ((h=S/3600))
 ((m=S%3600/60))
 ((s=S%60))
 printf "%dh:%dm:%ds\n" $h $m $s
}

_hms ${secs}
~~~
