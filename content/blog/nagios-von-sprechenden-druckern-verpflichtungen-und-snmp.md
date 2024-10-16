---
aliases:
- /archives/954
- /blog/2010/04/02/nagios-von-sprechenden-druckern-verpflichtungen-und-snmp
comments:
- author: Christian
  content: "<p>Und woran erkennt der Drucker, das Papier nachgef\xFCllt wurde wenn
    der Schacht noch gar nicht leer war?<br>Anhand des Herausziehens des Schachtes
    oder wie geht das?</p>"
  date: '2010-04-04T02:50:49'
- author: noqqe
  content: "<p>Hallo Christian,</p><p>das kann ich dir nicht mit 100%iger Sicherheit
    sagen. Die Information die ich via SNMP abhole, wird vom Drucker bereits bereitgestellt.
    </p><p>Ich kann vermuten und denke, das es anhand der im Tray eine ungef\xE4hre
    Gewichtsmessung gibt. </p><p>Wie gesagt, nur Vermutung. :)</p>"
  date: '2010-04-04T14:16:34'
- author: Leser
  content: "<p>Hallo,</p><p>ein sch\xF6ner Artikel. Die gro\xDFen Multifunktionger\xE4te
    bieten im Webinterface teilweise auch schon die M\xF6glichkeit per Email bei bestimmten
    Ereignissen, z.B. kein Papier, Benachrichtigungen zu verschicken.</p>"
  date: '2010-04-04T15:59:38'
- author: noqqe
  content: "<p>Stimmt, w\xE4re aber dann kein Blogpost geworden. Auch kein Lerneffekt
    ;)</p>"
  date: '2010-04-04T21:40:23'
- author: Jan
  content: <p>Danke, </p><p>nette Idee, wenn man eh ein Nagios am Laufen hat...  ;-)</p><p>Jan</p>
  date: '2011-02-06T06:24:26'
date: '2010-04-02T17:24:51'
tags:
- fach
- "nachf\xFCllen"
- snmp
- mib
- administration
- devops
- nagios
- paper
- ubuntu
- printer
- papier
- tray
- debian
- oid
title: Nagios | Von sprechenden Druckern, Verpflichtungen und SNMP
---

![Logo_Nagios](/uploads/2010/04/Logo_Nagios.gif)

Ich teile mir in meiner Funktion als Auszubildender mit einem Kollegen die
ehrenvolle Aufgabe, ab und an den Papierstand unserer 4-Fächer Ricoh
Drucker bei Gelegenheit zu "überwachen". Im vorbeigehen sozusagen. Nun, das
ganze ist ein etwas undankbarer Job - denn egal wie oft man zum
Kontrollieren kommt - es kommt immer der Moment in dem nichts mehr im Fach
/ den Fächern ist, wenn Kollegen 200 Seiten am Stück ausdrucken.

Um an diesem Umstand etwas zu ändern kam uns neulich eine
"Notification"-Idee. In der Beschreibung des Druckers steht "SNMP-fähig".
Das [Simple Network Management Protokoll](http://de.wikipedia.org/wiki/Simple_Network_Management_Protocol)
bietet allerlei Möglichkeiten Informationen von Routern, Switches, Server
oder Desktoprechnern abzufragen. In unserem Fall bietet auch der Drucker
diese Möglichkeiten. Nach kurzer Recherche, scheint es etwas wie
[OIDs](http://de.wikipedia.org/wiki/Object_Identifier) /
[MIBs](http://de.wikipedia.org/wiki/Management_Information_Base) für
Drucker zu geben. Eine Liste dieser OIDs lässt sich mit snmpwalk ausgeben.

```
snmpwalk -Os -c public -v 1 192.168.1.200
```

Herauskommt eine _Menge_ an Informationen, mit der man erst mal umzugehen
wissen muss. Nach etwas suchen, fiel mir folgender Block ins Auge:

```
mib-2.43.8.2.1.10.1.1 = INTEGER: 55
mib-2.43.8.2.1.10.1.2 = INTEGER: 385
mib-2.43.8.2.1.10.1.3 = INTEGER: 55
mib-2.43.8.2.1.10.1.4 = INTEGER: 385
mib-2.43.8.2.1.10.1.5 = INTEGER: 0
mib-2.43.8.2.1.11.1.1 = INTEGER: 0
mib-2.43.8.2.1.11.1.2 = INTEGER: 0
mib-2.43.8.2.1.11.1.3 = INTEGER: 0
mib-2.43.8.2.1.11.1.4 = INTEGER: 0
mib-2.43.8.2.1.11.1.5 = INTEGER: 9
mib-2.43.8.2.1.12.1.1 = ""
mib-2.43.8.2.1.12.1.2 = STRING: "Briefbogen"
mib-2.43.8.2.1.12.1.3 = ""
mib-2.43.8.2.1.12.1.4 = ""
mib-2.43.8.2.1.12.1.5 = ""
mib-2.43.8.2.1.13.1.1 = STRING: "Tray 1"
mib-2.43.8.2.1.13.1.2 = STRING: "Tray 2"
mib-2.43.8.2.1.13.1.3 = STRING: "Tray 3"
mib-2.43.8.2.1.13.1.4 = STRING: "Tray 4"
```

Kurzes systematisches Trial-and-Error am Papierfach des Druckers und
Kontrolle der Veränderungen der SNMP-Ausgabe brachten mehr Aufschluss. Es
schien also so als würde ein Teil der SNMP-Ausgabe die geschätzten Werte an
Papier im Drucker zurückliefern.

Genau genommen die OIDs :

```
mib-2.43.8.2.1.10.1.1
mib-2.43.8.2.1.10.1.2
mib-2.43.8.2.1.10.1.3
mib-2.43.8.2.1.10.1.4
```

Damit lässt sich arbeiten. Für jedes Papierfach des Druckers eine
OID-Nummer. snmpget lässt ein Gerät explizit nach einer OID fragen, oder
ihren Descriptor.

    snmpget -v1 -Cf -c public 192.168.1.200 mib-2.43.8.2.1.10.1.1
    SNMPv2-SMI::mib-2.43.8.2.1.10.1.1 = INTEGER: 385

Im ersten Fach sind also 385 Blätter. Schätzungsweise. So ist es also
möglich den aktuellen Stand der Papierfächer abzufragen, ohne aufstehen zu
müssen. Das war schonmal was. Aber wie das jetzt mit Benachrichtigung
laufen lassen? CronJob? [Nagios](http://www.nagios.org/)! Wo sonst httpd's,
smb-Freigaben, Erreichbarkeiten oder sonstige Dienste abgefragt und
monitored werden, ließen sich auch die Papierfächer einbinden.

Zuerst die Drucker als Hosts einbinden, die Nagios überwachen soll:

    /etc/nagios3/conf.d/host-printer1.intern.cfg
    define host {
            host_name  printer1.intern.firma.de
            alias       printer1
            address     192.168.1.200
            use         generic-host
            }

Jetzt kennt Nagios den Host. Weiss aber weder was dort überwacht werden
soll, noch mit welchem Plugin und ab welchem Schwellwert es Alarm schlagen
soll. Nagios enthält ein Plugin für SNMP-Abfragen. Zu finden unter
/usr/lib/nagios/plugins/check_snmp. Diesem Plugin kann ich die selben
Fragen stellen wie mit snmpget. Einzige Veränderung: Schwellwerte für
Warning (-w) und Critical (-c) müssen mitgegeben werden. Ich erstellte also
einen sogenannten "Check".

    /etc/nagios3/conf.d/z_check_papertray:
    define command{
             command_name    check_papertray
             command_line    /usr/lib/nagios/plugins/check_snmp -H '$HOSTADDRESS$' -C '$ARG1$' -o mib-2.43.8.2.1.10.1.$ARG2$ -w '$ARG3$': -c '$ARG4$':
    }

Die -H Hostadresse wird aus der angegebenen Adresse im Hostfile  gewonnen
in dem der Check später als "Service" eingebunden wird. Außerdem noch jede
Menge Argumente die auf den ersten Blick vielleicht verwirren.

* Argument1: Community für die SNMP-Anfrage. In dem Fall "public"
* Argument2: Fachnummer mit der die OID ergänzt wird.
* Argument3: Warning-Wert bei dem Nagios Alarm schlägt.
* Argument4: Critical-Wert bei dem Nagios Alarm schlägt.

In der Hostdatei kann der definierte Check jetzt als Service eingebunden werden.

    define service {
            use                             generic-service
            host_name                      printer1.intern.firma.de
            service_description             PAPERTRAY 1 DinA4
            check_command                   check_papertray!public!1!150!56
            }

Das ganze 4 mal. Für jedes Papierfach einmal. check_command ist dabei der
ausschlaggebende Punkt. Nagios zieht aus jedem Wert nach ! seine Argumente.
Community public, Fach 1, Schwellwert Warning 150 und Critical 56 oder
kleiner.

Und im Falle eines Falles. Emails:

> ** PROBLEM Service Alert 2: printer1/PAPERTRAY 1 DinA4 is CRITICAL **
>
>     ***** Nagios *****
>     Notification Type: PROBLEM
>     Service: PAPERTRAY 1 DinA4
>     Host: printer1
>     Address:
>
>     192.168.1.200
>
>     State: CRITICAL
>     Date/Time: Thu Apr 1 16:30:42 CEST 2010
>     ACK by:
>     Comment:
>     Additional Info:
>     SNMP CRITICAL - *0*

Schön, denn wir werden jetzt immer benachrichtigt wenn kein Papier mehr im
Fach ist, oder es bereits aufgefüllt wurde. Und Chef fands auch gut :)
