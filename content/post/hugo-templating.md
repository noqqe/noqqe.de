---
categories:
- blog
- development
- osbn
date: 2015-04-11T15:09:53+02:00
tags:
- Hugo
- Templates
- Archive
- Tags
- Categories
- Templating
- Go
title: Hugo Templating und Themes
---

Ein bisschen Hands-On an meine neue Blogging Engine. Bevor ich den Blog
migrieren konnte fehlten noch einige Anpassungen am Theme.

{{< figure src="/uploads/2015/04/webhipster.jpg" >}}

Da keines der Themes mir Out-of-the-Box gefallen hat, musste ich das
[Hyde](https://github.com/spf13/hyde) Theme ziemlich aufbohren.

### Pagination

Pages. Alle Blogposts auf der Startseite laden ist echt ungünstig.
Pagination ist allerdings ziemlich gut dokumentiert.

``` html
  <div class="posts">

    {{ $paginator := .Paginate (where .Data.Pages "Type" "post") }}
    {{ range $paginator.Pages }}

    <div class="post">
      <h1><a href="{{ .Permalink }}">{{ .Title }}</a></h1>
      <span class="post-date">{{ .Date.Format "2006-01-02" }}</span>
      {{ .Content }}
    </div>

    {{ end }}
  </div>

  {{ partial "pagination.html" $paginator }}
```

Kurz gesagt, Variable `$paginator` mit Content befüllen und mit `range`
drüber integrieren. `Paginate` greift dabei auf die Variablen aus der Config
zurück.

``` yaml
$ cat config.yaml
paginate: 10
paginatePath: "page"
```

Danach kann man durch die Seiten pagen.

{{< figure src="/uploads/2015/04/Pagination.png" >}}

Das komplette Fils gibts auf
[github](https://github.com/noqqe/noqqe.de/blob/master/themes/noqqe/layouts/index.html)

### Archive Page

Damit man aber trotzdem nicht alle Seiten einzeln durchkramen muss, fand
ich 'ne Archivseite auf Blogs schon immer sinnvoll. Gerade wenn die Site an
sich statisch ist und es somit keine Suche gibt.

{{< figure src="/uploads/2015/04/archive.jpg" >}}

Was man im Endeffekt erreichen möchte ist also eine statische Seite
erzeugen, die kein Blogpost und aber zur Build-Zeit dynamischen Content
hat.

Pages können dazu konfiguriert werden einem bestimmten Typ zu entsprechen.

``` yaml
$ hugo new archives.md
$ vim content/archives.md
---
date: "2012-03-31T21:38:03+02:00"
draft: false
title: "Archives"
type: archives
url: "/archives/"
---
```

Für den `type` muss ein Layout gleichen Namens im Theme enthalten sein:

```
$ tree themes/noqqe/layouts/
themes/noqqe/layouts/
├── 404.html
├── _default
│   ├── list.html
│   └── single.html
├── archives
│   └── single.html
├── index.html
├── page
│   ├── list.html
│   └── single.html
├── partials
│   ├── head.html
│   ├── isso.html
│   ├── pagination.html
│   └── sidebar.html
└── post
    ├── list.html
    └── single.html
```

Hier erkennt man auch, dass ich neben `archives` auch noch `post` und
`page` als Types definiert habe.

Der Single View für das Archiv ist das einzige was benötigt wird.

``` html
{{ partial "head.html" . }}
<body class="theme-base-0b">

{{ partial "sidebar.html" . }}

  <div class="content container">
  <div class="post">
  <h1>{{ .Title }}</h1>
      {{ range (where .Site.Pages "Type" "post").GroupByDate "2006"  }}
        <h2>{{ .Key }}</h2>
        <ul>
          {{ range .Pages }}
          <li>
              <span>{{ .Date.Format "02.01" }}</span>
              <a href="{{ .Permalink }}">{{ .Title }}</a>
          </li>
            {{ end }}
        </ul>
      {{ end }}
  </div>
  </div>
</body>
</html>
```

Im Endeffekt ist die Syntax ziemlich ähnlich wie bei der Pagination. Nur
der `range`-Query ist etwas mehr advanced. Im Endeffekt `for`-Loop über
allen Content mit `WHERE` und `GROUP BY` Command.

{{< figure src="/uploads/2015/04/Archive.png" >}}

An dem "2006" nicht stören. Die Templating Engine funktioniert nicht mit
wilden `%Y-%m-%d` Commands, sondern einem Referenz-Datum. Schön sehen kann
man das auch
[hier](https://github.com/noqqe/noqqe.de/blob/master/themes/noqqe/layouts/post/single.html#L9)


### Tags und Categories

Was auch gefehlt hat war, dass Tags und Kategorien im Single-View der
einzelnen Posts angezeigt werden. In `themes/noqqe/layouts/post/single.html`:

``` html
 <div class="post">
 <h1>{{ .Title }}</h1>
   <span class="post-date">{{ .Date.Format "2006-01-02" }}</span>
   {{ .Content }}

   <span class="meta">Categories:
   {{ range .Params.categories}}
     <a href="/categories/{{ . | urlize }}">{{ . }}</a>
   {{ end }}
   </span>
   <br/>
   <span class="meta">Tags:
   {{ range .Params.tags }}
     <a href="/tags/{{ . | urlize }}">{{ . }}</a>
   {{ end }}
   </span>
 </div>
```

Genau die Funktion `urlize` macht dann noch Links aus den Tagnamen.

{{< figure src="/uploads/2015/04/tagsandcats.png" >}}

So schön <3 Macht Spaß
