---
title: OpenBSD Network Configuration
date: 2012-11-04T12:31:38
tags:
- OS
- OpenBSD
---

## Interfaces

#### DHCP

``` bash
$ cat /etc/hostname.em0
dhcp
```

### Static

``` bash
$ cat /etc/hostname.fxp0
inet 10.0.0.38 255.255.255.0 NONE
```

### Route

in /etc/mygate

### Änderungen übernehmen

das ist wohl das pendant zum networking restart aus Debian

``` bash
sh /etc/netstart
```

## DNS

### /etc/resolv.conf

```
search example.com
nameserver 125.2.3.4
nameserver 125.2.3.5
lookup file bind
```

### /etc/myname

hier steht der Hostname
