---
date: 2009-10-12T18:57:37+02:00
type: post
slug: debian-etclogrotate-d
status: publish
comments: true
title: Debian | /etc/logrotate.d/*
aliases:
- /archives/684
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
