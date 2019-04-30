---
title: find
date: 2014-02-20T17:12:57
tags:
- Software
- find
---

Files finden die älter als 3 Tage sind

    find . -type f -mtime +3

Files finden die jünger als 3 Tage sind

    find . -type f -mtime -3

Ordner finden welche vom User und Other schreibbar sind

    find . \( ! -perm -o+w -and -type d \) -print

Löschen via Gnu Tools

    find . -type f -delete

Löschen via xargs

    find . -type f -print0 | xargs -0 rm

Löschen via exec (wenns nicht anders geht). Gefährlich mit Spaces!

    find . -type f -exec rm {} \;
