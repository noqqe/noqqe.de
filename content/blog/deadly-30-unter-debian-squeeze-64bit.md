---
comments:
- author: Michael
  content: "<p>Hab mir das auch gerade gekauft, ist ja ganz lustig :)<br>Danke f\xFCr
    deine Anleitung, ein Problem besteht bei mir allerdings: Als normaler User \xF6ffnet
    sich beim Start so ein standard-GTK-Errorfenster mit \u201Ecan't open file '/tmp/<a
    href=\"http://mdm9867.so\" rel=\"nofollow\">mdm9867.so</a>' (error 13: Permission
    denied)\u201C irgendeine Idee was da faul ist? Wenn man es frevelhafterweise mit
    Rootrechten ausf\xFChrt geht es.</p>"
  date: '2012-05-07T09:04:21'
- author: noqqe
  content: "<p>H\xF6rt sich doch an als w\xFCrdest du als Benutzer nicht in /tmp/
    schreiben d\xFCrfen, oder? :) Kontrollier doch mal die Permissions dort.\_</p>"
  date: '2012-05-07T11:36:12'
- author: Michael
  content: "<p>\_Interessant. Hatte tats\xE4chlich keine Schreibrechte mehr auf /tmp.
    Nach dem ich die bisher immer hatte und auch Dateien von meinem User drin waren
    hatte ich an ein generelles Problem garnicht gedacht. Keine Ahnung was da schiefgelaufen
    ist. Jetzt funktioniert alles bestens :)</p>"
  date: '2012-05-07T13:02:25'
- author: Michael
  content: "<p>(die L\xF6sung war nat\xFCrlich ein chmod 1777 /tmp)</p>"
  date: '2012-05-07T13:05:47'
- author: oz123
  content: <p>Hmm.... ok... danke, dein Post hat doch geholfen... Deadly 30 lauft
    jetzt unter mein 64bit Wheezy.</p>
  date: '2012-05-12T19:48:02'
date: '2012-05-06T19:51:00'
tags:
- web
- shooter
- 2d
- game
- linux
- debian
title: Deadly 30 unter Debian Squeeze 64bit
---

Gerade von [Deadly 30](http://www.deadly30.com/index.php) gelesen. Linux Version
ist fertig. Hurra!

Als ich den Trailer gesehen habe musste ich sofort an die [Metal Slug
Reihe](http://de.wikipedia.org/wiki/Metal_Slug) denken, die ich total mag.

Egal, es ist Sonntag Abend und hatte die 4 Euro übrig. Ich hatte aber bereits
erwartet, dass es nicht nur auf ein `./deadly\ 30` hinauslaufen würde. Damit lag
ich auch richtig.

```
$ ./Deadly\ 30
./Deadly 30: error while loading shared libraries: libgstreamer-0.10.so.0:
cannot open shared object file: No such file or directory
```

Mh. Da ich nicht zwischen 32 und 64bit wählen durfte wahrscheinlich 32bit?

```
$ file Deadly\ 30
Deadly 30: ELF 32-bit LSB executable, Intel 80386, version 1 (SYSV), dynamically
linked (uses shared libs), for GNU/Linux 2.6.15, stripped
```

Also die nötigen 32bit Libraries runterladen und einbauen.

## 1. /usr/lib32/ Backup erstellen

```
$ tar cfvz $HOME/usr-lib32.tar.gz /usr/lib32
```

## 2. libgstreamer installieren

* [http://packages.debian.org/squeeze/libgstreamer0.10-0](http://packages.debian.org/squeeze/libgstreamer0.10-0)

```
$ cd /tmp
$ wget http://ftp.de.debian.org/debian/pool/main/g/gstreamer0.10/libgstreamer0.10-0_0.10.30-1_i386.deb
$ ar -x libgstreamer0.10-0_0.10.30-1_i386.deb
$ tar xfvz data.tar.gz
$ sudo cp -av /tmp/usr/lib/* /usr/lib32/
```

## 3. libgstreamer-plugins installieren

* [http://packages.debian.org/squeeze/libgstreamer-plugins-base0.10-0](http://packages.debian.org/squeeze/libgstreamer-plugins-base0.10-0)

```
$ cd /tmp
$ wget http://ftp.de.debian.org/debian/pool/main/g/gst-plugins-base0.10/libgstreamer-plugins-base0.10-0_0.10.30-1_i386.deb
$ ar -x libgstreamer-plugins-base0.10-0_0.10.30-1_i386.deb
$ tar xfvz data.tar.gz
$ sudo cp -av /tmp/usr/lib/* /usr/lib32/
```

Das ist zwar alles immer total unschön, aber es funktioniert. Außerdem, was tut
man nicht alles für ein bisschen 2D Shooter und Zombies?