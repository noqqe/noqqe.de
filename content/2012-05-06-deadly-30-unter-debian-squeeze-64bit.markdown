---
layout: post
title: "Deadly 30 unter Debian Squeeze 64bit"
date: 2012-05-06T21:51:00+02:00
comments: true
categories:
- ubuntuusers
- Linux
- web
- debian
- code
keywords: "deadly 30, deadly, zombies, linux, indie, debian, squeeze, 64bit,
gstreamer, libgstreamer, plugins, metal slug, metal, slug, 30, game"
---

Gerade von [Deadly 30](http://www.deadly30.com/index.php) gelesen. Linux Version
ist fertig. Hurra!

Als ich den Trailer gesehen habe musste ich sofort an die [Metal Slug
Reihe](http://de.wikipedia.org/wiki/Metal_Slug) denken, die ich total mag.

Egal, es ist Sonntag Abend und hatte die 4 Euro übrig. Ich hatte aber bereits
erwartet, dass es nicht nur auf ein `./deadly\ 30` hinauslaufen würde. Damit lag
ich auch richtig.

{% codeblock %}
$ ./Deadly\ 30
./Deadly 30: error while loading shared libraries: libgstreamer-0.10.so.0:
cannot open shared object file: No such file or directory
{% endcodeblock %}

Mh. Da ich nicht zwischen 32 und 64bit wählen durfte wahrscheinlich 32bit?

{% codeblock %}
$ file Deadly\ 30
Deadly 30: ELF 32-bit LSB executable, Intel 80386, version 1 (SYSV), dynamically
linked (uses shared libs), for GNU/Linux 2.6.15, stripped
{% endcodeblock %}

Also die nötigen 32bit Libraries runterladen und einbauen.

## 1. /usr/lib32/ Backup erstellen

{% codeblock %}
$ tar cfvz $HOME/usr-lib32.tar.gz /usr/lib32
{% endcodeblock %}


## 2. libgstreamer installieren

* [http://packages.debian.org/squeeze/libgstreamer0.10-0](http://packages.debian.org/squeeze/libgstreamer0.10-0)

{% codeblock %}
$ cd /tmp
$ wget http://ftp.de.debian.org/debian/pool/main/g/gstreamer0.10/libgstreamer0.10-0_0.10.30-1_i386.deb
$ ar -x libgstreamer0.10-0_0.10.30-1_i386.deb
$ tar xfvz data.tar.gz
$ sudo cp -av /tmp/usr/lib/* /usr/lib32/
{% endcodeblock %}

## 3. libgstreamer-plugins installieren

* [http://packages.debian.org/squeeze/libgstreamer-plugins-base0.10-0](http://packages.debian.org/squeeze/libgstreamer-plugins-base0.10-0)

{% codeblock %}
$ cd /tmp
$ wget http://ftp.de.debian.org/debian/pool/main/g/gst-plugins-base0.10/libgstreamer-plugins-base0.10-0_0.10.30-1_i386.deb
$ ar -x libgstreamer-plugins-base0.10-0_0.10.30-1_i386.deb
$ tar xfvz data.tar.gz
$ sudo cp -av /tmp/usr/lib/* /usr/lib32/
{% endcodeblock %}

Das ist zwar alles immer total unschön, aber es funktioniert. Außerdem, was tut
man nicht alles für ein bisschen 2D Shooter und Zombies?
