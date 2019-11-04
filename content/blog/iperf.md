---
title: iperf
date: 2019-02-01T08:17:55
tags:
- Software
- iperf
---

`iperf` ist ein Tool mit dem man Netzwerk Geschwindikeiten testen kann.

{{< figure src="/uploads/2019/01/iperf.png" >}}

Dazu braucht man 2 Linux Boxen. Einen Server:

```
$ iperf -s -p 4001
------------------------------------------------------------
Server listening on TCP port 4001
TCP window size: 85.3 KByte (default)
------------------------------------------------------------
```

und einen Client

```
$ iperf -c 172.27.37.26 -n 1000M -p 4001
------------------------------------------------------------
Client connecting to 172.27.37.26, TCP port 4001
TCP window size: 45.0 KByte (default)
------------------------------------------------------------
[  3] local 172.26.176.10 port 36114 connected with 172.27.37.26 port 4001
[ ID] Interval       Transfer     Bandwidth
[  3]  0.0-18.4 sec  1000 MBytes   455 Mbits/sec
```

der Client bestimmt dabei wie viel Daten mit welcher Geschwindigkeit
übertragen wird. Dabei kann es zwischen UDP und TCP entscheiden bzw auch
parallel messen, siehe hierzu `-P` oder bidirektional `-d`.
