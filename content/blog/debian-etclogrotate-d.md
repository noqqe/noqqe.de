---
date: 2009-10-12T18:57:37+02:00
comments: true
title: Debian | /etc/logrotate.d/*
aliases:
- /archives/684
- /blog/2009/10/12/debian-etclogrotate-d
- /blog/2009/10/12/debian-etclogrotate-dot-d-star
tags:
- Administration
- Linux
- logs
- logging
- rotate
- filesystem
- weekly
---

Eigene Logs rotieren zu lassen ist mit logrotate eigentlich ziemlich entspannt.
File angeben, Optionen definieren und tut was es soll.

```
/var/log/backup.log {
    nocompress
    missingok
    notifempty
    rotate 4
    weekly
}
```

Schön. Selbsterklärend. Irgendwie.
