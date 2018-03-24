---
title: "10 Jahre Blog"
date: 2018-03-24T10:30:49+01:00
tags:
- blog
---

Kaum zu glauben das [mein erster Blogpost](/blog/2008/03/24/hallo-welt-2/) heute genau
10 Jahre her ist. Aus diesem Anlass heute ein kurzes Recap. Im Lauf der Zeit
hat dieser Blog einige Veränderungen durch.

Anfangs habe ich einfach auf [wordpress.com](https://wordpress.com) einen Blog
unter [seufz.wordpress.com](http://seufz.wordpress.com) erstellt und angefangen
Gedanken zu meinen Problemen die ich mit Linux auf meinem Laptop hatte.  Viele
davon handeln von Ubuntu Problemen, Grub und sonstigem Kram. Die typischen
Einsteiger Probleme von Linux Usern damals™ zumindest. Ich muss gestehen das
einige davon so peinlich sind, dass ich sie vor ein paar Jahren offline nahm.

{{< figure src="/uploads/2018/03/seufzbanner.gif" caption="Banner von 2008" >}}

Irgendwann hatte ich dann einen eigenen Server. Anfangs eine VM mit Debian Etch
auf als Xen [domU](https://wiki.xen.org/wiki/DomU), später als eigene Hardware. Dort lief dann auch irgendwann
meine [eigene Instanz](https://noqqe.de/blog/2009/02/23/noqqede-relaunch-des-blogs/) und ich
registrierte noqqe.de.  So langsam ging es dann auch los mit Server
Administration (was auch daran lag, dass ich endlich in die Technik-Abteilung
meines damaligen Arbeitgebers wechseln durfte). Ein
[Squid](https://noqqe.de/blog/2008/08/30/squid-opensource-proxyserver/), ein
[Apache2](https://noqqe.de/blog/2008/09/10/mein-erster-apache2/) und ein
Postfix. Warum ich davon immer die ganze Config im Blog gepastet hatte ist mir
aber ein Rätsel. Wahrscheinlich war ich stolz auf das Gefrickel.


{{< figure src="/uploads/2018/03/noqqewhite.png" caption="Banner von ca. 2010 ">}}

Wie man das als anständiger  Blogger so macht, musste natürlich auch an
Wordpress herumgepfuscht werden. So hat dann auch der
[eine](https://noqqe.de/blog/2010/04/05/wordpress-archive-page-erstellen/)
oder [andere](https://noqqe.de/blog/2009/12/17/wordpress-rss-und-planet/)
Artikel über [Wordpress](https://noqqe.de/tags/wordpress/) hier Einzug
gefunden. Es war furchtbar wenn ich so darüber nachdenke. Ungefähr zu der Zeit
wurde mein Blog auch Teil diverser Blog Planeten/Netzwerke
([ubuntuusers.de](https://ubuntuusers.de),
[planetenblogger.de](http://planetenblogger.de), [osbn.de](https://osbn.de),
usw..) von den übrigens nichts übrig geblieben ist.


{{< figure src="/uploads/2018/03/cloudtheme.png" caption="Eigenes Wordpress Theme von 2009" >}}

{{< figure src="/uploads/2018/03/airnoqqe.jpg" caption="Banner von ca. 2011" >}}


Als ich dann den Arbeitgeber wechselte, in die Nähe von Erlangen zog und anfing
mich im [K4CG](https://k4cg.org) herumzutreiben, migrierte  ich den Blog zu
[octopress](https://octopress.org). Statische Blog-Generatoren waren 2012
schliesslich der heiße Scheiss ([Switched to Octopress
- noqqe.de](https://noqqe.de/blog/2012/03/05/switched-to-octopress/))! Auch
sonst entfernte sich der Blog immer mehr von Linux Desktop Anwender hin zu
Server Systemen und insbesondere [OpenBSD](https://noqqe.de/tags/openbsd/).

{{< figure src="/uploads/2018/03/octopress.png" caption="Octopress Screenshot" >}}

In diesem Schema hat sich der Blog dann eingependelt. Und naja, letztes Jahr
ist hier etwas weniger passiert, was an anderen Interessensgebieten liegt.
Insgesamt hab ich mir mal kurz die Mühe gemacht ein bisschen Stats aus den 10
Jahren zusammenzutragen.

Blog Posts

* 322 Posts
* 26.902 Zeilen
* 128.355 Wörter
* 988.677 Zeichen

Sammelsurium Posts

* 175 Artikel
* 11.920 Zeilen
* 38.400 Wörter
* 306.910 Zeichen

Btw, auch das nochmal schön weil ich das mit `wc`

```
wc content/blog/*
ls content/blog/* | wc
```

einfach generieren konnte. PlainText <3 Und hier noch ein paar Besucher Stats,
zumindest solange ich Piwik in Verwendung habe..

{{< figure src="/uploads/2018/03/screenshot.png" caption="Besucher Stats" >}}

Es ist selten das man mal ein Projekt wirklich so lange durchzieht. Zumindest
bei mir. Die Intention “Linux Fortschritt dokumentieren” hat sich (wenn mans
nicht ganz so genau nimmt, wegen nebenher BSD) eigentlich seit 10 Jahren nicht
verändert. Blogging ist eine schöne Sache und ich muss sagen das ich beim
Zusammensuchen der Screenshots und dem Lesen alter Posts fast etwas
nostalgisch wurde. Fast ein bisschen wie ein Tagebuch.

Die einzige Sachen die ich beim Bloggen für wirklich wichtig halte ist die
folgende: Sucht euch Software von der man leicht wieder weg migrieren kann.
Denn das wird man sicher eines Tages.
