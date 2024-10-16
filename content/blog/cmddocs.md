---
comments:
- author: dAnjou
  content: "Du realisierst aber schon, dass du das ganze auch mit einer \xFCberschaubaren
    Anzahl Aliasen haben k\xF6nntest? Falls du das Tool `tree` nicht kennst, hab ich
    es dir soeben empfohlen. Und zusammen mit einer besseren Shell wie z.B. zsh sehe
    ich in diesem Projekt absolut keinen Mehrwert. Kl\xE4re mich auf, falls ich es
    noch nicht ganz verstanden habe."
  date: '2015-01-24T20:18:18.947328'
- author: Martin
  content: "Der Vorteil einer dedizierten \"Wiki-Shell\" ist z.B. das automatische
    comitten von \xC4nderungen, z.B. nach `edit Foobar` folgt sofort ein Commit. Nat\xFCrlich
    l\xE4sst sich das auch mit einem Alias realisieren. Auch das ANSI-Highlighting
    hast Du mit `cat` nicht direkt ootb."
  date: '2015-01-25T00:38:13.749454'
- author: noqqe
  content: Warum gibt es `htop`. Da gibts die selben Informationen wie in `top`
  date: '2015-01-25T13:38:14.669727'
- author: dAnjou
  content: "Du hast halt ziemlich viel Code geschrieben, den man durch simple Tools,
    die meistens schon vorinstalliert sind, ersetzen k\xF6nnte."
  date: '2015-01-30T19:13:15.198285'
- author: noqqe
  content: "Das da w\xE4re?"
  date: '2015-01-31T13:53:54.871927'
- author: bo
  content: Ja fett :)
  date: '2015-02-13T14:02:25.353568'
date: '2015-01-24T08:23:00'
tags:
- development
- wiki
- git
- openbsd
- python
- docs
- gitit
- haskell
- markdown
title: cmddocs
---

"Ich habe mein eigenes Wiki programmiert. So einfach ist das."

{{< figure src="/uploads/2015/01/disturbing.jpg" >}}

In letzter Zeit hab ich damit den ein oder anderen seltsamen Blick
kassiert. Tatsächlich ist aber alles etwas anders.

### gitit

Ich war lange ein zufriedener User von [gitit](https://gitit.net). Zugegeben,
nur wenn es dann wirklich mal kompiliert war. Ich wollte nicht zum 3. mal nach
neuem OpenBSD Release 4-5 Stunden verschwenden, das Stück Haskell gebaut zu
bekommen. Außerdem war ich das Webinterface irgendwie Leid.

### Übergang

Also was will ich überhaupt? Ich will auf CLI `markdown` schreiben und die Files mit `git`
unter Versionskontrolle halten. Möglichst einfach durchsuchen und bearbeiten
können. Ich dachte okay, `gitit` speichert Markdown Files. Ich kopier einfach
alles weg, lege mir in einem Verzeichnis auf meiner OpenBSD CLI Maschine an, in
dem ich alles finde was ich brauche. Das passt super ins Konzept mit
[mutt](http://www.mutt.org), [taskwarrior](http://taskwarrior.org),
[jrnl](http://maebert.github.io/jrnl/), [weechat](http://weechat.org) und was
ich sonst so verwende.

Auf längere Sicht wurde es aber ziemlich Tippintensiv. Jedes mal alle `git add`,
`git commit`, `grep`, `tree` Commands nutzen. Es hat sich irgendwie müßig
angefühlt. Hab nicht mehr gerne damit gearbeitet, kam zu dem Schluss das ich
auf Commandline bleiben, aber mir das dokumentieren/notieren etwas erleichtern
will.

### cmddocs

Zuerst machte ich mich auf die suche nach CLI-Wikis, aber die da sowas so
gut wie garnicht zu existieren scheint, fing ich an mir ein kleines Commandline
Interface mit `Python` zu bauen. Dazu benutzt hab ich das `Cli` Modul, was so
gefühlt der heilige Gral der CLI Module ist. Tabcompletion, Helps, History,
CTRL-R (unendlich wichtig), ESC-., ist alles schon drin.

Die Codebase von [cmddocs](https://github.com/noqqe/cmddocs) war nicht besonders
schön, aber tat das was ich wollte. Ich war erstmal zufrieden. Dann kam #31c3,
ich traf [posativ](https://posativ.org) wieder, der mich erstmal darauf Hinwies,
wie hässlich eigentlich. Nach seinen Tipps fing ich dann an das ganze etwas
umzustrukturieren. Eigentlich wollte ich das gute Stück nie irgendwo publishen,
aber durch die Änderungen sieht der Code jetzt doch einigermaßen annehmbar aus.
Ein Github Repo+Readme angelegt und gepushed.

Demo hier: [asciinema cmddocs demo](https://asciinema.org/a/15168) (etwas outdated)

Es fehlen noch ein paar Features, der View-Mode bzw. der aus 6 Zeilen bestehende
Markdown-&gt;ANSI Converter hat noch ein paar Problemchen, aber alles in Allem
bin ich wirklich zufrieden was daraus geworden ist.
