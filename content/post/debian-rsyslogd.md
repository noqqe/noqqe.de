---
date: 2009-10-12T19:23:22+02:00
type: post
slug: debian-rsyslogd
status: publish
comments: true
title: Debian | rsyslogd
alias: /archives/686
categories:
- Coding
- General
- Linux
---

Man erzählt hier ja sonst nichts weiter, warum also nicht mal ein Stück weit Aufklärung. Eigentlich habe ich bei den Bash-Skripten immer eigene Logs geschrieben. Via Stout in ein File ausgegeben. Ziemlich simpel und stupide zugleich. Die zentrale Logverwaltung übernimmt in Debian Lenny ein Daemon. Der Spass heisst ziemlich Daemon untypisch (haha-.-) rsyslogd.

Es besteht also die Möglichkeit diesem Daemon mit einem Programm Meldungen zu übergeben. Das da "logger" heisst. Logger leitet (unter angabe verschiedener Handlingdaten) an den syslog Daemon Nachrichten weiter, die dieser dann anhand von Dringlichkeit und Quelle einordnet.
Dringlichkeiten wären zb.
```
0       Emergency
1       Alert
2       Critical
3       Error
4       Warning
5       Notice
6       Informational
7       Debug
```

Absteigend sortiert. Ausserdem lässt sich eine Quelle definieren. Welches Programm/Dienst übermittelt diese Nachricht?
```
0       kernel messages
1       user-level messages
2       mail system
3       system daemons
4       security/authorization messages
5       messages generated internally by syslogd
6       line printer subsystem
7       network news subsystem
8       UUCP subsystem
9       clock daemon
10       security/authorization messages
11       FTP daemon
12       NTP subsystem
13       log audit
14       log alert
15       clock daemon
16       local0
17       local1
18       local2
19       local3
20       local4
21       local5
22       local6
23       local7
```


So ergibt sich eine Schreibweise wie zb:

```
logger -p local0.err -t FILEBACKUP Files Backup failed
```


mit der Quelle local0 und Stufe Error wird nun das Thema FILEBACKUP mit dem Inhalt "Backup failed" an den Daemon.
Das wäre jetzt so eigentlich nicht so das phänomenale Überfeature. Allerdings lässt sich unser rsyslogd sagen was wer wie wann wo und warum mit welcher Meldung aus welcher Quelle und mit welcher Stufe er Logmeldungen in welches File verarbeiten soll.

rsyslog.conf
```
*.err  /var/log/error.log
```

 gibt zb alle Fehlermeldungen mit Stufe Error oder höher in das nachher angegebene File aus.
kernel-messages.* /var/log/kernelmessages, muss ich glaub ich nicht erläutern.

Interagieren mit Logs gefällt mir auf die Weise aufjedenfall besser als echo "ERROR" >> /tmp/schauichniewiederan
