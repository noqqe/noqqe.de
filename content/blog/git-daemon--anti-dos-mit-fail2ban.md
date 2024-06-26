---
aliases:
- /archives/991
date: '2010-04-21T09:03:29'
tags:
- development
- daemon
- git
- gitweb
- fail2ban jail
- git-daemon
- verbieten
- "einf\xFCgen"
- filter
- fail2ban
- linux
- ban
- dos
- unban
title: Git-daemon | Anti-DOS mit fail2ban
---

Seit gestern versuche ich mittels Fail2ban zu vermeiden, dass mein
git-daemon "geDOSt" wird. Also bei ca 10 Downloads die Klappe für die IP
schliessen.  unter:

```
git clone git://zwetschge.org/roborobo.git
```

lässt sich über den Daemon ein Repo auschecken. Der Logeintrag bei Access
sieht folgendermaßen aus:

```
2010-04-20_11:37:52.05907 [16810] Connection from 200.200.200.200:54283
```

Nun lässt sich über ein einfaches Script, beispielsweise:

``` bash
for i in $(seq 1 100) ; do git clone git://zwetschge.org/roborobo.git gitrepodos$i ; done
```

den Server total auslasten. Gerade bei grossen Repos wäre das fatal.  Über
RegExp und fail2ban-regexp lässt sich der Ausdruck im Logfile auch filtern
und testen:

```
fail2ban-regex '2010-04-20_19:52:01.41131 [26818] Connection from 200.200.200.200:54283' '.*Connection from <HOST>:.{4,5}$'
```

gibt zurück der Ausdruck würde matchen. Meine Filterregel sieht auch
dementsprechend aus:

```
failregex = .*Connection from <HOST>:.{4,5}$
```

und der Eintrag in der jail.local (damit als jail erkannt wird):

```
[git-daemon]
enabled  = true
port     = git
filter   = git-daemon
logpath  = /var/log/git-daemon/current
maxretry = 5
```

ist eingerichtet. Keine Fehler im Fail2ban-Log. Alles erfolgreich
gestartet. Aber der Filter matched einfach nicht wenn ich das Script
ausprobiere.

Statusabfrage des Jails:

```
$ fail2ban-client status git-daemon

Status for the jail: git-daemon
|- filter
|  |- File list:    /var/log/git-daemon/current
|  |- Currently failed:    0
|  `- Total failed:    0
      - action
|- Currently banned:    0
|  `- IP list:
- Total banned:    0
```

Ich weiss mir gerade nicht zu helfen. Obwohl ich das selbe Spiel mit vsftp
auch gemacht habe. [LINK](/?p=904)
