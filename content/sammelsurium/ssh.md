---
title: ssh
date: 2020-03-09T13:42:07
tags:
- Software
- ssh
---

## Client Options

Passwort Auth erzwingen

```
ssh host -o PasswordAuthentication=yes -o PubkeyAuthentication=no
```

<!--more-->

Speziellen Port angeben

```
ssh host -p22
```

Connection Timeout spezifieren

```
ssh -o ConnectTimeout=10 <host>
```

Host Key verwerfen für automatisierung oder ähnliches

```
ssh -o StrictHostKeyChecking=no <host>
```

Speziellen Key angeben. Nicht das .pub!

```
ssh host -l noqqe -i ~/.ssh/private
```

## Public Key aus Private Key generieren

Zuweilen kommt es vor, dass man Public Keys nicht mehr findet und nur noch
die Private Keys hat. Zum Glück lassen sich diese 1:1 wieder generieren

```
ssh-keygen -y -f ~/.ssh/id_rsa > ~/.ssh/id_rsa.pub
```

## Secure ciphers

Vorsicht. Ändert sich auch immer mal wieder.

Im Zweifel mal hier nachschauen. [Crypto Hardening](https://bettercrypto.org/static/applied-crypto-hardening.pdf)

```
Ciphers aes256-gcm@openssh.com,aes128-gcm@openssh.com,aes256-ctr,aes128-ctr
MACs    umac-128-etm@openssh.com,hmac-sha2-512,hmac-sha2-256,hmac-ripemd160
KexAlgorithms diffie-hellman-group-exchange-sha256,diffie-hellman-group14-sha1,diffie-hellman-group-exchange-sha1
```

## Escape Characters

* `~.` Disconnect.
* `~^Z` Background ssh.
* `~##` List forwarded connections.

```
$ ~#
The following connections are open:
  #0 client-session (t4 r0 i0/0 o0/0 e[write]/4 fd 7/8/9 sock -1 cc -1)
```

* `~&` Background ssh at logout when waiting for forwarded connection / X11 sessions to terminate.
* `~?` Display a list of escape characters.
* `~B` Send a BREAK to the remote system (only useful for SSH protocol version 2 and if the peer supports it).
* `~C` Open command line. Currently this allows the addition of port forwardings using the -L and -R options (see above).
* `~R` Request rekeying of the connection (only useful for SSH protocol version 2 and if the peer supports it).
```

## Terminal Description

Um den sauberen Hostname im Tab zu haben.

```
Host *
PermitLocalCommand yes
LocalCommand if [[ $TERM == screen* ]]; then printf "\033k%h\033\\"; fi
```
