---
date: 2010-06-24T21:32:29+02:00
type: post
slug: nagios-ricoh-drucker-tonerstand-per-snmp-abfragen
status: publish
comments: true
title: Nagios | Ricoh-Drucker Tonerstand per SNMP abfragen
aliases:
- /archives/1066
categories:
- General
- Linux
- PlanetenBlogger
tags:
- auffüllen
- drucker
- nagios
- ricoh
- snmp
- stand
- status
- toner
---

[Schon wieder Nagios](http://zwetschge.org/blog/?p=954). Diesmal aber nur als kleine Notiz für mich. Vor kurzem hab ich erst die Zählerstände der Papierfächer in unser firmeninternes Nagios eingebunden. Dasselbe funktioniert natürlich auch mit den Tonern.

Beschreibung der Fächer mit snmpwalk abholen:
```
snmpwalk -Os -c public -v 1 192.168.1.200
```

```
mib-2.43.11.1.1.6.1.1 = STRING: "Toner Schwarz"
mib-2.43.11.1.1.6.1.2 = STRING: "Resttoner"
mib-2.43.11.1.1.6.1.3 = STRING: "Toner Cyan"
mib-2.43.11.1.1.6.1.4 = STRING: "Toner Magenta"
mib-2.43.11.1.1.6.1.5 = STRING: "Toner Gelb"
```


Status der Toner als Integerwerte (0 = leer, -3 = voll)
```
mib-2.43.11.1.1.9.1.1 = INTEGER: 0
mib-2.43.11.1.1.9.1.2 = INTEGER: -3
mib-2.43.11.1.1.9.1.3 = INTEGER: -3
mib-2.43.11.1.1.9.1.4 = INTEGER: 0
mib-2.43.11.1.1.9.1.5 = INTEGER: -3
```


Kommando für Nagios konfigurieren:
```
define command{
command_name check_toner
command_line /usr/lib/nagios/plugins/check_snmp -H '$HOSTADDRESS$' -C  '$ARG1$' -o mib-2.43.11.1.1.9.1.$ARG2$ -w '$ARG3$': -c '$ARG4$':
}
```


Service für den Host einbinden:
```
define service {
use generic-service ; Name of service template to use
host_name druckerxyz
service_description TONER YELLOW
check_command check_toner!public!5!2!1
}
```

