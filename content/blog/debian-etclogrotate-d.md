---
aliases:
- /archives/684
- /blog/2009/10/12/debian-etclogrotate-d
- /blog/2009/10/12/debian-etclogrotate-dot-d-star
date: '2009-10-12T16:57:37'
tags:
- rotate
- logging
- logs
- linux
- administration
- filesystem
- weekly
title: Debian | /etc/logrotate.d/*
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