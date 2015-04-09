---
date: 2010-10-17T12:14:05+02:00
type: post
comments: true
title: Bash | the zombie-revolution-environment
aliases:
- /blog/2010/10/17/bash-the-zombie-revolution-environment
- /archives/1314
categories:
- Development
- Debian
- Shell
tags:
- git
- 8bit
- bash
- battle
- brawl
- environment
- game
- gameplay
- git
- github
- humans
- revolution
- shell
- simulator
- zombie
- zombies
- zre
---

Wer kennt eigentlich noch [MUD](http://en.wikipedia.org/wiki/MUD)'s?  Diese
textbasierten Rollenspiele über Telnet. Uralt und genau genommen dürfte mir
das auch nichts sagen, denn das war vor meiner Zeit.

Wie dem auch sei. Vor kurzem saß ich nachts vor meiner Shell und suchte
irgendeine Ablenkung. Dachte dann an ein [Python Projekt](/archives/118),
welches ich zusammen mit Chris vor ca. 2 Jahren mal aufgegriffen hatte. Ich
fand die Idee damals schon gut, aber es mangelte einfach an Python
Kenntnissen.

Ich habe dann (auch wenn es schon spät war) versucht, das nochmal in Bash
abzubilden. Nur ein bisschen interessanter. Eine kleine Welt kam dabei
heraus, wie schon damals ausgedacht. In dieser Welt passieren Dinge, die
durch Zufall ausgewählt und gewichtet werden, was sich auf die Bewohner
dieser "Welt" auswirkt. In meiner jetzigen Version sollte allerdings
bisschen was passieren. Irgendwer musste da kämpfen oder sowas. Ich ließ
also Zombies und Humans auf dieser Welt gemeinsam leben. Es werden neue
Lebewesen geboren, sie sterben wieder, greifen einander an und so weiter.

{{< figure src="/uploads/2010/10/ice-cream-zombie.jpg" >}}

[](/uploads/2010/10/ice-cream-zombie.jpg)Im Klartext habe ich nichts
anderes in dieses Bash-Script gepackt als die Anzahl der Menschen/Zombies
auf der Welt und eine Hand voll Funktionen definiert, die "Events"
darstellen, die von einer Schleife und einem Case bei jeder Runde zufällig
ausgewählt werden. Das war's auch schon. Eigentlich hat mich mehr der Grad
an Zufälligkeit interessiert, der den Verlauf der erstellten Welt
beeinflusst, wie viele Zombies greifen wie viele Menschen auf einer
imaginären Farm an, können sich die Menschen wehren oder werden Sie von der
Übermacht der Zombies einfach überrannt? Wie viele Opfer gibt es? Wer
rottet wen zuerst aus?

Das war letzten Endes auch der Punkt der mich an ZRE (so hab ich das Script
genannt) immernoch unterhält und amüsiert. Jedesmal wenn ich es starte,
beginnt ein total neuer Verlauf dieser Welt, wer wann wo wen angreift und
letzten Endes besiegt.

Ich hab das zombie-revolution-environment auf github hochgeladen. Ich weiss
nicht wirklich was es ist, aber es unterhält mich.

[http://github.com/noqqe/zombie-revolution-environment](http://github.com/noqqe/zombie-revolution-environment)
[zombie-revolution-environment als RAW](http://github.com/noqqe/zombie-revolution-environment/raw/master/zre.bash)

Usage ist relativ einfach.

```
$ wget http://github.com/noqqe/zombie-revolution-environment/raw/master/zre.bash
$ ./zre.bash
```

[Ein beispielhafter Verlauf](http://github.com/noqqe/zombie-revolution-environment/blob/master/zre.example)
klemmt auch in dem github Repository.

Für neue Ideen, Bugreports, Meinungen, Flamewars oder Shitstorms bin ich
wie immer dankbar.
