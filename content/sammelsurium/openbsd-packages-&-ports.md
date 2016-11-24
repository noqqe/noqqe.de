---
title: OpenBSD Packages & Ports
date: 2014-01-24T10:43:59.000000
tags: 
- OS
- OpenBSD
---


    export PKG_PATH=http://openbsd.cs.fau.de/pub/OpenBSD/5.3/packages/amd64/

## Alternatvie GCC

Für zum Beispiel taskwarrior

    rm -rf CMakeFiles/
    rm -rf CMakeCache.txt
    cmake . -DCMAKE_CXX_COMPILER=eg++
    make
