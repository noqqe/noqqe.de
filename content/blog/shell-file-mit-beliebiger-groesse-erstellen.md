---
date: 2009-06-04T23:11:54+02:00
comments: true
title: Shell | File mit beliebiger Groesse erstellen
aliases:
- /blog/2009/06/04/shell-file-mit-beliebiger-groesse-erstellen
- /archives/627
categories:
- Development
- Linux
tags:
- bash
- dd
- Linux
- shell
---

Manchmal kommt man nicht drum rum. Ich brauche ein File um etwas zu testen.
Das File sollte ca 30 MB Gross sein. Aber wo bekomm ich sowas jetzt her?

``` bash
dd if=/dev/zero of=testfile.dat bs=1M count=30
```

* bs=einheit
* count=zähler der Einheit
* 1M x 30 = 30 MB

