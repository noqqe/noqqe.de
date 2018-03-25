---
aliases:
- /archives/104
comments:
- author: Dr. Azrael Tod
  content: "<p>zum produktiven Arbeiten empfehle ich \xFCbrigens Google App-Engine
    ;-)<br>Kostenloses Webhosting mit Datenbank in Python, dazu gibts auch noch eine
    ziemlich gute Einf\xFChrung.</p>"
  date: '2008-07-10T00:15:04'
- author: zypral
  content: <p>Ein Buch das ich mir bestellt habe und empfehlen kann:<br><a href="http://www.amazon.de/Objektorientierte-Programmierung-Python-Michael-Weigend/dp/382661660X/ref=sr_1_7?ie=UTF8&amp;s=books&amp;qid=1215634971&amp;sr=1-7"
    rel="nofollow">http://www.amazon.de/Objektorientierte-Programmierung-Python-Michael-Weigend/dp/382661660X/</a></p>
  date: '2008-07-10T00:23:26'
- author: noqqe
  content: "<p>Danke f\xFCr die Empfehlungen! Ein neues Buch kauf ich mir erst wenn
    ich mti meinem Linux-Befehle Buch durch bin. Dauert nichtmehr lange! Aber wenn
    hab ich da schon was im Auge.<br>Eventuell sogar <a href=\"http://www.galileocomputing.de/openbook/python/\"
    rel=\"nofollow\">http://www.galileocomputing.de/openbook/python</a></p><p>Google App-Engine
    werd ich mir auch mal anschaun!</p>"
  date: '2008-07-10T01:04:28'
- author: zypral
  content: <p>Hatte ich auch angefangen, bis ich das gelesen habe:<br><a href="http://www.python-forum.de/topic-13008.html"
    rel="nofollow">http://www.python-forum.de/topic-13008.html</a></p>
  date: '2008-07-10T22:30:18'
- author: CracKPod
  content: "<p>Hey, was zypral sagte stimmt, solltest du nicht kaufen.</p><p>Ich w\xFCnsche
    dir viel Spa\xDF mit Python und hoffe du hast relativ schnell Erfolge. :)</p><p>*Wollte
    eigentlich nur einen Kommentar machen und hat einen vern\xFCnftigen Vorwand gesucht*</p><p>MfG,<br>CracKPod</p>"
  date: '2008-07-11T00:46:12'
- author: charlysan
  content: "<p>Python ist wirklich genial, bin auch am rumspielen damit und dokumentiere
    dabei so das ein oder andere auf meinem Blog, also wen sowas interessiert, der
    kann ja mal vorbeischauen.<br>Das mit der Google App Engine ist ein cooler Tip,
    das werd ich mir mal angucken.<br>Als Buch kann ich btw \"Programming Python\"
    von O'Reilly empfehlen, ist zwar mit 1500 Seiten ein ganz sch\xF6ner Schinken
    und auch nicht gerade preiswert, aber daf\xFCr wirklich gut.</p><p>Gru\xDF Charlysan</p>"
  date: '2008-07-16T13:58:41'
date: '2008-07-09T20:09:40'
slug: ich-lerne-python
tags:
- development
- python
- linux
title: Ich lerne Python
---

Ist eigentilch Recht einfach. Aber ich glaube das liegt auch ein bisschen
an meinen 2 Jahren C++ in der Berufsschule. Hier meine ersten paar kleinen
Skripte:

Variablen und Operatoren

``` python
#!/usr/bin/python
laenge = 5
breite = 2
flaeche = 2 * (laenge + breite)
print 'laenge' , laenge
print 'breite' , breite`` print 'flaeche ist: ' , flaeche
```

Schleifen und Breaks

``` python
#!/usr/bin/python

while True:
  s = raw_input ('geben sie etwas ein ')
  if s == 'ende' :
    break
  if len(s) < 3:
    continue
  print 'Laenge ist ausreichend', len(s)
  print 'fertig'
```

Globale und Interne Variablen

``` python
#!/usr/bin/python

def funk():
  global x
  print 'x ist', x
  x = 2
  print 'lokales x ist jetzt' , x
  x = 50

funk()
print 'der wert von x ist' ,x
```

und hier das erste Skript mit Unterhaltungswert :D
Zahlenraten.

``` python
#!/usr/bin/python
zahl = 23
weiter = True
  while weiter:
    geraten = int(raw_input('geben sie eine zahl ein'))

  if zahl == geraten:
    print 'glueckwunsch, du hasts!'
    weiter = False
  elif geraten < zahl:
    print 'nein die zahl ist etwas hoeher'
  elif geraten > zahl:
    print 'nein die zahl ist etwas niedriger'
  else:
    print 'schleife beendet'
print 'Fertig'
```

Ich hoff ich kann bald produktiv damit arbeiten :)
