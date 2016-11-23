---
title: MongoDB Shard und Replica Sets
date: 2013-08-13T18:54:51.000000
tags: 
- Databases
- MongoDB
---


* 4 shards
* 3 configservers
* 1 mongos (router)

### enable auth

~~~
openssl rand -base64 741 > /etc/mongodb.keyfile
~~~

add to config

~~~
configdb = mongo01.example.com,mongo02.example.com,mongo03.example.com
where to log
logpath=/var/log/mongodb/mongos.log
logappend=true
port = 27017
pidfilepath = /var/run/mongodb/routesrv.pid
keyFile = /etc/mongodb.keyfile ## KEYFILE
~~~

### create replica sets

in the config file do:

    ## on 01 and 02
    replSet = rs0
    ## on 03 and 04
    replSet = rs1

restart all shard servers. than connect to the local shard directly


#### on 01

~~~
rs0:PRIMARY> rs.initiate()
rs0:PRIMARY> rs.add("mongo02:27018")
rs0:PRIMARY> rs.conf()
{
        "_id" : "rs0",
        "version" : 4,
        "members" : [
                {
                        "_id" : 0,
                        "host" : "mongo01:27018"
                },
                {
                        "_id" : 1,
                        "host" : "mongo02:27018"
                }
        ]
}
rs0:PRIMARY> rs.status()
{
        "set" : "rs0",
        "date" : ISODate("2014-03-10T16:09:09Z"),
        "myState" : 1,
        "members" : [
                {
                        "_id" : 0,
                        "name" : "mongo01:27018",
                        "health" : 1,
                        "state" : 1,
                        "stateStr" : "PRIMARY",
                        "uptime" : 2876,
                        "optime" : {
                                "t" : 1394466931,
                                "i" : 1
                        },
                        "optimeDate" : ISODate("2014-03-10T15:55:31Z"),
                        "self" : true
                },
                {
                        "_id" : 1,
                        "name" : "mongo02:27018",
                        "health" : 1,
                        "state" : 2,
                        "stateStr" : "SECONDARY",
                        "uptime" : 818,
                        "optime" : {
                                "t" : 1394466931,
                                "i" : 1
                        },
                        "optimeDate" : ISODate("2014-03-10T15:55:31Z"),
                        "lastHeartbeat" : ISODate("2014-03-10T16:09:08Z"),
                        "lastHeartbeatRecv" : ISODate("2014-03-10T16:09:08Z"),
                        "pingMs" : 0,
                        "syncingTo" : "mongo01:27018"
                }
        ],
        "ok" : 1
}
~~~

#### on 03

~~~
rs0:PRIMARY> rs.initiate()
rs0:PRIMARY> rs.add("mongo04:27018")
rs1:PRIMARY> rs.conf()
{
        "_id" : "rs1",
        "version" : 4,
        "members" : [
                {
                        "_id" : 0,
                        "host" : "mongo03:27018"
                },
                {
                        "_id" : 1,
                        "host" : "mongo04:27018"
                }
        ]
}
rs1:PRIMARY> rs.status()
{
        "set" : "rs1",
        "date" : ISODate("2014-03-10T16:08:52Z"),
        "myState" : 1,
        "members" : [
                {
                        "_id" : 0,
                        "name" : "mongo03:27018",
                        "health" : 1,
                        "state" : 1,
                        "stateStr" : "PRIMARY",
                        "uptime" : 2207,
                        "optime" : {
                                "t" : 1394466997,
                                "i" : 1
                        },
                        "optimeDate" : ISODate("2014-03-10T15:56:37Z"),
                        "self" : true
                },
                {
                        "_id" : 1,
                        "name" : "mongo04:27018",
                        "health" : 1,
                        "state" : 2,
                        "stateStr" : "SECONDARY",
                        "uptime" : 735,
                        "optime" : {
                                "t" : 1394466997,
                                "i" : 1
                        },
                        "optimeDate" : ISODate("2014-03-10T15:56:37Z"),
                        "lastHeartbeat" : ISODate("2014-03-10T16:08:51Z"),
                        "lastHeartbeatRecv" : ISODate("2014-03-10T16:08:52Z"),
                        "pingMs" : 0,
                        "syncingTo" : "mongo03:27018"
                }
        ],
        "ok" : 1
}
~~~

### add shards to replica sets

~~~
mongos> sh.addShard("rs0/mongo01:27018,mongo02:27018")
mongos> sh.addShard("rs1/mongo03:27018,mongo04:27018")
mongos> sh.status()
--- Sharding Status ---
  sharding version: {
        "_id" : 1,
        "version" : 3,
        "minCompatibleVersion" : 3,
        "currentVersion" : 4,
        "clusterId" : ObjectId("531ddff023ab5da608d4c665")
}
  shards:
        {  "_id" : "rs0",  "host" : "rs0/mongo01:27018,mongo02:27018" }
        {  "_id" : "rs1",  "host" : "rs1/mongo03:27018,mongo04:27018" }
  databases:
        {  "_id" : "admin",  "partitioned" : false,  "primary" : "config" }



~~~

### add sharded database

~~~
use DATABASE
sh.enableSharding("DATABASE")
mongos> sh.status()
--- Sharding Status ---
  sharding version: {
        "_id" : 1,
        "version" : 3,
        "minCompatibleVersion" : 3,
        "currentVersion" : 4,
        "clusterId" : ObjectId("53205448d1731e5d141a9262")
}
  shards:
        {  "_id" : "rs0",  "host" : "rs0/mongo01:27018,mongo02:27018" }
        {  "_id" : "rs1",  "host" : "rs1/mongo03:27018,mongo04:27018" }
  databases:
        {  "_id" : "admin",  "partitioned" : false,  "primary" : "config" }
        {  "_id" : "Daphne",  "partitioned" : true,  "primary" : "rs0" }
        {  "_id" : "Loyalty",  "partitioned" : true,  "primary" : "rs0" }
        {  "_id" : "test",  "partitioned" : false,  "primary" : "rs1" }
~~~

### mongo Primary switch

on primary to become secondary

~~~
rs.stepDown()
~~~


