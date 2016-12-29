---
aliases:
- /blog/2014/03/24/privacy-plus-plus
comments:
- author: tux.
  content: "Und warum genau l\xE4dst du immer noch Twitterwidgets nach?"
  date: '2014-03-24T22:17:57.474257'
- author: noqqe
  content: wo denn?
  date: '2014-03-24T22:26:23.665112'
- author: tux.
  content: "Unter dem &lt;/footer>.\n\n(Hab' das nun nur \xFCberflogen, nicht analysiert,
    Verzeihung.)"
  date: '2014-03-24T22:29:29.452432'
- author: noqqe
  content: "ah. sollte Discomnect Plugin ausmachem wenn ich nach sowas suche. \n\nDas
    kommt von Octopress, da konnte man Tweets einbetten. \n\ndanke f\xFCr den Hinweis.
    Morgen dann. "
  date: '2014-03-24T22:35:43.769282'
- author: tux.
  content: 'Octopress ist furchtbar. ;-)


    Keine Eile. Rennt nicht weg.'
  date: '2014-03-24T22:56:51.097115'
- author: Martin
  content: \o/
  date: '2014-03-24T23:39:45.186908'
- author: edik
  content: Die Aenderungen sind wirklich zu begruessen. Es kann nicht sein, dass man
    im auf dezentral getrimmten Internet beim Besuch der meisten Webseiten staendig
    zu Servern von Google und co. verbindet. So kann ich in letzter Zeit nur schlecht
    zu cdn.sstatic.net verbinden. Seiten wie Stackoverflow seh ich dann oldschoolmaessig
    ohne CSS und JavaScript.
  date: '2014-03-25T00:12:00.805500'
- author: Georg
  content: "Wie ich sehe, nutzt du hier isso als Kommentarsystem, oder?\nWie wird
    dieses denn in Octopress integriert?\nDas Fehlen einer lokalen Disqus-Alternative
    hat mich bisher immer davon abgehalten, Blogcompiler \xFCberhaupt mal auszuprobieren
    ;-)"
  date: '2014-03-25T11:02:11.207727'
- author: Georg
  content: "Sobald man mal GStatic und all die anderen Google-Seiten blockiert (habe
    ich mit Ausnahme von Youtube mal gemacht), sieht nahezu das gesamte Netz oldschoolm\xE4\xDFig
    aus. Sehr deprimierend, dass sogar selbsternannte \"Professionals\" ohne Google
    nicht mal einen h\xFCbschen Webseitenhintergrund basteln k\xF6nnen."
  date: '2014-03-25T11:03:47.802195'
- author: noqqe
  content: "Integration in Octopress funktioniert eigentlich sehr straight forward.
    Irgendwo im Header die embed.min.js inkludieren, iwo im Footer deines Post Templates
    den isso-thread definieren. Fertig. \n\nGibt aber auch Anleitungen wo das jemand
    mal beschrieben hat:\n\nhttp://blog.alboh.de/blog/2014/01/31/octopress-mit-isso-als-alternative-zu-disqus/\n\nF\xFCr
    die entsprechenden (aktuellen) Docs w\xFCrde ich dich an http://posativ.org/isso/docs/
    verweisen. Bei Fragen auch jederzeit in #isso auf Freenode. H\xE4nge da auch rum. "
  date: '2014-03-25T11:07:49.104566'
- author: Georg
  content: "Dankesch\xF6n. :) Dann fuchse ich mich da mal rein. Auf Wordpress habe
    ich keine Lust mehr...zu langsam, zu wenig Neues zu entdecken. ;-)"
  date: '2014-03-25T13:53:55.535548'
- author: noqqe
  content: 'btw: done.'
  date: '2014-03-26T09:10:49.924155'
date: '2014-03-24T18:55:00'
tags:
- development
- jquery
- github
- privacy
- tls
- pinboard
- webfonts
- openssl
- selfhosted
- blog
- ssl
- google fonts
- wget
- startssl
title: Privacy++
---

"Machste halt mal kurz SSL am Webserver an". So einfach ists halt leider
nicht.  Ich habe das so lange nicht in getan, weil weder Logins vorhanden
sind noch geheimer Content publiziert wird.  Wer möchte kann aber jetzt
[https://noqqe.de](https://noqqe.de) benutzten.

{{< figure src="/uploads/2014/03/wedontsell.jpg" >}}

Folgende Dinge haben sich geändert:

* Google Web Fonts now selfhosted (mit [Clemens Skript](https://neverpanic.de/blog/2014/03/19/downloading-google-web-fonts-for-local-hosting/))
* Hart kodierte Links mit `sed` auf relative Links umgestellt. Siehe
  [RFC3986](https://tools.ietf.org/html/rfc3986#section-4.2)
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

Gerade die letzten beiden Punkte waren garnicht so einfach. Für alles was
ich sonst so hoste benutze ich meistens die Domain n0q.org. Nur leider
möchte mir [StartCom](https://startssl.com) kein Zertifikat für diese
Domain ausstellen, da sie fürchten ich könnte eine Fakesite für
[noq.org](http://noq.org) hosten.  Valide Begründung aber doof für mich.