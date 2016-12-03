---
date: '2015-05-08T16:25:17'
tags:
- osx
- openbsd
- mutt
- file
- mysql
- imapfilter
- bogofilter
- newsbeuter
- spam
- macbook
- blog
- mail
- irc
- taskwarrior
- mongodb
- safari
- thinkpad
- bsd
- php
- imap
- rss
- jrnl
- weechat
title: Dinge die ich bediene.
---

Ich glaube ich habe mich noch niemals bzgl. meiner privaten Situation mit
Technik so wohl gefühlt wie bisher. Das Equipment, das OS. Macht Spaß und
funktioniert.

Ein bisschen wie bei [usesthis.com](http://usesthis.com) (wtf, wie langweilig
ist bitte das [Setup](http://bruce.schneier.usesthis.com) von Bruce Schneier?)
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

Ich habe es mir angewöhnt alles was ich so für den
Alltag brauche auf einer klein dimensionierten VM irgendwo im Internet zu
hosten.

{{< figure src="/uploads/2015/05/cli.gif" >}}

Das ist extrem praktisch, da ich egal wo ich bin, egal an welchem Rechner ich
sitze immer alles da habe. Software die ich dort auf dem bei
[rootbsd](http://rootbsd.net) gehosteten System nutze ist unter Anderem:

* ToDo: [taskwarrior](http://taskwarrior.org)
* Wiki: [cmddocs](https://github.com/noqqe/cmddocs)
* Bookmarks: [bm](https://github.com/noqqe/bm)
* IRC: [weechat](http://weechat.org)
* Mail: mutt
* AntiSpam: [bogofilter](/blog/2013/10/26/spammer-vs-statistik-mit-bogofilter/)
* Mail-Rules: [imapfilter](https://github.com/lefcha/imapfilter)
* Journal: [jrnl](http://maebert.github.io/jrnl/)
* RSS: [newsbeuter](http://www.newsbeuter.org)

Auf der Maschine befindet sich sonst nichts. Alles läuft unter meinem User,
kein Daemon der lauscht, nichts. Gesichert wird die Kiste mittels `tarsnap`

Klar hat das auch Nachteile, ich kann auf meine Todoliste nicht zugreifen wenn
ich nur mit dem iPhone bestückt im Supermarkt stehe, aber diesen Use-Case habe
ich auch einfach nicht. Mit `newsbeuter` Urls im Browser öffnen ist auch
bescheiden, daher muss ich dort immer klicken. Wenn jemand hierfür eine Lösung
hat, immer her damit.

### OpenBSD Server

Der Normal-Nerd hat natürlich auch Bedürfnisse Dinge zu hosten. Deshalb gibts
eine zweite Maschine, die alle meine Dienste bereitstellt die ich so brauche,
diverse PHP/MySQL Applikationen für den Eigengebrauch.

{{< figure src="/uploads/2015/05/obsd.png" >}}

OpenBSD, brauche ich jetzt nicht erwähnen, ist dafür momentan so mein liebstes
OS. Sicher per default. Die Devs hauen immer wieder allerhand nützliche Sachen
wie lustiges [Crypto für Ping](https://twitter.com/dlgwynne/status/589784636714143745) oder seit neuestem
[privilege separated](http://marc.info/?l=openbsd-cvs&m=142989267412968&w=2) `file`

* Meine privaten `git` Repos mit `gitolite`
* der Blog
* Zwei Instanzen von [nichtparasoup](http://github.com)
* [Isso](https://posativ.org/isso) Kommentarsystem
* [MongoDB](http://mongodb.org)
* und diverse andere Websites

Demnächst kommst vielleicht noch etwas DNS hinzu, was ich dort hoste.

Dinge die mich Nerven gibt es auch hier. Nämlich die fehlende SNI Funktionalität
bei `httpd`, `relayd` und `libTLS`. Somit muss ich bisher noch `nginx` für die
Websites nutzen. Aber das ist nur eine Frage der Zeit.