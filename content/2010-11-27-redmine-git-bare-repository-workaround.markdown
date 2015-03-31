---
date: '2010-11-27 13:49:42'
layout: post
slug: redmine-git-bare-repository-workaround
status: publish
comments: true
title: Redmine | Git Bare Repository Workaround
alias: /archives/1396
categories:
- Coding
- git
- PlanetenBlogger
tags:
- bare
- barerepo
- issue
- redmine
- repositories
- repository
- webinterface
- workaround
- working copy
---

Die letzten beiden Tage habe ich mich etwas mit [Redmine](http://www.redmine.org/)
beschäftigt. Das Webinterface für Versionskontrollsysteme hat allerdings eine kleine
Macke. Zumindest ist das mein Verständnis des Umstands.

Git-Repositories sind in Redmine nur über Working Copies einzulesen. Der gute Stil (und gitolite) stellt serverseitige Git-Repositories als Bare Repos dar. Bare Repos sind im Grunde nur das, was das .git/ Verzeichnis in jeder Working Copy darstellt. Es befinden sich also keine Klartext Files darin.

{% img center /uploads/2010/11/2431284305_9c4952e7f6.jpg %}
Creative Commons by [Impact Tarmac](http://www.flickr.com/photos/bbcolin/)

Diese Bare Repos lassen sich in Redmine aber nicht hinzufügen. Ob es einfach technisch
nicht machbar ist, oder eine Macke seitens Redmine kann ich nicht ausmachen. Fakt ist allerdings, dass ich momentan dazu gezwungen bin Working-Copies für Redmine vorzuhalten.

Ein bisschen Dokumentation zu lesen hat mich da schon um einiges weiter gebracht. Es gibt unter der Redmine Community anscheinend zwei verbreitete Lösungsansätze.

	
  * Cronjob der alle heilige Zeit auf Änderungen prüft
  * Hook in Bare Repos einfügen, welcher auslöst wenn gepushte Änderungen vorliegen


Kurz: Ich entschied mir für die zweite Lösung. Um Ressourcen zu sparen. Oder weniger Cronjobs laufen zu haben. Erschien mir effizienter.


## Hook für Bare Repo


In Hook-verzeichnis wechseln. Beispielsweise: /home/git/repositories/repo1.git/hooks/. Richtigen Hook auswählen. **post-recieve** führt Skripte aus, welche _nach_ dem erhalt von Commits ausgeführt werden.  Das Skript sieht vor, das ServerRepo an einen zentralen Sammelpunkt auszuchecken, Rechte anzupassen und für Redmine schreibbar zu machen.



Werden jetzt neue Commits gepushed, wird automatisch ein neues Repo gecloned und die Änderungen sind sofort in Redmine ersichtlich.
