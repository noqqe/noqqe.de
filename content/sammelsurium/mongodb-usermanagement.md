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

    > use Project1
    switched to db Project1
    > db.createUser( { user: "Project1", pwd: "XXX", roles: [ "readWrite", "dbAdmin" ] } )

## Show existing users

    > db.system.users.find()

after all, verify login

    mongo --username Project1 -pXXX Project1

## Update a user

    > db.users.update({"user" : "testAdmin"}, {$addToSet: {'otherDBRoles.TestDB': 'dbAdmin'}}, false, false)

## Add new role to user


    > db.grantRolesToUser("admin", [ {role: "clusterAdmin",db: "admin"} ] )

## Delete a user

    > use Project1
    > db.dropUser("Project1", {w: "majority", wtimeout: 5000})
