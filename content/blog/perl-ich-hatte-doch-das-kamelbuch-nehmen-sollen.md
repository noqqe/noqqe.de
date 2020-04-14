---
aliases:
- /blog/2010/07/04/perl-ich-hatte-doch-das-kamelbuch-nehmen-sollen
- /archives/1087
comments:
- author: otzenpunk
  content: "<p>Die Perl-B\xFCcher von O'Reilly sind wirklich hervorragend. Allerdings
    ist das \"Kamel-Buch\" eher eine umfassende Referenz der Sprache, als ein Einsteigerbuch.
    Von den O'Reilly-Perl-B\xFCchern, die ich besitze, war es mir eigentlich mit am
    wenigsten von Nutzen. \"Einf\xFChrung in Perl\" war da weit fl\xFCssiger zu lesen.
    Solltest du dir noch ein Perl-Buch zulegen wollen, empfehle ich von O'Reilly auf
    jeden Fall das \"Perl Kochbuch\" und/oder \"Einf\xFChrung in Perl-Objekte, Referenzen
    &amp; Module\". (Und nat\xFCrlich \"Regul\xE4re Ausdr\xFCcke\". Das ist zwar nicht
    perl-spezifisch, aber Perl ohne RegExes ist nur eine halbe Sprache.)</p>"
  date: '2010-07-08T14:41:48'
- author: noqqe
  content: "<p>Hi otzenpunk,</p><p>na das h\xF6rt sich doch mal an als w\xFCsstes
    du von was du sprichst ;) </p><p>Regul\xE4re Ausdr\xFCcke hab ich schon durch.
    Hat auch wirklich sehr geholfen :)</p><p>Ich weiss nur nicht ob das Einf\xFChrung
    in Perl jetzt nicht schon zu... naja. \xDCberfl\xFCssig ist, da ich das andere
    Buch und Tutorial des Perlboards schon gelesen habe.</p><p>oder was meinst du?</p>"
  date: '2010-07-09T00:02:23'
- author: otzenpunk
  content: "<p>&gt; Ich weiss nur nicht ob das Einf\xFChrung in Perl jetzt nicht schon
    zu\u2026 naja. \xDCberfl\xFCssig ist</p><p>Richtig. So hatte ich das auch gemeint.
    Ich w\xFCrde mit dem Kochbuch fortfahren.</p>"
  date: '2010-07-12T21:22:22'
- author: noqqe
  content: <p>Ah, okay :) <br>Ich denke das werde ich auch :)</p>
  date: '2010-07-16T19:53:06'
date: '2010-07-04T07:17:00'
tags:
- development
- programmieren
- use perl
- perl
- linux
- oreilly
- coden
title: "Perl | Ich h\xE4tte doch das Kamelbuch nehmen sollen"
---

...denn das wäre mir jetzt viel lieber, als das ["Programmieren in Perl" von Rainer Krienke vom Hanser Verlag](http://www.amazon.de/Programmieren-Perl-Rainer-Krienke/dp/3446220135/ref=pd_sim_b_17).
Fehler in den gedruckten Skripten wie fehlende Klammernpaare oder ähnliches
können einen Perl-Anfänger leicht zum verzweifeln bringen. Ganz abgesehen
von den irgendwie distanziert klingenden Erklärungen des Buchs. Ich sollte
aber vom Anfang anfangen. Nicht mittendrin. Hab gehört das macht man so in
Blogs.

{{< figure src="/uploads/2010/07/perl_logo.gif" >}}

Perl hat angefangen mich zu interessieren. Oder anders rum? Vor ... ca.
einem viertel Jahr. Ich las teilweise kurz How-To's und auch mal längere
Beschreibungen, weil mir manchmal die Schwerfälligkeit und nicht vorhandene
Flexibilität von Bash auf den Keks ging.

Das mittlerweile als etwas veraltete Skriptsprache für Administratoren
gehandelte Perl hat mich neugierig gemacht. Die Module, die Handlichkeit im
Umgang mit Strings, Hashwerten und RegExp gefielen mir. Beim
herumexperimentieren damit bin ich nun auf den [PerlGuide vom deutschen Perlboard](http://www.perlboard.de/perlguide/Inhalt.html) gestossen. Ein
wunderschönes Tutorial. Jedes Kapitel endet mit Übungen, Praxis-Programmen
und Zusatzfragen.

``` perl
print "Perl hat manchmal geile Syntax" and die if ! defined @ARGV ;
```

Danach, fand ich brauchte ich noch ein Buch. Also ein analoges Medium aus dem
ich auch mal auf der Couch lesen kann. Meine Wahl fiel nicht (wie meistens)
auf ein [O'Reilly Buch](http://www.amazon.de/Programmieren-mit-Perl-Larry-Wall/dp/3897211440/ref=sr_1_5?ie=UTF8&s=books&qid=1278231152&sr=8-5).
Sondern auf das bereits oben erwähnte Programmieren in Perl. Nachher erfuhr
ich (unter anderem aus dem wunderbaren [1. O'Reilly Podcast](http://community.oreilly.de/blog/2010/06/25/kol001-das-oreilly-universum/)
mit [Tim Pritlove](http://tim.geekheim.de/)) dass das Perl-Buch von O'Reilly
eines der Besten sein soll, die dort je herausgebracht wurden.

Bereue meinen Kauf. Allerdings mache ich trotzdem Fortschritte mit Perl.
Wie immer arbeite ich mit den Skripten in einem Git-Repo. Wen'S
intressiert:
[git.zwetschge.org](http://git.zwetschge.org/?p=learning-perl.git;a=tree;h=671b98e403d952d9ed2730ac1221e867039127cc;hb=671b98e403d952d9ed2730ac1221e867039127cc)
