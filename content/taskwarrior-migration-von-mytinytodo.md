---
date: 2011-09-25T13:58:52+02:00
layout: post
slug: taskwarrior-migration-von-mytinytodo
status: publish
comments: true
title: 'Taskwarrior | Migration von MyTinyTodo '
alias: /archives/1766
categories:
- Bash
- Coding
- Debian
- General
- PlanetenBlogger
- SQL
- Ubuntu
- Web
tags:
- GTD
- import
- mig
- Migration
- mirgration
- MyTinyTodo
- Task
- Taskwarrior
- todo
- todolist
- tutorial
- ubucon
---

Im [Programm der Ubucon](http://ikhaya.ubuntuusers.de/2011/09/06/das-programm-der-ubucon-2011-steht/) habe ich einen [Vortrag](http://www.ubucon.de/programm/taskwarrior) von [Dirk](http://www.deimeke.net/dirk/blog/) über [Taskwarrior](http://taskwarrior.org) gefunden. Aus Gründen hat mich das Tool interessiert. Ich habe zwar bereits eine Todo-Verwaltung (nämlich [MyTinytodo](http://mytinytodo.net/)) aber trotzdem.

{{< figure src="/uploads/2011/09/patrick.jpg" >}}

Nach zwei Tagen herumspielen und einem wirklich hervorragendem [How-To](http://taskwarrior.org/projects/taskwarrior/wiki/Tutorial) hat mir Taskwarrior allerdings so gefallen, dass ich mich entschlossen habe MyTinyTodo aufzugeben und zu Taskwarrior zu wechseln. Da ich meine TodoListe aber auch gerne als eine Art "Log" verwende musste ich die Aufgaben erst aus MyTinyTodo migrieren.

Realisiert habe ich das über ein kleines Bash-Script, welches die Daten aus der MyTinyTodo MySQL Datenbank ausliesst und daraus automatisch Taskwarrior Statements formuliert. Allerdings: Ich habe absichtlich nur abgeschlossene Tasks über die "LOG"-Funktion von Taskwarrior einfuegt.



Aus dem Datenbank Eintrag:



    mysql> select d_completed,  mtt_lists.name as liste, title from mtt_todolist, mtt_lists where mtt_lists.id = mtt_todolist.list_id order by d_completed desc limit 1;
    +-------------+-------+--------------+
    | d_completed | liste | title        |
    +-------------+-------+--------------+
    |  1307714692 | Home  | Miete zahlen |
    +-------------+-------+--------------+
    1 row in set (0.00 sec)



formuliert das Skript beispielsweise folgendes Kommando für Taskwarrior:

```
task log due:20110610 pro:IMPORT-Home Miete zahlen
```


Das passiert eben für jede Eintrag in der Todo-Datenbank, damit ich dort nachschlagen kann wann was passiert ist.
