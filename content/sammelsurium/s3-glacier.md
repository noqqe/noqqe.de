---
title: "AWS S3 Glacier"
date: 2022-08-18T10:23:07+02:00
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

## Glacier Objekt in "Standard" Class überführen

To overwrite the existing object with the Amazon S3 Standard storage class, run the following command:

    aws s3 cp s3://awsexamplebucket/dir1/example.obj s3://awsexamplebucket/dir1/example.obj --storage-class STANDARD

To perform a recursive copy for an entire prefix and overwrite existing objects with the Amazon S3 Standard storage class, run the following command:

    aws s3 cp s3://awsexamplebucket/dir1/ s3://awsexamplebucket/dir1/ --storage-class STANDARD --recursive --force-glacier-transfer


