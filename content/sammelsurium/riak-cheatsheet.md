---
title: Riak Cheatsheet
date: 2016-05-30T15:14:45
tags:
- Databases
- Riak
---

#### Cluster Informations

    riak-admin ring-status
    riak-admin vnode-status
    riak-admin aae-status
    riak-admin diag
    riak-admin status
    curl http://127.0.0.1:8098/stats | python -m json.tool | less

#### Queries

List all buckets

    curl 127.0.0.1:8098/buckets?buckets=true

List all keys - BE CAREFUL - DONT DO ON PRODUCTION

    curl 127.0.0.1:8098/buckets/bucket/keys?keys=true

Add Content

    curl -v 127.0.0.1:8098/buckets/training/keys/test\?returnbody=true\&pw=1 -X PUT -H "content-type: text/plain" -d "testing"

or

    curl -XPUT -H "Content-Type: text/plain" -d "vroom" http://localhost:8098/types/default/buckets/dodge/keys/viper?w=3

Reading Objects

    curl http://localhost:8098/types/default/buckets/dodge/keys/viper?r=3 ; echo
    vroom

#### Riak Cluster Setup

Initial

    riak-admin cluster clear
    riak-admin cluster status

Make sure each nodes can see all other nodes. Then join each node to one
specific node.

    riak-admin cluster join riak@10.20.30.40
    riak-admin cluster status

Commit Changes and make them permanent

    riak-admin cluster plan
    riak-admin cluster commit

Verify

    watch -n 1 "riak-admin transfers"
    riak-admin test
    riak-admin top

