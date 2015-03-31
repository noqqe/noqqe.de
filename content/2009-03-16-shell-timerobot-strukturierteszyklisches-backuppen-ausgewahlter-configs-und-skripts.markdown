---
date: '2009-03-16 12:06:41'
layout: post
slug: shell-timerobot-strukturierteszyklisches-backuppen-ausgewahlter-configs-und-skripts
status: publish
comments: true
title: Shell | TimeRobot - Strukturiertes/Zyklisches Backuppen ausgewählter Configs
  und Skripts
alias: /archives/548
categories:
- Coding
- Linux
tags:
- .conf
- Back-UP
- backup
- bash
- config
- configfiles
- debian
- Linux
- shell
- shell script
- skript
- structure
- time
- timerobot
- ubuntu
---

Moin,
aufgrund diverser Handling-Probleme mit RCS und anderen Structure-Backup-Tools, war ich einfach so frei und hab mein eigenes kleines System zusammen gebastelt. Denkbar einfach ist der Sinn des ganzen. Configdateien und Skripte ändern sich. Und manchmal funktioniert dank einer kleinen Änderung das ganze Programm nicht mehr. Wenn man nur wüsste welche kleine Änderung man nochmal getätigt hat? TimeRobot (so hab ich das ganze getauft) sichert sowohl manuell (kurzer Aufruf), als auch Zyklisch (Woche Tag Stunde, in timerobot.conf konfigurierbar) entsprechende Configfiles, Skripte und sonstige Dateien die vom User im .conf File hinterlegt wurden. Somit erstellt sich eigenständig eine Basis von Files auf die man im Notfall zurückgreifen kann.

Wer intressiert ist:
Runterzuladen gibts das ganze unter: [http://zwetschge.org/timerobot/timerobot-0.0.5/timerobot-0.0.5-all.deb
](http://zwetschge.org/timerobot/timerobot-0.0.5/timerobot-0.0.5-all.deb)

Und hier noch ein kurzes Anwendungsbeispiel:

```
$ dpkg -i timerobot-0.0.5-all.deb # Installation
```


```
$ vim /etc/timerobot.conf # Erstanpassung der Pfade (selbsterklärend)
```


```
$ timerobot -t # Erst-Einrichtung der in timerobot.conf gesetzen zyklischen Sicherung
```


```
$ timerobot -a X11 # Ersten Eintrag hinzufügen
TimeRobot-Verzeichnis X11 wurde angelegt.
Editiere /etc/timerobot.conf um die Pfade anzupassen
```


```
$ timerobot -a aliases # Zweiter Eintrag
TimeRobot-Verzeichnis aliases wurde angelegt.
Editiere /etc/timerobot.conf um die Pfade anzupassen
```


```
$ timerobot -a hosts # Dritter Eintrag (usw.)
TimeRobot-Verzeichnis hosts wurde angelegt.
Editiere /etc/timerobot.conf um die Pfade anzupassen
```


```
$ timerobot -l # Struktur des Sicherungsverzeichnisses anzeigen
/home/user/timerobot
|-- X11
|-- aliases
```
-- hosts
3 directories, 0 files
```


```
$ timerobot -u X11 # Erstmalige Sicherung zb der xorg.conf
```


```
$ timerobot -l
/home/user/timerobot
|-- X11
|   `-- 2009-03-16-10-51-28-X11
|-- aliases
```
-- hosts
3 directories, 1 file
```


```

$ timerobot --autoall # Komplett-Sicherung aller angegebenen Files
```


```

$ timerobot -l
/home/user/timerobot
|-- X11
|   |-- 2009-03-16-10-51-28-X11
|   `-- 2009-03-16-10-52-07-X11
|-- aliases
|   `-- 2009-03-16-10-52-07-aliases
```
-- hosts
```
-- 2009-03-16-10-52-07-hosts
3 directories, 4 files
```

