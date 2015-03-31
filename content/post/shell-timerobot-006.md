---
date: 2009-03-18T12:02:12+02:00
type: post
slug: shell-timerobot-006
status: publish
comments: true
title: 'Shell | timeRobot 0.0.6 '
alias: /archives/561
categories:
- Coding
- Linux
tags:
- archivierung
- backup
- config
- shell
- skript
- timerobot
---

Der Robot ist bei mir jetzt ca 4 Tage im Einsatz und ich muss sagen das ganze wird relativ schnell unübersichtlich. Jede angegebene Config wird bei jedem Cronjob komplett gesichert. Da sammeln sich nach einer weile Recht viele Daten. Um das zu umgehen gibts jetzt die neue Version von timeRobot:

Changelog:

	
  * Via md5-Summe erkennt timeRobot automatisch ob die letzte gesicherte Datei der aktuellen Datei entspricht. Falls dies Eintritt erfolgt keine Sicherung der Datei, aber eine Ausgabe und ein Log-Eintrag.

	
  * timeRobot ist im allgemeinen "gesprächiger" geworden. Sowohl bei Einträgen, Updates als auch bei Auto-updates usw.

	
  * Es wird mehr geLoggt und vor allem detailierter.

	
  * Bugfix: Automatische Cron-Job Einrichtung gefixt.


Rausgekommen ist also eine effizientere, humanere Version :D

Zu finden wie bei den anderen Projekten auch unter:

[http://zwetschge.org/timerobot/timerobot-0.0.6/timerobot-0.0.6-all.deb](http://zwetschge.org/timerobot/timerobot-0.0.6/timerobot-0.0.6-all.deb)

to be continued..
Flo
