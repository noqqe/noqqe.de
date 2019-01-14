---
title: MongoDB Performance Test
date: 2014-03-31T14:38:08
tags:
- Databases
- MongoDB
---

writes

``` { .python }
import time
import pymongo
m = pymongo.MongoClient()

doc = {'a': 1, 'b': 'hat'}

i = 0

while (i < 1000000):

  start = time.time()
  m.tests.insertTest.insert(doc, manipulate=False, w=1)
  end = time.time()

  executionTime = (end - start) * 1000 ## Convert to ms

  print executionTime

  i = i + 1
```

> times.txt

[Benchmarks](https://blog.serverdensity.com/mongodb-benchmarks/)

[API Tutorial](http://api.mongodb.org/python/current/tutorial.html)

## Script with Auth

```
!/usr/bin/python
import time
import pymongo
import random
import string
import sys

uri = "mongodb://user:pw@localhost:27017/db"
m = pymongo.MongoClient(uri)

x = "e"
for i in xrange(10000):
    x = x + random.choice(string.letters)

doc = {'a': 1, 'b': x}

i = 0
while (i < 1000000):
        m.Loyalty.insertTest.insert(doc, manipulate=False, w=1)
        i = i + 1
```

## Verify the tests

``` { .json }
mongos> use devopstest
mongos> db.stats()
{
        "raw" : {
                "rs0/mongo01:27018,mongo02:27018" : {
                        "db" : "devopstest",
                        "collections" : 3,
                        "objects" : 110053,
                        "avgObjSize" : 40.00101769147592,
                        "dataSize" : 4402232,
                        "storageSize" : 11194368,
                        "numExtents" : 8,
                        "indexes" : 2,
                        "indexSize" : 8380400,
                        "fileSize" : 201326592,
                        "nsSizeMB" : 16,
                        "dataFileVersion" : {
                                "major" : 4,
                                "minor" : 5
                        },
                        "ok" : 1
                },
                "rs1/mongo03:27018,mongo04:27018" : {
                        "db" : "devopstest",
                        "collections" : 3,
                        "objects" : 110159,
                        "avgObjSize" : 40.0010167122069,
                        "dataSize" : 4406472,
                        "storageSize" : 11194368,
                        "numExtents" : 8,
                        "indexes" : 2,
                        "indexSize" : 8519392,
                        "fileSize" : 201326592,
                        "nsSizeMB" : 16,
                        "dataFileVersion" : {
                                "major" : 4,
                                "minor" : 5
                        },
                        "ok" : 1
                }
        },
        "objects" : 220212,
        "avgObjSize" : 40,
        "dataSize" : 8808704,
        "storageSize" : 22388736,
        "numExtents" : 16,
        "indexes" : 4,
        "indexSize" : 16899792,
        "fileSize" : 402653184,
        "ok" : 1
}
```

## Overview 1 Mio Queries from 1 Machine

Auswertung mit R

```
> x <- read.csv("C:/cygwin64/home/noqqe/single.txt", as.is=T, header=F)
> head(x)
         V1
1 1.1410713
2 0.6830692
3 0.7622242
4 0.6301403
5 0.6930828
6 0.6759167
> hist(x)
> mean(x)
> mean(x[,1])
[1] 0.5329964
>
> summary(x[,1])
     Min.   1st Qu.    Median      Mean   3rd Qu.      Max.
   0.3819    0.4399    0.4652    0.5330    0.5040 1054.0000

> x <- x[,1]
> sort(x[x>=100])
 [1]  102.5641  106.1161  114.6011  139.0510  144.7840  192.0590  193.9261  195.7710  203.4011  204.2670  209.1570  209.5540  210.3591  211.7610  212.6400  486.0430  555.4299
[18]  599.1330 1053.8390
hist(x[x>=100&x<1000],main="Distribution of big queries")

> summary(y)
     Min.   1st Qu.    Median      Mean   3rd Qu.      Max.
    0.325     0.598     0.786     1.773     1.186 23870.000
```
