---
aliases:
- /blog/2009/02/24/mail-via-telnet-persohnlich-abliefern
- /archives/521
comments:
- author: Dominik
  content: "<p>Dann versuch mal:</p><p>  EHLO hostname</p><p>und schon spricht man
    ESMTP...</p><p>Sehr hilfreich f\xFCr die Fehlersuche des eigenen Systems sind
    auch:</p><p>  VRFY Email-Adresse  -&gt; \xDCberpr\xFCft ob eine Adresse g\xFCltig
    ist.<br>  EXPN Email-Adresse  -&gt; Expandiert eine Email-Adresse, so kann man
    Umleitungen testen.</p><p>Leider m\xFCssen diese Funktionen heutzutage wegen der
    Spammer bei \xF6ffentlichen Mail-Servern abgeschaltet sein. Im eigenen Netz (Hinter
    dem Firewall!) ist das nat\xFCrlich kein Problem.</p><p>Gru\xDF,<br>  D.</p>"
  date: '2009-03-01T17:29:46'
- author: lego xbox 360
  content: <p>Took me time and energy to research every one of the comments, but I
    definitely loved the actual write-up. That proved to become quite good to me and
    I am guaranteed to every one of the commenters below! Its constantly good when
    you are able not only be informed, but additionally amused! I'm good you had pleasure
    scripting this write-up.</p>
  date: '2011-09-18T13:02:16'
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
