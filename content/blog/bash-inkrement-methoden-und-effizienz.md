---
aliases:
- /blog/2010/05/26/bash-inkrement-methoden-und-effizienz
- /archives/1014
comments:
- author: Paradiesstaub
  content: "<p>Danke f\xFCr die n\xFCtzliche Recherche. Mein Favorit war bisher ((i++)).</p>"
  date: '2010-05-27T00:26:21'
- author: noqqe
  content: "<p>:) wir wurde auch gezwitschert, dass deine Methode eine Integer-Variable
    erzwingt. </p><p>Dann w\xFCrde ich auch dahin umsatteln. Gerne wieder.</p>"
  date: '2010-05-27T00:39:44'
- author: user unknown
  content: "<p>Da finde ich doch <br>[code]j=5 ; k=12 ; i=$((j*10+k/2)); echo $i<br>[/code]<br>k\xFCrzer
    und praktischer.</p><p>Der Vollst\xE4ndigkeit halber w\xFCrde ich noch [code]<br>((i+=1));
    echo $i[/code]<br>erw\xE4hnen, aber nur der Vollst\xE4ndigkeit halber. </p><p>F\xFCr
    besonders gro\xDFe Zahlen <br>[code]<br>echo \"6^6^6\" | bc<br>[/code]<br>und
    Nachkommastellen empfiehlt sich dann bc, wie man hier sieht. <br>[code]<br>echo
    \"scale=42;pi=4*a(1);pi\" | bc -l <br>[/code]</p>"
  date: '2010-05-27T15:19:10'
date: '2010-05-26T18:22:56'
tags:
- development
- inkrementieren
- shell
- tricks
- increment
- linux
- decrement
- bash
title: Bash | Inkrement-Methoden und Effizienz
---

Des öfteren schlage ich mich in Bash-Scripting mit den Meldungen "unärer
Operator" und ähnlichen Fehlern herum. Inkrement über `i=$[$i+1]`
funktionieren unter entsprechenden Umständen nicht. In Bash mit Zahlen
umzugehen ist garnicht so leicht, finde ich. Manchmal werden die Zahlen als
Char statt Int interpretiert und manchmal darf man zu einer Variable mit
Wert 0 keine Addition durchführen. Das kann natürlich an meinen mangelnden
Fähigkeiten liegen, aber heute habe ich mich auf die Suche nach Lösungen zu
diesem Thema gemacht.

``` bash
# 1
let i+=1 # increment
let i-=1 # decrement
```

``` bash
# 2
let i++ # increment
let i-- # decrement
```

``` bash
# 3
i=$[$i+1] # increment
i=$[$i-1] # decrement
```

``` bash
# 4
i=$((i+1)) # increment
i=$((i-1)) # decrement
```

``` bash
# 5
((i++)) # increment
((i--)) # decrement
```

``` bash
# 6
i=$(expr $i + 1) # increment
i=$(expr $i - 1) # decrement
```

``` bash
# 7
: $[ n = $n + 1 ] # increment
: $[ n = $n - 1 ] # decrement
```

Das waren so die üblichen Varianten mit denen ich mir die letzten Monate
geholfen habe. Mitunter wirklich unschöne Sachen. Auf
[http://tldp.org/LDP/abs/html/declareref.html](http://tldp.org/LDP/abs/html/declareref.html)
bin ich dann allerdings auf eine sehr elegante Lösung gestossen. Via
"declare" lässt sich eine Variable auf einem bestimmten Typ festnageln.
"declare -i VAR" deklariert (haha) die Variable "VAR" nach Integer. So
lässt sich quasi direkt damit rechnen. Anschliessend das Beispiel um im
Bild zu bleiben:

``` bash
# Declare Variante
declare -i i # -i = integer
i=$i+1 # increment
i=$i-1 # decrement
```

Das geht wiederrum auch mit mehreren Variablen, Zahlen und
Rechenoperationen. Ein kleines einzeiliges Beispiel:

``` bash
j=5 ; k=12 ; declare -i i ; i=$j*10+$k/2 ; echo $i
```

Natürlich gibt es auch noch andere lustige Modi die man seiner Variable
aufs Auge drücken kann:

``` bash
declare -r # readonly
declare -a # array
declare -x # export variable
declare -f # function
```

Alles in Allem sehr nett. Werde jetzt erstmal ein paar Skripte umbauen.
Denke ich.
