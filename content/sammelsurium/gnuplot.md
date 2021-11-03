---
title: "gnuplot"
date: 2021-11-03T09:15:18+01:00
tags:
- Software
---

<!--more-->

## Einfaches Barchart

Die Daten (zum veranschaulichen) im File (`f.txt`)

```
1995 417
1996 547
1997 803
1998 925
1999 907
2000 066
2001 830
2002 729
2003 734
2004 742
2005 699
2006 605
2007 563
2008 546
2009 630
2010 447
2011 372
2012 614
2013 508
2014 578
2015 560
2016 501
2017 33
```

Das Gnuplot File `plot.gp`

```
set boxwidth 0.5
set style fill solid
plot "f.txt" using 1:2:xtic(1) with boxes
```

Aufruf unter macOS (das `-p` ist wichtig, sonst geht das Fenster gleich
wieder zu)

```
gnuplot -p plot.gp
```

Ergebnis

{{< figure src="/uploads/2021/11/gnuplot.png" >}}
