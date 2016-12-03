---
title: MongoDB Usermanagement
date: 2014-03-11T14:17:07
tags: 
- Databases
- MongoDB
---

## Authentication

there are two directives indicating authentication

keyFile und auth

keyFile implies enabled auth!
Always create a admin user _BEFORE_ any other user

## Create a new Admin User

for mongodb 2.4

    use admin
    db.addUser( { user: "admin", pwd: "XXX", roles: [ "userAdminAnyDatabase", "clusterAdmin", "dbAdminAnyDatabase", "readWriteAnyDatabase" ], otherDBRoles: { config: [ "readWrite", "dbAdmin" ] } } )
    db.system.users.find()

for mongodb 2.6

    db.createUser( { user: "admin", pwd: "XXX", roles: ["root", "restore"] } )

## Create new user for a single database

MongoDB 2.4

    mongos> use Project1
    switched to db Project1
    mongos> db.addUser( { user: "Project1", pwd: "XXX", roles: [ "readWrite", "dbAdmin" ] } )
    mongos> db.system.users.find()
    { "_id" : ObjectId("532062fc07f14d1ead116815"), "user" : "Project2", "pwd" : "cb4a562cdaf742b68b9ecded34bc898f", "roles" : [ "readWrite", "dbAdmin" ] }

MongoDB 2.6

    mongos> use Project1
    switched to db Project1
    mongos> db.createUser( { user: "Project1", pwd: "XXX", roles: [ "readWrite", "dbAdmin" ] } )
    mongos> db.system.users.find()
    { "_id" : ObjectId("532062fc07f14d1ead116815"), "user" : "Project2", "pwd" : "cb4a562cdaf742b68b9ecded34bc898f", "roles" : [ "readWrite", "dbAdmin" ] }


after all, verify login

    mongo --username Project1 -pXXX Project1

## Update a user

full

    > db.system.users.remove({"user":"testAdmin"})
    > db.addUser( { user: "testAdmin",
                    pwd: "[whatever]",
                    roles: [ "clusterAdmin" ],
                    otherDBRoles: { TestDB: [ "readWrite", "dbAdmin" ] } } )

and in one line

    db.users.update({"user" : "testAdmin"}, {$addToSet: {'otherDBRoles.TestDB': 'dbAdmin'}}, false, false)

## Add new role to user

with update query

    db.system.users.update({ user:"admin" }, { $set:{ roles: [ "userAdminAnyDatabase", "clusterAdmin", "dbAdminAnyDatabase", "readWriteAnyDatabase" ] } })

MongoDB 2.6

    db.grantRolesToUser("admin", [ {role: "clusterAdmin",db: "admin"} ] )

## Delete a user

MongoDB 2.6

    use Project1
    db.dropUser("Project1", {w: "majority", wtimeout: 5000})

## Object workaround

~~~
mongos> obj=db.system.users.findOne({user:"admin"})
{
        "_id" : ObjectId("53205d22c90d1e727eb499eb"),
        "otherDBRoles" : {
                "config" : [
                        "readWrite",
                        "dbAdmin"
                ]
        },
        "pwd" : "0d15bab61475ebbd480074eea7a5b8bf",
        "roles" : [
                "userAdminAnyDatabase",
                "clusterAdmin",
                "dbAdminAnyDatabase",
                "readWriteAnyDatabase"
        ],
        "user" : "admin"
}
mongos> obj
{
        "_id" : ObjectId("53205d22c90d1e727eb499eb"),
        "otherDBRoles" : {
                "config" : [
                        "readWrite",
                        "dbAdmin"
                ]
        },
        "pwd" : "0d15bab61475ebbd480074eea7a5b8bf",
        "roles" : [
                "userAdminAnyDatabase",
                "clusterAdmin",
                "dbAdminAnyDatabase",
                "readWriteAnyDatabase"
        ],
        "user" : "admin"
}
mongos> obj.roles
[
        "userAdminAnyDatabase",
        "clusterAdmin",
        "dbAdminAnyDatabase",
        "readWriteAnyDatabase"
]
mongo> db.system.users.save(obj)
~~~
