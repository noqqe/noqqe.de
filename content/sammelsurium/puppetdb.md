---
title: "Puppetdb"
date: 2017-12-18T15:58:58+01:00
draft: false
tags:
  - Puppet
  - PuppetDB
---

Delete a node

```
puppetdb=# delete from catalog_resources where title like '%hostname%';
DELETE 138
puppetdb=# delete from resource_params where value like '%hostname%';
DELETE 29
```
