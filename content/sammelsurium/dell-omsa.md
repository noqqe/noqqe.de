---
title: Dell OMSA
date: 2012-06-05T13:32:53
tags: 
- Software
- Dell
- Omsa
---

#### Allgemeine Hilfe

    $ omreport -?
    $ omreport storage -?
    $ omreport chassis -?
    $ omreport system -?

#### Storage Commands

Informationen über Storage Controller

    $ omreport storage controller

Informationen über Disks von Controller

    $ omreport storage pdisk controller=0

    $ omreport storage pdisk controller=0 | egrep '^(ID|Status)'

#### Chassis Commands

RAM Zustand anzeigen lassen

    $ omreport chassis memory

#### System Commands

Allgemeines Alertlog anzeigen

    $ omreport system alertlog

#### Alertlog für Storage exportieren

    $ omconfig storage controller action=exportlog controller=0
    $ less /var/log/lsi*.log
