---
title: XFS
date: 2013-03-19T18:47:59
tags: 
- Filesystems
---

XFS Filesystem erstellen

    ## time mkfs.xfs /dev/mapper/data
    log stripe unit (524288 bytes) is too large (maximum is 256KiB)
    log stripe unit adjusted to 32KiB
    meta-data=/dev/mapper/data       isize=256    agcount=32, agsize=45782272 blks
             =                       sectsz=512   attr=2, projid32bit=0
    data     =                       bsize=4096   blocks=1465032704, imaxpct=5
             =                       sunit=128    swidth=384 blks
    naming   =version 2              bsize=4096   ascii-ci=0
    log      =internal log           bsize=4096   blocks=521728, version=2
             =                       sectsz=512   sunit=8 blks, lazy-count=1
    realtime =none                   extsz=4096   blocks=0, rtextents=0

    real    6m13.184s
    user    0m0.016s
    sys 0m1.420s
