---
date: 2009-01-14T21:40:24+02:00
type: post
slug: awk-und-syncn
status: publish
comments: true
title: awk und syncN
alias: /archives/446
categories:
- Coding
- Linux
tags:
- awk
- bash
- shell
- skript
- Sync
- SyncN
---

awk hat mir tatsächlich geholfen ;) wer hätte das gedacht. Ich glaube ich habe die kleine Skriptsprache deutlich unterschätzt.

Aber von Anfang an: In meinem Skript [syncN](http://zwetschge.org/syncN/) gibt es eine Funktion "--check". Die zwei Files miteinander vergleicht. Bisher nur über das Datum. Das sieht bis jetzt so aus :


    check ()
    {
    	ssh user@host "ls -lah /Route/to/SyncDir/ > /tmp/checkreturn"
    	scp user@host:/tmp/checkreturn /tmp/checkreturn
    	cat /tmp/checkreturn
    	rm /tmp/checkreturn
    }


Die Ausgabe des ganzen ist relativ unschön.

insgesamt 352K
drwxr-xr-x 2 noqqe noqqe 4,0K 2009-01-09 16:10 .
drwxr-xr-x 7 noqqe noqqe 4,0K 2009-01-13 12:29 ..
-rw-r--r-- 1 noqqe noqqe 337K 2009-01-14 20:07 $FILE

Jetzt gibt es da wunderschön, awk.


    ssh user@host "ls -lah /Route/to/SyncDir/ | awk '{ print $6 " " $7 }' > /tmp/checkreturn"


und schon bekommt man nur noch das Datum und die Uhrzeit ;)


    2009-01-14 20:09


Demnächst gibts übrigens zusätzlich zur wunderschönen awk Ausgabe noch einen md5 Summen Vergleich ;)
