---
title: Apache Solr
date: 2013-04-25T17:53:44
tags:
- Software
- solr
---

## Collection API

Erstellen

     curl "http://solr01.example.com:8080/solr/admin/collections?action=CREATE&name=#NAME#&numShards=1&replicationFactor=1"

Löschen

     curl "http://solr01.example.com:8080/solr/admin/collections?action=DELETE&name=#NAME#"

Mass Delete

      for x in $PSS ; do curl --silent "http://solr01.example.com:8080/solr/admin/collections?action=DELETE&name=${x}_${i}" ; sleep 2 ; done

## Cores API

* CREATE
* RELOAD
* UNLOAD
* SWAP
* RENAME

Erstellen

    curl "http://solr01.example.com:8080/solr/admin/cores?action=CREATE&name=#NAME#_shard1_replica1&collection=#COLLECTION#&shard=shard1&collection.configName=db"

Löschen

    curl "http://solr01.example.com:8080/solr/admin/cores?action=UNLOAD&core=#NAME#"

    &deleteInstanceDir=TRUE will do what both &deleteDataDir=TRUE and &deleteIndex=TRUE

Status

    curl "http://solr01.example.com:8080/solr/admin/cores?action=STATUS&core=#NAME#"

## ContextLinks Queries

Löschen

    curl --silent "http://admin:xxx@cl01.example.com:8080/db/contextLinks?command=deleteIndex&indexId=#NAME#"