---
title: Windows Zeitgeber Dienst
date: 2012-01-05T10:37:08
tags: 
- OS
- Windows
---

Zum Anhalten und starten des W32Time Dienstes kann man aufrufen:

    net stop w32time net start w32time

Um die Zeit manuell abzurufen, kann man folgenden Befehl nutzen:

    net time /SET

Dieser Befehl verlangt eine Bestätigung mit "J" vom Benutzer; dies kann
umgangen werden mit folgendem Befehl:

    net time /SET /YES

Der erwartet keinerlei Bestätigung (danke an Rene Frauchiger für den Tip).
Um die Zeit vom Domain Controller abzugleichen, egal, welche externe Quelle
gerade eingestellt ist, kann man aufrufen:

    net time /DOMAIN /SET

Auch hier ist analog die Bestätigung mit "J" nötig.

Möchte man die externe Zeitquelle konfigurieren per Kommandozeile, dann
kann man dies erreichen über:

    net time /setsntp:SERVERNAME

wobei Servername der Name oder die IP - Adresse des Server ist, der als
Zeitquelle dient. Die Dokumentation ist hier teilweise widersprüchlich, es
gibt zwei Quellen, die besagen, man könne hier eine Liste von Servern
eingeben, sie nennt aber kein Trennzeichen und auch keine
Iterationsvorschrift. Eine anderer Microsoft - KnowledgeBase - Artikel
weist darauf hin, daß diese Aussage falsch sei, und stets nur ein Server
konfiguriert werden könne.

Der Befehl "NET TIME" ist aber als deprecated erklärt worden und wird bald
nicht mehr unterstützt. Seit Windows XP gibt es den Befehl w32tm, der den
NET TIME - Befehl ersetzen soll. Hier ein Beispielaufruf:

    C:\>w32tm /config /syncfromflags:manual /manualpeerlist:ptbtime1.ptb.de,ptbtime2.ptb.de C:\>w32tm /config /update

Dieser Befehl legt unsere beiden PTB-Server als Zeitquelle fest und
aktualisiert die Konfiguration des W32Time - Dienstes.

    C:\>w32tm /resync

Erzwingt ein Abrufen der Uhrzeit vom eingestellten Zeitserver, und ein
resynchronisieren der lokalen Uhr. Erzeugt auch einen Eintrag im Systemlog.