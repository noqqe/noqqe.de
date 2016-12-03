---
title: "Mehr BSD, Bye Uberspace"
date: 2014-06-19T17:20:00+02:00
comments: true
tags:
- Mail
- osbn
- BSD
- OpenBSD
- Linux
- Uberspace
- imapfilter
- mailbox
- bogofilter
- dns
- blog
- domain
- apache
- nginx
- osbn
---

Ab und an passiert es, dass Dinge die einfach funktionieren anfangen mich
zu langweilen.  So geschehen mit [Uberspace](http://uberspace.de). Das
[CentOS](http://centos.org), der [Apache2.2](http://httpd.apache.org/), die
[Sache mit den Domains](https://wiki.uberspace.de/philosophy:domains).

{{< figure src="/uploads/2014/06/dagobert.jpg" >}}

Im Grunde habe ich die 3 Dienste, die ich bei Uberspace nutze (Webserver für Blog,
Mail, Domains) nun zu anderen Dienstleistern migriert.

### Blog Webserver

Der erste "Dienstleister" bin ich gewissermaßen selbst. Die 1HE Dell
Maschine die ich seit einiger Zeit bei meinem alten Arbeitgeber mit
[OpenBSD](http://openbsd.org) betreibe, hostet mit
[nginx](http://nginx.org) den Blog. Selbige Kiste delivered auch
devnull-as-a-service.com und coffeestat.org. Warum also eine extra Lösung.

### Mails

Bei Uberspace [hatte ich](https://noqqe.de/blog/2013/10/26/spammer-vs-statistik-mit-bogofilter/)
die Mglichkeit die Mails direkt von Qmail über
[Maildrop](http://www.courier-mta.org/maildrop/) umzusortieren und dort
meine Tools zu platzieren.

Mein neuer Mailprovider [neomailbox.net](https://neomailbox.net) hat seinen
Firmensitz auf den Seychellen, Server in der Schweiz und setzt ebenfalls
auf OpenBSD als OS.  Nicht weil ich ernsthaft glaube, das ich das bräuchte,
sondern weils mir gefällt.

{{< figure src="/uploads/2014/06/underwood.jpg" >}}

Mit dem Wechsel ändert sich aber die Architektur meines Setups. Statt
direkt am IMAP Server arbeiten zu können bin ich nun nur noch Konsument des
IMAP Services.  Blöd für maildrop-Regeln und Spamfiltering mit
[bogofilter](http://bogofilter.sourceforge.net/).

Gerade beim Abschied von Bogofilter tat ich mich schwer, weils so gut und
einfach ist. Aber dann entdeckte ich
[imapfilter](https://github.com/lefcha/imapfilter).

imapfilter ist primär dazu gedacht Mails zwischen einem oder mehrere IMAP
Accounts zu verschieben.  Meine maildrop Regeln sind sehr leicht umgesetzt:

``` lua
neomailbox = IMAP {
    server = 'neomailbox.net',
    username = 'user',
    password = 'pw',
    ssl = 'ssl3',
}

-- get set of mails from the last day
res = neomailbox.INBOX:is_newer(1)

-- move techmails to Tech/
tech = res:contain_to('tech@openbsd.org') +
       res:contain_cc('tech@openbsd.org') +
       res:contain_from('Cron Daemon <root@') +
       res:contain_from('Charlie Root <root@o0.n0q.org>') +
       res:contain_from('root@z0idberg.') +
       res:contain_from('noc.n0q.org')
tech:move_messages(neomailbox.Tech)

-- spam by vipul razor to Spam/ (offered by neomailbox.net)
-- header: X-SA-Status: Yes
spamvipul = res:contain_field('X-SA-Status', 'Yes')
spamvipul:move_messages(neomailbox.Spam)
```

Mit den Ergebnissen von vipul razor Antispam war ich aber nicht so 100%ig
zufrieden, weshalb ich anfing zu googeln und
[fand](https://gist.github.com/battlemidget/5758764) was ich suchte.

Imapfilter piped jede Mail des letzten Tages zu dem lokal laufenden
bogofilter.

``` lua
-- bogofilter spam mails to Spam
res = neomailbox.INBOX:is_newer(1)
spam = Set {}
unsure = Set {}
  for _, mesg in ipairs(res) do
        mbox, uid = unpack(mesg)
        text = mbox[uid]:fetch_message()
        -- subject = mbox[uid]:fetch_field('subject')
        -- print(subject)
        flag = pipe_to('/usr/bin/bogofilter', text)
        if (flag == 0) then
          table.insert(spam, mesg)
        elseif (flag == 2) then
          table.insert(unsure, mesg)
        end
  end
neomailbox['INBOX']:move_messages(neomailbox['Spam'], spam)
```

H00ray!

### Domains

Relativ unspektakulär, habe ich meine Domains zu
[inwx.de](https://www.inwx.com/en) umgezogen.  Preise und Webinterface sind
okay. Und falls ich mal Lust habe eine Art DynDNS selbst zu bauen haben die
auch gleich eine API.

Alles in allem kann ich mit dem neuen Setup gut Leben. Mehr OpenBSD,
weniger Linux.
