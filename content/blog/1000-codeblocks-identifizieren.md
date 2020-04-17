---
title: ">1000 Codeblocks identifizieren"
date: 2020-04-17T15:19:40+02:00
tags:
- python
- markdown
---

Endlich kann [Hugo](https://gohugo.io) natives Syntax Highlighting. Endlich? Seit
[2017](https://gohugo.io/news/0.28-relnotes/)! Scheinbar ging das immer an
mir vorbei.

<!--more-->

Früher hatte ich über Javascript `pygments` geladen und damit die Codeblöcke
bunt angemalt, als wären sie Besucher eines Holi Festivals. Die clientseitige
JS Lösung hatte mir nie gefallen, aber war damals die einzige Möglichkeit. Es
hat sie Seite langsamer gemacht. Ich baute es bei *irgendeinem* Layoutwechsel
wieder aus.

Als Engine benutzt Hugo [Chroma](https://github.com/alecthomas/chroma). Aber
als ich es aktivierte, blieb trotzdem das meiste in tristem Schwarz & Weiss
zurück. Natürlich hatte ich über Jahre die jeweilige Sprache in den [Fenced Codeblocks](https://spec.commonmark.org/0.29/#code-fence)
nicht gepflegt. Ohne Auszeichnung kein Highlighting.

"Das müsste ich jetzt alles mal nachholen", dachte ich so. Wie? `vim
content/*.md`? Das wären 600 Dateien! "Und wenn ich nur die Files mit
Codeblöcken darin öffne?"

``` bash
$ grep ``` -R content/ | wc -l
    2226
```

Holy Shit. Über 1000 Codeblöcke! Was ich brauchte, war ein Weg um
schnell durch **unidentifizierte** Codeblocks zu gehen und diesen die
jeweilige Syntax zuzuordnen.

Ich setzte mich kurz hin und schrieb etwas `python`. Die Idee: Mein Script
findet alle untypisierten Codeblöcke eines Blogpostm, zeigt mir diese
nacheinander an und fragt mich, welche Sprache das ist.

``` python
for pos, line in enumerate(data):

  # does line contain a code fence?
  if re.findall('^```', line):

    # ignore closing ``` fences
    no = no + 1
    if no % 2 == 0:
        continue

    # ignore already typed codeblocks
    if re.findall('^```\s?\S', line):
        continue

    print("Found Codeblock:")
    print(line, end = '')

    # print lines until next occurence of ```
    print_codeblock(data[pos+1:])
```

Oben ist der eigentliche Auszug der Logik des Skripts. Das Ergebnis für einen
gefundenen Codeblock sah so aus:

{{< figure src="/uploads/2020/04/codeblock.png" caption="Ausgabe des Skripts" >}}

Das ganze Script ist auf
[github](https://gist.github.com/noqqe/534cb1b63c821b66f4df555180a06fc7) zu
finden. Angewendet habe ich es auf alle Dateien und es hat ungefährt 20
Minuten gedauert alle zu identifizieren:

``` fish
> for x in (ag -l ```); identify-codeblock.py $x ; end
> git diff e9a2924bb57bd59ddf3b975baa0cb9d4f864d00b 947398b3c4930f33d9620de6e2831540c107d0fd --shortstat
 159 files changed, 376 insertions(+), 376 deletions(-)
```

kkthxbye :wq!
