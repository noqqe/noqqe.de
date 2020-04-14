---
title: wget
date: 2020-03-09T13:42:29
tags:
- Software
- wget
---

## Download fortfahren

Wenn man einen Download beginnt, kann dieser (sofern vom Server unterstützt)
wieder fortgesetzt werden.

<!--more-->

``` bash
wget https://ftp.fau.de/kiwix/zim/wikipedia/wikipedia_de_all_maxi_2019-08.zim
<ctrl+c>
wget --continue https://ftp.fau.de/kiwix/zim/wikipedia/wikipedia_de_all_maxi_2019-08.zim
```

## Webpage offline sichern

Für den Offlinegebrauch bietet sich diese wunderbare Zeile an

``` bash
wget -E -H -k -K -p -nd -nH -p -np http://stackoverflow.com/questions/29681631/tomcat-database-connection-derby
```

Sie lädt eine Webseite mit samt allen Mediatypes und Abhängigkeiten lokal
herunter.
