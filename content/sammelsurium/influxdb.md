---
title: InfluxDB
date: 2016-10-20T15:46:50
tags: 
- InfluxDB
- Software
---

#### Databases

Create

    curl -i -XPOST http://localhost:8086/query --data-urlencode "q=CREATE DATABASE besucher"

Drop

    curl -i -XPOST http://localhost:8086/query --data-urlencode "q=DROP DATABASE besucher"

#### Data

Select

    curl -i -XPOST http://localhost:8086/query  --data-urlencode "db=sensors" --data-urlencode 'q=SELECT "besucher" FROM "besucher"'

Insert

    curl -i -XPOST 'http://localhost:8086/write?db=besucher' --data-binary "besucher besucher=15 "

Insert with specific time

    curl -i -XPOST 'http://localhost:8086/write?db=besucher' --data-binary "besucher besucher=15 1470866400000000"

Delete all data from Table

    curl -i -XPOST http://localhost:8086/query  --data-urlencode "db=sensors" --data-urlencode 'q=DROP SERIES FROM tinker_light'
