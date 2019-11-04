---
title: Emacs Cheatsheet
date: 2016-02-16T21:17:51
tags:
- Software
- emacs
---

* C = CTRL
* M = META (ALT auf OSX)

[OSX Keymapping](http://xor.lonnen.com/2013/01/04/emacs-on-osx.html)

Alle Befehle die mit dem gleichen Kürzel anfangen
können einfach durchgehend gedrückt werden! WTF!

#### Generelles

* Exit: c-x c-c
* Kommando abbrechen: c-g
* Neues File öffnen: c-x c-f

* Neuen Buffer erstellen: c-x b `<name>`
* Buffer schliessen: c-x k
* Alle Buffers anzeigen: c-x c-b
* Switch Buffer: c-x b
* Speichern: c-x c-s

#### Cursor bewegen

* Zeile runter: c-n
* Zeile hoch: c-p
* Rechts: c-f
* Links: c-b
* Seite runter: c-v
* Seite hoch: m-v
* Zeile zentrieren: c-l

#### Bearbeitung

* Undo: c-/
* Visual Block: c-space (set marker), move cursor, c-w (löscht alles):w

#### Package Manager

* Name el-get
* Paket installieren: m-x el-get-install

#### Misc

* c-h h - Zeigt lustige UTF-8 Zeug an
