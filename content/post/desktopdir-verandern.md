---
date: 2008-12-09T10:18:33+02:00
type: post
slug: desktopdir-verandern
comments: true
title: DesktopDir verändern
aliases:
- /archives/355
categories:
- Linux
tags:
- config
- desktop
- dir
- home
- Linux
- x.org
- X11
---

Sollte es jemand mal ähnlich gehen und nicht 45 min das Internet
durchforsten wollen :)

in  /home/user/.config/user-dirs.dirs sind alle Orte definiert ;)

Falls jemand also bock hat das home verzeichnis als Desktop zu haben:
Einfach wie ich (*depp*) /home/user/Desktop nach /dev/null verschieben oder
einfach oben definieren ;) (Tipp: ich würde letzteres wählen...)

Das ganze sieht dann ungefähr so aus :


    XDG_DESKTOP_DIR="$HOME/<span class="highlight">Desktop</span>"
    XDG_DOWNLOAD_DIR="$HOME/<span class="highlight">Desktop</span>"
    XDG_TEMPLATES_DIR="$HOME/Templates"
    XDG_PUBLICSHARE_DIR="$HOME/Public"
    XDG_DOCUMENTS_DIR="$HOME/Documents"
    XDG_MUSIC_DIR="$HOME/Music"
    XDG_PICTURES_DIR="$HOME/Pictures"
    XDG_VIDEOS_DIR="$HOME/Videos"


Aber naja ich sag das nur mal so ... *grrr*
