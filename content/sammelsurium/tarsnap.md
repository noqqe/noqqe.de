---
title: tarsnap
date: 2014-12-12T13:52:35
tags:
- Software
- tarsnap
---

Tarsnap ist eine Backupsoftware und sein Service

<!--more-->

### Register

    tarsnap-keygen --keyfile /root/tarsnap.key --user wa1@noqqe.de --machine noc

### Config

    cachedir /var/cache/tarsnap/
    keyfile /etc/tarsnap.key
    humanize-numbers
    print-stats
    totals
    normalmem

## Useful Info

    tarsnap --print-stats --configfile /etc/tarsnap.conf
    tarsnap --list-archives --configfile /etc/tarsnap.conf

## Files anzeigen in einem dump

    tarsnap --configfile /etc/tarsnap.conf -tf vim2

Nach bestimmten File suchen mit Liste ausgeben

```bash
for x in $(tarsnap --list-archives | sort | grep home) ; do
  echo $x ; tarsnap -tf $x ;
done
```

## Dump an aktueller stelle auspacken

    tarsnap --configfile /etc/tarsnap.conf -xf vim2

## Testbackup

    tarsnap -c -f vim --configfile /etc/tarsnap.conf /home/noqqe/.vim/

## Cleanup

Alle Backups löschen

```bash
for x in $(tarsnap --list-archives); do
  tarsnap -d -f $x
done
```
