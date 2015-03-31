---
layout: post
title: "Hartes Discard Protokoll"
date: 2014-11-04T19:30:00+02:00
comments: true
categories:
- OpenBSD
- inetd
- discard
- devnull-as-a-service
- port
- connection
- osbn
---

Ich habe verloren. Als ich gestern das OpenBSD unter
[devnull-as-a-service.com](http://devnull-as-a-service.com) upgegraded habe,
musste ich die Kiste durchbooten.

Nachdem ich das Discard Protokoll über den OpenBSD inetd auf Port 9 aktiviert
habe, haben ein paar Leute dort dauerhaft Connections geöffnet.

{% img center /uploads/2014/11/daas-inetd.png %}

Ich dachte, mal sehen wer längern kann. Client oder Server. Stellte sich heraus:
die Clients.

Aber da ich ein guter Verlierer sein kann, Respekt und Gratz.
