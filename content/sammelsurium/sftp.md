---
title: sftp
date: 2013-01-10T16:28:46
tags:
- Software
- sftp
---

## Chroot

~~~
#Subsystem      sftp    /usr/lib/ssh/sftp-server
Subsystem       sftp    internal-sftp

## This enables accepting locale enviroment variables LC_* LANG, see sshd_config(5).
AcceptEnv LANG LC_CTYPE LC_NUMERIC LC_TIME LC_COLLATE LC_MONETARY LC_MESSAGES
AcceptEnv LC_PAPER LC_NAME LC_ADDRESS LC_TELEPHONE LC_MEASUREMENT
AcceptEnv LC_IDENTIFICATION LC_ALL

Match group sftpacc
    ChrootDirectory %h
    X11Forwarding no
    AllowTcpForwarding no
    ForceCommand internal-sftp
~~~

Startverzeichnis DARF NICHT für die User schreibbar sein. sprich das
ChrootDirectory Schreiben darfst du immer erst eine Verzeichnisebene
tiefer, sonst lehnt sftp sofort die Verbindung ab.