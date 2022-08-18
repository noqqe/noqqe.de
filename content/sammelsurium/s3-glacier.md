---
title: "AWS S3"
date: 2022-08-18T10:23:07
tags:
- AWS
- S3
---

Ein paar Kniffe und Tricks für AWS S3 Glacier

<!--more-->

## Alle Objekte eines Buckets in Glacier finden

```bash
aws s3api list-objects \
  --bucket <bucketname> --query 'Contents[?StorageClass==`GLACIER`]'`
```

## Aus Glacier restoren

```bash
aws s3api restore-object --bucket <bucketname> --key path/to/file \
  --restore-request '{"Days":60,"GlacierJobParameters":{"Tier":"Standard"}}'
```

## Restore Progress anzeigen

Das wichtige Feld ist hier `Restore`

```json
aws s3api head-object --bucket <bucketname> --key path/to/file
{
    "AcceptRanges": "bytes",
    "Restore": "ongoing-request=\"true\"",
    "LastModified": "2022-05-13T09:45:39+00:00",
    "ContentLength": 73340,
    "VersionId": "EtPT875z8WWtmjimqHZ_BU",
    "ContentType": "image/png",
    "BucketKeyEnabled": true,
    "StorageClass": "GLACIER"
}
```


