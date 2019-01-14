---
title: Hive
date: 2014-05-09T11:22:05
tags:
- Software
- Hadoop
---

Das Hive CLI Tool ist mittlerweile deprecated. Der Ersatz heisst irgendwas
mit Bees

typical operations

    hive -e "show databases;"
    hive -e "show tables;"
    hive -e "select * from hh limit 20;"

Turn on debug mode

    hive -hiveconf hive.root.logger=DEBUG,console
