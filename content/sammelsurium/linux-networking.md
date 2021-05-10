---
title: Linux Networking
date: 2021-03-29T14:24:19
tags:
- OS
- Linux
---

Ein paar kleine Kniffe und Linux Networking Foo.

<!--more-->

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

DNS in `resolv.conf`

    nameserver 10.42.10.50
    nameserver 217.145.99.9


### Löschen eines Interface

    ip link delete docker0

### Löschen von Bridge Interfaces

Mehrere Möglichkeiten

    ip link set br0 down
    brctl delbr br0

oder

     ip link delete br0 type bridge

### Einzelnes Programm

Um zu verfolgen welchen Traffic ein einzelnes Programm erzeugt lässt sich gut
mit `strace` nachweisen.

    strace -p <pid> -f -e connect

