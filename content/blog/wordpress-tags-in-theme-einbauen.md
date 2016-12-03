---
date: 2010-02-15T08:50:01+02:00
comments: true
title: Wordpress | Tags in Theme einbauen
aliases:
- /archives/886
- /blog/2010/02/15/wordpress-tags-in-theme-einbauen
tags:
- Development
- Blog
- PHP
- tagging
- tags
- wordpress
---

Ich tagge ;). Und zwar In der Einzelansicht der (mehr oder minder) guten
Posts hier. Bewerkstelligen lies sich das mit der Hilfe der [Codex
Wordpress Site](http://codex.wordpress.org/Template_Tags/the_tags) und der
Datei single.php (Name kann je nach Theme variieren). Vordefinierte
Funktion von Wordpress:

```
<?php the_tags('Tagging: ',' | ','<br/>'); ?>
```

