---
aliases:
- /blog/2010/04/13/regexp-domainendungen
- /archives/983
date: '2010-04-13T18:07:32'
tags:
- development
- endungen
- shell
- grep
- regular expressions
- tld
- cat
- sort
- regular
- linux
- domains
- regexp
- sed
- uniq
- debian
- ubuntu
title: RegExp | Domainendungen
---

Heute mal etwas weniger spektakulär. Das ist lustig, wenn ich sowas
schreibe. Denn es impliziert, dass es hier schon mal etwas spektakuläres
gab. Reingefallen.

Heute wollte ein Kunde mit etwas mehr Domains wissen, welche verschiedenen
TLDs genau dabei sind. Hatte auch kostentechnische Gründe. Jedenfalls
wollte ich nach dem exportieren nicht wirklich die Liste durchgehen und
rausschreiben.

Wie beschreibe ich also Domains in Regular Expressions?

```
^.*.(.*)
```

Würde das Format beschreiben. Im Grunde alles was nach dem ersten Punkt
einer Zeile kommt in $1 ablegen (durch () markiert). Kämen keinen doppelten
Domains wie .co.uk in die Quere und Subdomains kommen nicht aus der
Domainliste. Ziemlich low-level-regexp.

```
cat KundenDomainliste.txt | sed -e 's/^.*.(.*)/1/'
```

Gibt die komplette Domainliste (nur mit TLDs) aus. Sonderzeichen wie ( )
müssen für die Bash noch escaped werden. Desweiteren noch den Ausdruck auf
den die Beschreibung zutrifft durch 1 ersetzen. Wurden aber nicht weniger
Domains.

```
cat KundenDomainliste.txt | sed -e 's/^.*.(.*)/1/' | uniq | sort
```

Die Ausgabe von sed an uniq zu übergeben, behebt diesen Umstand. Um noch
alphabetisch zu sortieren diese Ausgabe wiederrum an sort übergeben. Nett.
Raus kommt eine Liste von Domainendungen:

```
.au
.co.uk
.com
.de
.hu
.it
```

usw...  Übrigens bin ich mir der Ironie bewusst, cat zu benutzen. [Useless
use of cat](http://sial.org/howto/shell/useless-cat/)