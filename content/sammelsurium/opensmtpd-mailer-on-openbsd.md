---
title: OpenSMTPD Mailer on OpenBSD
date: 2016-09-16T12:30:49
tags:
- OpenBSD
- Mail
- Opensmtpd
---

#### Local Forwards...

Hab mehrere Mailadressen

mail@devnull-as-a-service.com
mail@organic-entropy.org

die ich einfach auf meine reguläre Adresse weiterleiten will.
Dazu in aliases

    mail: flo@example.com

eintragen und die folgenden Regeln setzen

    table rdrdomains { "example.com", "devnull-as-a-service.com", "organic-entropy.org", "entbehrlich.es" }
    accept from any for domain <rdrdomains> alias <aliases> deliver to mbox

#### SSL enable

SSL Cert generieren

    openssl genrsa 4096 > /etc/ssl/private/mail.example.com.key
    openssl req -new -sha256 -key /etc/ssl/private/mail.example.com.key -subj
    "/CN=mail.example.com" > /etc/ssl/csrs/mail.example.com.csr

Einbinden

    pki mail.example.com certificate "/etc/ssl/certs/mail.example.com.crt"
    pki mail.example.com key "/etc/ssl/private/mail.example.com.key"

#### Encrypted Queues

    queue compression

#### Spamd

Greylisting ist fürn Arsch. Ich hab keinen Bock 10.000 Jahre auf ne Mail zu warten.
Blacklist only mode.

#### Mailregeln

Scheinbar muss ich von maildrop weg und zu sieve gehen. Gut das ist jetzt
nicht sonderlich schlimm. Aber trotzdem.

Um Sieve zu aktivieren

    plugin {
      sieve = /home/%Ln/mail.sieve
    }

    protocol lmtp {
      ## Space separated list of plugins to load (default is global mail_plugins).
      mail_plugins = $mail_plugins sieve
    }


#### Auth Bruteforce


    ## Deny and track SMTP Relay bruteforce
    table <bruteforce-relay> persist
    pass in on $extif proto tcp from any to any port submission \
      flags S/SA keep state \
      (max-src-conn 30, max-src-conn-rate 30/60, \
      overload <bruteforce-relay> flush global)

#### DKIM

#### SMTP Relay

Wenn kein Auth table gegeben, dann wird system auth genommen.

    listen on xnf0 port submission tls pki mail.example.com auth

#### Dovecot Users

Delivery durch einfach nur local auth

    mail_location = maildir:~/mail
    auth_username_format = %Ln

    passdb {
        driver = bsdauth
    }

    userdb {
        driver = passwd
    }

und im smtpd

    table mailboxdomains { "example.com" }
    accept from any for domain <mailboxdomains> deliver to lmtp "/var/dovecot/lmtp" rcpt-to

Hurra!

#### Reverse Lookup

mail.example.com -> IP
IP -> mail.example.com

und unbedingt an IPv6 denken!

#### Dovecot Security

IMAPS only.

    ## TLS
    ssl = yes
    ssl_cert = </etc/ssl/mail.example.com.crt
    ssl_key = </etc/ssl/private/mail.example.com.key

    ## Disable unsecure IMAP
    service imap-login {
      inet_listener imap {
        port = 0
      }
      inet_listener imaps {
        port = 993
      }
    }

#### Monitoring

SMTP Incoming - Port 25
SMTP Relay - Port 587
IMAPS - Port 993

#### Tests

Forwarding works.

    swaks --to flo@example.com --server mail.example.com --timeout 1800
    swaks --to noqqe@example.com --server mail.example.com

Mailbox delivery Works

    swaks --to flo@example.com --server mail.example.com

Relaying  on port 25 does not work

    swaks --to tax@z.2om.us --server mail.example.com

Forwarding with TLS

    swaks --to noqqe@example.com --server mail.example.com --tls

Delivery to mbox with tls

    swaks --to flo@example.com --server mail.example.com --tls

Authenticating works

    swaks --to noqqe@example.com --server mail.example.com --tls --port 587
    --auth --auth-user noqqe

Unauthenticated relay does not work

    swaks --to noqqe@example.com --server mail.example.com --tls --port 587
    --auth --auth-user noqqe

Authenticated relaying works

    swaks --to tax@z.2om.us --from flo@example.com --server mail.example.com
    --tls --port 587 --auth --auth-user flo --auth-pass FOOLOL
