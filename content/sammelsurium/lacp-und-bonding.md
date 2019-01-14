---
title: LACP und Bonding
date: 2012-02-14T10:07:31
tags: 
- LACP
- Network
- Bonding
- HowTos
---

## Beispielhafte Ausgabe für LACP gebondete Interfaces

bloeder bondage witz hier einfuegen.

    $ ifconfig
    bond0     Link encap:Ethernet  HWaddr 00:1e:68:4a:07:3c
              inet addr:10.1.1.50  Bcast:10.1.1.255  Mask:255.255.255.0
              inet6 addr: fe80::21e:68ff:fe4a:73c/64 Scope:Link
              UP BROADCAST RUNNING MASTER MULTICAST  MTU:1500  Metric:1
              RX packets:297092625 errors:0 dropped:0 overruns:0 frame:0
              TX packets:108449744 errors:0 dropped:0 overruns:0 carrier:0
              collisions:0 txqueuelen:0
              RX bytes:377310521891 (377.3 GB)  TX bytes:7322758046 (7.3 GB)

    eth0      Link encap:Ethernet  HWaddr 00:1e:68:4a:07:3c
              UP BROADCAST RUNNING SLAVE MULTICAST  MTU:1500  Metric:1
              RX packets:281017381 errors:0 dropped:0 overruns:0 frame:0
              TX packets:108438375 errors:0 dropped:0 overruns:0 carrier:0
              collisions:0 txqueuelen:1000
              RX bytes:373260450833 (373.2 GB)  TX bytes:7321348290 (7.3 GB)
              Memory:fc5e0000-fc600000

    eth1      Link encap:Ethernet  HWaddr 00:1e:68:4a:07:3c
              UP BROADCAST RUNNING SLAVE MULTICAST  MTU:1500  Metric:1
              RX packets:16075244 errors:0 dropped:0 overruns:0 frame:0
              TX packets:11369 errors:0 dropped:0 overruns:0 carrier:0
              collisions:0 txqueuelen:1000
              RX bytes:4050071058 (4.0 GB)  TX bytes:1409756 (1.4 MB)
              Memory:fc5c0000-fc5e0000

## Hilfreiche Befehle

    $ aptitude install ifenslave
    $ ifenslave bond0
    The result of SIOCGIFFLAGS on bond0 is 1443.
    The result of SIOCGIFADDR is 00.00.0a.01.
    The result of SIOCGIFHWADDR is type 1  00:1e:68:4a:07:3c.

### Status abfragen eines Bonding Interfaces

    $ cat /proc/net/bonding/bond0
    Ethernet Channel Bonding Driver: v3.5.0 (November 4, 2008)

    Bonding Mode: IEEE 802.3ad Dynamic link aggregation
    Transmit Hash Policy: layer2 (0)
    MII Status: up
    MII Polling Interval (ms): 100
    Up Delay (ms): 0
    Down Delay (ms): 0

    802.3ad info
    LACP rate: slow
    Aggregator selection policy (ad_select): stable
    Active Aggregator Info:
            Aggregator ID: 3
            Number of ports: 2
            Actor Key: 17
            Partner Key: 14
            Partner Mac Address: 78:fe:3d:41:b5:00

    Slave Interface: eth0
    MII Status: up
    Link Failure Count: 1
    Permanent HW addr: 00:1e:68:4a:07:3c
    Aggregator ID: 3

    Slave Interface: eth1
    MII Status: up
    Link Failure Count: 1
    Permanent HW addr: 00:1e:68:4a:07:3d
    Aggregator ID: 3

### Informationen zum genutzten Kernel Modul Bonding

    $ modinfo bonding
    filename:
    /lib/modules/2.6.32-27-server/kernel/drivers/net/bonding/bonding.ko
    author:         Thomas Davis, tadavis@lbl.gov and many others
    description:    Ethernet Channel Bonding Driver, v3.5.0
    version:        3.5.0
    license:        GPL
    srcversion:     E5E5261A12E39B584F85F66
    depends:
    vermagic:       2.6.32-27-server SMP mod_unload modversions
    parm:           max_bonds:Max number of bonded devices (int)
    parm:           num_grat_arp:Number of gratuitous ARP packets to send on
    failover event (int)
    parm:           num_unsol_na:Number of unsolicited IPv6 Neighbor Advertisements
    packets to send on failover event (int)
    parm:           miimon:Link check interval in milliseconds (int)
    parm:           updelay:Delay before considering link up, in milliseconds (int)
    parm:           downdelay:Delay before considering link down, in milliseconds
    (int)
    parm:           use_carrier:Use netif_carrier_ok (vs MII ioctls) in miimon; 0
    for off, 1 for on (default) (int)
    parm:           mode:Mode of operation : 0 for balance-rr, 1 for active-backup,
    2 for balance-xor, 3 for broadcast, 4 for 802.3ad, 5 for balance-tlb, 6 for
    balance-alb (charp)
    parm:           primary:Primary network device to use (charp)
    parm:           lacp_rate:LACPDU tx rate to request from 802.3ad partner
    (slow/fast) (charp)
    parm:           ad_select:803.ad aggregation selection logic: stable (0,
    default), bandwidth (1), count (2) (charp)
    parm:           xmit_hash_policy:XOR hashing method: 0 for layer 2 (default), 1
    for layer 3+4 (charp)
    parm:           arp_interval:arp interval in milliseconds (int)
    parm:           arp_ip_target:arp targets in n.n.n.n form (array of charp)
    parm:           arp_validate:validate src/dst of ARP probes: none (default),
    active, backup or all (charp)
    parm:           fail_over_mac:For active-backup, do not set all slaves to the
    same MAC.  none (default), active or follow (charp)

### Starten des bonding modules mit xmit_hash_policy

fuer layer 3+4 hash map

    $ rmmod bonding
    $ modprobe bonding xmit_hash_policy=1

### Wichtig ist der Mode des Modules Bonding

weil hier nur der Modus 4 der ist, der LACP macht (802.3ad)

    $ dmesg
    [530622.396615] bonding: bond0: enslaving eth1 as a backup interface with a down
    link.
    [530622.397442] bonding: bond0: setting mode to 802.3ad (4).
    [530622.397479] bonding: bond0: Setting MII monitoring interval to 100.

## Ein Beispiel mit iperf und aktiviertem layer3+4 xmit_hash_policy

    $ modinfo bonding
    xmit_hash_policy:XOR hashing method: 0 for layer 2 (default), 1 for layer 3+4
    (charp)

    $ rmmod bonding
    $ modprobe bonding xmit_hash_policy=1

    ## iperf -c 10.1.1.50 -n 1000M -P 6
    ------------------------------------------------------------
    Client connecting to 10.1.1.50, TCP port 5001
    TCP window size: 16.0 KByte (default)
    ------------------------------------------------------------
    [  8] local 10.1.1.12 port 44541 connected with 10.1.1.50 port 5001
    [  4] local 10.1.1.12 port 44536 connected with 10.1.1.50 port 5001
    [  3] local 10.1.1.12 port 44537 connected with 10.1.1.50 port 5001
    [  5] local 10.1.1.12 port 44538 connected with 10.1.1.50 port 5001
    [  6] local 10.1.1.12 port 44539 connected with 10.1.1.50 port 5001
    [  7] local 10.1.1.12 port 44540 connected with 10.1.1.50 port 5001
    [ ID] Interval       Transfer     Bandwidth
    [  7]  0.0-26.2 sec  1000 MBytes   321 Mbits/sec
    [  6]  0.0-26.5 sec  1000 MBytes   317 Mbits/sec
    [  4]  0.0-26.5 sec  1000 MBytes   317 Mbits/sec
    [  8]  0.0-26.6 sec  1000 MBytes   315 Mbits/sec
    [  5]  0.0-26.7 sec  1000 MBytes   314 Mbits/sec
    [  3]  0.0-26.7 sec  1000 MBytes   314 Mbits/sec
    [SUM]  0.0-26.7 sec  5.86 GBytes  1.88 Gbits/sec

    ## iperf -s
    ------------------------------------------------------------
    Server listening on TCP port 5001
    TCP window size: 85.3 KByte (default)
    ------------------------------------------------------------
    [  4] local 10.1.1.12 port 5001 connected with 10.1.1.50 port 36476
    [  5] local 10.1.1.12 port 5001 connected with 10.1.1.50 port 36478
    [  6] local 10.1.1.12 port 5001 connected with 10.1.1.50 port 36479
    [  7] local 10.1.1.12 port 5001 connected with 10.1.1.50 port 36480
    [  8] local 10.1.1.12 port 5001 connected with 10.1.1.50 port 36477
    [  9] local 10.1.1.12 port 5001 connected with 10.1.1.50 port 36481
    [ ID] Interval       Transfer     Bandwidth
    [  6]  0.0-26.6 sec  1000 MBytes   315 Mbits/sec
    [  4]  0.0-26.7 sec  1000 MBytes   314 Mbits/sec
    [  8]  0.0-26.7 sec  1000 MBytes   314 Mbits/sec
    [  5]  0.0-26.7 sec  1000 MBytes   314 Mbits/sec
    [  7]  0.0-26.7 sec  1000 MBytes   314 Mbits/sec
    [  9]  0.0-26.7 sec  1000 MBytes   314 Mbits/sec
    [SUM]  0.0-26.7 sec  5.86 GBytes  1.88 Gbits/sec

## Beispiel Konfiguration für bonding in /etc/network/interfaces #

    ## LACP
    auto bond0
    iface bond0 inet static
            address 10.1.1.12
            network 10.1.1.0
            netmask 255.255.255.0
            broadcast 10.1.1.255
            gateway 10.1.1.3
            dns-nameservers 10.1.1.1
            dns-search gd
            slaves eth0 eth1
            bond-mode 4
            bond-miimon 100
            xmit_hash_policy 1

    ## MGMT
    auto eth3
    iface eth3 inet static
            address 192.168.15.201
            netmask 255.255.255.0