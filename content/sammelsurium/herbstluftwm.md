---
title: herbstluftwm
date: 2014-02-21T20:08:43
tags: 
- Software
- herbstluftwm
---

#### Resolution

Setze Auflösung für Dell Monitor zuhause

``` bash
xrandr --output DP1 --primary
xrandr -s 1920x1200
herbstclient detect_monitors
```

#### Auflösung und multiple Bildschirme

```
detect_monitors
list_monitors
focus_monitor
...
```