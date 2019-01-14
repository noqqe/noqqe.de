---
title: Hadoop Cheatsheet
date: 2014-03-19T15:59:03
tags: 
- Software
- Hadoop
---

#### HDFS Benutzung

Files anzeigen

    hadoop fs -ls /

File hochladen

    hadoop fs -copyFromLocal <local_FS_filename> <target_on_HDFS>
    hadoop fs -put local.txt hdfs://nn.example.com/user/hadoop/hadoopfile

File löschen

    hadoop fs -rm <target_on_HDFS>

File verschieben

    hadoop fs -mv <source_on_HDFS> <target_on_HDFS>

Berechtigungen aendern

    hadoop fs -chmod 777 <target_on_HDFS>

Tail auf ein File im HDFS

    hadoop fs -tail <target_on_HDFS>

Replication Factor eines Files ändern

    hadoop fs -setrep 2 <target_on_HDFS>

Einen Ordner erstellen

    hadoop fs -mkdir <target_on_HDFS>

Größe eines Ordners anschaeun

    hadoop fs -du <target_on_HDFS>

#### Administratives

Übersicht

    hdfs dfsadmin -report | head -11

Alle decommissioned in progress

    hdfs dfsadmin -report | grep -B 3 progr | grep -E '(^Name|progr)'

alle fertig decommissioned

    hdfs dfsadmin -report | grep -B 3 -i decommissioned | grep -E '(^Name|progr)'

Hadoop FSCK ausführen

    hadoop fsck / | grep -v '^\.'

Namenode initalisieren

    hdfs namenode -format

Zookeeper Initalisieren

    hdfs zkfc -ZKformat

Standbyhost initalisieren

    hdfs namenode -bootstrapStandby

Datanode reformat

    /etc/init.d/hadoop-hdfs-datanode stop ;  for x in /mnt{1..4}/dfs ; do rm -rf $x ; mkdir $x ; chown hdfs:hadoop $x ; done ; /etc/init.d/hadoop-hdfs-datanode start

Stop all Datanodes

    for x in $hosts ; do ssh $x "sudo /etc/init.d/hadoop-hdfs-datanode stop" ; done

Get current master

    hdfs haadmin -getServiceState nn1
    > active

    hdfs haadmin -getServiceState nn2
    > standby

#### Architektur

* Namenode + Namenode mit Zookeeper und Journal Server
  ist sowas wie das DRBD mit MySQL Setup. Kein Secondary aber dafür einen Standby Namenode.
  Die initalisierung des Secondary erfolgt über den Journal server nach dem der
  primary initalisiert worden ist.

* Standalone ist halt einfach ein einzelner Namdenode

* Primary Namenode + Secondary.
  Secondary ist der wohl irreführendeste begriff überhaupt
  er wird für allerhand meta daten information missbrauchtn.
  Wenn der Primary down ist, hilft auch der Secondary nichts.

#### Namenode Reformat

    n1: /etc/init.d/hadoop-hdfs-namenode stop
    n2: /etc/init.d/hadoop-hdfs-namenode stop
    j1: /etc/init.d/hadoop-hdfs-journalnode stop
    n1: rm -rf /mnt1/dfs/ /mnt/dfs/ ; mkdir /mnt1/dfs/ /mnt/dfs/ ; chown hdfs /mnt1/dfs/ /mnt/dfs/
    j1: rm -rf /mnt/dfs/journal/
    j1: /etc/init.d/hadoop-hdfs-journalnode start
    n1: hdfs namenode -format
    n1: rsync /mnt/ und /mnt1/ to n2
    n1: /etc/init.d/hadoop-hdfs-namenode start
    n2: /etc/init.d/hadoop-hdfs-namenode start

#### Overview Script

    #!/bin/bash
    echo "##### REPORT"
    hdfs dfsadmin -report | head -11
    echo
    echo "##### REPORT Decommisioning in Progress"
    hdfs dfsadmin -report | grep -B 3 progr | grep -E '(^Name|progr)'
    echo
    echo "##### REPORT Decommisioned"
    hdfs dfsadmin -report | grep -B 3 -i decommissioned | grep -E '(^Name|progr)'
    echo
    echo "##### REPORT fsck / "
    hadoop fsck / | grep -v '^\.'

#### List all jobs in queue

as a superuser in hadoop (hdfs)

```
$ mapred job -list
2 jobs currently running
JobId   State   StartTime       UserName        Priority        SchedulingInfo
job_201403191443_0007   1       1395240869633   ubuntu  NORMAL  NA
job_201403191443_0010   1       1395241066090   ubuntu  NORMAL  NA
```