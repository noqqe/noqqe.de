---
categories:
- openbsd
- osbn
- planetenblogger
- ubuntuusers
date: 2016-11-13T10:50:02+01:00
tags:
- smtp
- mail
title: Reverse DNS SMTP Gedönz
---

Seit 2-3 Monaten betreibe ich meinen Mailserver nun selbst. Dann wollte ich
eine mail zu web.de schicken und wurde abgewiesen.

```
550-Requested action not taken: mailbox unavailable
invalid DNS MX or A/AAAA resource record
```

Was denn nun? Egal. Ich habs natürlich falsch gemacht. Für die Domain
`noqqe.de` gabs den MX `mail.n0q.org` und der war ein CNAME auf
`aax.n0q.org`. Letzterer hat eine saubere Rückwärtsauflösung für v4 und v6.

{{< figure src="/uploads/2016/11/look.jpg" >}}

Nach etwas genauerem Recherchieren in der [K4CG](https://k4cg.org) zeigt
sich dann in [RFC2181](https://tools.ietf.org/html/rfc2181#section-10.3)
der treffenderweise "Clarifications to the DNS Specification" heisst:

> The domain name used as the value of a NS resource record, or part of the
> value of a MX resource record must not be an alias.

Und weiter bzgl. des Warum:

> This can cause extra queries, and extra network burden, on every query.
> It is trivial for the DNS administrator to avoid this by resolving the
> alias and placing the canonical name directly in the affected record just
> once when it is updated or installed.  In some particular hard cases the
> lack of the additional section address records in the results of a NS
> lookup can cause the request to fail.

Also ists dann **falsch**:

```
$ host -t mx noqqe.de
noqqe.de mail is handled by 10 mail.n0q.org.

$ host -t A mail.n0q.org
mail.n0q.org is an alias for aax.n0q.org.
aax.n0q.org has address 185.34.0.152
```

Man muss halt direkt den Hostname (in meinem Fall aax.n0q.org) hinterlegen. Schade, denn eigentlich
nehm ich immer den Weg über diese Indirektion.

Folgt man dem RFC klappts auch wiedder mit den Nachbarn.
