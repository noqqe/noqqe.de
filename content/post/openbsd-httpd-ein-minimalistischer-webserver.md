---
layout: post
title: "OpenBSD httpd"
date: 2014-09-27T14:30:00+02:00
comments: true
aliases:
- /blog/2014/09/27/openbsd-httpd-ein-minimalistischer-webserver
- /blog/2014/09/27/openbsd-httpd
categories:
- OpenBSD
- ubuntuusers
- osbn
tags:
- httpd
- Base
- webserver
---
[Reyk Floeter](https://twitter.com/reykfloeter) hat zuletzt begonnen seinen
[relayd](http://bsd.plumbing) zu forken und einen minimalistischen Webserver
daraus zu bauen. Langfristig soll
[httpd](http://www.openbsd.org/cgi-bin/man.cgi/OpenBSD-current/man8/httpd.8) in
OpenBSD den erst kürzlich in Base gewanderten nginx ersetzen.

Die Hintergründe dazu kann man gut im [BSDNOW Podcast 053](http://www.bsdnow.tv/episodes/2014_09_03-its_hammer_time) nachhören.
Zuerst denkt man so "Was? Noch ein HTTP Daemon?". Zusammengefasst soll der neue httpd
aber genau das werden (und vor allem bleiben) wie nginx angefangen hat. Plain,
Free, minimalistisch, einfach. So wurden auch schon mehrere Diffs/Features vom Entwickler
abgelehnt.

{{< figure src="/uploads/2014/09/httpdps.png" >}}

Konfiguration OpenBSD gemäß sehr straight forward. `pf`/`relayd` like.
Hab mir nen 5.6 Snapshot vom Mirror meiner Wahl besorgt und das Teil mal
ausprobiert.

Nach etwas herumprobieren: grinsen. Comic Sans in den default Error Messages.

{{< figure src="/uploads/2014/09/httpderror.png" >}}

Bin mir nicht sicher ob das so bleibt. Die Config Parameter sind wie bei OpenBSD Software
zu erwarten gut dokumentiert und eingängig.

```
prefork 5

server "default" {
        listen on em0 port 80
        log syslog
        log access default-access.log
        directory auto index
        connection timeout 30
        connection max requests 120
}

server "httpd1.noqqe.de" {
        listen on 192.168.1.14 port 80
        root "/htdocs/httpd1/"
        log syslog
        directory auto index
}

server "httpd2.noqqe.de" {
        listen on 192.168.1.14 port 80
        root "/htdocs/httpd2/"
        log syslog
        directory auto index
        connection timeout 3600
}
```

Noch ist das gute Stück nicht Feautre-Complete bzw. Production-Ready.
Dinge die noch fehlen, aber kommen werden ist zum Beispiel Basic HTTP Auth.
SSL und ein bisschen mit [beeswithmachineguns](https://github.com/newsapps/beeswithmachineguns)
Performance austesten hab ich bisher noch
nicht gemacht. `curl`-`for-loop` mit `time` zum Ausprobieren kann man kaum Performance Test nennen ;)
Demnächst dann vielleicht.

`httpd` wird er ab 5.6 in Base mit nginx koexistieren. Portable Version ist
ebenfalls geplant. Freu mich drauf.
