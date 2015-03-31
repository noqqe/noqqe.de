---
date: '2009-05-26 18:43:25'
layout: post
slug: shell-timerobot-007-scripts-und-singlefiles-strukturiert-sichern
status: publish
comments: true
title: Shell | TimeRobot 0.0.7 Scripts und SingleFiles strukturiert sichern
alias: /archives/620
categories:
- Coding
- General
- Linux
tags:
- backup
- bash
- shell
- timerobot
---

Hab grade die nächste Version von TimeRobot "released". (Gott hört sich das übertrieben an :D ) Im Endeffekt ist die Logik etwas überarbeitet und das Logfile übersichtlicher und schöner geworden.

Wer sich nicht mehr so wirklich erinnern kann was TimeRobot tut: Es sichert einzelne Files und Scripte die in einem /etc File konfiguriert werden. Wenn der Cronjob(Stunde/Woche oder Tag) jetzt durchläuft der eingerichtet wird vergleicht das Skript die md5 Summe des Aktuellen Files (zb. /home/user/.bashrc) mit dem des zuletzt gesicherten File (/time/robot/backuppfad/bashrc/2009-05-26-15-12bashrc) und erkennt je nach dem den unterschied oder auch nicht. Falls ja wird natürlich eine Kopie mit dem aktuellen Datum und Uhrzeit erstellt und ins verzeichnis gesichert.

Das Verzeichnis lässt sich dann mit "timerobot -l" abrufen und dann sieht das so aus:

[TimerobotScreen](http://zwetschge.org/pic/timerobotpic.png)

Zum download wie immer hier:

[Timerobot Projekt](http://zwetschge.org/timerobot/)
