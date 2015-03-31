---
date: '2008-09-15 22:24:15'
layout: post
slug: zwetschgeorg-monitoring
status: publish
comments: true
title: zwetschge.org monitoring
alias: /archives/246
categories:
- Hardware
- Linux
tags:
- debian
- monitoring
- Netzwerk
- server
- zwetschge
---

Habe mich heute mal mit [collectd](http://collectd.org/) und[ rrdtool](http://oss.oetiker.ch/rrdtool/) befasst und eine art Monitoring für meine [zwetschge eingerichtet](http://zwetschge.org).

Dank eines [kleinen Tutorials von Trinec](http://trinec.de/index.php/2008/08/17/einfaches-monitoring-statistiken-wie-mrtg/) war das aber ganz einfach ;)

Im groben kurz nochmal:

* rrdtolls und collectd installieren

* /etc/collectd/collectd.conf anpassen

* /etc/monitoring anlegen und mit :


``` bash
#!/bin/bash
# Webstatistiken erzeugen auf Webserver - Umwandlung von collectd zu Grafiken
# Speicherort der Grafiken auf einem Webserververzeichnis
GPATH=/var/www/monitoring/

# Speicherort der collectd Daten die von RRDTool umgewandelt werden
DPATH=/var/lib/collectd/

# Anzeige der Load (Tagesstatistik)
rrdtool graph ${GPATH}load-d.png -w 300 -h 100 -t "Tagesstatistik Load" --start -1d DEF:shortterm=${DPATH}load.rrd:shortterm:AVERAGE LINE1:shortterm#00ff00:Kurz DEF:midterm=${DPATH}load.rrd:midterm:AVERAGE LINE1:midterm#0000ff:Mittel DEF:longterm=${DPATH}load.rrd:longterm:AVERAGE LINE1:longterm#ff0000:Lang

# Anzeige des Arbeitsspeichers (Tagesstatistik)
rrdtool graph ${GPATH}memory-d.png -w 300 -h 100 -t "Tagesstatistik Memory" --start -1d DEF:used=${DPATH}memory.rrd:used:AVERAGE LINE1:used#ff0000:Benutzt DEF:free=${DPATH}memory.rrd:free:AVERAGE LINE1:free#00ff00:Frei

# Anzeige des Traffics auf eth0 (Tagesstatistik)
rrdtool graph ${GPATH}traffic-eth0-d.png -w 300 -h 100 -t "Tagesstatistik Eth0" --start -1d DEF:outg=${DPATH}traffic-eth0.rrd:outgoing:AVERAGE LINE1:outg#ff0000:Ausgehend DEF:inc=${DPATH}traffic-eth0.rrd:incoming:AVERAGE LINE1:inc#00ff00:Eingehend
```

befüllen.

* in die /etc/cron.d/monitoring und */5 * * * * root sh /etc/monitoring eintragen

Schon aktualisiert sich alle 5 min der Ticker.

Danke Trinec! und ab jetzt auch in Blogroll zu finden ;)
