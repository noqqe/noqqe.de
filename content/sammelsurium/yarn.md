---
title: "Yarn"
date: 2018-08-10T12:04:10+02:00
tags:
- Software
- yarn
- hadoop
---

Yarn ist der "neue" Jobscheduler

Anschauen aller Jobs

``` bash
yarn application -list

Total number of applications (application-types: [] and states: [SUBMITTED, ACCEPTED, RUNNING]):7

                Application-Id      Application-Name        Application-Type          User       Queue               State         Final-State         Progress

application_1533623521450_1214                distcp               MAPREDUCE          hdfs  root.hdfs.hdfs             RUNNING           UNDEFINED           94.67%
application_1533623521450_1635         Hive on Spark                   SPARK          hdfs  root.hdfs.hdfs             RUNNING           UNDEFINED              10%
application_1533623521450_1590         Hive on Spark                   SPARK          hdfs  root.hdfs.hdfs             RUNNING           UNDEFINED              10%
application_1533623521450_1589         Hive on Spark                   SPARK          hdfs  root.hdfs.hdfs             RUNNING           UNDEFINED              10%
```

Beenden eines Jobs

``` bash
yarn application -kill $id
```