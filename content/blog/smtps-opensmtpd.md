---
title: "SMTPS OpenSMTPD"
date: 2014-12-01T12:02:00+02:00
comments: true
categories:
- osbn
- ubuntuusers
- OpenBSD
- BSD
- Mail
tags:
- openbsd
- neomailbox
- mutt
- smtp
- auth
---
Mein Mailprovider ist mittlerweile ein halbes Jahr
[neomailbox.net](https://neomailbox.net). Seit einiger Zeit hatte ich aber
Probleme beim Einliefern von Mails zum SMTP Server. Mutt resettet beim SMTP mit
CRAM-MD5 über SSL immer wieder die Verbindung. Kein Einliefern möglich.

{{< figure src="/uploads/2014/11/tgvfu.gif" >}}

Eigentlich ist die `.muttrc` ziemlich straight-forward was das betrifft

```
set smtp_url="smtp://user@neomailbox.net"
set smtp_pass="PW"
set ssl_starttls = yes
set smtp_authenticators = "cram-md5"
set ssl_force_tls = yes
```

### Debugging

Also erstmal `openssl` angeworfen um damit zu schauen was
die Serverseite so erzählt. Zuvor aber Username und Password
in BASE64 encodiert vorbereiten:

```
perl -MMIME::Base64 -e 'print encode_base64("passwort");'
perl -MMIME::Base64 -e 'print encode_base64("username");'
```

Dann Verbindung zum SMTPS aufbauen und bisschen Tippern.

```
openssl s_client -connect 5.148.176.58:465
---
220 s3.neomailbox.net ESMTP
ehlo noc.n0q.org
250-s3.neomailbox.net Hello noc.n0q.org [127.0.0.1]
250-SIZE 52428800
250-PIPELINING
250-AUTH CRAM-MD5 PLAIN LOGIN
250-STARTTLS
250 HELP
AUTH LOGIN
334 VXNlcm5hbWU6
XYZABCDEFGHIJ
334 UGFzc3dvcmQ6
ABCDEFGHIJKLMNO
235 Authentication succeeded
MAIL FROM:example@example.org
250 OK
RCPT TO:example@example.org
RENEGOTIATING
depth=1 C = BE, O = GlobalSign nv-sa, CN = AlphaSSL CA - SHA256 - G2
verify error:num=20:unable to get local issuer certificate
verify return:0
```

Ich bin mir bis jetzt nicht sicher ob das renegotiaten nach `MAIL FROM:` normal
ist. Danach war jedenfalls auch meine Plaintext-Session vorbei. Fand ich
komisch. Ich dachte auch ob es vielleicht am `LibreSSL` des OpenBSD auf der
Kiste liegt, die ich benutze. Ein Test mit Debian bewies aber dann das
Gegenteil.

### OpenSMTP als lokaler MTA

So konnte das ja auch nicht bleiben. Mails verschicken können wär
schon schön. Da es sowieso ein OpenBSD ist, auf der mein `mutt` läuft war der
`OpenSMTPD` schon da.

Was folgt ist eine kurze Anleitung, alle Mails an einen remote SMTP Server mit
Authentifizierung weiterzuleiten.

Das `secrets` File erstellen

``` bash
$ echo "neo user:pw" > /etc/mail/secrets
$ chown root:_smtpd /etc/mail/secrets
$ chmod 640 /etc/mail/secrets
$ makemap /etc/mail/secrets
```

und die `smtpd.conf` wie folgt anpassen:

```
listen on lo0
table aliases db:/etc/mail/aliases.db
table secrets db:/etc/mail/secrets.db
accept for local alias <aliases> deliver to mbox
accept for any relay via secure+auth://neo@neomailbox.net:465 auth <secrets>
```

Die `.muttrc` auf localhost umbiegen

```
set smtp_url="smtp://localhost:25"
```

Das hat jetzt nicht nur den Vorteil, dass ich wieder Mails versenden kann. Mir
gefällt auch, dass ich jetzt bei eventualler nicht-Verfügbarkeit des
Provider-SMTPs eine queuende Instanz habe. Ausserdem Logfiles in denen ich
wirklich sehen kann wann eine Mail mein System verlassen hat. Negativ: Eine
Komponente mehr die potentiell kaputt gehen kann.
