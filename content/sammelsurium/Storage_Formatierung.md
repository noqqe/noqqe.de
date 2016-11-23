---
title: Storage Formatierung
date: 2013-02-08T20:52:40.000000
tags: 
- HowTos
---


Öfters mal gesehen

       Device Boot      Start         End      Blocks   Id  System
    /dev/sdb1              63  3907029167  1953514552+  83  Linux
    Partition 1 does not start on physical sector boundary.

mit parted kann man das dann richten:

Falsch:

    (parted) mkpart primary 32.3kb 2000GB
    Warning: The resulting partition is not properly aligned for best performance.

Richtig:
Bei Blocksize von 512bytes * 2048 Blöcke:

    echo $((512*2048))
    select /dev/sdd
    rm 1
    mkpart primary 1048576b 2000GB

Wenns fertig ist sieht das so aus:

~~~

## fdisk -l

Disk /dev/sda: 2000.4 GB, 2000398934016 bytes
255 heads, 63 sectors/track, 243201 cylinders, total 3907029168 sectors
Units = sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disk identifier: 0x00071476

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1            2048  3907028991  1953513472   83  Linux

Disk /dev/sdb: 2000.4 GB, 2000398934016 bytes
255 heads, 63 sectors/track, 243201 cylinders, total 3907029168 sectors
Units = sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disk identifier: 0x000d37ee

   Device Boot      Start         End      Blocks   Id  System
/dev/sdb1            2048  3907028991  1953513472   83  Linux

Disk /dev/sdd: 2000.4 GB, 2000398934016 bytes
255 heads, 63 sectors/track, 243201 cylinders, total 3907029168 sectors
Units = sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disk identifier: 0x000e5ff6

   Device Boot      Start         End      Blocks   Id  System
/dev/sdd1            2048  3907028991  1953513472   83  Linux

Disk /dev/sdc: 2000.4 GB, 2000398934016 bytes
255 heads, 63 sectors/track, 243201 cylinders, total 3907029168 sectors
Units = sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disk identifier: 0x000f1151

   Device Boot      Start         End      Blocks   Id  System
/dev/sdc1            2048  3907028991  1953513472   83  Linux

Disk /dev/sde: 250.1 GB, 250059350016 bytes
255 heads, 63 sectors/track, 30401 cylinders, total 488397168 sectors
Units = sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disk identifier: 0x000eac13

~~~

