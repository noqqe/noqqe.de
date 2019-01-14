---
comments:
- author: Anonymous
  content: Man kann OpenBSD also durch discard DoSsen???
  date: '2015-02-04T22:07:46.173307'
date: '2014-11-04T17:30:00'
tags:
- openbsd
- network
- devnull-as-a-service
- connection
- discard
- inetd
- port
title: Hartes Discard Protokoll
---

Ich habe verloren. Als ich gestern das OpenBSD unter
[devnull-as-a-service.com](http://devnull-as-a-service.com) upgegraded habe,
musste ich die Kiste durchbooten.

Nachdem ich das Discard Protokoll über den OpenBSD inetd auf Port 9 aktiviert
habe, haben ein paar Leute dort dauerhaft Connections geöffnet.

{{< figure src="/uploads/2014/11/daas-inetd.png" >}}

Ich dachte, mal sehen wer längern kann. Client oder Server. Stellte sich heraus:
die Clients.

Aber da ich ein guter Verlierer sein kann, Respekt und Gratz.
