---
aliases:
- /blog/2010/06/24/nagios-ricoh-drucker-tonerstand-per-snmp-abfragen
- /archives/1066
comments:
- author: Anonymous
  content: "Hallo !  Erstmal vorweg danke f\xFCr die tollen Infos :) ... Habe heute
    begonnen Ricoh Drucker in unser Nagios System einzubinden und bin bei der Recherche
    auf Ihren Artikel gesto\xDFen.  \nLeider scheint es bei mir nicht ohne weiteres
    zu funktionieren.\nIch erhalte als R\xFCckgabewert f\xFCr den check wenn ich ihn
    genauso anlege: \n\nExternal command error: Bad Operator (INTEGER): At line 73
    in /usr/share/mibs/ietf/SNMPv2-PDU \n\nSagt Ihnen das zuf\xE4llig etwas ? W\xE4re
    hier f\xFCr jede Hilfe dankbar ! :)"
  date: '2015-08-20T16:18:13.746999'
- author: Anonymous
  content: "So wie es aussieht hab ich es hinbekommen. Ich erhalte jetzt einen R\xFCckgabewert.
    \nLeider ist der Wert aber immer der selbe  (  3 )  . Ich habe es mit dem vorgeschlagenen
    MIB  auf 2 Ger\xE4ten getestet und jeweils erhalte ich egal wie der Tonerstand
    ist immer den gleichen R\xFCckgabewert... Was mache ich noch falsch ?"
  date: '2015-08-20T18:46:38.448844'
- author: noqqe
  content: "Hast du mal nachgeschaut ob du wirklich die richtige OID abfragst? Der
    Post ist immerhin 5 Jahre alt, vielleicht haben sich die OIDs vom Hersteller ge\xE4ndert?\n\nHab
    da auch lange gesucht. Am Besten mal nach den Strings greppen die du suchst und
    dann dr\xFCber oder drunter nachschauen, ober der Status dazu passt.\n\nWahlweise
    auch etwas Trial and Error Verfahren anwenden. Anfangs hab ich den Toner mal Rausgezogen,
    dann wieder SNMP angefragt und geschaut welcher Wert sich \xE4ndert."
  date: '2015-08-21T09:05:29.092663'
- author: jigsaw
  content: "mib-2.43 ist die Standardadresse f\xFCr Printer. Also wird sich an der
    Adresse relativ wenig \xE4ndern. Was bei relativ vielen Druckern vorkommt, dass
    man einerseits die mib-2.43 abfragen kann und zus\xE4tzlich noch eine Herstellerspezifische
    Adresse.\nWenn ich einen neuen Drucker finde, bem\xFChe ich immer 2 Abfragen mit
    snmpwalk\nEinerseits 'snmpwalk -Os -c public -v public <Host> .1.3.6.1.4.1' \nDamit
    kriege ich als Ausgabe in der Regel die Herstellerspezifischen Adressen. Da findet
    man ab und zu auch mal den Tonerstand, je nach Hersteller. (Hier muss man wirklich
    t\xFCfteln oder google f\xFCr entsprechende Guides/Mib-Datenbanken bem\xFChen,
    weil der Tonerstand mal \xFCber 2 Werte als %Wert errechnet werden muss, mal als
    %-Wert dasteht oder eben wirklich nur der \"Nagios\"-Zustand [0,1,3])\nAndererseits
    'snmpwalk -Os -c public -v public <Host> .1.3.6.1.2.1'\nDort bekommt man neben
    Haufenweise Informationen \xFCber Schnittstellen eben auch die Universellen PrinterOID's
    geliefert."
  date: '2015-08-21T10:33:17.947086'
date: '2010-06-24T19:32:29'
tags:
- status
- drucker
- snmp
- toner
- administration
- devops
- nagios
- blog
- ricoh
- stand
- linux
- "auff\xFCllen"
title: Nagios | Ricoh-Drucker Tonerstand per SNMP abfragen
---

[Schon wieder Nagios](http://zwetschge.org/blog/?p=954). Diesmal aber nur
als kleine Notiz für mich. Vor kurzem hab ich erst die Zählerstände der
Papierfächer in unser firmeninternes Nagios eingebunden. Dasselbe
funktioniert natürlich auch mit den Tonern.

Beschreibung der Fächer mit snmpwalk abholen:

``` bash
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
