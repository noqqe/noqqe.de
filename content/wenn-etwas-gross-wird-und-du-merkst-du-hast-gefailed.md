---
layout: post
title: "Wenn etwas groß wird und du merkst du hast gefailed."
date: 2013-02-15T13:15:00+02:00
comments: true
categories:
- php
- google charts
- mysql
- fail
- Debian
- ubuntuusers
- osbn
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

{% img center /uploads/2013/02/ohgod.gif %}

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

{% blockquote Koushik Nandy, http://blog.codez.in/php5-time-zone-solution/web-development/2010/07/09 blog.codez.in %}
Time zone settings are normally not the first thing you think of when working on a web project with PHP
{% endblockquote %}

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
