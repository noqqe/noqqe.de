---
layout: post
title: "Kenne deine Repositories"
date: 2012-11-18T20:27:00+02:00
comments: true
categories:
- osbn
- bash
- git
- stats
---

Um einen Überblick der git Repositories zu bekommen,
kann man zuerst mal `locate` benutzen und diese zu finden.

{% codeblock lang:bash %}
locate --regex '.*\.git$'
{% endcodeblock %}

Aber wie viele Objekte halte ich in meinen Repos? Und wie viele Commits
eigentlich? Da ich momentan wieder vermehrt Bücher über Statistik lese,
brauchte ich Daten zum Spielen. git Repos bieten sich für sowas an.

{% img center /uploads/2012/11/git.jpg %}

### Wie viele Objects?

{% codeblock lang:bash %}
for x in $(locate --regex '.*\.git$'); do
    echo -e "$(locate "$x/objects" | wc -l)\t$(echo $x | sed 's/.git$//') "
done | sort -rn
{% endcodeblock %}

Dabei heraus kommt eine nette Übersicht über die Anzahl der git Objects

{% codeblock %}
7435  /home/noqqe/Code/ueberstatistic/
3580  /home/noqqe/.minecraft/
2649  /home/noqqe/Code/Tomcat-Schulung/
2402  /home/noqqe/Code/octopress/
1013  /home/noqqe/Code/dwarf-fortress/
415   /home/noqqe/Code/LaTeX-Dokumente/
[...]
{% endcodeblock %}

### Wie viele Commits?

{% codeblock lang:bash %}
for x in $(locate --regex '.*\.git$'); do
    echo -e "$(git --git-dir="$x" log --oneline | wc -l)\t$(echo $x | sed 's/.git$//')"
done | sort -rn
{% endcodeblock %}

Dass hier der ausgecheckte Linux Kernel alles in den Schatten stellt wundert
mich auch nicht ;)

{% codeblock %}
288160  /home/noqqe/Code/linux-stable/
2330    /home/noqqe/Code/fish-shell/
1090    /home/noqqe/Code/mixminion/
724     /home/noqqe/Code/octopress/
530     /home/noqqe/Code/bash-it/
[...]
{% endcodeblock %}


### Wie viele Zeilen pro Commit bearbeitet?

Innerhalb eines Repos sieht man auch oft interessante Werte über seine Commits.
Daran kann ich dann lauter andere lustige statistische Kennzahlen festmachen wie
die Standartabweichung usw. Aber darauf geh ich jetzt hier nicht näher ein.

{% codeblock lang:bash %}
for y in $(git log --format="%h");
  do git show $y | egrep '^(\+|\-)' | wc -l
done
{% endcodeblock %}
