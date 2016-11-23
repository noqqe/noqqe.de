---
title: rsync
date: 2012-01-12T13:20:11.000000
tags: 
- Software
- rsync
---


#### System migration

Hier ein kleiner Hint für eine System migration.

    rsync -avz --exclude=/proc --exclude=/tmp --exclude=/sys --exclude=/dev / root@host:/mnt/root

oder

    rsync -axHSKDvz --exclude=/proc --exclude=/tmp --exclude=/sys --exclude=/dev -e ssh / root@host:/mnt/root

