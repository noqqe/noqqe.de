---
date: 2010-10-24T21:06:33+02:00
type: post
comments: true
title: Bash | minimal-bash-debug
aliases:
- /blog/2010/10/24/bash-minimal-bash-debug
- /archives/1332
categories:
- Development
- Debian
- git
- Linux
- PlanetenBlogger
- Ubuntu
tags:
- bash
- bash debug script
- debug
- debugging
- minimal debug
- script
- shell debug script
- shell debugging
---

Bei Bash Scripten gibt es ja genug Möglichkeiten zu debuggen. Jeder der
Bash Scripte schreibt, hat da so seine eigenen Vorlieben. Nachdem ich etwas
geforscht hatte fand ich diverse Methoden a la:

```
#!/bin/bash
_DEBUG=true
if [ "$_DEBUG" == "true" ]; then echo "some debug informations" ; fi
```

Auch schön fand ich den eingebauten Debug Modus der Shell:

```
$ sh -x script.sh
```


Aber auch sehr umfangreiche Tools wie
[log4sh](http://log4sh.svn.sourceforge.net/svnroot/log4sh/trunk/source/1.5/doc/log4sh.html)
sind vorhanden. Ganz zu Schweigen von wahllosen "echo's" die über das ganze
Skript verteilt sind. Irgendwie gefiel mir aber keine der Möglichkeiten so
wirklich. Der Code wird unlesbar, die Ausgabe ist zu kryptisch oder der
Lernaufwand ist jenseits von Gut und Böse. Deshalb habe ich mich hingesetzt
und überlegt, wie ich eine Debugging-Engine gerne hätte. Wie müsste sowas
aussehen, damit ich es verwenden würde?

  * Code sollte nicht verunstaltet werden.
  * Debug Informationen sollten gut leserlich ausgegeben werden.
  * Debug Nachrichten sollten sich leicht zu erzeugen sein.
  * Sollte sehr einfach abschaltbar und skalierbar sein.
  * Minimaler Umfang des Gesamtsystems

Nun ja, im typischen Sonntag Nachmittag Größenwahn hab ich mich bemüht meine
Anforderungen an den Bash Debugger umzusetzen. Ob es mir gelungen ist lasse
ich jetzt mal dahingestellt, denn das ist mit Sicherheit Ansichtssache.
Aber das aus dem Bedarf heraus entstandene
**[minimal-bash-debug](http://github.com/noqqe/minimal-bash-debug/)** hat
mich dann doch einigermaßen zufrieden gestellt.

### Download

    $ cd path/to/your/bashscript
    $ wget http://github.com/noqqe/minimal-bash-debug/raw/master/.minimal-bash-debug
    $ chmod +x .minimal-bash-debug

### Implementation

Das runtergeladene (wirklich sehr schmale) Script dient jetzt als Auswerter
für die Debug Informationen. In das Bash Script, welches es zu debuggen
gilt, muss nun folgende Funktion eingefügt werden. Mein Ziel Dabei war es
den Snippet der Implementierung so klein wie nur irgendwie möglich zu
halten, um das Skript nicht zu verunstalten.

``` bash
debug() {
  debug=2 # set debug level 0|1|2|3
  if [ -x .minimal-bash-debug ]; then
  ./.minimal-bash-debug $debug $1 $2 "$3"
  fi
}
```

### Usage

Die Implementation der Funktion debug() in das eigene Skript hat mehrere
Vorteile. Aber zuerst zur Benutzung.

```
debug 2 syslog "Variable VAR is $VAR"
debug 3 echo "Variable FOO is $FOO"
```

Das ganze ist simpel.

  * debug aufrufen (debug)
  * debug level bestimmen (2)
  * debug mode auswählen (echo|syslog)
  * debug message übergeben (Variable FOO is $FOO)

Das war auch schon der ganze Zauber. Das schöne daran ist, ist das
.minimal-bash-debug Script einmal nicht mehr vorhanden/nicht ausführbar
oder das debug level einfach auf 0 gesetzt, verfallen alle debug Zeilen in
sofortiges Stillschweigen.

Der kurze Auszug der Erklärung ist natürlich relativ wenig. Deswegen gibt
es auf Github
[http://github.com/noqqe/minimal-bash-debug/](http://github.com/noqqe/minimal-bash-debug/)
ein Beispiel der Benutzung und ein README.

example.sh:
[http://github.com/noqqe/minimal-bash-debug/raw/master/example.sh](http://github.com/noqqe/minimal-bash-debug/raw/master/example.sh)
