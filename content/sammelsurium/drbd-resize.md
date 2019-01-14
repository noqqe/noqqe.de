---
title: DRBD-Resize
date: 2013-02-19T17:13:54
tags: 
- Filesystems
---

~~~
$ echo "- - -" > /sys/class/scsi_host/host2/scan
$ fdisk -l
$ cfdisk /dev/sdb
$ pvcreate /dev/sdb1
$ vgextend vg01 /dev/sdb1
$ lvresize -L +40G /dev/vg01/data
$ drbdadm resize data
$
$ cat /proc/drbd
version: 8.3.7 (api:88/proto:86-91)
GIT-hash: ea9e28dbff98e331a62bcbcc63a6135808fe2917 build by root@db01, 2012-11-07 05:48:54

 1: cs:SyncSource ro:Primary/Secondary ds:UpToDate/Inconsistent A r----
    ns:1691598948 nr:441184 dw:1687765316 dr:505882802 al:1257161 bm:734 lo:0 pe:9 ua:0 ap:0 ep:1 wo:b oos:38172436
        [=>..................] sync'ed: 10.1% (37276/41456)M
        finish: 0:54:05 speed: 11,704 (10,240) K/sec
~~~