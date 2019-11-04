---
title: "Die Fish Shell"
date: 2018-01-28T10:16:54+01:00
tags:
- fish
- bash
- shell
---

Seit 2 Monaten ist die [fish Shell](https://fishshell.com) sowohl @work als
auch zuhause meine Login Shell. Dazu gekommen ist es durch
[yogan](http://zogan.de) und [Mullet](https://twitter.com/mulletti) die
ebenfalls umgestiegen sind und mich mit dem Thema angefixt haben.

So eine neue Shell ist schon eine Umstellung. Seitdem ich Linux mache heisst
eben diese Shell für mich Bash. Ab und an hatte ich mal versucht mich mit
[zsh](https://www.zsh.org/) anzufreunden, hab es aber immer gleich wieder
verworfen.

Es dauert natürlich eine ganze Zeit bis man sich dort häuslich einrichtet. Die
Configs die man seit Jahren mit sich herumschleift, die Funktionen und Aliases.
Das schreckt erstmal ab. Ausserdem dachte ich Jahre das ich mich nie an eine
Shell gewöhnen könnte, die anders ist als die auf all den Servern auf denen ich
den ganzen Tag herumfuhrwerken muss. Scheinbar irrte ich.

Die Migration der Skripte ging ich eins nach dem anderen an. Im Grunde von
[nano-bash](https://github.com/noqqe/nano-bash) nach
[fishstaebchen](https://github.com/noqqe/fishstaebchen) migrieren. Ja. Sorry
für den Namen. Wenn man die beiden Repos vergleicht, kann man auch gut sehen
wie sich die Code Teile im Vergleich zu Bash verändert haben. Eigentlich nur
marginal.

Auch ansonsten gewöhnte ich mich überraschend schnell an das etwas andere
Verhalten von `fish`. Größter Blocker war `ctrl+r` für die Rückwärtssuche.
Dafür habe ich aber dann die gängige Lösung mit
[fzf](https://github.com/junegunn/fzf) gewählt und in die Config eingebaut.

```
set -U FZF_DEFAULT_OPTS "--height 5 -e --inline-info"
```

{{< figure src="/uploads/2018/01/fzf.png" >}}

Wenn ich nun `ctrl+r` drücke, kann ich (halbwegs) wie gewohnt durch meine
History durch. Ich muss aber dazu sagen, dass ich mit `fish` die Rückwärtssuche
deutlich weniger benutze, da die Completion so schön ist.

{{< figure src="/uploads/2018/01/completion.png" >}}

Der markierte Teil in lila(?) wird mir einfach vorgeschlagen und mit `ctrl-e`
(yay emacs..) übernommen.

Extrem schön finde ich auch den Teil der Autocompletion welcher Optionen gleich
erklärt. Lieblingsbeispiel dafür ist natürlich `gpg`.

{{< figure src="/uploads/2018/01/gpg.png" >}}

Aber darauf brauche ich jetzt nicht weiter eingehen da es wohl eines der Main
Features ist, weswegen Leute `fish` benutzen.

Ein bisschen Strange finde ich die Webinterface Konfiguration und dieses
`fishd` File in dem dann dynamische Config Ergebnisse daraus landen, aber ich
hatte noch keine Zeit zu verstehen/recherchieren ob/wie es sich abschalten
lässt.

Die Funktionalität Plugins mittels
[fisher](https://github.com/fisherman/fisherman) zu installieren empfand ich
als sehr angenehm, da ich eine ähnliche Vorgehensweise schon in meinem `vim`
habe. So kann ich mir via

```
fisher barnybug/docker-fish-completion
```

einfach die Tab Completion für `docker` holen, so ich sie denn brauche.
Außerdem zu Erwähnen ist noch die (wesentlich) angenehmere Implementation von
[z](https://github.com/fisherman/z). Ich nutzte `z` auch schon in der Bash
aber durch doofes verhalten nie wirklich sinnvoll.

Im Moment bin ich mit `fish` sehr zufrieden. Wie lange das so bleibt lasse ich
mal offen. Wahrscheinlich so lange bis mir wieder etwas anderes zwischen die
Finger kommt. Aber das ist ja normal.
