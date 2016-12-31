---
aliases:
- /blog/2012/07/29/star-dot-wma-zu-star-dot-mp3
comments:
- author: bakudan
  content: "<p>was mir etwas mutig erscheint ist das du davon ausgehst das der mplayer
    Aufruf keinen Fehler zur\xFCck gibt, geht hier etwas schief ist das Origenal weg,
    oder etwa nicht?<br>Zumindest f\xE4llt mir hier jetzt nichts auf was daf\xFCr
    sorgt das das \"rm -f\" nicht ausgef\xFChrt wird.</p>"
  date: '2012-08-01T11:50:01'
- author: noqqe
  content: "<p>Oh \xE4hm. Da hast du vollkommen recht, sollt ich wenigstens mit ein
    paar logischen UNDs verkn\xFCpfen.. </p><p>Danke f\xFCr den Hinweis!</p>"
  date: '2012-08-02T06:12:49'
- author: bakudan
  content: "<p>Bitte, ich war mir hier nur nicht sicher ob die shell hier vielleicht
    irgendwie abbricht, im Moment mache ich viel mit Windows Batch und da muss man
    extrem aufpassen, besonders auf solche F\xE4lle (und auf &amp; in File Namen ....).</p><p>Kannst
    du mir hier ein Beispiel zeigen wie man das in bash sauber macht?<br>Ich glaub
    es gibt hier bessere M\xF6glichkeiten als wie unter Windows die Error-Variable
    mit If zu \xFCberpr\xFCfen.</p>"
  date: '2012-08-02T08:50:46'
- author: noqqe
  content: "<p>Da muss man nicht nur bei Batch aufpassen :) Hier eigentlich auch.</p><p>lame
    --preset standard audiodump.wav -o \"`basename \"$i\" .wma`.mp3\" &amp;&amp; rm
    -f $i</p><p>W\xE4re sauberer :)</p>"
  date: '2012-08-03T12:24:16'
date: '2012-07-29T15:39:00'
tags:
- web
- convert
- lame
- konvertierung
- konvertieren
- mp3
- wma
- debian
- bash
title: '*.wma zu *.mp3'
---

Man denkt ja immer "Gott, wie viele von diesen $MUSIKFORMAT_A zu $MUSIKFORMAT_B
Formatierungsblogposts kann mein Reader eigentlich vertragen".
Das denkt man zumindest so lange, bis man mal in die Verlegenheit gerät sowas zu
brauchen.

{{< figure src="/uploads/2012/07/koala.jpg" >}}

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

Das wurde gerade im Basis Verzeichnis meiner Biblio gestartet und rödelt
jetzt erstmal vor sich hin.  Obacht, gerade die
[IFS](http://stackoverflow.com/questions/4128235/bash-shell-scripting-what-is-the-exact-meaning-of-ifs-n)
Umgebungsvariable ist bei dieser Variante extrem wichtig.

Eigentlich will man eher *.ogg. Würd ich auch echt gern. Aber mein mobiles
Abspielgerät will nicht.