---
title: SSL/TLS Cheatsheet
date: 2011-11-15T14:46:15
tags: 
- SSL
- TLS
- Software
---

#### HTTP vs HTTPS

* Brauch ich ja keinem erklären.
* Sicher & Toll.
* SSLv2 und SSLv3 ist tot. Deshalb macht es eigentlich keinen
  Sinn mehr das ganze noch SSL zu nennen, da es kein funktionierendes SSL
  Protokoll mehr gibt. Gängige Projekte gehen dazu über
  alles einfach TLS zu nennen.

#### Verbindungsaufbau

* TCP Verbindung, SSL Handshake, HTTP über Secure Layer.

https://www.identrustssl.com/images/learn_ssl_diagram.gif

#### Protokolle

* SSLv1 .. bitte nicht.
* SSLv2 .. kaputt gegangen bei DROWN-Lücke
* SSLv3 .. kaputt gegangen bei Poodle-Lücke
* TLS1.0 .. geht schon
* TLS1.1 .. geht
* TLS1.2 .. total super
* TLS1.3 .. kam letztens(tm) raus.

#### Algorithmen

Innerhalb der Protokolle werden Ciphers oder Algorithmen gesprochen.

Prüfsummendinge

> CBC, ECB, MD5

Symmetrische Verfahren

> AES, Twofish, Blowfish, IDEA
> Camellia, Salsa20, Chacha20, Cast

Asymetrische Verfahren

> RSA, DSA, Diffie Hellman, XTR

Elyptische Kurven

> Curve1174, Curve25519, Curve383187, Curve41417, Ed448-Goldilocks,
> NISTP-384, NIST P-256, ECDSA

#### Ciphersuites

* Wie gesagt, SSL/TLS ist hybrid.
* Für verschiedene Sachen braucht man verschiedene Algorithmen
* Zussammenstellung

> Schlüsselaustausch: RSA, DH (auch ADH, ECDH), PSK, SRP
> Authentifizierung: RSA, DSA (auch ECDSA), PSK
> Hashfunktion: (MD5, SHA)
> Verschlüsselung: (keine, RC4, DES, 3DES, IDEA, AES)

ABER: Zum Glück muss man das nicht alles selber machen.
Es gibt fertige und ständig aktualisierte Cyphersuites die man fuer
Webserver, SSH und sonstigen Käse verwenden kann: https://bettercrypto.org

#### SSL/TLS Software

* openssl
* gnutls (GNU Projekt)
* libressl (OpenBSD)
* BoringSSL (Google)
* MatrixSSL (unwichtig...)
* WolfSSL (unwichtig...)

und noch viel Kram für verschiedene Anwendungszwecke
wie die Java Library oder Libs für irgendwelche Embedded Devices.

#### Erstellen

CSR erstellen

    openssl genrsa -out host.key 4096 ## ohne pw
    openssl genrsa -des3 -out 185.34.0.152:443.key 4096 ## mit pw
    openssl req -new -nodes -key host.key -out host.csr

CRT Selfsigned ausstellen (1095 == 3 Jahre)

    openssl x509 -req -days 365 -in host.csr -signkey host.key -out host.crt

Entfernen der PEM pass phrase eines zuvor mit PEM pass phrase generierten Schlüssels

    openssl rsa -in host.key -out host_np.key

#### Auslesen

CRT auslesen:

    openssl x509 -in host.crt -text
    openssl x509 -noout -text -in host.crt

RootCA auslesen

    openssl x509 -in host.pem.cer -text

CSR ansehen

    openssl req -noout -text -in host.csr

#### Verifizierung

Verifizieren ob mit RootCA zu Certificate passt:

    openssl verify -CAfile provider.crt -verbose domain.crt
    domain.crt: OK

Verifizieren ob Cert zu Key passt (wenn gleicher Key dann ok):

    (openssl x509 -noout -modulus -in host.pem | openssl md5 ; openssl rsa -noout -modulus -in host.key | openssl md5)

Client SSL auf URL prüfen:

    openssl s_client -connect www.domain.com:443

Review der Certificate Chain

    openssl s_client -connect myweb.com:443 -showcerts -CApath /etc/ssl/certs

Wenn man wissen will ob intermediate Zertifikat zum eigentlichen Zertifikat passt müssen diese beiden nummern gleich sein:

    $ openssl x509 -in domain.crt -issuer_hash -noout
    a302054c

    $ openssl x509 -in intermediate.pem -subject_hash -noout
    b204d74a


#### Selbsterstellte Zertifikate

* Generating key
* Generating certificate signing request
* Generating self signed certificate
* Sign signing request to get signed certificate
* Use Signed Certificate + Private Key
* Configure web server for those.

## PEM Files und die Reihenfolge

Do not use openssl verify for .pem files!

    $ openssl verify all.pem
    all.pem: OU = Domain Control Validated, OU = PositiveSSL Wildcard, CN = *.example.com
    error 20 at 0 depth lookup:unable to get local issuer certificate

After reading some ssl3 documentation it seems pem files have to be in
a very specific order (see manpage SSL_CTX_use_certificate_chain_file). The
defined order is

* private key
* certificate
* 1 chain cert (e.g. Intermediate)
* 2 chain cert (rootca)

#### SNI (Server Name Indication)

* Probleme früher hart wenn mehrere Hosts (vHosts) unter einer IP laufen.
* Da TCP -> SSL -> HTTP gesprochen wird.
* Hostname wird erst in HTTP uebermittelt
* Problem lief dann wie folgt:
  Server macht tcp und ssl
  Server weiss nicht fuer welchen VHOST
  Server liefert _irgendein_ Zertifikat aus
  Server hoert dann welcher HTTP Host angefragt wird
  Client faellt auf die Fresse, weil Zertifikat nicht zum Hostname passt
* SNI ist hierfuer die Loesung.

#### Lets Encrypt & ACME

* Neues Verfahren
* API gesteuert zum Abholen des CERTS

Ablauf:

* Vorbereiten des Private Key + CSR
* Generieren einer Challenge und ablegen auf
  $domain/.well-known/challenge.txt
* Query gegen den Austeller (Server) mit uebermitteltem CSR
* Server besucht und verifiziert die challenge.txt
* Server liefert unterschriebenes Certificate zurück

So zum Beispiel

    H=p.example.com
    openssl genrsa 4096 > /etc/ssl/private/$H.key
    openssl req -new -sha256 -key /etc/ssl/private/$H.key -subj "/CN=$H" > /etc/ssl/csrs/$H.csr

    TACME=/home/noqqe/Code/acme-tiny/acme_tiny.py
    python $TACME --account-key /etc/ssl/private/letsencrypt.key --csr \
      /etc/ssl/csrs/${x}.csr --acme-dir /var/www/htdocs/challenges/ \
      > /tmp/signed.crt

#### Speedtest

Ausführen

    $ openssl speed

Speedcheck auswerten

    $ cat speedcheck \
    | sed -e 's#Doing\s\(.*\)\s\(.*\)\sfor#Doing \1\-\2 for#' -e 's#10s:\s\(.*\)\s.*bit#10s: \1 \1 \1 bit#g' \
    | awk '{ print $9":"$2 }' \
    | sort -rn \
    | awk -F: '{print $2":"$1}' \
    | statistical \
    | head -50
