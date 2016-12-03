---
aliases:
- /archives/886
- /blog/2010/02/15/wordpress-tags-in-theme-einbauen
date: '2010-02-15T06:50:01'
tags:
- development
- tags
- blog
- wordpress
- php
- tagging
title: Wordpress | Tags in Theme einbauen
---

Ich tagge ;). Und zwar In der Einzelansicht der (mehr oder minder) guten
Posts hier. Bewerkstelligen lies sich das mit der Hilfe der [Codex
Wordpress Site](http://codex.wordpress.org/Template_Tags/the_tags) und der
Datei single.php (Name kann je nach Theme variieren). Vordefinierte
Funktion von Wordpress:

```
<?php the_tags('Tagging: ',' | ','<br/>'); ?>
```