---
title: "age"
date: 2021-05-17T10:30:17+02:00
tags:
- Software
---

`age` ist ein Commandline-Tool zum verschlüsseln von Daten. Das neue PGP
kümmert sich intern um alles was Sicherheit angeht. Crypto Standards, Ciphers
und ähnliches ist nicht konfigurierbar und soll so simpel bleiben.

<!--more-->

## Passphrase

Encrypt

    age -e -a -p  merk.txt > merk.txt.age
    Enter passphrase (leave empty to autogenerate a secure one):
    Confirm passphrase:

und Decrypt.

     age -d merk.txt.age > merk.txt

Passphrase wird automatisch erkannt

## Keys und Recipients

`age` kann sowohl bestehende `ssh` keys oder eigenes mit `age-keygen`

### Age Keys

Key generieren

    > age-keygen -o age.txt
    # created: 2021-05-17T12:00:30+02:00
    # public key: age1hx2q3czpucxwtpld9lpt0848rarycr9rjdzp0sd6jjd2pwzun5eqdc3k5r
    AGE-SECRET-KEY-XXX

Encrypt

    age -r age1hx2q3czpucxwtpld9lpt0848rarycr9rjdzp0sd6jjd2pwzun5eqdc3k5r merk.txt > merk.txt.age

Decrypt

    age -d -i age.txt merk.txt.age


### SSH Keys

Kleines Beispiel mit eigenen SSH Schlüssel

Encrypten

    age -R .ssh/id_rsa.pub -o merk.txt.age merk.txt

Decrypt

    age -i .ssh/id_rsa -d merk.txt.age
