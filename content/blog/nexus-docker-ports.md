---
title: "Nexus Docker Registry Ports anzeigen"
date: 2023-06-26T12:05:40+02:00
tags:
- Docker
- Nexus
---

@ Work betreiben wir ein relativ großes Nexus. Für jede gehostete Docker
Registry wird in Nexus ein neuer Port benötigt. 

Da es leider nicht möglich ist, sich alle bereits vergebenen Ports anzeigen
zu lassen, habe ich dafür mal die API bemüht um diese inkl Protokoll dann
auszugeben.

<!--more-->

```python
#!/usr/bin/env python3

# Usage
# ./show-docker-ports.py https://adminuser:adminpass@nexus.example.net

import sys
import json
import requests

base_url=sys.argv[1]
repolist=base_url + '/service/rest/v1/repositories/'

rl = json.loads(requests.get(repolist).content)

for repo in rl:
    if repo["format"] == 'docker':
        repourl = base_url + "/service/rest/v1/repositories/docker/" + repo["type"] + "/" + repo["name"]
        detail = json.loads(requests.get(repourl).content)

        if detail["docker"]["httpPort"] is not None:
            print(detail["docker"]["httpPort"], detail["name"], "http")

        if detail["docker"]["httpsPort"] is not None:
            print(detail["docker"]["httpsPort"], detail["name"], "https")
```

Und der Output sieht dann beispielsweise so auszugeben

```
8083 docker-sr-hosted http
8483 docker-sr-hosted https
8082 docker-sr-group http
8482 docker-sr-group https
8884 docker-fo-hosted https
8885 docker-fo-group https
8486 docker-rot-hosted https
```
