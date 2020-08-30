---
title: "Gleichzeitigkeit messen mit Hyperfine"
date: 2020-08-30T10:06:23+02:00
tags:
- hyperfine
- golang
- go routines
---

Bildbearbeitung ist ein tolles Thema geworden. Seit dem Mehrteiler hier im
Blog über
[abandoned-tuples](https://noqqe.de/blog/2020/01/29/abandoned-tuples/) in
Python ist noch ein bisschen was passiert. Ich habe ungefähr verstanden wie
bestimmte Effekte herbeizuführen sind beispielsweise.

<!--more-->

Während des Lockdowns habe ich dann angefangen Go zu lernen und habe dazu
auch die Bildbearbeitung nach Go portiert. Das entstandene Projekt heisst
[nept](https://github.com/noqqe/nept). Nept ist 50% einen Backronyms. Noch
unklar was es bedeuten soll.

So weit, so gut.

# Go Routines

Warum habe ich das alles gemacht? Neben anderen Features wie Tests, Modules
und Interfaces wollte ich auch Gos eingebaute Gleichzeitigkeit mit [Go
Routines](https://gobyexample.com/goroutines) ausprobieren.

Der
[Diff](https://github.com/noqqe/nept/commit/b833d1eb86fbe6dd29c7486627c30930e85e5c37)
war schnell eingebaut (und auch echt Spass gemacht!) aber danach nicht so
ganz gewusst wie ich den Unterschied messen kann.

{{< figure src="/uploads/2020/08/diff.png" caption="Go Routines" >}}

Anstatt das mit `time` per Hand zu messen, gibt es das Tool
[hyperfine](https://github.com/sharkdp/hyperfine). Es hat tolle Features und
rechnet gleich aus wie viel schneller/langsamer 2 Kommandos im Vergleich
sind.

Ich habe mein Git Repository also zurückgesetzt und die alte Version von
`nept` nocheinmal kompiliert.

{{< figure src="/uploads/2020/08/hyperfine.png" caption="Vergleich Go vs Go mit Gleichzeitigkeit" >}}

Und tatsächlich, es sind 5 statt 6 Sekunden.

Zum Vergleich habe ich dann auch noch die Go Implementierung gegen die
Vormalige Python Variante dagegen gehalten:

{{< figure src="/uploads/2020/08/govspy.png" caption="Vergleich Python vs Go" >}}

Und der Unterschied hierbei krass, ganze 8 mal schneller ist die Go Variante.

Tolles Tool, kann ich jedem empfehlen der mal Tools vergleichen muss.
