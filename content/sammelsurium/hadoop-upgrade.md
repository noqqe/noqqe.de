---
title: Hadoop Upgrade
date: 2015-07-06T08:41:52
tags:
- Software
- Hadoop
---

#### Mirror

    wget --no-parent -r http://archive.cloudera.com/cdh5/ubuntu/precise/amd64/cdh/

#### Backups

On Namenode Primary

    sudo -u hdfs hdfs dfsadmin -safemode enter
    sudo -u hdfs hdfs dfsadmin -saveNamespace

Save Meta Data

    cd /mnt1/dfs/nn/
    sudo tar cfvz /root/mnt1-nn.tar.gz .
    cd /mnt/dfs/nn
    sudo tar cfvz /root/mnt-nn.tar.gz .

#### Turn off Services

    /etc/init.d/hadoop-hdfs-namenode stop
    /etc/init.d/hadoop-hdfs-zkfc stop
    /etc/init.d/hadoop-yarn-resourcemanager stop

#### Stop Journal, History and Zookeeper Cluster

    for x in /etc/init.d/hadoop-* ; do sudo service $x stop ; done
    /etc/init.d/hadoop-hdfs-journalnode stop
    /etc/init.d/hadoop-mapreduce-historyserver start

#### Stop Datanodes

    for x in ip-10-4-{3,4}-{11..13}.example.com; do ssh $x "sudo /etc/init.d/hadoop-hdfs-datanode stop" ; done

#### Update

Using puppet

```
     apt::source { "cloudera":
     location          => $repository,
-    release           => "precise-cdh5.0.1",
+    release           => "precise-cdh5.4.3",
     repos             => "contrib",
     include_src       => false,
     require           => Apt::Key['cloudera'],
     }
```

Alternatives to change back to normal Hadoop Configuration

    sudo update-alternatives --install /etc/hadoop/conf hadoop-conf /etc/hadoop/conf.noqqe/ 100
    for x in /etc/init.d/hadoop-* ; do sudo $x restart; done

Namenode

    apt-get install avro-libs bigtop-jsvc bigtop-utils flume-ng hadoop hadoop-hdfs hadoop-hdfs-namenode hadoop-hdfs-zkfc hadoop-mapreduce hadoop-yarn hadoop-yarn-resourcemanager parquet parquet-format zookeeper

Datanode

    apt-get install hadoop-hdfs-datanode

Job, Journal and History Server

    apt-get install hadoop-hdfs-journalnode zookeeper zookeeper-server

Misc

    apt-get install bigtop-utils avro-libs bigtop-jsvc parquet-format parquet

Also useful

     puppet agent -t ; aptitude install hadoop-hdfs-datanode hadoop-yarn-nodemanager hadoop-mapreduce ; puppet agent -t

Upgrade Namenode ACTIVE (with running history and journal server)

    sudo service hadoop-hdfs-namenode upgrade

Upgrade Namenode STANDBY (with running history and journal server)

    sudo -u hdfs hdfs namenode -bootstrapStandby

Turn Safemode off

    sudo -u hdfs hdfs dfsadmin -safemode leave

#### Verification

    $  hdfs haadmin -getServiceState nn2
    standby
    $  hdfs haadmin -getServiceState nn1
    active

#### After Works

After a few weeks

    sudo -u hdfs hdfs dfsadmin -finalizeUpgrade
    Finalize upgrade successful for hd1.example.com/10.4.3.10:8020
    Finalize upgrade successful for hd2.example.com/10.4.4.10:8020