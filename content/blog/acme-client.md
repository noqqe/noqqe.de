---
date: 2017-04-22T09:51:51+02:00
tags:
- letsencrypt
- acme
- libressl
title: acme-client(1)
---

Auf was ich mich am Meisten im neuen OpenBSD Release 6.1 gefreut hab ist
[acme-client(1)](http://man.openbsd.org/acme-client). Das Binary welches
direkt aus dem `base` System kommt ersetzt bei mir einen hässlichen Wust
aus Shell Skripten und anderen Schnipseln die ich mir irgendwo notiert
hatte.

``` bash
#!/usr/local/bin/bash
TACME=/home/noqqe/Code/acme-tiny/acme_tiny.py

wget --quiet -O - https://letsencrypt.org/certs/lets-encrypt-x3-cross-signed.pem > /tmp/intermediate.pem

D="organic-entropy.org entbehrlich.es devnull-as-a-service.com noqqe.de"
for x in $D ; do
  python $TACME --account-key /etc/ssl/private/letsencrypt.key --csr /etc/ssl/csrs/${x}.csr --acme-dir /var/www/htdocs/challenges/ > /tmp/signed.crt || exit
  cat /tmp/signed.crt /tmp/intermediate.pem > /etc/ssl/${x}.crt
  /etc/rc.d/nginx reload
  rm /tmp/signed.crt
done

rm /tmp/intermediate.pem
```

Das lief alle heilige Zeit als Cronjob. Mit `acme-client` kann ich in der
Config `/etc/acme-client.conf` definieren welche Domain mit welcher
Authority registriert werden soll und welche Pfade ich gerne hätte.

```
authority letsencrypt {
  agreement url "https://letsencrypt.org/documents/LE-SA-v1.1.1-August-1-2016.pdf"
  api url "https://acme-v01.api.letsencrypt.org/directory"
  account key "/etc/ssl/acme/private/letsencrypt.key"
}

domain organic-entropy.org {
  domain key "/etc/ssl/private/organic-entropy.org.key"
  domain certificate "/etc/ssl/organic-entropy.org.crt"
  domain full chain certificate "/etc/ssl/organic-entropy.org.pem"
  challengedir "/var/www/htdocs/challenges/"
  sign with letsencrypt
}

[...]
```

Was `acme-client` **nicht** versucht (Gott sei dank) ist deinen HTTP Server zu
konfigurieren. Weiss nicht wer auf die glorreiche Idee kam, aber es ist
definitiv eine gute Entscheidung das nicht zu tun. Abhilfe schafft hier die
Manpage, in der ein paar Beispiele zur Konfiguration von `httpd(8)` zu
finden sind, falls man wirklich nicht weiter kommt.

{{< figure src="/uploads/2017/04/ssl.jpg" >}}

Eigentlich ging ich anfangs davon aus, dass es noch eine Art Daemon gibt
der immer mal wieder kontrolliert ob das Cert erneuert werden muss. Die
Manpage zeigt aber, dass man am besten einfach einen Cronjob nehmen soll
pro Domain. Das gefällt mir ehrlich gesagt nicht so. Wenn man jetzt **richtig** viele Domains
hat brauche ich doch wieder irgendein Shell Script. Für jetzt reicht es mir
aber das zu machen wie beschrieben.

```
# Acme Client
0 5 * * * acme-client noqqe.de && rcctl reload nginx
```

Jeden Tag? Ja. Sollte das Cert noch zu lange gültig sein, wird es erst
garnicht erneuert.

``` bash
# acme-client -v noqqe.de
acme-client: /etc/ssl/noqqe.de.crt: certificate valid: 86 days left
```

Zertifikate die früher per Hand erstellen musste (siehe [TLS Cheatsheet](https://noqqe.de/sammelsurium/ssltls-cheatsheet/))
kann nun ebenfalls `acme-client` direkt.

``` bash
acme-client -dv newdomain.org
```

Ich werde das jetzt mal ein paar Wochen testen, bin aber zuversichtlich™.
