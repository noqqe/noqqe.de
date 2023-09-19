---
title: "Nexus"
date: 2023-09-19T13:12:48+02:00
tags:
- Software
---

Sonatype Nexus Repository Manager 

<!--more-->

## Blob Repository

Ein File herunterladen

    curl -vL -u user:pw -X GET "https://nexus.acme.com/repository/acme-blobs/<file>

Ein File hochladen

    curl -vL -u user:pw --upload-file $x  https://nexus.acme.com/repository/acme-blobs/<file>

Ein File löschen

    curl -L -u user:pw -X DELETE https://nexus.acme.com/repository/acme-blobs/<file>
