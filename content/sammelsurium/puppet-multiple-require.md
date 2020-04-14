---
title: Puppet Multiple Require
date: 2013-10-29T15:07:21
tags: 
- Puppet
---

``` puppet
 require    => [ Group[users], Group[admin] ],
```