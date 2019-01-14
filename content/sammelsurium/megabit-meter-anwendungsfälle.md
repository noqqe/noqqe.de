---
title: Megabit Meter Anwendungsfälle
date: 2011-07-24T14:58:57
tags: 
- Programming
- Bash
---

https://noqqe.de/blog/2011/08/27/arduino-ich-baute-ein-megabitmeter/

Zufaellige Zahlen

    while true; do RND=$(($RANDOM % 99 * 10)); printf "$RND\n" > /dev/ttyUSB0 ;echo $RND; sleep 2 ; done

Zombie Kill Meter

    while true ; do mysql -u zombies_re -pXXX -e \"SELECT kills from zombies.zre_kills ORDER BY id DESC LIMIT 1;\" | grep -v ^kills ; sleep 3; done > /dev/yUSB0

Port 80 Verbindungen:

    while true; do echo \$(( $(netstat -tapn | grep -c -e ':80\s*') * 100 )) ; sleep 2; done  > /dev/ttyUSB0