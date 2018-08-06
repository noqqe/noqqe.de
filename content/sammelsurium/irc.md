---
title: IRC
date: 2014-08-08T13:29:49
tags:
- Software
- irc
---

## OP

Channel Operator werden

    /msg chanserv OP #chan

OP wieder wegmachen

    /deop

Auto OP via ChanServ

    /query ChanServ
    aop #channel add <username>

## Anzahl der User in Channel

    /list #k4cg
    /names #k4cg

Das geht aber (auf freenode) nur ohne den Mode "s" (secure)

## Channel Infos

    /mode #channel

Wer ist OP

    /msg ChanServ ACCESS #channel LIST

## User Infos

Generelle Infos

    /whois <user>

Überprüfen ob User registriert ist


    /msg NickServ ACC <username>

     0 - account or user does not exist
     1 - account exists but user is not logged in
     2 - user is not logged in but recognized (see ACCESS)
     3 - user is logged in


## Topic setzen

    /topic foo bar

