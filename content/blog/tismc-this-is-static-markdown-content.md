---
aliases:
- /archives/1243
- /blog/2010/09/19/tismc-this-is-static-markdown-content
comments:
- author: Holger
  content: "<p>Ich \xE4\xE4h nixe Disain\xE4r. Aber FLEISCHk\xE4se okay .D</p>"
  date: '2010-09-21T20:23:05'
date: '2010-09-19T14:09:00'
tags:
- development
- web
- cronjob
- markdown
- md
- tismc
- content
- html
- static
- debian
- bash
title: tismc | This is static markdown content!
---

Vorher noch nie gehört? Kann auch gar nicht sein.

Das sich gerade formierende Hobby-Eishockey-Team, dem ich angehöre, war
bisher ohne großartige HomePage ausgekommen. Es bestanden kleine Basteleien
unseren Captains, aber so wirklich vollständig war das dann auch nicht. Ich
machte mich auf die Suche nach einem minimalen CMS. Irgendwas kleines
Schönes, in dem wir unseren minimalen Content unterbringen können. Ausmaße
wie etwa [Mytinytodo](http://mytinytodo.net) für Todo-Listen (von dem ich
sehr begeistert bin ;) ). Ich wurde nicht wirklich fündig. Alle Tipps die
ich bekam waren zu umfangreich und somit einfach Overkill.

Die Sache schlief irgendwie 1-2 Monate ein, bis ich von der
[Markdown](http://daringfireball.net/projects/markdown/)-Language gehört
habe. Diese Mischung aus .txt-File Syntax und HTML-Generator gefiel mir und
ich dachte wieder an die Hockey-Page. Muss es immer PHP und SQL sein?
Eigentlich nicht.

Der Plan war: In einen Unterordner content/ für jede Seite ein
*.markdown-File zu laden, aus welchem generisch die .html Files gebaut
werden. Umgeben Natürlich von header.tmpl und footer.tmpl. Auf dem
Filesystem sieht das in etwa so aus:

```
    |-- content
    |   |-- 001.impressum.markdown
    |   |-- 001.index.markdown
    |   `-- 002.index.markdown
    |-- css
    |   |-- readme.txt
    |   `-- style.css
    |-- impressum.html
    |-- index.html
    |-- tismc.bash
    |-- tmpl
          |-- header.tmpl
          `-- footer.tmpl
```

Ein kleines Bashscript durchsucht den content/-Ordner und baut für jedes
gefundene .markdown-File die HTML-Page.

``` bash
$ ./tismc.bash
page-dir is now /path/to/homepage
--- parsing impressum.html
adding header
parsing content/001.impressum.markdown
adding footer
--- parsing index.html
adding header
parsing content/002.index.markdown
parsing content/001.index.markdown
adding footer
```

So kann auch mein Captain oder andere Dritte, neue Inhalte in die Page
hochladen. Denn Markdown-Syntax ist nicht wirklich
[kompliziert](http://markdown.de/syntax/). Allerdings gab es da noch ein
Problem. Eigentlich drei. Die hießen header.tmpl, footer.tmpl und
style.css. Ich bin in solchen Design-Dingen sehr schlecht, aber mein
Berufsschulkollege [Holger](http://savier.zwetschge.org), der öfters mal
wunderschöne Templates aus dem Ärmel schüttelt, sagte natürlich nicht nein
und erstellte mir die nötigen Files <3. Unter
[http://ice-bullocks.zwetschge.org](http://ice-bullocks.zwetschge.org) kann
man das Ergebnis begutachten. Sind Leberkäse-Brötchen eigentlich anerkannte
Währung unter Designern ?

Den Source zum Anschauen gibts unter:
[bash-script](http://git.zwetschge.org/?p=this-is-static-markdown-content.git;a=blob;f=tismc.bash)
[Homepage-Dir](http://git.zwetschge.org/?p=this-is-static-markdown-content.git;a=tree)
