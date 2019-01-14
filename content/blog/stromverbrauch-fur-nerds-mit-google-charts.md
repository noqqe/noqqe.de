---
comments:
- author: Higgi
  content: "<p>W\xFCrde ein Nerd es nicht mit gnuplot machen ;)<br>Spa\xDF beiseite.
    Ich w\xFCrde es glaube ich mit gnuplot angehen. Da bin ich unabh\xE4ngiger von
    einem Dienst wie Google. Ich wei\xDF allerdings gerade nicht ob Gnuplot auch Tortendiagramme
    und Balkendiagramme kann. Bisher hab ich nur ganz klassisch x/y manchmal x/y/z
    aufgetragen.</p>"
  date: '2012-06-17T18:10:14'
- author: noqqe
  content: "<p>Hehe, jop die Argumentation kann ich verstehen. Aber mit Gnuplot hab
    ich schon gearbeitet, mit Google Charts noch nicht. </p><p>Gnuplot kann Balkendiagramme
    auch... hab n Template daf\xFCr. Torte hab ich noch nicht probiert, mag ich aber
    eh nicht wirklich ... ^^</p>"
  date: '2012-06-18T04:05:35'
date: '2012-06-16T10:31:00'
tags:
- web
- google
- stats
- html
- charts
- js
- json
- api
- opensource
- debian
title: "Stromverbrauch f\xFCr Nerds mit Google Charts"
---

Seit nun ca. 8 Monaten wohne ich nicht mehr bei meinen Eltern. Wenn man
Zuhause auszieht kommen immer ein Haufen neuer Aufgaben auf einen zu. Unter
anderem auch den Verbrauch des Stroms regelmäßig abzulesen und zu notieren.

Ich wollte wenn's irgendwie geht vermeiden irgendein Excel File zu pflegen
oder das einfach nur auf Papier zu schreiben. Also habe ich den Weg für
Statistik Nerds via [Google Charts](https://developers.google.com/chart/)
gewählt.

Konkret schick ich per JavaScript meine Daten zu einer Google API und
daraus werden dann wunderschöne(!), interaktive(!) Graphen in den Browser
gezaubert.

Am Anfang gibts erstmal einen Konfigurationsabschnitt. Dort trage ich die
Werte ein definiere Variablen. Ich habe hier zu Demonstrationszwecken
relativ zufällige Zahlen eingefügt.

``` js
<script type="text/javascript" src="https://www.google.com/jsapi"></script>
<script type="text/javascript">
  // Zaehlerstaende
  var g31082011 = 131079;
  var g30092011 = 131259;
  var g31102011 = 131445;
  var g31112011 = 132014;
  var g31122011 = 132150;
  var g31012012 = 132386;
  var g29022012 = 132804;
  var g31032012 = 133305;
  var g30042012 = 133514;
  var g30052012 = 133811;
  // Werte bzw. Differenzen errechnen
  var d30092011 = g30092011 - g31082011;
  var d31102011 = g31102011 - g30092011;
  var d31112011 = g31112011 - g31102011;
  var d31122011 = g31122011 - g31112011;
  var d31012012 = g31012012 - g31122011;
  var d29022012 = g29022012 - g31012012;
  var d31032012 = g31032012 - g29022012;
  var d30042012 = g30042012 - g31032012;
  var d30052012 = g30052012 - g30042012;
  // Preis pro kWh
  var kwhpreis = 0.23;
</script>
```

Das alles fülle ich dann in die Data Arrays der Google Chart Funktion ein,
definiere eine `div id` und beschrifte die Spalten und den Graphen an sich.

``` js
<script type="text/javascript">
  google.load("visualization", "1", {packages:["corechart"]});
  google.setOnLoadCallback(drawChart);
  function drawChart() {
    var data = new google.visualization.DataTable();
    data.addColumn('string', 'Date');
    data.addColumn('number', 'kWh');
    data.addRows([
      ['30.09.2011',    d30092011],
      ['30.10.2011',    d31102011],
      ['31.11.2011',    d31112011],
      ['31.12.2011',    d31122011],
      ['31.01.2012',    d31012012],
      ['29.02.2012',    d29022012],
      ['31.03.2012',    d31032012],
      ['30.04.2012',    d30042012],
      ['30.05.2012',    d30052012],
    ]);
     var options = {
      width: 550, height: 300,
      title: 'Monatlicher Stromverbrauch in kWh',
      vAxis: {title: 'Datum'}
    };
 var chart = new google.visualization.BarChart(document.getElementById('barchart_div'));
 chart.draw(data, options);
 }
</script>
```

Wenn alles fertig befüllt ist wird die .js Funktion clientseitig ausgeführt
und somit eine Anfrage an die Google Charts API gestellt. Was diese dann
antwortet sieht man hier:

{{< figure src="/uploads/2012/06/c1.png" >}}

So oder so ähnlich befülle ich dann auch alle anderen Funktionen die nötig
sind um mir andere Graphen bauen zu lassen. Im Grunde ist es jedesmal nur
das Einfügen von Variablen in ein Array. Die eigentliche Google Magie in
den Funktionen kann man sich sehr einfach auf der [Google Page](https://google-developers.appspot.com/chart/interactive/docs/quick_start)
zusammen klicken.

{{< figure src="/uploads/2012/06/c2.png" >}}

Ein wunderbar nichts sagendes Tortendiagramm darf natürlich nicht fehlen.

{{< figure src="/uploads/2012/06/c3.png" >}}

Ich hab mir auch noch einen Graphen gebaut der mir gleich noch die
vorraussichtlichen Preise pro Monat errechnet.

{{< figure src="/uploads/2012/06/c4.png" >}}

[Hier gibts das File](/uploads/2012/06/strom.html), so wie ich es bei mir
lokal Pflege. Da kann man auch nochmal genau den Source ansehen, wie welche
Funktion befüllt werden muss.  Und gerade weil das so schön .js ist, kann
man das auch wunderbar in seinen Blog einbinden, cool mh?
