---
title: MongoDB Usermanagement
date: 2014-03-11T14:17:07
tags: 
- Databases
- MongoDB
---

## Authentication

`keyFile` und `auth`. `keyFile` implies enabled `auth`!

Always create a admin user _BEFORE_ any other user

## Create a new Admin User

    > use admin
    > db.createUser({user: "admin", pwd: "admin", roles: [ { role: "userAdminAnyDatabase", db: "admin" } ] })

## Create new user for a single database

read only user

    > use test
    > db.createUser({user: "user", pwd:  "pass", roles: [ { role: "readWrite", db: "test" } ]})

read only user

    > use test
    > db.createUser({user: "user", pwd:  "pass", roles: [ { role: "readWrite", db: "test" } ]})

## Show existing users

    > db.system.users.find()

## Update a user

    > db.users.update({"user" : "testAdmin"}, {$addToSet: {'otherDBRoles.TestDB': 'dbAdmin'}}, false, false)

## Add new role to user


    > db.grantRolesToUser("admin", [ {role: "clusterAdmin",db: "admin"} ] )

## Delete a user

    > use Project1
    > db.dropUser("Project1", {w: "majority", wtimeout: 5000})
