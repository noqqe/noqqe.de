---
title: R Graphs
date: 2013-03-31T12:50:57
tags: 
- Programming
- R
---

``` R
par(las=2)
## create the barchart with colorschema
barplot(d[,2], cex.names=0.8, names.arg=d[,1], main="Stromverbauch",
          ylim=c(0,max(d[,2])), xlab="Date", col=mintcolor)
```

## ggplot2

#### Barchart

``` R
ggplot(d, aes(x=factor(d[,1], levels=d[,1]),y=Verbrauch, colour=Verbrauch, fill=Verbrauch)) +
       geom_bar(stat="bin") +
       xlab("Monat") +
       ylab("Verbrauch")
```

#### LineChart

``` R
ggplot(d, aes(x=factor(d[,1], levels=d[,1]),y=x[,2][2:length(x[,2])])) +
          geom_line(aes(group=1)) +
          geom_point() +
          xlab("Monat") +
          ylab("Verbrauch")
```