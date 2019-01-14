---
aliases:
- /blog/2010/08/29/explain-shell-kommandos-visualisiert-erklaren
- /archives/1224
comments:
- author: noqqe
  content: <p>Ich kritisiere meinen Blogpost an dieser Stelle selbst und frage mich,
    ob es nun ein Markdown oder ein Markup File bzw. Mechanismus ist, den ich oben
    beschreibe.</p>
  date: '2010-08-30T01:11:50'
- author: vain
  content: "<p>Also, ich tendiere da eher zu \u201Eup\u201C. :)</p><p>Cheers!</p>"
  date: '2010-08-30T14:26:15'
- author: Eddard
  content: "<p>Kurze verst\xE4ndnissfrage. was ist der unterschied bzw. worum geht
    es bei Markup/down</p>"
  date: '2010-08-30T23:40:07'
- author: noqqe
  content: "<p>Markup: Aus Daten einer bestimmten Form werden neue Daten generiert,
    welche das Endprodukt darstellen. Sozusagen Rohdaten-&gt;Ergebnis</p><p>Markdown:
    sch\xE4tze ich umgekehrt.</p><p>Soweit ich das richtig verstanden habe. Lasse
    mich aber gerne korrigieren. :)</p>"
  date: '2010-08-30T23:58:07'
- author: vain
  content: "<p>Hmmm, ich kenne den Begriff \u201EMarkup\u201C als Kurzform f\xFCr
    \u201EMarkup Language\u201C oder eben eine \u201EAuszeichnungssprache\u201C, mit
    der du dem Inhalt eines Dokuments Eigenschaften hinzuf\xFCgen kannst. Also Beispiel
    LaTeX: Mit \u201Etextbf{Hallo}\u201C sagst du aus, dass das Wort \u201EHallo\u201C
    bitte fett gedruckt werden soll. Wei\xDF jetzt nicht, ob das ein hartes Kriterium
    ist, aber f\xFCr mich geh\xF6rt zu einer Markup-Sprache auch immer dazu, dass
    die Dokumente mit einem einfachen Texteditor bearbeitbar sind.</p><p>\u201EMarkdown\u201C
    kenne ich jetzt nur als Namen einer konkreten Markup-Sprache. Ob die das nun aus
    Spa\xDF an der Freude \u201Edown\u201C statt \u201Eup\u201C genannt haben oder
    ob es da einen tieferen Sinn gibt, wei\xDF ich nicht. ;) Vielleicht liegt es daran,
    dass Markdown einige Dinge implizit aus der Plain-Text-Formatierung \xFCbernimmt,
    Abs\xE4tze oder Listen zum Beispiel. Es ist also n\xE4her am Text (\u201Edown\u201C,
    weiter unten, auf keiner so hohen Abstraktionsebene) als zum Beispiel HTML \u2013
    oder so? Keine Ahnung, reine Spekulation. ;)</p><p>Das, was beim Explain-Skript
    als Eingabe verlangt wird, ist in gewissem Sinne eine ganz, ganz einfache Markup-Sprache.
    Der Inhalt deines \u201EDokuments\u201C ist der zu erkl\xE4rende Befehl samt den
    Kommentaren. Und das versiehst du durch einfache Mittel mit weiteren Informationen:
    Markierung des Bereiches f\xFCr jeden Kommentar und Trennung der einzelnen Kommentare
    mit einer Leerzeile. Das alles kannst du auch ohne weitere Hilfsmittel in einem
    einfachen Texteditor machen.</p><p>W\xE4re jetzt zumindest meine Interpretation
    der Dinge. :)</p>"
  date: '2010-08-31T02:27:44'
- author: noqqe
  content: "<p>Der Preis f\xFCr den l\xE4ngsten Kommentar im Blog geht \xFCbrigens
    an vain ;)</p><p>Sch\xF6n erkl\xE4rt. Jetzt hab ich's auch verstanden :)</p>"
  date: '2010-08-31T11:23:18'
date: '2010-08-29T20:31:32'
tags:
- development
- python script
- dokumentation
- explain
- find
- shell
- png
- ubuntu
- vain
- debian
- bash
title: "Explain | Shell-Kommandos visualisiert erkl\xE4ren"
---

Neulich bin ich über das [github Profil](http://github.com/vain) von [Peter Hofmann](http://uninformativ.de)
gestolpert. Darin befand sich ein Projekt,
welches ich sehr interessant fand.

[Explain](http://github.com/vain/explain) versucht Shell Kommandos zu
erklären und zu visualisieren. Gerade für Blogs oder andere Dokumentationen
finde ich das mehr als sinnvoll. Es erstellt aus einem simpel gestricktem
Markdown File eine ASCII-Art ähnliche Erläuterung des Kommandos.
Beispielsweise:

```
$ ./explain.py command.markdown
find . -iname '*.png' -exec echo '<br><img src="{}">' ; > gallery.html
\_/ | ___________/  ________/ ___________________/ |  ___________/
  |  |       |             |               |           |        |
  |  |       |             |               |           |        - Ausgeben nach
  |  |       |             |               |           |           gallery.html
  |  |       |             |               |           |
  |  |       |             |               |           - find Syntax Ende.
  |  |       |             |               |
  |  |       |             |               - mit folgendem Inhalt aus.
  |  |       |             |
  |  |       |             - und führe echo
  |  |       |
  |  |       - alle Dateien die mit .png enden
  |  |
  |  - im aktuellen Verzeichnis
  |
  - Finde (via find)
```

(PlainText: [/uploads/2009/09/015](/uploads/2009/09/015))

Die Syntax des Files das zur Deklaration der Ausgabe dient:

```
find . -iname '*.png' -exec echo '<br><img src="{}">' ; > gallery.html
---- ! -------------  ---------- --------------------- ! -------------

Finde (via find)

im aktuellen Verzeichnis

alle Dateien die mit .png enden

und führe echo

mit folgendem Inhalt aus.

find Syntax Ende.

Ausgeben nach gallery.html
```

Die Trennzeichen  sind via Parameter austauschbar und auch ansonsten tut
das kleine Python Script seinen Job hervorragend. Sollte demnächst mal
wieder ein Kommando erläutert werden müssen, werde ich definitiv darauf
zurückgreifen. Weitere Beispiele auch unter:

[1] [http://www.uninformativ.de/?section=news&ndo=single&newsid=118](http://www.uninformativ.de/?section=news&ndo=single&newsid=118)
[2] [http://github.com/vain/explain](http://github.com/vain/explain)
