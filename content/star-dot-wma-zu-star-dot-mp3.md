---
layout: post
title: "*.wma zu *.mp3"
date: 2012-07-29T17:39:00+02:00
comments: true
categories:
- planetenblogger
- Web
- Debian
- Bash
keyworkds: "lame, mp3, wma, convert, konvertieren, konvertierung, linux, cmd,
komplette Verzeichnisse"
---

Man denkt ja immer "Gott, wie viele von diesen $MUSIKFORMAT_A zu $MUSIKFORMAT_B
Formatierungsblogposts kann mein Reader eigentlich vertragen".
Das denkt man zumindest so lange, bis man mal in die Verlegenheit gerät sowas zu
brauchen.

{% img center /uploads/2012/07/koala.jpg %}

Tatsächlich ist das alles andere als einfach. In 90% der Szenarien werden 0815
Probleme gelöst wie:

> Ich möchte gerne ein Lied konvertieren. Wie geht das?

Dankeschön. Eins. Ich muss hier 25% meiner f*cking Musikbibliothek converten. Warum
finde ich daüber keine Blogposts?

Naja. Dann eben selbst. Der Plan ist einfach. Verzeichnisse herausfinden, die
*.wma Dateien enthalten und in eine Liste schreiben.

``` bash 
OIFS=$IFS
IFS=$'\n'
for i in $(find . -type f -iname '*.wma'); do
  dirname $i
done | sort | uniq > list.txt
```

Diese Liste mit Verzeichnissen abarbeiten, sodass ich die Lösung
die ich im [UbuntuUsers-Wiki](http://wiki.ubuntuusers.de/Audiodateien_umwandeln#WMA)
gefunden habe einbetten kann:

``` bash 
for x in $(cat list.txt) ; do
  cd $x
  for i in *.wma ; do
    mplayer -vo null -vc  dummy -af resample=44100 -ao pcm:waveheader "$i"
    lame --preset standard audiodump.wav -o "`basename "$i" .wma`.mp3"
    rm -f $i
  done
  rm -f audiodump.wav
  cd - &>/dev/null
done
IFS=$OIFS
```

Das wurde gerade im Basis Verzeichnis meiner Biblio gestartet und rödelt jetzt
erstmal vor sich hin.
Obacht, gerade die [IFS](http://stackoverflow.com/questions/4128235/bash-shell-scripting-what-is-the-exact-meaning-of-ifs-n) Umgebungsvariable ist bei dieser Variante extrem wichtig.

Eigentlich will man eher *.ogg. Würd ich auch echt gern. Aber mein mobiles
Abspielgerät will nicht.
