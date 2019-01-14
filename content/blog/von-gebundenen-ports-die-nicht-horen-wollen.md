---
aliases:
- /blog/2014/11/27/von-gebundenen-ports-die-nicht-h%C3%B6ren-wollen/
comments:
- author: "Rainer M\xFCller"
  content: "Interessant, unter OS X 10.10 Yosemite wird mir so ein Port \xFCbrigens
    als CLOSED angezeigt. Getestet hab ich das mit einem Einzeiler in Python.\n\n<pre>\n$
    python -c 'import socket; s = socket.socket(socket.AF_INET, socket.SOCK_STREAM);
    s.bind((\"\", 8888)); raw_input();'\n\n$ lsof -i -n -P |grep 8888\nPython    97609
    raimue    3u  IPv4 0x9efe4c51c483d339      0t0  TCP *:8888 (CLOSED)\n</pre>\n\nUnter
    Linux kann ich das Problem auch nachvollziehen, ich seh allerdings noch nicht
    mal irgendwas mit \"identify\" im lsof. Eine L\xF6sung kenn ich demnach also auch
    nicht."
  date: '2014-11-27T22:49:25.114504'
- author: noqqe
  content: Mh, CLOSED macht auf jedenfall mehr sinn als garnichts. Danke!
  date: '2014-11-28T13:23:40.156244'
date: '2014-11-27T19:45:00'
tags:
- suse
- socket
- bind
- administration
- linux
- sles
- listen
- debian
- network
title: "Von gebundenen Ports die nicht h\xF6ren wollen"
---

Ich komme immer mal wieder in die Verlegenheit in der Arbeit
[solo](http://timkay.com/solo/solo) zu benutzen.
Um doppelte Läufe von Cronjobs zu verhindern beispielsweise.

Der betroffene Cronjob konnte nun aber schon mehrere Stunden nicht mehr
ausgeführt werden.

``` bash
./solo -port=3005 ./script.sh
solo(3005): Address already in use
```

Scheinbar läuft dieser Prozess noch, vermutlich ist er der eigentliche
Prozess aber auf irgendeine Weise gestorben.  Also machte ich mich auf die
Suche nach dem Prozess, der den Port blockiert.

``` bash
$ netstat -tapn | grep 3005
$ lsof -Pnl +M -i4 | grep 3005
$ fuser 3005/tcp
$ lsof -i :3005
$ socklist | grep 3005
$ cat /proc/net/tcp
$ cat /proc/net/udp
$ ss -pl |grep 3005
```

Man sieht schon, ich hab "ein bisschen was" versucht um das herauszubekommen.
Ich habe schon an `solo` an sich gezweifelt, und danach angefangen zu
verifizieren, ob der Port wirklich in Verwendung ist.

``` bash
$ netcat -lvp 3005
retrying local 0.0.0.0:3005 : Address already in use
```

Ich habe auch mit `strace` mitgehört, was Perl da tut.

```
$ strace -o /tmp/foo ./solo -port=3005 "echo anfang ; sleep 5 ; echo ende"
[...]
> bind(3, {sa_family=AF_INET, sin_port=htons(3005), sin_addr=inet_addr("127.0.0.1")}, 16) = 0
```

Okay. Der Port wird verwendet. Bis mir der Unterschied zwischen `bind()`
und `listen()` auf OS-Ebene wirklich klar wurde verging eine peinlich lange
Zeit. Was bedeutet das innerhalb des Betriebssystems eigentlich. Ich kann
diesen Port/Socket mit den oben genannten Tools garnicht finden. Weil er
lediglich "gebunden" ist aber nicht "hört", nicht für andere zur
Kommunikation bereitsteht.

Das war auch so dieser Moment, in dem ich mir dachte dass mir ein paar
Uni-Vorträge in Betriebssystem-Lehre wohl nicht geschadet hätten.

{{< figure src="/uploads/2014/11/ohlol.gif" >}}

Die einzige Weise mit der ich den Prozess finden konnte war mit einem Hinweis
aus dem Netz. Bei nicht erkennbarem Protokoll werden in `lsof` mit `can't
identify protocol` gekennzeichnet.

```
$ lsof | grep identify
2727            root    3u     sock                0,6      0t0   41804186 can't identify protocol
```

Das war aber auch schon das einzige was mir einfiel. Problem hatte sich dann für
mich erledigt. Prozess beseitigt. Funktionierte wieder.
Sollte jemand aber einen Hint haben, wie man gezielter danach
suchen kann, bitte her damit!
