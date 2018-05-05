---
title: "ElasticSearch Monitoring"
date: 2018-05-05T09:20:48+02:00
tags:
- ElasticSearch
- Icinga2
---

Wir benutzen die gehostete Version des [elastic](https://www.elastic.co/)
Stack. Dort gibt es einige Möglichkeiten via
[Watcher](https://www.elastic.co/about/press/elastic-introduces-watcher-alerting-for-elasticsearch)
Details des Clusters monitoren und ggf. an Menschen Emails zu verschicken. Was
damit nicht geht ist den verbleibenden Speicherplatz monitoren.

{{< figure src="/uploads/2018/05/devops.jpg" >}}

Das musste irgendwie gelöst werden und ich habe ein Icinga2 Plugin geschrieben
um eben jenes zu bewerkstelligen

```
#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import sys
import requests
import humanfriendly

warn = sys.argv[1]
crit = sys.argv[2]
api = sys.argv[3]
max_size = sys.argv[4]
sum_bytes = 0

res = requests.get(api).text
indices = res.split("\n")

# go throug all lines of result
# get size and convert to bytes
# collect everything in sum_bytes
for index in indices:
    try:
        size = index.split()[9]
        size = humanfriendly.parse_size(size)
        sum_bytes = sum_bytes + size
    except (TypeError, IndexError):
        pass

# calculate
max_size_bytes = humanfriendly.parse_size(max_size)
human_sum = humanfriendly.format_size(sum_bytes)
human_max = humanfriendly.format_size(max_size_bytes)

# calc percentage
filled = 100 * float(sum_bytes)/float(max_size_bytes)

# return value
print("%.2f%% used (%s/%s) | percentage=%.2f; used=%s; max=%s;" %
  (filled, human_sum, human_max, filled, sum_bytes, max_size_bytes))

# check if crit
if filled > float(crit):
    sys.exit(2)

# check if warn
if filled > float(warn):
    sys.exit(1)
```

Es fordert im Endeffekt die Liste aller Indices und deren Size an, welche ich
dann mit der tollen Library `humanfriendly` in Bytes konvertiere und
zusammenrechnen kann.

```
> humanfriendly.parse_size('140g')
> 140000000000
>
> humanfriendly.format_size(150000)
> '150 KB'
```

Das Script ist natürlich total spartanisch, aber es tut was es soll. Um das
angeben der Max Size komme ich nicht herum, da das Cluster leider nicht unter
meiner Kontrolle steht.

```
$ ./check_elastic_storage.py <url>/_cat/indices?v 80 90 1TB
> 50% used (500 GB/1 TB) | percentage=50; used=500000000000; max=1000000000000;
```

In Icinga2 einbauen kann glaube ich jeder selbst, die Doku dafür ist ja
ziemlich schick :) Habe mir auch mal Icinga2 Plugin Libraries angesehen für
Python, tatsächlich war mir das für ein bisschen Return Code und formatierte
Ausgabe zu viel Overhead. Falls ihr aber Empfehlungen habt welche der 80 Libs
man davon benutzen sollte, immer her damit.
