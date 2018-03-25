---
comments:
- author: Thomas
  content: <p>"fucking lesbar" bedeutet schlecht lesbar oder?</p>
  date: '2013-09-09T08:32:18'
- author: noqqe
  content: <p>Ne, gut! :)</p>
  date: '2013-09-09T08:43:06'
- author: nuit
  content: '<p>*applaus* Fazit: zum Debuggen sinnvoll, wirkliche Verwendung in hardcoded
    Stuff geht garnicht.</p>'
  date: '2013-09-09T10:00:47'
- author: Erik Wasser
  content: "<p>Ich fand den Award schon immer \xFCberbewertet. Bashzeilen mit \"cat\"
    sind lesbarer und k\xF6nnen f\xFCr das Debugging leichter ausgetauscht werden.</p>"
  date: '2013-09-12T09:29:11'
- author: mina86
  content: <p></p>
  date: '2013-10-30T12:33:55'
- author: mina86
  content: <p>Issues with less then sign it seems. Also, sorry, I actually don't speak
    German. ;) Either way, maybe it'll work this time, what I'm trying to point out
    is:</p><p>&lt; file-name comma</p>
  date: '2013-10-30T12:36:52'
date: '2013-09-08T18:28:00'
tags:
- development
- shell
- useful
- use
- cat
- unix
- useless
- linux
- bash
- usability
title: Useful use of cat
---

Eine dieser langen Traditionen unter Unix Usern ist es über
[useless-use-of-cat](https://en.wikipedia.org/wiki/Cat_\(Unix\)#Useless_use_of_cat)
zu ragen. Unbegründet finde ich.

Auf Twitter erklärt jemand Map-Reduce mithilfe einer Versinnbildlichung
unter Unix.

<blockquote class="twitter-tweet"><p>Explaining map/reduce to Unix folks <a
href="http://t.co/mtF3SvFMaC">pic.twitter.com/mtF3SvFMaC</a></p>&mdash; Todor Genov (@tgenov) <a
href="https://twitter.com/tgenov/statuses/374915872932237312">September 3,2013</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

Was in den Replies alles auftaucht sollte jetzt keine Überraschung mehr sein.
Mich beschleicht das Gefühl, dass mittlerweile mehr aus Gewohnheit drauf los
getrollt wird. Dabei gibt es so viele Vorteile.

{{< figure src="/uploads/2013/09/ahm.gif" >}}

### Lesbarkeit

Ganz ehrlich, `bash`-Zeilen die mit `cat` beginnen sind fucking lesbar. Es
ergibt sich sofort dieses Gefühl der Streamline `stdout` String Verarbeitung.
Gerade für Leser von Tutorials / HowTos / Blogposts sind solche Kommandofolgen
wesentlich verständlicher. Sachverhalte sind so leichter darstellbar,
educational use, sozusagen.

### Benutzbarkeit

Usage Gewinn. Ultimativ. Für den Tippenden vor der Tastatur ist das führende `cat`
oft einfacher zu handhaben.

Warum? Änderung der Reihenfolge bei Verarbeitung ist einfacher. Entscheide ich mich
die `sed` Regex nun doch erst nach dem umsortieren mit `awk` oder `tr`
durchzuführen brauche ich nicht erst die Parameter der Commands anpassen sondern
einfach nur die nachgelagerten Filter umsortieren.

Ein anderer Punkt der unter Benutzbarkeit fällt, ist das "Entwickeln" von
Mechanismen bei großen Files. Wenn ich meine Filter ausprobieren möchte nutze ich oft das Bash Suche/Ersetze Feature.
Beispielsweise `^cat^head`. Aber auch einfachem Drücken der `Pos 1` bzw
`Ctrl + a` Taste ist `cat` durch die Syntax ein sehr dankbarer Befehl.

Bei Logfile-Auswertungen bastle ich oft länger an
der passenden Regex, wobei ich nicht jedes verdammte mal von `sed` die vollen
10GB durchwühlen lassen möchte. Stattdessen arbeite ich eher nach folgendem
Prinzip:

``` bash
head access.log | sed 's/foo/bar/g' | [...]
tail access.log | sed 's/foo/bar/g' | [...]
# wenn ok
^head^cat
# oder über Ctrl-a tippen
cat access.log | sed 's/foo/bar/g' | [...]
```

Das spart mir Zeit, Nerven und Tastendrücken.

### Performance

Einen zusätzlichen Prozess forken tut heutzutage niemandem
mehr weh. Gerade bei lustigen Einzeilern die Formate in Ordnung bringen oder IO
generieren ist der neue Prozess das geringste Problem.

Wo sich meine Meinung allerdings komplett ändert, ist wenn es um `cat` in
Scripts geht. Wartbarkeit und Effizienz erhalten Vorrang gegenüber der
Benutzbarkeit und dem Bildungsauftrag.

So, meine 2 Cent. Ihr dürft jetzt loslegen.
