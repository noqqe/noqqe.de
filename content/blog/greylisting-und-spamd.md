---
categories:
- administration
- bsd
- mail
- openbsd
- osbn
- planetenblogger
- ubuntuusers
date: 2016-10-23T16:34:04+02:00
tags:
- mail
- smtpd
- opensmtp
- openbsd
- spamd
- greylisting
title: Greylisting und spamd
---

Nachdem ich nun als Mailhoster schon 1&1, Uberspace und neomailbox.net
durch habe und sich mein Stack immer und immer wieder verändert kam es nun
doch mal dazu, dass ich meinen eigenen Mailserver betreibe.

Der Stack (ich sage so oft Stack, weil das so hip ist) besteht aus Dovecot
und OpenSMTPD. Daher sah ich mir auch gleich mal die Lösung aus `base` in
OpenBSD an.

{{< figure src="/uploads/2016/10/spam.png" >}}

`spamd` ist eine wunderbare kleine Lösung zum Blocken von IPs die sehr eng
mit `pf` zusammenarbeitet. Im Detail wird die Software sehr anschaulich
auch
[hier](http://protoc.org/blog/2014/12/22/trapping-spammers-with-the-openbsd-spam-deferral-daemon/)
erklärt. Eigentlich reicht aber auch die Manpage locker aus. Deswegen wird
das hier auch kein Tutorial.

### Greylisting

Um zu starten braucht man eigentlich nur 4 Dinge.

* Ein `pf` Ruleset aus der Manpage,
* einen Cronjob der `/usr/libexec/spamd-setup` aufruft
* den laufenden Daemon `spamd`
* und `spamlogd`

Man muss eigentlich wirklich nichts machen. Ein bisschen Hirn einschalten
vielleicht. Ich hatte das auch ein paar Tage so laufen. Immer mal wieder in
`spamdb` geschaut. Eigentlich ganz nett.

```
spamd[83667]: 190.158.6.160: To: wa1@noqqe.de
spamd[83667]: 190.158.6.160: Subject: contact information
spamd[83667]: 190.158.6.160: From: "Lillie Shields" <Shields.5201@cable.net.co>
spamd[83667]: 190.158.6.160: disconnected after 382 seconds. lists: nixspam
```

Es hat mir sehr viel Spaß gemacht zu sehen wie man Spammern das nimmt, was
wohl am meisten schmerzt. Zeit. Aber eben auch sowas:

```
[...]
GREY|173.0.84.225|xxx|service@paypal.com|1419811111|1419814111
GREY|173.0.84.226|xxx|service@paypal.com|1419834111|1419834111
GREY|173.0.84.228|xxx|service@paypal.com|1419811111|1419814111
```

Klar, dass man als riesen Firma mehrere MX'e hat. Das diese auch eine
verteile Mailqueue haben wundert mich auch nicht. Für Greylisting ist das
halt ein riesen Problem, da es nur beim Auftauchen gleichen IP auf Durchzug
schaltet. So kann es passieren das ich meine Mails nach 2 Wochen oder halt
einfach garnicht bekomme.

Was ist also die Lösung? Es soll Menschen geben die tatsächlich anfangen
Whitelists für solche Firmen zu pflegen.
Aus [Binärgewitter](https://binaergewitter.de) erfuhr ich dass Microsoft es mit
seiner Mailing Infrastruktur genauso hält. MX IPs sammeln? Hallo? Ich wollte kein neues
Hobby, ich wollte nur meinen eigenen Mailserver.

Und überhaupt. Ich bin zwar kein Typ der sich Push Notifications seiner
Mikrowelle schicken lässt, aber gibt es was Schlimmeres als auf Emails zu
warten die man braucht? Bei Registrierungen und Sonstigem. Einfach
furchtbar. Nein das wollte ich nicht.

### Blacklisting

Zum Glück gehts auch ohne. `spamd` lässt sich im `blacklist` Mode
konfigurieren. Anstatt einfach jede `SMTP` Connection zu `spamd`
weiterzuleiten, legt man einen `pf` Table an und leitet somit nur IPs aus
der Blacklist an den stotternden Daemon.

```
table <spamd> persist
pass in on $extif inet proto tcp from <spamd> to any port smtp divert-to
127.0.0.1 port spamd
```

Sowohl `spamd` also auch `spamd-setup` müssen dafür mit `-b` gestartet
werden. Die Config dafür:

```
all:\
  :nixspam:

nixspam:\
  :black:\
  :msg="Your address %A is in the nixspam list\n\
  See http://www.heise.de/ix/nixspam/dnsbl_en/ for details":\
  :method=http:\
  :file=www.openbsd.org/spamd/nixspam.gz
```

So werde ich meinen Spam los ohne Greylisting zu brauchen. Vielleicht
schaue ich mir irgendwann noch die BGP Variante an.
[bgp-spamd.net](http://bgp-spamd.net).
Vorerst reicht mir die Spamliste von Heise aber.

