---
title: "Sudo"
date: 2018-03-28T13:19:06+02:00
tags:
- Software
- sudo
---

Vollen root Access für User konfigurieren

```
user ALL=(ALL:ALL) ALL
```

Vollen root Access für Gruppe konfigurieren

```
%gruppe ALL=(ALL:ALL) ALL
```

User ohne Passwort

```
user ALL=(ALL) NOPASSWD: ALL
```

User fuer einen bestimmten Command

```
user ALL=(ALL) NOPASSWD: /sbin/reboot
```
