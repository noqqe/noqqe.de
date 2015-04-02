---
date: 2009-12-02T18:37:16+02:00
type: post
slug: bash-timerobot-0-0-8
comments: true
title: bash | timeRobot 0.0.8
aliases:
- /archives/769
categories:
- Coding
- Linux
tags:
- backup
- bash
- debian
- shell
- timemachine
- timerobot
- ubuntu
---

Nach den [Erkenntnissen über rsyslogd und logger](/?p=686), hab ich das Logging von timeRobot überarbeitet.

Wo früher stand:

``` bash
echo `date +%d-%m-%Y-%H:%M:%S` --- NEW Action ---  >> $LOGFILE
echo `date +%d-%m-%Y-%H:%M:%S` TimeRobot-Verzeichnis $PATHNAME hinzugefuegt >> $LOGFILE
echo `date +%d-%m-%Y-%H:%M:%S` --- END Action ---  >> $LOGFILE
```

Steht jetzt:

```
logger -p local0.info -t TIMEROBOT added $PATHNAME directory
```

Die Umstellung hab ich an allen Enden an denen timeRobot was ausspuckt eingefügt/ersetzt. Das macht den Source schöner und mit den Fehlermeldungen lässt sich leichter weiterarbeiten. Bzw lassen Sie sich dadurch auch unterdrücken :)

Der komplette Source als TextFile zum anschauen hier: [http://zwetschge.org/projects/timerobot/timerobot-0.0.8/timerobot.source](http://zwetschge.org/projects/timerobot/timerobot-0.0.8/timerobot.source)

Als Debian-Paket hier: [http://zwetschge.org/projects/timerobot/timerobot-0.0.8/timerobot_0.0.8-all.deb](http://zwetschge.org/projects/timerobot/timerobot-0.0.8/timerobot_0.0.8-all.deb)
