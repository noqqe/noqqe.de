---
title: Linux Kernel Schedulers
date: 2012-02-27T15:08:27.000000
tags: 
- Filesystems
---


Citrix XenServer 6 als Dom0 mit Debian Squeeze in einer DomU.
Hat keine Veraenderung gebracht. Getestet mit dd und hdparm.
Sehen welchen IO-Scheduler man verwendet kann man so:

## IO Scheduler auslesen

so gehts...

    root ~ ## cat /sys/block/xvda/queue/scheduler
    noop anticipatory deadline [cfq]

Das was in "[]" steht, wird aktuell verwendet.
Ein Test auf VMware waere mal interessant ob es sich hier anders verhaelt.

## Links

Hintergrund:

http://lonesysadmin.net/2008/02/21/elevatornoop/

Umsetzen in Live Betrieb:

http://www.cyberciti.biz/faq/linux-change-io-scheduler-for-harddisk/
