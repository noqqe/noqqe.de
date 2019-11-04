---
aliases:
- /blog/2010/08/16/git-dotfiles-des-home-dirs-verwalten
- /archives/1192
comments:
- author: Vain
  content: "<p>Moin,</p><p>so als Idee: Wahrscheinlich k\xF6nntest du dir das \"coming-home.bash\"-Skript
    sparen, wenn du stattdessen Symlinks setzt. Also statt einer echten ~/.bashrc
    einfach einen Symlink nach ~/git/dotfiles/.bashrc. Habe zumindest keine schlechten
    Erfahrungen damit gemacht. :)</p><p>Cheers!</p>"
  date: '2010-08-16T23:07:54'
- author: noqqe
  content: "<p>@Vain: Moin, mit der \xDCberlegung hatte ich wirklich schonmal gespielt.
    Erschien mir aber irgendwie zu \"unstable\". Son lokales Repo kommt dann doch
    mal weg, oder sonst was... :/ <br>Hab mich da ehrlich gesagt nicht dran gewagt.
    </p><p>In wie weit hast du das im Einsatz? F\xFCr mehrere Configs? Bzw auch \xFCber
    fremde Partitionen ? Gibts da eventuell Leistungseinbu\xDFen?</p>"
  date: '2010-08-16T23:12:10'
- author: Keba
  content: "<p>Nat\xFCrlich kann ein lokales Repo mal wegkommen, aber bei DVCS wie
    git gibt es zum Gl\xFCck kein \"Single Point of failure\". Einfach neu klonen
    und alles ist gut. ;)</p>"
  date: '2010-08-17T01:35:38'
- author: Vain
  content: "<p>Naja, alle meine Dotfiles sind Symlinks in ein Git-Repo. Allerdings
    nicht \xFCber Partitionsgrenzen hinweg, nein. \xDCber ein Verlorengehen des Repos
    hab ich eigentlich auch noch nie nachgedacht \u2013 eben, weil man\u2019s halt
    wieder klonen kann. :)</p><p>Mir f\xE4llt aber gerade ein, dass es doch manchmal
    zu Problemen kommen kann. Manche Programme schreiben ja ihre Configs neu, nachdem
    man sie beendet hat (auch, wenn sich eigentlich nichts ge\xE4ndert hat). Und manchmal
    machen sie das so d\xE4mlich, dass sie die Datei erst komplett l\xF6schen und
    dann neu anlegen \u2013 dabei geht der Symlink halt kaputt. Dadurch entsteht kein
    Datenverlust, nur Git bekommt keine \xC4nderungen mehr mit.</p><p>Ich glaube,
    WeeChat war so ein Sonderfall, was das so gemacht hat. Kann man dann aber auch
    leicht umschiffen, wenn man das Verzeichnis (in dem Fall ~/.weechat/) verlinkt
    statt einzelner Dateien. Man muss es nur erst mal bemerken. ;)</p><p>Aber das
    ist so selten, dass ich es schon erfolgreich aus dem Ged\xE4chtnis verdr\xE4ngt
    hatte.</p><p>Cheers!</p>"
  date: '2010-08-17T02:15:29'
- author: noqqe
  content: "<p>@Keba: Jepp. Denke mal es wird Distributed Version Control System bedeuten
    oder? :) Wenn das Repo mal weg ist und der Rechner gebootet wird, werden warscheinlich
    meine ganzen Links auch weg sein oder? Standard-Config wieder eingesetzt evtl?
    </p><p>@Vain: Okay, also mit einschr\xE4nkungen h\xF6rt sich das aber doch trotzdem
    ganz nett an :) Ich denke ich werds mal testen und sehen wie ich damit zurecht
    komme. Danke f\xFCr den Tipp :)</p>"
  date: '2010-08-17T11:12:02'
- author: nbkr
  content: "<p>F\xFCr sowas nutze ich gerne \"Puppet\". Meine User lege ich dort alle
    an und bestimme \xFCber das Manifest welcher User auf welchem System zu finden
    sein soll.</p><p>Au\xDFerdem speichere ich mir noch welche Dateien im Homedir
    des Users liegen sollen und wie die aussehen.</p><p>Muss ich dann an diesen Dateien
    etwas anpassen (z.B. die .vimrc) dann mache ich das auf der Puppetmaster-Maschine
    und die anderen System holen sich die Updates automatisch und ohne zu tun.</p>"
  date: '2010-08-17T14:16:38'
- author: bka
  content: <p>hi,</p><p>da kommt aber mehr als nur die dotfiles raus...</p><p>ls -la
    | grep -v ^d | awk '{print $8}'</p><p>alle dotfiles bekommt man mit:<br>ls -a
    ~/ | grep '^.'</p>
  date: '2010-08-18T20:32:51'
- author: noqqe
  content: <p>Hallo @bka, <br>die Ausgabe soll den Inhalt des Git-Repos filtern. Nicht
    die des Home-Verzeichnisses.</p>
  date: '2010-08-18T21:58:06'
- author: marian
  content: <p>&gt; ls -la | grep -v ^d | awk '{print $8}'| grep -v ^coming-home.bash$<br>find
    existiert :)</p>
  date: '2010-08-27T06:25:44'
date: '2010-08-16T16:42:48'
tags:
- development
- dotfiles verwalten
- git
- homedir
- post-update
- dotfiles
- git hook
- hook
- automatisierung
- gitrepo
- ubuntu
- pre-commit
- debian
- xubuntu
- post-commit
title: Git | Dotfiles des Home-Dirs verwalten
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
