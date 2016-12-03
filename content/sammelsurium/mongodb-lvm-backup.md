---
title: MongoDB LVM Backup
date: 2014-12-09T17:32:00
tags: 
- Databases
- MongoDB
---

created some snapshots (without data and with)
Create the Backups after configuring

    mongolvmbackup.sh -g data -v data

see output

```
## ls -lah
total 26M
drwxr-xr-x  2 root root 4.0K Dec  9 15:54 .
drwxr-xr-x 25 root root 4.0K Dec  9 14:14 ..
-rw-r--r--  1 root root  11K Dec  9 14:21 data-2014-12-09_1415-snap.tbz2
-rw-r--r--  1 root root  11K Dec  9 14:35 data-2014-12-09_1430-snap.tbz2
-rw-r--r--  1 root root  26M Dec  9 15:54 data-2014-12-09_1543-snap.tbz2
lrwxrwxrwx  1 root root   30 Dec  9 15:54 latest.tbz2 -> data-2014-12-09_1543-snap.tbz2
```

created the data with

```
for (var i = 1; i <= 1000000; i++) { db.test.insert({1:2, 3:4, 5: "foo", 6: false }) }
> db.test.stats()
{
        "ns" : "test.test",
        "count" : 1000000,
        "size" : 112000000,
        "avgObjSize" : 112,
        "storageSize" : 174735360,
        "numExtents" : 12,
        "nindexes" : 1,
        "lastExtentSize" : 50798592,
        "paddingFactor" : 1,
        "systemFlags" : 1,
        "userFlags" : 1,
        "totalIndexSize" : 32458720,
        "indexSizes" : {
                "_id_" : 32458720
        },
        "ok" : 1
}
```

im going to add / change a few things now

```
> db.test.re
db.test.reIndex(           db.test.remove(            db.test.renameCollection(
> db.test.remove({"_id":ObjectId("54870c6f64b8ed4efdbf2f46")})
WriteResult({ "nRemoved" : 1 })
> db.test.stats()
{
        "ns" : "test.test",
        "count" : 999999,
        "size" : 111999888,
        "avgObjSize" : 112,
        "storageSize" : 174735360,
        "numExtents" : 12,
        "nindexes" : 1,
        "lastExtentSize" : 50798592,
        "paddingFactor" : 1,
        "systemFlags" : 1,
        "userFlags" : 1,
        "totalIndexSize" : 32458720,
        "indexSizes" : {
                "_id_" : 32458720
        },
        "ok" : 1
}
```

and update alle records

```
> db.test.update({},{$set: {6:true}}, {multi: true})
WriteResult({ "nMatched" : 999999, "nUpserted" : 0, "nModified" : 999999 })
> db.test.find({6:true}).count()
999999
> db.test.find({6:false}).count()
0
```

and now i do the restore


test restore:

```
## rm -rf *
## ls -lah
total 8.0K
drwxr-xr-x  2 mongodb mongodb 4.0K Dec  9 16:14 .
drwxr-xr-x 25 root    root    4.0K Dec  9 14:14 ..
## tar xfvj /snapshots/latest.tbz2
journal/
journal/prealloc.2
journal/prealloc.1
journal/prealloc.0
local.0
local.ns
lost+found/
mongod.lock
test.0
test.1
test.2
test.ns
_tmp/
## /etc/init.d/mongod start
```

and see the results

```
> db.test.stats()
{
        "ns" : "test.test",
        "count" : 1000000,
        "size" : 112000000,
        "avgObjSize" : 112,
        "storageSize" : 174735360,
        "numExtents" : 12,
        "nindexes" : 1,
        "lastExtentSize" : 50798592,
        "paddingFactor" : 1,
        "systemFlags" : 1,
        "userFlags" : 1,
        "totalIndexSize" : 32458720,
        "indexSizes" : {
                "_id_" : 32458720
        },
        "ok" : 1
}
> db.test.find({6:true}).count()
0
> db.test.find({6:false}).count()
1000000
```

worked.
