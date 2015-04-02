---
date: 2010-12-23T16:16:03+02:00
type: post
slug: mac-fraise-app-fur-latex-konfigurieren
comments: true
title: Mac | Fraise App für LaTeX konfigurieren
aliases:
- /archives/1437
categories:
- Bash
- Development
- Mac
- PlanetenBlogger
- Web
tags:
- app
- Apple
- Build
- Fraise
- Latex
- Mac
- R
- Shortcut
- Smultron
- Tastenkürzel
---

Nachdem meine 30 Tage von [TextMate App](http://macromates.com/) ausgelaufen sind, habe ich mich irgendwie an das Apfel+R für LaTeX setzen, durchbauen und PDF anzeigen gewohnt.

Allerdings hab ich keine 50$ auf der hohen Kante um dir TextMate zu kaufen. Aber es geht ja auch anders. Was früher [ Smultron](http://en.wikipedia.org/wiki/Smultron) hieß und mittlerweile vom eigentlichen Entwickler eingestellt wurde, gibt es heute als Fork unter dem Namen Fraise.

[Fraise](https://github.com/jfmoy/Fraise/) ist OpenSource und eigentlich ein schöner schmaler Editor für Mac OS X.  Meine vermisste Funktion des Apfel+R lässt sich auch hier "nachbilden". Mit Apfel+B werden alle Kommandos angezeigt die mit Apfel+R verfügbar sind. Standardmäßig gibt es dort auch einen Satz von Dummies und sonstigem Nützlichem.

{{< figure src="/uploads/2010/12/Screen-shot-2010-12-23-at-15.09.171.png" >}}

Neue Collection namens LaTeX angelegt und das Command-Set "Build LaTeX" eingerichtet. Alles weitere ist im Endeffekt nur ein Shell-Skript, welches das PDF erstellt, die überflüssigen Dateien aufräumt und das generierte pdf mittels "open" öffnet.

Extrem hilfreich und billiger als TextMate. Das Skript habe ich bei github gepasted.


