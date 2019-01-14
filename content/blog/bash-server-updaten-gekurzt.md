---
aliases:
- /blog/2009/11/12/bash-server-updaten-gekurzt
- /archives/709
date: '2009-11-12T14:01:29'
tags:
- development
- upgrade
- update
- apt
- shell
- aptitude
- linux
title: "bash | Server updaten gek\xFCrzt"
---

So faul wie ich also bin, mich bei jedem Server einzeln einzuloggen,
upzugraden und Cronjob auszuführen, wollt ich ein Stück weit
automatisieren. Dabei raus kam:

``` bash
#!/bin/bash

uhost=$(echo $1)

if [ $(echo $#) -gt 1 ]
  then
  echo "too many parameters"
else
  if [ -n "$uhost" ]
  then
    echo "Connecting to Host $(host $uhost | awk '{print $1" "$4}')
    ssh root@$uhost "if [ -x /etc/cron.daily/apt-update ]; then aptitude upgrade && /etc/cron.daily/apt-update ; fi"
  fi
fi
```

Usage: supgrade zwetschge.org
Automatisierungen bieten zwar Fehlerquellen, aber Faulheit siegt.
