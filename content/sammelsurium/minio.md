---
title: "Minio"
date: 2019-07-18T14:06:14+02:00
draft: true
---

Minio ist ein S3 kompatibler Object Store

Neue Minio Cluster zu `mc` hinzufügen

    mc config host add ppdbp http://example.com:9000 access_key secret_key --api S3v4

Alle Buckets in einem Minio Cluster anzeigen

    mc ls myminio/

Neuen User für Bucket anlegen

    mc admin user add ppdbp access secret readwrite

Neuen Bucket erstellen

    mc mb --region=eu-west-1 ppdbp/bucketname

Bucket löschen (mit Inhalt)

    mc rb --force myminio/bucketname

Alle konfigurierten Cluster ansehen

    mc config host list
