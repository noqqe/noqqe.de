---
title: "On Hackernews"
date: 2013-12-29T13:50:00+02:00
comments: true
tags:
- OpenBSD
- osbn
- Stats
- acrylamid
- hackernews
- reddit
- twitter
- plotly
- usage
- graphs
- GET
- POST
- DAAS
- mail
- stats
- statistik
---

Eigentlich wollte ich nur Acrylamid ausprobieren. Ich evaluierte
neue Software wegen des ständigen Fuckups mit Octopress und RVM.

{{< figure src="/uploads/2013/12/daas.png" >}}

Ich lud also Acrylamid herunter, spielte damit herum, bastelte eine Demo Site
mit etwas Layout Änderungen. Content bestand anfangs aus ein paar random Texten und "/dev/null"
im Header.
Irgendwann formte sich der Inhalt aber zu etwas Rundem und ich dachte,
egal jetzt kauf ich mir die Domain zum entstandenen Text. [devnull-as-a-service.com](http://devnull-as-a-service.com).
Ein kleines Spass Projekt am Rande schadet ja keinem.

Als ich dann fertig war und es
[vertweetete](https://twitter.com/noqqe/status/395252504033624065), freute ich mich noch, dass die Site
von [@slith76](https://twitter.com/slith76) erwähnt wurde. Legte mich aber dann schlafen.

Als ich am nächsten Morgen aufwachte sah das alles irgendwie anders aus. Ein
Blick in Piwik in der S-Bahn zeigte mir über 20.000 Visit. O_o. Dann wurde es
etwas verrückt; was folgt ist eine stichpunktartige Zusammenfassung.

### Facts

* die Zugriffszahlen des 1. Monats

<iframe id="igraph" src="https://plot.ly/~noqqe/47/640/420/" width="640" height="420" seamless="seamless" scrolling="no"></iframe>

* 127.326 Visits
* 438.461 Page Views (Actions)
* 1.249.375 GET Requests (4.385 MB)
* 97.318 POST Requests

{{< figure src="/uploads/2013/12/HTTP_Requests.png" >}}

* ~2 Stunden #1 auf HackerNews ([Ranking](http://hnrankings.info/6637480/))
* Zugriffe aus 118 verschiedenen Ländern

{{< figure src="/uploads/2013/12/locations.png" >}}

### Reaktionen

Als Response zu der Site bekam ich insgesamt ~40 Emails unter anderem:

* Jemand schrieb eine [Wrapper Library](https://gist.github.com/ryancdotorg/7241987) für Linux via LD_PRELOAD
* 3 Job/Praktika Anfragen bzgl. der Carreers-Page
* Ergänzungsvorschläge bzgl. des HTTP Methodenverhaltens
* Kooperationsanfrage des Dienstes ZAAS (/dev/zero as a Service)
* 4 lustige Supportanfragen bzgl. des offerierten Serivces
* 3 (ironisch gemeinte?) Jobangebote irgendwo in China-World
* Requestete Features: Kerberos, SOAP Interface, 10GE Interface für die
  Enterprise Appliance, Sharding Support, "is it webscale?")
* anderer related Content via Mail wie

{{< figure src="/uploads/2013/12/devnullnummernschild.jpg" >}}

* paar Dankes und/oder "LOL" Mails
* *.gif bei [SecurityReactions](http://securityreactions.tumblr.com/post/66371497808/when-a-client-asks-if-they-should-consider-using)
* 16 Github Issues
* 10 Pull Requests

### Wohin hat sich die Site verbreitet?

Wie alles anfing lässt sich nur sehr schwer rekonstruieren. Nicht das ich es
nicht versucht hätte, aber über mehrere Tweets und andere Socialmedia Kanäle zu
vergraphen hätte viel Aufwand bedeutet.

* [HackerNews](https://news.ycombinator.com/item?id=6637480) 271 Points \o/
* [Reddit](http://www.reddit.com/search?q=site%3Adevnull-as-a-service.com&restrict_sr=off&sort=relevance&t=all)
  Ich bin nicht sonderlich Redditaffin, die Kommentare sind aber teils lesenswert. 785 Upvotes \o/
* [Twitter](https://twitter.com/search?q=devnull-as-a-service.com&src=typd&f=realtime)
  Insgesamt gab es 990 Tweets die einen Link zu DAAS enthielten. Die Twitter
  Suche dazu erheitert mich immer wieder :)
* [Heise Security](http://www.heise.de/security/meldung/lost-found-Was-von-der-Woche-uebrig-blieb-2037941.html)
  hat DAAS in einem Wochenrückblick erwähnt
* [Habrahabr](http://habrahabr.ru/post/200230/) Ist scheinbar eine lustige
  Russische Tech-News-Site von der _immens_ viele Menschen kamen. Scheinbar
  grosses Ding da.

* Top Referrer als Graph

{{< figure src="/uploads/2013/12/Top_Referrers.png" >}}

### Schlechte Entscheidungen die ich traf

{{< figure src="/uploads/2013/12/shocked.gif" >}}

Es ist eine grundsätzliche Frage wie sehr man ein neues Projekt over-engineert.
Hielt die Aufwand/Nutzen Relation für angemessen, über ein paar Dinge ärgere ich
mich aber trotzdem.

* Ich hätte das OpenFiles Limit in OpenBSD für nginx anpassen sollen. So kam es
  dazu dass bei Lastspitzen sporadisch 500er Fehler geworfen wurden.
* Flattr. Mal ein Projekt bei dem es sinnvoll gewesen wäre Flattr zu integrieren
* [RFC 863](http://tools.ietf.org/html/rfc863) Discard Service nicht gleich zu nennen.
  Habe nicht gezählt wie oft, wurde aber gefühlt 10.000 mal an mich
  herangetragen.

### Gute Entscheidungen die ich traf

{{< figure src="/uploads/2013/12/schneewittchen.gif" >}}

* Plain Files bzw. Static Content deployed. In PHP oder etwas ähnlichem
  serverseitig Generiertem wäre ich viel früher hart hingeflogen, als durch die
  OpenFiles Limitierung.
* Kein gif auf der Startseite einbinden. Gifs benutze sich sehr gerne in Dingen
  die ich ins Internet stelle. Das hier nicht zu tun war unter dem Aspekt meines limitiert zur Verfügung stehenden Traffics eine sehr gute Entscheidung
* Die maximale POST-Size in [nginx](http://nginx.org) nicht aufzubohren. Es
  stellte sich heraus das _echt_ viele Leute Daten POSTeten.
* Einrichtung einer separaten Mailadresse unter Contact, da dort nun auch schon
  wieder echt viel Spam ankommt.
* Source auf GitHub zu stellen

Rückwirkend muss ich sagen, es war mal eine Erfahrung wert :) Ich wundere mich
auch immernoch wie wenig Kritik oder "mhhh lahm" Comments kamen. Genauswenig
wie Kommentare über das schlechte HTML oder das schlechte Layout.
Was auch immer. Ich hatte Spass. A lot.
