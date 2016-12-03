---
title: Restic
date: 2016-06-16T12:18:08
tags: 
- restic
- backup
- Software
---

#### Configuration

    export RESTIC_REPOSITORY=/path/to/repo
    export RESTIC_PASSWORD=lustigesPW

#### Initialization

    restic -r /path/to/repo init

#### Backup

Ohne env variablen

    restic -r /path/to/repo backup /path/to/backup

Mit

    restic backup /path/to/backup

#### Restore a single file

    restic restore 62324d64 --include=path/within/snapshot/foo.csv -t /tmp/
