---
layout: post
title: "Trennung vom bash-it Framework - Koche nun eigenes Süppchen"
date: 2012-08-12T20:12:00+02:00
comments: true
categories:
- Bash
- Coding
- Debian
- ubuntuusers
keywords: "bash-it, bash, framework, nano-bash, nano, bashrc, config, tricks,
tipps, oh-my-zsh, zsh, github, startup"
---

Manchmal muss man das auch. [Einige](/blog/2010/11/24/bash-it-n0qorg-theme-und-git_info/)
[Zeit](/blog/2010/12/07/github-mitarbeit-an-bash-it/) hab ich
Hirnschmalz in das Community Framework [bash-it](https://github.com/revans/bash-it) gesteckt.

Anfangs mit der aktiven Weiterentwicklung beschäftigt, hab ich später eher
massiv Aufwand betreiben müssen um die neuen Dinge mit meiner Version kompatibel zu halten.
Man konnte bash-it früher entweder komplett als eine Art Distribution für die
Bash benutzen, oder sich wie in einer Art Supermarkt an einzelnen Schnipseln,
Funktionen und kleinen Skripten bedienen die man gut fand. Und das war super.
Man konnte Dinge lernen, Schnippsel von anderen adaptieren und hatte immer
Trickreiche Funktionen am Start.

{% img center /uploads/2012/08/frameworks.jpg %}

Der Rip-Off von [oh-my-zsh](https://github.com/robbyrussell/oh-my-zsh) hat sich
aber mit der Zeit (und Dank der großen Aktivität der Community) m.E. zu einem
für mich nicht mehr funktionierendem Konzept entwickelt. Konkret:

* Die Überarbeitung des Plugin Systems mit Meta Informationen war ein riesen
Fehler
* Aktivierung und Deaktivierung einzelner Komponenten ist unnötig und verwirrend
für neue User (und für mich auch.. :P )
* Etwas, dass ich gerne den "theme-color-vcs-prompt-complex" nenne. Aber das
würde den Rahmen sprengen...
* Der Wust aus 1. Mio Dingen, den mittlerweile niemand (mehr)
braucht/versteht
* Fehlende Ziele/grobe Richtung für die Entwicklung
* Durch die stark wachsende Komplexität funktioniert bash-it nichtmal mehr als
Supermarkt Bedienkonzept.
* Ca. 5 Sekunden bis zum Prompt

All diese Probleme und Punkte haben mich dazu veranlasst, meine eigene Suppe zu
kochen, die ich in meiner gnadenlosen Unkreativität
[nano-bash](https://github.com/noqqe/nano-bash) genannt habe. Es hat einen
ganzen Nachmittag gedauert die wenigen Dinge die ich wirklich brauche aus bash-it zu
extrahieren und aus Funktionen zu lösen, damit das Ganze in eine neue (für mich
pflegbare) Form passt. Denke das spricht auch schon für sich...

Schade :(
