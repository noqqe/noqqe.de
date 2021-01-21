---
title: Restic
date: 2016-06-16T12:18:08
tags:
- restic
- backup
- Software
---

Restic ist eine Backup Software die so ziemlich alle Kriterien erfüllt die
ich brauche.

<!-- more -->

## Configuration

    export RESTIC_REPOSITORY=/path/to/repo
    export RESTIC_PASSWORD=lustigesPW

## Initialization

    restic -r /path/to/repo init

## Backup

Ohne env variablen

    restic -r /path/to/repo backup /path/to/backup

Mit

    restic backup /path/to/backup

## Cleanup / Prune

Um 7 Tage, 4 Wochen und 3 Monate Backups zu behalten:

    restic forget --prune --keep-daily 7 --keep-weekly 4 --keep-monthly 3

#### Restore a single file

Weil ich damit immer kämpfe...

    restic restore 62324d64 --include=path/within/snapshot/foo.csv -t /tmp/
