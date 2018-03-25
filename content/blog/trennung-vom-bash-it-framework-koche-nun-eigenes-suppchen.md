---
aliases:
- /blog/2012/08/12/trennung-vom-bash-it-framework-koche-nun-eigenes-s%C3%BCppchen/
comments:
- author: Martin
  content: "<p>5 Sekunden zum Prompt ist ja *noch* langsamer als oh-my-zsh mit ~2
    Sekunden + Bedenkzeit je Tab-Anschlag (unter FDE).</p><p>Gab es jemals die Intention,
    alle plugins lazy zu machen? [1] finde ich ziemlich geil. Gerade noch vertretbare
    0.45 Sekunden f\xFCr die Bash ;)</p><p>[1]: <a href=\"http://fahdshariff.blogspot.de/2011/09/speeding-up-bash-profile-load-time.html\"
    rel=\"nofollow\">http://fahdshariff.blogspot.de/2011/09/speeding-up-bash-profile-load-time.html</a></p><p></p>"
  date: '2012-08-12T20:00:42'
- author: Andreas
  content: "<p>Ich kann dich verstehen, geht mir auch so. Mit SSD geht wenigstens
    der Prompt etwas schneller, aber bei gro\xDFen SVN oder GIT-Archiven dauert das
    wechseln in ein Verzeichnis unglaublich lange (&gt;10 sec)</p>"
  date: '2012-08-12T20:44:04'
- author: noqqe
  content: "<p>Ok, 5 Sekunden sind auch etwas \xFCbertrieben gewesen. Aber es f\xFChlt
    sich eben an wie eine kleine Ewigkeit :)</p><p>Dieses <a href=\"http://dnycomp.sh\"
    rel=\"nofollow\">dnycomp.sh</a> ist ja geiler Scheiss. Schau ich mir heut Abend
    mal an. Danke!</p>"
  date: '2012-08-13T07:55:52'
- author: Martin
  content: "<p>Leider hat der Entwickler sich inzwischen von der bash abgewandt und
    das Repository gel\xF6scht... und <a href=\"http://dyncomp.sh\" rel=\"nofollow\">dyncomp.sh</a>
    hat noch ein paar merkw\xFCrdige Bugs, die ich nicht so ganz zu fixen wei\xDF.
    So funktionieren nur ~80% aller completions, `rsync` oder `hub` geh\xF6ren nicht
    dazu.</p>"
  date: '2012-09-30T11:54:36'
- author: noqqe
  content: "<p>Ach, und ich dachte ich w\xE4re zu bl\xF6de :) </p>"
  date: '2012-09-30T16:37:33'
date: '2012-08-12T18:12:00'
tags:
- development
- nano-bash
- github
- oh-my-zsh
- nano
- bashrc
- zsh
- tricks
- framework
- bash-it
- config
- debian
- bash
title: "Trennung vom bash-it Framework - Koche nun eigenes S\xFCppchen"
---

Manchmal muss man das auch. [Einige](/blog/2010/11/24/bash-it-n0qorg-theme-und-git_info/)
[Zeit](/blog/2010/12/07/github-mitarbeit-an-bash-it/) hab ich
Hirnschmalz in das Community Framework [bash-it](https://github.com/revans/bash-it) gesteckt.

Anfangs mit der aktiven Weiterentwicklung beschäftigt, hab ich später eher
massiv Aufwand betreiben müssen um die neuen Dinge mit meiner Version
kompatibel zu halten.  Man konnte bash-it früher entweder komplett als eine
Art Distribution für die Bash benutzen, oder sich wie in einer Art
Supermarkt an einzelnen Schnipseln, Funktionen und kleinen Skripten
bedienen die man gut fand. Und das war super.  Man konnte Dinge lernen,
Schnippsel von anderen adaptieren und hatte immer Trickreiche Funktionen am
Start.

{{< figure src="/uploads/2012/08/frameworks.jpg" >}}

Der Rip-Off von [oh-my-zsh](https://github.com/robbyrussell/oh-my-zsh) hat sich
aber mit der Zeit (und Dank der großen Aktivität der Community) m.E. zu einem
für mich nicht mehr funktionierendem Konzept entwickelt. Konkret:

* Die Überarbeitung des Plugin Systems mit Meta Informationen war ein
  riesen Fehler
* Aktivierung und Deaktivierung einzelner Komponenten ist unnötig und
  verwirrend für neue User (und für mich auch.. :P )
* Etwas, dass ich gerne den "theme-color-vcs-prompt-complex" nenne. Aber
  das würde den Rahmen sprengen...
* Der Wust aus 1. Mio Dingen, den mittlerweile niemand (mehr)
  braucht/versteht
* Fehlende Ziele/grobe Richtung für die Entwicklung
* Durch die stark wachsende Komplexität funktioniert bash-it nichtmal mehr
  als Supermarkt Bedienkonzept.
* Ca. 5 Sekunden bis zum Prompt

All diese Probleme und Punkte haben mich dazu veranlasst, meine eigene
Suppe zu kochen, die ich in meiner gnadenlosen Unkreativität
[nano-bash](https://github.com/noqqe/nano-bash) genannt habe. Es hat einen
ganzen Nachmittag gedauert die wenigen Dinge die ich wirklich brauche aus
bash-it zu extrahieren und aus Funktionen zu lösen, damit das Ganze in eine
neue (für mich pflegbare) Form passt. Denke das spricht auch schon für
sich...

Schade :(
