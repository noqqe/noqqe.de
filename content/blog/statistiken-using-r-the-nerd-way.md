---
aliases:
- /blog/2011/02/04/statistiken-using-r-the-nerd-way
- /archives/1458
- /blog/2011/02/04/statistiken-using-r-the-nerd-way/
- /blog/2011/02/04/statistiken--using-r-the-nerd-way/
comments:
- author: Fredo
  content: "<p>Yeah, R rocks!</p><p>R ist wirklich eine ziemlich interessante Sprache.
    Die Grafikm\xF6glichkeiten sind ziemlich gut, aber auch die F\xE4higkeiten in
    Bezug auf Datentransformation, Matrixalgebra etc. sind sehr ausgefeilt.</p><p>Nur
    als Hinweis am Rande: Kreisdiagramme sind zwar ganz h\xFCbsch, aber Balkendiagramme
    sind in der Regel zweckm\xE4\xDFiger: Bei Kreisdiagrammen l\xE4sst sich die Gr\xF6\xDFe
    von Segmenten optisch relativ schwer erfassen und vergleichen, w\xE4hrend Balkendiagramme
    sehr gut gelesen und interpretiert werden k\xF6nnen.</p>"
  date: '2011-02-04T21:45:02'
- author: noqqe
  content: "<p>hey fredo!<br>Ja, fand ich auch. Hab mir das Manual schon zum gro\xDFen
    Teil durchgelesen, aber irgendwie keine wirkliche Verwendung f\xFCr die (echt
    umfangreichen) mathematischen M\xF6glichkeiten gefunden. </p><p>Ja, bei den Kreisdiagrammen
    muss ich dir recht geben, aber ich meine hey... KUCHEN! :)</p>"
  date: '2011-02-05T00:40:39'
- author: HmpfCBR
  content: "<p>Schau dir auch mal 'littler' an, gerade f\xFCr Interaktion bash und
    R recht hilfreich.</p><p>Ansosnten kann ich Frodo bzgl. Kreis und Balkendiagrammen
    nur zustimmen. Wenigstens hast du keine 3D-Torten benutzt. ;) Vergleich deine
    Grafiken und frag dich selbst, wo du was wie gut erkennst ...</p><p>Torten taugen
    nur bei ganz wenigen Teilen und auch dann nur, wenn man klar machen will, dass
    es sich um Teile eines Ganzen handelt.</p>"
  date: '2011-02-05T04:07:23'
- author: thomas
  content: "<p>Wow, R ist echt nice :D</p><p>Vielleicht k\xF6nnte man bei \"Commands\"
    und auch der Statistik deiner Timeline statt den Rest \"abzuschneiden\" auch R
    anweisen, nur die 10 gr\xF6\xDFten Anteile einzeln anzuzeigen und den Rest zusammengefasst
    darzustellen, eben als \"Rest\" wie man das von herk\xF6mmlichen Statistiken kennt.<br>Damit
    w\xFCrdest du das Ergebnis nicht verf\xE4lschen und trotzdem die \xDCbersicht
    bewahren.</p><p>Super Artikel, wieder was gelernt!</p>"
  date: '2011-02-07T01:19:10'
- author: noqqe
  content: "<p>@HmpfCBR: Ja, wenn ich so im nachhinein dr\xFCber nachdenke, habt ihr
    wohl Recht. </p><p>@thomas: Ja ich weiss was du meinst. Sone art \"Other\"-Spalte
    im Kuchen wie man das von Wahlen kennt. Hab ich mich auch kurz schlaugelesen,
    aber dazu h\xE4tte ich (wenn ich das richtig verstanden hab) ein zus\xE4tzliches
    Paket aus dem CRAN gebraucht. Damit hab ich mich dann f\xFCrn Anfang doch noch
    nicht auseinander setzen wollen ;) Aber sch\xF6n mal wieder was von dir zu h\xF6ren!</p>"
  date: '2011-02-07T18:09:21'
date: '2011-02-04T10:18:35'
tags:
- development
- web
- analyiseren
- shell
- twitter
- linux
- r-script
- blog
- statistiken
- r
- ubuntu
- mail
- maillog
- statistics
- bash
- history
title: Statistiken | Using R - The Nerd Way
---

Momentan interessiere ich mich für [R](
http://de.wikipedia.org/wiki/R_%28Programmiersprache%29). Die
Programmiersprache für Statistiken genießt aktuell ziemliche
[Popularität](http://www.nytimes.com/2009/01/07/technology/business-computing/07program.html).

Persönlich habe ich (in der kurzen Zeit, seitdem ich es benutze und kenne)
R als eine Sprache kennengelernt, die nur begrenzt für coole Sachen benutzt
wird. Kann an meiner Deklaration von "cool" liegen oder das sich die coolen
Sachen bisher meiner Kenntnis entzogen. Aber mal davon abgesehen dass
Statistiken immer interessant sind, wird R anscheinend meistens von
irgendwelchen Sales-Druiden oder Marketing-Dudes verwendet, die Daten aus
der DB ihres Online-Shops kratzen lassen um festzustellen wie viele Rosa
Einhörner letzten Monat verkauft worden sind.

Nunja, sagen wirs so. Ich hab überlegt für was ich R nutzen würde und etwas
damit gespielt. Prinzipiell ist die vorgehensweise sehr einfach. Ich hab
irgendwo ein LogFile o.ä. rumliegen (sozusagen die Rohdaten) und bringe
diese mit ein bisschen Bash/Shell Zauber in ein Format welches ich R
vorschmeissen kann. R unterstützt... naja so ziemlich alles, was einen
halbwegs anständigen Delimiter vorweisen kann.

R arbeitet in diesem Fall mit einer Interaktiven Shell, die die Eingaben interpretiert.

### Tweetstatistik meiner Timeline analysieren (via twidge)

```
# Letzte 200 Tweets ausgeben
$ twidge lsrecent --all > tweets
# Usernamen sortieren und Tweets zählen
$ grep '^<' tweets | awk '{print $1}' | sort | uniq -c | sort -rn
# In R Shell wechseln
$ R
# Twitter Timeline einlesen
> twittertimeline <- read.table("twitter/tweets", head=T, sep="" )
# Kuchenansicht erstellen
> pie(twittertimeline$Tweets, label=twittertimeline$User)
```

{{< figure src="/uploads/2011/02/timeline-500.png" >}}

An dieser Stelle kann es passieren, dass es bei zu vielen Daten wirklich
sehr unübersichtlich wird. Was genau ich da falsch mache ist mir allerdings
etwas schleierhaft. An der Skalierung lässt sich sicher noch arbeiten. Aber
mit einem kleineren Satz von Daten, sieht es schon besser aus.

```
# 20 meisten Twitterer heraussuchen
$ grep '^<' tweets | awk '{print $1}' | sort | uniq -c | sort -n  | tail -20 > tweets-20
# Einlesen
> twittertimeline20 <- read.table("twitter/tweets-20", head=T, sep="")
# Balkendiagram erstellen
> barplot(as.matrix(twittertimeline20$Tweets), main="Timeline", ylab="Tweets", xlab="Users", beside=TRUE, col=rainbow(20))
```

Und damit das ganze auch was aussagt, kann man sogar eine kleine Legende anfügen.

```
> legend( 1, 60,twittertimeline20$User, cex=0.6,  fill=rainbow(20))
```

{{< figure src="/uploads/2011/02/timeline-legend-500.png" >}}

### History der Bash visualisieren

Wer auf Kuchen steht, kann auch eifnach mal seine Bash History einlesen.

``` bash
# Command sortieren und zählen
$ history | awk '{print $4}' | sort | uniq -c | sort -rn | head > sys/history
   Used Command
   1106 ls
    889 vim
    490 cd
    423 git
    116 screen
    116 rm
     81 ./tismc.bash
     76 ps
     75 less
     72 ./coming-home.bash
# Das generierte File in R einlesen
> history <- read.table("sys/history", head=T, sep="")
# Aus dem Object history ein neues Object für Labels erstellen
# in denen die prozentualen Werte der Commands stehen
> history_labels <- paste(history$Command , " " , round(history$Used/sum(history$Used) * 100, 1) , "%", sep="")
# Object und Labels zu einem Kuchendiagramm formen
> pie(history$Used, main="Commands", label=history_labels, cex=0.8)
```

{{< figure src="/uploads/2011/02/commands-500.png" >}}

### Mailadressen aus mail.log

Um noch ein 3. mal den Anwendungsbereich zu
wechseln, kann man damit auch wunderbar das mail.log seines Mailservers
unter die Lupe nehmen.

    > mailaddresses <- read.table("mail/mail_addresses", head=T, sep="")
    > pie(mailaddresses$Mails, main="Mail Adressen", col=rainbow(length(mailaddresses$Address) ), label=mailaddresses$Address)

{{< figure src="/uploads/2011/02/addresses-advanced2-500.png" >}}

Egal, mein BWL Lehrer sagte immer: Traue nie einer Statistik, die du nicht
selbst gefälscht hast.
