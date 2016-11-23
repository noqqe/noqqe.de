---
title: Redis
date: 2015-04-09T14:45:47.000000
tags: 
- Databases
- Redis
---


#### Server Details

print out all kinds of informations

    > INFO

#### DB Queries

Show all Values

    > KEYS *

Show summary of all data

    > KEYSPACE

Add some Data

    SET Key0 Value0
    SET Key1 Value1
    SET Key2 Value2

#### Replication

Configure a Slave and check afterwards

    > SLAVEOF host 6379
    > INFO

#### Sentinel

The following is a list of accepted commands:

This command simply returns PONG.

    > PING

Show a list of monitored masters and their state.

    > SENTINEL masters

Show the state and info of the specified master.

    > SENTINEL master <master name>

Show a list of slaves for this master, and their state.

    > SENTINEL slaves <master name>

Show master ip and port

    > SENTINEL get-master-addr-by-name <master name>

Force a failover as if the master was not reachable,

    > SENTINEL failover <master name>

Query a Redis Sentinel for current master

    > sentinel get-master-addr-by-name master01
    1) "1.2.3.4"
    2) "6379"

#### Commandline

Access a remote database

    $ redis-cli -h host -p 6379

Authenticate with password

    $ redis-cli -a CXXX

Useful queries from commandline example

    $ echo "SENTINEL get-master-addr-by-name master01" | redis-cli -h host -p 26379 -x

