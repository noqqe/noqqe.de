---
title: doas
date: 2016-06-07T18:36:00
tags: 
- Software
- OpenBSD
- doas
---

#### Commandline Options

Get a new shell as root

    doas -s

execute command as root

    doas echo foo
    doas whoami

#### doas.conf

Allow full access to every command

    permit nopass noqqe as root
    permit nopass noqqe as otheruser
    permit nopass keepenv root as root

Access to a single command

    permit nopass noqqe as root cmd /bin/whoami
