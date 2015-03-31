---
layout: post
title: "Packages zu R 3.1 migrieren"
date: 2014-05-11T19:05:00+02:00
comments: true
categories:
- R
- stats
- Statistik
- Statistics
- rprojects
- packages
- upgrade
- osbn
---

Die [R](http://r-project.org) Version 3.1.0 "Spring Dance" wurde released.
Damit ändern sich auch die Modulpfade, sofern nicht global definiert.

Paketmigration einfach und pragmatisch

``` r 
R> x <- list.files("~/R/x86_64-pc-linux-gnu-library/3.0/")
R> sapply(x, install.packages)
```

Macht mehr R.
