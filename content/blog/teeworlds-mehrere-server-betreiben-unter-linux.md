---
date: 2009-12-09T15:57:37+02:00
type: post
comments: true
title: Teeworlds | Mehrere Server betreiben unter Linux
aliases:
- /archives/800
- /blog/2009/12/09/teeworlds-mehrere-server-betreiben-unter-linux
categories:
- Development
- Linux
tags:
- cronjob
- debian
- mehrere Server
- multi server
- paralell
- server
- skript
- Teeworlds
- teeworldsserver
---

Die 4 Teeworldsserver die auf zwetschge.org laufen, stellten ihre Dienste
ehrilchgesagt in einem ziemlichen WirrWarr aus Binaries und Configfiles.
Nach Neuorganisition, der technical Overview hier:

Dateistruktur:
Ausgabe mit tree -L 3 in /home/teeworlds/:
[FileTree](/uploads/2009/09/4)

Configfiles:
Individuell für jeden Server
[Configs](/uploads/2009/09/3)

Startskript:
Nach Reboot automatisch ausgeführt
[init](/uploads/2009/09/2)

Syntax:

    binary -f configfile >> ausgabe_logfile & #(& = als hintergrundprozess)

Sollte jemand auch mal mehrere Teeworldsserver nebeneinander betreiben und
den Überblick behalten wollen ;)
