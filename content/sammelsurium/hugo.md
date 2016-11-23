---
title: hugo
date: 2015-04-10T10:58:48.000000
tags: 
- Software
- Hugo
---


#### Neuer Post

    hugo new post/postname.md

#### Neue Seite

    hugo new seitenname.md

#### Hugo Preview Server

    hugo server -vw

oder auch ohne watch

    hugo server -v

custom server and bind

    hugo server --baseURL=noc.example.com --bind=0.0.0.0

#### Deploy

    rm -rf public/ ; hugo ; /usr/local/bin/rsync -avi --delete --iconv=utf-8-mac,utf-8 public/ host:/var/www/htdocs/noqqe.de/
