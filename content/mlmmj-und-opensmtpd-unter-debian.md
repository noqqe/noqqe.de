---
layout: post
title: "mlmmj und OpenSMTPD unter Debian"
date: 2014-05-29T11:24:00+02:00
comments: true
categories:
- OpenSMTPD
- Mailingliste
- Debian
- mlmmj
- jessie
- ubuntuusers
- osbn
---

Für die Mailingliste der [k4cg](http://k4cg.org) zieht demnächst um. Weswegen
ich mich mit einem dementsprechenden Setup auseinander setzen wollte.

### mlmmj

Bisher läuft die ML mit [mlmmj](http://mlmmj.org).
Ich kannte das gute Stück vorher garnicht, macht aber einen sehr
netten Eindruck. Einfach gestrickt, wenig Overhead, Plaintext Files ohne viel
TamTam.
Bei der Konfiguration kann man sich ohne Bedenken von `mlmmj-make-ml` leiten lassen.

``` bash
$ sudo aptitude install mlmmj
$ mlmmj-make-ml
```

Nachdem die selbsterklärende Installation abgeschlossen ist, noch in `/etc/aliases`
eine Pipe einfügen für den entsprechenden User.

``` bash
k4cg:     "|/usr/bin/mlmmj-receive -L /var/spool/mlmmj/k4cg"
```


### OpenSMTPD

{{< figure src="/uploads/2014/05/opensmtpd.png" >}}

Den aus dem OpenBSD Umfeld entstandenen [OpenSMTPD](http://opensmtpd.org) wollte ich mir schon länger ansehen.
Für Postfix läge die mlmmj Konfigurationsanleitung zwar bei, aber hat ja irgendwie auch jeder und ist
für unsere Zwecke viel zu bloated.

``` bash
$ sudo aptitude install opensmtpd
```

Die einzige Config, die es bei OpenSMTPD gibt, `/etc/smtpd.conf` liesst sich schön im Stil von `pf`.

``` bash
# interfaces to listen
listen on localhost
listen on eth0

# if you edit the file, you have to run "smtpctl update table aliases"
table aliases file:/etc/aliases

# recieve mails only for mailinglist domain and from any host
accept from any for domain "k4cg.org" alias <aliases> deliver to mbox

# other local mailboxes
accept from local for local alias <aliases> deliver to mbox

# allow to sent out mails to subscribed users
accept from local for any relay
```

Habe etwas mit dem smtpd herumgespielt, gefällt mir richtig gut.
Minimal gehalten und selbsterklärend. Danach noch das newaliases Pendant `smtpctl update table
aliases` ausführen. Ansehen will man sich auch mal `smtpctl monitor` &lt;3

### Tests mit Swaks

Gerade bei Mailsetups sind die Testszenarien etwas unschön abzuarbeiten. Das
Swiss-Army-Knife-for-SMTP `swaks` hilft einem, das Zeug nicht jedesmal
selbst über `telnet` eintippern zu müssen.

``` bash
$ swaks --server 56.78.90.46 --to k4cg+subscribe@k4cg.org --from noqqe@example.org
=== Connected to 56.78.90.46.
 -> MAIL FROM:<noqqe@example.org>
<-  250 2.0.0: Ok
 -> RCPT TO:<k4cg+subscribe@k4cg.org>
<-  250 2.1.5 Destination address valid: Recipient ok
 -> DATA
<-  354 Enter mail, end with "." on a line by itself
 -> Date: Wed, 28 May 2014 23:12:58 +0200
 -> To: k4cg+subscribe@k4cg.org
 -> From: noqqe@example.org
 -> Subject: test Wed, 28 May 2014 23:12:58 +0200
 -> X-Mailer: swaks v20130209.0 jetmore.org/john/code/swaks/
 ->
 -> This is a test mailing
 ->
 -> .
<-  250 2.0.0: ac3d1ccf Message accepted for delivery
 -> QUIT
<-  221 2.0.0: Bye
```

Nach Test für subscribe/unsubscribe sollte man ebenfalls überprüfen, ob man nicht unter Umständen ein OpenRelay konfiguriert hat.

``` bash
$ swaks --server 56.78.90.46 --to irgendwer@gmail.com --from noqqe@example.org
[...]
 -> RCPT TO:<irgendwer@gmail.com>
<** 550 Invalid recipient
 -> QUIT
<-  221 2.0.0: Bye
```

