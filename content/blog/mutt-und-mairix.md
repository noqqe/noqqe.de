---
comments:
- author: nuit
  content: <p>notmuch und emacs  als reader :D</p>
  date: '2013-03-19T08:09:31'
- author: Anaximander
  content: <p>Auch wenn es etwas offtopic ist, was hat aus deiner Sicht gegen Claws
    Mail gesprochen?</p>
  date: '2013-03-19T08:15:53'
- author: Martin
  content: "<p>mutt (mit wie vielen externen Patches inzwischen? 10?) + offlineimap
    (+ tool f\xFCr das Auslesen der Keychain) + mairix ist jetzt ein weniger retardiertes
    Setup als out-of-the-box Tools wie Thunderbird oder ClawsMail?</p><p>Da kann ich
    ja gleich sowas wie <a href=\"http://www.ulm.ccc.de/ChaosSeminar/2012/04_mmh\"
    rel=\"nofollow\"><code>mmh</code></a> aufsetzen.</p>"
  date: '2013-03-19T08:30:20'
- author: noqqe
  content: <p>Kranke Sau! :D</p>
  date: '2013-03-19T08:40:51'
- author: noqqe
  content: "<p>Ich kann das garnicht so genau auf einen Punkt bringen. Evtl. h\xE4tte
    ich nur noch mit Clawsmail warm werden m\xFCssen. Bzw. mich an ihn gew\xF6hnen
    m\xFCssen - bei Thunderbird ging das ja damals auch.  Ganz aufgegeben hab ich
    Clawsmail bisher noch nicht.</p>"
  date: '2013-03-19T08:42:43'
- author: noqqe
  content: "<p>Naja, wollte ich sowas wie Thunderbird br\xE4uchte ich ja: offlineimap,
    mutt-superbonus-patches, abook, mutt-print, pdf2text, mairix, w3m mailcap, usw...
    </p><p>Was ich sagen will: Ich finds nicht retardiert, weil ich selektiv Komponenten
    zu mutt hinzufuegen kann. </p><p>Was gibts denn noch f\xFCr externe Patches f\xFCr
    mutt, au\xDFer sidebar?</p>"
  date: '2013-03-19T08:46:05'
- author: Martin
  content: "<p>Mindestens 25, die zwar die meisten Distributionen einpflegen, aber
    nicht Upstream gehen. Dazu gibt es halt noch den sidebar, trash, pgp-verbose-mime
    und ignore-thread Patch.</p><p>Bei den &gt;100k Mails ist die Performance beim
    Navigieren in Mutt fl\xFCssig? Mir kam das so vor, als ob das alles single-threaded
    ist.</p>"
  date: '2013-03-19T08:51:18'
- author: noqqe
  content: "<p>Ah, jetzt weiss ich was du meinst. Ja benutze auch das mutt-patched
    \ Paket aus Debian. Wusste aber nicht, dass da mehr drinsteckt als der Sidebar
    Patch. </p><p>Muss ich mir mal durchlesen :) </p><p>Die 100k Mails sind gr\xF6\xDFtenteils
    in meinen Archives Ordnern. Hebe in Inbox nur das aktuelle Jahr auf und auch da
    sortiere ich Teilweise schon Themenorientiert in den Archives 2013 Ordner weg.</p>"
  date: '2013-03-19T08:56:03'
- author: Anaximander
  content: "<p>Erinnert mich an meinen Umstieg von Kmail zu Claws Mail. Da war ich
    anfangs auch nicht zufrieden, ohne wirklich sagen k\xF6nnen wieso. Inzwischen
    kann ich mir nicht mehr vorstellen Kmail zu nutzen.</p>"
  date: '2013-03-19T13:12:24'
- author: dakira
  content: "<p>Hab ich was verpasst? Thunderbird wird doch weiterhin von der Community
    weitergepflegt. Von Mozilla kam auch schon bei den letzten \"offiziellen\" releases
    nichts substanzielles mehr. Die meisten \xC4nderungen kamen von Extern. Die haben
    es doch nur offiziell gemacht, dass andere sich besser um Thunderbird k\xFCmmern.
    Und diese \"anderen\" haben auch direkt gesagt, dass sie weiter in Thunderbird
    investieren.<br>Edit: <a href=\"http://www.soeren-hentzschel.at/mozilla/thunderbird/2012/07/23/die-zukunft-von-thunderbird-ausblick-auf-kommende-features/\"
    rel=\"nofollow\">Hier</a> mehr Informationen dazu und <a href=\"http://www.soeren-hentzschel.at/mozilla/thunderbird/2013/03/08/mozilla-veroffentlicht-thunderbird-20-0-beta/\"
    rel=\"nofollow\">Thunderbird 20</a> ist auch gerade raus.</p>"
  date: '2013-03-19T16:33:06'
- author: Jochen
  content: "<p>Warum syncst du nicht dein IMAP dauerhaft in das Maildir und l\xE4sst
    den Mutt auf das locale Maildir los? Was bringt es dir zwei imap-Clients und verschiedene
    Datenbasen zu haben?</p>"
  date: '2013-03-22T19:00:09'
- author: noqqe
  content: "<p>naja, Mails am Handy und an anderen Rechnern. Das ist dann nicht mehr
    so wirklich Konsistent. So hab ich das Maildir nur f\xFCrs Suchen, Rest benutzt
    weiterhin IMAP.</p>"
  date: '2013-03-24T12:45:42'
date: '2013-03-18T20:00:00'
tags:
- index
- stats
- mutt
- suche
- mail
- mairix
title: mutt und mairix
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

Das dürfte eigentlich ja nicht so schwer sein. `mutt` bietet mir zwei
dieser 3 Punkte. Die Suche funktioniert zwar über `l` auch ganz gut, aber
nicht wirklich schnell und auch nicht in meinem ganzen IMAP Account.

Was her musste war ein Mailindexer. Textbasiert gibt es da verschiedene
Lösungen. Unter anderem [Sup](http://sup.rubyforge.org/),
[notmuch](http://notmuchmail.org/) und
[mairix](http://www.rpcurnow.force9.co.uk/mairix/).  Habe alle Tools
angetestet und mich für `mairix` entschieden. Eingängige Syntax, leicht
verständlich.

Einziges Problem: IMAP sprechen ist nicht drin. Local maildir'n'stuff. Deshalb
habe ich mir dann die folgende Lösung überlegt.

{{< figure src="/uploads/2013/03/mairix-mutt.png" >}}

`mairix` durchsucht also ausschliesslich das lokale Maildir, welches ich mit
offlineimap täglich (reicht mir) synce, und gibt dir Ergebnisse in eine
separate mbox `Search`.

``` bash
sudo aptitude install mairix offlineimap
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

Noch 2 Cronjobs für regelmäßiges mairix und offlineimap sync und alles ist
entspannt.  Für Verbesserungsvorschläge und Inspirationen aus euren
Mail-Setups bin ich wie immer gerne zu haben! :)
