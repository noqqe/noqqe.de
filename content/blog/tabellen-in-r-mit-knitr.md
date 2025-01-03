---
date: '2013-08-03T11:21:00'
tags:
- development
- latex
- markdown
- stats
- tex
- r
- pdf
- statistik
title: Tabellen aus R mit knitr und TeX
---

Eines der Dinge die `R` wirklich gut kann ist Daten darstellen. Gerade im
Interactive Mode.

{{< figure src="/uploads/2013/08/termdata.png" >}}

Konkret migriere ich gerade meine Graphen/Statistiken in ein PDF.  Um mit
`R` portablen Output zu generieren ist [knitr](http://yihui.name/knitr/)
zum DeFacto Standard geworden. Möglich sind zum Beispiel
[Markdown](http://daringfireball.net/projects/markdown/) und
[TeX](https://en.wikipedia.org/wiki/TeX). Allerdings funktioniert dort die
einfache `print()` Methode für Daten nicht so gut, dafür gibts das Plugin [xtable](http://cran.r-project.org/web/packages/xtable/index.html).

Man muss ja zum Glück kein TeX Gott sein um ein kleines Template zu ergoogeln.
Wichtig ist nur der Teil am Ende. Das Einbetten von R Code in das Dokument.

``` tex
\documentclass[a4paper]{article}
\usepackage[british]{babel}
\begin{document}
\title{Quantifying Myself}
\author{Florian Baumann}
\maketitle\thispagestyle{empty}

\section{Runkeeper}

<<Biking_Table, echo=FALSE, results="asis">>=
source(file="../Sports/.R")
xtable(tail(a))
@

\end{document}
```

Das Plugin `xtable` versteht so ziemlich alle Datentypen
und generiert natives TeX.

Um den R Teil im Dokument auszuführen wird das so-called `.Rnw` File durch R gejagt.
Es ensteht pures TeX. Für die komplette Verarbeitung nutze ich ein R Script.

``` r
#!/usr/bin/Rscript

# loading libraries
library(xtable)
library(knitr)

# Rnw to TeX
knit("Stats.Rnw", out="tmp/Stats.tex" )

# create and open pdf via system call
system("pdflatex -output-directory tmp/ Stats.tex && evince tmp/Stats.pdf")
```

Der Output:

{{< figure src="/uploads/2013/08/texdata.png" >}}
