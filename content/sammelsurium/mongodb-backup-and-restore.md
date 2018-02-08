---
title: MongoDB Backup and Restore
date: 2015-03-17T09:48:07
tags:
- Databases
- MongoDB
---


## Backup

    mongodump --db test --out /tmp/dump

## Restore

    mongorestore --db test /tmp/dump/test/


## Full Cluster Backup

* Pick one shard+configserver
* Stop Balancer via mongos
* Lock Shardserver
* Lock Configserver
* LVM Snapshot
* Unlock Shardserver
* Unlock Configserver
* mount snapshot
* tar content
* unmount snapshot
* delete lvm snapshot

## Full Cluster Restore

There are multiple scenarios of backing up

#### Shard Cluster Full Restore

* Shutdown mongos
* Shutdown configservers
* Shutdown shardservers
* extract shardserver dbpath to every server
* extract configserver dbpath to every configserver
* start up shard servers
* start up config servers
* start up mongos

#### Temp MongoDB Instance

if you have a fully sharded cluster and have backup from data dir of shard
servers you can

    mkdir /data/mongo1/testrestore
    cd /data/mongo1/testrestore
    tar xfvz /tmp/mongo04-2015-03-13.tar.gz ./mongodb
    mongod --bind_ip 127.0.0.1 --port 30000 --dbpath mongodb/
    mongodump --host 127.0.0.1 --port 30000 --db Project1_Test --out dump

Fun part is that you acutally import the dumped data into the mongos

    mongorestore --host 127.0.0.1 --port 27017 --username admin -pxxx --authenticationDatabase admin --db Project1RESTORE dump/Project1_Test/

#### Dump oplog

Obacht, --oplog does not dump the oplog collection.

    mongodump --db local --host localhost:27018 --username xxx -pxxx --collection oplog.rs --authenticationDatabase admin --out /tmp/

