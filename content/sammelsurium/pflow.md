---
title: pflow
date: 2015-02-21T00:03:45
tags:
- OS
- OpenBSD
---

#### Sensor installieren

Dynamisch zur Laufzeit konfigurieren

    ifconfig pflow0 flowsrc 127.0.0.1 flowdst 127.0.0.1:3001

oder in /etc/hostname.pflow0 (ungetestet)

    cat hostname.pflow0
    flowsrc 127.0.0.1 flowdst 127.0.0.1:3001

und in der pf.conf

    set state-defaults pflow

Obacht, nur Regeln die mit state pflow versehen sind, werden wirklich zum Collector geschickt. Mehr dazu hier

[Google Books](https://books.google.de/books?id=oSQBBQAAQBAJ&pg=PA177&lpg=PA177&dq=openbsd+pf+pflow+setup&source=bl&ots=cBYa2yah88&sig=MPPwQWcx3tFpVgXDgD6Uh4O-_sI&hl=de&sa=X&ei=ILnnVOnrOYqqPNucgIAO&ved=0CGUQ6AEwCA#v=onepage&q=openbsd%20pf%20pflow%20setup&f=false)

#### Collector installieren

    $ pkg_add flowd
    $ cat /etc/flowd.conf | grep -v ^#
    logfile "/var/log/flowd"
    listen on 127.0.0.1:3001
    listen on [::1]:3001
    flow source 127.0.0.1
    store ALL
    discard all
    accept agent 127.0.0.1

#### Anschauen

    $ tcpdump -nettti pflow0
    $ tcpdump -nettti pflog0 host 192.168.1.19 and port 9995
