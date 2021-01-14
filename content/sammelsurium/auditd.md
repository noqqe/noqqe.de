---
title: "auditd"
date: 2021-01-14T16:36:28+01:00
tags:
- Software
- auditd
---

Unter Linux gibt es einen Auditing Daemon. Konfigurierbar und aus 2 Teilen
bestehend. Kernel Modul und ein Userland Daemon.

<!--more-->

Es kann folgende Kategorien monitoren

* System calls: See which system calls were called, along with contextual information like the arguments passed to it, user information, and more.
* File access: This is an alternative way to monitor file access activity, rather than directly monitoring the open system call and related calls.
* Select, pre-configured auditable events within the kernel

## Userland

Temporär deaktivieren, nützlich bei `Backlog exceeded` Machine Freezes.

    auditctl -e 0

Temporär alle Regeln deaktivieren

    auditctl -D

Oder einfach den Daemon stoppen

    systemctl stop auditd

Das heisst der Kernel nimmt die Messages zwar noch auf, aber werden nicht
peristiert.

Status von `auditd` anzeigen

```
> auditctl -s
enabled 0     # <----- this means that auditd logging is disabled
failure 1
pid 478
rate_limit 0
backlog_limit 64
lost 0
backlog 0
loginuid_immutable 0 unlocked
```

## Kernel

Da das Kernelmodul `audit` geladen werden muss, kommt in die `grub` Config:

```
audit_backlog_limit=256 audit=1
```

## Rules

File

```
-w /etc/shadow -p raw -k access
-w /var/log/auth.log -p aw -k access
```

Events ([Liste](https://access.redhat.com/articles/4409591#audit-record-types-2))

Syscalls disablen

```
-a always,exit -F arch=b32 -S init_module,finit_module -F key=module-load
-a always,exit -F arch=b64 -S init_module,finit_module -F key=module-load
-a task,never
```

## Was wird geloggt?

Die letzten Events kann man einfach über das Logfile lesen.

    /var/log/audit/audit.log

Es gibt ein tool `aureport` mit dem alle Einträge ausgewerten werden können

```
> aureport --start today --event --summary -i

Event Summary Report
======================
total  type
======================
46260  SYSCALL
45  SERVICE_STOP
14  LOGIN
14  USER_ACCT
14  CRED_ACQ
14  USER_START
14  USER_END
13  CRED_DISP
4  SERVICE_START
2  USER_LOGIN
1  CRED_REFR
1  SYSTEM_SHUTDOWN
1  DAEMON_END
1  CWD

> aureport --start this-week
```

## Links

* [capsule8.com](https://capsule8.com/blog/auditd-what-is-the-linux-auditing-system/)
