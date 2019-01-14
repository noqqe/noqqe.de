---
title: dpkg
date: 2012-01-12T12:53:04
tags:
- Software
- dpkg
---

## Overwrite bei Konflikten

    dpkg --force-depends -r rsync
    dpkg --force-overwrite -i <paket>.deb

## dpkg 1.16 version-policy enforcing

Policy Problematik. Habe ich das Problem evtl bald?

```
$ apt-cache policy dpkg
dpkg:
  Installed: 1.15.8.12
  Candidate: 1.15.8.12
  Version table:
 *** 1.15.8.12 0

$ LANG=C dpkg -l 2>/dev/null | awk '{print $3}' | egrep -v '^[[:digit:]]'
```

### Lösung 1

Wenns im Available vorkommt. [DPKG Error](http://www.linuxquestions.org/questions/debian-26/dpkg-is-dead-error-in-var-lib-dpkg-available-225508/)

```
dpkg --clear-avail
apt-get update
```

### Lösung
