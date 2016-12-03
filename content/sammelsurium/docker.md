---
title: Docker
date: 2014-09-22T12:28:18
tags: 
- Software
- docker
---

Docker Build multiple Instances

    for x in {1..5}; do docker build -t devnull$x . ; done

Run multiple instances

    for x in {1..5} ; do docker run -p 8$x:80 -p 1220$x:22 -p 1109$x:9 -d devnull$x ; done

Docker Stop all instances

    docker stop $(docker ps -a -q)
