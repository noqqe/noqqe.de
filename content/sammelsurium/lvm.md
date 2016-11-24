---
title: LVM
date: 2011-11-28T14:42:21.000000
tags: 
- Filesystems
---


Informationen von Volume anzeigen

    $ vgdisplay

Kurze Info über Free/Used des Volumes:

    $ vgdisplay -s

Alle Platten anzeigen

    $ lvscan
    ACTIVE            '/dev/toshi/root' [6,96 GB] inherit
    ACTIVE            '/dev/toshi/swap_1' [616,00 MB] inherit
    ACTIVE            '/dev/toshi/home' [10,82 GB] inherit

### Volume anlegen

Erstellen, formatieren und mounten.

    lvcreate -n srv -L 50G lvm
    mkfs.ext4 /dev/lvm/srv
    mount /dev/lvm/srv /media/newsrv
    rsync -avP /srv/* /media/newsrv/

### Volume Resize

    $ lvresize /dev/toshi/root --size +2G
    ## oder auch mit -L +10G
    > Extending logical volume root to 6,96 GB
    > Logical volume root successfully resized
    $ resize2fs /dev/mapper/root
    $ resize2fs /dev/system/root

### Alle LV und Typen listen

    blkid
    lvs

### Volumegroup rename

    vgrename gnar1 gnar2
    dpkg-reconfigure linux-image-2.6.

### Volles Howto

Erweiterung der VM im VCenter um +100G Speicherplatz auf na5-vmstor1

#### Rescan des SCSI-Hostbuses

    ls /sys/class/scsi_host/
    host0  host1  host2

    echo "- - -" > /sys/class/scsi_host/host0/scan
    echo "- - -" > /sys/class/scsi_host/host1/scan
    echo "- - -" > /sys/class/scsi_host/host2/scan

    lsscsi
    [1:0:0:0]    cd/dvd  NECVMWar VMware IDE CDR10 1.00  /dev/sr0
    [2:0:0:0]    disk    VMware   Virtual disk     1.0   /dev/sda
    [2:0:1:0]    disk    VMware   Virtual disk     1.0   /dev/sdb

#### PV erstellen und VG erweitern:

    ## pvcreate /dev/sdb
    Physical volume "/dev/sdb" successfully created
    ## vgextend vg01 /dev/sdb
    Volume group "vg01" successfully extended

#### LVM erweitern

    ## lvextend -L +50G /dev/mapper/vg01-data
    Extending logical volume data to 89,00 GiB
    Logical volume data successfully resized
    ## resize2fs /dev/mapper/vg01-data
    resize2fs 1.41.12 (17-May-2010)
    Filesystem at /dev/mapper/vg01-data is mounted on /data; on-line resizing
    required
    old desc_blocks = 3, new_desc_blocks = 6
    Performing an on-line resize of /dev/mapper/vg01-data to 23330816 (4k) blocks.
    The filesystem on /dev/mapper/vg01-data is now 23330816 blocks long.

