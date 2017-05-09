---
title: LVM
date: 2011-11-28T14:42:21
tags:
- Filesystems
---

## Auslesen

Informationen von Volume anzeigen

    $ vgdisplay

Kurze Info über Free/Used des Volumes:

    $ vgdisplay -s

Alle Platten anzeigen

    $ lvscan
    ACTIVE            '/dev/toshi/root' [6,96 GB] inherit
    ACTIVE            '/dev/toshi/swap_1' [616,00 MB] inherit
    ACTIVE            '/dev/toshi/home' [10,82 GB] inherit

Alle LV und Typen listen

    blkid
    lvs

## Volume anlegen

Erstellen, formatieren und mounten.

    lvcreate -n srv -L 50G lvm
    mkfs.ext4 /dev/lvm/srv
    mount /dev/lvm/srv /media/newsrv
    rsync -avP /srv/* /media/newsrv/

## Volume Resize

    $ lvresize /dev/toshi/root --size +2G
    ## oder auch mit -L +10G
    > Extending logical volume root to 6,96 GB
    > Logical volume root successfully resized
    $ resize2fs /dev/mapper/root
    $ resize2fs /dev/system/root


## Volume Group Rename

    vgrename gnar1 gnar2
    dpkg-reconfigure linux-image-2.6.


## Rescan des SCSI-Hostbuses

    ls /sys/class/scsi_host/
    host0  host1  host2

    echo "- - -" > /sys/class/scsi_host/host0/scan
    echo "- - -" > /sys/class/scsi_host/host1/scan
    echo "- - -" > /sys/class/scsi_host/host2/scan

    lsscsi
    [1:0:0:0]    cd/dvd  NECVMWar VMware IDE CDR10 1.00  /dev/sr0
    [2:0:0:0]    disk    VMware   Virtual disk     1.0   /dev/sda
    [2:0:1:0]    disk    VMware   Virtual disk     1.0   /dev/sdb

