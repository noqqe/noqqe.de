---
title: Elasticsearch
date: 2016-03-02T09:48:08.897000
tags: 
- Software
- elasticsearch
---


#### Cluster

* Does multicast on same clustername
* you can also use unicast with ips
* master will be elected

Get cluster health

    curl -XGET 'http://localhost:9200/_cluster/health?pretty=true'

Show infos about host

    curl localhost:9200/_nodes/stats/process?pretty
