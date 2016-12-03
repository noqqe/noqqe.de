---
title: OpenBSD Keyboard / Timesettings
date: 2012-11-03T17:51:55
tags: 
- OS
- OpenBSD
---

## Timezone

    ln -fs /usr/share/zoneinfo/Europe/Berlin /etc/localtime

## Umlauts Encoding Charset UTF8

    cat /etc/wsconsctl.conf
    keyboard.encoding=de

und in .bash_profile und .profile

    export LC_CTYPE="en_US.UTF-8"
    export LC_MESSAGES="en_US.UTF-8"
