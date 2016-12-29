---
comments:
- author: tux.
  content: "Ein \xFCber 30 Jahre altes St\xFCck Software, das seinen Ursprung (!)
    in einer bald 22 Jahre alten Software hat?\n\nHast du keine Taschenrechnersoftware
    gefunden? ;-)"
  date: '2015-01-24T15:31:13.681734'
- author: noqqe
  content: "Naja in dem LICENSE file im Repo liegt steht halt 1980 an ner Universit\xE4t"
  date: '2015-01-25T13:36:26.058430'
- author: tux.
  content: "Das sind die Standardangaben der BSD-Lizenz. Die gelten auch f\xFCr sp\xE4tere
    Programme. :)"
  date: '2015-01-25T15:51:20.557570'
- author: Martin
  content: "> Das sind die Standardangaben der BSD-Lizenz.\n\nF\xE4llt mir schwer
    das zu glauben, wenn es die Lizenz vor 1988 noch gar nicht \xF6ffentlich gab.\n\n`leave`
    gab's wohl schon in 3BSD, also wohl \xE4lter als 30 Jahre ;)"
  date: '2015-01-25T19:02:49.250528'
- author: noqqe
  content: du bist einfach der Beste <3
  date: '2015-01-26T08:15:24.504784'
date: '2015-01-24T08:34:00'
tags:
- openbsd
- cli
- commandline
- leave
- netbsd
title: leave
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