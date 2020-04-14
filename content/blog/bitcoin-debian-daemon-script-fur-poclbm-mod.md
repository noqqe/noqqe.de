---
aliases:
- /blog/2011/06/12/bitcoin-debian-daemon-script-fur-poclbm-mod
- /archives/1689
comments:
- author: Maxi
  content: "<p>\xD6hm, dir ist aber schon bekannt, dass CPUs seit 8 Jahren bei Nichtauslastung
    runtertakten und a) deutlich weniger Strom verbrauchen und b) gleichfalls weniger
    Elektrosmog produzieren? Passt nicht zu deiner Einleitung und der Folgerung \"Aus
    diesem Grund...\".</p>"
  date: '2011-06-12T16:49:27'
- author: noqqe
  content: <p>BitCoin Mining benutzt keine CPU sondern die GPU der Grafikkarte. Ist
    dir schon bekannt oder?</p><p>Meine Argumentation richtet sich auf die vorher
    genannten Punkte (separate Rechner, neue Grafikkarten, Euro Tausch) wogegen der
    Skriptansatz so deutlich sparsamer ist.</p>
  date: '2011-06-12T16:56:47'
- author: Maxi
  content: "<p>Naja, die n\xE4chstliegendste Art der Berechnung ist erst mal die CPU,
    dann denkt man an die GPU, aber stimmt schon, ich hab das Script nicht gelesen.
    Funktioniert das eigentlich \xFCber OpenCL oder GL und funktioniert es auch mit
    freien Treibern oder AMD-Karten, oder ben\xF6tigt es spezielle Nvidia-APIs wie
    CUDA?</p><p>Dennoch: So weit ich das verstehe, lastet es die GPU so stark aus
    wie m\xF6glich, und das geb\xFChrt einer Warnung bei den \xFCblichen Monster-Karten
    und ihrem Energiehunger.</p>"
  date: '2011-06-12T17:29:16'
- author: noqqe
  content: "<p>Ja es funktioniert zum Gl\xFCck mit freien (normalen :P) Treibern und
    OpenCL (pyopencl). </p><p>Ich habe auch keine besonders leistungsf\xE4hige Grafikkarte
    (8600 GS - f\xFCr Office Anwendungen ausgelegt), weshalb bei mir nur bis zu 900KH/s
    m\xF6glich sind. Aktuelle Grafikkarten bewegen sich da eher im 80MH/s Bereich.
    Was da im SLI Verbund so passiert fehlen mir genauere Informationen. </p><p>Klar,
    die GPU wird belastet, denn irgendwoher muss die Rechenkapazit\xE4t kommen. Die
    GPU ist allerdings auf solche Aufgaben eher ausgelegt als normale CPUs. Was die
    ben\xF6tigte Leistung angeht ist auch klar, dass im Ruhezustand weniger verbraucht
    wird als unter Auslastung. Ich k\xF6nnte aber auch nicht sagen das ich das jetzt
    merken w\xFCrde was meine GPU da treibt. Kein starkes r\xF6deln der L\xFCfter
    oder anderes was darauf hinweisen k\xF6nnte.</p><p>Man k\xF6nnte den Rechner auch
    einfach auslassen, right? :P </p><p>Wir sind uns doch denke ich einig das es wohl
    keine effzientere (weniger verschwenderische) M\xF6glichkeit gibt an einem BitCoin
    Mining Pool teilzunehmen ohne massiv extra Strom zu verbrauchen, oder?</p>"
  date: '2011-06-12T18:57:04'
- author: Ghostbuster
  content: "<p>Dir ist schon aufgefallen, da\xDF Dein Init-Script ohne Abh\xE4ngigkeit
    zu einer funktionierenden Netzwerkanbindung ist? :P</p>"
  date: '2011-06-14T17:07:46'
- author: noqqe
  content: <p>Ich verstehe nicht ganz ? Meinst du den Debian related Part? Oben in
    dem Skript? </p><p>Oder was?</p>
  date: '2011-06-14T18:33:13'
- author: Ghostwriter
  content: <p>Im Init Info Header?</p>
  date: '2011-06-14T18:51:57'
- author: Lonesome Walker
  content: <p>kriegt man das auch auf der cpu zum laufen?</p><p>weil mein server hat
    leider nur 4MB video-ram ;-)</p>
  date: '2012-01-07T05:04:01'
- author: noqqe
  content: "<p>@Walker: </p><p>Ja es \"gibt\" CPU Mining Software. Allerdings w\xFCrd
    ich dir empfehlen bisschen im Wiki zu st\xF6bern und dich dazu schlauzulesen.</p><p>Da
    wird dir relativ schnell klar, dass du das nicht machen willst. Auf deinem Server
    schon gleich 3 mal nicht.</p>"
  date: '2012-01-07T15:59:10'
- author: Lonesome Walker
  content: "<p>Oh doch :-)<br>bei entsprechender CPU und das Mining \xFCber einen
    Pool macht das durchaus Sinn ;-)<br>(Server l\xE4uft ohnehin 24h nur im Leerlauf,
    von daher...)</p><p>Habe durch das Wiki gew\xFChlt, sind aber nur Windows und
    GUI-Miner :-(</p>"
  date: '2012-01-07T19:39:57'
date: '2011-06-12T11:09:28'
tags:
- mining
- crypto
- bitcoin
- ubuntu
- debian
- bash
title: "BitCoin | Debian Daemon Script f\xFCr poclbm-mod"
---

Man muss Dinge über Hirnschäden von Menschen lesen die
[Nachts guten Gewissens neben einem 4 Grafikarten im SLI-Verbund schlafen](http://www.bitcoinminingaccidents.com/?p=196)
und zugleich die [Bedenken von sinnierenden Typen](http://tav.espians.com/why-bitcoin-will-fail-as-a-currency.html) die
an den Limitierungen von 21 Billionen maximal möglichen BitCoins zweifeln.
Umweltverschmutzung ist natürlich auch ein Thema. Klar. Gerechtfertigter
Weise.

{{< figure src="/uploads/2011/06/Tux-G2_bitcoin.png" >}}

Wenn Ihr mich fragt, schiesst die BitCoin Mining Gesellschaft am Ziel
vorbei. Separate Rechner betreiben schiesst am Ziel vorbei. Hunderte von
Euros für neue Grafikkarten ausgeben um 5 MegaHashes/s mehr rechnen zu
können schiesst am Ziel vorbei. Euro's für BitCoins bezahlen schiesst
sowieso am Ziel vorbei. Man kann BitCoins meines Erachtens auch benutzen
ohne Kopfstände zu machen. Immer wenn der Rechner sowieso gerade läuft, mit
der Hardware die man zur Verfügung hat.

Aus diesem Grund habe ich ein kleines Skript gebastelt. Einen
Start-Stop-Daemon für **/etc/init.d/**. Es startet automatisch wenn mein
Rechner hochfährt und hört auf wenn ich Ihn herunterfahre. Ganz einfach

``` bash
wget -O /etc/init.d/bitcoin https://gist.github.com/raw/1007794/bitcoin.sh
chmod +x /etc/init.d/bitcoin
update-rc.d bitcoin defaults
```

[https://gist.github.com/1007794](https://gist.github.com/1007794)

Ich möchte aber dazu sagen, dass ich dies Funktionstüchtigkeit des Skripts
nicht auf anderen Rechnern/Betriebssystemen getestet habe. Vor Benutzung
also bitte lesen, verstehen ggf. anpassen. Außerdem wird ein Account bei
einem Mining Pool benötigt und der Mining Client an sich
([poclbm-mod](https://en.bitcoin.it/wiki/Poclbm-mod)). Während ich diesen
Post geschrieben habe, hat mein BitCoin Mining Client entspannt auf meiner
Geforce 8600 GS mit nahezu niedlichen 950 KiloHashes/s vor sich hin
gemined.
