---
aliases:
- /archives/1425
- /blog/2010/12/14/bash-backup-eines-files-mit-zeitstempel-einfach-gemacht
comments:
- author: FreakErn
  content: "<p>Hey Noqqe,</p><p>coole Sache, vielen Dank!</p><p>Ich w\xFCrde diese
    Funktion allerdings in eine leere Textdatei packen mit ner entsprechenden shebang-zeile.</p><p>Der
    Vorteil des ganzen ist, dass du dir sch\xF6n einen kompletten Ordner anlegen kannst,
    in dem alle deine Skripte drin sind welchen du wiederum auf einen SVN-Server packen
    kannst.</p><p>Da ich zwei Komputa besitze und regelm\xE4\xDFig an beiden arbeite,
    hab ich so, auf dem zweiten PC die M\xF6glichkeit mit:<br>svn co svn://myServer/bin/trunk
    ~/bin &amp;&amp; source ~/.profile</p><p>direkten zugriff auf meine skripte!</p><p>Da
    in der ~/.profile eine \xDCberpr\xFCfung auf das Verzeichnis ~/bin existiert und
    diese, wenn vorhanden, in die $PATH eingetragen wird, brauch man auch sonst, unter
    Ubuntu, nix weiter machen.</p><p>Greetz</p>"
  date: '2010-12-14T19:57:05'
- author: noqqe
  content: "<p>Hi Freakern, </p><p>Da hast du Recht :) Die Meisten haben so Ihre eigene
    Verwaltung von hilfreichen Schnippseln. </p><p>Ich pers\xF6hnlich habe das auch
    lange so gel\xF6st wie du, bin dann aber zu bash-it gewechselt. </p><p>Wenn du
    Interesse hast schau dir mal <a href=\"http://github.com/revans/bash-it\" rel=\"nofollow\">http://github.com/revans/bash-...</a>
    an. Das ist im Grunde das selbe in Gr\xFCn und Strukturiert :)</p>"
  date: '2010-12-14T20:13:36'
- author: Lars Moelleken
  content: <p>... hab mir auch mal ein paar Shell-Schnippseln zusammengeschrieben
    :-) <br>-&gt; <a href="http://suckup.de/blog/2010/07/30/bashrc/" rel="nofollow">http://suckup.de/blog/2010/07/...</a></p><p>bash-it
    ist ne gute Idee! Thx!</p>
  date: '2010-12-14T22:49:36'
- author: Alex
  content: "<p>Quoting! Sonst erlebst du unangenehme \xDCberaschungen.</p>"
  date: '2010-12-15T01:02:37'
- author: "Gr\xFCni"
  content: "<p>Die kleinen Bash-Einzeiler in einer Datei zu vereinen halte ich f\xFCr
    eine ausgezeichnete Idee.</p><p>Ich habe das folgenderma\xDFen gel\xF6st:</p><p>
    -) Eintrag in der .bashrc:<br>\"if [ -f ~/.bash_snippets ]; then<br>        .
    ~/.bash_snippets<br>fi\"</p><p> -) und die Bash-Einzeiler in der Datei .bash_snippets
    aufgelistet.</p>"
  date: '2010-12-15T03:24:07'
- author: noqqe
  content: "<p>@Alex: Quoting ist erf\xFCllt ;)<br>@Gr\xFCni: Auch gute M\xF6glichkeit
    :) Aber wenns umfangreicher wird wird es auch un\xFCbersichtlich. Das war zumindest
    meine Erfahrung damit :)</p>"
  date: '2010-12-15T16:29:38'
- author: Tim
  content: <p>Super! Habe ich gleich in meine dotfiles eingebaut.</p>
  date: '2011-01-18T23:02:29'
date: '2010-12-14T14:33:45'
tags:
- development
- function
- shell
- buf
- commandline
- commandlinefu
- kommandozeile
- bashrc
- admins
- dirty
- schnell
- ubuntu
- quick
- backup
- debian
title: 'Bash | Backup eines Files mit Zeitstempel in unkompliziert. '
---

[Auf eine besondere kleine Funktion](http://www.commandlinefu.com/commands/view/7292/backup-a-file-with-a-date-time-stamp)
möchte ich kurz aufmerksam machen, weil sie viele Administratoren oder
Sonst-Wie-Konfigurationen-vornehmende-Benutzer wahrscheinlich auch brauchen
könnten:

``` bash
buf () { filename=$1; filetime=$(date +%Y%m%d_%H%M%S); cp ${filename} ${filename}_${filetime}; }
```

Man spielt an der main.cf des Postfix herum oder bastelt an der xorg.conf.
Und wie man das eben von dem erfahrenen Linux-Benutzer seines Vertrauens
(oder dem UbuntuUsersWiki) gelernt hat, sichert man vorher die
Konfigurationsdatei mit dem obligatorischen:

```
cp xorg.conf xorg.conf.bak
```

Je nach Intensität der Anpassungen häufen sich auch Backupfiles a la
"xorg.conf.bak.bak" oder ".bak2 bis .bak5000"

Fügt man die obige Funktion in seine .bashrc hinzu, erledigt sich das Thema
relativ schnell.

```
buf xorg.conf
```

Erstellt automatisch ein File mit dem aktuellen Zeitstempel im aktuellen
Verzeichnis. Nochmal im formatierten Zustand:

``` bash
buf () {
filename=$1
filetime=$(date +%Y%m%d_%H%M%S)
cp ${filename} ${filename}_${filetime}
}
```

Dass [commandlinefu.com](http://commandlinefu.com) eine großartige Site
ist, ist wahrscheinlich für die meisten nichts Neues. Teilweise regnet es
dort wunderschöne kleine Kommandozeilen für den Alltag.