---
title: "Gnome Terminator und Terminus Bitmap Font Rendering"
date: 2012-10-19T19:50:00+02:00
comments: true
tags:
- osbn
- Shell
- Linux
- Terminal
- Bash
- Font
- X11
---

Mir wurde letzten Donnerstag vom
[Datenkleptomanen](http://datenkleptomanie.info) als Terminal Font
[terminus](http://terminus-font.sourceforge.net/) empfohlen. Sogar in
Debian [paketiert](http://packages.debian.org/squeeze/console-terminus)!

Nach der Installation habe ich mich erstmal gefreut. Danach fiel mir auf
das mein CPU irgendwie viel zu tun hat, wenn ich viel Text auf `STDOUT`
mithilfe meines (mir ans Herz gewachsenen) [Gnome
Terminator](http://www.tenshu.net/p/terminator.html) ausgebe.

Habe dann jeweils mit dem alten Font und mit Terminus die selbe Datei
ausgegeben und micht etwas gewundert:

```
$ wc -l /var/log/syslog
946 syslog
$ cat /var/log/syslog
```

Terminus:
{{< figure src="/uploads/2012/10/terminus.png" >}}

Monospace:
{{< figure src="/uploads/2012/10/monospace.png" >}}

Irgendwie scheint die Ausgabe wesentlich länger zu dauern, wenn der Font
Terminus ist. Das ist auch beliebig oft reproduzierbar. Es "ruckelt"
regelrecht.

Das fand ich komisch. Komisch und nervend, weil mir dieser Font gefällt und
ich mich nicht von Terminator trennen möchte.

Zwangsläufig musste ich Terminus aber doch dran glauben. Anscheinend gibt
es einen [Bug](https://bugs.freedesktop.org/show_bug.cgi?id=48395) in
libcairo der für das CPU intensive Bitmap Font Rendering verantwortlich
ist.  Ich werde erstmal warten, bis eine neue Version in meinem Debian
Testing ankommt. In Ubuntu VM gehts wohl :/
