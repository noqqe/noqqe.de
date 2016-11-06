---
date: 2010-06-13T14:23:20+02:00
comments: true
title: 'Bash | mailrequest '
aliases:
- /blog/2010/06/13/bash-mailrequest
- /archives/1042
categories:
- Development
- Linux
- Shell
tags:
- bash
- GTD
- mail
- script
- sick
- work
---

Mal wieder was aus der "Faule Sysadmins" und "wie mache ich mir _noch_
weniger Arbeit"-Sparte. Des öfteren rufen Kunden an die Probleme mit ihrem
Mailkonto haben. Meistens drehen sich die Probleme um die selben Themen.
Angeblich nicht versandte Mails, Login-Probleme oder um Postfach Grössen,
die zu klein, zu voll oder gross sind.

Die Prozedur ist immer die gleiche. Logfiles untersuchen. Quota ausfindig
machen. Mail nach verfolgen. Je nach User individueller Mailserver. Für die
grobe Arbeit, hab ich mir jetzt ein kleines Skript geschrieben, mit welchem
ich (easy-usage-like ;D) mir ca 5-6 Befehle spare. Ich nannte es
mailrequest. Die Kreativität hat mich nämlich einfach überrannt -_-

Das ganze läuft jetzt wie folgt ab.

```
mailrequest -s stichwort
#Nach Queue-id, Absender, Empfaenger im aktuellen Logfile suchen.
```

```
mailrequest -r stichwort
#Für alle Logfiles am Mailserver
```

```
mailrequest -q user@foobar.de
#Postfachgroesse ermitteln
```

Bis dahin hat der Plan ganz gut geklappt. Nur die Auswahl spezieller
Mailserver war mir noch nicht schön genug geregelt. Mailserver jedesmal per
Parameter mit übergeben müssen? Fand ich zu doof. Aber Hard-Coded im Source
ist fast noch ungünstiger gewesen. Den Mittelweg hat mir die Bash
ermöglicht.

```
$host=$3
$defaulthost=mail.domain.de
${host:-$defaulthost}
```

Klartext: Wenn kein 3. Parameter definiert (oder leer) ist, wird
automatisch der $defaulthost gewählt. Sollte der Mailserver nun vom
Standard abweichen, kann ich ihn mitgeben. Falls nicht, kann ich mir den 3.
Parameter des Aufrufs sparen.

```
mailrequest -s spam@zwetschge.org mail.zwetschge.org
```

Das ganze Script zum begutachten, Kritik äußern, anflamen, besser wissen:
[/uploads/2009/09/013](/uploads/2009/09/013)
