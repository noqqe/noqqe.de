---
title: vim
date: 2015-03-05T08:17:55
tags:
- Software
- vim
---

Some of the shortcuts here may refer to my vim configuration on github.
This is far away from being complete.

## Commands
* `ZZ` - nicer for closing

## Editing
* `CTRL+V` - Block select
* `gg+=+G` -  Whitespace entfernen
* `<select>+gq` - Block formatieren nach `textwidth`

## Moving
* `#` - zu nächstem Vorkommen des Wortes springen
* `K` - öffnet Manpage zu dem Wort über dem sich der Cursor befindet
* `H` - Ganz oben im screen springen
* `M` - Mitte des Screens springen
* `L` - Ganz unten im screen springen

* `mx` - Markierung setzen in aktueller Zeile
* `'x` - Zur markierten Zeile zurückspringen
* `''` - Zurück zur Stelle vorher springen

## Windows
* `CTRL+ww` - Switch between windows
* `CTRL+ws` - Split windows
* `CTRL+ww` - switch between windows
* `CTRL+wq` - Quit a window
* `CTRL+wv` - Split windows vertically

## Panes
* `ALT+Arrow` - Move between Splitpane

## Files
* `:bn` - next File
* `:bp` - previous File

## Filetypes

Konvertieren des Filetyps

    :set ff=unix
    :set ff=mac

Encoding reparieren

    :write ++enc=utf-8

## RegEx

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
