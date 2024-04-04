---
title: PostgreSQL
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

Create new Database with user:

```
CREATE DATABASE mydb;
CREATE USER myuser WITH ENCRYPTED PASSWORD 'mypass';
GRANT ALL PRIVILEGES ON DATABASE mydb TO myuser;
ALTER DATABASE mydb OWNER TO admin;
```

Use Database

    postgres=## \c teamvault

Delete all Tables

```
DROP SCHEMA public CASCADE;
CREATE SCHEMA public;
GRANT ALL ON SCHEMA public TO <role>;
GRANT ALL ON SCHEMA public TO public;
```


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

Backup only the Schema

    pg_dump --schema-only -d db1 > /tmp/foo.sql

Backup only Schema of one Table

    pg_dump --schema-only -t settings -d db1 > /tmp/foo.sql

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

Show all roles available on the instance

```sql
SELECT rolname FROM pg_roles;

       rolname
----------------------
 pg_monitor
 pg_read_all_settings
 pg_read_all_stats
 pg_stat_scan_tables
 pg_signal_backend
 rds_superuser
 rds_replication
 rds_iam
 rds_password
 rdsadmin
 rdsrepladmin
 lambda
 postgres
```

## Docker-Compose

Ein einfaches Postgres `docker-compose` Setup mit Backup

```
  postgres:
      image: postgres
      restart: always
      volumes:
        - postgres_data:/var/lib/postgresql/data
      environment:
        POSTGRES_DB: postgres
        POSTGRES_USER: root
        POSTGRES_PASSWORD: 'xxx'

  pgbackups:
    image: prodrigestivill/postgres-backup-local
    restart: always
    ports:
      - "9876:9876"
    volumes:
      - 'pgbackup:/backups'
    depends_on:
      - postgres
    environment:
      POSTGRES_HOST: postgres
      POSTGRES_DB: postgres
      POSTGRES_USER: root
      POSTGRES_PASSWORD: 'xxx'
      POSTGRES_EXTRA_OPTS: '-Z9 --schema=public --blobs'
      SCHEDULE: '@daily'
      BACKUP_KEEP_DAYS: 7
      BACKUP_KEEP_WEEKS: 4
      BACKUP_KEEP_MONTHS: 6
      HEALTHCHECK_PORT: 9876`
```
