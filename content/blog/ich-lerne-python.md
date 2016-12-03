---
aliases:
- /archives/104
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