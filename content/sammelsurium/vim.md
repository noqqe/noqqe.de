---
title: vim
date: 2015-03-05T08:17:55
tags:
- Software
- vim
---

Some of the shortcuts here may refer to my vim configuration on github.
This is far away from being complete.

#### Commands
`ZZ` - nicer for closing

#### Plugin shortcuts
* `-r` - Reload all configurations (after config change)
* `-P` - Cleanup + Install + Update all Plugins
* `-n` - Toggle line numbers
* `-p` - Toggle paste mode
* `-s` - Toggle spell checking
* `-w` - Strip trailing whitespace
* `-g` - Toggle git gutter line markers
* `-e` - Check for syntaxerrors with syntastic

#### Editing
* CTRL+V - Block select
* gg+=+G - Remove all whitespace
* Selektieren+gq - Block formatieren nach textwidth

#### Moving
* `#` - zu nächstem Vorkommen des wortes springen
* `K` - oeffnet Manpage zu dem Wort ueber dem sich der Curor befindet

#### Windows
* `CTRL+ww` - Switch between windows
* `CTRL+ws` - Split windows
* `CTRL+ww` - switch between windows
* `CTRL+wq` - Quit a window
* `CTRL+wv` - Split windows vertically

#### Panes
* ALT+Arrow - Move between Splitpane

#### Files

* `:bn` - next File
* `:bp` - previous File

#### Filetypes

Convert file format

    :set ff=unix
    :set ff=mac

and to switch encoding

    :write ++enc=utf-8

#### RegEx

* `%s`   - Ganzes File
* `1,5`  - Zeile 1 - 5
* `^,10` - Erste 10 Zeilen
* `20,$` - Zeile 20 bis Ende

## Sonderzeichen und Krams

Newlines im File durch Komma ersetzen:

     %s/\s*\n/,/g

Newline durch RegEx einfügen:

     %s/`/```^M/

Der Trick dabei ist: STRG+V und STRG+M Drücken. Das macht in das Kommando
das Umschaltzeichen ^M rein.
