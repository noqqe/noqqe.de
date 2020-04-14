---
date: '2016-12-03T09:30:16'
tags:
- razor
- pyzor
- bogofilter
- spam
title: jeezus Content Spam Filtering
---

Die ganze Geschichte mit [spamd](/blog/2016/10/23/greylisting-und-spamd/)
als Spamfilter ist schon etwas her. Nach bisschen über 2 Monaten bin ich
eigentlich ganz zufrieden wie es läuft mit dem eigenen Mail Setup.

`spamd` bleibt aber eben nur ein IP basierter Filter. Irgendeine Art von
Contentfiltering braucht man aber trotzdem. Zumindest war meine Mailbox
trotzdem ziemlich voll mit Mist. Da ich schon
[immer](https://noqqe.de/blog/2013/10/26/spammer-vs-statistik-mit-bogofilter/) ein Fan
von [Bayesian Spam Filtering](https://en.wikipedia.org/wiki/Naive_Bayes_spam_filtering)
war wollte ich das wieder benutzen. Und das serverseitig. Allerdings wollte
ich auch der riesen Bloatware Spamassassin aus dem Weg gehen.

[Bisher](/blog/2014/06/19/mehr-bsd/) hatte ich `imapfilter` benutzt um
Bogofilter clientseitig auszuführen. Das hat den gravierenden Nachteil das
Mails in die Mailbox kommen und erst danach weg sortiert werden. Im Zweifel
hab ich zu dem Zeitpunkt schon eine Notification auf einem Endgerät
erhalten.

Jedenfalls googelte ich wild durchs Netz nach serverseitigen Lösungen für
Bogofilter. Ich probierte `.forward` Files und sogar `maildrop` versuchte
ich in `opensmtpd` einzubetten. Was eigentlich funktionieren sollte aber
seltsame `authdaemon: s_connect() failed:` Errors wirft.

Letztendlich war ich dann gelangweilt und ich dachte wie schwer
das wohl sein kann einen Spam Daemon zu schreiben. Ich weiss, so fangen die
besten Geschichten an.

Stellt sich heraus: Mit den richtigen Libraries braucht man eigentlich fast
nichts tun. Ich bastelte `jeezus`. Das Python Script `bind()`ed sich auf
einen Port, spricht dort SMTP, liesst den Content ein, wendet Bogofilter
darauf an und sendet die Mail via SMTP wieder raus. Wie ein kleiner Proxy.

``` bash
./jeezus --listen 127.0.0.1:12001 \
  --pipe-command '/usr/local/bin/bogofilter -u -e -R' \
  --deliver 127.0.0.1:12002
```

Das kann man mit <150 Zeilen Python halt machen. Eingebettet wird das
eigentlich genauso wie Spamassassin in `opensmtpd`. Der relevante Auszug:

```
# Local delivery port for jeezus
listen on lo0 port 12002 tag Filtered

# Rule for mailboxes
accept tagged Filtered for domain <mailboxdomains> deliver to lmtp "/var/dovecot/lmtp" rcpt-to

# Filter through jeezus
accept from any for domain <mailboxdomains> relay via "smtp://127.0.0.1:12001"
```

Das ganze loggt schön nach `syslog`

```
Dec  2 13:19:05 aax jeezus[36563]: Results rcpt=wa1@noqqe.de
from=btv1==1448ffcb8c9==oliviaborreg123@gmail.com command=Spam pyzor=Spam
subj="URGENT MAIL FROM HOSPITAL "
Dec  2 13:59:57 aax jeezus[36563]: Results rcpt=wa1@noqqe.de
from=bounce-5144-68847596-248-248@iuguqf.tk command=Spam pyzor=Spam
subj="December Deals: UGG Boots Special Deal 60% OFF!"
```

Tatsächlich läuft der kleine Daemon seit bisschen mehr als einem Monat.
"[Works for me](https://events.ccc.de/2016/11/22/hello-this-is-33c3-works-for-me/)" :)

Zusätzlich hab ich noch [pyzor](http://pyzor.readthedocs.io) in das Skript
eingebaut. Vereinfacht gesagt erstellt `pyzor` einen Hash vom Mail Body und
gleicht diesen mit einem public Server ab. Die Erkennungsraten sind ganz
gut und der Einbau war wegen der Python Library auch sehr einfach.

Solang mir nicht anderes über den Weg läuft werde ich das wohl erstmal so
weiter betreiben. Eine Bayesian Library mit der ich die Erkennung direkt in
Python machen kann ohne einen externen Command zu callen wäre noch schön.
Mal sehen ob sich da was findet.

Ansonsten ist der Source auf [Github](https://github.com/noqqe/jeezus) zu finden.
