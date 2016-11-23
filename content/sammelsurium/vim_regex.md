---
title: vim regex
date: 2012-03-04T14:39:47.000000
tags: 
- Software
- Vi
---


* %s Ganzes File
* 1,5 Zeile 1 - 5
* ^,10 Erste 10 Zeilen
* 20,$ Zeile 20 bis Ende

## Sonderzeichen und Krams

Newlines im File durch Komma ersetzen:

     %s/\s*\n/,/g

Newline durch RegEx einfügen:

     %s/`/```^M/

Der Trick dabei ist: STRG+V und STRG+M Drücken. Das macht in das Kommando das Umschaltzeichen ^M rein.
