---
layout: post
title: "Das CAP Theorem"
date: 2012-04-29T12:22:00+02:00
comments: true
categories:
- Web
- Datenbanken
- Code
keywords: "Berkeley, CAP, Theory, Theorem, Databases, Datenbanken, Knoten,
wikipedia"
---
Das CAP Theorem ist mir das erste mal wirklich begegnet im Podcast
[Binärgewitter #1
NoSQL](http://www.radiotux.de/index.php?/archives/5497-Binaergewitter-1-NoSQL.html)
in dem es kurz und knackig an einfachen Beispielen gut erklärt wurde.

{{< figure src="/uploads/2012/04/cap-theorem.png" >}}

Im Grunde hat sich schon kurz nach der Jahrtausendwende ein Professor in
Berkeley Gedanken gemacht was man von einem Datenbank System erwarten kann.
Im wesentlichen gehts um 3 Faktoren, die aber niemals alle zugleich erfüllt
werden können.

{% blockquote wikipedia %}
Konsistenz (C): Alle Knoten sehen zur selben Zeit dieselben Daten.
Verfügbarkeit (A): Alle Anfragen an das System werden stets beantwortet.
Partitionstoleranz (P): Das System arbeitet auch bei Verlust von Nachrichten, einzelner Netzknoten oder Partition des Netzes weiter.
{% endblockquote %}

Will man sich mal anschauen, wenn man Datenbanken betreut.

Links

* [CAP Theorem - Wikipedia Deutsch](http://de.wikipedia.org/wiki/CAP-Theorem)
* [CAP Theory - Wikipedia Englisch](http://en.wikipedia.org/wiki/CAP-Theory)

