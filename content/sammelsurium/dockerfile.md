---
title: Dockerfile
date: 2021-01-14T09:34:31
tags:
- Software
- docker
- dockerfile
---

Um einen Container zu bauen wird ein Dockerfile zu einem Image gebaut und das
Image kann dann ausgeführt werden zum laufenden Container.

<!--more-->

## Beispiel

``` Docker
FROM jfloff/alpine-python:3.8-onbuild

RUN apk add --update gnupg

# Install App
RUN mkdir /app
WORKDIR /app
COPY teamvaultexport.py /app/
RUN chmod +x /app/teamvaultexport.py

# Create Backup folder
RUN mkdir -p /var/backups/teamvault/

# Run Entrypoint
COPY entrypoint.sh /entrypoint.sh
CMD ["/entrypoint.sh"]
```

## Pakete unter Debian/Ubuntu

Um die Installation so gut wie möglich auf einen Automatisierten Prozess
umzubiegen empfiehlt sich dieses `RUN` Statement.

``` Docker
RUN set -x \
	&& DEBIAN_FRONTEND=noninteractive apt-get update --quiet \
	&& DEBIAN_FRONTEND=noninteractive apt-get install --quiet --yes --no-install-recommends \
	apt-transport-https \
	build-essential \
	ca-certificates \
	curl \
	gettext \
	git \
	ssh \
	libffi-dev
```

## ARG vs ENV

`ARG` ist nur zur Build-Zeit (siehe unten)

`ENV` ist auch im laufenden Container vom Image eingebettet.

## Bauen

Ein Dockerfile bauen und somit zum Image machen

    docker build -t r-devel .
