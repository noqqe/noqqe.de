---
layout: post
title: "graylog2 Hacks"
date: 2012-08-28T20:24:00+02:00
comments: true
categories:
- planetenblogger
- Web
- Debian
- Ubuntu
- CentOS
- Coding
---

Eigentlich sollte im Titel eher "Tipps und Tricks" stehen, aber mit Hacks im
Titel fühl ich immer so gut. Seit neuestem hab ich graylog2 für mich entdeckt.
Im Grunde ein Syslog Server der Logs in einer mongodb speichert und via Elasticsearch und einem
Webinterface zur Verfügung stellt. Geiler Scheiss - gerade @ work.

## Einbindung in rsyslog

Nach der super dokumentierten Einrichtung von
[Mongodb](http://docs.mongodb.org/manual/tutorial/install-mongodb-on-linux/),
[Elasticsearch](http://www.elasticsearch.org/tutorials/2010/07/01/setting-up-elasticsearch.html),
[graylog2-web-interface](https://github.com/Graylog2/graylog2-web-interface/wiki/Installing-the-web-interface-on-Debian-5.0) und
[graylog2-server](https://github.com/Graylog2/graylog2-server/wiki/Installing)
steht die Konfiguration der Clients an.
Bei rsyslog ists ein Einzeiler um das Logging sowohl lokal als auch an
graylog2 zu schicken:

{% codeblock %}
*.* @graylog2.domain.example
{% endcodeblock %}

## Einbindung in syslog-ng

Wer SuSe/SLES/CentOS Kisten am laufen hat kommt leider nicht drum herum:

{% codeblock %}
destination graylog2 { udp("1.2.3.4" port(514)); };
log { source(src); destination(graylog2); };
{% endcodeblock %}

## Volle DNS Namen/FQDN auflösen

Hab relativ lange googeln müssen, da per default der Shortname via Syslog
übermittelt wird. Mit folgender Konfiguration versucht graylog2 den Hostnamen
via Rückwärtsauflösung zu ermitteln (Obacht bei Masquerading/SNAT!)

{% codeblock %}
$ vi /etc/graylog2.conf
[...]
force_syslog_rdns = true
[...]
{% endcodeblock %}

## Nützliche Queries in der Analytics Shell

Die Analytics Shell imho das Überfeature an graylog2. Die Queries finde ich aber
etwas... dürftig dokumentiert.

Zähle alle Logs die "Permission" beinhalten
{% codeblock %}
all.count(message = "Permission")
20:57:29 - Completed in 8.57ms - Count result: 137
{% endcodeblock %}

Zeige die letzten 500
{% codeblock %}
all.find(message = "Permission denied")
{% endcodeblock %}

Zähle alle Logs eines Hosts die PHP Fehler "undefined variable" enthält
{% codeblock %}
all.count(host = "web01.example.org", message = "Undefined variable" )
20:57:29 - Completed in 58.57ms - Count result: 13657
{% endcodeblock %}

Suche z.B. nach "CPU17: Core power limit notification (total events = 487717)"
{% codeblock %}
stream(503cd70cc029b530ef000015).find(message = "power")
16:22:13 - Completed in 339.36ms
{% endcodeblock %}

Generell hab ich aber das Gefühl, dass die Analytics Shell noch nicht so
wahnsinnig ausgereift ist. Facitily hab ich bis jetzt noch nicht geschafft
wirklich zu bedienen trotz Data Type Beschreibung auf der Wiki Page.
