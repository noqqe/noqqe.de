---
title: Linux Networking
date: 2011-06-07T15:42:20
tags:
- OS
- Linux
---

### Offene Verbindungen anzeigen

Es gibt mehrere Wege

    lsof -i tcp
    netstat -tap
    netstat -tulpan
    ss -tulapn

### Network Interfaces konfigurieren

Weil ich es einfach immer wieder vergesse.

    ## The primary network interface
    #allow-hotplug eth1
    auto eth1
    iface eth1 inet static
        address 10.42.1.104
        netmask 255.255.0.0
        gateway 10.42.10.1

    resolv.conf
    nameserver 10.42.10.50
    nameserver 217.145.99.9


### Löschen eines Interface

    ip link delete docker0

### Einzelnes Programm

Um zu verfolgen welchen Traffic ein einzelnes Programm erzeugt lässt sich gut
mit `strace` nachweisen.

    strace -p <pid> -f -e connect

