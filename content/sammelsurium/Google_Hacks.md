---
title: Google Hacks
date: 2012-01-06T12:58:40.000000
tags: 
- Software
---


#### Einfache Suchausdrücke

In Google Suchabfragen werden die eingegebenen Wörter automatisch mit einem
logischen "UND" miteinander verknüpft. Das bedeutet, dass Google nach allen
eingegebenen Wörtern sucht. Die Groß- bzw. Kleinschreibung der eingegebenen
Suchbegriffe wird von Google nicht beachtet. Somit sind die folgenden
beiden Suchabfragen identisch:

    Haus AND STAHL AND BeToN
    haus stahl beton

Um nach einem der eingegebenen Suchbegriffe zu suchen, wird der logische
Operator "OR" verwendet. Alternativ kann an Stelle von "OR" auch das
Pipe-Symbol ("|") verwendet werden. Die Suchabfrage

    haus OR beton OR stahl

findet Seiten in denen entweder das Wort Haus, Beton oder Stahl vorkommt.

Suchabfragen können mehrere unterschiedliche Operatoren enthalten. Weiters
können Begriffe mit runden Klammern auch gruppiert werden. Um nach dem Wort
Haus, sowie nach Beton oder Stahl zu suchen, kann die Suchabfrage

    haus (beton | stahl)

bzw.

    haus (beton OR stahl)

verwendet werden.

Google behandelt die eingegebenen Worte alle einzeln und es ist egal, in
welcher Reihenfolge die eingegebenen Suchbegriffe im Dokument vorkommen.
Will man nach einem ganz bestimmten Ausdruck suchen, muss dieser Ausdruck
in Anführungszeichen eingeschlossen werden. Die Suchabfrage

    "unix time stamp"

findet Dokumente, in denen genau diese Wortfolge enthalten ist.

Oft werden bei einer Suche auch viele Seiten gefunden, die nichts mit dem
gesuchten Thema zu tun haben. Um Google anzuweisen, dass Seiten mit einem
bestimmten Wort nicht gefunden werden dürfen, muss vor diesem Wort ein
Minuszeichen gesetzt werden. Die folgende Suchabfrage findet Seiten, in
denen "unix time stamp" vorkommt aber das Wort "system" nicht enthalten
ist:

    "unix time stamp" -system

Möchte man zwar nach einer bestimmten Wortkombination suchen, kennt aber
ein oder mehrere Wörter innerhalb des Suchbegriffes nicht, kann ein
Stern-Zeichen (*) als Platzhalter verwendet werden. Der Suchbegriff

    "alles * feinsten"

findet z.B. "Alles vom Feinsten", aber auch "Hier stimmte einfach alles,
NATURLAUB vom Feinsten!".

Möchte man nach einem oder mehreren Sonder- oder Satzzeichen zwischen
Buchstaben suchen, kann ein Punkt als Platzhalter verwendet werden:

    vicky.21

findet unter anderem "Vicky (21) ", "Vicky, 21" oder "Vicky :: 21".
Spezielle Suchausdrücke

Mit Google ist es auch möglich, in bestimmten Teilen einer Webseite zu
suchen. Dies kann beispielsweise der Titel, der eigentliche Text, oder ein
Link sein. Dazu müssen vor dem Suchausdruck bestimmte Schlüsselwörter
angegeben werden, welche mit einem Doppelpunkt vom Suchausdruck getrennt
werden. Zwischen dem Schlüsselwort bzw. Doppelpunkt und dem Suchbegriff
darf kein Leerzeichen enthalten sein, außer wenn diese in Anführungszeichen
eingeschlossen werden.

#### intitle:

Durchsucht nur die Titel der Webseiten. Dabei wird jedoch nur der Begriff
beachtet, der direkt an "intitle" folgt. Der Suchausdruck

    intitle:"unix time" stamp

sucht "unix time" im Titel der Seiten und "stamp" irgendwo auf der Seite.

#### allintitle:

Sucht alle angegebenen Wörter nur im Titel einer Webseite. Die Suche

    allintitle:"unix time" stamp

findet den Ausdruck "unix time" sowie "stamp", wenn sie im Titel einer
Seite vorkommen. "allintitle" sollte nicht mit anderen Spezialausdrücken
verwendet werden.

#### inurl:

Durchsucht nur die URLs der Seiten nach dem angegebenen Begriff. Die
Funktionsweise ist analog zu der von "intitle".

#### allinurl:

Sucht nur in den URLs nach den darauffolgenden Wörtern. Die Funktionsweise
ist gleich jener von "allintitle". Auch "allinurl" sollte nicht mit anderen
Spezialausdrücken verwendet werden.

#### intext:

Durchsucht nur den Seiteninhalt nach dem angegebenen Begriff. Nach Links,
in der URL und dem Titel der Seite wird dabei nicht durchsucht.

#### allintext:

Durchsucht nur den Seiteninhalt nach den darauffolgenden Begriffen.
"allintext" sollte nicht mit anderen Spezialausdrücken verwendet werden.

#### inanchor:
Mit diesem Ausdruck können nur die Linkbeschreibungen in den Seiten
durchsucht werden. Lautet der HTML-Quelltext eines Links

    <a href="http://www.gaijin.at/">Gaijin's Homepage</a>

wird nur im Text "Gaijin's Homepage" gesucht.

#### site:

Mit dem Ausdruck "site:" wird nur nach dem angegebenen Host, der Domain
oder der Top-Level Domain gesucht werden. Hier einige Beispiele für diesen
Suchausdruck:

    site:www.gaijin.at
    site:gaijin.at
    site:at

#### link:

Gibt alle Seiten aus, die auf die angegebene Seite verweisen.

    link:www.gaijin.at

findet beispielsweise alle Seiten, die einen Link zu http://www.gaijin.at/
enthalten. Das Präfix "http://" kann jedoch weggelassen werden. Die Suche
mit "link:" ist jedoch nicht sehr genau.

#### cache:

Suche eine Seite im Cache von Google. Die Anfrage

    cache:www.gaijin.at

gibt die zuletzt gespeicherte Version der Seite aus. Vor allem wenn die
Seite vor kurzem komplett geändert oder entfernt wurde, kann dieser
Suchausdruck hilfreich sein.

### daterange:

Beschränkt die Suche auf einen bestimmten Zeitraum, in dem die Seite
indiziert wurde. Da Google ein julianisches Datum erwartet, ist diese
Funktion eher unhandlich..

#### filetype:

Mit diesem Suchausdruck ist es möglich, nur nach bestimmten Dateitypen zu
suchen. Nach "filetype:" wird die gewünschte Dateinamenerweiterung
angegeben. So sucht

    filetype:doc

beispielsweise nur Worddokumente. Der Suchbegriff

    "google hacks" filetype:pdf

sucht nur nach PDF-Dokumente, die "google hacks" enthalten.

#### related:

Der Suchausdruck "related" sucht nach Seiten, die mit der angegebenen Seite
inhaltlich in Zusammenhang stehen. Die Suche nach

    related:google.com

liefert zum Beispiel eine Liste mit diversen Suchmaschinen.

#### info:

Mit dem Ausdruck "info" zeigt Google zur angegebenen Seite Links zu
weiteren Informationen, wie etwa einen Link zur Seite im Cache und weitere
nützliche Abfragen an. Der Ausdruck

    info:www.gaijin.at
