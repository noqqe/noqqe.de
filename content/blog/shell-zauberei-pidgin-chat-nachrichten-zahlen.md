---
aliases:
- /archives/1672
date: '2011-05-03T10:04:34'
slug: pidgin-chat-nachrichten-zahlen
tags:
- development
- web
- shell
- nachrichten
- messages
- pidgin
- shell-zauberei
- chat
- ubuntu
- icq
- purble
- libpurple
title: "Pidgin Chat-Nachrichten z\xE4hlen"
---

### Code

``` bash
cd $HOME/.purple/logs/icq/987654321
cat */** | grep "^<font" | sed -e 's#.* <b>(.*):</b>.*#1#g' | grep -v "^<"|sort | uniq -c | sort -rn
```

### Hintergrund

Heute ist der 3. Mai und morgen ist meine Abschlussprüfung
zum Fachinformatiker. Natürlich sollte man da andere Sachen machen, als
Shell Zeilen zu schreiben, aber mein Kopf fühlt sich von den letzten beiden
Wochen, die ich während meines Lernurlaubs damit verbracht habe jegliches
während meiner Ausbildung vermitteltes theoretisches Wissen in mich
aufzusaugen, sowas von überfüllt an, dass ich während einer Lernpause kurz
den Instant-Messenger meiner Wahl angeschmissen habe. Kurz darauf wollte
ich wissen, wie viele Nachrichten ich wohl seit Neuinstallation schon
erhalten und gesendet habe.

### Funktion

Die Zeile bewirkt im Grunde nicht anderes als alle Files
auszugeben, nach Zeilen zu suchen die auf das Schema passen die dem einer
"Nachricht gesendet" Zeile ähneln und sortiert, zählt und sortiert diese
wieder. Vorrausgesetzt die Logfiles sind durch Pidgin im HTML Format
abgespeichert.
