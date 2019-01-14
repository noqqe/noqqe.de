---
title: ktrace
date: 2015-12-03T14:10:03
tags: 
- ktrace
- Software
---

Unter OpenBSD gibts kein strace
Deswegen gibts ktrace.

Um einen Prozess zu tracen:

    ktrace -p PID

Tracing wieder beenden

    ktrace -C

Den Output dann lesen

    kdump -f ktrace.out