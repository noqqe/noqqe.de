---
date: 2010-04-05T20:01:16+02:00
type: post
slug: wordpress-archive-page-erstellen
comments: true
title: Wordpress | Archive Page erstellen
aliases:
- /archives/970
categories:
- Development
- PHP
- PlanetenBlogger
tags:
- archiv
- archive
- clouds
- contents
- html
- PHP
- theme
- wordpress
- wp
---

War 'ne schwere Geburt, was sich da jetzt unter [Archive](/?page_id=927) versteckt. Zum einen (ungewöhnlicher Weise) schlecht und nicht umfassend genug dokumentiert bei [codex.wordpress.org](http://codex.wordpress.org/Creating_an_Archive_Index) und zum anderen ist mein aktuelles Theme schon etwas widerspenstig was solche Sachen angeht. Unter anderem diese Widerspenstigkeit war der Punkt, warum ich kein [Plugin](http://wordpress.org/extend/plugins/search.php?q=archive&sort=) verwenden konnte.

Umständlich aber trotzdem zielführend. Die single.php stellt die Einzelansicht eines Blogposts. Super das will ich ja. Die Archiv-Template-Seite kann man in dem Theme-Dir ruhig liegen lassen:

```
cd /var/www/blog/wp-contents/themes/clouds/
cp single.php archive.php
```


Darauf folgend habe ich die Post-Aufrufe aus dem Quelltext entfernt. Der Source ist relativ selbsterklärend. Somit besteht nur noch das Grundgerüst einer Einzelansicht. Wo sich vorher die Aufrufe für die Posts befand, fügte ich folgendes ein:

```
<br/><b>Monthly outline</b>
<?php wp_get_archives('type=monthly&show_post_count=1'); ?>
<br/><b>Post outline by date</b>
<?php wp_get_archives('type=postbypost');?>
```


Die Funktion wp_get_archives ist allerdings [sehr schön dokumentiert](http://codex.wordpress.org/Template_Tags/wp_get_archives). Die Kunst an dem eigentlich Spass war aber das einbinden dieser archive.php-Datei.

Damit Wordpress erkennt, dass es sich bei archive.php um eine Template-Datei handelt:

```
<?php
/**
* @package WordPress
* @subpackage Default_Theme
*/
/*
Template Name: Archives
*/
?>
```


am Anfang der Datei einfügen. Jetzt noch in Wordpress die Seite erstellen. Dazu musste ich eine Seite anlegen mit leerem Inhalt. Anschliessend bei Seiten -> bearbeiten -> Quickedit(!) -> Template: Archive auswählen.
