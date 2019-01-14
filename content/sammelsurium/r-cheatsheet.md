---
title: R Cheatsheet
date: 2013-03-27T19:37:01
tags:
- Programming
- R
---

## Statistics Functions

```
rnorm(100) ## Random Normal Distribution
mean(x) ## Durchschnitt
sd(x) ## Standard Deviation
lm(x[,2] ~ x[,1] + x[,3]) ## Regression for Prediction of a column based on 1 and 3
seq(100) ## zähle bis 100
rep(98,4) ## wiederholt werte (repeat)
all(x > 8) ## Makro für IF Conditions, wenn ALLE dann
any(x > 8) ## wenn IRGENDEINS dann
sqrt(9) ## Wurzel ziehen
round(12.2) ## Runden auf nächsten Wert
head(x) ## einfach wie in Unix
tail(x) ## same here
subset(x,x > 5) ## Filtern innerhalb Vectoren nach Conditions
which(x,x > 5) ## Selbes, gibt aber die relativen positionen der Werte aus.
ifelse($condition, if-action, else-action) ## if
diff(x) ## berechnet die unterschiede innerhalb nummerischer Vektoren
sign(x) ## "begradigt" negative und positive werte
length(x) ## länge von vectoren abfragen
sort(x)
order(x) ## sortieren mit realtiven index angaben als return
```

## Import Funktions

```
read.table("bla.csv",header=FALSE) ## Import von CSV Dateien
read.table("bla.txt",header=FALSE) ## Import von TXT
scan(tf,"") ## import von textdateien
```

## R Internal Functions for data

```
str(x) ## welche Struktur hat der Vector?
mode(x) ## was für eine Struktur liegt vor?
summary(x) ## erzähl mir alles was so geht über x?
print(x) ## gib mir die print methode für den jeweiligen Datentypen
head(x) ## Zeigt den Header an
class(x) ## zeigt die Class an zb. "data.frame"
matrix(x,y) ## erstelle Matrix
apply(x,y,z) ## funktion für jeden Wert in vecotr ausführen
lapply(x,y,z) ## selbes für liste
sapply(x,y) ##  für liste, aber rückgabe als vector
cbind(matrix,vector) ## column an matrix binden
rbind(matrix,vector) ## row an matrix binden
dim(matrix) ## columns und rows anzahl einer matrix ausgeben
nrow(x) ##  nur rows
ncol(x) ## nur columns
attributes(x) ## attribute einer klasse abfragen (zb. dim bei matrix)
as.matrix
as.numeric
as.character
as.factor
colnames(matrix) ## columns namen geben (von matrix)
rownames(matrix) ## für rows
names(j) ## beschreibungen ausgeben bei listen
unlist(j) ## liste zu vector konvertieren
class(j) ## classe abfragen
unname(j) ## alle namen entfernen
```

## Misc Functions

```
source("file.R") ## Include
pdf("out.pdf") ## PDF File Output definieren
hist(x) ## Histogram vom Vector x erstellen
dev.off() ## PDF schreiben
data() ## Mitgelieferte Datensätze anzeigen
help()
help.search("normal distribution") ## Google like search
?print ## alternative
?"print" ## online manual
source("file.R") ## incldue
library("libaryname",quietly=T) ## lib nachladen und Fresse halten
suppressPackageStartupMessages(library("libaryname",quietly=T)) ## und dann auch noch wirklich die fresse halten!
```

## Befüllen der Variablen

```
x <- c(1,2,4) ## Itegers
z <- paste("foo", "bar") ## Strings
m <- rbind(c(1,4),c(2,3)) ## Matrix
l <- list(drinks=c("Mate","Coffee"),consumption=c(23,42))
d <- data.frame(list(drinks=c("Mate","Coffee"),consumption=c(23,42))) ## data Frames
```

## Variablen Bearbeitung

```
x ## Vector ausgeben
print(x) ## Vector ausgeben
x[3] ## 3. Variable des Vectors ausgeben
x[2:3] ## Range ausgeben
m[2,1] ## Matrix specific
```

## Funktionen

Funktionsdefinition

```
oddcount <- function(x) {
  k <- 0
  for (n in x) {
    if (n %% 2 == 1 ) k <- k+1
  }
  return(k)
}
```

## Data Structures

#### Vectors

```
> x <- c(1,2,4,5)
> x
[1] 1 2 4 5
```

#### Character Strings

```
> y <- "aber"
> y
[1] "aber"
> y <- paste("foo", "bar")
> y
[1] "foo bar"
```

#### Matrices

Matrix-artige Datenstruktur - rbind = row bind

```
> m <- rbind(c(1,4),c(2,3))
> m
     [,1] [,2]
[1,]    1    4
[2,]    2    3
> m[,1]
[1] 1 2
> m[1,1]
[1] 1
> m[2,1]
[1] 2
> m[2,1]
```

#### Lists

Gut für mehrere versch. Datentypen

```
> l <- list(u=2, v="asdf")
> l
$u
[1] 2
$v
[1] "asdf"
> mode(l)
[1] "list"
> l[2]
$v
[1] "asdf"
> l$u
[1] 2
>
```

#### Dataframes

```
> d <- data.frame(list(drinks=c("Mate","Coffee"),consumption=c(23,42)))
> d
  drinks consumption
1   Mate          23
2 Coffee          42
```

Sehr gutes Beispiel für DataFrames

```
> people = data.frame (
age = c(32,34,12,41,18,23,43,22,19,24),
height = c(177,166,165,174,156,184,191,179,182,180),
sex = c('M','F','M','M','F','M','F','F','M','F'))

> people
   age height sex
1   32    177   M
2   34    166   F
3   12    165   M
4   41    174   M
5   18    156   F
6   23    184   M
7   43    191   F
8   22    179   F
9   19    182   M
10  24    180   F
```

#### Classes

Es ist kompliziert :P