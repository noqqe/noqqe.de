---
type: post
title: "mutt und mairix"
date: 2013-03-18T22:00:00+02:00
comments: true
categories:
- Mail
- mutt
- mairix
- suche
- index
- ubuntuusers
- osbn
---

Ich war auf der Suche nach einem neuen Mailclient. Thunderbird
wird nur noch mit [Security Fixes versorgt](http://www.golem.de/news/e-mail-client-mozilla-will-thunderbird-nicht-mehr-weiterentwickeln-1207-93038.html).
Und außerdem hängt er mir zum Hals raus. Bloated Shitfuck.

Probierte viel aus. [KMail](http://userbase.kde.org/KMail), [Geary](http://yorba.org/geary/),
[ClawsMail](http://www.claws-mail.org/). Wobei letzteres mir noch am Besten
gefallen hat. Vollkommen überzeugt war ich aber nicht.

Ich musste mir erstmal klar werden was ich will. Nach etwas hin und her habe ich
mich auf folgende Punkte festgelegt:

* Minimalistischer Client
* GPG/PGP-Fähigkeit
* Schnelle und bedienbare Suche

Das dürfte eigentlich ja nicht so schwer sein. `mutt` bietet mir zwei dieser 3
Punkte. Die Suche funktioniert zwar über `l` auch ganz gut, aber nicht wirklich
schnell und auch nicht in meinem ganzen IMAP Account.

Was her musste war ein Mailindexer. Textbasiert gibt es da verschiedene Lösungen. Unter
anderem [Sup](http://sup.rubyforge.org/), [notmuch](http://notmuchmail.org/) und [mairix](http://www.rpcurnow.force9.co.uk/mairix/).
Habe alle Tools angetestet und mich für `mairix` entschieden. Eingängige Syntax,
leicht verständlich.

Einziges Problem: IMAP sprechen ist nicht drin. Local maildir'n'stuff. Deshalb
habe ich mir dann die folgende Lösung überlegt.

{{< figure src="/uploads/2013/03/mairix-mutt.png" >}}

`mairix` durchsucht also ausschliesslich das lokale Maildir, welches ich mit
offlineimap täglich (reicht mir) synce, und gibt dir Ergebnisse in eine
separate mbox `Search`.

```
$ sudo aptitude install mairix offlineimap
```

Meine mairix-Config:

``` bash
# local Maildir
base=/home/noqqe/Maildir/
database=/home/noqqe/Maildir/.mairixdb

# Set this to a list of maildir folders within 'base'.
maildir=INBOX
maildir=Archives*
[...]

# Drop search results here:
mfolder=Search
mformat=mbox
```

Die Konfiguration für `offlineimap` ist auch kein Kunststück, deshalb lass ich
das hier weg. Mailindex durch mairix auf das bestehende Maildir initial aufbauen.
Auch gleich ein schöner Überblick über das eigene Postfach.

```
$ mairix -v
Wrote 126633 messages
Wrote 0 mbox headers
Wrote 0 bytes of mbox message checksums
To: Wrote 8458 tokens
Cc: Wrote 4289 tokens
From: Wrote 18506 tokens
Subject: Wrote 36705 tokens
Body: Wrote 620414 tokens
Attachment Name: Wrote 3313 tokens
(Threading): Wrote 133934 tokens
```

Angenommen man möchte alle Mails von Amazon finden, die im Subject "Linux"
enthalten kann man

```
$ mairix f:Amazon s:Linux
Matched 3 messages
```

benutzen. Der springende Punkt ist allerdings: Wie kann ich das jetzt sinnvoll
in meinen Client integrieren? Mein Mutt verbindet sich mit einem IMAP Account
und die Suchergebnisse liegen irgendwo lokal auf der Platte.

Dafür habe ich einen Macro in meiner `.muttrc` definiert.

```
macro index,pager L "<change-folder-readonly>/home/noqqe/Maildir/Search<enter><shell-escape>mairix " "search via mairix"
unset wait_key # do not require additional enter for shell commands
```

Noch 2 Cronjobs für regelmäßiges mairix und offlineimap sync und alles ist entspannt.
Für Verbesserungsvorschläge und Inspirationen aus euren Mail-Setups bin ich wie immer gerne zu haben! :)

