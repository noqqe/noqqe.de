---
title: Docker
date: 2018-01-09T12:28:18
tags:
- Software
- docker
---

Ein kleines Cheatsheet für Docker

<!--more-->

## Zum laufenden Container verbinden

Je nach Image...

    docker exec -it <name> /bin/bash
    docker exec -it <name> /bin/sh

## Neuen Container

Interaktiven Container starten

    docker run --rm ubuntu

Verbinden

    docker run -it --rm google/cloud-sdk:alpine /bin/bash

Docker Image mit Port Mappings

    docker run -p 80:80 -p 1220:22 -p 1109:109 -d r-devel

## Docker Images

Delete all images

    docker rmi $(docker images -q)
    docker images prune -a

Images die nicht mehr referenziert sind älter als 6 Stunden löschen.

    docker image prune -a -f --filter "until=6h"

## Networking

Alle netzwerke entfernen. Manchmal ist ein Bridge Device Stuck

    docker volume prune -f

## Custom Registry Image Upload

Dafür muss mann das Repo anlegen, einen User+Role und dann nen HTTPS Port
freigeben

``` bash
> docker login nexus.acme.com:8087
Username: max-docker
Password:
Login Succeeded
```

Upload Docker image to Nexus

``` bash
docker tag 916a0128c7e4 nexus.acme.com:8087/library/r35:0.0.1
docker push nexus.acme.com:8087/library/r35:0.0.1
```

## Disk Space

Wenn man sehen will wie viel Space Docker Installationen verbrauchen

    docker system df

## Logs

Es kann passieren dass der Docker Container viel loggt.

Log truncaten

```
echo "" > $(docker inspect --format='{{.LogPath}}' <name>)
```

Und für Produktion empfielt sich dieses Limit auch gleich via den Daemon
automatisch einzubauen.

```json
{
  "log-driver": "json-file",
  "log-opts": {"max-size": "50m", "max-file": "3"}
}
```
