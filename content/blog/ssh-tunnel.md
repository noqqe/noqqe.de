---
date: 2008-10-09T13:43:16+02:00
slug: ssh-tunnel
comments: true
title: ssh-Tunnel
aliases:
- /archives/297
categories:
- Linux
- Administration
- Network
tags:
- Linux
- ssh
- ssh tunnel
- ubuntu
---

Bin ja richtig begeistert für was sich das alles verwenden lässt.

Problem war folgendes: Ich möchte an ein Netzwerkdrucker-WebInterface von
einem Kunden in einem internen Netzwerk(Von außen nicht erreichbar)um
Einstellungen vorzunehmen. Wenn möglich ohne hinzufahren. Ein
Arbeitskollege hat mir dann geraten einen ssh-Tunnel aufzubauen.

    ssh -L lokalerport:zielrechner:zielport login@gateway(firewall)

Bedeutet im Klartext:

Ich baue eine Verbindung zum Gateway auf(`login@gateway(firewall)`in diesem
Falle die Firewall des Kunden) und steuere über diese eine beliebige
Interne IP(`zielrechner)` an. Dann kann ich den Port 80 (`zielport)(`HTTP -
eben fürs Webinterface) der Drucker IP auf einen frei wählbaren
Port(`lokalerport)` auf localhost legen und anschließend simpl im Browser
unter http://localhost:port abrufen! Und das alles mit diesem einen Befehl
:)

Enter gedrückt und schon kann ich den Internen NetzDrucker vom Kunden
bequem vom Arbeitsplatz aus konfigurieren.

Geht übrigens mit allem was nen Webinterface hat... Router ...Switches ...
Modems... und wenn der Port stimmt sogar noch andere Scherze... also eine
riesen Erleichterung.
