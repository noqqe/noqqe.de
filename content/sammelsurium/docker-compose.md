---
title: docker-compose
date: 2021-01-14T09:21:26
tags:
- Software
- docker
---

Mit `docker-compose` lassen sich mehrere Docker Container in einem File
spezifizieren und dadurch eine "Deployment Definition" erstellen.

## Commands

```
docker-compose up -d
docker-compose down
docker-compose restart <servicename>
```

## Struktur

3 Basis Sektionen. Services, Volumes, Networks.

```yaml
version: '3'

services:
  tls:
    image: caddy:latest
    restart: always
    ports:
      - "80:80"

volumes:
  certs:
    driver: local

networks:
  mynetwork:
    driver: bridge

```

## Builds einbetten

Wenn man sein `Dockerfile` mit im selben Repo ausliefert kann man statt der
Definition eines `Image` auch einen Docker Buildstep definieren

```yaml
version: '3'

services:
  srv:
    build:
      context: teamvault/
```

Anwenden dann so:

```bash
docker-compose build --no-cache
# or
docker-compose up -d --build
```

## Volumes

Volumes können auf verschiedene Weisen deklariert werden

### Bind Volume

Schönste Variante, da man weiss wo die Nutzdaten liegen und sie auch
außerhalb vom Container nutzen kann.

```yaml
version: '3'

services:
  tls:
    image: caddy:latest
    volumes:
      - 'certs:/root/.caddy/'

volumes:
  certs:
    driver: local
    driver_opts:
      type: none
      device: "$PWD/certs"
      o: bind
```

### Direct Bind

hier ist es nicht nötig ein Bolume mit einem expliziten Ort zu verwenden

```yaml
services:
  tls:
    image: caddy:latest
    volumes:
      - './tls/Caddyfile:/etc/caddy/Caddyfile'
```

### Native Volume

Die (finde ich) beschissenste Variante. Die Nutzdaten hängen dann irgndwo
hinter einem ID Volume unter `/var/lib/docker/volumes` herum... und man
findet nichts nie wieder.

```yaml
version: '3'

services:
  tls:
    image: caddy:latest
    volumes:
      - 'certs:/root/.caddy/'

volumes:
  certs:
```

## Interaktion zwischen 2 Netzwerken

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
