---
title: vim
date: 2015-03-05T08:17:55
tags:
- Software
- vim
---

Ein paar Hilfen von Dingen die ich immer wieder vergesse.

## Commands

* `ZZ` - nicer for closing

## Editing

* `CTRL+V` - Block select
* `gg+=+G` -  Whitespace entfernen
* `<select>+gq` - Block formatieren nach `textwidth`

## Moving

`:help motion` lesen. Sahneschnittchen sind hier.

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

## Maps

Wenn ich mal wieder Keys neu mappen muss

    :help map-overview

## Registers

In Registern werden Dinge aus Clipboards und yank, delete Aktionen gespeichert

    :help registers

Zusammenfassung

> 1. The unnamed register ""
> 2. 10 numbered registers "0 to "9
> 3. The small delete register "-
> 4. 26 named registers "a to "z or "A to "Z
> 5. three read-only registers ":, "., "%
> 6. alternate buffer register "#
> 7. the expression register "=
> 8. The selection and drop registers "*, "+ and "~
> 9. The black hole register "_
> 10. Last search pattern register "/

Um aus verschiedenen OS Clipboards lesen zu können, will man sich auch

    :help clipboard

anschauen

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

## vim-sandwich

Da ich mir das einfach nicht merken kann, muss ich das jetzt hier aufschreiben.

* `saiw"` - Wort umklammern wie "hier"
* `sr"(` - Doublequotes mit `(` ersetzen wie (hier)
* `sd"` - löschen der Umrandung " wie hier
