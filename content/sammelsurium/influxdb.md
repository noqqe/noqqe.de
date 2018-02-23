---
title: InfluxDB
date: 2016-10-20T15:46:50
tags:
- InfluxDB
- Software
---

## Datenbanken

DB auswählen

    show database
    use telegraf

Oder auch beim connecten

    influx -database telegraf

Erstellen

    create database telegraf

Löschen

    drop database telegraf

## Series

Show all series

    use telegraf
    show series

## Measurements

    show measurements

## Queries

Alle series anzeigen

    select * from mqtt_consumer

Alle Series löschen

    drop series from mqtt_consumer

Where ist wie bei SQL

    select * from mqtt_consumer where topic='sensors/spacestatus/temp'

## curl Queries

### Databases

Create

    curl -i -XPOST http://localhost:8086/query --data-urlencode "q=CREATE DATABASE besucher"

Drop

    curl -i -XPOST http://localhost:8086/query --data-urlencode "q=DROP DATABASE besucher"

### Data

Select

    curl -i -XPOST http://localhost:8086/query  --data-urlencode "db=sensors" --data-urlencode 'q=SELECT "besucher" FROM "besucher"'

Insert

    curl -i -XPOST 'http://localhost:8086/write?db=besucher' --data-binary "besucher besucher=15 "

Insert with specific time

    curl -i -XPOST 'http://localhost:8086/write?db=besucher' --data-binary "besucher besucher=15 1470866400000000"

Delete all data from Table

    curl -i -XPOST http://localhost:8086/query  --data-urlencode "db=sensors" --data-urlencode 'q=DROP SERIES FROM tinker_light'
