---
title: R JSON
date: 2014-01-15T16:40:30
tags: 
- Programming
- R
---

## Installing required library

    install.packages("rjson")
    library("rjson")

## Using it

Creating the file (downloaded from somewhere) and assing the name to a value and read it with readChar

~~~
f <- '/tmp/pin.json'
j <- readChar(f,file.info(f)$size)
~~~

after that you can use this string object for the library

    fromJSON(j)
    x <- fromJSON(json)

resulting in a list object
