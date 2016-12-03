---
aliases:
- /blog/2009/02/24/mail-via-telnet-persohnlich-abliefern
- /archives/521
date: '2009-02-24T19:32:32'
tags:
- zwetschge
- eee-pc
- telnet
- hardware
- linux
- mail
- email
- mailserver
title: "Mail | via Telnet pers\xF6hnlich abliefern"
---

Als ich vor kurzem meine Anmeldung zur Abschlussprüfung über die Post
abgeschicken wollte, hab ich beim eintragen der Adresse gemerkt das die IHK
nur 2 Häuser neben der Post ist bei der ich den Brief aufgeben wollte =)
Also hab ich den Brief kurzerhand selbst zur IHK gebracht.  So ähnlich
funktioniert das auch mit E-Mails.

Via Telnet-Verbindung zum Port 25 des Mailservers:

```
telnet zwetschge.org 25
```

Identifizierung via helo:

```
helo hostname
```

Absender übermitteln

```
mail from: flo@noqqe.de
```

Empfänger übermitteln

```
rcpt to: root@zwetschge.org
```

Mail-Inhalt übermitteln

```
data
text
text text
```

Inhalt stoppen:

```
.
```

Message fertig übermittelt.Beenden:

```
quit
```

Ein netter alternativ-Weg zum versenden von Mails. Wird hauptsächlich zu
Testzwecken von Mailservern verwendet. Wenn jemand es probieren will :)
Meine email habt ihr ja jetzt :)