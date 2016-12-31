---
aliases:
- /archives/854
- /blog/2010/01/26/web-paste-service-via-commandline-sprunge-us
comments:
- author: Maxe
  content: "<p>Mein Lieblingspastebin ( <a href=\"http://paste.pocoo.org/\" rel=\"nofollow\">http://paste.pocoo.org/</a>
    ) bietet auch ein Konsolentool \u2192 <a href=\"http://paste.pocoo.org/help/integration/\"
    rel=\"nofollow\">http://paste.pocoo.org/help/in...</a></p><p>:P</p>"
  date: '2010-01-27T01:30:18'
- author: Sven
  content: "<p>Das ist ja gro\xDFartig.<br>Ich habe bei Opera gerade eine neue Suche
    definiert und f\xFCr diesen Dienst angepasst. Hat keine halbe Minute gedauert.<br>So
    kann ich im Browser ein Text markieren und statt nach ihm zu suchen eben zu <a
    href=\"http://sprunge.us\" rel=\"nofollow\">sprunge.us</a> schicken. Oder in die
    Adresszeile per \"sprunge text der gepastet werden soll\".<br>Vielen Danke daf\xFCr!</p>"
  date: '2010-01-27T05:29:23'
- author: noqqe
  content: '<p>@Maxe: Ja paste.pocoo hab ich auch lang benutzt. aber das tool kannte
    ich garnicht...</p><p>@sven: Ja? So hab ich das noch garnicht gesehen. Super improvisiert!
    :D</p>'
  date: '2010-01-29T00:33:32'
date: '2010-01-26T16:28:38'
tags:
- development
- web
- python-setuptools
- shell
- script
- python
- py
- ubuntu
- skript
- alias
- easy_install
- sprunge
- sprang
- linux
- php
- debian
- bash
title: Web | Paste-Service via CommandLine (Sprunge.us)
---

[Sprunge.us](http://sprunge.us) ist ein Paste-Service den ich heute von
[Chris](http://cryzed.de) gezeigt bekommen habe. Sprunge ist aber außerdem
noch _awesome_, weil er ohne Registrierung oder Umstände alles annimmt was
man ihm via `curl -F `übergibt. Von den Entwicklern ist das wie folgt
vorgesehen:

``` bash
<command> | curl -F 'sprunge=<-' http://sprunge.us
INFO: Code: gJIJ
INFO: URL: http://sprunge.us/gJIJ

```

Und man kann unter der ausgespuckten URL den SourceCode begutachten. Den
curl-Aufruf finde ich persöhnlich ziemlich lang und nicht wirklich
eingängig. Das fanden anscheinend auch die Entwickler von
"[sprang](http://github.com/jingleman/sprang)". Usage ungefähr so:

``` bash
cat /usr/local/scripts/script.sh | sprang
INFO: Code: gJIJ
INFO: URL: http://sprunge.us/gJIJ
```

[sprang](http://github.com/jingleman/sprang) ist ein Python-Script das mit
dem sprunge.us Pastebin-Dienst interagieren kann. Man kann ihm zum
Bleistift auch mit sprang -f ein Fileübergeben, mit -L Logfiles definieren
oder ähnliches bewerkstelligen (genaueres mit sprang --help). Durch die
Installation des python-setuptools bzw dem Kommando

``` bash
aptitude install python-setuptools; easy_install sprang
```

wird der Helfer für den Dienst nutzbar. Ich muss ehrlich gestehen ich bin
kein Fan von Fremdpaketsystemen. Aber diesbezüglich muss es eben sein.
Alternative ist natürlich ein Bash-alias

``` bash
alias sprang="curl -F 'sprunge=<-' http://sprunge.us"
```

Wobei somit die Restfunktionalität des sprang-scripts verloren geht.
Besonders schön ist auch das Syntax Highlightning. Je nach Eingespeisten
Source kann man der URL beispielsweise ein ?bash oder ?py mitgeben

```
http://sprunge.us/gJIJ?bash
http://sprunge.us/gJIJ?py
```

und erhält schön bunt und leserlich ge-Highlightete Versionen des gesendeten.