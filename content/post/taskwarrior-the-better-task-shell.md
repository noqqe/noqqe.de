---
date: 2011-10-22T14:05:57+02:00
layout: post
slug: taskwarrior-the-better-task-shell
status: publish
comments: true
title: Taskwarrior | The better-task-shell
alias: /archives/1810
categories:
- Bash
- Coding
- Debian
- PlanetenBlogger
- Ubuntu
- Web
tags:
- bash
- better shell
- shell
- Task
- task-shell
- Taskwarrior
---

Eigentlich wollte ich das Projekt task-shell-ng nennen. Aber so gut ist es dann doch nicht geworden. Stattdessen hat es sich aber den Prefix better verdient ;)

Als ich vor ca. einem Monat Taskwarrior für mich entdeckt habe, war eigentlich alles gut. Ich hab mich über den integrierten interactive Mode wirklich gefreut. Anfangs. Mit der Zeit habe ich aber festgestellt, dass mich dieses "Ding" fast in den Wahnsinn treibt. Mir persönlich fehlen einfach elementare Features wie einfaches Cursor bewegen nach vorne und zurück. Überhaupt eine History zu haben wäre schon ein enormer Vorteil.

{{< figure src="/uploads/2011/10/1733_f644.gif" >}}

Ich hab mir dann kurzer Hand selber eine Taskwarrior Shell Variante gebaut, die im großen und ganzen auf einer Bash basiert.

**Features:**



	
  * History vorwärts und rückwärts via Pfeiltasten

	
  * Cursorbewegung vorwärts und rückwerts in der aktuellen Zeile

	
  * Alle Kommandos nativ benutzbar ( $ add pri:H pro:Living Miete zahlen )

	
  * ID's direkt nutzbar ( $ 34 edit oder $ 12 pri:H )

	
  * separate Logging Funktion in $HOME/.better-task-shell_history

	
  * OS Befehle weiterhin nutzbar! ( $ vim /home/user/foobar.txt )

	
  * Automatische Erkennung von doppelten Aliases

	
  * Automatische Alias Generierung fuer os-binaries ( $ ls  = task ls  ; os-ls = /bin/ls )

	
  * Auto-Komplettierung aller Taskwarrior Befehle und definierte Aliase


**Known Bugs:**



	
  * Neu angelegt tasks  können derzeit noch nicht via ID aufgerufen werden.
$ add Uberweisung einwerfen
Created Task 45
$ 45 pri:H
bash: 45: Kommando nicht gefunden
Für beim Start bestehende Einträge funktioniert dies allerdings problemlos.

	
  * Mode -v ist bis jetzt noch nicht benutzbar aber bereits implementiert.


Das ganze gibts jetzt unter [http://github.com/noqqe/better-task-shell](http://github.com/noqqe/better-task-shell)

**Usage:**

```
git clone git@github.com:noqqe/better-task-shell.git
$ cd better-task-shell
$ ./better-task-shell
```

