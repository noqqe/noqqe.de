---
date: 2010-08-15T12:44:27+02:00
layout: post
slug: bash-nerdtool-config-for-macosx
status: publish
comments: true
title: Bash | NerdTool Config for MacOSX
alias: /archives/1182
categories:
- Coding
- Mac
- PlanetenBlogger
tags:
- bash
- cpu
- daemons
- established
- geektool
- gist
- github
- hostname
- listen
- memory
- nerdtool
- Network
- processes
- ps aux
- ram
- script
- skript
- uptime
---

[GeekTool](http://projects.tynsoe.org/en/geektool/) und [NerdTool](http://www.macupdate.com/info.php/id/31909/nerdtool) für MacOSX sind schöne Programme, welche Ausgaben von Bash-Scripten auf den Desktop ausgeben und ständig aktualiseren. Kein Geheimnis und nichts Neues.
Das Netz ist voll von schönen Spielereien für diese Tools. Hier mein Setup. (Benutze Nerdtool, Script ist aber unabhängig von der Software)

{{< figure src="/uploads/2010/08/Bildschirmfoto-2010-08-15-um-13.19.43-300x187.png" >}}

Bash-Script: [http://gist.github.com/525385](http://gist.github.com/525385)


``` bash
ESC=$(printf "e")
echo "$ESC[34;47mDATE$ESC[0m"
date
echo ""
echo "$ESC[34;47mUPTIME$ESC[0m"
uptime
echo ""
echo "$ESC[34;47mSTATUS$ESC[0m"
top -l1 -u -o cpu -S | head -n 12
echo ""
echo "$ESC[34;47mEstablished$ESC[0m"
lsof -i -n | grep -i established | awk '{print $1" "$8" "$9 }' | head -n 18
echo ""
echo "$ESC[34;47mListen$ESC[0m"
lsof -i -n | grep -i listen | awk '{print $1" "$8" "$9 }' | head -n 18
```



Besonderheit hierbei: Der Escape-Character muss so _zwingend_ wie beschrieben eingesetzt werden. Normale Ausgabe wird nicht entsprechend wahrgenommen. Er ist nötig um die Farbtöne innerhalb des Scripts zu managen ;)
