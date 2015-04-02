---
date: 2010-09-29T18:34:51+02:00
type: post
slug: twitter-oauth-twidge-reanimieren
comments: true
title: Twitter OAuth | Twidge reanimieren
aliases:
- /archives/1262
categories:
- Debian
- General
- Linux
- PlanetenBlogger
- Ubuntu
tags:
- API
- auth
- debian
- Key
- OAuth
- twidge
- twitter
- ubuntu
---

Da Twitter ja glorreicher Weise das Auth-Verfahren [für Dritt-Software umgestellt](http://disfunctions.de/tutorials/twitter-oauth-gwibber-upgraden/) hat, funktioniert mein Twitter-Logfile nun nicht mehr. Wie so manch anderer Client oder Schnittstelle. Die in Debian-Stable ausgelieferte Version 0.99.3 von Twidge unterstützt noch kein OAuth. Das mag jetzt niemandes Schuld sein. Denn warscheinlich liegt es weder am [Author der Software](http://github.com/jgoerzen/twidge), noch am [Packaging](http://packages.debian.org/squeeze/twidge). Wohl eher an der Zeit die ein Debian Paket braucht, um von Testing o. Unstable in Stable zu gelangen. Was auch gut so ist.

Da ich Twidge aber auch außerhalb des Logfiles sehr gerne benutze, behob ich das Problem durch Installieren der Testing-Version. Zu finden unter http://packages.debian.org/squeeze/twidge

```
$ twidge -v
[...]
This is Twidge, version 0.99.4.  Copyright (c) 2008 John Goerzen
```


```
$ wget http://ftp.de.debian.org/debian/pool/main/t/twidge/twidge_1.0.5_i386.deb
```


```
$ dpkg -i twidge_1.0.5_i386.deb
Vorbereiten zum Ersetzen von twidge 0.99.4 (durch twidge_1.0.5_i386.deb) ...
Entpacke Ersatz für twidge ...
dpkg: Abhängigkeitsprobleme verhindern Konfiguration von twidge:
twidge hängt ab von libffi5 (>= 3.0.4); aber:
Paket libffi5 ist nicht installiert.
dpkg: Fehler beim Bearbeiten von twidge (--install):
Fehler traten auf beim Bearbeiten von:
twidge
```


```
$ aptitude install libffi5
Die folgenden NEUEN Pakete werden zusätzlich installiert:
libffi5
Die folgenden teilweise installierten Pakete werden konfiguriert:
twidge
Richte libffi5 ein (3.0.7-1) ...
Richte twidge ein (1.0.5) ...
Aktueller Status: 0 kaputt [-1].
```


Die aktuelle Version von Twidge wäre also somit installiert. Zwar mit dpkg, aber wenigstens wird nachher bei neuereren Versionen geupdated :). Twidge stellt jetzt einen (wie ich finde) hervorragend gelösten Konfigurationsweg für die neue OAuth.

```
$ twidge setup
Please cut and paste this URL and open it in a web browser
Click Allow when prompted.  You will be given a numeric
key in your browser window.  Copy and paste it here.
```
`https://api.twitter.com/oauth/authorize?oauth_token=xxx
Authorization key: 123456789
Successfully authenticated!
Twidge has now been configured for you and is ready to use.
```


Hat irgendwie fast ein bisschen was von Wizard ;). Im Endeffekt sind es aber nur ein paar kleine Schritte:



	
  * Url öffnen

	
  * "App" autorisieren oder "Erlauben"

	
  * App-Authkey kopieren

	
  * Twidge vorwerfen


Danach kann man das Internet wieder mit dem eigenen daily Nonsense füllen. Mit einem CLI-Client :). Was mich nebenbei noch fasziniert hat: [http://packages.ubuntu.com/search?lang=de&searchon=names&keywords=twidge](http://packages.ubuntu.com/search?lang=de&searchon=names&keywords=twidge) . Es gibt Twidge für Ubuntu also für Jaunty, Karmic und Maverick. Warum nicht für Lucid? Ich konnte es in meinen Quellen in Xubuntu (Lucid) auch nicht finden.
