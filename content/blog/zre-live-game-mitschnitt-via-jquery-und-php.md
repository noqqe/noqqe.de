---
aliases:
- /archives/1520
- /blog/2011/03/14/zre-live-game-mitschnitt-via-jquery-und-php
comments:
- author: Knorkebrot
  content: "<p>Netter Codingstil :P<br>Deine for-Schleife im PHP ist kaputt. Die Clientseite
    h\xE4tte ich anders gemacht, vor allem in so fern, dass ich kein Framework daf\xFCr
    eingebunden h\xE4tte. F\xFCr solche Requests gibt es in Javascript die XMLHttpRequest-Klasse:</p><p>var
    xhr = new XMLHttpRequest();<br>function update() {<br> xhr.open('GET', '/statics/gamequery.php?mode=part',
    false);<br> xhr.send();<br> document.getElementById('gamebox').innerHTML = xhr.response;<br>
    return true;<br>}<br>setInterval(update, 2000);</p><p>jQuery ist einfach nur ein
    dicker, fetter Overhead f\xFCr sowas, denke ich.</p><p>M\xDFG</p>"
  date: '2011-03-17T20:14:44'
- author: Knorkebrot
  content: "<p>Oh, sieht ja fein aus, gibt es sowas wie <code> f\xFCr Kommentare?<br>M\xDFG,
    bzw. ist HTML erlaubt?</code></p>"
  date: '2011-03-17T20:16:36'
- author: Knorkebrot
  content: "<p>Na das artet ja schon in Spam aus.<br>W\xE4re sicher sinnvoll unter
    das Kommentarfeld zu schreiben, dass HTML (in Teilen?) erlaubt ist.<br>Du darfst
    meine Beitr\xE4ge gerne reparieren, wenn du willst :)<br>M\xDFG, kaputtgespielt.</p>"
  date: '2011-03-17T20:22:02'
date: '2011-03-14T19:37:05'
tags:
- development
- web
- javascript
- environment
- jquery
- ubuntu
- zre
- blog
- game
- zombies
- linux
- output
- php
- zombie
- debian
- bash
title: ZRE | Live Game Mitschnitt via jQuery und PHP
---

Eine der Aufgaben, von denen ich wirklich wenig bis keine Ahnung habe, war die Umsetzung der Live Gamebox für ZRE. Die Anforderung hat mich ehrlich gesagt mal in die Nähe des Web2.0 Wahns getrieben, von dem ich mich sonst fernzuhalten versuche. Zumindest von Entwicklerseiten her.

{{< figure src="/uploads/2011/03/4625_7f89_390.jpeg" >}}

ZRE als Daemon (ja, der Teil kommt auch noch) hat die Eigenschaft im zufällig wiederkehrendem Rhythmus Output in ein Textfile zu produzieren. Dieses Textfile (game.txt) hat einen Platz im DocumentRoot des Webauftritts zombies.n0q.org (Ist das eigentilch das erste Mal, dass ich die URL erwähne? Ich glaube ja.). Diese möchte ich parsen und über asynchrone Aktualisierung in den Browser des Heimanwenders bringen.

Zuerst musste ich also den PHP Parser Teil schreiben, der mir die letzten 20 Zeilen, oder falls weniger eben diese aus meiner game.txt ausließt.


    $input = file($zreoutput);
    $resultArray = "";
    $index = 20;
    if (count($input) < 20)
    {
       $index = count($input);
    }
    for($i=count($input)-$index; $i
    {
       $resultArray .= $input[$i]."<br/>";
    }
    return $resultArray;


Als nächstes war es dann noch nötig diesen zurecht gebogenen Output dann noch selbst aktualisierend in meine Index Datei zu portieren. Dazu ist dann eine JavaScript nötig um die Funktion zu definieren.


    setInterval(
        function() {
          $('#gamebox').load('../statics/gamequery.php');
          }, 2000
    );


Jetzt mussten sowohl jQuery als auch das JavaScript Snippet in die Index Datei eingebunden werden. Über Kritik und Anregung freue ich mich natürlich wie immer.


    <!-- game informations -->
    <script src="js/jquery-min.js" type="text/javascript"></script>
    <script src="js/gamebox.js" type="text/javascript"></script>
    <!-- game results -->
    <div id="gamebox"></div>



An der Stelle hatte ich gleich mal tatkräftige Hilfe eines Arbeitskollegen ;) Danke an dieser Stelle.