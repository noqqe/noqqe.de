---
aliases:
- /blog/2013/09/03/statistical-und-showterm-dot-io
comments:
- author: Martin
  content: <p><a href="http://Showterm.io" rel="nofollow">Showterm.io</a> ist zudem
    auch noch <a href="https://github.com/ConradIrwin/showterm.io" rel="nofollow">quelloffen</a>
    :)</p>
  date: '2013-09-03T22:21:44'
- author: noqqe
  content: <p>Yay!</p>
  date: '2013-09-05T05:43:59'
date: '2013-09-03T19:20:00'
tags:
- recording
- shell
- stats
- showterm
- terminal
- barchart
- github
- statistical
- graph
- bar
- ascii
- bash
title: statistical und showterm.io
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

[Showterm](http://showterm.io/0d0f510cb43b206350679#fast)

Installation und Bedienung sind unendlich einfach.
Was man bekommt ist ein Link zu einer Site, die das aufgenommene mit HTML5
Zauber und etwas JS abspielt. Was sich dann wie oben über einen iframe einfach
einbetten lässt.

Einzig was mich stört ist, dass bei viel Output (gerade bei der v-barplot Loop
gen Ende) der Inhalt manchmal sehr schleichend ausgegeben wird. Hier nochmal der
[Direktlink](http://showterm.io/0d0f510cb43b206350679).
