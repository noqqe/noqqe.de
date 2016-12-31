---
aliases:
- /blog/2011/02/07/duckduckgo-suche-via-firefox-addressleiste-umstellen
- /archives/1475
comments:
- author: Linux, Laptops und Ubuntu
  content: "<p>Cooler Tip. Duckduckgo ist wirklich eine Wucht. <br>Die erste Suchmaschine
    die f\xFCr mich Google ersetzen kann.</p><p>F\xFCr Geeks ist es die bessere Suchmaschine.</p>"
  date: '2011-02-07T22:42:04'
- author: Jensel
  content: "<p>Wo ist der Vorteil gegen\xFCber K\xFCrzeln, die bei Firefox direkt
    unter \"Suchmachinen verwalten ...\" vergeben werden k\xF6nnen? Nur, dass man
    die DuckDuckGo K\xFCrzel in der Suchleiste statt der Addressleiste nutzt?</p>"
  date: '2011-02-07T23:22:32'
- author: noqqe
  content: "<p>@rockiger: yeah!<br>@Jensel: Das Plugin f\xFCr die Suchmaschinen \xE4ndert
    eben leider nicht die Standartsuche via Addressleiste. :/ Oder hab ich dich falsch
    verstanden?</p>"
  date: '2011-02-07T23:39:31'
- author: Stefan
  content: "<p>\xDCbrigens:<br>STRG+K = Suchleiste<br>STRK+L = URL-Bar (Awesome Bar)</p>"
  date: '2011-02-07T23:55:28'
- author: noqqe
  content: '<p>@Stefan: yeah! Nach den Shortcuts hab ich lange gesucht. Danke!</p>'
  date: '2011-02-08T00:09:27'
- author: Jensel
  content: "<p>Ich meinte die plugin-freie, standardm\xE4\xDFig eingebaute Funktion,
    einem Suchmaschineneintrag ein K\xFCrzel vergeben zu k\xF6nnen. </p><p>Beispiele:<br>*
    \"g wort\" um nach Wort bei Google zu suchen<br>* \"w ding\" um ding bei Wikipedia
    nachzuschlagen<br>* \"en begriff\" um das DE-EN W\xF6rterbuch von Leo nach Begriff
    zu fragen</p><p>Wohlgemerkt in der Addresszeile. Den Suchkasten brauch ich net.
    Wenn ich dich richtig verstanden hab, macht DuckDuckGo was \xE4hnliches, oder
    nicht? Daher die Frage, wo ist der Unterschied?</p>"
  date: '2011-02-08T00:46:14'
- author: noqqe
  content: "<p>@Jensel: Achso jetzt verstehe ich. Du hast Recht DDG macht etwas \xE4hnliches,
    allerdings im viel breiterem Spektrum. Es fragt auch nicht _nur_ andere Suchfunktionen,
    sondern l\xE4sst sich auch irgendwie nach Themen sortieren. Schau dir doch mal
    die Liste von den !Bangs an: <a href=\"http://duckduckgo.com/bang.html\" rel=\"nofollow\">http://duckduckgo.com/bang.htm...</a></p>"
  date: '2011-02-08T00:49:49'
- author: name
  content: "<p>Nach ein paar Wochen mit Dukgo bin ich wieder zur\xFCck zu Google,
    weil ich sowieso oft genug auf Google zur\xFCckgreifen musste, weil die Relevanz
    der Suchresultate bei DuckDuckGo (gegen\xFCber Google) einfach nicht gut genug
    war.</p>"
  date: '2011-02-08T01:22:49'
- author: Stefan
  content: "<p>Achso noch eins: wenn man mit STRK+K auf die Suchleiste wechselt, dann
    kann man mit STRG+Pfeil-Hoch/Runtertaste die Suchmaschine w\xE4hlen</p>"
  date: '2011-02-08T01:43:07'
- author: kdeuser
  content: <p>wow..... with wolfram alpha integration... and the search results are
    good!!!  wow... that is a really hot candiadate to replace google.</p>
  date: '2011-02-08T02:19:57'
- author: grk
  content: "<p>Noch ein Tipp:<br>mit einem ! im Suchstring z.B. \"ubuntuusers !\"
    wird eine auf-gut-gl\xFCck suche ausgef\xFChrt. auch manchmal sinnvoll.</p>"
  date: '2011-02-08T02:35:27'
- author: DarkMetatron
  content: "<p>&gt; \"<a href=\"http://duckduckgo.com/?q=\" rel=\"nofollow\">http://duckduckgo.com/?q=</a>\"
    <br>&gt; als String einf\xFCgen</p><p>Also ich pers\xF6nlich nutze ja lieber den
    verschl\xFCsselten Zugang zu duckduckgo unter \"<a href=\"https://duckduckgo.com/?q=\"
    rel=\"nofollow\">https://duckduckgo.com/?q=</a>\"</p>"
  date: '2011-02-08T13:35:22'
- author: Matthias
  content: '<p>Im neuen Firefox 23 ist das jetzt umgestellt: Ab sofort suchst du immer
    mit der Suchmaschine, die in der Standartsuchleiste gerade verwendest und keyword.URL
    ist verschwunden.</p>'
  date: '2013-08-16T01:20:11'
date: '2011-02-07T16:04:06'
tags:
- development
- web
- shell
- duckduckgo
- privacy
- addressleisten suche
- searchengine
- geek
- bang
- sicherheit
- suche
- bang search
- alternative
- jabber
- nerd
title: DuckDuckGo | Suche via Firefox Addressleiste umstellen
---

Seit einiger Zeit benutze ich die Suchmaschine [duckduckgo.com](http://duckduckgo.com ).
Wenn man von dem völlig bescheuerten Namen und
dem Logo mal absieht, versteckt sich dahinter eine wunderbare Suchmaschine.
Dazu gehört alles was der typische Paranoia Geek [so braucht](http://duckduckgo.com/goodies.html).

{{< figure src="/uploads/2011/02/nduck.v105-300x52.png" >}}

Um jetzt zum Punkt zu kommen: Meine Standardsuchleiste habe ich via Plugin
schon auf Duckduckgo umgestellt. Ich habe aber die blöde Angewohnheit meine
Suchanfrage einfach direkt in die Suchleiste des Firefox zu schreiben.
Bisher wurde mir das immer brav von Google beantwortet. Ich dachte mit der
Installation des Plugins von DuckDuckGo wäre das auch umgestellt. Dem ist
aber nicht so. Leider.

Stattdessen gilt es, noch folgende Anpassungen zu machen:

  * about:config in die Adresszeile tippen
  * Nach dem Begriff "keyword.URL" suchen
  * Das Feld doppelklicken (standardmäßig dürfte es leer sein)
  * "**http://duckduckgo.com/?q=**" als String einfügen

Anschließend Firefox Neustart. Nein ich habe kein Geld für diesen Post
bekommen. Schon wieder nicht.