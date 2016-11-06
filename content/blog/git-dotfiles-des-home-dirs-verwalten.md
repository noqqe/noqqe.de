---
date: 2010-08-16T18:42:48+02:00
comments: true
title: Git | Dotfiles des Home-Dirs verwalten
aliases:
- /blog/2010/08/16/git-dotfiles-des-home-dirs-verwalten
- /archives/1192
categories:
- Development
- Debian
- Ubuntu
- ubuntuusers
tags:
- xubuntu
- automatisierung
- dotfiles
- dotfiles verwalten
- git
- git hook
- gitrepo
- homedir
- hook
- post-commit
- post-update
- pre-commit
---

Permanent an verschiedenen Workstations zu sitzen hat Nachteile.
Heimrechner(Xubuntu), Arbeits-PC(Xubuntu), Macbook und im blödsten Fall
noch (private) Server. An all diesen Maschinen entwickelt man Vorlieben,
bestimmte Software zu bedienen, entsprechend anzupassen oder zu
konfigurieren. Es müssen nicht mal elementare Sachen sein. Selbst nicht
vorhandene triviale Kleinigkeiten wie zum Beispiel Anpassungen der
.dircolors oder .bashrc bzw. .bash_aliases sind manchmal extrem nervig.

> "Warum funktioniert mein Alias hier eigent.....ah. Falscher Host :/ "

Das brauche ich hier warscheinlich niemandem weiter erläutern ;)

Grundsätzlich bemühte ich mich einen kleinen Satz ausgewählter .dotfiles
meiner Home-Verzeichnisse auf neue Systeme zu kopieren bzw zu aktualiseren.
Aber bei einer Änderung >4 Hosts neu umkopieren? Oder Änderungen sogar per
Hand separat ausführen? Jedenfalls entwickelt mit der Zeit jeder, der auf
diese Problemstellung trifft, seinen eigenen Weg alle Dotfiles synchron zu
halten. Hier ist nun der (seit neuestem) Meinige.

Ich benutze ein Git-Repository. In diesem Repo befinden sich alle Dotfiles
die ich im alltäglichen Gebrauch benötige. Das sieht ungefähr so aus:

```
dotfiles/
|-- .bashrc
|-- coming-home.bash
|-- .csshrc
|-- .dircolors
|-- .dmrc
|-- .git/
|-- .gitconfig
-- .vimrc
```


(Aus Gründen der Lesbarkeit gekürzt)

Alle Änderungen, die für die Configs anstehen, kann ich nun hier tätigen.
Füge neue Aliase hinzu, ändere meine Editor einstellungen für Vim oder
ähnliches. Für die automatische Einrichtung habe ich
"[coming-home.bash](http://gist.github.com/527325)" geschrieben. Das Script
filtert automatisch alle Verzeichnisse(., .., .git) und Files (z.B. das
Script selbst) die nicht ins $HOME-Verzeichnis gehören und kopiert alles
was übrig bleibt in das Home-Dir des aktuellen Benutzers.

``` bash
#!/bin/bash
# get dotfiles
dotfiles=$(ls -la | grep -v ^d | awk '{print $8}'| grep -v ^coming-home.bash$ )
mode=${1:-normal}

for x in $dotfiles; do
    # fishing errors
    if [ ! -e $x ] ; then echo "error reading files" ; exit 1 ; fi
    # keep old .bashrc
    if [ $x = ".bashrc" ]; then
        if [ ! -e $HOME/.bashrc_old ]; then
            cp $HOME/.bashrc $HOME/.bashrc_old
        fi
    fi
    # forced hook
    if [ $mode = "--hook" ]; then echo "hook: copying $x" ; /bin/cp -r $x $HOME/ ; fi
    # standard run
    if [ $mode != "--hook" ]; then /bin/cp -i $x $HOME/ ; fi
done
```


Git bietet weiterhin noch schöne Möglichkeiten, Operationen bei Änderungen
automatisch ausführen zu lassen.
[Hooks](http://www.kernel.org/pub/software/scm/git/docs/githooks.html). Ein
post-commit-Hook führt das Skript jetzt nach jeder eingespeicherten
Änderung automatisch aus.

.git/hooks/post-commit:

    #!/bin/bash
    $GIT_DIR/../coming-home.bash --hook

Anwendungsbeispiel:

    $ git commit -a -m "added aliases and changed colors in vimrc"
    hook: copying .bashrc
    hook: copying .csshrc
    hook: copying .dircolors
    hook: copying .dmrc
    hook: copying .gitconfig
    hook: copying .vimrc

So kann ich jetzt Änderungen der Dotfiles im Git-Repo durchführen und habe
diese gleich darauf im Home-Verzeichnis. Dieses Repo existiert (über einen
Git-Server verwaltet) auf jedem meiner Hosts.  Um im Fallbeispiel zu
bleiben: Gehe ich morgen in die Arbeit, aktualisiere ich das (auch dort
vorhandene) Repo und führe coming-home.bash aus.

By the Way: An dieser Stelle habe ich auch über Automatisierung
nachgedacht. Git-Hooks gibt es für die verschiedensten Situationen:

> Pre-Commit (vor speichern einer Änderung)
> Post-Update (Nach dem pushen auf den Server)
> usw.

Für eine Situation darf es diese Automatisierung aber nicht geben. Nämlich
nach dem post-clone (unmittelbar nach dem Herunterladen eines Repos).
Automatische Skripte die nach dem Herunterladen ausgeführt werden könnten
große Schäden anrichten, wenn man so drüber nachdenkt.
