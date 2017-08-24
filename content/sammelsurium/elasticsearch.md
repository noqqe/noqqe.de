---
title: Elasticsearch
date: 2016-03-02T09:48:08
tags:
- Software
- elasticsearch
---

#### Cluster

* Cluster bauen findet über Multicast im Netz statt (furchtbar)
* Man _kann_ unicast statt multicast benutzen.
* `master` wird gewählt

Informationen eines CLUSTERS anzeigen

    curl -XGET 'http://localhost:9200/_cluster/health?pretty=true'

Informationen eines HOST anzeigen

    curl localhost:9200/_nodes/stats/process?pretty

Alle Indices anzeigen

    curl localhost:9200/_aliases

Einen Index löschen

    http --auth elastic:xxx DELETE 'http://localhost:9200/beat-2017.08.19'
