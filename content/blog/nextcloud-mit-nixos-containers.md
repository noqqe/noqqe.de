---
title: "Nextcloud mit NixOS Containern"
date: 2021-02-06T08:57:32+01:00
tags:
- NixOS
- Podman
- Docker
- Nextcloud
---

Ich nutze schon länger eine kleine [Nextcloud](https://nextcloud.com) Instanz
um hier und da ein paar kleine Dateien zwischen verschiedenen Endgeräten im
Sync zu halten.

Wie das nach ein paar Jahren so ist, wird aus "klein" aber irgendwann "doch
nicht mehr ganz so klein wie ich anfangs dachte, oh shit das kost ja Geld"

<!--more-->

Block-Storage ist teuer. 25GB davon kosten bei meinem Provider
[2,50$/Monat](https://www.vultr.com/products/block-storage/).

Nextcloud ist aber in der Lage Object Stores als Primär Storage zu verwenden.
Dieser muss AWS S3 Kompatibel sein. Bei
[backblaze.com](https://backblaze.com) kosten diese 25GB nur
[0,125$/Monat](https://www.backblaze.com/b2/cloud-storage.html). Das ist 20x
Billiger?[^1]

{{< figure src="/uploads/2021/02/ruff.png" >}}

## Design

Bevor ich jetzt aber vollends zum Buchhalter werde, lieber ein bisschen was
Technisches. Ich dachte, cool ich könnte das auch mit Containern machen und
wenn ich schon dabei bin unter [NixOS](https://nixos.org) und dann
auch gleich via [Podman](https://podman.io).

Warum Podman? Letztens gab es einen tollen Artikel darüber bei
[zwischenzugs.com: Goodbye Docker](https://zwischenzugs.com/2019/07/27/goodbye-docker-purging-is-such-sweet-sorrow/).

## NixOS

Zuerst habe ich mir überhaupt mal das Interface von NixOS zu Containern
angesehen. Die Doku ist hierfür wie immer wunderbar
[lesbar](https://search.nixos.org/options?channel=20.09&from=0&size=50&sort=relevance&query=oci-containers).

Um mal ein Gefühl dafür zu bekommen hier ein gekürzter Auszug:

```nix
{
  virtualisation.oci-containers = {
    containers = {
      ingress = {
        image = "library/caddy:2.2.1";
        ports = ["80:80" "443:443"];
        [...more...];
      };
      linx = {
        image = "andreimarcu/linx-server:version-2.3.8";
        [...more...];
      };
    };
  };
}
```

Jeder der schonmal mit `docker-compose` oder ähnlichem gearbeitet hat, kommt
das bekannt vor.

## Podman

Ich will aber jetzt lieber Podman anstelle des Docker Daemons. Die Container
Schnittstelle ist zum Glück von Haus aus gleich Backend-agnostisch designed,
sodass die Umstellung auf Podman nur eine einzige Config Option ist.

```nix
{
  virtualisation.podman.enable = true;
  virtualisation.podman.dockercompat = true;

  virtualisation.oci-containers = {
    backend = "podman";
    containers = {
      [...];
    };
  };
}
```

Das `dockerCompat` Setting macht dann zusätzlich einen `alias docker=podman`
in dein System, sodass man wirklich **garnichts** mehr umlernen muss.

## System Interaktion

NixOS baut mir aus den obigen Definitionen somit also ein kleines Container
Setup auf. Ich kann mir diese natürlich (mit `docker` alias) anzeigen lassen:

```
# docker ps
CONTAINER ID  IMAGE                                            COMMAND               PORTS               NAMES
3c8e5f6505a0  docker.io/library/caddy:2.2.1                    caddy run --confi...  0.0.0.0:80->80/tcp  ingress
11a11f9d66ab  docker.io/andreimarcu/linx-server:version-2.3.8  -config /data/lin...                      linx
```

Jeder Container endet in einem schön generierten Podman `systemd` Service den
ich hauptsächlich nutze um mit dem Containern zu interagieren.

```
# systemctl status podman-linx.service
# journalctl -fu podman-ingress.service
```

Nice!

## Nextcloud

Zur eigentlichen Sache zurück. Ich will ja schliesslich Geld sparen mit
Nextclouds Object Store Backend.

Der Nextcloud Container wird genauso initialisiert wie alle anderen auch, nur
mit einem Set an Umgebungsvariablen. Datenbank, Initiales Nextcloud Passwort
und natürlich der Object Store.

```nix
{
  virtualisation.podman.enable = true;
  virtualisation.podman.dockercompat = true;

  virtualisation.oci-containers = {
    backend = "podman";
    containers = {
      nextcloud = {
        image = "nextcloud:20.0.4";
        extraOptions = [
          "--ip=10.88.10.12"
          "--add-host=db:10.88.10.13"
        ];
        volumes = [
          "/data/nextcloud/docroot:/var/www/html"
        ];
        environment = {
          POSTGRES_HOST = "db";
          POSTGRES_PORT = "5432";
          POSTGRES_DB = "nextcloud";
          POSTGRES_USER = "";
          POSTGRES_PASSWORD = "";
          TRUSTED_PROXIES = "10.88.10.10";
          NEXTCLOUD_ADMIN_USER = "admin";
          NEXTCLOUD_TRUSTED_DOMAINS = "<domain> 10.88.10.*";
          NEXTCLOUD_ADMIN_PASSWORD = "";
          OBJECTSTORE_S3_HOST = "s3.us-west-001.backblazeb2.com";
          OBJECTSTORE_S3_BUCKET = "this-bucket-does-not-exist";
          OBJECTSTORE_S3_KEY = "";
          OBJECTSTORE_S3_SECRET = "";
          OBJECTSTORE_S3_PORT = "443";
          OBJECTSTORE_S3_SSL = "true";
          OBJECTSTORE_S3_REGION = "us-west-001";
          OBJECTSTORE_S3_USEPATH_STYLE = "false";
        };
      };

      # Postgres
      db = {
        image = "postgres:13.1";
        extraOptions = [
          "--ip=10.88.10.13"
        ];
        environment = {
          POSTGRES_DB = "postgres";
          POSTGRES_USER = "root";
          POSTGRES_PASSWORD = "";
        };
        volumes = [
          "/data/db/postgres:/var/lib/postgresql/data/"
        ];
      };


```

Den S3 Kompatiblen Store muss man vorher natürlich im Webinterface des
Anbieters (in meinem Fall Backblaze) anlegen und die
Credentials dann in die `_S3_` Slots einfügen.

Eine kleine `PostgreSQL` Datenbank brauche ich auch, welche ebenfalls als
Container läuft. Alles kein Hexenwerk.

## Der Rest

Zusammen mit [Caddy](https://caddyserver.com/) für SSL Zertifikate und
Webproxy ist nach einem `nixos-rebuild switch` die Konfiguration aktiv und
die Container fahren sich hoch und das Setup läuft.

Die alte Kiste konnte ich abschalten und jetzt habe ich sozusagen 500GB
Platz fürs gleiche Geld (2,50$).

Mittlerweile hoste ich noch ein paar andere Dinge, wie ein
[Miniflux](https://miniflux.app/) und ein [Linx Paste
Service](https://github.com/andreimarcu/linx-server) auf der Maschine,
einfach weil es so leicht ist.

## TL;DR

NixOS mit Podman Containern finde ich super schön gemacht.
Es fühlt sich an als hätte jemand `docker-compose` in mein
OS hineinintegriert. Das macht es super einfach und komfortabel zu bedienen.

Was ich nicht so toll finde ist das hantieren mit IPs und Hosts einträgen.
`docker-compose` setzt hier automatisch alle Container namen im selben
`network` als Hosts eintrag und ich kann (z.B. die Datenbank) mittels `db`
aus jedem anderen Container ansprechen. Diese Arbeit muss ich mir hier via
`--ip` in `extraOptions` selbst machen. Oder ich hab das richtige Setting
noch nicht gefunden. Hints welcome!

[^1]: Ja, ok noch ein bisschen Transferkosten 0,010$ im Monat, dann sagen wir
  eben 10x billiger im Extremfall.
