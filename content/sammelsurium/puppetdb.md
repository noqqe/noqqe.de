---
title: "PuppetDB"
date: 2017-12-18T15:58:58+01:00
draft: false
tags:
  - Puppet
  - PuppetDB
---

Delete a node

``` sql
puppetdb=# delete from catalog_resources where title like '%hostname%';
DELETE 138
puppetdb=# delete from resource_params where value like '%hostname%';
DELETE 29
```

Find resources based on their name

``` sql
SELECT catalog_resources.file, certnames.certname
FROM catalog_resources
INNER JOIN certnames ON catalog_resources.certname_id = certnames.id
WHERE title LIKE '%<name of resource>%';
```
