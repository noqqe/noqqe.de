---
layout: post
title: "Switched to Octopress"
date: 2012-03-05T21:22:00+02:00
comments: true
categories:
- Web
- Blog
- Ruby
- Python
tags:
- Blogging
- Posts
- Framework
- rvm
- switched
- exitwp
- Wordpress
---

Hatte von [Octopress](http://octopress.org) gelesen. Dann nochmal im [k4cg](http://k4cg.org) drüber
gesprochen. Irgendwie. So halb und anschliessend für super befunden.

{% img center /uploads/2012/03/05/zoidberg.png %}

Ich weiss nicht warum, aber irgendwie hat mich das Thema dann einfach nicht mehr
los gelassen. Hier probiert, da gemacht. Das Ende vom Lied war dann, dass ich
ein fast vollstänidg migriertes Blog in Octopress lokal gebaut hatte. Dann wars
auch schon egal und ich hab mein Wordpress eingestampft.

Eigentlich ging alles ganz fix.

* Ruby 1.9.2 via rvm installiert
* Die Anleitung auf [octopress.org](http://octopress.org/docs/setup/) befolgt
* In Wordpress Inhalte in wordpress.xml exportiert
* Dem Konvertierungstool [exitwp](https://github.com/thomasf/exitwp) vorgeworfen
* Alles in `source/_posts` platziert
* wordpress.xml nach [Disqus.com](http://disqus.com) hochgeladen
* Viele sed/awk Spielchen für die richtigen ImagePfade und Code Block
Formattierung betireben
* Deployed auf Webserver
* Apache Redirects für alte Posts eingerichtet

Lediglich das umgemappe und durch die Gegend redirecte war müßig. Aber naja.
Jetzt ist schön und auch schon umgezogen. Mir gefällts.

Was mir bei der Recherche noch aufgefallen ist:

{% blockquote BalusC, http://stackoverflow.com/questions/3912694/using-markdown-how-do-i-center-an-image-and-its-caption stackoverflow  %}
The &lt;center&gt; element is deprecated in HTML4 since 1998. Just use CSS to make it
or its wrapper a block element with a fixed width and margin: 0 auto
{% endblockquote %}

WTF. Ich habe ca. 7 Jahre danach das erste mal ein Buch gelesen, dass sich mit
HTML beschäftigt und es war dort schon deprecated. Das war mir nicht klar.
