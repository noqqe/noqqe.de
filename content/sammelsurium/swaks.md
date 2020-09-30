---
title: "Swaks"
date: 2020-09-24T16:20:38+02:00
tags:
- Software
---

Swaks is einfach das beste Tool um irgendwas mit Email zu machen.
Das dieser Sammelsuriumseintrag noch fehlt ist eigentlich eine Schande.

<!--more-->

## Mailserver debuggen

Um einen Mailserver zu debuggen der keine validen DNS Records besitzt kann
man

```
swaks --from from@example.net --to to@example.net --server 127.0.0.1:12002
```

## Empfänger Email verifizieren

Verifizieren ob eine Bestimmte Emailadresse existiert

```
swaks --from test@example.net --to <empfaenger> -q RCPT
```

`-q` ist hierbei ein `--quit-after`. Sobald also `RCPT` gesendet wird,
beendet swaks die Verbindung.

