---
title: OpenBSD IPv6
date: 2013-06-14T11:56:41
tags: 
- OS
- OpenBSD
---

Interfaces

    /etc/hostname.em0
    inet 213.95.21.200 255.255.255.0 NONE
    inet6 alias 2001:780:3:5::122 64
    inet6 alias 2001:780:132::1 48

Defaultgateway

    cat /etc/mygate
    213.95.21.1
    2001:780:3:5::1

Neustart

    sh /etc/netstart em0
