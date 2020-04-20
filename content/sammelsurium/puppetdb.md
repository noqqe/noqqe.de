---
title: "PuppetDB"
date: 2017-12-18T15:58:58+01:00
draft: false
tags:
- Puppet
- PuppetDB
- Software
---

Delete a node

``` sql
delete from catalog_resources where title like '%hostname%';
delete from resource_params where value like '%hostname%';
```

Find resources based on their name

``` sql
SELECT catalog_resources.file, certnames.certname
FROM catalog_resources
INNER JOIN certnames ON catalog_resources.certname_id = certnames.id
WHERE title LIKE '%<name of resource>%';
```
