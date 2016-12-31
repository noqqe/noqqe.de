---
aliases:
- /blog/2010/11/27/redmine-git-bare-repository-workaround
- /archives/1396
comments:
- author: git spasti
  content: "<p>Das ganze l\xE4uft effizienter wenn nicht geklont sondern lediglich
    gepulled wird.</p><p>Mir f\xE4llt momentan kein Grund ein warum man an dieser
    Stelle auf die Vorz\xFCge von Git (als SCM) verzichten sollte. Es ist nicht n\xF6tig
    jedes mal einen Full-Clone zu machen.</p>"
  date: '2011-12-22T23:01:26'
- author: noqqe
  content: "<p>Da hast du Recht. Das war auch mein erster Ansatz. </p><p>Aber ich
    h\xE4tte mir da einen Check \xFCberlegen m\xFCssen obs das Repo schon gibt oder
    ob es erst angelegt wurde und ich deshalb clonen m\xFCsste. </p><p>Darauf hatte
    ich damals keine Lust und habs dann einfach alles wie es ist erschlagen.</p>"
  date: '2011-12-27T14:12:49'
- author: chris
  content: "<p>Eigentlich nur Rechte Webserver auf die Dirs pr\xFCfen und ggf. bare
    repo richtig chmoden!<br>Geht bei mir..!</p>"
  date: '2012-11-28T18:59:21'
date: '2010-11-27T11:49:42'
tags:
- development
- repositories
- git
- workaround
- repository
- barerepo
- administration
- working copy
- redmine
- bare
- webinterface
- issue
title: Redmine | Git Bare Repository Workaround
---

Die letzten beiden Tage habe ich mich etwas mit
[Redmine](http://www.redmine.org/) beschäftigt. Das Webinterface für
Versionskontrollsysteme hat allerdings eine kleine Macke. Zumindest ist das
mein Verständnis des Umstands.

Git-Repositories sind in Redmine nur über Working Copies einzulesen. Der
gute Stil (und gitolite) stellt serverseitige Git-Repositories als Bare
Repos dar. Bare Repos sind im Grunde nur das, was das .git/ Verzeichnis in
jeder Working Copy darstellt. Es befinden sich also keine Klartext Files
darin.

{{< figure src="/uploads/2010/11/2431284305_9c4952e7f6.jpg" >}}
Creative Commons by [Impact Tarmac](http://www.flickr.com/photos/bbcolin/)

Diese Bare Repos lassen sich in Redmine aber nicht hinzufügen. Ob es
einfach technisch nicht machbar ist, oder eine Macke seitens Redmine kann
ich nicht ausmachen. Fakt ist allerdings, dass ich momentan dazu gezwungen
bin Working-Copies für Redmine vorzuhalten.

Ein bisschen Dokumentation zu lesen hat mich da schon um einiges weiter
gebracht. Es gibt unter der Redmine Community anscheinend zwei verbreitete
Lösungsansätze.

  * Cronjob der alle heilige Zeit auf Änderungen prüft
  * Hook in Bare Repos einfügen, welcher auslöst wenn gepushte Änderungen vorliegen

Kurz: Ich entschied mir für die zweite Lösung. Um Ressourcen zu sparen.
Oder weniger Cronjobs laufen zu haben. Erschien mir effizienter.

## Hook für Bare Repo

In Hook-verzeichnis wechseln. Beispielsweise:
/home/git/repositories/repo1.git/hooks/. Richtigen Hook auswählen.
**post-recieve** führt Skripte aus, welche _nach_ dem erhalt von Commits
ausgeführt werden.  Das Skript sieht vor, das ServerRepo an einen zentralen
Sammelpunkt auszuchecken, Rechte anzupassen und für Redmine schreibbar zu
machen.

Werden jetzt neue Commits gepushed, wird automatisch ein neues Repo
gecloned und die Änderungen sind sofort in Redmine ersichtlich.