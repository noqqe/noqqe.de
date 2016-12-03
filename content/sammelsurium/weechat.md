---
title: weechat
date: 2015-02-23T10:33:18
tags: 
- Software
- weechat
---

#### Quickies

* switch through Buffers in multi displays

    CTRL + X

* Show weechat internal tags

    /debug tags

#### Filters

ignore modes for a specific server

    /filter add modes irc.slack.* irc_mode *

delete a filter

    /filter del modes

#### Create new Server

    /server add slack irc.example.com/6667 -autoconnect -ssl -ssl_dhkey_size=512 -password=xxx -username=user -nicks=user
    /set irc.server.slack.ssl_verify off

#### Alias

    /alias hidebars /bar hide nicklist ; /bar hide buffers
    /alias showbars /bar show nicklist ; /bar show buffers

#### FISH Encryption

    /blowkey set user2 pass
    /blowkey set -server freenode #CHAN PASS

oder einfach Keyless

    /blowkey exchange user
