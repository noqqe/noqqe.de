---
date: '2016-03-02T18:30:50'
tags:
- development
- code
- python
- vim
- opensource
- pip
title: Minimales Python Development
---

Wegen `rvo` und [cmddocs](https://github.com/noqqe/cmddocs) musste ich in
letzter Zeit ziemlich an meinen Development Kenntnissen arbeiten. Dabei ist
eine Menge Zeug heraus gekommen.

{{< figure src="/uploads/2016/03/django.jpeg" >}}

Das ganze läuft relativ unprofessionell ab. Ich nutze weder PyCharm noch
sonst irgendwelche fancy IDEs. Ich entwickle meinen Python Kram unter
OpenBSD auf der Commandline mit `vim`. Der Rest sind kleine Zusatztools.

### Setuptools

Zugegeben, in das Thema `pip` Packages einzutauchen ist irgendwie unschön.
Undurchsichtig, solange man das nicht mal kurz gezeigt bekommt.

Ich nutze dafür [Cookiecutter](https://github.com/audreyr/cookiecutter), um
nicht alles selbst machen zu müssen.

{{< figure src="/uploads/2016/03/cookiecutter.png" >}}

### Imports und Virtualenv

Package Scopes waren für mich immer etwas fragwürdig. Vor allem Imports. Wie
importiere ich jetzt selbstgeschrieben Module?

~~~ python
# import foo ? nein.
import rvo.foo
~~~

{{< figure src="/uploads/2016/03/breaking.png" >}}

Um diese Schreibweise während dem Development nutzen zu können sollte man sich
ein Virtualenv initialisieren und darin das Paket deployen in dem man arbeitet.

    pip install virtualenv
    virtualenv .venv
    source .venv/bin/activate
    python setup.py develop

Schon klappts auch mit den Nachbarn^WImports

### Changelog

Changelog schreiben. Meh. [gitchangelog](https://github.com/vaab/gitchangelog)
hilft. Die kleine Software generiert automatisch ein Changelog aus den `git`
Commits. `gitchangelog` nutzt dafür Regexes um die Commits zu kategorisieren
und git tags um die Versionen zu kennzeichnen. Alles Dinge, die ich sowieso
schon tue.

~~~ python
$ vim .gitchangelog.rc
section_regexps = [
  ('Feature',
    [ r'^[fF]eature\s*:\s*([^\n]*)$' ]
  ),
  ...
]
tag_filter_regexp = r'^[0-9]+\.[0-9]+(\.[0-9]+)?$'
output_engine = mustache("restructuredtext")

$ gitchangelog > CHANGELOG.rst
~~~

Man braucht zwar etwas Disziplin beim committen, aber das ist halbwegs erträglich.

### Help!

Das auf [pypi](http://pypi.python.org) nur ReStructuredText okay ist, ist
so eine andere Geschichte.

Sollte man (so wie ich) vorher ein Markdown Readme gehabt haben, kann man
dieses einfach mit `pandoc` konvertieren.

    pandoc README.md --to rst > README.rst

### Editor

Die beiden größten Faktoren für ein entspanntes Entwicklungsumfeld sind für
mich [syntastic](https://github.com/scrooloose/syntastic) und
[snippets](https://github.com/honza/vim-snippets).

Snippets sind nichts anderes als Autocompletions. Es ist einfach wahnsinnig
komfortabel `try`s und ganze Klassen per `Tab` automatisch generieren zu
lassen. Das verlinkte Repo ist nur der Provider der eigentlichen
Schnippsel. Das Plugin, welches diese Snippets abruft und einfügt, ist
[UltiSnips](https://github.com/SirVer/ultisnips). So ein Snippets Setup
wird aber nie die Intelligenz einer full-blown IDE erreichen, die auch
eigene Funktionen, Variablen oder weiss der Teufel noch was komplettieren
können.

Wer trotz Snippets Code mit Syntax Errors baut (zum
Beispiel ich), greift zu Syntastic. Syntastic zeigt einem (mittels
Pylint und Pydocstyle) Importerrors, fehlende Dokumentation oder falsche
Methodenparameter.

{{< figure src="/uploads/2016/03/vim.png" >}}

Per Default werden diese Checks beim Öffnen und Schreiben des Files
angewendet. Das "entschleunigt" das Programmiererleben so sehr, dass es
nervt. Ich habe das so umgebogen, dass die Überprüfung nur angestoßen wird,
wenn ich es will.

~~~ vim
let b:syntastic_mode = "passive"
let g:syntastic_always_populate_loc_list = 1
let g:syntastic_check_on_open = 0
let g:syntastic_check_on_wq = 0
let g:syntastic_check_on_w = 0
let g:syntastic_enable_signs = 0
let g:syntastic_auto_loc_list = 2
let g:syntastic_loc_list_height = 5
let g:syntastic_aggregate_errors = 1

nmap <silent> <leader>e :SyntasticCheck<CR>:Errors<CR>
~~~

Und das wars auch schon. Mehr braucht es eigentlich nicht. Mein `vim` ist
mittlerweile wieder etwas schmaler geworden. Bis vor kurzer Zeit hatte ich
auch noch die [Tagbar](https://github.com/majutsushi/tagbar) aktiv
geschaltet. Es fehlt mir ein bisschen, alle Klassen, Funktionen, Variablen und
Imports direkt in einer Seitenleiste zu sehen. Aber die Performance lässt
einfach extrem zu wünschen übrig. Das Plugin spawnt erstmal einen
Subprozess, der das ganze File analysiert. Wenn die Tagbar nicht auf
[exubertant-ctags](http://ctags.sourceforge.net) basieren würde, wäre das
vielleicht anders.

Empfehlen kann ich außerdem noch
[vim-python-pep8-indent](https://github.com/hynek/vim-python-pep8-indent),
damit man auch Indents und Formatierung richtig macht. Um die
Versionsverwaltung in den Editor einzubinden kann man noch
[gitgutter](https://github.com/airblade/vim-gitgutter) nutzen (highlightet
geänderte Zeilen) und [fugitive](https://github.com/tpope/vim-fugitive),
womit sich `git` direkt aus `vim` bedienen lässt.

### Versionierung

Versionierung ist auch noch so ein Ding. Wie kann ich meine Software jetzt
möglichst ohne Aufwand um eine Version erhöhen? Wenn man aktuellen Browser
Versionsnummern folgt, passiert das durchaus öfter.

{{< figure src="/uploads/2016/03/versioning.jpg" >}}

Wo will ich meine Versionsinformationen überall haben?

* Commandline mit --version
* In der Help Message
* Im `pip` Paket
* Im Git Changelog

Die stressfreiste Variante ist, sie _nur_ in der setup.py zu definieren und
sie an allen anderen Stellen zu importieren.

~~~ python
setup(
    name='rvo',
    version='10.0.0',
    ...
)
~~~

In der `__init__.py` holt man sie sich via

~~~ python
#!/usr/bin/env python2.7
from pkg_resources import get_distribution
__version__ = get_distribution('rvo').version
~~~

Und kann es an den verschiedensten Stellen einbauen mit

~~~ python
from rvo import __version__
~~~

### PyPI

{{< figure src="/uploads/2016/03/piepie.jpg" >}}

In der `.pypirc` sollte man, bevor man irgendwas tut, erstmal das Repo eintragen.

``` ini
[pypitest]
repository=https://testpypi.python.org/pypi
username=your_username
password=your_password
```

Danach kann das Pakt in die unendlichen Weiten des Internets freigelassen
werden.

    python setup.py register -r pypitest
    python setup.py sdist upload -r pypitest

Nach ein paar (erfolgreichen) Tests, kann man das nochmal wiederholen und
auf PyPI hochladen. Macht einfach einen gedanklichen sed auf
`s#pypitest#pypi#`.

### Tests

Unter anderem auch noch... Tests. Aber das Thema ist so groß und ich noch
so unerfahren, dass ich das mal weg lasse. Will man auf jedenfall haben, da
Tests im Zweifel die einzigen Indikatoren für Paketmaintainer diverser
Betriebssysteme sind, um sicherzustellen, dass die Software auch wirklich
funktioniert.
