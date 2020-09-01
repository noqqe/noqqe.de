---
title: PostgreSQL Cheatsheet
date: 2015-07-09T15:28:29
tags:
- Databases
- PostgreSQL
---

Connect

    ## su - postgres
    $ psql

Create Database

    postgres=## CREATE DATABASE db1
    or
    $ createdb db1

List all databases

```
postgres=## \l
         List of databases
   Name    |  Owner   |   Access privileges
-----------+----------+-----------------------
 airflow   | postgres |
 configs   | postgres | =Tc/postgres         +
           |          | postgres=CTc/postgres+
           |          | lambda=c/postgres
 postgres  | postgres |
```

Make root a super user (ignore switching to postgres first)

    CREATE USER root;
    ALTER USER root WITH SUPERUSER;

Use Database

    postgres=## \c teamvault

Show Tables

```
configs=> \dt
          List of relations
 Schema |  Name   | Type  |  Owner
--------+---------+-------+----------
 public | airflow | table | postgres
```

List Roles

    postgres=## \du

Backup / Dump a single DB

    pg_dump db1 > backup.sql

Backup all Databases

    pg_dumpall > pgbackup.sql

Backup a remote database

    pg_dump -d <dbname> -h <hostname> -p5432 --username postgres  > /tmp/all.sql

Restore

    psql db1 < backup.sql

Get Permissions for a specific table

```sql
SELECT grantee, privilege_type
FROM information_schema.role_table_grants
WHERE table_name='airflow';

 grantee  | privilege_type
----------+----------------
 postgres | INSERT
 postgres | SELECT
 postgres | UPDATE
 postgres | DELETE
 postgres | TRUNCATE
 postgres | REFERENCES
 postgres | TRIGGER
 lambda   | INSERT
 lambda   | SELECT
 lambda   | UPDATE
 lambda   | DELETE
 lambda   | TRUNCATE
 lambda   | REFERENCES
 lambda   | TRIGGER
(14 rows)
```

