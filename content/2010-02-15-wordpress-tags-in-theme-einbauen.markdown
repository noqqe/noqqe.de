---
date: '2010-02-15 08:50:01'
layout: post
slug: wordpress-tags-in-theme-einbauen
status: publish
comments: true
title: Wordpress | Tags in Theme einbauen
alias: /archives/886
categories:
- Coding
- General
- PHP
tags:
- einbauen
- PHP
- tagging
- tags
- werdasliesthatverstandenumwasesimpostgeht
- wordpress
---

Ich tagge ;). Und zwar In der Einzelansicht der (mehr oder minder) guten Posts hier. Bewerkstelligen lies sich das mit der Hilfe der [Codex Wordpress Site](http://codex.wordpress.org/Template_Tags/the_tags) und der Datei single.php (Name kann je nach Theme variieren). Vordefinierte Funktion von Wordpress:
```
<?php the_tags('Tagging: ',' | ','<br/>'); ?>
```

