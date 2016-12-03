---
title: cntlm
date: 2015-02-16T14:03:05
tags: 
- Software
- cntlm
---

reverse proxy for authentication

#### starten

unter homebrew

    sudo launchctl unload /Library/LaunchDaemons/homebrew.mxcl.cntlm.plist
    sudo launchctl load /Library/LaunchDaemons/homebrew.mxcl.cntlm.plist

#### autostart

Damit sich das ding automatisch startet

    sudo cp -fv /usr/local/opt/cntlm/*.plist /Library/LaunchDaemons
    sudo chown root /Library/LaunchDaemons/homebrew.mxcl.cntlm.plist

#### Authoxy wieder deinstallieren

Gefunden auf https://paikialog.wordpress.com/2012/05/23/uninstall-authoxy/

    $ ps aux | grep cnt
    root           11328   0.0  0.0  2451268    376   ??  Ss    1:59PM 0:00.00 /usr/local/opt/cntlm/bin/cntlm
    $ sudo rm -rf /Library/PreferencePanes/Authoxy.prefPane/
    $ sudo rm -rf /Applications/startAuthoxy.app
    $ sudo rm -rf ~/Library/Preferences/net.hrsoftworks.AuthoxyPref.plist

Dieses authoxy geschmarre
