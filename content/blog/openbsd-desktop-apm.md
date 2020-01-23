---
title: "OpenBSD Desktop - APM"
date: 2019-12-27T14:19:01+01:00
tags:
- apm
- OpenBSD
---

Ein weiterer Dienst mit dem man so eigentlich nichts zu tun hat, wenn man
sonst nur OpenBSD Server betreibt. Das "Advanced Power Management Control
Program" hat diverse Features die mich anfangs ueberrascht hatten.

```
> apm
Battery state: high, 69% remaining, 145 minutes life estimate
A/C adapter state: not connected
Performance adjustment mode: manual (2501 MHz)
```

So kann ich beispielsweise mittels `zzz` in den Suspend bzw mit `ZZZ` in den
Hibernate Modus schalten. `apmd` quatscht dann mit der Hardware und leitet
das entsprechend ein.

Wenn man zum Beispiel eine Statusbar oder so etwas bauen und nur den
Prozentsatz der verbleibenden Batterielaufzeit haben will:

```
> apm -l
69
```

APM kontrolliert ausserdem auch die CPU Power. Die Performance auf 100%
aufdrehen geht zum Beispiel so:

```
apm -H
```

Nach der Installation hatte ich mit der CPU Power aber auch ein Problem. APM
hatte diese auf `1` gestellt, sodass ich nur mit 1% der Leistung unterwegs
war.

```
$ doas sysctl hw.setperf=100
hw.setperf: 1 -> 100
```

In welcher Art und Weise der Daemon diesen Wert bestimmt verstehe ich noch
nicht ganz, das muss ich noch herausfinden. Was ich jetzt aber unternommen
habe ist, die Flags via `rcctl` zu setzen.

```
apmd_flags=-H -t 30 -Z 5
```
Also, CPU Performance mit 100% starten, alle 30 Sekunden den Batterie Status
loggen und bei 5% verbleibender Batterie die Maschine in Hibernate Modus
schicken. Ziemlich geil!

Ich will auch wenn ich meinen "Deckel" zu mache das mein Bildschirm gesperrt
ist. Ich kann via folgenden Dateien


```
/etc/apm/suspend
/etc/apm/hibernate
/etc/apm/standby
/etc/apm/resume
/etc/apm/powerup
```

Scripte hinterlegen die ausgefuehrt werden, sobald eines der folgenden Events
den `apmd` triggern. Um nun den Bildschrim zu locken, baue ich erst

```
xidle -timeout 300 -program "/usr/local/bin/slock"
```

Was nebenbei den Charme hat, das automatisch nach 5 Minuten auch mein
Bildschirm gesperrt wird.  Um das Locking beim Decek zuklappen zu triggern
schicke ich in `/etc/apm/hibernate/` via kill noch ein Signal an xidle, um
das Lock Programm sofort zu triggern.

```
pkill -USR1 xidle
```

APM ist fuer mich eins der besten Beispiele warum es eine gute Idee ist
Kernel und Userland an der selben Stelle entwickelt werden. Der Daemon tut
allerlei Dinge. So eine einheitliche Erfahrung bzgl. Device Management hatte
ich bei Linux selten. Und ich feiere es sehr.


