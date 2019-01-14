---
title: Filesystem Caches
date: 2012-10-29T15:45:22
tags:
- Filesystems
---

### Drop

    $ sync
    $ echo 1 > /proc/sys/vm/drop_caches

[linuxquestions.org](http://www.linuxquestions.org/questions/linux-kernel-70/how-to-disable-filesystem-cache-627012/)

### drop_caches

Writing to this will cause the kernel to drop clean caches, dentries and
inodes from memory, causing that memory to become free.

To free pagecache:

    echo 1 > /proc/sys/vm/drop_caches

To free dentries and inodes:

    echo 2 > /proc/sys/vm/drop_caches

To free pagecache, dentries and inodes:

    echo 3 > /proc/sys/vm/drop_caches

As this is a non-destructive operation and dirty objects are not freeable, the
user should run `sync' first.
