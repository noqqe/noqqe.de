---
title: ZFS
date: 2012-01-27T15:42:15
tags:
- Filesystems
---

# Generelle Befehle

Status ausgeben

    $ zpool status

Alle pools anzeigen

    $ zpool list

# Platten ersetzen

## Platten anzeigen

    $ zpool status | less

    NAME                STATE     READ WRITE CKSUM
      data                DEGRADED     0     0     0
        raidz1-0          ONLINE       0     0     0
          c4t0d0          ONLINE       0     0     0
          c4t1d0          ONLINE       0     0     0
          c4t2d0          ONLINE       0     0     0
          c4t3d0          ONLINE       0     0     0
          c4t4d0          ONLINE       0     0     0
          c4t5d0          ONLINE       0     0     0
          c4t6d0          ONLINE       0     0     0
        raidz1-1          DEGRADED     0     0     0
          c4t7d0          ONLINE       0     0     0
          c5t0d0          ONLINE       0     0     0
          c5t1d0          ONLINE       0     0     0
          c5t2d0          ONLINE       0     0     0
          spare-4         DEGRADED     0     0     0
              c5t3d0s0  FAULTED      0     0     0  too many errors
            c3t6d0        ONLINE       0     0     0
          c5t4d0          ONLINE       0     0     0
          c5t5d0          ONLINE       0     0     0
        raidz1-2          DEGRADED     0     0     0
          c5t6d0          ONLINE       0     0     0
          c5t7d0          ONLINE       0     0     0
          c6t0d0          ONLINE       0     0     0
          c6t1d0          ONLINE       0     0     0
          c6t2d0          ONLINE       0     0     0
          c6t3d0          ONLINE       0     0     0
          spare-6         DEGRADED     0     0     0
              c6t4d0s0/o  FAULTED      0     0     0  corrupted data
            c3t7d0        ONLINE       0     0     0
        raidz1-3          ONLINE       0     0     0
          c6t5d0          ONLINE       0     0     0
          c6t6d0          ONLINE       0     0     0
          c6t7d0          ONLINE       0     0     0

### Platten replacen

    zpool replace data c5t3d0
    zpool replace data c6t4d0

### Platten werden resilvered

        NAME                STATE     READ WRITE CKSUM
        data                DEGRADED     0     0     0
          raidz1-0          ONLINE       0     0     0
            c4t0d0          ONLINE       0     0     0
            c4t1d0          ONLINE       0     0     0
            c4t2d0          ONLINE       0     0     0
            c4t3d0          ONLINE       0     0     0
            c4t4d0          ONLINE       0     0     0
            c4t5d0          ONLINE       0     0     0
            c4t6d0          ONLINE       0     0     0
          raidz1-1          DEGRADED     0     0     0
            c4t7d0          ONLINE       0     0     0
            c5t0d0          ONLINE       0     0     0
            c5t1d0          ONLINE       0     0     0
            c5t2d0          ONLINE       0     0     0
            spare-4         DEGRADED     0     0     0
              replacing-0   DEGRADED     0     0     0
                c5t3d0s0/o  FAULTED      0     0     0  too many errors
                c5t3d0      ONLINE       0     0     0  8,26G resilvered
              c3t6d0        ONLINE       0     0     0
            c5t4d0          ONLINE       0     0     0
            c5t5d0          ONLINE       0     0     0
          raidz1-2          DEGRADED     0     0     0
            c5t6d0          ONLINE       0     0     0
            c5t7d0          ONLINE       0     0     0
            c6t0d0          ONLINE       0     0     0
            c6t1d0          ONLINE       0     0     0
            c6t2d0          ONLINE       0     0     0
            c6t3d0          ONLINE       0     0     0
            spare-6         DEGRADED     0     0     0
              replacing-0   DEGRADED     0     0     0
                c6t4d0s0/o  FAULTED      0     0     0  corrupted data
                c6t4d0      ONLINE       0     0     0  8,28G resilvered
              c3t7d0        ONLINE       0     0     0
          raidz1-3          ONLINE       0     0     0
            c6t5d0          ONLINE       0     0     0
            c6t6d0          ONLINE       0     0     0
            c6t7d0          ONLINE       0     0     0

### Platten sind resilvered

      pool: data
     state: ONLINE
     scrub: resilver completed after 5h52m with 0 errors on Wed Feb 29 19:22:58 2012
    config:

        NAME        STATE     READ WRITE CKSUM
        data        ONLINE       0     0     0
          raidz1-0  ONLINE       0     0     0
            c4t0d0  ONLINE       0     0     0
            c4t1d0  ONLINE       0     0     0
            c4t2d0  ONLINE       0     0     0
            c4t3d0  ONLINE       0     0     0
            c4t4d0  ONLINE       0     0     0
            c4t5d0  ONLINE       0     0     0
            c4t6d0  ONLINE       0     0     0
          raidz1-1  ONLINE       0     0     0
            c4t7d0  ONLINE       0     0     0
            c5t0d0  ONLINE       0     0     0
            c5t1d0  ONLINE       0     0     0
            c5t2d0  ONLINE       0     0     0
            c5t3d0  ONLINE       0     0     0  85,0G resilvered
            c5t4d0  ONLINE       0     0     0
            c5t5d0  ONLINE       0     0     0
          raidz1-2  ONLINE       0     0     0
            c5t6d0  ONLINE       0     0     0
            c5t7d0  ONLINE       0     0     0
            c6t0d0  ONLINE       0     0     0
            c6t1d0  ONLINE       0     0     0
            c6t2d0  ONLINE       0     0     0
            c6t3d0  ONLINE       0     0     0
            c6t4d0  ONLINE       0     0     0  170G resilvered
          raidz1-3  ONLINE       0     0     0
            c6t5d0  ONLINE       0     0     0
            c6t6d0  ONLINE       0     0     0
            c6t7d0  ONLINE       0     0     0
            c7t0d0  ONLINE       0     0     0
            c7t1d0  ONLINE       0     0     0
            c7t2d0  ONLINE       0     0     0
            c7t3d0  ONLINE       0     0     0
          raidz1-4  ONLINE       0     0     0
            c7t4d0  ONLINE       0     0     0