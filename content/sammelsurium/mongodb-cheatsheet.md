---
title: MongoDB Cheatsheet
date: 2014-10-11T17:59:52
tags: 
- Databases
- MongoDB
---

## Query

    db.coll.find( {WHERECLAUSE}, {FIELDSELECTION,FIELDSELECTION2})

## Methoden

    .limit(10)
    .skip(12)
    .sort()

Werden serverseitig aus der MongoShell ausgeführt mit Lazy Evaluation des eigentlichen Queries.

## Update

    db.links.update({"url": /\.gif/},{$set: {"title": "GIF"}},{multi: true})

## Remove

    db.links.remove({"_id": ObjectId("567075b16815b419f8e35bba")})

## Operators in queries

Zum Beispiel

    db.coll.find({$gte:200})

also

    $gte: greater then or equal
    $gt
    $lte
    $lt
    $in
    $type
    $or
    $exist

default query ist AND mit ","

## mongoimport

    mongoimport --db lolo --collection products < products.json
