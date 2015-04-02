---
date: 2011-07-04T20:42:09+02:00
type: post
slug: minecraft-git-bash
comments: true
title: 'Minecraft + Git + Bash = '
aliases:
- /archives/1693
categories:
- Bash
- Coding
- Debian
- General
- git
- Ubuntu
- ubuntuusers
- Web
tags:
- Autosave
- bash
- git
- Minecraft
- Minecraft and git
- Minecraft und git
- Save
- Speichern
---

Seit mittlerweile erstaunlich langer Zeit spiele ich [Minecraft](http://minecraft.net). Minecraft hält seine Daten in **~/.minecraft** vor. Also Levels, Statistiken, Items. Das Minecraft Home Directory unter Versionskontolle zu stellen hat unter Umständen mehrere Vorteile, die ich hier kurz erläutern möchte :)

**Initiales Setup**

Als erstes muss das Verzeichnis initial eingerichtet werden. Initialisierung, hinzufügen aller Dateien und ersten Commit erstellen.

```
$ cd $HOME/.minecraft
$ git init
$ git add .
$ git commit -a -m "Initialer Commit"
```


**Spielstände manuell Laden und Verwalten (Commits)**

Einer der gravierendsten Vorteile. Wer wie ich oft an Klippen hinunter stürzt oder an einem (oder auch mehreren :P ) Creeper(n) scheitert wird das bestätigen können. Einmal gefallen/gestorben gibt es kein zurück mehr. Bis jetzt.

{{< figure src="/uploads/2011/07/minecraft-hill-500.png" >}}

Die hypothetische "Herausforderung" scheint sich gerade aufzutun. Ob jetzt Creeper, Berg oder sonst was ist erstmal egal. Könnte auf jeden Fall kritisch für meinen Minecraft Character enden.

```
$ git commit -a -m "Ob man den Sprung ueberlebt?"
```


Nach einem kurzen Tab in die Konsole, sollte das Spiel erstmal gesichert sein und ich kann den Sprung wagen.

{{< figure src="/uploads/2011/07/minecraft-gameover-500.png" >}}

Anscheinend überlebt man nicht, aber genau das war auch der kritische Punkt. Genau jetzt bin ich in der Lage meinen alten Spielstand wiederherzustellen. Mit nachfolgendem Kommando verwerfe ich alle seit dem letzten Commit entstandenen Änderungen an meinem Spielstand. Vorher dringend aufs Minecraft Titelmenü zurückkehren!

```
$ git stash
# Update
# oder alternativ:
$ git checkout -f
```


{{< figure src="/uploads/2011/07/minecraft-hill-again-500.png" >}}

Dieses Szenario lässt sich nicht nur auf gerade geschehene Ereignisse abbilden sondern auch zwischen Commits die längere Zeit her sind. Wenn nach einer halben Stunde/einem Monat klar wird, das der Minecraft Char gerade nur Müll verzapft hat, kann auch zwischen mehrere Commits hin und her gesprungen werden. Mit welchen git Kommandos das bewerkstelligt wird, bleibt jedem selbst überlassen.

**git-revert** macht den letzten Commit rückgängig, erstellt dabei einen neuen in dem die Änderungen enthalten sind. Das ist in soweit gefährlich, dass zwischenliegende Commits unberührt bleiben und eventuell in einen großen Haufen Datenmüll zerfallen(!). Eher Anwendung für den "Warp" an einen früheren Zeitpunkt X findet daher **git-reset**.

```
$ git reset 66a2594
# oder
$ git reset HEAD^
```


Das Working Directory wird damit auf einen Stand gebracht, wie es zum Zeitpunkt des angegebenen Commits aussah. Dieser kann somit auch weiter in der Vergangenheit liegen.

**Automatische Speicherung (Bash-Einzeiler)**

Allerdings muss ich zugeben, dass diese Praxis relativ schnell aufwendig wird. Immer zwischen Fenstern hin und her zappen ist ja auf Dauer auch eher zermürbend. Daher habe ich mir diese "Arbeit" von einer kleinen Bash Zeile abnehmen lassen.

```
$ SEKUNDEN=10 ; while true ; do git add . ; git commit -a -m "AutoSave $(date)" ; sleep $SEKUNDEN ; done
```


Ich denke es ist Geschmacksache wie oft bzw. in welcher Frequenz die Commits abgesetzt werden können. Bis jetzt bin ich mit ca 300 Sekunden (5 Minuten) am besten Gefahren. Die Commits rieseln vor sich hin und beeinträchtigen so in keinster Weise den Spielfluss.

```

[master bf9dd85] AutoSave Mo 4. Jul 19:56:17 CEST 2011
4 files changed, 15 insertions(+), 12 deletions(-)
rewrite saves/0pen_Running/level.dat (100%)
rewrite saves/0pen_Running/level.dat_old (100%)
[master 5ddedc8] AutoSave Mo 4. Jul 19:56:27 CEST 2011
4 files changed, 17 insertions(+), 19 deletions(-)
rewrite saves/0pen_Running/level.dat (100%)
rewrite saves/0pen_Running/level.dat_old (100%)
[master 2d33023] AutoSave Mo 4. Jul 19:56:37 CEST 2011
4 files changed, 10 insertions(+), 11 deletions(-)
rewrite saves/0pen_Running/level.dat (100%)
rewrite saves/0pen_Running/level.dat_old (100%)

```


**Parallele Welten (Branches)**

Um einfach mal ein Anwendungsbeispiel zu nennen: Wer in seinem virtuellen Minecraft-Keller mal raue Mengen an TNT gebunkert hat, möchte es nach Möglichkeit auch mal benutzen, right? Aber danach das ganze Dorf wieder aufbauen? Nee... Branching!

Das ist der Punkt an dem die Geschichte der lokalen Minecraft Map sich in zwei Teile spaltet. In einer wird das eigene Bauwerk Sodom und Gomorra mäßig untergehen und in der anderen weiterhin existierenden Welt tollen sich Pigs und Sheeps in Minecarts herum. Die Abzweigung lässt sich wie folgt bewerkstelligen.

```
$ git branch blowup
$ git checkout blowup
```


Jetzt kann man in aller Seelen Ruhe TNT verteilen und auch mal Destroyer statt Builder spielen. Tipp: Commit **vor** der Sprengung setzen :P Explosion immer und immer wieder von vorne genießen ;) Bemerkenswert sind außerdem die unterschiedlichen Abläufe von ein und der selben Explosion, aber dazu vielleicht wann anders ein Blogpost. Irgendwann wird aber auch das dann zur Routine und man wechselt via

```
$ git checkout master
```


wieder zu den Schäfchen. Der Branch "blowup" bleibt aber bestehen und lässt sich auch nach weiteren Spielständen immer wieder herbeirufen. Ich habe mittlerweile eine Art Branchset meiner "Lieblingssituationen" im Game, die ich immer wieder durchspielen kann, wie es mir gerade gefällt. Und nein es sind nicht immer nur Explosionen :)

**Networking, Baby! (Remotes)**

Mein Minecraft Setup mit allen Einstellungen und Spielständen zentral an einem Ort zu haben war ehrlich gesagt meine erste Motivation git einzusetzen. Ich spiele Minecraft auf 3 verschiedenen Maschinen (Ubuntu, Debian und sogar Mac OSX) und wollte keine 3 unterschiedlichen Maps pflegen müssen. Deshalb fing ich an auf meinem Server zwischen zu lagern. Ein eigens laufender git-Server ist hier aber Vorraussetzung! Freies Hosting bei beispielsweise [Github](http://github.com) fällt wegen der großen Datenmengen (ca. 300MB bei mir derzeit) und der fehlenden Privatsphäre flach. Remote-Server hinterlegen und aktuellen Stand pushen:

```
$ git remote add origin git@gitserver.com:minecraft
$ git push origin master
```


Remote-Server auf anderen hosts klonen:

```
# Ubuntu/Debian
$ git clone git@gitserver.com:minecraft $HOME/.minecraft
# Mac OSX
$ git clone git@gitserver.com:minecraft $HOME/Library/Application Support/minecraft
```


**At least**

Ich möchte nicht sagen, dass dies hier der ultimative Weg zum heiligen Gral in Minecraft ist. Manchmal weckt eben diese erzeugte "Sicherheit" durch den Reset eine gewisse "Wayne..." Einstellung in einem, die dem Spielspaß ein Kleinwenig den kitzel raubt. Gerade am Anfang hat es mir aber extrem geholfen, nicht bei jedem Wipe alle Items zu verlieren oder sich nach einem Ausflug in den Wald wieder "zurück warpen" zu können.

Auf weitere Ideen im Umgang mit Minecraft und Git freue ich mich natürlich wie immer :)
