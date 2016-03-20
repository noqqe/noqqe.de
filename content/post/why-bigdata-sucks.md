---
categories:
- administration
- devops
- linux
- opensource
- osbn
- planetenblogger
- ubuntuusers
date: 2016-03-20T14:36:09+01:00
tags:
- Hive
- BigData
- Flume
- Hadoop
- Yarn
title: Why BigData sucks.
---

In der Arbeit komme ich häufiger mit BigData in Berührung.
BigData. Ein hartnäckiges Buzzword, das sich seit Jahren in der
IT Branche hält.

{{< figure src="/uploads/2016/03/bigdata.jpeg" >}}

Irgendwie geht es um Hadoop. Der Standardstack ist ziemlich robust. Also
Hadoop HDFS und Hadoop MapReduce. Aber es gibt noch 100 andere
Software Projekte die sich BigData auf die Fahne schreiben. Ein kleiner
Auszug:

> Hive, Spark, Impala, Tez, HortonWorks, Cloudera, Redshift, Kinesis,
> ElasticSearch, Tableau, Flume, Storm, Druid, Flink, HBASE,
> Drill, Kafka, Oozie, Knox, Parquet, Sqoop, Crunch, Lumify, Solr,
> Samoa, Yarn

Und das sind nur ein paar Wenige. All diese Projekte interagieren in
irgendeiner Weise mit HDFS, MapReduce oder werden im selben Kontext
angewendet. Alle 2-3 Wochen kommt ein neues Projekt auf den Markt und wird
als "the next big thing" gehandelt. Das stellt einige Anforderungen an den
Admin.

Auch Ansprüche die man als Admin an sich selbst stellt. Das Setup soll ja schliesslich
automatisiert, ausrollbar und hochverfügbar sein.

### Administration Hell

Und jetzt einfach mal vorstellen, dass man als Wald-und-Wiesen Admin auf
sowas losgelassen wird. Jede Software mit eigener Architektur, eigenen
Wording und eigenen Konzepten.

Noch dazu ist das meiste Zeug gerade erst
2-3 Monate alt, seit es aus den Tiefen des Open Apache Graveyards
aufkeimte. Da lässt die Softwarequalitätssicherung grüßen.

Ein kleines Spiel. Die folgenden 6 Worte aus dem Hadoopumfeld stellen
entweder Master- oder Slave-daemon dar. Ordne zu.

* Nodemanager
* Namenode
* Resourcemanager
* Datanode
* Jobtracker
* Tasktracker

Von oben nach unten: Slave, Master, Master, Slave, Slave, Master, Slave.
Witzig ist auch, dass der "Secondary Namenode" so ziemlich alles macht aber
kein Secondary ist.

Man merkt öfter, dass das Stück Software einfach nicht zuende Gedacht ist.
Mal fehlt das Logrotate, mal ein anständiges Interface um das Cluster zu
Administrieren. In den meisten Fällen fliegt einem Software X einfach wegen
Kleinigkeiten um die Ohren und man findet sich selbst beim Googeln nach Java Stacktraces
wieder.

### Quality!

Wenn ich mit dem Anspruch von Unix Software an BigData denke, dreht sich
mir alles um. Das Gefühl beschleicht einen, dass nahezu alle
Grundsätze von "Simple, Robust, Secure, Fast, Well documented" ignoriert
werden.

{{< figure src="/uploads/2016/03/classpath.png" >}}

### Community

Ich würde behaupten, mich halbwegs in dem Thema BigData auszukennen. Mit
Sicherheit bin ich kein Spezialist, aber ich finde mich zurecht. Selbst
wenn ich in den tiefsten Ecken meines Verstandes grabe, fällt es mir bei
typischen BigData Artikeln schwer zu verstehen was der Post eigentlich
sagen möchte. Ein Beispiel:
[16 Top Big Data Platforms](http://www.informationweek.com/big-data/big-data-analytics/16-top-big-data-analytics-platforms/d/d-id/1113609)

{{< figure src="/uploads/2016/03/iphone.jpg" >}}

Ich meine ich lese hier kein Tolstoi. Was ist denn los zur Hölle? Und das
geht *17* Seiten so weiter!

### Conclusion

Ich glaube, das Geschäftsfeld BigData braucht einfach Zeit. Zeit bis sich die
"Big Players" herauskristallisiert haben. Bis genügend Startups die Pleite
gegangen sind. Bis Data Scientists & Developer nicht mehr gezwungen sind
130 Tools zu evaluieren um Mehrwert zu generieren. Dann wird BigData
annehmbar.

