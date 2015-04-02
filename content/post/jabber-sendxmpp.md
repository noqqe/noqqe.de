---
date: 2009-05-09T19:28:30+02:00
type: post
slug: jabber-sendxmpp
comments: true
title: Jabber | sendXMPP
aliases:
- /archives/606
categories:
- Development
- Hardware
- Linux
tags:
- backup
- jabber
- jabber.ccc.de
- jabber.org
- timerobot
- unpack
- xmpp
- zwetschge
- zwetschge.org
---

Bis noch für kurzem hielt ich Jabber und das XMPP Protocol eher für eine
Art OpenSource -ICQ-Ersatz für Geeks. Mittlerweile... bin ich da immernoch
der selben Meinung :D Allerdings hab ich in den letzten paar Tagen erkannt
wie toll Jabber sein kann. Aus diesem Anlass hier einmal kurz vorgestellt:

sendxmpp ([http://sendxmpp.platon.sk/](http://sendxmpp.platon.sk/))

sendxmpp ist dazu da Jabbernachrichten über bestehende Konten per shell zu
versenden. Syntax mäßig funktioniert das ganze genauso wie beispielsweise
bei mail.

  * Zu haben ist das ganze im Debian Repo ganz normal. (apt-get install
    sendxmpp)
  * Um Nachrichten verschicken zu können wird zuerst ein bestehendes
    Jabber-Konto benötigt. (Ich hab zwetschge.org mal ein Jabberkonto bei
    jabber.org gesponsort)
  * Als nächstes mit den Konto-Daten eine Config erstellen die sehr einfach
    aufgebaut ist:
  * Das wars eigentlich schon. Nun kann versendet werden:

    echo "Hello - Jabber sendxmpp Test" | sendxmpp empfänger@jabber.org

Dabei gibts noch die schönsten Parameter (in der MAN-Page vermerkt) wie
Resource / Subject and so on.  Im ersten Moment klingt das als Jabberclient
recht kompliziert. Aber für mich als Server-Inhaber ist das ganze sehr
sinnvoll. Nur so aus Spaß hab ich eine Datei erstellt mit folgendem Inhalt:


```
 /usr/bin/jabberscript:
echo $1 | sendxmpp -f /root/.sendxmpprc -r zwetschge-generated noqqe@jabber.ccc.de
```

so kann ich in meinen backupscripten/timerobot/cronjobs/unpack mit dem
Aufruf "jabberscript "Backup erfolgreich" (oder eben individueller Inhalt)"
Mir Jabbernachrichten auf mein Konto senden. Nur so zum Spaß lass ich mir
jeden Morgen per cronjob von zwetschge einen schönen Tag wünschen und
andere Scherze - einfach weil ichs kann :) .

<3 Jabber/XMPP

PS: Hilfe.. mein PC spricht mit mir O_o
