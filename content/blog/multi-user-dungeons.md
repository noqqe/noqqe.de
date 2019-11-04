---
comments:
- author: Anonymous
  content: "Toll, dass du nochmal an sowas erinnerst! Mir geht es bei Spielen um die
    erlebte Geschichte, und daf\xFCr braucht man keine Grafik."
  date: '2015-09-12T16:25:10.916720'
- author: chris_blues
  content: "Hihi! Lang ists her. Hab auch vor ein paar Monaten mal wieder eins von
    den bsd-games angeworfen - wei\xDF nicht mehr wie's hie\xDF. Da war meine Geduld
    allerdings ziemlich schnell ersch\xF6pft. Bin halt doch kein Zocker mehr... Aber
    ich hab noch allerhand Erinnerungen an n\xE4chtelanges shell-zocken mit Star-Trek
    & Co... :D"
  date: '2015-09-12T17:05:06.724260'
- author: nosports
  content: 'Herzlich Willkommen im MG. :)

    (SirJow gruesst)'
  date: '2015-09-14T09:26:23.409261'
date: '2015-09-12T12:47:14'
tags:
- multi user dungeon
- web
- shell
- network
- mud
- games
- opensource
- rpg
title: Multi User Dungeons
---

[MUDs](https://de.wikipedia.org/wiki/Multi_User_Dungeon) sind textbasierte RPGs
aus der ganz frühen Zeit der Videospiele.  Man verbindet sich ganz altmodisch
per `telnet` auf Port 23 und es kann los gehen.

Ich habe das vor Jahren mal ausprobiert, aber gleich wieder verworfen. Vor ein
paar Wochen musste ich aber wieder daran denken und angefangen das
[MorgenGrauen](http://mg.mud.de) zu spielen.

{{< figure src="/uploads/2015/09/hafen.png" >}}

Typischerweise befindet man sich in Räumen bzw Feldern, die mit einem 3-4 Zeilen
langem Beschreibungstext die Umgebung spezifizieren. Das Reicht vom Hobbit Wald
zur Drachenhöhle, über den Sumpf der Orks, bis hin zur Zwergenhochburg im
Gebirge. Darin können sich allerhand Sachen befinden. NPCs, Tische, Tümpel, ein
Strauch, ein abgetrennter Trollkopf oder verzauberte Kochlöffel. Danach hat man meist ein paar
Möglichkeiten. Man untersucht Details im Raum, spricht mit NPCs, tötet
irgendetwas oder verlässt den Raum in eine Himmelsrichtung.

{{< figure src="/uploads/2015/09//Colossal_Cave_Adventure_on_VT100_terminal.jpg" >}}

Man bewegt sich also von Raum zu Raum. Am Anfang geht das noch und man fängt
Stück für Stück an die Welt zu erforschen aber schnell merkt man, dass die Welt
einfach zu groß ist um durch fiese Wälder zu tollen und sich alle Wege zu
merken. Deshalb macht man einfach ein paar Skizzen und Notizen.

{{< figure src="/uploads/2015/09/map2.JPG" >}}

Charm hat auch Stufenpunkte System. Statt World of Warcraft like pausenlos auf
dem immer gleichen Mob zu kloppen um EPs zu grinden, setzt sich das Levelsystem
aus einem komplexen Mix allerhand Indikatoren zusammen. Natürlich gibt es EP,
aber es gibt auch separate Punkte für Quests, gefundene Zaubertränke (die Stats
erhöhen), geführte Dialoge, Besuche in Orten der Weltkarte und und und.
Alles ist darauf abgezielt die Erfahrung des Spielers in allen Bereichen des
Spiels im Level widerzuspiegeln.

### TinTin++

Mit der Zeit ergeben sich aber natürlich Schwächen von einfachem Telnet. Kein
Zurück, keine History, Text des flappt einfach zwischen dem Getippten usw.
Dabei bietet so ein einfaches Textinterface do so viel potiential für Tooling!

* Autocomplete
* History
* Alias
* Loops
* Highlights für Farben
* Actions
* Reverse Search a la Bash
* Path Recording (was ich noch nicht nutze)
* Prompts

Es gibt wirklich eine Menge Clients. Das Meiste davon ist totaler Bullshit.
Einer aber hat meinen Geschmack recht gut getroffen. Ncurses Interface und sau
schön skriptbar. [Tintin++](http://tintin.sourceforge.net). Er kann so ziemlich
alles was man sich wünschen kann.

Ein Auszug der Konfiguration von Tintin++

```
#HIGHLIGHT  {%w teilt Dir mit: %*}  {light green}  {5}
#HIGHLIGHT  {Du teilst %w mit: %*}  {light green}  {5}
#HIGHLIGHT  {Es gibt %w sichtbaren Ausgang:}  {light blue}  {5}
#HIGHLIGHT  {Es gibt %w sichtbare Ausgaenge:}  {light blue}  {5}
#HIGHLIGHT  {^[%w:%*]%*}  {green}  {5}
#SPLIT      {bottom row}
#TAB        {steinschlag}
#TAB        {unt}
#TAB        {info}
#TAB        {frage}
#TAB        {kurzinfo}
#ALIAS      {k} {kurzinfo} {5}
#ALIAS      {u} {unt} {5}
#ALIAS      {tmo} {tm ootato} {5}
#ALIAS      {ZT} {zaubertraenke} {5}
#ACTION     {Eine Stechmuecke kommt an} {toete muecke} {5}
#SESSION    MorgenGrauen mg.mud.de 23
```

Un noch einiges mehr. Wenn ich zum Beispiel in einer alten Scheune befinde und
50 Kisten zum Untersuchen habe, löse ich das auch öfters mal mit einer For-Loop

    #loop 1 50 x {ziehe kiste $x ; oeffne kiste}

Oder ähnliches. Es macht Spass im Moment einen kleinen Zwerg durch die Gegend
schicken. Ich werde da wohl noch die eine oder andere Stunde verbringen :)

Bildquelle: VT-100 von
[Wikipedia](https://de.wikipedia.org/wiki/Datei:Colossal_Cave_Adventure_on_VT100_terminal.jpg)
