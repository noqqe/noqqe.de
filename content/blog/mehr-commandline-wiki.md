---
categories:
- bsd
- code
- opensource
- osbn
- shell
- ubuntuusers
- web
date: 2015-06-06T18:28:32+02:00
tags:
- cmddocs
- wiki
- commandlinedocs
- cmd
- python
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
$ pip install cmddocs
$ cmddocs
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
