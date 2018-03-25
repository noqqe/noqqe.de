---
aliases:
- /archives/1833
- /blog/2011/12/21/arduino-ich-bau-mir-einen-synthesizer
comments:
- author: Knorkebrot
  content: "<p>Mit Verlaub, das ist fett. Da bekommt man glatt auch Lust mal was in
    der Richtung zu machen.<br>Steuert man den Ton nur \xFCber die Potenziometer oder
    ist das auch programmierbar? Braucht man dazu ein gr\xF6\xDFeres Vorwissen in
    Sachen Elektronik? In Physik war ich trotz astreinem Mathe immer eine Niete. :/</p>"
  date: '2011-12-22T17:23:21'
- author: noqqe
  content: "<p>Hi Knorkebrot: Naja also generell kommen die analogen Signale von den
    Potenziometern. </p><p>Diese werden dann (vermute ich) gegen die Werte Tabellen
    in der Firmware gedr\xFCckt (oben im Code) und unten werden alle diese Werte dann
    miteinander verheiratet/verwurstet. </p><p>Du kannst also auch die Tables bzw.
    den ganzen Code anpassen um andere T\xF6ne zu erhalten. </p><p>Aber ich finde
    die Software daf\xFCr sehr Tricky weil Sie so Hardware nah ist und ich von Arduino
    Programmierung auch noch nicht sooooo die Welt Ahnung hab. </p><p>Hab mal mit
    ein paar Unterschiedlichen Werten im MidiTable experimentiert, aber im Endeffekt
    bin ich bei der Original Firmware geblieben. </p><p>Solltest dus auch bauen: Bloggen!
    :P</p>"
  date: '2011-12-22T18:19:51'
- author: noqqe
  content: "<p>Achja und: Nein man braucht kein gr\xF6\xDFeres Vorwissen in Elektronik.
    Das besitze ich n\xE4mlich auch definitiv _nicht_. </p><p>Lediglich die einfachsten
    Grundlagen sollte man kennen.</p>"
  date: '2011-12-22T18:25:07'
- author: Knorkebrot
  content: "<p>Hm, ja, h\xE4tte ich auch gleich mal in den Source gucken k\xF6nnen,
    nicht?<br>(*.pde ist \xFCbrigens nur ein C Source File, falls dar\xFCber noch
    jemand stolpert und sich extra die dicke Arduino IDE runterl\xE4d...)</p><p>Das
    sieht eigentlich nicht sonderlich schwer aus, es werden halt einfach nur Werte
    ausgelesen, gemapt und f\xFCr die Ausgabe zwischengespeichert. Was da dann passiert,
    muss einen ja nicht wirklich interessieren, aber auch das ist an sich keine Magie
    - musst nur wissen was mit Dreieckswelle, Phase, Grain(?) und alles meint. Keine
    Ahnung von dem Schnick Schnack. ;)</p><p>Aber man k\xF6nnte sich auf jeden Fall
    die Potenziometer sparen, wenn man eine vorprogrammierte Melodie haben will. Man
    m\xFCsste nur die Werte kennen, die dort herausfallen und sie sich abtippen.<br>Oder
    um die Idee etwas weiter zu spinnen, man k\xF6nnte einen Ethernet Arduino nehmen,
    dem einen Schalter verpassen, der in der loop() abgefragt wird. Ist er aktiv,
    wird die Tonausgabe ausgeschaltet und der Interrupttimer deaktiviert. Jetzt kann
    man auf dem Ethernetport lauschen und annehmen, was kommt, das schreibt man sich
    in das Array und schon hat man eine neue Melodie.<br>Legt man den Schalter wieder
    um, wird die Interruptschleife wieder gestartet.</p><p>Ich glaube ich spreche
    heute Abend mal mit dem Arduino-M\xE4n im \xF6rtlichen Chaostreff.</p>"
  date: '2011-12-22T21:49:04'
- author: noqqe
  content: "<p>Mh naja. Also das gibts ja oft zu h\xF6ren, dass die IDE so \xFCberladen
    w\xE4re und auch sonst nicht so toll. </p><p>Aber man darf nicht vergessen das
    Arduino als Lernplattform gegr\xFCndet wurde. Und genau das kann Arduino wirklich
    gut. </p><p>Ich brauch mir da keine riesen Gedanken machen wie die Software auf
    den Flash kommt usw. Ich dr\xFCcke drauf und es geht. Das finde ich schon sehr
    sehr gut. Eben weil ich keine Ahnung davon habe wie sowas in der Industrie funktioniert.</p><p>Jop
    das k\xF6nnte man auch in der Loop mit vordefinierten werten machen. Sicherlich.</p><p>Was
    hat denn dein \xF6rtlicher Arudino Mensch so gesprochen?</p>"
  date: '2011-12-27T14:08:14'
- author: Knorkebrot
  content: "<p>Oh, ich kenne die IDE ja gar nicht, ich hab mich nur ge\xE4rgert, dass
    mir das nicht vorher aufgefallen ist, dass ich nicht mehr als einen einfachen
    Editor brauchte, vim oder so. Ich hab mal schnell gegooglet, es gibt auf jeden
    Fall auch einige nette CLI-Tools:<br><a href=\"http://arduino.cc/playground/FreeBSD/CLI\"
    rel=\"nofollow\">http://arduino.cc/playground/FreeBSD/CLI</a></p><p>Ardunio-M\xE4n sagt:
    Im Bereich des M\xF6glichen. (Er ist aber auch noch ziemlich am Anfang mit Ardunino,
    wusste ich nicht :P ) Ich werde mich mal nach Silvester nach der Hardware umsehen,
    da sollte ich etwas mehr Zeit im Alltag haben, richtig cool w\xE4re das dann ja
    auch, wenn der noch PoE kann. Gibt ja auch dazu Erweiterungen f\xFCr den Arduino
    und ein PoE Injector ist ja nicht sonderlich teuer.</p>"
  date: '2011-12-28T19:42:35'
- author: noqqe
  content: <p>test</p>
  date: '2011-12-28T21:30:59'
- author: noqqe
  content: <p>test2</p>
  date: '2011-12-28T21:31:19'
- author: Guest
  content: "<p>Hey, ich w\xFCrde den auch gerne nachbauen, allerdings finde ich auf
    der Seite den Download f\xFCr den Source-code nicht ... Oder ist der komplette
    code schon das kleine Popel ding da, welches man nurnoch in eine schleife setzen
    muss, wenn ja wie muss ich die inputs und alles definieren?<br>Am besten du postest
    einfach mal den kompletten Code, w\xE4re sehr nett :)</p><p>Au\xDFerdem: kann
    man als Potis auch x-beliebige lineare 4.7kohm potis nehmen oder braucht man die
    vom Tinkershop?</p>"
  date: '2012-12-09T20:03:45'
- author: Anonymous
  content: "hey ;)\nIch wei\xDF, dass das bei dir jetzt schon eine Weile her sein
    mag, allerdings w\xFCrde ich den auch gerne nachbauen, hab aber absolut keine
    Ahnung von Arduino und Programmieren. Bekomm ich das trotzdem hin und wie lange
    wird es dauern? Weisst du das?\nUnd welche Bauteile w\xFCrde ich daf\xFCr ben\xF6tigen?\n"
  date: '2014-02-23T11:53:45.354058'
- author: Anonymous
  content: 'Wo finde ich genau den Code?

    Alles super aber den Code finde ich nicht'
  date: '2014-04-12T10:32:14.294218'
date: '2011-12-21T15:30:58'
tags:
- kabel
- web
- "l\xF6ten"
- auduino
- arduino
- lfo
- wiring
- community
- verkabeln
- hardware
- synthesizer
- soldering
- synthi
- low frequenzy synthesizer
title: Arduino | Ich bau mir einen Synthesizer
---

Vor ca. 2 Wochen habe ich auf der [Suppe vom K4CG](http://k4cg.soup.io) ein
[Video](http://vimeo.com/2266458) über einen auf Arduino basierten
Synthesizer. Die "Firmware" die darauf läuft nennt sich "Auduino".

Auf deren [Projektseite](http://code.google.com/p/tinkerit/wiki/Auduino)
habe ich mich dann etwas schlau gelesen und wie auch schon bei dem
[Megabitmeter](/archives/1751) über tinkersoup.de meine Teile bestellt. Auf
der Projektseite ist die Konstruktion des ganzen finde ich zwar nicht
sonderlich gut beschrieben, aber man kommt mit ein bisschen Googeln und
Reverse Engineering schon weiter.

Habe dabei aber einen [Arduino Nano](http://arduino.cc/en/Main/ArduinoBoardNano) benutzt, weil mir die
Anschlüsse bzw "Architektur" besser gefällt und ich nicht erst ein Breakout
Board von Seriell auf USB nachkaufen musste. Entgegen aller Erwartungen
musste ich die Firmware dafür nichtmal modifizieren, da auch bei diesem
Board ein ATMega328 verbaut ist.

Ich habe mir wegen der einfacheren Anbringung am Nano so eine Art
Halterung/Breadboard mitbestellt, in dem ich die Adern mit Schrauben
einfacher verbauen konnte.

{{< figure src="/uploads/2011/12/IMG_0604.jpg" >}}

Die Potenziometer (wieder was, das ich gelernt habe) sind in Reihe an den
Ground und den 5V Pin geschlossen. Der jeweils mittlere Pin der
Drehschalter kommt an die Analog Pins 0 bis 4.

{{< figure src="/uploads/2011/12/IMG_0606.jpg" >}}

Danach kam der (für mich) kniffligere Teil. Der Audio Jack (bzw. Klinke
Buchse) hat von Haus aus 5 Pins. Auf der Projekte Seite von Auduino nur
Input und Ground. Nach bisschen schlaulesen in Wikis und Foren scheint es,
als würden die verschiedenen Revisionen von Klinke andere Features mit sich
bringen. für den Mini Synthesizer hätte vollkommen Klinke Mono ausgereicht.
Diverse Zusatzfunktionskanäle sind da eigentlich überflüssig aber im
[Audio Jack bei TinkerSoup](http://www.tinkersoup.de/product_info.php?products_id=74&osCsid=3c2172e4114e78d30b2788b3cd0d6077)
integriert.

Nach etwas Trial and Error Verfahren den weg für Doofe gewählt. Ich hab
ehrlich gesagt einfach ein altes Klinke Stecker auf Buchse Kabel
aufgeschnitten und mir die Belegung auf der Steckerseite angesehen.

{{< figure src="/uploads/2011/12/IMG_0609.jpg" >}}

Bei 3poligen Klinken Steckern sind die vorderen beiden Kontakte fürs Signal
(Links, Rechts) und hinten für Ground. Habe dann die beiden Signaladern auf
der Buchsenseite verdrillt und wie vorgesehen in den Digitalen Pin 3
geklemmt. Ground natürlich an seine Stelle.

{{< figure src="/uploads/2011/12/IMG_0611.jpg" >}}

Im Endeffekt wars dann schon fertig. Firmware mit dem Arduino IDE auf den
Chip geladen und hat auch schon funktioniert.  Aber weil ich dann ständig
die Potenziometer durcheinander gebracht habe, hab ich noch eine alte
Plastikbox aus dem Baumarkt meiner Wahl benutzt, die entsprechenden Löcher
gebohrt dort das ganze eingebaut.

{{< figure src="/uploads/2011/12/IMG_0616.jpg" >}}

Etwas smoother ;) Noch ein paar kleine Kostproben von einem wirklich
unbegabten Synthesizer-Bediener. Beim hören etwas aufpassen, ab und zu ist
mir da ein Ton entglitten.

[Auduino Sample Mp4](/uploads/2011/12/record.mp4)
[Auduino Sample WAV 16 bit 22 Kkhz](/uploads/2011/12/record16bit-22kkhz.wav)
