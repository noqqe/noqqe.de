---
title: Linux Networking
date: 2021-03-29T14:24:19
tags:
- OS
- Linux
---

Ein paar kleine Kniffe und Linux Networking Foo.

<!--more-->

## Networking mit `ip`

IP setzen

    ip address add 192.168.1.200/24 dev eth0

Gateway setzen

    ip route add 0.0.0.0/0 via 192.168.1.1 dev eth0


## Networking mit /etc/network/interfaces

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

## Networking mit netplan

DHCP

```yaml
network:
  version: 2
  renderer: networkd
  ethernets:
    enp0s3:
      dhcp4: yes
```

Static IP

```yaml
network:
  version: 2
  renderer: networkd
  ethernets:
    enp0s3:
     dhcp4: no
     addresses: [192.168.1.222/24]
     gateway4: 192.168.1.1
     nameservers:
       addresses: [8.8.8.8,8.8.4.4]

```

## Offene Verbindungen anzeigen

Es gibt mehrere Wege

    lsof -i tcp
    netstat -tap
    netstat -tulpan
    ss -tulapn


## Löschen eines Interface

    ip link delete docker0

## Löschen von Bridge Interfaces

Mehrere Möglichkeiten

    ip link set br0 down
    brctl delbr br0

oder

     ip link delete br0 type bridge

## Einzelnes Programm

Um zu verfolgen welchen Traffic ein einzelnes Programm erzeugt lässt sich gut
mit `strace` nachweisen.

    strace -p <pid> -f -e connect

