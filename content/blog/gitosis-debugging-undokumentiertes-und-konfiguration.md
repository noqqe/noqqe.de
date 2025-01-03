---
aliases:
- /blog/2010/08/12/gitosis-debugging-undokumentiertes-und-konfiguration
- /archives/1175
comments:
- author: surtic
  content: "<p>Moin,</p><p>Du bringst es echt auf den Punkt, als ich vor etwa 1 Jahr
    f\xFCr mich Gitosis Konfiguriert und dann entlich mal zum laufen gebracht habe
    verging so seine Zeit. Vor allem das mit dem Hook post-update war keine einfache
    Sache.</p><p>Wie w\xE4re es mit einen Wiki Eintrag bei <a href=\"http://ubuntuusers.de\"
    rel=\"nofollow\">ubuntuusers.de</a> zu erstellen?</p><p>greez</p>"
  date: '2010-08-13T10:10:36'
- author: noqqe
  content: "<p>Hi surtic: Danke, sch\xF6n dass es bei sowas nicht nur mir so geht.
    <br>Ob bei ubuntuusers Bedarf f\xFCr so einen Eintrag besteht weiss ich nicht.
    Vielleicht erstelle ich einen :)</p>"
  date: '2010-08-13T10:41:21'
- author: philipp
  content: "<p>Es gibt BTW eine example.conf</p><p><a href=\"http://eagain.net/gitweb/?p=gitosis.git;a=blob;f=example.conf;h=87bd822fc17ebfe0132f7060afa8036f84c3be99;hb=HEAD\"
    rel=\"nofollow\">http://eagain.net/gitweb/?p=gitosis.git;a=blob;f=example.conf;h=87bd822fc17ebfe0132f7060afa8036f84c3be99;hb=HEAD</a></p><p>Die ist interessant,
    da steht z.B. drinne, wie gruppen mitglied einer anderen gruppe werden k\xF6nnen!...</p><p>(subdirs
    behandeln sie nur implizit, weil eines mit subdir angegeben wird)</p>"
  date: '2011-01-20T19:49:43'
- author: noqqe
  content: "<p>Hallo Phillip,</p><p>Danke :) Ich hab mittlerweile aber nicht mehr
    wirklich Bedarf f\xFCr gitosis. </p><p>Bin auf gitolite umgestiegen, da gitosis
    an sich ja auch nicht mehr weiter entwickelt wird. :(</p>"
  date: '2011-01-20T20:49:31'
date: '2010-08-12T17:46:44'
tags:
- gitserver
- how to
- hack
- gitdaemon
- subdirs
- gitosis hacking
- hook
- post-update
- accounting
- git
- gitosis
- pubkey
- development
- subdirectories
- gitweb
- format
- syntax
- keydir
- debian
- loglevel
- setup
- documentation
- unterordner
- public key
- debug
title: Gitosis | Debugging, Undokumentiertes und Konfiguration
---

Aktuell setze ich mich mit
[Gitosis](http://eagain.net/gitweb/?p=gitosis.git;a=summary) auseinander. Ja,
auseinander setzen ist gut ausgedrückt. Dieses  widerspenstige,
(standardmäßig) wenig gesprächige und nur [oberflächlich dokumentierte](http://eagain.net/gitweb/?p=gitosis.git;a=blob;f=README.rst;h=92047762c38cdf018a901b48a5a092796f51500e;hb=dedb3dc63f413ed6eeba8082b7e93ad136b16d0d)
Stück Software sträubt sich vehement gegen den  tieferen Einsatz und
komplexeren Ordnerhierarchien.

Also im Klartext: Ich finde Gitosis super. Der Ansatz ist gut. Definierbare
Rechte und Gruppenorganisation für Git-Repositories. Im Web findet man
[unzählige](http://scie.nti.st/2007/11/14/hosting-git-repositories-the-easy-and-secure-way)
[How](http://bogdan.org.ua/2009/02/20/gitosis-how-to-add-new-repository.html)-[To's](http://www.mantisbt.org/wiki/doku.php/mantisbt:gitosis_management)
die eine Standardinstallation wirklich gut und übersichtlich  dokumentieren
bzw. dazu anleiten. Speziellere Anpassungen und kleinere Grauzonen hingegen
leider gar nicht.

Um den Überblick über kleinere Mängel zu behalten und auch zu
dokumentieren:

### (1) post-update - Hook-Problematik

Die Funktionsweise von Gitosis ist eigentlich denkbar simpel.
Konfigurationsdatei syntaxgerecht anpassen  und innerhalb des admin-repos
committen+pushen. Ist die neue Config  gepusht und der Benutzer bzw. das
neue Repo nicht ansprechbar, verbringt man lange Zeit damit, Configs und
Pubkeys (siehe 2) zu  kontrollieren und erneut zu initialisieren. An dieser
Stelle begegnet einem  schon das erste Problem, welches Kenntnis über die
Funktionsweise von  Gitosis voraussetzt und hervorragend schlecht
dokumentiert ist.

```
ERROR:gitosis.serve.main:Repository read access denied fatal: The remote end hung up unexpectedly
```

Nach dem  das Master-Repo die gepushte Version von Gitosis erhält, führt es
einen  sogenannten
[Hook](http://www.kernel.org/pub/software/scm/git/docs/githooks.html) aus.
Dieser Hook ist ein Skript (bzw. ein Symlink dazu) welches die  neue
Konfiguration der Authentifizierungsstelle von Gitosis einließt. Der erste
Fehler der (wirklich häufig) passiert ist, dass dieses Skript schlicht weg
einfach nicht ausführbar  ist. Das lässt sich natürlich sehr einfach durch

``` bash
chmod 755  /home/git/repositories/gitosis-admin.git/hooks/post-update
```

lösen. Aber der knackende Punkt stellt (wie ich finde) die fehlende Meldung
dieses Fehlers dar. Gitosis teilt einem einfach nicht mit,  dass der Hook
fehlschlug und die getätigten Änderungen komplett für  die Katz waren.

### (2) Pub-Key-Format

Ein weiteres undefiniertes Loch der Konfiguration ist das Format, in  dem
der PublicKeys im conf-File angegeben werden muss.

Abgelegte Keys in gitosis-admin.git/keydir/ **müssen** mit .pub  enden.
Sonst werden diese nicht als Keys erkannt. Das ist das  kleinere Übel. Die,
für mich etwas unklare, Dokumentation darüber findet sich  in den
verschiedensten Varianten. Wie soll der PublicKey in der  gitosis.conf
hinterlegt werden?

"Wer wird Millionär"-mäßig kann ich jetzt nach dem
Trial-and-Error-Verfahren auflösen. Die endlosen Variationen von  Filenamen
und Config-Aufruf, die ich testen musste, damit ich mich  anmelden durfte,
haben sich also gelohnt.

( ) Name des Key-Files user@host.pub
( ) Im File hinterlegtes Suffix user@host
(X) FQDN user@host.domain.com
( ) Nur User zB. jdoe wie dokumentiert

### (3) Gitosis, sprich mit mir.

Ein weniger behütetes Geheimnis, ist die Gesprächigkeit von Gitosis. Im
Konfigurationsfile lässt sich das LogLevel deklarieren.

``` ini
[gitosis]
loglevel = DEBUG
gitweb = no
git-daemon = no
```

Siehe da, Informationen!

    DEBUG:gitosis.serve.main:Got command "git-receive-pack 'repo1'"
    DEBUG:gitosis.access.haveAccess:Access check for 'user@domain.com' as 'writable' on 'repo1'...
    DEBUG:gitosis.group.getMembership:found 'user@domain.com' in 'domain.com'
    DEBUG:gitosis.access.haveAccess:Access ok for 'user@domain.com' as 'writable' on 'repo1'
    DEBUG:gitosis.access.haveAccess:Using prefix 'repositories' for 'repo1'
    Initialized empty Git repository in /home/git/repositories/repo1.git/
    DEBUG:gitosis.gitdaemon:Global default is 'deny'
    DEBUG:gitosis.gitdaemon:Walking '.', seeing ['repo1', 'repo2', 'repo3', 'gitosis-admin.git']
    DEBUG:gitosis.gitdaemon:Deny 'gitosis-admin'
    DEBUG:gitosis.gitdaemon:Walking 'repo1', seeing ['justatest.git']
    DEBUG:gitosis.gitdaemon:Deny 'repo1/justatest'
    DEBUG:gitosis.serve.main:Serving git-receive-pack 'repositories/repo1.git'

Diese Infos werden nun zu fast jeder Gelegenheit ausgegeben. Pushen lokal,
Authentifizieren per Remote usw. Für weitere Administration mit Gitosis
unabdingbar.

### (4) Subdirectories handhaben

Angenommen ich habe oder möchte eine Ordnerstruktur meiner Projekte die
nicht alle in _repositories/_ liegen. Dieser kleine total untriviale
Umstand, lässt sich einfach nirgends nachlesen. Um das Ganze zu
verdeutlichen, habe ich kurz eine kleine Umgebung angelegt und mit _tree -L
2_ ausgegeben:

```
-- repositories
        |-- subdir1
        |   |-- repo1.git
        |   |-- repo2.git
        |   `-- repo3.git
        |-- subdir2
        |   `-- test.git
        -- subdir3
```

Angelege und initalisierte Repos werden zwar in der DEBUG-Ausgabe von
Gitosis wargenommen (siehe 3.), aber einfach übergangen. Stattdessen wird
ein neues Repo in _repositories/_ angelegt. Zumindest, wenn man der
Konfiguration folgt, wie sie in 95% der Fällen im Netz zu finden ist. Als
Faustregel für Unterordner gilt also: Zwingend jeden(!) Pfad mit Angabe des
Subdirs angeben.  Remote add:

``` bash
git remote add origin git@gitserver.org:subdir1/repo2.git
```

gitosis.conf:

``` init
[group subrepo]
members = user@host.com
writable = subdir1/repo2.git
```

Es klingt logisch. Aber ohne Dokumentation, ist es hart herauszubekommen.

Letztlich möchte ich kurz noch erwähnen, dass dieser Post keinerlei
Vorwurf, Flame oder Sh!t-Storm gegen Gitosis darstellen soll. Gitosis ist
ein wunderbarer Ansatz einer Benutzerverwaltung für Git-Remotes. Allerdings
abenteuerlich bei nicht alltäglicher Nutzung. In diesem Sinne. Happy
committing.
