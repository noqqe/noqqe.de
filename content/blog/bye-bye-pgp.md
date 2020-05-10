---
title: "Bye Bye PGP"
date: 2020-05-10T12:09:43+02:00
tags:
- PGP
- GPG
---

Dieser Blogpost könnte mit einer Fülle an Links über Probleme von PGP
beginnen. Tut er aber nicht. Es wäre weder komplett, noch wirklich wichtig.
Probleme existieren auf diversen unterschiedlichen Ebenen wie: Usability,
Tooling, Integration, und Crypto Algorithmen und letzten Endes sogar
Keyserver Security.

<!--more-->

{{< figure src="/uploads/2020/05/patrick.png" >}}

[2014 saß ich am Debian Booth](/blog/2014/03/16/enlarge-your-p-dot-gp-key/)
auf den Chemnitzer Linuxtagen habe meinen Key generiert.
Damals hatte sich das irgendwie richtig angefühlt. Seitdem habe ich genau
**73** Mails mit GPG verschlüsselte Mails bekommen. 12 davon Tests.

Ich glaube nicht an die Zukunft von PGP und nutzen tue ich es auch nicht.
Deshalb habe ich meinen Key jetzt revoked und auf die Keyserver gepushed.

```
> gpg --pinentry-mode loopback --gen-revoke CDA4B775

sec  rsa4096/BFE78B80CDA4B775 2014-03-15 Florian Baumann (noqqe)
[...]
Reason for revocation: Key is no longer used
PGP is dead
Is this okay? (y/N) y

> gpg --list-secret-keys
--------------------------------
sec   rsa2048 2010-06-25 [SC] [revoked: 2014-03-15]
sec   rsa4096 2014-03-15 [SC] [revoked: 2020-05-10]
```

Bye bye.
