---
date: 2011-03-14T21:37:05+02:00
layout: post
slug: zre-live-game-mitschnitt-via-jquery-und-php
status: publish
comments: true
title: ZRE | Live Game Mitschnitt via jQuery und PHP
alias: /archives/1520
categories:
- Bash
- Coding
- Debian
- General
- Linux
- PHP
- PlanetenBlogger
- Ubuntu
- Web
tags:
- environment
- game
- javascript
- jquery
- Output
- PHP
- zombie
- zombies
- zre
---

Eine der Aufgaben, von denen ich wirklich wenig bis keine Ahnung habe, war die Umsetzung der Live Gamebox für ZRE. Die Anforderung hat mich ehrlich gesagt mal in die Nähe des Web2.0 Wahns getrieben, von dem ich mich sonst fernzuhalten versuche. Zumindest von Entwicklerseiten her.

{% img center /uploads/2011/03/4625_7f89_390.jpeg %}

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
