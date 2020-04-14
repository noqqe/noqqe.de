---
title: "LVM Disks verkleinern"
date: 2018-04-20T16:48:42+02:00
tags:
- lvm
- shrink
- resize
---

Ich kam in die Verlegenheit ein Volume verkleinern zu müssen. LVM vergrößern,
klar, 100 mal gemacht. Aber verkleinern war immer so...ehh... lieber auf neue
Disk syncen und alte Disk wegwerfen. Da es diesmal aber Produktiver Service
ist und es sich um Millionen von Files handelt (Danke Graphite..) geht das
nicht.

## Testaufbau

Da es dabei um 5TB geht und wir noch kein Backup davon haben (daher ja der
ganze Aufwand) wollte ich das vorher mal ausprobiert haben.

```
# Init disks
pvcreate /dev/sdb
pvcreate /dev/sdc

# Create VG
vgcreate test /dev/sdb /dev/sdc
vcreate -L 20000M -n test test

# Init Filesystem and mount
mkfs.ext4 /dev/mapper/test-test
mount /dev/mapper/test-test /tmp/test

# Fill with demo data
cd /tmp/test
for x in {1..100} ; do dd if=/dev/zero of=$x.txt bs=1M count=500 ; done

# Remove some data
/dev/mapper/test-test      20G   20G     0 100% /tmp/test
rm 2* 3* 6* 7*
/dev/mapper/test-test      20G  8.4G  9.8G  47% /tmp/test
```

## Shrink

Diese besteht aus 2 Disks die nun beide gefüllt sind. In der Produntkion sind
das 4 Disks, aber das sollte zum Vergleichen erstmal reichen.

```
pvs
  PV         VG     Fmt  Attr PSize  PFree
  /dev/sdb   test   lvm2 a--  10.00g      0
  /dev/sdc   test   lvm2 a--  10.00g 472.00m
```

Um einen resize durchzuführen muss das FS erstmal umounted werden.

```
umount /dev/mapper/test-test
```

FS überprüfen (muss sein)

```
e2fsck -f /dev/mapper/test-test
e2fsck 1.43.4 (31-Jan-2017)
Pass 1: Checking inodes, blocks, and sizes
Pass 2: Checking directory structure
Pass 3: Checking directory connectivity
Pass 4: Checking reference counts
Pass 5: Checking group summary information
/dev/mapper/test-test: 89/1281120 files (0.0% non-contiguous), 2300477/5120000 blocks
```

Hier wird nun der erste Layer, das Filesystem, auf eine bestimmte größe
reduziert. Das funktioniert überraschend gut.

``` bash
resize2fs /dev/mapper/test-test 9G
```

Dann das Logical Volume auf dem das Filesystem liegt kleiner machen.

``` bash
lvreduce -L 9g /dev/mapper/test-test
```

Das nun kleinere FS wieder mounten.

``` bash
/dev/mapper/test-test  172M  8.4M  151M   6% /tmp/test
```

Die Daten die evtl auf der Platte liegen müssen aber noch migriert werden auf
die Platte die wir noch behalten wollen. Warum hier noch die alten Used Zahlen
(FS voll) angezeigt werden kann ich nur vermuten. Denke aber weil diese im FS
nur als gelöscht markiert worden sind und die eigene Disk darunter garnichts
vom löschen mitbekommen hat. Anzeigen kann ich mir das aber mit dem LVM Tool
`pvs`

```
pvs -o+pv_used
  /dev/sdb   test   lvm2 a--  10.00g      0  10.00g
  /dev/sdc   test   lvm2 a--  10.00g 472.00m  9.54g

```

Bevor ich eine Disk entfernen kann, muss ich die verbliebenen Daten noch auf
die Platte verschieben die übrig bleibt.

```
pvmove /dev/sdc /dev/sdb
```

Dann müsste mehr Platz im VG sein als vorher  und wir können eine Disk aus dem
VG rausziehen

```
vgreduce test /dev/sdc
```

Und fertig! Gerade der `lvreduce` ist unendlich hilfreich, im vergleich zu
früher wo man mit `fdisk` den Partition Table gelöscht und neu angelegt hat.
*grusel*

{{< figure src="/uploads/2018/04/hdd.jpg" >}}

Jetzt nur noch hoffen das es auch produktiv gut wird.
