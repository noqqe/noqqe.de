---
date: 2010-08-25T20:06:47+02:00
type: post
slug: gitosis-zweischneidigkeit-des-auth-verfahrens
comments: true
title: Gitosis | Zweischneidigkeit des Auth-Verfahrens
aliases:
- /archives/1208
categories:
- Development
- Debian
- git
- Linux
- PlanetenBlogger
tags:
- auth
- denied
- deny
- dsa
- ERROR:gitosis.serve.main:Need SSH_ORIGINAL_COMMAND in environment
- git
- gitosis
- gitosis-serve
- problem
- PTY allocation request failed on channel 0
- publickey
- pulic
- rsa
- ssh
- sshlogin
---

Ein [Nachtrag](/archives/1175) und zugleich ein ganz besonders unschöner Zustand kam mir gestern unter die Finger. Gitosis benutzt bekanntermaßen SSH-Public-Keys zum authentifizieren der User, die in Git-Repositories arbeiten dürfen. Dieser Austausch zwischen Reporitory und Arbeitskopie passiert ebenfalls über SSH-Port 22. Die Benutzer, die sich dort anmelden, dürfen allerdings keinen direkten SSH-Zugriff bekommen. Soweit die Theorie.

Wenn man seinen Public-Key also Gitosis zur automatischen Authentifizierung vorwirft, wird man in das System der Git-Benutzer eingespeißt.

```
cp id_rsa.pub ~/gitosis/keydir/user@host.pub
git add keydir/*
git commit -a -m "user hinzugefügt"
git push
```


Bei erneuter Anmeldung an das System passiert folgendes:

```
$ ssh user@gitserver.domain.com
PTY allocation request failed on channel 0
ERROR:gitosis.serve.main:Need SSH_ORIGINAL_COMMAND in environment.
Connection to git closed.
```


Ich darf mich also nicht mehr einloggen. Bin ich normaler Benutzer, der wirklich nur mit git arbeiten darf, ist das auch gut so. Denn so wird die Sicherheit des Systems gewahrt. Bin ich allerdings Administrator des git-Remote-Servers sieht das anders aus. Ich habe ab diesem Zeitpunkt keine Möglichkeit mehr mein System (auf gewohntem Wege) zu pflegen.

Die Verbose-Ausgabe von ssh lässt darauf schließen was passiert:

```
$ ssh -v user@gitserver.domain.com
debug1: Authentications that can continue: publickey,password
debug1: Next authentication method: publickey
debug1: Trying private key: /home/user/.ssh/identity
debug1: Offering public key: /home/user/.ssh/id_rsa
debug1: Remote: Forced command: gitosis-serve user@host
debug1: Remote: Port forwarding disabled.
debug1: Remote: X11 forwarding disabled.
debug1: Remote: Agent forwarding disabled.
debug1: Remote: Pty allocation disabled.
debug1: Server accepts key: pkalg ssh-rsa blen 277
debug1: read PEM private key done: type RSA
debug1: Remote: Forced command: gitosis-serve user@host
debug1: Remote: Port forwarding disabled.
debug1: Remote: X11 forwarding disabled.
debug1: Remote: Agent forwarding disabled.
debug1: Remote: Pty allocation disabled.
debug1: Authentication succeeded (publickey).
PTY allocation request failed on channel 0
ERROR:gitosis.serve.main:Need SSH_ORIGINAL_COMMAND in environment.
debug1: client_input_channel_req: channel 0 rtype exit-status reply 0
debug1: client_input_channel_req: channel 0 rtype eow@openssh.com reply 0
debug1: channel 0: free: client-session, nchannels 1
Connection to git closed.
```


Die Authentifizierung mit meinen Public-Key klappt zwar, aber ich werde in eine gitosis-serve ssh-session gezwungen und damit bleibt mir der ssh-zugang ins System verwehrt. Nicht mit dieser Situation rechnend, starrte ich völlig perplex auf mein Terminal und die Reverse-Engeneering-Abteilung in meinem Kopf ratterte vor sich hin. Was passiert da und warum passiert das? Und vor allem: Wie komme ich jetzt wieder auf den Server?



### Solve it!


**1. Public-Key Auth deaktivieren**
Ohne PubKey Auth, wird der ssh-daemon nicht erkennen, das er mir eine git-serve session geben müsste. Dem lokalen ssh-client beizubringen sich nicht mit dem Public-Key am entfernten System anzumelden, wäre also eine Lösung (aber keine Schöne). Folgende Konfiguration führt dazu.
```
$ vi ~/.ssh/config
Host git
HostName gitserver.domain.com
User root
pubkeyauthentication no

```


**2. Different User**
Die Alternative zu dieser dauerhaften Veränderung ist (wenn vorhanden) einen anderen Benutzer zu verwenden um sich ins System einzuloggen und erst anschließend zu root zu werden.

**3. gitosis-serve zurechtstutzen**
Nachdem der Zugriff auf das System  wiederhergestellt ist, gehts zum Bugfix (gitosis-serve). gitosis muss diesen Umstand in irgendeiner ssh-config erzwingen. Ich verstehe nicht ganz warum, aber gitosis schrieb mir diese Änderungen in /root/.ssh/authorized_keys.

```
command="gitosis-serve user@host",no-port-forwarding,no-X11-forwarding,no-agent-forwarding,no-pty ssh-rsa AABBB3NzaC1yc2EAAAABIwAAAQEAyjwZCinCmB4oJJZ4RuiSqrQmiYE8+C+JKpTmiPkdfojUbiB9gm3BOhsYAdu99vP7yDOaIqg9e2dk/4HGm+P8obUR7lVrinMf5NvoRkOa8EfGdPJRz4ABOGRDte454bwestyWlvLhnKyWd+a9lU07siDJg5b1NbitIXkXa76V+lGMrqkixaDC6meZQEjZlxnVMpgzC5wyEQy2cVwUnX+Swiw68gsHsMYKBNsiVgNQ7nY8fa5lhV13E6L2aYAIorVpudS1bTiQfvfXCpVtJkJVSNPP6RzUtuSSErhsqOn1o2QtVjWhH5J/Y0D1b4eeEAgmdhq7554kQupJ9LgRww== user@host
```


Dieser Eintrag ist für das Verhalten verantwortlich. Auskommentieren oder entfernen aller Parameter bis ssh-rsa fixt das Problem . Happy Committing.
