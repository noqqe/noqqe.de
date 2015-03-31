---
layout: post
title: "Privacy++"
date: 2014-03-24T20:55:00+02:00
comments: true
categories:
- osbn
- privacy
- openssl
- ssl
- jquery
- google fonts
- webfonts
- selfhosted
- startssl
- tls
- wget
- pinboard
- github
---

"Machste halt mal kurz SSL am Webserver an". So einfach ists halt leider nicht.
Ich habe das so lange nicht in getan, weil weder Logins vorhanden sind noch geheimer Content publiziert wird.
Wer möchte kann aber jetzt [https://noqqe.de](https://noqqe.de) benutzten.

{% img center /uploads/2014/03/wedontsell.jpg %}

Folgende Dinge haben sich geändert:

* Google Web Fonts now selfhosted (mit [Clemens Skript](https://neverpanic.de/blog/2014/03/19/downloading-google-web-fonts-for-local-hosting/))
* Hart kodierte Links mit `sed` auf relative Links umgestellt. Siehe [RFC3986](https://tools.ietf.org/html/rfc3986#section-4.2)
* Flattr Button durch statische Variante ersetzt
* Kein jQuery nachladen von extern mehr
* Kein Github nachladen mehr
* Kein Pinboard nachladen mehr
* ~50MB verwaiste Uploads entfernt
* Mit `wget --spider` tote Links entdeckt und korrigiert.
* Extern gehostete Bilder aus früheren Posts in /uploads/ migriert
* SSL noqqe.de Zertifikat erstellt
* Isso Kommentarsystem auf comments.noqqe.de umgezogen und SSL
* Piwik auf analytics.noqqe.de umgezogen und SSL

Gerade die letzten beiden Punkte waren garnicht so einfach. Für alles was ich sonst so hoste benutze ich meistens die Domain n0q.org. Nur leider
möchte mir [StartCom](https://startssl.com) kein Zertifikat für diese Domain ausstellen, da sie fürchten ich könnte eine Fakesite für [noq.org](http://noq.org) hosten.
Valide Begründung aber doof für mich.
