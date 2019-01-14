---
title: s3cmd
date: 2016-04-04T13:13:46
tags: 
- Software
- S3
- Amazon
- s3cmd
---

List all buckets

    s3cmd la

List files within one bucket

    s3cmd ls s3://com.example.int.backup/

Download a file

    s3cmd get s3://BUCKET/OBJECT LOCAL_FILE

Upload a file

    s3cmd put FILE [FILE...] s3://BUCKET[/PREFIX]

Delete a file

    s3cmd del s3://com.example.int.backup/git_bitbucket_saved.tar.gz