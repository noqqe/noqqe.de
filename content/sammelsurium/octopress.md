---
title: Octopress
date: 2012-03-04T14:45:37
tags: 
- Software
- octopress
---

In '_config.yml' einfügen:

    sidebar: collapse

## Isolate

Damit nur ein einzelner Post neu generiert wird:

    rake isolate[path/to/post.markdown]

Damit wieder alles neu generiert wird

    rake integrate

## Bilder

    [![](http://noqqe.de/uploads/2012/03/05/zoidberg.png)](http://noqqe.de/uploads/2012/03/05/zoidberg.png)

## Links

    [Link](http://link.com)

## Haudrauf Hammer wenn nix mehr geht

``` bash
git pull octopress master
bundle update
rake update_source
rake update_style
rake preview
```

## Disqus neben TIME

```
in octopress/source/_includes/article.html
{% if site.disqus_short_name and page.comments != false and post.comments != false and site.disqus_show_comment_count == true %}
| <a href="{% if index %}{{ root_url }}{{ post.url }}{% endif %}#disqus_thread">Comments</a>
{% endif %}
```

## Navigation ausschalten

```
octopress/source/_layouts/default.html
```