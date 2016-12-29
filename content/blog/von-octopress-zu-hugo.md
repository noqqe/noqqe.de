---
comments:
- author: Martin
  content: "Gef\xE4llt mir!"
  date: '2015-04-01T21:29:27.238207'
- author: tux.
  content: "Gesetzt den Fall, dass man Go genauso wenig kann wie Python: Was spricht
    f\xFCr Hugo statt Pelican?"
  date: '2015-04-01T22:11:30.401686'
- author: noqqe
  content: "Hab ich mir vor ein paar Monaten mal angsehen. Objektiv kann ich aber
    jetzt keine Vor-/Nachteile vorbringen. Irgendwie sahen mir Codebeispiele und beim
    Docs durchbl\xE4ttern Hugo sympathischer aus."
  date: '2015-04-01T22:30:39.169485'
- author: Alex
  content: "Noch einfacher / sch\xF6ner als Hugo ist IMHO Kirby (http://getkirby.com/)
    Kostet allerdings schlappe 15,- Euros..."
  date: '2015-04-04T17:23:28.206269'
- author: noqqe
  content: Aber das ist ja PHP O_o
  date: '2015-04-11T13:12:37.477396'
- author: dakira
  content: "Also PHP bashing is ja auch so'n bisschen 2005 oder? Da hat sich in den
    letzten 10 Jahren einiges getan. Muss ich mal dr\xFCber bloggen. ;-) (PS: reply
    notification gibt es nicht, oder?"
  date: '2015-04-12T18:58:00.843856'
date: '2015-04-01T18:55:43'
draft: false
tags:
- development
- web
- octopress
- python
- hugo
- blog
- pygments
- go
- ruby
- gem
title: Von Octopress zu Hugo
---

Ich war genervt. Oh mein Gott, wie war ich genervt. Von diesem ständigem
Ruby, Octopress, Gem Fuckup, Pygments Fuckup.

### Octopress

{{< figure src="/uploads/2015/04/octopress.png" >}}

* 2-3 Minuten Blog-Compiles für poplige 300 Posts.
* Man kann nie irgendwas updaten, da sonst der ganze Blog bricht.
* Theming ist scheisse
* Mit einem Ruby Gem eine Software bedienen die eigentlich Python ist.
* Abhängigkeiten
* Kryptische Fehlermeldungen
* SCSS und Konsorten
* Hauseigenes, nicht Markdown-konformes Auszeichnen von Images, Quotes,
  Codeblocks

Aber am Meisten von der Struktur. Man muss sich das vorstellen. Man will
einen Blog eröffnen, klont den Source des Bloggingframeworks, fängt an
innerhalb des Repos hinein zu committen. Ändert Themes, Configs,
fummelt Plugins mit Git-Submodulen hinein. Und versucht dabei noch Upstream
in das Repo mit den eigentlichen Posts hineinzumergen.

Das haben mittlerweile auch die Entwickler von Octopress erkannt und
Octopress 3 angekündigt, mit dem alles besser werden soll. Aber irgendwie
fühlt sich das an wie die neue
[Version](http://octopress.org/2015/01/15/octopress-3.0-is-coming/) des
Internet Explorer.

### Hugo

3 Jahre hab ich mir das jetzt angetan. Und jetzt? Alles schön. So wunderschön.
Nachdem mir [posativ](https://posativ.org) auch letztens nochmal [gohugo.io](https://gohugo.io)
empfohlen hat, hab ich mir die in [go](http://golang.org) geschriebene
Software mal angesehen. Und sie machen einfach alles richtig. Alles.

{{< figure src="/uploads/2015/04/hugo.png" >}}

Das bauen meines ganzen Blogs dauert nur noch 400ms, eine aktive
Community (mit [Discourse Installation](http://discuss.gohugo.io) \o/),
[Dokumentation](http://gohugo.io/overview/introduction/) ohne Ende (im Gegensatz zu 4 Seiten Octopress Doku). Keine
1000 Subkomponenten, im Blog-Repo ist nur das Theme, Config/Meta Daten und die
Posts selbst. Es wirkt im Moment wie ein schnelles, straight-forward,
monolithisches Stück Software.

Migrationsschritte:

* Neuen Hugo Blog erstellen
* Demo Theme in `themes/` herunterladen
* `config.toml` mit Meta Daten des Blogs befüllen
* Umbenennen und CSS nach eigenen wünschen anpassen
* Logo und Favicon einbauen
* `uploads/` und `_posts/` aus Octopress umziehen
* Date Format anpassen
* Codeblock Syntax von `{% codeblock %}` zu Standard Markdown ändern
* Quote Syntax von `{% blockquote %}` zu Standard Markdown `>` ändern
* Image Syntax von `{% img center %}` zu
  [figure](http://gohugo.io/extras/shortcodes/) ändern
* Default Template auf `YAML` ändern
* Pagination in das Theme patchen
* Pages in die Sidebar einbauen und aus Octopress umziehen
* JavaScript Client-Side Syntax Highlightning (highlightjs) einbauen
* CSS für Codeblocks
* Switch von GoogleFonts zu lokalen Fonts
* Archive Page erstellen (denke hierzu gibts nen separaten Blogpost)
* Comments ([isso](https://posativ.org/isso)) als `partial`

Gerade das gebaue der Archiv Seite und der Pagination hat mir gezeigt wie
unfassbar geil die Templating Sprache ist. Zum Beispiel sind `GROUPS` und
`WHERE` Clauses ohne Probleme möglich.

``` go
{{ range (where .Site.Pages "Type" "post").GroupByDate "2006"  }}
```

Würde mich jetzt jemand fragen, welche Blogging Engine ich empfehlen würde,
wäre es defintiv Hugo. Aber ich glaube das merkt man auch an meiner
Schreibweise.