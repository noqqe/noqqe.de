---
title: "Enlarge your P..GP Key"
date: 2014-03-16T17:20:00+02:00
comments: true
tags:
- Crypto
- pgp
- gpg
- key
- revoke
- certificate
- debian
- aptitude
- signing-party
- ps
- pdf
- pubkey
- publickey
- replace
- migrate
aliases:
- /blog/2014/03/16/enlarge-your-p-dot-gp-key
---

Mittlerweile ist mein 2048bit PGP Key mit ein paar Signaturen
ca. 4 Jahre alt und auch etwas kurz. State of the Art ist wohl 4096bit.
Der Migrationspfad auf einen neuen Key.

### Sicherung

Sichern des alten Key Environments

    $ cp -a ~/.gnupg ~/.gnupg-old

### Neuen Key generieren

Wesentlich für einen tollen neuen 4096bit Key ist ausreichend Entropy.
Für aus der CPU generierte Entropie, kann
[havege](http://www.irisa.fr/caps/projects/hipsor/) installiert werden.

    $ cat /proc/sys/kernel/random/entropy_avail
    > 153
    $ aptitude install havege
    $ /etc/init.d/havege start
    $ cat /proc/sys/kernel/random/entropy_avail
    > 1649

und anschliessend den Key erstellen

    $ gpg --gen-key

Der interaktive Dialog sollte einen eigentlich halbwegs verständlich durch die
Generierung führen.

### Alten Key revoken

Beim Revoken des alten Keys, ist es sinnvoll die korrekte Begründung
auszuwählen. In meinem Fall Ersetzung durch einen neuen Key, also superseded.

```
$ gpg --gen-revoke oldkeyid

Create a revocation certificate for this key? (y/N) y
Please select the reason for the revocation:
  0 = No reason specified
  1 = Key has been compromised
  2 = Key is superseded
  3 = Key is no longer used
  Q = Cancel
(Probably you want to select 1 here)
Your decision? 2
Enter an optional description; end it with an empty line:
> replaced by newkeyid
>
Reason for revocation: Key is superseded
replaced by newkeyid
```

Heraus fällt ein Public-Key Block, der in einer Datei importiert werden kann.

    $ gpg --import revoke.txt

### Migration

Auf meinem alten Key befinden sich wie gesagt ein paar Unterschriften. Auch
deswegen macht es Sinn den neuen Key mit dem Alten zu unterschreiben.

    $ gpg --default-key oldkeyid --sign-key newkeyid

Dadurch wird nicht nur die Referenzkette gebildet, sondern auch gleich
nachvollziehbar, dass der neue Key wirklich der Nachfolger meiner "originalen"
KeyID ist.

### Keyserver aktualisieren

Wenn der alte Publickey nun revoked ist, kann es auf den Keyserver geladen
werden.

    $ gpg --list-key oldkeyid
    $ gpg --send-keys oldkeyid

Genauso der neue Publickey.

    $ gpg --list-key newkeyid
    $ gpg --send-keys newkeyid

### Nacharbeiten

* Default Key in der gpg.conf ändern `default-key newkeyid`
* PGP Plugin in mutt (oder Mailclient) updaten
* Zum Austausch des Fingerprints fürs einfaches Signing, gibt es im Debian Paket
  signing-party die entsprechende key2ps Tools. Diese kann dann zum PDF
  umgewandelt und ausgedruckt werden.

```
$ aptitude install signing-party ghostscript
$ gpg-key2ps -p a4 CDA4B775 > key.ps
$ ps2pdf key.ps
```

* Signatur im Mailclient updaten
* Informationen auf der Website updaten
* [caff](https://wiki.debian.org/caff) Konfiguration anpassen

{{< figure src="/uploads/2014/03/w0000t.gif" >}}
