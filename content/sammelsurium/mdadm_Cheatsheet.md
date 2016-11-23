---
title: mdadm Cheatsheet
date: 2012-01-10T15:42:58.000000
tags: 
- Software
---


Wie ich das Raid auf z0idberg erstellt hab:

    mdadm --create /dev/md0 --level=5 --raid-disks=4 /dev/sdb1 /dev/sdc1 /dev/sdd1 /dev/sde1

## mismatch_cnt Reparieren


    cat /proc/mdstat
    cat /sys/block/md0/md/mismatch_cnt

    echo repair >/sys/block/md0/md/sync_action
    watch cat /proc/mdstat

jetzt resynced er erstmal. Kontrolle ob alles gut ist:

    echo check >/sys/block/md0/md/sync_action
    watch cat /proc/mdstat
    cat /sys/block/md0/md/mismatch_cnt
