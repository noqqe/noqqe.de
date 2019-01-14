---
title: Swap
date: 2011-11-22T10:08:45
tags:
- OS
- Linux
- Swap
---

# Temporäres Swapfile hinzufuegen

Das Filesystem muss Swapfiles auch unterstützen. BTRFS tut das zum Beispiel nicht.

```
swapon -s # anzeigen lassen
dd if=/dev/zero of=tempswap bs=1M count=1024
sudo mkswap tempswap
sudo swapon tempswap
```
