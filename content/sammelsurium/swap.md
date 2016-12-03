---
title: Swap
date: 2011-11-22T10:08:45
tags: 
- Linux
- Swap
---

## Temporäres Swapfile hinzufuegen

    $ swapon -s ## anzeigen lassen
    $ dd if=/dev/zero of=tempswap bs=1M count=1024
    $ sudo mkswap tempswap
    $ sudo swapon tempswap
