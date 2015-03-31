---
date: 2009-01-29T17:33:06+02:00
layout: post
slug: syncn-110
status: publish
comments: true
title: syncN 1.1.0
alias: /archives/450
categories:
- Coding
- Linux
tags:
- automatische synchronisation
- debian
- server
- shell
- shell skript
- Sycronisation
- Sync
- SyncN
---

Morgen,

einmal mehr ist es dann soweit. syncN ist überarbeitet. Gibt viele schöne neue Sachen und aufgrund der vielen positiven Resonanz von der letzten Version, hoffe ich das auch Version 1.1.0 wieder gefällt. Unter anderem gibts folgende Bugs..äh. Features natürlich:

check: Mithilfe von awk und ls Parametern eine viel deutlichere und schönere Ausgabe des Dateidatums zustande gekommen(beidseitig).

md5 Summe: Desweiteren gibt es eine MD5 Summenprüfung die eindeutig prüft ob die Dateien identisch sind.

sha224/512 Summe: Im Grunde das gleiche wie MD5 Summe, nur eine andere Methode.

autosyncronisation: Mein ganzer Stolz. Mit sha224 Summen und Datumsvergleichen erkennt das Skript automatisch welche Datei älter/jünger oder sogar identisch sind. Je nach Situation leitet das Skript entsprechende Schritte ein(upload/download/Ausgabe der Files wenn Identisch)

Wie immer ist das ganze hier: [http://zwetschge.org/syncN/](http://zwetschge.org/syncN/) herunterzuladen und natürlich auch in "projects" vorzufinden.
Habe die Ehre.
