---
date: 2011-05-03T12:04:34+02:00
type: post
slug: pidgin-chat-nachrichten-zahlen
comments: true
title: Pidgin Chat-Nachrichten zählen
aliases:
- /archives/1672
categories:
- Bash
- Coding
- Jabber
- Linux
- PlanetenBlogger
- Shell-Zauberei
- Ubuntu
- Web
tags:
- chat
- ICQ
- libpurple
- messages
- nachrichten
- pidgin
- purble
- shell
- Shell-Zauberei
---

**Code**
```
cd $HOME/.purple/logs/icq/987654321
cat */** | grep "^<font" | sed -e 's#.* <b>(.*):</b>.*#1#g' | grep -v "^<"|sort | uniq -c | sort -rn
```


**Hintergrund**
Heute ist der 3. Mai und morgen ist meine Abschlussprüfung zum Fachinformatiker. Natürlich sollte man da andere Sachen machen, als Shell Zeilen zu schreiben, aber mein Kopf fühlt sich von den letzten beiden Wochen, die ich während meines Lernurlaubs damit verbracht habe jegliches während meiner Ausbildung vermitteltes theoretisches Wissen in mich aufzusaugen, sowas von überfüllt an, dass ich während einer Lernpause kurz den Instant-Messenger meiner Wahl angeschmissen habe. Kurz darauf wollte ich wissen, wie viele Nachrichten ich wohl seit Neuinstallation schon erhalten und gesendet habe.

**Funktion**
Die Zeile bewirkt im Grunde nicht anderes als alle Files auszugeben, nach Zeilen zu suchen die auf das Schema passen die dem einer "Nachricht gesendet" Zeile ähneln und sortiert, zählt und sortiert diese wieder. Vorrausgesetzt die Logfiles sind durch Pidgin im HTML Format abgespeichert.

