---
title: tmux
date: 2013-06-29T12:18:46
tags: 
- Software
- tmux
---

Reattach:

    alias re="exec tmux attach -d"

Reload tmux without restarting

    C-a :source-file /absolute/path/to/your/.tmux.conf

oder innerhalb der session:

    tmux source-file /absolute/path/to/your/.tmux.conf

Automatisch weechat starten

    @reboot tmux -u new-session -s base -d -n irc '/bin/bash -lc "weechat-curses"'

