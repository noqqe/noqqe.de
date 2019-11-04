---
aliases:
- /blog/2013/02/15/wenn-etwas-gross-wird-und-du-merkst-du-hast-gefailed/
- /blog/2013/02/15/wenn-etwas-gro%C3%9F-wird-und-du-merkst-du-hast-gefailed./
- /blog/2013/02/15/wenn-etwas-gro%C3%9F-wird-und-du-merkst-du-hast-gefailed/
comments:
- author: Thomas Do
  content: <p>All of old.<br>Nothing else ever. </p><p>Ever tried. </p><p>Ever failed.
    </p><p>No matter. </p><p>Try again. </p><p>Fail again. </p><p>Fail better.</p><p>Samuel
    Beckett, Worstward Ho</p>
  date: '2013-02-15T14:40:14'
- author: spion
  content: <p>wie waere folgendes:</p><p>Du baust eine alternative Funktion, die das
    Ganze in Unix-Epoch realisiert und migrierst dann deine Datenbank? ;)</p>
  date: '2013-02-15T15:26:17'
- author: noqqe
  content: "<p>Right, das ginge. </p><p>Problem ist jedoch, dass ich jegliche Kalkulation
    von zeitlichen Values ins Backend also die Datenbank ausgelagert habe. Daten werden
    w\xE4hrend des SQL Querys bereits so modifiziert wie ich sie in der PHP Anwendung
    brauche. So kann ich Fehler fr\xFCh abfangen. </p><p>Ich m\xFCsste also bei so
    einem Umzug _alle_ Queries und alle Graphen/Charts anpassen und umbauen. </p><p>Das
    m\xF6chte ich nicht! :P</p>"
  date: '2013-02-15T17:35:14'
- author: spion
  content: <p>naja ... einen tot wirst du wohl sterben muessen .... nichts destotrotz
    solltest du so oder so deine zeitstempel pruefen, bevor du sie in die datenbank
    tackerst ;)</p>
  date: '2013-02-15T17:53:28'
- author: noqqe
  content: <p>Wer sagt denn das ich sie nicht validiere? :) Genau diese Mechanismen
    umzubauen will ich mir ja ersparen.</p>
  date: '2013-02-15T17:57:30'
- author: ChristophLSA
  content: "<p>Wieso l\xE4sst du den User bei der Registrierung nicht einstellen,
    in welcher Zeitzone er lebt? Ggf. kannst du bei jeder neuen Angabe der zugenommenen
    Menge das Feld auch einbauen (falls man wirklich gerade wo anders ist). Wenn du
    dann erreichst, dass die bestehenden User die timezone updaten, dann kannst du
    die alten Daten ja fixen, denn wie ich sehe hast du die Userids mitabgespeichert.
    ;)</p>"
  date: '2013-02-15T18:06:48'
- author: noqqe
  content: "<p>Coole Idee, aber folgende Punkte daran find ich nicht so prall:</p><p>*
    Unn\xF6tige Verkomplizierung der Bedienung der Site f\xFCr den User<br>* \xC4nderung
    im DB Schema daf\xFCr n\xF6tig<br>* Ich muss jeden Timestamp erst umrechnen in
    der Anwendung um vereinheitlichbare  werte f\xFCr die Graphen zu haben. <br>*
    Ich kann bestehende User nicht dazu zwischen ihre Zeitzone einzustellen :)</p><p>Mit
    dem JS kommt automatisch die richtige Zeit f\xFCr den User und f\xFCr mich an.
    Und niemand muss ich darum wirklich k\xFCmmern (au\xDFer ich beim JS schreiben
    kurz mal...) Keine DB Schema Anpassung und kein herumgerechne.</p>"
  date: '2013-02-16T08:48:24'
- author: noqqe
  content: <p>Sehr geil :)</p>
  date: '2013-02-16T20:50:11'
- author: spion
  content: <p>fauli!</p>
  date: '2013-02-17T21:38:30'
- author: Maik
  content: <p>Uebrigens geht auch das Captcha bei der Registrierung nicht. Scheint
    auf irgendeine Nicht-SSL-Ressource zuzugreifen ;)</p>
  date: '2013-02-19T16:56:40'
- author: Maik
  content: <p>Das momentane JavaScript auf der Seite ist uebrigens immernoch leicht
    problematisch. Geht der User auf die Update-Seite, laesst sie ein paar Stunden
    offen und klickt erst dann auf update, ist d immernoch das Datum vom Aufruf der
    Seite, nicht vom Klick auf den Absenden-Button. Hatte ich vorhin. Hat mich arg
    verwirrt. ;)</p><p>Kann man damit beheben, indem man die Zeile mit "d=new Date()"
    in die AddPostData-Funktion schiebt. :)</p>
  date: '2013-02-20T14:54:41'
- author: mike
  content: '<p>Auf gleichen Fehler wollte ich auch hinweisen. Aber das hat Maik ja
    schon getan.</p><p>Weiterhin: Ich finde das Projekt Klasse. :-)</p><p>moikmoikmoik</p>'
  date: '2013-02-21T10:02:58'
- author: noqqe
  content: "<p>Danke f\xFCr die Tipps und das Lob :) </p><p>Ich werd das fixen. Glaub
    der Javascript Hint ist gut :)</p>"
  date: '2013-02-21T11:04:47'
- author: noqqe
  content: <p>Mal schauen ob ich die ReCaptcha Library dazu bringe ssl Content von
    Google zu laden :)</p>
  date: '2013-02-21T11:05:23'
- author: Sebastian Gaul
  content: '<p>Einfacher ist manchmal einfacher: Warum nicht einfach anonym melden
    lassen?</p>'
  date: '2013-02-23T11:10:54'
- author: noqqe
  content: <p>Ich verstehe nicht was du meinst</p>
  date: '2013-02-23T11:31:28'
- author: Sebastian Gaul
  content: "<p>Na dass Besucher ohne Registrierung und Anmeldung auf einen dicken
    Button \"Hatte nen Kaffee\" dr\xFCcken k\xF6nnen. Oder so \xE4hnlich.</p>"
  date: '2013-03-04T21:28:56'
date: '2013-02-15T11:15:00'
tags:
- development
- google charts
- fail
- mysql
- php
- debian
title: "Wenn etwas gro\xDF wird und du merkst du hast gefailed."
---

Zugegeben, ein Projekt über das ich im Blog noch nie gesprochen habe ist
[coffeestats.org](https://coffeestats.org). Holger wollte sich mit PHP beschäftigen, ich mit [Google Charts](https://developers.google.com/chart/). Wir taten uns zusammen. Das ist
nun über ein Jahr her. Es geht grundsätzlich nur darum den eigenen Koffeein-Konsum zu
tracken und statistisch auszuwerten. Auf der Site haben im Grunde nur Holger,
ein paar Kollegen und ich rumgelurked.

Vor kurzem empfand
[Jan](http://blog.waja.info/2013/02/13/tracking-coffee-consumption/) es für sinnvoll einen kleinen
Blogpost über das Projekt zu schreiben. In [planet.debian.org](http://planet.debian.org).
Klar die Hits würden ein bisschen hochgehen. Aber so wirklich an Registrierungen hab ich
nicht geglaubt.

{{< figure src="/uploads/2013/02/ohgod.gif" >}}

Tatsächlich haben sich aber wirklich Menschen angemeldet. 40 Stück in am
ersten Tag. Von London, Montreal, Zürich, Schweden über Griechenland zu Madrid.
Alles dabei. Fand ich witzig. Bis ich mir die "Overall Statistics" ansah. Wer
zur Hölle säuft um 3 Uhr nachts Kaffee... und bucht das auch noch auf
coffeestats.org? Erstmal etwas rumdebuggt.

``` sql
INSERT INTO cs_coffees VALUES ('', '$userid' , NOW() );
```

Tjo. Problem gefunden. Wann immer jemand (egal wo) auf "Ich hab gerade ne
leckere Tasse Kaffee vor mir stehen" drückte, hab ich diese unter der aktuell
vorherrschenden ServerTime abgespeichert. Für Holger und mich hätte das ja super
funktioniert. Wir hätten damals nicht gedacht das sich dort überhaupt jemand
anmeldet. Schon garnicht jemand der nicht in unserer Zeitzone wohnt (und kein Bot
ist).

> Time zone settings are normally not the first thing you think of when
> working on a web project with PHP

Was tun? Auch wenn ich mir sicher bin dass die registrierten User da nicht auf
Dauer aktiv sein werden, hatt mich das trotzdem genervt. Was ist wenn ich mal in
den Urlaub fahre und Kaffee trinke ... mit einem falschen Timestamp?
Schrecklich!

Nach etwas herumgoogeln stellte sich heraus, dass man die Client Zeit wohl mit
JavaScript feststellen möchte. Und das Date Formatting mit JS unglaublich
hässlich ist.

``` javascript
function coffeetime(d){
  function pad(n){return n<10 ? '0'+n : n}
  return d.getFullYear()+'-'
    + pad(d.getMonth()+1)+'-'
    + pad(d.getDate())+' '
    + pad(d.getHours())+':'
    + pad(d.getMinutes())+':'
    + pad(d.getSeconds())}

var d = new Date();

function AddPostData() {
  document.getElementById('coffeetime').value = coffeetime(d);
  document.getElementById('form').submit();
}
```

Resultiert in dem tollen Stamp `2013-02-15 12:47:41` den ich an die gewünschte HTML Form bastle.
Mit viel Handliebe Nullen padden müssen, 6 Funktionen callen und getMonth Index Workaround. Pfui.

Dann noch durch den XSS und SQL Injection Validator schicken und fertig. Jetzt
sind die User im Zweifel selbst für falsche Timestamps verantwortlich und ich
kann beruhigt $irgendwohin in Urlaub fahren.

Sollte ich nochmal vor so eine Aufgabe gestellt werden würde ich die Timestamps aber in Unix-Epoch in UTC
Zeitzone abspeichern und dann nur noch die Zeitzonen zum herumrechnen benutzen.
Die Variante gefällt mir (gestern von [Mullet](https://twitter.com/mulletti)
erklärt bekommen) aber der Zug is jetzt definitiv schon abgefahren. Ums mit
seinen Worten zu sagen "Fehler, die man nur einmal macht." :)
