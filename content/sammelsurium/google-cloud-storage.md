---
title: "Google Cloud Storage"
date: 2021-05-27T15:24:49+02:00
tags:
- Software
- GCP
---

Wie AWS S3, nur mit anderen Commands.

<!--more-->

Retrieve a banner image from a publicly accessible Cloud Storage location:

    gsutil cp gs://cloud-training/gcpfci/my-excellent-blog.png my-excellent-blog.png

Copy the banner image to your newly created Cloud Storage bucket:

    gsutil cp my-excellent-blog.png gs://$DEVSHELL_PROJECT_ID/my-excellent-blog.png

Modify the Access Control List of the object you just created so that it is readable by everyone:

    gsutil acl ch -u allUsers:R gs://$DEVSHELL_PROJECT_ID/my-excellent-blog.png

