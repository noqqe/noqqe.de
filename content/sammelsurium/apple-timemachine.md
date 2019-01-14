---
title: Apple TimeMachine
date: 2011-06-10T14:47:55
tags:
- Software
- Apple
- Timemachine
---

Das was ich zuvor erwähnt hatte, und was ab 10.6.3 nicht mehr funktioniert
war kein Konfigurationshack, sondern einfach die Möglichkeit für die
"Sparsebundles" die von TimeMachine angelegt werden eine feste Größe
anzugeben. Seit 10.6.3 wird bei jedem Backup diese Einstellung
überschrieben. Es gibt aber wohl einen Workaround, den ich persönlich noch
nicht getestet habe.

1. Time Machine erstellt für das Backup ein Sparsebundle (mitwachsendes
   Image) in welchem die Daten abgelegt werden.
2. Mit dem Befehl hdiutil kann man die maximale Größe dieses Images neu
   setzen, dies geht über den Terminalbefehl "hdiutil resize -size 100g
   `<sparsebundle name>`". Time Machine am besten vorher abschalten und darauf
   achten dass das Image gerade nicht gemountet ist. Wenn das Image bereits
   größer als die gewünschte Zielgröße ist, kann man vorher in der normalen
   Time Machine Oberfläche alte Backups löschen und mit "hdiutil compact
   `<sparsebundle name>`" das Image verkleinern.
3. Um zu verhindern dass diese maximale Größe beim nächsten Backup einfach
   überschrieben wird, muss man die Info.plist und Info.bckup Files die am
   Sparsebundle hängen gegen Schreibzugriff sperren. Dies geht am
   effektivsten in dem man eine Sperre darauf setzt. Hierfür gibt es den
   Befehl SetFile: "SetFile -a L sparsebundle/Info.*"

Es kann sein das SetFile nicht standardmäßig vorhanden ist, sondern
zusammen mit den Developer Tools installiert wird. Auf meinem MBP ist der
Befehl da, ich habe aber auch XCode installiert.

Mehr dazu in diesem [Thread](http://discussions.apple.com/thread....readID=2383738)
