---
date: '2009-10-12 18:57:37'
layout: post
slug: debian-etclogrotate-d
status: publish
comments: true
title: Debian | /etc/logrotate.d/*
alias: /archives/684
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
