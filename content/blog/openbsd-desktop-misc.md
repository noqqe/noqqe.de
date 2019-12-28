---
title: "OpenBSD Desktop - Die Software"
date: 2019-12-28T11:57:48+01:00
tags:
- cwm
- OpenBSD
- feh
- surf
- x11
- suckless
---

Es ist Tag 2 und ich bin ziemlich zufrieden mit dem Arbeiten am groestenteils
textbasierten OpenBSD Desktop.

{{< figure src="/uploads/2019/12/linux.jpeg" >}}

An meinem `cwm` Window Manager habe ich ehrlich gesagt wenig bis garnichts
herumkonfiguriert. Ich will lediglich ein schwarzes Terminal haben, weshalb
ich den "Neues Terminal"-Shortcut ueberschreiben und einen weiteren um meinen
Bildschrim zu sperren.

```
> cat .cwmrc
bind-key CM-Return "xterm -bg black -fg white +sb"
bind-key CM-q "slock"
```

Wo wir gerade ueber `xterm` sprechen. Das Terminal bietet mir alles was ich
brauche, ist schnell, klein, miniamlistisch. Es besteht garkein Bedarf mich
aus den Ports an einer alternative Bedienen zu muessen. Die unendlich lange
Manpage dazu musste ich bisher nicht lesen, da ich einfach keine
Anforderungen habe. Der Font ist schoen, weiss und der Hintergrund schwarz.

Browsen. Wie `slock` von [suckless.org](https://suckless.org), benutze ich
`surf`. Der Browser bringt zwar eine Unmenge von Abhaengigkeiten via
`pkg_add` mit, ist aber trotzdem noch das kleinste Stueck Software die man in
Webbereich so haben kann. Die Standard Keybindings sind gut und Navigieren
und Suchen funktionert mittels `dmenu`.

{{< figure src="/uploads/2019/12/surf.png" >}}

Was bei `surf` nicht geht ist Adblocking. Hierzu habe ich mittels `pf`
([pf-badhost](https://www.geoghegan.ca/pfbadhost.html)) und `unbound`
([unbound-adblock](https://www.geoghegan.ca/unbound-adblock.html)) eine
Blacklist aus dem Internet gezogen, die die meisten Werbenetzwerke
herausfiltert.

Fuer die "normale" Arbeit, also Text Editieren und die Shell bin ich von dem
integrierten `vi(1)` und `ksh(1)` aus Gewohnheit und Komfort zu `fish` und
`neovim` gewechselt. Ich war fast etwas erschrocken wie sehr ich mich
mittlerweile an die neue Shell gewoehnt habe. Die Features die ich moechte
und brauche bezahle ich mit etwas aufblasenerer Software als es eigentlich
sein muesste. An der Stelle finde ich das aber in Ordnung.

Todos & Notizen. Anfangs habe ich alles was ich so an Baustellen im Kopf
hatte in eine todo.txt Datei geschrieben. Einfach so mit `vi`. Das hat gut
funktioniert, normalerweise benutze ich ja [Todoist](https://todoist.com)
aber extra mit einer Webseite herummachen war mir auf dem Desktop zu bloed.
Ich habe vor Todoist 5 Jahre lang [taskwarrior](https://taskwarrior.org)
benutzt, welchen ich mir nun auch mal wieder aus dem Package System
installiert hab. Es passt ja in das textbasierte Konzept und die Befehle sind
immernoch in meinem Muskelgedaechtnis. Meine todo.txt Datei habe ich dann in
taskwarrior umgezogen.

Home Directory. [blinry](https://morr.cc) hat vor ein paar Jahren einen
Blogpost ueber ein wunderschoenes Konzept zur Verwaltung seines
Heimatverzeichnisses geschrieben. Siehe [$HOME, sweet
$HOME](https://morr.cc/home-sweet-home/). Dieses Konzept habe ich mir zueigen
gemacht und mein `$HOME` dementsprechend eingerichtet.

```
> find ~ -type d -maxdepth 2
library
library/xresources
library/obsd
tmp
work
work/archive
work/permanent
work/permanent/entbehrlich.es
work/permanent/noqqe.de
```

Neben dem Browser, ist die andere Grafische Anwendung `feh`. Um Bilder
anzuschauen und sogar zu bearbeiten.

```
feh . # alle bilder im folder anschauen
feh --edit linux.png # Bild mit < rotieren! :)
```

Wie gesagt, ganz viele neue Software zum ausprobieren.
Es ist eine andere Erfahrung. Ein anderes Gefuehl als das MacBook mit dem ich
ueblicherweise durch den Tag komme. Die Probleme sind anders und die
Loesungen dazu auch. Das macht ziemlich viel Spass und bringt einen auch
weiter.

Was ich immernoch nicht habe sind Umlaute. UTF-8 Detox sozusagen.
