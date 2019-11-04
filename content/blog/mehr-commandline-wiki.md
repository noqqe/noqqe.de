---
comments:
- author: Martin
  content: "> Extrem fragw\xFCrdig finde ich aber die Downloads/Zugriffszahlen.\n\nDie
    Statistiken lassen sich nur begrenzt verwerten, sind eben genau Bots und Mirrors."
  date: '2015-06-06T21:44:42.379442'
- author: P-Tronic
  content: 'Ich wollte es gerade installiere, werde aber aus dem Error nicht schlau,
    benutze Archlinux im Configparser scheint aber ein Syntaxerror zu sein, liege
    ich da richtig ?

    http://paste.pwhl.tk/img/23-31-04-502190502f155c4637bf96873b9f8f96.png'
  date: '2015-06-06T23:33:22.671053'
- author: noqqe
  content: hast du Python wirklich auf 2.7?
  date: '2015-06-06T23:36:40.230792'
- author: P-Tronic
  content: "Ich Hab Python 3.4, l\xE4uft cmddocs auf 2.7, Wenn es so ist, habe ich
    keinen Hinweis dazu entdeckt."
  date: '2015-06-07T12:12:23.314409'
- author: noqqe
  content: ist im moment nur in 2.7. Die Information ist auch in pip.
  date: '2015-06-07T13:03:51.894236'
- author: noqqe
  content: Alles klar, danke!
  date: '2015-06-07T13:04:30.285276'
- author: Anonymous
  content: 'Hi! versucht es doch mal mit:

    pip2.7 install cmddocs

    Dann klappts auch mit verschiedenen Python versionen

    (...zumindest auf Arch)'
  date: '2015-06-07T13:42:14.262148'
- author: Anonymous
  content: "andererseits bekomme ich dann folgenden Fehler:\n\ncmddocs - press ? for
    help\ncmddocs> dirs\nTraceback (most recent call last):\n  File \"/usr/sbin/cmddocs\",
    line 9, in <module>\n    load_entry_point('cmddocs==0.10.5', 'console_scripts',
    'cmddocs')()\n  File \"/usr/lib/python2.7/site-packages/cmddocs/__init__.py\",
    line 246, in main\n    Cmddocs().cmdloop()\n  File \"/usr/lib/python2.7/cmd.py\",
    line 142, in cmdloop\n    stop = self.onecmd(line)\n  File \"/usr/lib/python2.7/cmd.py\",
    line 221, in onecmd\n    return func(arg)\n  File \"/usr/lib/python2.7/site-packages/cmddocs/__init__.py\",
    line 112, in do_dirs\n    return list_directories(dir)\n  File \"/usr/lib/python2.7/site-packages/cmddocs/articles.py\",
    line 16, in list_directories\n    subprocess.call([\"tree\", \"-d\", dir])\n  File
    \"/usr/lib/python2.7/subprocess.py\", line 522, in call\n    return Popen(*popenargs,
    **kwargs).wait()\n  File \"/usr/lib/python2.7/subprocess.py\", line 710, in __init__\n
    \   errread, errwrite)\n  File \"/usr/lib/python2.7/subprocess.py\", line 1335,
    in _execute_child\n    raise child_exception\nOSError: [Errno 2] No such file
    or directory\n\n und danach ist beendet sich cmddocs.\n\nIst das vielleicht irgend
    ein Modul, das ich in python3-Version benutze, und es sollte Version2 sein?"
  date: '2015-06-07T15:46:03.205676'
- author: Martin
  content: "`tree` per *pacman* installieren, ist 'ne externe Abh\xE4ngigkeit."
  date: '2015-06-07T16:42:11.667842'
- author: Anonymous
  content: super! Vielen Dank!
  date: '2015-06-07T16:44:45.020207'
- author: noqqe
  content: das sollte ich vielleicht Dokumentieren...
  date: '2015-06-07T20:33:17.709410'
- author: Dr. Azrael Tod
  content: '%s/Packet/Paket ;-)'
  date: '2015-06-08T06:52:52.859658'
date: '2015-06-06T16:28:32'
tags:
- wiki
- web
- shell
- python
- cmd
- cmddocs
- code
- opensource
- bsd
- commandlinedocs
- documentation
title: Mehr Commandline Wiki
---

Wenn ich nicht gerade mit [IRC-Bot](https://github.com/k4cg/Rezeptionistin)
schreiben oder das [Wiki für k4cg.org](https://k4cg.org) schön/neu mache, bastle
ich immer wieder an [cmddocs](https://github.com/noqqe/cmddocs).

<!--more-->

{{< figure src="/uploads/2015/06/search.png" >}}

Der Code wurde umfangreicher, robuster, paar Features kamen hinzu. Tollste Sache
ist aber, dass ich mal was auf
[pypi](https://pypi.python.org/pypi/cmddocs/0.10.5) hoch geladen hab.
Das ist tatsächlich das erste Packet was ich dort hochlade. Demnach kann man
`cmddocs` jetzt einfach über `pip` installieren.

``` bash
pip install cmddocs
cmddocs
```

\o/, Extrem fragwürdig finde ich aber die Downloads/Zugriffszahlen.

{{< figure src="/uploads/2015/06/pip.png" >}}

So schön ich es fände wenn es so wäre, aber das kann einfach nicht sein. Wenn
jemand nen Link hat wie diese Stats generiert werden oder eine Erklärung hat
(Bots, Mirrors?) bitte, her damit.

Wesentlichste Änderung ist wahrscheinlich, dass Markdown->ANSI Highlightning
jetzt nicht mehr nur über 4 Zeilen Regex realisiert ist. Ich hab den Markdown
Lexer von `mistune` verwendet und die Output-Funktionen so überschrieben, dass
statt HTML ANSI Codes herauskommen.

``` python
class md_to_ascii(mistune.Renderer):

    # Pagelayout
    def block_code(self, code, lang):
        return '\n\033[92m%s\033[0m\n' % code.strip()
    def header(self, text, level, raw=None):
        if level == 2 or level == 1:
            head = '\n\033[4m\033[1m\033[37m%s\033[0m\n' % text.strip()
        else:
            head = '\n\033[1m\033[37m%s\033[0m\n' % text.strip()
        return head
    def block_quote(self, text):
        return '\n%s\n' % text.strip()
    def block_html(self, html):
        return '\n%s\n' % html.strip()
    def hrule(self):
        return '---'
    def list(self, body, ordered=True):
        return '\n%s' % body
    def list_item(self, text):
        return '* %s \n'  %  text.strip()
```

Schönen dank an Flowwar für die Idee an der Stelle ;)
Ansonsten noch folgende Dinge eingebaut:

* `undo` Feature implementiert
* Suchfunktion gibt auch Ergebnisse in Dateinamen wieder (siehe Screenshot oben)
* Anständige Fehlermeldungen an mehreren Stellen
* Nicht vorhandene Configoptionen mit sinnvollen Defaults eingebaut
* Error Handling für nicht vorhandene Config/Datadir
* Description für `pip` eingefügt
* Readme nach RST konvertiert
* Promptfarbe konfigurierbar

Wems gefällt und Pull- oder Feature Requests hat, bitte immer her damit.
