---
aliases:
- /blog/2013/06/09/openbsds-pf-gegen-ssh-bruteforces
comments:
- author: zralb
  content: <p>Genau so wird das gemacht! :-)</p>
  date: '2013-06-09T15:28:57'
- author: Anonymous
  content: "Du kannst auch expiretables (Port security/expiretables) daf\xFCr benutzen,
    das tut das Ganze auch ohne cronjob"
  date: '2013-11-28T22:30:51.933127'
- author: noqqe
  content: ui, nice! das kenn ich garnicht, guter hinweis, danke!
  date: '2013-11-29T08:29:07.189039'
date: '2013-06-09T14:55:00'
tags:
- bruteforce
- openbsd
- network
- firewall
- sshguard
- pf
- ssh
- block
- fail2ban
title: OpenBSD pf gegen ssh Bruteforces
---

Unter Linux hatte ich meist
[fail2ban](http://www.fail2ban.org/wiki/index.php/Main_Page) benutzt um die Maschine gegen `ssh`
Bruteforce Attacken zu sichern. Für OpenBSD hatte ich bisher zu
[sshguard](http://www.sshguard.net/)
gegriffen, was aber Overkill ist. Ich muss nicht erst Logs parsen um zu wissen
wer zu häufig am ssh Port anklopft.

{{< figure src="/uploads/2013/06/itsyouagain.gif" >}}

## Block

Wenn ich an `pf` rumspiele, aktiviere ich einen Cronjob
der alle 10 Minuten `/sbin/pfctl -d` ausführt (denn er weiss nicht, was er da
tut!). So schliesse ich mich wenigstens nicht für immer aus wenn was schief
geht.

Die nachfolgende Regel bedeutet im Grunde, dass bei mehr als 5 Connections
in 50 Sekunden die Source-IP in den Table `<bruteforce>` eingepflegt wird.

``` bash
# Allow and track ssh brute force
pass in on $extif proto tcp from any to any port ssh \
  flags S/SA keep state \
  (max-src-conn 5, max-src-conn-rate 5/50, \
  overload <bruteforce> flush global)
```

Den aufgebauten Table `<bruteforce>` muss man aber auch noch verarbeiten.
Alle Einträge darin werden geblockt.

``` bash
table <bruteforce> persist
block quick from <bruteforce>
```

## Whitelist

Die Whitelist-Funktionalität in fail2ban lässt sich auch
mit nativem `pf` nachbauen.

```
table <admins> { 1.2.3.4/32 }
pass in on $extif from <admins> to any port ssh
```

Kurz gesagt, ein Table in dem sich eine SRC-IP befindet,
die von den `block` Rules unberührt bleibt. Aufpassen bei der Reihenfolge.

## Expire Tables

Fehlerhafte Logins passieren. Auf meinem oBSD loggen sich auch Menschen ein,
die sich von `gitolite` ein Repo abholen wollen. Nur weil mal jemand 5 Repos in einer
Minute haben wollte, ist das noch kein Grund ihn für immer auszusperren.

Seit OpenBSD 4.1 gibt es das native `expire` Feature für Tables.

``` bash
$ pfctl -t bruteforce -T show
   69.110.96.21

$ pfctl -vt bruteforce -T expire 3600
0/0 addresses expired.

$ pfctl -vt bruteforce -T expire 600
1/1 addresses expired.
D  69.110.96.21
```

Expired wird alle 10 Minuten oder öfter per Cronjob.

```
*/10 * * * * /sbin/pfctl -t bruteforce -T expire 3600
```

Der relevante Teil meiner Config nochmal auf [gist.github.com](https://gist.github.com/noqqe/5743740).
Und nicht vergessen den `--disable` Cronjob für `pf` nach basteln und
ausführlichem(!) Testen wieder abzuschalten. Sonst bringt das herzlich wenig.
