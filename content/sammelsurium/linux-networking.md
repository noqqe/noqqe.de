---
title: Linux Networking
date: 2011-06-07T15:42:20
tags: 
- OS
- Linux
---

### Offene Verbindungen anzeigen

    lsof -i tcp
    netstat -tap

### Network Interfaces konfigurieren

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