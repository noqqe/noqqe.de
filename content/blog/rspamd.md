---
title: "Rspamd"
date: 2019-05-20T20:20:46+02:00
tags:
- spam
- mail
---

Zugegeben, ich war in den letzten Monaten etwas nachlässig mit meiner
privaten Infrastruktur. Was das Spamvolumen anging entstand aber zuletzt dringender
Handlungsbedarf.

2016 hatte ich zusätzlich zum IP-basierten
[bpgd-spamd](https://noqqe.de/blog/2016/10/23/greylisting-und-spamd/) eine
eigene kleine Software bastelt: [jeezus Content Spam Filtering](https://noqqe.de/blog/2016/12/03/jeezus-content-spam/).

Also bisschen Bayesian, bisschen Checksum, bisschen IP-Blocking. Scheinbar
reicht das aber nicht mehr aus und es ging wieder los. `jeezus` hab ich
wieder entfernt und suchte nun eine allumfassende Spamlösung. Nachdem mir
Spamassassin immernoch Angst macht, hatte ich [rspamd](https://rspamd.com)
ins Auge gefasst.

Ich habe gefühlt unglaublich viel Zeit versemmelt um zu verstehen wie man
diesen Wust aus Konfigurationsoptionen dazu bewegt zu tun was man gerne
möchte. Das hatte ich mir deutlich einfacher vorgestellt.

Nach etwas spicken bei
[vedetta-com/caesonia](https://github.com/vedetta-com/caesonia) habe ich es
jedoch geschafft eine Konfiguration zu basteln die mir gefallen hat (übrigens ohne
Redis Queue).

Spannend war auch die Anbindung an OpenSMTPD, welcher kurz vorher noch eine
neue Syntax bekommen hat.

```
table mailboxdomains { "noqqe.de" }

# Bind to interfaces
listen on lo0
listen on vio0 port 25 tls pki mail.n0q.org  # mailin
listen on vio0 port submission tls-require pki mail.n0q.org auth # relaying

# MDA Wrapper
mda wrapper dovecot "rspamc -h 127.0.0.1:11334 -t 120 --mime -e '%{mda}'"

# Actions
action "mbox" mbox alias <aliases>
action "relay" relay
action "rspam" mda "/usr/local/bin/lda2lmtp -r %{user.username} -s %{sender}" wrapper "dovecot"

# Allow redircect for local aliases (root, hostmaster ..)
match for local action "mbox"

# Filter through jeezus
match from any for domain <mailboxdomains> action "rspam"

# Allow local sending of mails
# This rule also allows authenticated users to relay.. somehow.
match auth from any for any action "relay"
```

Man hätte jetzt auch direkt via Dovecot LDA die Mail in die Mailbox delivern
können, irgendwie wollte ich aber `LMTP`. Deshalb habe ich da dieses kleine
Wrapper Script
[lda2lmtp](https://brnrd.eu/freebsd/2018-05-27/improving-my-mail-server-setup.html).
Losgelöst ob das jetzt eine schöne Lösung ist oder nicht, ist der verlinkte Blogpost jedenfalls lesenswert.

Auch der Bayesian Filter von `rspamd` will angelernt werden.
Dazu ein bisschen Shell Script via Cronjob.

```shell
for x in $(find ~/mail/.Spam/cur -type f -mtime -1) ; do
  rspamc learn_spam $x
done

for x in $(find ~/mail/cur -type f -mtime -1) ; do
  rspamc learn_ham $x
done
```

Für mich als Stats-Liebhaber gab Frei Haus noch ein paar Goodies geliefert.
Auch wenn ich mir das Webinterface auf `localhost` gebunden habe und nur ab
und zu via `ssh -L` mal auf den lokalen Host forwarde. Sicher ist sicher.

{{< figure src="/uploads/2019/05/rspamd1.png" >}}

{{< figure src="/uploads/2019/05/rspamd2.png" >}}

Was soll ich also sagen. Email Hosting ist mehr so Spaß am Basteln. Es
rechnet sich weder Zeitlich noch Qualitativ, die Freude am Lernen aber
bleibt.
