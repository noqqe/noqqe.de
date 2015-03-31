---
date: 2010-12-14T16:33:45+02:00
layout: post
slug: bash-backup-eines-files-mit-zeitstempel-einfach-gemacht
status: publish
comments: true
title: 'Bash | Backup eines Files mit Zeitstempel in unkompliziert. '
alias: /archives/1425
categories:
- Bash
- Coding
- Debian
- Ubuntu
- ubuntuusers
tags:
- admins
- backup
- bashrc
- buf
- commandline
- commandlinefu
- dirty
- function
- kommandozeile
- quick
- schnell
---

[Auf eine besondere kleine Funktion]( http://www.commandlinefu.com/commands/view/7292/backup-a-file-with-a-date-time-stamp) möchte ich kurz aufmerksam machen, weil sie viele Administratoren oder  Sonst-Wie-Konfigurationen-vornehmende-Benutzer wahrscheinlich auch brauchen könnten:

```
 buf () { filename=$1; filetime=$(date +%Y%m%d_%H%M%S); cp ${filename} ${filename}_${filetime}; }
```


Man spielt an der main.cf des Postfix herum oder bastelt an der xorg.conf. Und wie man das eben von dem erfahrenen Linux-Benutzer seines Vertrauens (oder dem UbuntuUsersWiki) gelernt hat, sichert man vorher die Konfigurationsdatei mit dem obligatorischen:

```
cp xorg.conf xorg.conf.bak
```


Je nach Intensität der Anpassungen häufen sich auch Backupfiles a la "xorg.conf.bak.bak" oder ".bak2 bis .bak5000"

Fügt man die obige Funktion in seine .bashrc hinzu, erledigt sich das Thema relativ schnell.

```
buf xorg.conf
```


Erstellt automatisch ein File mit dem aktuellen Zeitstempel im aktuellen Verzeichnis. Nochmal im formatierten Zustand:

```
buf () {
filename=$1
filetime=$(date +%Y%m%d_%H%M%S)
cp ${filename} ${filename}_${filetime}
}
```


Dass [commandlinefu.com](http://commandlinefu.com) eine großartige Site ist, ist wahrscheinlich für die meisten nichts Neues. Teileweise regnet es dort wunderschöne kleine Kommandozeilen für den Alltag.
