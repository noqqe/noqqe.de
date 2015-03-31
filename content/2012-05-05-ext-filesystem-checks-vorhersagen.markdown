---
layout: post
title: "Ext Filesystem Checks vorhersagen"
date: 2012-05-05 11:33
comments: true
categories:
- ubuntuusers
- web
- hardware
- linux
- debian
- ubuntu
keywords: "Debian, kernel, fsck, ext2, ext3, ext4, Superblock, dumpe2fs, restore, check, filesystem, fs"
---

Kernel Updates am Server seiner Wahl einspielen und dann...warten. Warten auf die erlösenden
ICMP Antworten. Oder bootet die Kiste doch nicht? Alles dauert verdächtig lange. Auch nach 10 Minuten noch nichts.
Gerade wenn man die Zugangsdaten fürs DRAC/ILO rausgekramt hat zeigt sich: Filesystem Check.

Extended Filesystems machen an den folgenden 2 Punkten fest wann fsck's passieren:

* Maximale Anzahl von Mounts
* Zeitliche Abstände zwischen den Routine Checks

Die Informationen stehen im Superblock einer Parition. Das Tool dumpe2fs stellt diese zur Verfügung:

    $ dumpe2fs -h /dev/sda1

Das ganze hab ich dann verskriptet. Es "scanned" alle gemounteten ext Parititonen und warnt einen, falls die
maximalen Mounts oder die zyklischen Checks anstehen. [ext-verify.sh](https://gist.github.com/noqqe/2601222)

Das sieht dann unter Umständen so aus:

{% codeblock %}
WARNING: Max mount count on /dev/sda1 has been reached. (29/29)
WARNING: /dev/sda1 has reached the next periodically filesystemcheck. (Sa Apr 7 13:37:58 2012)
RESULT: A fsck will be executed at the next reboot for /dev/sda1.
{% endcodeblock %}

Natürlich für jede gefundene Partition. Um das jetzt noch am besten irgendwie zu
automatisieren hab ich mich entschieden das in `apt` zu integrieren.

{% codeblock %}
$ vim /etc/apt/apt.conf.d/09extverfiy
DPkg::Pre-Install-Pkgs   { "if [ -x /usr/local/bin/ext-verify.sh ]; then echo 'Verifying ext Filesystems' ; /usr/local/bin/ext-verify.sh ; fi"; };
{% endcodeblock %}

Bei jedem apt-get/aptitude wird das nun ausgeführt.

{% codeblock %}
$ aptitude install whois
Die folgenden NEUEN Pakete werden zusätzlich installiert:
  whois
0 Pakete aktualisiert, 1 zusätzlich installiert, 0 werden entfernt und 0 nicht aktualisiert.
Muss 0 B/64,9 kB an Archiven herunterladen. Nach dem Entpacken werden 406 kB zusätzlich belegt sein.
Verifying ext Filesystems
RESULT: Everything's fine on /dev/sda1.
Vormals abgewähltes Paket whois wird gewählt.
(Lese Datenbank ... 48680 Dateien und Verzeichnisse sind derzeit installiert.)
Entpacken von whois (aus .../whois_5.0.10_amd64.deb) ...
Trigger für man-db werden verarbeitet ...
whois (5.0.10) wird eingerichtet ...
{% endcodeblock %}


