---
title: "Der Docker-Netzwerk Zwischenfall"
date: 2020-03-24T15:00:35+01:00
tags:
- Docker
- Netzwerk
---

"Die AWS Umgebung steht! Die Container Registry kann die Docker Images nicht
bereitstellen!". Anstelle eines heißen Kaffees sollte mir morgens erst einmal
der Schweiss auf der Stirn stehen.

<!--more-->

Tags zuvor stellte ich [Nexus](https://sonatype.com/) auf Docker um. Ich
hatte gut getestet, alles funktionierte. Alles - bis auf das Netzwerk?

```
CannotPullContainerError: Error response from daemon: Get
https://<nexus>/v2/: net/http: request canceled while waiting for
connection (Client.Timeout exceeded while awaiting headers)
```

Was war passiert? Soweit ich mich erinnerte, hatte ich netzwerkseitig keine
Änderungen vorgenommen? Durch `ifconfig` wurde ich auf ein Interface
aufmerksam das neu war.

```
br-82dd056a2d72: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 172.17.0.1  netmask 255.255.0.0  broadcast 172.17.255.255
        ether 02:42:dc:61:ec:ea  txqueuelen 0  (Ethernet)
        RX packets 2214  bytes 6816264 (6.5 MiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 2920  bytes 478870 (467.6 KiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0 
```

Fuck. Die IP-Range kam mir etwas zu bekannt vor. In der AWS Umgebung
nachgesehen um sicherzugehen. Auch dort benutzte ich die `172.17.x.x/24` und
diese liegt im selben Bereich wie `172.17.0.1/16` welche sich der Docker
Daemon erwählt hat.

## Theorie

Wenn man in der Berufsschule eines lernt, dann das IP Konflikte nie gut sind.

{{< figure src="/uploads/2020/03/conflict.jpg" >}}

Der Traffic von AWS zum Nexus System funktioniert zwar, aber vom System kommt
keine Antwort. Der Grund: IP Paket für 172.17.x.x kommt an, Antwort auf jenes
Paket durchläuft aber den lokalen Routing Table und wir an das Interface
`br-82dd056a2d72` geschickt.

```
$ route -n
Kernel IP routing table
Destination  Gateway  Genmask      Iface
172.17.0.0   0.0.0.0  255.255.0.0  br-6cfdc8482c79
```

**Die Chance eben dieses Subnet zu treffen lag übrigens bei 1:65536**.
So viele `/16` Subnetze kann man im v4 Adressbereich theoretisch
herausschneiden. Aber das nur nebenbei erwähnt.

Ähnlich alternativlos wie die derzeitigen Ausgehbegrenzungen wegen COVID-19
war auch die Lösung des Problems. Mein lokaler Docker Daemon brauchte einen
eigenen, unbenutzen(!) Adressbereich.

## Docker Daemon

Zur Kommunikation zwischen den Containern legt Docker ein neues Netzwerk
`docker0` an. Wie man dieses umkonfiguriert zeigt die
[Dokumentation](https://docs.docker.com/network/bridge/)

``` json
$ cat /etc/docker/daemon.json
{
  "bip": "192.168.1.5/24",
  "fixed-cidr": "192.168.1.5/25"
}
```

## Docker Compose

Docker Daemon war soweit entschärft, aber mein eigentliches Problem entsprang
einer anderen Quelle. Docker Compose baut eigene Netzwerke. So müssen auch
diese konfiguriert werden. `docker-compose.yaml`:

``` yaml
networks:
  nexus:
    driver: bridge
    ipam:
      driver: default
      config:
      - subnet: "192.168.6.0/24"
```

Dem Treiber `bridge` lassen sich Konfigurationsdetails im `yaml` Format
mitgeben. \o/

Meine beiden Docker Interfaces haben somit keine Schnittmenge mit dem AWS
Umfeld mehr. Setup wieder Up!
