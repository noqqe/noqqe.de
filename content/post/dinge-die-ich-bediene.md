---
categories:
- blog
- bsd
- openbsd
- osbn
- osx
- ubuntuusers
date: 2015-05-06T22:25:17+02:00
tags:
- null
title: Dinge die ich bediene.
---

Ich glaube ich habe mich noch niemals bzgl. meiner privaten Situation mit
Technik so wohl gefühlt wie bisher. Das Equipment, das OS. Macht Spaß und
funktioniert.

Ein bisschen wie bei [usesthis.com](http://usesthis.com) (wtf, wie langweilig
ist bitte das [Setup](http://bruce.schneier.usesthis.com) von Bruce Schneier)
beschreibe ich mal was ich so benutze.

### MacBook Pro 13"

Seit Ende 2014 benutze ich ein MacBook. Mein Zweites. Zwischenzeitlich hatte ich
ein Lenovo Thinkpad x201 was zwar seitens der Hardware super geil war, mir aber
das mit dem Betriebssystem und alles zum Hals raushing. Hatte ich schon erwähnt
das 2015 das Jahr von Linux auf dem Desktop ist?

{{< figure src="/uploads/2015/05/macbook.jpg" >}}

All diese Dinge die bei OS X einfach so angenehm sind. `Safari` synchronisiert
einfach alles zwischen iPhone und anderen MacBooks hin und her. `Airmail2` als
Mailclient ist großartig. Die Kalender &amp; Adressbuch Synchronisation rockt.
Nutze `1Password`, `Spotify`  und seit dem Switch von iPhoto zu `Photos` kann
man das auch wieder benutzen. Ich mag es sehr wie alles sich integriert,
funktioniert und man eigentlich nie irgendwas tun muss. Ganz besonders toll: Die
`iMessage`/SMS/`FaceTime` vom Handy zum OS. Ich hätte nie gedacht, dass ich das so
intensiv benutze. Mit `iMessage` beweist Apple auch noch, dass
benutzerfreundliche Crypto nicht unmöglich ist. Die Meisten wissen es
nichteinmal.

Genug Kuschelkurs, OS X Liebesschwüre gibts schon genug im Netz. Die uralte Software
z.B. `rsync` oder `OpenSSH` ist ein echtes Unding. Auch iTunes. Ich benutze es
einfach nicht. Es fehlt mir nicht.

### OpenBSD CLI VM

Im Stil eines "traditionellen" Arbeitsplatzes an der Konsole,
hab ich es mir angewöhnt alles was ich so für den
Alltag brauche auf einer klein dimensionierten VM irgendwo im Internet zu
hosten.

{{< figure src="/uploads/2015/05/cli.gif" >}}

* ToDo: taskwarrior
* Wiki: cmddocs
* Bookmarks: bm
* IRC: weechat
* AntiSpam: bogofilter
* Mail: mutt
* Journal: jrnl
* RSS: newsbeuter
* Mail-Rules: imapfilter

Extrem praktisch, egal wo ich bin, ich hab meistens alles da.
und mobil sicher weil openbsd

URL OPEN
Klar hat das auch Nachteile, ich kann auf meine Todoliste nicht zugreifen wenn
ich nur mit dem iPhone bestückt im Supermarkt stehe, aber diesen Use-Case habe
ich auch einfach nicht. Mit

### OpenBSD Server

* git Repositories
* Websites
* nicht.parasoup.de

Dinge die mich Nerven gibt es auch hier. Nämlich die fehlende SNI Funktionalität
bei `httpd`, `relayd` und `libTLS`. Somit muss ich bisher noch `nginx` für die
Websites nutzen. Aber das ist nur eine Frage der Zeit.

### FreeNAS

Was früher dein Debian mit `mdcrypt` und &gt;25 LXC Containern war,
ist nun nur noch ein FreeNAS.

*
* Usenet
