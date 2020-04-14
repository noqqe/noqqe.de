---
title: "iptables"
date: 2018-07-09T11:00:49+02:00
tags:
- Software
- iptables
---

Alle Regeln anzeigen

    iptables -nvL

Alle Regeln entfernen

``` bash
iptables -P INPUT ACCEPT
iptables -P FORWARD ACCEPT
iptables -P OUTPUT ACCEPT
iptables -F
iptables -X
iptables -t nat -F
iptables -t nat -X
iptables -t mangle -F
iptables -t mangle -X
iptables iptables -t raw -F
iptables -t raw -X
```

Redhat default iptables deaktivieren

    systemctl disable firewalld
