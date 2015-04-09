---
date: 2009-11-12T16:01:29+02:00
type: post
comments: true
title: bash | Server updaten gekürzt
aliases:
- /blog/2009/11/12/bash-server-updaten-gekurzt
- /archives/709
categories:
- Development
- Linux
- Shell
tags:
- apt
- upgrade
- update
- aptitude
---

So faul wie ich also bin, mich bei jedem Server einzeln einzuloggen,
upzugraden und cronjob auszuführen, wollt ich ein Stück weit
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
