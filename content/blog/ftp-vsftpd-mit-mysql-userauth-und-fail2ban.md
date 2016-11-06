---
date: 2010-02-24T19:40:51+02:00
comments: true
title: FTP | vsftpd mit MySQL-Userauth und fail2ban
aliases:
- /blog/2010/02/24/ftp-vsftpd-mit-mysql-userauth-und-fail2ban
- /archives/904
categories:
- Development
- Hardware
- Linux
- ubuntuusers
tags:
- daemon
- debian
- Fail2ban
- ftp
- logfile
- mysql
- ubuntu
- ubuntuusers
- vsftpd
---

Ein Kollege aus dem lokal vertretenen Eishockey-Hobbyverein hatte eine
kleine Page mit HTML gebastelt und wollte diese irgendwo hosten.  Hier
würde sich von den Mitgliedern um den Informationsfluss gekümmert und da
ich selbst öfters an den Spielen teilnehme, half ich natürlich gerne. Ich
benutzte bis dato allerdings nie FTP und hatte auch keinen FTP-Server
installiert. "Wenn dann schon richtig" war meine Intention. Über ein
[How-To auf HowtoForge.com](http://www.howtoforge.de/howto/virtual-hosting-mit-vsftpd-und-mysql-auf-debian-etch/)
richtete ich einen [vsftpd](http://vsftpd.beasts.org/) mit mysql-userauth
ein. Das war innerhalb 15 Minuten geschafft. FTP-Server lief wunderbar und
die (noch dürftige) Site ist auch fast online. Mir gefiel die
Auth-Möglichkeit über MySQL.

Nichtsahnend durchforstete ich heute Morgen die Logfiles meiner Zwetschge.
vsftpd-Logfiles innerhalb 15 Stunden relativ voll. Irgendwas war faul.
Nachdem ich die fehlerhafte Konfiguration des logrotated ausschliessen
konnte sah ich mir die Logs mal an.

```
CONNECT: Client "xxx"
[Administrator] FAIL LOGIN: Client "xxx"
[Administrator] FAIL LOGIN: Client "xxx"
[Administrator] FAIL LOGIN: Client "xxx"
CONNECT: Client "xxx"
[Administrator] FAIL LOGIN: Client "xxx"
[Administrator] FAIL LOGIN: Client "xxx"
[Administrator] FAIL LOGIN: Client "xxx"
…
```

Ich zählte nicht, wie oft genau. Jedenfalls zu oft um von fehlerfreier
Konfiguration meines fail2ban ausgehen zu können. Außerdem ist es
beachtlich wie schnell Bots einen existierenden FTP-Server ausmachen
können. Was solls. Zur Erinnerung:
[Fail2ban](http://www.fail2ban.org/wiki/index.php/Main_Page) verbietet
(anhand Logfileanalyse) Clients die Verbindung, wenn sie  zu oft
abgewiesene Verbindungsversuche gestartet haben. Sprich: Zu viele falsche
Passwörter. [Stichwort Bruteforce-Attacke](http://de.wikipedia.org/wiki/Brute-Force-Methode)

Dies veranstaltet fail2ban mit einem Configfile (/etc/fail2ban/jail.local)
und Filtern (/etc/fail2ban/filters.d/*). Ich habe länger überlegt, Config
erneuert, fail2ban-server neu gestartet bis mir kam warum die übermäßig
vorhandenen failed-logins meines FTP-servers nicht geblockt wurden. Die
Ausgabe im Loggingfile hatte sich durch die Umstellung auf MySQL geändert
und [fail2ban](http://www.fail2ban.org/wiki/index.php/Vsftpd) greift nicht
mehr:

```
auth.log(Standard): Jan 23 14:04:14 vsftpd: pam_unix(vsftpd:auth): authentication failure; logname= uid=0 euid=0 tty=ftp ruser=Administrator rhost=xxx
---
auth.log(mysqlauth): Feb 24 12:33:29 zwetschge vsftpd: pam_mysql - SELECT returned no result.
```

Nach etwas erfolglosen herumgegoogle und anderem, beschloss ich die RegExp
für den neuen Filter selbst zu konfigurieren. Der neue Filter basiert
nichtmehr auf dem auth.log sondern auf dem vsftpd.log(im jail.local-File
vermerken!). fail2ban bietet eine wunderschöne Möglichkeit selbstgecodete
Filter auszuprobieren. Via fail2ban-regexp wird ein zu filternder
Logeintrag auf ein regexp geprüft.

```
fail2ban-regexp 'logeintrag' 'regexp zum logeintrag'
```

[/uploads/2009/09/011](/uploads/2009/09/011)

In filters.d: die die Regular-Expression des Zugriffs für das
StandardLogfile ersetzen:

```
alt:auth.log(stdregexp): failregex = vsftpd: (pam_unix) authentication failure; .* rhost=<HOST>(?:s+user=S*)?s*$
---
neut:vsftpd.log(mysqlregexp): failregex = .* FAIL LOGIN: Client "<HOST>"$
```


Fail2ban neu starten, glücklich sein.  Um zukünftigen Usern diesen Schritt
zu erleichtern habe ich natürlich die [Änderungen unter das How-To kommentiert](http://www.howtoforge.com/vsftpd_mysql_debian_etch_p2#comment-22234).
Awating Moderation btw.
