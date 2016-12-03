---
aliases:
- /archives/55
date: '2008-05-13T09:12:06'
slug: eeepc-urlaubstuchtig-media-tauglich-machen
tags:
- hardware
- xmms
- xubuntu
- totem
- linux
title: "EeePC urlaubst\xFCchtig / Media tauglich machen!"
---

Wie man sicher schon merkt mutiere ich hier immer mehr zum EeePC Blogger :) Aber
ich kann versichern das das nur momentan so ist.

### MP3 Player

Vorinstallierter Player ist [Totem](http://de.wikipedia.org/wiki/Totem_(Programm)).
Gefiel mir natürlich nicht so wirklich. Aber gut. Zuerstmal den Codec installiert:

```
apt-get install libxine1-plugins
```

Wollte aber nicht so recht. Genau wie andere Pakete. Aber hat nicht lang
gedauert bis ich geschnallt hab das eeeXubuntu aus irgendeinem Grund
Paketquellen "aus dem Internet herunterladbar" kein einziges Häkchen hatte.

  * Von Canoncial unterstützte OpenSourceSoftware
  * Von der Gemeinschaft betreute OpenSourceSoftware
  * Proprietäre Gerätetreiber
  * Urheberrechtlich eingeschränkte Software
  * Quelltext

Nichts von alle Dem war angekreuzt. Also gut. MP3s liefen und um nicht
Totem verwenden zu müssen hab ich mich dann für xmms entschieden. Hab ich
noch nicht probiert, soll Platz sparen und handlicher sein. Musik läuft
schonmal.

### Flash Player

Das ganze ging auch recht einfach von der Bühne. Zuerst den
Adobe-Flash-Player mit Hilfe der Browser Plugin Bar installiert und dann
noch den Gnash SWF Viewer.  Erst nach dem 2. ging dann alles. Warum ? KA.
Aber es geht.

### Externe Festplatte

_"_Cannot mount volume"_ Invalid mount option when attempting to mount the
volume "Externe"._

Warum nicht? Darum muss ich mich umbedingt noch kümmern.

### Videos

Auch schnell gegessen:

```
apt-get install libxvidcore4
```
Ausprobiert hab ich das jetzt zwar noch nicht... aber wird sich zeigen sobald
ich meine Externe zum laufen gebracht habe auf der mein ganzes Zeug liegt!