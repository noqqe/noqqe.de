---
aliases:
- /archives/72
comments:
- author: Dominik
  content: "<p>Was f\xFCr einen Bildschirmtreiber hat displayconfig erkannt? Manchmal
    muss man die Verwendung der propriet\xE4ren Treiber erzwingen und mit denen klappt's
    bei mir recht gut bei 1680x1050...<br>Allerdings nur solange ich meinen (analogen)
    Monitor nicht via DVI-Adapter anschliesse, sondern nur direkt am (2.) VGA-Ausgang.
    \xDCber den DVI-Adapter klappt wohl die Abfrage der Bildschirmparameter mittels
    DDC nicht.<br>Ansonsten k\xF6nnte man noch versuchen displayconfig die Windows
    INF-Datei f\xFCr den Monitor anzugeben. Ich meine mich zu erinnern, da\xDF das
    geht.</p>"
  date: '2008-06-13T01:44:16'
- author: seufz
  content: "<p>Aha ;)<br>Um es mit den Worten des Kommentators des Deutschlandspiels
    auszudr\xFCcken \"Eine Verkettung von Zuf\xE4llen hat zum Tor (Ziel) gef\xFChrt!\"</p><p>Danke
    f\xFCr die Anst\xF6\xDFe!<br>Ganze Geschichte kommt im n\xE4chsten Blogeintrag!</p>"
  date: '2008-06-13T20:53:30'
date: '2008-06-12T12:38:24'
slug: hardy-heron-und-die-auflosung
tags:
- hardware
- hardy heron
- ubuntu
- linux
title: "Hardy Heron und die Aufl\xF6sung"
---

Never Ending Story möchte ich fast schon sagen. Kann an der aktuellen
Beruffschulwoche und dem Lernstress liegen aber ich bekomm es einfach nicht
hin!!!

Die Einstellungen wollen einfach nicht so wie ich mir das vorstelle!
Breitbildauflösüng = Fehlanzeige.Ich sitz hier allen Ernstes mit 1280x1024
und Tippe diesen Eintrag. Ich schaff es einfach nicht und ich hab auch
nicht genug Zeit mich wirklich darum zu kümmern.

Jedesmal wenn ich eine falsche Einstellung wähle muss ich sowieso erst neu
booten im recoverymodus, xserver fixen und dann normal wieder mit 640x400
neustarten.

Ich geh dran kaputt, ich sags euch. Im allgemeinen läufts aber auch nich
umbedingt besser. Der Kartenleser für SD funktioniert auch hier nicht. Also
steht schonmal ein Hardware problem fest. Die MS geht wunderbar. Aber mein
Gott.

Ich werds noch etwas versuchen.

Genauer übrigens :

Plug and Play Monitor: Nur 640x400 auswählbar!? (Danke video7 Bildschirm)

LCD 1680x1050 : Auch nicht mehr auswählbar.

Haken bei "widescreen" gesetzt: folgende Fehlermeldung:

```
$ npx@kiwi:~$ sudo gksu displayconfig-gtk
[sudo] password for npx:
Traceback (most recent call last):
File "/usr/lib/python2.5/site-packages/displayconfiggtk/DisplayConfig.py", line 764, in on_treeview_screens_cursor_changed
self.get_rates_for_current_resolution()
File "/usr/lib/python2.5/site-packages/displayconfiggtk/DisplayConfig.py", line 873, in get_rates_for_current_resolution
for rate in device.getAvailableRefreshRates():
File "/var/lib/python-support/python2.5/displayconfigabstraction.py", line 2092, in getAvailableRefreshRates
return self.getAvailableRefreshRatesForResolution(self.currentsizeindex)
File "/var/lib/python-support/python2.5/displayconfigabstraction.py", line 2104, in getAvailableRefreshRatesForResolution
isize = self.available_sizes[resolution_index]
IndexError: list index out of range
```

Absolut Klasse. Im moment verbring ich meine Zeit damit entwerder
Herrauszufinden wo das problem liegt bei GKSU oder ich probiere Wahllos
irgendwelche 22" WideScreens aus ob es klappen könnte!
