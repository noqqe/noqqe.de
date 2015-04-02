---
date: 2008-12-14T14:33:07+02:00
type: post
slug: flashspeicher-ist-auch-nichtmehr-das-was-er-mal-war
comments: true
title: FlashSpeicher ist auch nichtmehr das was er mal war.
aliases:
- /archives/359
categories:
- Hardware
- Linux
tags:
- Error
- Flash
- MemoryStick
- ProDuo
- PSP
- Sony
---

Meine Tolle 8GB MemoryStickProDuo Magicgate Mark2 Karte (kurzer treffender Name mh?...) die mir gestern geliefert wurde sieht ziemlich im Arsch aus.

Mounten klappt alles. Solang bis ich daten schreibe. Dann:

[  263.799077] sd 6:0:0:0: Attached scsi generic sg8 type 0
[ 3648.888900] sd 4:0:0:3: [sdf] 15900672 512-byte hardware sectors (8141 MB)
[ 3648.890024] sd 4:0:0:3: [sdf] Write Protect is off
[ 3648.890027] sd 4:0:0:3: [sdf] Mode Sense: 03 00 00 00
[ 3648.890030] sd 4:0:0:3: [sdf] Assuming drive cache: write through
[ 3648.891146] sd 4:0:0:3: [sdf] 15900672 512-byte hardware sectors (8141 MB)
[ 3648.892019] sd 4:0:0:3: [sdf] Write Protect is off
[ 3648.892022] sd 4:0:0:3: [sdf] Mode Sense: 03 00 00 00
[ 3648.892025] sd 4:0:0:3: [sdf] Assuming drive cache: write through
[ 3648.892029]  sdf: sdf1 # Bis hierher ist Mounten....Schreibvorgang folgt :
[ 3758.531417] sd 4:0:0:3: [sdf] Device not ready: Sense Key : Not Ready [current]
[ 3758.531425] sd 4:0:0:3: [sdf] Device not ready: Add. Sense: Medium not present
[ 3758.531432] end_request: I/O error, dev sdf, sector 6272
[ 3759.606683] sd 4:0:0:3: [sdf] Device not ready: Sense Key : Not Ready [current]
[ 3759.606691] sd 4:0:0:3: [sdf] Device not ready: Add. Sense: Medium not present
[ 3759.606697] end_request: I/O error, dev sdf, sector 15872
[ 3789.675274] usb 5-8: reset high speed USB device using ehci_hcd and address 5
[ 3819.924931] usb 5-8: reset high speed USB device using ehci_hcd and address 5
[ 3850.167730] usb 5-8: reset high speed USB device using ehci_hcd and address 5
[ 3880.410527] usb 5-8: reset high speed USB device using ehci_hcd and address 5
[ 3910.653324] usb 5-8: reset high speed USB device using ehci_hcd and address 5
[ 3940.897116] usb 5-8: reset high speed USB device using ehci_hcd and address 5
[ 3941.077889] sd 4:0:0:3: [sdf] Result: hostbyte=DID_ABORT driverbyte=DRIVER_OK,SUGGEST_OK
[ 3941.077898] end_request: I/O error, dev sdf, sector 16352
[ 3941.083123] sd 4:0:0:4: timing out command, waited 180s
Defekte Sektoren und Input/Output Errors...

Hat jemand ne Idee? *zu Dominik schiel*
