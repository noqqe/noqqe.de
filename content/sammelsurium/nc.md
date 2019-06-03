---
title: nc
date: 2014-04-03T15:20:17
tags:
- Software
- netcat
- nc
---

UDP Listen und Send Test

Host1:

    nc -u -l 0.0.0.0 51666

Host2:

    echo foo | nc -u Host1 51666

