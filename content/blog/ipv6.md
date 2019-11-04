---
aliases:
- /archives/363
date: '2008-12-17T21:54:39'
slug: ipv6
tags:
- "bin\xE4r"
- hexadezimal
- protocol
- network
- ip
- ipv6
- hex
- hardware
- icmpv6
- linux
title: IPv6
---

Mal ein Eck (den Vorhersagen nach, garnichtmehr soweit entfernter) Zukunft.
Schulbedingt habe ich schon länger damit zu tun und ehrlichgesagt finde ich
das Prinzip richtig genial. Vielleicht gibts einige unter den Lesern die
noch nicht mit allen Details der Zukunft vertraut sind. Dieser Artikel ist
also 2-fach von Nutzen. Für mich zum Einprägen und für den Rest der hier
liest :)

#### Wie sieht sowas eigentilch aus?

Eine IPv6-Adresse ist sensationelle 128Bit lang, eine IPv4 Adresse nur
32Bit. Und zwar sieht sie in etwa so aus:

    2003:0100:4008:0000:0000:00CD:A10B:4E87

8 Blöcke lang. Mit jeweils 4 Hexadezimalen Zahlen durch ":" getrennt.

Das ist die komplette Korrekte schreibweise. Allerdings wird diese mit
diversen Schreibweisen noch durch vereinfacht.  Dazu mehr weiter unten.

#### Wie bekomme ich eine IPv6-Adresse?

Eine IPv6-Adresse zu bekommen ist garnicht so schwer. Eine Site die unter
IPv6 läuft besuchen und schon bekommt man (unter Linux und Windows Vista)
eine IPv6-Adresse zugewiesen.

#### Brauche ich dafür besondere Hardware?

Nein. IP, als Protokoll, spielt sich auf der Schicht 3 des ISO/OSI Modells
ab.Die Netzwerkkarte befindet sich auf dem Layer 2 und bekommt vom
IP-Protokoll überhauptnichts mit. Sie überträgt weiterhin fröhlich ihre
Daten.  [Weitere Infos zum ISO/OSI Modell hier](http://de.wikipedia.org/wiki/OSI-Modell)

#### Was hat das ganze für Vorteile?

So gesehen wurde IPv6 entwickelt um die Adressknappheit von IPv4
abzuschaffen. Anfang 2010 soll es nach Aussagen der IANA(Das Institut das
IP-Adressen vergibt) keine öffentlichen IPv4 Adressen mehr geben. Diese
wurden alle an Firmen/Provider vergeben und somit landet das Internet in
einer Art Sackgasse. Im Grunde sind alle 4.294.967.296 öffentlichen IPv4
Adressen aufgebraucht.  Ein neues Modell muss her. Tada...IPv6.

Es gibt 340 Sextillion IPv6 Adressen. Das sind

    340.000.000.000.000.000.000.000.000.000.000.000.000 Adressen

Das sind im Gegensatz zu IPv4 relativ viele. Das sind sogar so viele, dass,
wenn man die Erde mit IPv6 Adressen bedecken wollte, es auf **einem
QuadratZentimeter der Erde 667 Billiarden Adressen** passen würden. Auch
ansonsten bietet das neue Protokoll noch viele Vorteile wie z.B.
Sicherheit.

#### Ist das nicht etwas umständlich zu merken?

Ja in der Tat. Aber auch dafür wurde eine Lösung gefunden:

    2003:0100:4008:0000:0000:00CD:A10B:4E87

Ist einmal die Ausgangsadresse.  Als nächstes werden Blöcke mit
ausschliesslich Nullen gekürzt und ein "::" dafür eingesetzt

    2003:0100:4008::00CD:A10B:4E87

So würde das ganze in etwas aussehen.  Der nächste Schritt, hat wieder mit
Nullenkürzung zu tun.  Jetzt werden führende Nullen innerhalb der Blöcke
weggekürzt.

    2003:100:4008::CD:A10B:4E87

Somit wird das ganze doch relativ Kompakt.

#### Kompabilität mit IPv4?

Selbstverständlich. IPv4 Adressen können einfach eingebettet werden.

    0:0:0:0:0:ffff:192.168.0.1

umgerechnet ins dual-System:

    0:0:0:0:0:ffff:11000000.10101000.00000000.00000001

umgerechnet ins Hexadezimale System:

    0:0:0:0:0:ffff:C0A8:0001

und jetzt noch gekürzt:

    ::ffff:C0A8:1

#### localhost?

127.0.0.1 wurde ganz einfach ersetzt durch die **::1** Schön einfach finde
ich :)

#### Private Adressen? PAT/NAT System?

Vermisst da jemand die 192.168.0.1 ? :) PAT/NAT ist definitv Schnee von
gestern. Warum auch ? Wir haben genügend IP Adressen um jedem Sandkorn der
Erde eine eindeutige Nummerzuzuweisen. Wenn man überlegt dass das NAT/PAT
System aber auch nur zu Ausdehnung der IPv4 Adressen erdacht wurde ist das
aber nachzuvollziehen. Dennoch gibts auch bei IPv6 wieder Private IPs.

Beginnen mit fc00: in einem /7 er "Netz" Genauso Site-locale Adressen.
Ersetzt ULA (Unique Local Adress)

#### MutliCasting wie vorher?

Mh nicht ganz.

    ff01:: Interface local (Loopback)
    ff02:: Link local
    ff05:: site local
    ff0e:: global Multicast

Im Allgemeinen sollte die Umstellung schon lange erfolgt sein. Aus welchen
Gründen genau IPv6 immernoch der Exot ist kann ich nicht genau sagen. Aber
es steht schonmal fest das es irgendwann soweit ist. Ja. Keine IPv4
Adressen mehr.  :)

Und ehrlichgesagt freue ich mich jetzt schon drauf :) Wird zwar nichtmehr
ganz so leicht sich eine IP zu merken, aber dafür haben wir wieder Adressen
und noch einige Vorteile.

Morgen Prüfung über IPv6. Ich hoffe ich bin gut vorbereitet :) Drückt mir
die Daumen :)
