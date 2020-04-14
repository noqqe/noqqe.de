---
date: 2016-12-31T14:45:25+01:00
tags:
- 33c3
- blog
- comments
- isso
title: yaml Kommentare
---

Auf dem [33c3](https://events.ccc.de/congress/2016/wiki/Main_Page) hatte
ich kurz Lust etwas umzusetzen, dass ich schon länger geplant hatte.

Seit der Umstellung des Blog Themes benutze ich kein
[isso](https://github.com/posativ/isso) mehr. Um trotzdem Feedback (z.B.
per Mail) an posts anhängen zu können hab ich mir überlegt Kommentare im
Yaml des Blogposts abzulegen.

``` yaml
---
title: Als ich begann die Wikipedia zu editieren.
comments:
- author: Anonymous
  content: "Wurde ja inzwischen gelöscht schade D:"
  date: '2016-02-01T19:57:44.804276'
- author: Bob
  content: "Test.."
  date: '2016-02-01T19:57:44.804276'
---
```

So zumindest die Idee. Zuerst hab ich mal Dummie Daten in einen Post
gebastelt und den Templating Code in `hugo` geschrieben. Dazu einfach in
dem `partial` meinen Footer angepasst:

``` html
{{ if isset .Params "comments" }}
  <div class="comments">
    <h4>Comments ({{ .Params.comments | len }})</h4>
  {{ range .Params.comments }}
    <div class="comment">
      <div class="itemdetails"><b>{{ .author }}</b> on {{ .date }}</div>
      <div>{{ plainify .content }}</div>
      <br/>
    </div>
  {{ end }}
{{ end }}
```

Als das fertig war, hatte ich aber noch keine Kommentare zum ausprobieren.
Daher hab ich dann noch alle Kommentare von Isso (aus der sqlit3 Datenbank)
via [frontmatter](https://github.com/eyeseast/python-frontmatter) in die
Blogposts umgezogen. Wen das interessiert kann sich das Script als
[Gist](https://gist.github.com/noqqe/7635337eb55e10cf3af9992101a3e26d)
ansehen.

Das Ergebnis kann man zum Beispiel mal in diesem
[Post](https://noqqe.de/blog/2015/10/26/als-ich-begann-die-wikipedia-zu-editieren/) begutachten.

Ich war auch tatsächlich etwas überrascht, über 8 Jahre kamen **1028**
Kommentare zu **267** Posts zusammen.

In `git log --shortstat` sah das ganze dann so aus

{{< figure src="/uploads/2016/12/comments.png" >}}

Als kleines Bonus feature kann man jetzt theoretisch auch einfach Pull
Requests für Kommentare auf github öffnen. Wenn man das will. Primär brauch
ich aber einfach keine Kommentare mehr im Blog. Wollte aber den
Datenbestand einfach nicht aufgeben. Für die Menschen!
