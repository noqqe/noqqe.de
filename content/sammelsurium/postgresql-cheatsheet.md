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

    postgres=## \l

Make root a super user (ignore switching to postgres first)

    CREATE USER root;
    ALTER USER root WITH SUPERUSER;

Use Database

    postgres=## \c teamvault

Show Tables

    postgres=## \dt

List Roles

    postgres=## \du

Backup / Dump a single DB

    pg_dump db1 > backup.sql

Backup all Databases

    pg_dumpall > pgbackup.sql

Restore

    psql db1 < backup.sql
