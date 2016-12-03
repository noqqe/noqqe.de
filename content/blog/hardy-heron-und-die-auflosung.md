---
aliases:
- /archives/72
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