---
title: OpenBSD strace
date: 2014-05-11T22:13:49
tags:
- OS
- OpenBSD
---

Wie bei Strace debugging

``` bash
ktrace -t + python 1.py    ## <- Got error like before: "ImportError: Cannot load specified object"
kdump -f ktrace.out | grep -2 -i 'permission denied'
```
