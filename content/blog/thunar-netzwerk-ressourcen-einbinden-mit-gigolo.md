---
aliases:
- /archives/1109
- /blog/2010/07/16/thunar-netzwerk-ressourcen-einbinden-mit-gigolo
comments:
- author: Darcie Companie
  content: <p>I am very thankful to this topic because it really gives useful information
    `~'</p>
  date: '2011-01-20T01:05:47'
- author: DuKane
  content: <p>Hallo Florian,</p><p>ich stand vor dem gleichen Problem, bin es aber
    etwas anders angegangen. <a href="http://dukuntu.blogspot.com/2011/10/xubuntu-1110-and-deja-dup.html"
    rel="nofollow">http://dukuntu.blogspot.com/2011/10/xubuntu-1110-and-deja-dup.html</a></p>
  date: '2011-10-23T22:29:49'
date: '2010-07-16T14:39:16'
tags:
- ressourcen
- gigolo
- network
- linux
- thunar
- netzwerk
- gio
- ubuntu
- xubuntu
title: Thunar | Netzwerk-Ressourcen einbinden mit Gigolo
---

Xubuntu hat ein schönes Tool von Haus aus dabei, welches sich
[Gigolo](http://www.uvena.de/gigolo/index.html) nennt. Gigolo soll dazu
dienen Netzwerk-Ressourcen zu verwalten und integrieren. Nun, Ressourcen
einbinden funktioniert wunderbar. Nur sollten sich Diese auch öffnen
lassen. Irgendwie. Oder auch nicht.

Gigolo möchte also [Thunar](http://thunar.xfce.org/index.html) zum Öffnen
dieser Orte benutzen (voreingestellt). Allerdings unterstützt Thunar in der
in Xubuntu vorliegenden Version noch keine Netzwerk-Ressourcen. In diversen
Threads bei UbuntuUsers fand ich auch Lösungsansätze die empfahlen Nautilus
zu installieren. Aber ich finde Thunar schön, weshalb ich mich gegen diesen
Ansatz entschied. Nachdem ich mich dann (erneut) 20 min durchs Netz
gegreppt habe, fand ich eine Alternative.

Die aktuelle Entwicklerversion von Thunar setzt auf Gio auf. Dem selben
Manager den auch Nautilus für Netzwerksachen nutzt.

Zu finden ist diese unter:
[https://launchpad.net/~xubuntu-dev/+archive/ppa](https://launchpad.net/~xubuntu-dev/+archive/ppa)

Neue Sources in /etc/apt/sources.list hinzugefügt

```
deb http://ppa.launchpad.net/xubuntu-dev/ppa/ubuntu lucid main
deb-src http://ppa.launchpad.net/xubuntu-dev/ppa/ubuntu lucid main
```

und Thunar updaten. Gigolo kann mittels Thunar die eingebundenen Ressourcen
öffnen. Gigolo funktioniert. Tada.

Allerdings interessiert es mich trotzdem, ob ich einfach nur Fehler im
Handling mache oder da seitens Xubuntu etwas Abstimmungstechnisches schief
lief. Immerhin ist es schon die zweite Xubuntu Installtion bei der dieser
Umstand eintritt.
