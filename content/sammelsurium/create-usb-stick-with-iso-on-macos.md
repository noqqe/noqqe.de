---
title: Create USB Stick with iso on macOS
date: 2016-10-28T11:44:46.234000
tags: 
- osx
- macOS
- mount
- dd
- iso
---


For openbsd boot in this example.

Look for device

    diskutil list

Unmount the disk

    sudo diskutil unmountDisk /dev/disk4

Write iso to raw device with `pv`. Progressbar ftw!

    pv -tpreb Downloads/install60.iso | sudo dd of=/dev/rdisk4 bs=1m
