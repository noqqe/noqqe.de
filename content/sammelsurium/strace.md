---
title: strace
date: 2017-05-22T08:17:55
tags:
- Software
- strace
---

`strace` an einen laufenden Prozess hängen.

``` bash
strace -p 6637
```

Mit `-e` können Syscalls getracked werden. Nützliche sind:

* `open` - File wird geöffnet
* `execve` - Subcommand wird ausgeführt
* `unlinkat` - File wird gelöscht

Beispiel

``` bash
strace -p 22212 -o output.txt -e open
```

Auch schön ist die statistische Auswertung mit `-c`.