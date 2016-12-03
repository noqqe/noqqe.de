---
title: "leave"
date: 2015-01-24T10:34:00+02:00
comments: true
tags:
- OpenBSD
- osbn
- ubuntuusers
- leave
- CLI
- NetBSD
- Commandline
---
Die Tage bin ich auf ein über 30 Jahre altes Stück Software gestoßen, dass
seinen
[Ursprung](http://ftp.rrzn.uni-hannover.de/pub/mirror/bsd/NetBSD/NetBSD-current/src/usr.bin/leave/)
in NetBSD hat. `leave` ist klein, simpel und nützlich.  Tools die aus
Jahrzehnten stammen, wo man als Admin noch Multi-User Großrechner administirert
hat und CLI Programme um das Arbeitsleben herum programmiert hat.

> leave waits until the specified time, then reminds you
> that you have to leave. You are reminded 5 minutes and 1 minute before the
> actual time, at the time, and every minute thereafter. When you log off,
> leave exits just before it would have printed the next message.

Auch heute noch ist es in NetBSD und OpenBSD per default enthalten. Unter Debian
nachinstallierbar.

{{< figure src="/uploads/2015/01/leave2.png" >}}

Es kann auch etwas penetrant werden, jetzt endlich heim zu gehen.

{{< figure src="/uploads/2015/01/leave1.png" >}}

In Zeiten in denen man mit am Retina Display mit 4 Splitwindows, &gt;15
Terminaltabs und multiplen Ebenen von `tmux`-Sessions sitzt aber leider eher
unpraktisch.
