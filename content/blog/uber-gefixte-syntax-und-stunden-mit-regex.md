---
comments:
- author: dakira
  content: "<p>Hey. Kannst du mal ein wenig erz\xE4hlen, wie du zu Octopress gekommen
    bist?</p>"
  date: '2012-11-23T10:59:39'
- author: noqqe
  content: "<p>Klar, dazu gibts nen Blogpost: <a href=\"http://noqqe.de/blog/2012/03/05/switched-to-octopress/\"
    rel=\"nofollow\">http://noqqe.de/blog/2012/03/0...</a>. Ham uns im \xF6rtlichen
    Hackerspace bissl \xFCber Blogging Software unterhalten und dann kam (zum Gro\xDFteil
    wegen Markdown und wegfallen von PHP) eins zum anderen :)</p><p>Wenn du konkrete
    Fragen hast meld dich ;)</p>"
  date: '2012-11-23T16:22:16'
- author: dakira
  content: "<p>Dort wird die Motivation auch nicht erkl\xE4rt. Meine Frage ist: Warum
    nicht mehr WP? Wieso ist PHP weggefallen, oder willst du nur sagen, dass octopress
    auf ruby setzt. Was heisst \"wegen Markdown\"? Schreibt man in octo die Artikel
    mit Markdown? Also kurz: Was war die Motivation?</p>"
  date: '2012-11-23T16:33:35'
- author: noqqe
  content: "<p>Naja:</p><p>* Das ich kein PHP mehr brauche hat mir Security Technischund
    Performance Technisch  gefallen. Die statischen HTML Files die ich per rsync einfach
    hochladen kann, find ich total nice. </p><p>* Markdown: Ich denke dabei ist der
    Workflow ganz wichtig. Das Wordpress \xFCber seine Admin GUI \"h\xE4ndeln\" ist
    mir schon immer schwer gefallen. Die beiden Editoren bedienen. Es hat sich immer
    wie ein Kampf angef\xFChlt. Wenn ich jetzt einen Post schreibe sitze ich vor meinem
    Vi mit Synatx Highlightning f\xFCr die (wie ich finde) wundersch\xF6ne Auszeichnugssprache
    Markdown und setze danach einen git commit ab. </p><p>Au\xDFerdem war WP einfach
    viel zu gro\xDF f\xFCr das was ich damit mache. Ich wollte was simples, minimalistisches.
    Damit f\xFChl ich mich jetzt wohler.</p>"
  date: '2012-11-24T07:56:36'
- author: Dr. Azrael Tod
  content: "<p>nicht zu vergessen wie ekelhaft vieles bei Wordpress ist...<br>Ich
    hab mir vor l\xE4ngerer Zeit mal den Aufwand gemacht und versucht die Datenbank
    in ein eigenes Projekt zu konvertieren. Dinge wie permanente, absolute URLs in
    dem HTML-Strings die dieser M\xFCll direkt in die Datenbank wirft sind noch die
    ertr\xE4glichsten Sachen.<br>Quellcode-M\xE4\xDFig sieht es noch schlimmer aus.</p>"
  date: '2012-11-26T10:59:48'
- author: bo
  content: "<p>Statt des ersten sed kannst du auch `fmt` benutzen, m\xFCsste man sich
    aber noch etwas einfallen lassen, wie man damit umgeht, wenn man h1 und h2 mit
    `=` und `-` unterstreicht, die zieht er dann n\xE4mlich in eine Zeile zusammen.
    K\xF6nnte man nat\xFCrlich auch nachtr\xE4glich wieder hinbiegen, aber ich f\xE4nde
    das deutlich sch\xF6ner, wenn das von Anfang an nicht passiert. Das Problem gibt
    es nicht beim Atx-Style mit Hashes.</p><p>Btw., die inline Bilder als {% img ...
    %} sind nicht Standard-Markdown,  die andere Variante ![alt](url) ist es aber
    schon. Hat also 'nen Grund. :)</p>"
  date: '2012-11-27T09:24:12'
- author: Martin
  content: "<p>Ich sekundiere den letzten Abschnitt. Diese nicht standardkonformen
    Extension in Jekyll sind eine Pest. Wenn sowohl das Bild als auch die URL identisch
    sind, kann man das einfach referenced machen: [![alt][ref]][ref]. Nicht sonderlich
    intuitiv, aber daf\xFCr mit einem alternativen Text, der seit HTML5 Pflicht ist
    ;-)</p>"
  date: '2012-11-30T09:31:50'
- author: noqqe
  content: <p>Pflicht bei html5? Wusst ich nicht O_o </p><p>Sagen wirs so, sollte
    "Markdown - The Spec" jemals erscheinen, werd ich meine Syntax daran ausrichten.
    Ging ja auf HN genug rum die Diskussion ;)</p>
  date: '2012-11-30T13:27:30'
- author: posativ
  content: <p>Ja, die Markown Specs sind sehr unbefriedigend. Aber auf Jekyll zu setzen
    macht es schwieriger, zu anderen blog engines zu wechseln (z.B. wenn Du die lange
    Kompilierungszeit von Octopress/Jekyll leid bist :p)</p>
  date: '2012-11-30T21:18:31'
- author: noqqe
  content: "<p>Naja zuletzt ging es ja drum einmal _einen_ Spec f\xFCr Markdown zu
    definieren. Schau dir mal bei Wikipedia die Liste der Implementierungen [1] an.
    Und jede davon funktioniert ein bisschen anders. </p><p>Sollte sich _ein_ Standard
    durchsetzen und tats\xE4chlich zum RFC werden m\xFCssen eh einige leute was tun.
    </p><p>Zum Thema Kompilierungszeit: Wie lange braucht dein Framework eigentlich
    so wenn du deinen Blog baust? H\xE4tt mich eh mal interessiert :) </p><p>[1]:
    <a href=\"http://en.wikipedia.org/wiki/List_of_Markdown_implementations\" rel=\"nofollow\">http://en.wikipedia.org/wiki/L...</a></p>"
  date: '2012-12-01T09:48:50'
- author: posativ
  content: "<p>Pro \xC4nderung/neuer Eintrag ~1s, from scratch ca. 10s, bei 170 Beitr\xE4gen.
    Wie lange braucht dein Blog mit Octopress?</p>"
  date: '2012-12-01T12:08:36'
- author: noqqe
  content: <p>ah. ok. Bei mir sinds zwischen 20 und 30 Sekunden from scratch bei 279
    Posts. </p><p>(gerade selber zum ersten mal wc -l gemacht. wow 280 schon O_o)</p>
  date: '2012-12-01T12:45:14'
- author: posativ
  content: "<p>Kann Octopress/Jekyll inzwischen was anderes als from scratch?</p><p>280
    sind schon ordentlich viel. Gratulation daf\xFCr!</p>"
  date: '2012-12-01T13:23:18'
- author: noqqe
  content: "<p>Ich weiss nicht ob ich jetzt ein bisschen Ironie raush\xF6re bei deinen
    \"280\" :D </p><p>Jo du kannst</p><p>rake isolate['post-name.markdown'] </p><p>machen.
    Dann werden alle anderen Posts weggeschoben und du kannst das f\xFCr </p><p>rake
    preview </p><p>sch\xF6n zum testen nutzen. Bevor du Uploadest musst du aber dann
    trotzdem </p><p>rake gen_deploy machen. (In dem Fall macht gen_deploy dann: integrate,
    generate, deploy) f\xFCr dich.</p>"
  date: '2012-12-01T13:35:09'
- author: noqqe
  content: <p>whatever. wollte deine Engine aber auch mal ausprobieren. :)</p>
  date: '2012-12-01T13:35:50'
- author: Martin
  content: "<p>Nein, das war keine Ironie, sondern echte Anerkennung! :-)</p><p>Das
    mit dem `rake isolate [/path/]` kannte ich gar nicht. Nat\xFCrlich eine effektive
    Umgehung der kompletten Rekompilierung. Dann ist Octopress ja gar nicht mal so
    \xFCbel wie ich dachte (abseits der Jekyll-only Syntax :P)</p>"
  date: '2012-12-02T17:33:13'
- author: Martin
  content: "<p>Freue mich \xFCber jeden Nutzer, ist jetzt aber vom Theme her nicht
    so h\xFCbsch wie Octopress und bietet auch nicht diese hunderttausend one-click-and-go
    Features. Vielleicht sieht man sich ja mal in #acrylamid auf Freenode.</p>"
  date: '2012-12-02T17:34:45'
date: '2012-11-22T20:27:00'
tags:
- development
- web
- markdown
- octopress
- regex
- blog
- html
- bash
title: "\xDCber gefixte Syntax und Stunden mit RegEx"
---

Manchmal sitze ich abends einfach so herum und repariere kaputte/unschöne
Syntax alter (via [exitwp](https://github.com/thomasf/exitwp)
[importierter](/blog/2012/03/05/switched-to-octopress/)) Blogposts unter
Zuhilfenahme regulärer Ausdrücke damit besagte Eintraege nacher genauso
aussehen wie vorher. Das hilft keinem. Nur mir, wenn ich mir die Markdown
Sourcen des Octopress Repos anschaue.

{{< figure src="/uploads/2012/11/bale.jpg" >}}

## Umbrüche

Einen Großteil der Zeit habe ich mir überlegt wie ich es schaffe nach 80
Zeichen einen automatischen Umbruch mit `sed` zu machen ohne dabei ein Wort
zu zerhacken. Lösung:

``` bash
sed -i -e 's#\(.\{80\}\S*\s\)#\1\n#g' *
```

## Ugly syntax is ugly

Folgendes ist eine sehr schlimme Variante von Bildeinbettung in Markdown
(imho). Exitwp ist da aber anscheinend anderer Meinung und hat damals alle
Bilder so eingebunden:

```
Original:
[![CC by Drunken Monkey](/uploads/2010/10/243104896_eb10db6e1d.jpg)](/uploads/2010/10/243104896_eb10db6e1d.jpg)

Wünschenswert:
&#123;&#37; img center /uploads/2010/10/243104896_eb10db6e1d.jpg &#37;&#125;
```

Aber auch das hab ich dann noch hinbekommen.

```
sed -i -e 's#.*\[![^\(]*(\([^\)]*\).*#&#123;&#37; img center \1 &#37;&#125;#g' *
```

## Misc

Zusätzlich noch die Domain zwetschge.org komplett aus den Posts gekratzt,
Alte removed und kaputte Syntax aus exitwp generierten Posts.

```
$ git commit -a -m "fixed every broken markdown img tag with epic sed line"
[master b0b1693] fixed every broken markdown img tag with epic sed line
 78 files changed, 137 insertions(+), 137 deletions(-)
```

Außerdem hab ich die Hälfte des Blogposts schon bei Twitter gespoilert.
Komisch das ich meinen Account da wieder reaktiviert hab.