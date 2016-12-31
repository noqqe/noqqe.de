---
aliases:
- /archives/1810
- /blog/2011/10/22/taskwarrior-the-better-task-shell
comments:
- author: Dirk Deimeke
  content: <p>Ich kann das Repository leider nicht klonen. Ist das eventuell privat?</p>
  date: '2011-10-23T14:41:07'
- author: noqqe
  content: '<p>Hi dirk, </p><p>ich hab gerade auch mal ausgecheckt, aber bei mir funktionierts.
    </p><p>Versuch doch alternativ mal:</p><p><a rel="nofollow">git://github.com/noqqe/better-...</a></p><p>oder:
    </p><p><a href="https://github.com/noqqe/better-task-shell.git" rel="nofollow">https://github.com/noqqe/bette...</a></p>'
  date: '2011-10-23T14:57:46'
- author: Dirk Deimeke
  content: <p>Die Alternative hat funktioniert, aber die Shell tut noch nicht ganz
    das, was sie soll. Wenn ich "ll" eintippe, wird mein selbst definierter Report
    nicht angezeigt.</p>
  date: '2011-10-23T15:27:08'
- author: noqqe
  content: "<p>Hi Dirk,</p><p>ich habe das irgendwie vergessen. Ist aber schon gepatched
    :)</p><p>W\xE4r cool wenn dus nochmal probieren k\xF6nntest und deine Meinung
    abgibst als Taskwarrior Entwickler ;)</p>"
  date: '2011-10-23T18:04:05'
- author: Dirk Deimeke
  content: "<p>Das ist f\xFCr mich noch nicht benutzbar, da es kaum Nutzen gibt. Deine
    Shell ist nur eine Reihe von aliases ...</p><p>Das ist mir zu wenig. \"?\" geht
    nicht \"help\" auch nicht.</p><p>Aber, wenn es Dir hilft, ist es genau richtig.</p>"
  date: '2011-10-23T18:19:35'
- author: noqqe
  content: <p>Ja richtig. Ich wollte lediglich die Features die ich von einer normalen
    Shell gewohnt bin weiter nutzen und hab mir die Umgebung von Taskwarrior nachgebaut.</p><p>Wie
    das realisier ist ist m.E. Egal.</p>
  date: '2011-10-23T19:30:55'
date: '2011-10-22T12:05:57'
tags:
- development
- web
- shell
- better shell
- taskwarrior
- task
- ubuntu
- task-shell
- debian
- bash
title: Taskwarrior | The better-task-shell
---

Eigentlich wollte ich das Projekt task-shell-ng nennen. Aber so gut ist es
dann doch nicht geworden. Stattdessen hat es sich aber den Prefix better
verdient ;)

Als ich vor ca. einem Monat Taskwarrior für mich entdeckt habe, war
eigentlich alles gut. Ich hab mich über den integrierten interactive Mode
wirklich gefreut. Anfangs. Mit der Zeit habe ich aber festgestellt, dass
mich dieses "Ding" fast in den Wahnsinn treibt. Mir persönlich fehlen
einfach elementare Features wie einfaches Cursor bewegen nach vorne und
zurück. Überhaupt eine History zu haben wäre schon ein enormer Vorteil.

{{< figure src="/uploads/2011/10/1733_f644.gif" >}}

Ich hab mir dann kurzer Hand selber eine Taskwarrior Shell Variante gebaut,
die im großen und ganzen auf einer Bash basiert.

### Features

  * History vorwärts und rückwärts via Pfeiltasten
  * Cursorbewegung vorwärts und rückwerts in der aktuellen Zeile
  * Alle Kommandos nativ benutzbar ( $ add pri:H pro:Living Miete zahlen )
  * ID's direkt nutzbar ( $ 34 edit oder $ 12 pri:H )
  * separate Logging Funktion in $HOME/.better-task-shell_history
  * OS Befehle weiterhin nutzbar! ( $ vim /home/user/foobar.txt )
  * Automatische Erkennung von doppelten Aliases
  * Automatische Alias Generierung fuer os-binaries ( $ ls  = task ls  ; os-ls = /bin/ls )
  * Auto-Komplettierung aller Taskwarrior Befehle und definierte Aliase

### Known Bugs


  * Neu angelegt tasks  können derzeit noch nicht via ID aufgerufen werden.
```
$ add Uberweisung einwerfen
Created Task 45
$ 45 pri:H
bash: 45: Kommando nicht gefunden
```
  * Mode -v ist bis jetzt noch nicht benutzbar aber bereits implementiert.


Das ganze gibts jetzt unter
[http://github.com/noqqe/better-task-shell](http://github.com/noqqe/better-task-shell)

### Usage

```
git clone git@github.com:noqqe/better-task-shell.git
$ cd better-task-shell
$ ./better-task-shell
```