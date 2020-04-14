---
title: DRBD Setup
date: 2012-10-19T10:13:27
tags:
- Filesystems
- DRBD
---

### Netzwerk Interface  für die Konfiguration

```
allow-hotplug eth5
auto eth5
iface eth5 inet static
        address 10.0.0.6
        netmask 255.255.255.252
        network 10.0.0.4
```

### Partition oder ganzes Device erstellen

```
/dev/sdb1
/dev/sdc
```

### Installation

``` bash
aptitude install drbd-utils
```

### DRBD Config

/etc/drbd.conf auf beiden nodes:

```
global {
    usage-count no;
}

common {
  syncer { rate 100M; }
}

resource db {
  protocol      C;

  startup { wfc-timeout 0; degr-wfc-timeout     120; }
  disk { on-io-error detach; }
  net {
        after-sb-0pri discard-older-primary;
        after-sb-1pri discard-secondary;
        after-sb-2pri disconnect;
        }

  syncer {
  }
  on db01 {
    device      /dev/drbd0;     ## gemeinsames DRBD-Device
    disk        /dev/sdb;       ## Partition die genutzt werden soll
    address     10.0.0.5:7791;  ## jeweilige lokale IP von eth2 auf db1
    meta-disk   internal;
  }
  on db02 {
    device      /dev/drbd0;     ## gemeinsames DRBD-Device
    disk        /dev/sdb;       ## Partition die genutzt werden soll
    address     10.0.0.6:7791;  ## jeweilige lokale IP von eth2 auf db2
    meta-disk   internal;
  }
}
```

### Initialisierung

```
drbdadm create-md db
oder
drbdadm create-md

und

/etc/init.d/drbd start
```

### Erster Sync

```
$ drbdadm -- --overwrite-data-of-peer primary
$ cat /proc/drbd
version: 8.3.11 (api:88/proto:86-96)
srcversion: 71955441799F513ACA6DA60
 0: cs:SyncSource ro:Primary/Secondary ds:UpToDate/Inconsistent C r-----
    ns:748544 nr:0 dw:0 dr:749208 al:0 bm:45 lo:0 pe:0 ua:0 ap:0 ep:1 wo:f oos:194018468
        [>....................] sync'ed:  0.4% (189468/190200)Mfinish: 5:14:06 speed: 10,284 (10,252) K/sec
```

### Dateisystem einrichten

    mkfs.ext4 /dev/drbd0
    mount /dev/drbd0 /mnt/
    cd /mnt
    touch 1 2 3 4 5
    umount /mnt

### Schwenken

    node1: drbdadm secondary db
    node2: drbdadm primary db
    node2: cat /proc/drbd && mounten

### MTU Problem

weil ping immer mtu 1500 macht, findet man das nicht.

```
[72358.666485] block drbd0: [drbd0_worker/23957] sock_sendmsg time expired, ko = 4294967292
[72364.665601] block drbd0: [drbd0_worker/23957] sock_sendmsg time expired, ko = 4294967291
[72370.664668] block drbd0: [drbd0_worker/23957] sock_sendmsg time expired, ko = 4294967290
[72376.663727] block drbd0: [drbd0_worker/23957] sock_sendmsg time expired, ko = 4294967289
[72382.662747] block drbd0: [drbd0_worker/23957] sock_sendmsg time expired, ko = 4294967288
[72388.661861] block drbd0: [drbd0_worker/23957] sock_sendmsg time expired, ko = 4294967287
[72394.660934] block drbd0: [drbd0_worker/23957] sock_sendmsg time expired, ko = 4294967286

allow-hotplug eth5
auto eth5
iface eth5 inet static
        address 10.0.0.6
        netmask 255.255.255.252
        network 10.0.0.4
        up ip link set eth5 mtu 1500 ## Right
        up ip link set eth5 mtu 9000 ## Wrong!

```

## Hearbeat

### Installation Heartbeat

``` bash
aptitude install heartbeat
```

### Konfiguration

Insgesamt nur 3 Files platzieren

#### /etc/ha.d/ha.cf:

```
logfacility     daemon          ## This is deprecated
keepalive       1               ## Interval between heartbeat (HB) packets.
deadtime        10              ## How quickly HB determines a dead node.
warntime        5               ## Time HB will issue a late HB.
initdead        120             ## Time delay needed by HB to report a dead node.
udpport         694             ## UDP port HB uses to communicate between nodes.
bcast           eth0            ## Which interface to use for HB packets.
auto_failback   off             ## Auto promotion of primary node upon return to cluster.
node            host1px         ## Node name 1.
node            host2px         ## Node name 2.

respawn hacluster /usr/lib/heartbeat/ipfail
## Specifies which programs to run at startup

use_logd        yes                             ## Use system logging.
logfile         /var/log/hb.log                 ## Heartbeat logfile.
debugfile       /var/log/heartbeat-debug.log    ## Debugging logfile.
```

#### /etc/ha.d/authkeys

```
auth 2
2 crc
```

#### /etc/ha.d/haresources

```
host1px IPaddr::172.19.15.209/27 drbddisk::db Filesystem::/dev/drbd0::/data/mysql::ext3::defaults mysql
```
