---
date: 2009-02-03T22:32:47+02:00
type: post
slug: teeworlds-server-howto
comments: true
title: Teeworlds - Server How To
aliases:
- /archives/459
- /blog/2009/02/03/teeworlds-server-how-to/
- /blog/2009/02/03/teeworlds-server-howto/
categories:
- Linux
- Administration
tags:
- codecocktail
- how to
- screen
- server
- ssh
- Teeworlds
- teeworlds server
- teeworlds server how to
- unpack
- virtual terminal
- zwetschge
---

Durch Crackpod und AtzeEgge von
[Codecocktail](http://codecocktail.wordpress.com) bin ich auf Teeworlds
aufmerksam geworden. [Crackpod der einen Aimbot dafür geschrieben
hat](http://crackpod.bplaced.net) und AtzeEgge der wie jeder normale Mensch
fair spielt ;) Das ganze ist ein extrem witzig und dynamisches 2D
ShooterSpiel geworden. Man könnte es auch Echtzeit Worms nennen :) Das es
auch noch OpenSource ist macht den Shooter nur noch attraktiver :) Es lief
letztlich drauf hinaus das mein Server zwetschge.org für einen GameServer
herhalten muss :D.

Der Server ist über die MasterServer von Teeworlds erreichbar unter dem
Namen **[JustForFun]zwetschge.org|noqqe.de ** und wird (was mich wundert)
mit gutem Ping relativ häufig bespielt(auch via IP-Connect mit
**zwetschge.org:8303** erreichbar). Sollte jemand das Bedürfnis bekommen
mich oder die anderen von CodeCocktail mal "ownen" oder "bashen" zu wollen
- kann dieser jemand das nun tun. Ich stell damit sogar die Mittel :).

Um aber mal zum technischen Teil zu kommen: Ist ein Teeworlds-Server sehr
einfach aufzusetzen.


  * Aktuelle Version runterladen :
    [http://teeworlds.com](http://teeworlds.com)
  * teeworlds0.x.x.tar.gz entpacken (wahlweise mit
    [unpack](http://zwetschge.org/unpack/) ;) keine Schleichwerbung
    eigentlich..)
  * mit dem Wechsel in das entsprechende Verzeichnis und ausführen "
    ./teeworlds_srv" kann der Server das erste mal gestartet werden
    (allerdings ohne Konfiguration! nicht zu empfehlen)
  * Um die gewünschte Konfiguration einzustellen gibt es eine Liste mit
    Config-Details auf
    [http://www.teeworlds.com/?page=docs&wiki=SettingUpAServer](http://www.teeworlds.com/?page=docs&wiki=SettingUpAServer)
  * "vim twserver.cfg" anlegen und  mit Konfigurationen befüllen. Beispiel
    Konfiguration:

```
sv_name Teeworlds sample dm
sv_map dm1
sv_scorelimit 20
sv_timelimit 10
sv_gametype dm
sv_rcon_password remember
sv_motd Teeworlds sample dm configuration
sv_max_clients 12
sv_spectator_slots 10
```

  * Den Server mit Configfile starten " ./teeworlds_srv -f twserver.cfg "
  * Happy Bashing.

So würde das gehen. Wenn da nicht noch ssh Session ins Spiel kommen würde.
Ich muss sagen das teeworlds_srv eine kleine Macke hat. Diese Executeable
öffnet einen Prozess. Schön oder? Wäre meine ssh-Session damit nicht
unbrauchbar wärs schön. Der Prozess spielt sich quasi im Vordergrund ab und
startet nicht wie üblich einen Hintergrundprozess mit dem ich weiter
arbeiten kann. Sobald ich nun aber mein Terminal Fenster schliesse... wars
das mit meinem Game-Server. Abhilfe schafft aber ein Paket namens screen
das im Debian Repository ist. Dieses Programm zielt genau auf solche Fälle
ab.

  * apt-get install screen
  * " screen  ./teeworlds_srv -f twserver.cfg " (öffnet den Vordergrund
    Prozess in einem Virtuellen Terminal in dem der TeeworldsServer jetzt
    läuft)
  * mit Strg+A+D kann man diesen Modus nun wieder verlassen und die SSH
    Session schliessen.
  * Angenommen man will den Server aber trotzdem nachher noch einmal
    begutachten: " screen -r " zeigt offene Screens an. So ist das Problem
    super gelöst

Der Server auf zwetschge.org bleibt denke ich längerfristig bestehen und
wird in die Projects hinzugefügt. Vielleicht sieht man sich ja mal :) Happy
Bashing
