---
aliases:
- /blog/2008/07/18/gobbymulti-cursor-file-editing-oder-so-ahnlich-d
- /archives/118
comments:
- author: Dr. Azrael Tod
  content: "<p>vlt nicht ganz so gut f\xFCr gr\xF6\xDFere Projekte geeignet wie CVS
    aber sicher sehr interessant zu lern- und spa\xDFzwecken...<br>Werd ich mir mal
    ansehen. :-)</p>"
  date: '2008-07-18T13:08:08'
- author: noqqe
  content: <p>Ach wir nehmen dich einfach mal in ein Session mit rein :)</p>
  date: '2008-07-18T13:14:56'
- author: charlysan
  content: <p>Naja, also mit CVS hat das ja eigentlich nichts zu tun, mit dem Editor
    kann man ja wirklich gleichzeitig live an einer Datei arbeiten (und hat nebenbei
    sogar noch einen ganz netten Chat).<br>Also ich hab damit auch schonmal rumgespielt
    und fand den echt cool, auch wenn man sowas ja nicht so oft braucht...</p>
  date: '2008-07-20T20:58:04'
date: '2008-07-18T08:38:15'
tags:
- development
- python
- ubuntu
- gobby
- linux
title: Gobby | Multi Cursor File Editing
---

Seit einigen Tagen spielen [CrackPod](http://crackpod.bplaced.net) und ich
gleichzeitig zusammen an PythonDateien herum. Ich hab auch noch nicht
verstanden wie das genau funktioniert.

Hier mal der Link zu Gobby:
[http://gobby.0x539.de/trac/](http://gobby.0x539.de/trac/)

> Gobby is a free collaborative editor supporting multiple documents in one
> session and a multi-user chat. It runs on Microsoft Windows, Mac OS X,
> Linux and other Unix-like platforms.

{{< figure src="/uploads/2008/04/250px-Gobby-0.4.0-dev-linux.png" >}}

Mithilfe eines Passworts lässt sich dann die Session betreten und mehrere
Dokumente bearbeiten/anlegen. Dadurch sind wir dann auf die Idee gekommen
uns unsere eigene kleine Welt zusammenzubasteln.

``` python
#!/usr/bin/python

class Person:
  bevoelkerung = 0
 def __init__(self, name, url):
  self.name = name
  self.url = url
  print'%s Wird geboren' % self.name
  Person.bevoelkerung +=1

  def blog(self):
  print 'Visit me on %s'  % self.url

  def __del__(self):
  print '%s ist tot!' % self.name
    Person.bevoelkerung -=1

    if Person.bevoelkerung == 0:
      print'ich bin der letzte'
    else:
      print'Es gibt noch %d Leute.' % Person.bevoelkerung
  def sagHallo(self):
    print 'Servus ich bin %s' % self.name
  def wieViele(self):
  if Person.bevoelkerung == 1:
    print 'ich bin ganz alleine hier*angst*'
  else:
    print 'Es leben hier %d Leute' % Person.bevoelkerung

noqqe = Person('noqqe', 'noqqe.de' )
noqqe.blog()
noqqe.sagHallo()
noqqe.wieViele()
crackpod = Person('crackpod', 'crackpod.bplaced.net' )
kpod.sagHallo()
crackpod.blog()
crackpod.wieViele()
noqqe.sagHallo()
noqqe.wieViele()
Julie = Person('Julie', 'julie.hat-gar-keine-homepage.de' )
Julie.sagHallo()
Julie.blog()
Julie.wieViele()
noqqe.__del__()
noqqe.wieViele()
```

oder wie immer hier :
[http://paste.pocoo.org/show/79759/](http://paste.pocoo.org/show/79759/)

Ausgabe sieht dann ungefähr wie folgt aus:

```
noqqe Wird geboren
Visit me on noqqe.de
Servus ich bin noqqe
ich bin ganz alleine hier*angst*
crackpod Wird geboren
Servus ich bin crackpod
Visit me on crackpod.bplaced.net
Es leben hier 2 Leute
Servus ich bin noqqe
Es leben hier 2 Leute
Julie Wird geboren
Servus ich bin Julie
Visit me on julie.hat-gar-keine-homepage.de
Es leben hier 3 Leute
noqqe ist tot!
Es gibt noch 2 Leute.
Julie ist tot!
Es gibt noch 1 Leute.
ich bin der letzte
crackpod ist tot!
```
