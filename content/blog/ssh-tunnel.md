---
aliases:
- /archives/297
comments:
- author: Thomas
  content: '<p>Irgendwie muss ich das falsch verstanden haben. Ich will jetzt aus
    Lust und guter Laune das ganze mal am Router testen, indem ich dan lokalen Port
    8080 auf den Port 80 vom Router (192.168.220.1) umleite. Meine Schlussfolgerung
    war dann folgende:<br>[thomas@jane ~]$ ssh -L 8080:192.68.220.1:80</p><p>aber
    als Antwort kriege ich:<br>usage: ssh [-1246AaCfgKkMNnqsTtVvXxY] [-b bind_address]
    [-c cipher_spec]<br>           [-D [bind_address:]port] [-e escape_char] [-F configfile]<br>           [-i
    identity_file] [-L [bind_address:]port:host:hostport]<br>           [-l login_name]
    [-m mac_spec] [-O ctl_cmd] [-o option] [-p port]<br>           [-R [bind_address:]port:host:hostport]
    [-S ctl_path]<br>           [-w local_tun[:remote_tun]] [user@]hostname [command]</p><p>Bin
    der Meinung, dass genau dass doch dort steht:<br>           [-i identity_file]
    [-L [bind_address:]port:host:hostport]</p><p>Wo ist mein Fehler?</p>'
  date: '2008-10-09T20:42:16'
- author: Dr. Azrael Tod
  content: "<p>Ja, diese kleinen Unix-Tools haben schon was...<br>Gut mit SSH l\xE4sst
    sich \xFCbrigens Screen kombinieren (wie mit jedem anderen Shellzugriff auch).
    Dann noch diverse Desktoptools auf Konsolenversionen (irssi, rtorrent, mcabber,
    ...) umstellen und schon hat man \xFCberall seine Programme griffbereit.<br>Manche
    Router lassen sich \xFCbrigens auch per ssh als unix-system fernwarten, wird noch
    viel flexibler ;-)</p>"
  date: '2008-10-10T10:40:10'
- author: noqqe
  content: '<p>@Azrael: Jepp - diese Linksys WRT dinger sind absolut genial :)<br>@Thomas:
    Gateway?! Du musst doch ssh irgendwasgeben wohin du connecten willst O_o</p>'
  date: '2008-10-10T13:33:59'
date: '2008-10-09T11:43:16'
slug: ssh-tunnel
tags:
- network
- linux
- administration
- ssh
- ubuntu
- ssh tunnel
title: ssh-Tunnel
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
