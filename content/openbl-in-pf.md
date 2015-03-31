---
layout: post
title: "OpenBL in pf"
date: 2015-02-04T20:31:00+02:00
comments: true
categories:
- osbn
- OpenBSD
- pf
- OpenBL
- Firewall
- Blocklist
---
Das Projekt [OpenBL](https://openbl.org), dass ursprünglich die OpenSSH
Blacklist war, hat auf der ganzen Welt verteilte Honeypots rumstehen und
publiziert regelmäßig Blacklisten sortiert nach Zeitraum und Protokoll des
Angriffs.

{{< figure src="/uploads/2015/02/assholes.jpg" >}}

Das Resultat daraus kann jeder frei benutzen und zum Beispiel in seine
Firewall einbauen. Realisiert hab ich das über einen `pf` Table.

```
table <blacklist> persist
pass in quick from <admins>
block quick from <blacklist>
```

Die Regel ist relativ unspektakulär, abgesehen davon, dass die pass-Regel
für die Admin IPs vor der Blacklist-Block Regel matchen sollte. Nur für den
Fall dass man mal selbst auf so einer Liste landen sollte.
Ein kleines Bash Script, dass mit Cron einmal die Woche die IPs updated.

``` bash 
#!/usr/local/bin/bash
STORE="/tmp/"
RULES="base.txt.gz"
TABLE="blacklist"
/usr/local/bin/wget -q https://www.openbl.org/lists/$RULES -P $STORE
gunzip $STORE/$RULES
pfctl -t $TABLE -T flush
pfctl -t $TABLE -T add -f $STORE/${RULES%.*}
rm $STORE/${RULES%.*}
```

Es gibt natürlich auch 10.000 andere Listen, die sowas bewerkstelligen.
Darunter auch ganze
[Wrapper](https://github.com/khainebot/Blocklist-Downloader/blob/master/blocklistdownloader.py)
die solche Listen parsen. War mir aber dann etwas zu hart overengineert.
