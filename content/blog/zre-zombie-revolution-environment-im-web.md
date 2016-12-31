---
aliases:
- /archives/1500
- /blog/2011/03/06/zre-zombie-revolution-environment-im-web
comments:
- author: F00b4r
  content: "<p>Ein Gl\xFCck, das die bash unter OSX lauff\xE4hig ist. Man k\xF6nnte
    aber auch Solaris nehmen ... oder cygwin!</p>"
  date: '2011-03-07T03:39:41'
date: '2011-03-06T14:06:38'
tags:
- web
- r
- shell
- revolution
- brains
- commandline
- brainzz
- zombie
- zre
- html
- emulator
- zombies
- databases
- php
- debian
- bash
- statistik
title: ZRE | Zombie Revolution Environment. Im Web!
---

Demnächst könnte es unter Umständen zu einer Vielzahl von Posts kommen, die
sich um das [damals erwähnte](/archives/1314) [Zombie Revolution
Environment](https://github.com/noqqe/zombie-revolution-environment)
drehen. Um jetzt nicht unnötig weit auszuholen versuche ich es mal so kurz
wie mögilch zu beschreiben.

Vor ca. 2 Wochen bekamen wir @ School einen Flyer für den [JOY-Dortmund
Wettbewerb](http://www.joy-dortmund.de/de/home/) für Informatik Projekte
von IT Auszubildenden. Was unter anderem zu dem hier erwähnten Projekt
führte. [Holger](http://savier.n0q.org/) und ich kamen auf die Idee im
Grunde ZRE für das Web umzusetzen. Aber nicht nur um die Welt auf
unterhaltende Art und Weise anzeigen zu lassen, sondern viel mehr eine
Analyse des Spielverlaufs mit Hilfe von Statistiken zu erstellen.

{{< figure src="/uploads/2011/03/2251_16ea_4801.png" >}}

Da das "Spiel" ausschließlich aus zufällig generierten Umständen besteht,
wäre die Auswertung der Entwicklungen der simulierten Welt evtl.
interessant. Grade wenn man so auf Statistiken steht. :/


## Anforderungen

  * ZRE (in Bash geschrieben) als lauffähigen Daemon umsetzen der Output an eine Stelle generiert. (Bash - Shell)
  * Eine geeignete Website erstellen bzw. designen (HTML, CSS, PHP)
  * Eine Schnittstelle die die generierten "Welten" im Web anzeigt. Am besten als selbst aktualisierende live Anwendung. (JavaScript, AJAX)
  * Ein SQL-Modul für ZRE entwickeln, welches statistische Werte an eine Datenbank übermittelt. (MySQL)
  * PHP-Funktionen definieren, die Ergebnisse aus Datenbank abholen. (PHP)
  * Statistiken visualisieren (R statistical Programming Language)

Im wesentlichen sieht die Aufgabenteilung vor, das sich Holger um Website,
Design und PHP kümmert und ich mich um die Bash-Module, ZRE an sich, SQL
Datenbank und statistische Auswertung bemühe. Selbstverständlich würden wir
auch jegliche andere Hilfe oder freiwillige Mitarbeiter dafür begrüßen,
falls Interesse besteht. Das wir an dem Wettbewerb teilnehmen wird zusehens
unwahrscheinlicher, da das Projekt an sich schon eher einen
Unterhaltungswert, statt dem wirklichen Nutzen eines IT-Projekt hat.

Wer sich ZRE an sich mal ansehen will, kann einfach mal das Git-Repo
auschecken und starten (auch unter Mac OS X lauffähig).

``` bash
git clone git://github.com/noqqe/zombie-revolution-environment.git
cd zombie-revolution-environment
./zre.bash
```