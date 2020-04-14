---
date: 2017-05-18T08:27:26+02:00
draft: false
title: NUMA & numactl
tags:
- OS
---

NUMA ist sozusagen Memory und CPU Caching auf Hardware Ebene. Das kann für
Produktionsdatenbanken (MySQL und MongoDB) ungünstig sein. Deswegen
schaltet man das aus.

## Status

Herausfinden ob NUMA aktiviert ist, lässt sich über den Kernel.

``` bash
$ dmesg | grep -i numa
[    0.000000] No NUMA configuration found
```

Manchmal ist die Meldung aber zu sehr am Anfang, weswegen man lieber im Log
nachsieht.

``` bash
zgrep -i numa /var/log/messages*
kernel: [    0.000000] NUMA: Node 0 [mem 0x00000000-0x0009ffff] + [mem 0x00100000-0x0fffffff] -> [mem 0x00000000-0x0fffffff]
```

So sieht das aus wenn NUMA aktiviert ist.

## Konfiguration

Es gibt eine Tool das NUMA Einstellungen pro Prozess oder global setzen
kann. Die nachfolgende Option verteilt den Memory Alloc gleichmässig auf
alle Nodes.

```
numactl --interleave=all <path> <options>
```

Auf Platte persistieren kann man das deaktivierte NUMA auch via

```
sudo sysctl -w vm.zone_reclaim_mode=0
```

## Details

Anzeigen der NUMA Statistiken (Cache Hits..)

``` bash
cat /sys/devices/system/node/node*/numastat
```

## Links

* [numactl](https://linux.die.net/man/8/numactl)
* [acm](https://queue.acm.org/detail.cfm?id=2513149)