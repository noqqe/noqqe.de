---
aliases:
- /archives/41
comments:
- author: Marcus Kober
  content: "<p>Habe auch das Problem, dass ich nach jedem Start gksu displayconfig-gtk
    ausf\xFChren und dort die richtige Aufl\xF6sung angeben muss. Hast Du mittlerweile
    einen Weg gefunden, die Einstellungen zu speichern?</p>"
  date: '2008-04-28T23:06:26'
- author: seufz
  content: <p><a href="http://seufz.wordpress.com/2008/04/28/hachjanvidia/" rel="nofollow">http://seufz.wordpress.com/2008/04/28/hachjanvidia/</a><br>;)</p><p>Sollte
    deine Frage beantworten.<br>Falls noch was ist sag Bescheid. Ich helf gerne.</p>
  date: '2008-04-29T00:11:59'
- author: Lukas
  content: <p>hi<br>die 3d-effekte kannst du normalerweise mit dem compizconfig-settings-manager
    einstellen/aktivieren.</p>
  date: '2008-05-01T19:21:51'
- author: seufz
  content: "<p>Danke f\xFCr den Tipp! werd ich mal probieren!</p>"
  date: '2008-05-02T10:45:36'
- author: noqqe
  content: test
  date: '2013-11-17T17:16:41.405467'
date: '2008-04-25T23:16:03'
slug: hardy-macht-spasironie-off
tags:
- hardware
- hardy heron
- ubuntu
- linux
title: "Hardy macht Spa\xDF!Ironie off."
---

Gestern hab ich Hardy bei durchschnittlich 50kb/s runtergeladen und installiert. Allein das war ja schon sehr aufmunternd.
Und auch ansonsten lief es ziemlich "worst case" like! Hier eine kurze Liste der Probleme die sich
mir mehr oder weniger in den Weg stellen:

  * **Bildschirmauflösung**: Hardy hatte ich gerade fertig installiert und gebootet als
  sich mir eine 800x600 Bildschirmauflösung bot."System -> Einstellungen -> Bildschirmauflösung -> maximal mögliche Auflösung: 800x600".
  Da freut man sich doch! Irgendwann bin ich dann drauf gekommen einen anderen Bildschirm einzustellen.
  Auch nicht so einfach in Hardy.... Wo ist denn dieser Abschnitt in der Systemverwaltung "_Bildschirme und Grafik_" hin ?! Dank Google hab ich dann den Befehl
  "gksu displayconfig-gtk" gefunden und konnte das ganze per Terminal starten :) Eingestellt und geht!
  Aber nach jedem Neustart wirds wieder zurück gesetzt... Speichern?! Ich hoffe ich finds noch raus.

  * Und wo wir gerade bei **Neustart** sind: Bei jedem reboot muss ich jedesmal das Standard Bootbild 4-5 mal ertragen bevor ich dann
  endlich in mein 800x600 System darf. Grund? Keine Ahnung! Ich hoff ich finds raus!

  * **3D-Effekte**: Auch die sind leider weg und ich muss mich erst wieder Schlaulesen
  wie der Manager hies mit dem ich diese wieder "enabeln" kann :P

Trotz allem gibts auch ein paar gute Sachen :). Vielleicht bilde ichs mir ein. Aber ich finde Hardy läuft
deutlich schneller!
