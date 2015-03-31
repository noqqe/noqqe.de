---
layout: post
title: "Tabellen aus R mit knitr und TeX"
date: 2013-08-03 13:21
comments: true
categories:
- ubuntuusers
- osbn
- R
- Stats
- Statistik
- TeX
- LaTeX
- PDF
- Markdown
---

Eines der Dinge die `R` wirklich gut kann ist Daten darstellen. Gerade im Interactive Mode.

{% img center /uploads/2013/08/termdata.png %}

Konkret migriere ich gerade meine Graphen/Statistiken in ein PDF.
Um mit `R` portablen Output zu generieren ist [knitr](http://yihui.name/knitr/)
zum DeFacto Standard geworden. Möglich sind zum Beispiel
[Markdown](http://daringfireball.net/projects/markdown/) und
[TeX](https://en.wikipedia.org/wiki/TeX). Allerdings funktioniert dort die
einfache `print()` Methode für Daten nicht so gut, dafür gibts das Plugin [xtable](http://cran.r-project.org/web/packages/xtable/index.html).

Man muss ja zum Glück kein TeX Gott sein um ein kleines Template zu ergoogeln.
Wichtig ist nur der Teil am Ende. Das Einbetten von R Code in das Dokument.

{% codeblock lang:tex Stats.Rnw %}
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
{% endcodeblock %}

Das Plugin `xtable` versteht so ziemlich alle Datentypen
und generiert natives TeX.

Um den R Teil im Dokument auszuführen wird das so-called `.Rnw` File durch R gejagt.
Es ensteht pures TeX. Für die komplette Verarbeitung nutze ich ein R Script.

{% codeblock lang:r create.R %}
#!/usr/bin/Rscript

# loading libraries
library(xtable)
library(knitr)

# Rnw to TeX
knit("Stats.Rnw", out="tmp/Stats.tex" )

# create and open pdf via system call
system("pdflatex -output-directory tmp/ Stats.tex && evince tmp/Stats.pdf")
{% endcodeblock %}

Der Output:

{% img center /uploads/2013/08/texdata.png %}
