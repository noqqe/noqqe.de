---
title: "Nexus Repository Storage analysieren"
date: 2019-11-18T15:54:26+01:00
tags:
- Nexus
- Python3
- Groovy
---

Aus finanziellen Gründen finde ich mich ab und an in der Verlegenheit wider,
eine [Nexus](https://www.sonatype.com/nexus-repository-oss) Instanz zu betreiben.

Oft ist bereits der Speicherplatz zur Neige gegangen (bitte mit der Anno 1604
Sprecher Stimme vorstellen). Was folgte war meist zielloses herumgestocher
und [Clean Up Policies](https://help.sonatype.com/repomanager3/cleanup-policies) tunen.
Das ist alles andere als geil, wenn man mehrere Repos hat. Ich will
eigentlich einfach nur eine Übersicht, welches Repo wie viel Storage
verbraucht.

In der offiziellen Dokumentation wird ein Groovy Script referenziert, welches
dann ein JSON Dokument ausspuckt, wenn man es als Task in der Nexus Admin Oberfläche angelegt hat.

{{< figure src="/uploads/2019/11/nexus.png" >}}

Dabei heraus kommt eine Datei in ungefähr folgender Struktur.

``` json
{
  "docker": {
    "repositories": {
      "docker-xxx-hosted": {
          "reclaimableBytes": 3460,
          "totalBytes": 193468724
      },
      "docker-yyy-hosted": {
          "reclaimableBytes": 1385881,
          "totalBytes": 46483379009
      },
      [...]
}
```

Um das jetzt aber irgendwie lesen zu können, habe ich mich kurz hingesetzt und ein
Python Script gebaut.

```python
#!/usr/bin/env python3

import sys
import json
import math

space = []

def convert_size(size_bytes):
   if size_bytes == 0:
       return "0B"
   size_name = ("B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB")
   i = int(math.floor(math.log(size_bytes, 1024)))
   p = math.pow(1024, i)
   s = round(size_bytes / p, 2)
   return "%s %s" % (s, size_name[i])

with open(sys.argv[1]) as json_file:
    data = json.load(json_file)


for repotype in data:
    for repo in data[repotype]['repositories']:
        space.append([data[repotype]['repositories'][repo]["totalBytes"],
          data[repotype]['repositories'][repo]["reclaimableBytes"], repo, repotype])

for entry in sorted(space, key=lambda x: float(x[0]), reverse=True):
    print('{} ({} reclaimable) {} ({})'.format(convert_size(entry[0]),
      convert_size(entry[1]), entry[2], entry[3]))
```

Falls noch jemand das Problem einer zu vollen Nexus Instanz hat und nicht
weiss welchen der Kollegen man nun hauen muss, kann man dieses Skript
benutzen um den Übeltäter zu finden.

{{< figure src="/uploads/2019/11/result.png" >}}

Ich habe absichtlich keine Libraries wie `hurry` oder `humanreadable` zur
Umrechnung der Bytes benutzt um möglichst portabel zu sein.
