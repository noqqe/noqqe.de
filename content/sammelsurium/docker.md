---
title: Docker
date: 2018-01-09T12:28:18
tags:
- Software
- docker
---

Dockerfile Beispiel

``` Docker
FROM rocker/r-devel:latest

[more changes...]
```

Ein Dockerfile bauen und somit zum Image machen

    docker build -t r-devel .

Docker Image starten

    docker run -t r-devel

Docker Images ansehen

```
> docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
hello-world         latest              fce289e99eb9        8 days ago          1.84kB
r-devel             latest              14efdbbdfc99        2 weeks ago         3.75GB
rocker/r-devel      latest              14efdbbdfc99        2 weeks ago         3.75GB
```

Docker Image mit Port Mappings

    docker run -p 80:80 -p 1220:22 -p 1109:109 -d r-devel

Docker Stop all instances

    docker stop $(docker ps -a -q)

Docker Login to Nexus

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

Delete all containers

    docker rm $(docker ps -a -q)

Delete all images

    docker rmi $(docker images -q)

## Docker Compose

Docker Network zu anderen docker-compose Installationen auf der selben
Instanz. Mittels externe Netzwerke für Beispielsweise einen HTTP Ingress
Reverse Proxy

``` yaml
services:
  tls:
    image: caddy
    restart: unless-stopped
    networks:
      - retro
      - status
    ports:
      - "80:80"
      - "443:443"

networks:
  status:
    external:
      name: cachet-docker_cachet
  retro:
    external:
      name: retro-board_retro
```
