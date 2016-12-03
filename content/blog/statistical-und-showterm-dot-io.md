---
title: "statistical und showterm.io"
date: 2013-09-03T21:20:00+02:00
comments: true
tags:
- Stats
- osbn
- Shell
- showterm
- bash
- ascii
- terminal
- recording
- statistical
- graph
- barchart
- bar
- github
aliases:
- /blog/2013/09/03/statistical-und-showterm-dot-io
---

Glaubt man dem `git log` wars wohl irgendwann 2011 als ich
[meinte](/blog/2011/04/14/statistical-statistiken-visualisieren-im-terminal/)
, ich bräuchte jetzt umbedingt eine Möglichkeit Barcharts in meinem
Terminal darstellen zu können.

Ich bastelte in `bash` die beiden kleinen Skripte die im GitHub Repository
[statistical](https://github.com/noqqe/statistical) sind.

Was damals sehr uncool war, war das ich (abgesehen von einem gif) nichts tun
konnte um schön darzustellen wie das Ergebnis aussieht. Abgesehen vom Pasten der
Ausgabe in das readme.markdown File, was wie man sieht total hart gescheitert
ist.

Im fancy Internet denken sich aber Leute ziemlich coole Dinge wie
[showterm.io](http://showterm.io) aus, mit denen man nicht den bloated Weg
des "Video aufzeichnens" gehen oder gar total kaputte animierte Bildformate
ausliefern muss. Was generiert und übertragen wird sind einfache
Characters. Das erzeugt fast keinen Overhead und funktioniert auch noch
ohne langes herumrechnen.

<iframe src="http://showterm.io/0d0f510cb43b206350679#fast" width="670" height="510"></iframe>

Installation und Bedienung sind unendlich einfach.
Was man bekommt ist ein Link zu einer Site, die das aufgenommene mit HTML5
Zauber und etwas JS abspielt. Was sich dann wie oben über einen iframe einfach
einbetten lässt.

Einzig was mich stört ist, dass bei viel Output (gerade bei der v-barplot Loop
gen Ende) der Inhalt manchmal sehr schleichend ausgegeben wird. Hier nochmal der
[Direktlink](http://showterm.io/0d0f510cb43b206350679).
