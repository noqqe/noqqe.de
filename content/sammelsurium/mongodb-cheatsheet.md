---
title: MongoDB Cheatsheet
date: 2014-10-11T17:59:52
tags:
- Databases
- MongoDB
---

## Query

```
db.coll.find( {WHERECLAUSE}, {FIELDSELECTION,FIELDSELECTION2})
db.invoices.find()
db.invoices.find({"value":{$gt:12000}})
```

## Insert

```
db.invoices.insert({ "company": "a", "value": 15000 })
db.invoices.insert({ "company": "b", "value": 12000 })
db.invoices.insert({ "company": "b", "value": 12000 })
db.invoices.insert({ "company": "c", "tax": 1000 })
```

## Update

```
db.links.update({"url": /\.gif/},{$set: {"title": "GIF"}},{multi: true})
db.invoices.update({company: "a"}, {$set: {"company": "z"}})
```

## Remove

```
db.links.remove({"_id": ObjectId("567075b16815b419f8e35bba")})
db.invoices.deleteOne({company:"b"})
```

## Funktionen

```
.limit(10)
.skip(12)
.sort()

db.invoices.find().limit(2)
db.invoices.find().sort({value: 1})
db.invoices.find().sort({value: 1}).limit(1)
db.invoices.find().pretty()
```

## Operators

Zum Beispiel

    db.coll.find({"value:" {$gte:200}})

also

```
$gte: greater then or equal
$gt
$lte
$lt
$in
$type
$or
$exist
```

default query ist AND mit ","

## mongoimport

    mongoimport --db lolo --collection products < products.json
