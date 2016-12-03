---
date: '2016-06-22T14:05:15'
tags:
- hdfs
- blocks
- administration
- devops
- hadoop
- datanode
- namenode
title: HDFS Storage Sizes
---

Hadoop Filesystem Informationen können echt ziemlich verwirrend sein.
Zuerst, sieht so aus als würden erstmal 3.69 TB benutzt werden. Als Admin ist dann erstmal alles
cool. Bis man das Cluster verkleinern will und Daten wegschmeissen muss.

```
$ hdfs dfsadmin -report | head -11
Configured Capacity: 10498624905216 (9.55 TB)
Present Capacity: 10496880074752 (9.55 TB)
DFS Remaining: 6441878704128 (5.86 TB)
DFS Used: 4055001370624 (3.69 TB)
DFS Used%: 38.63%
Under replicated blocks: 0
Blocks with corrupt replicas: 0
Missing blocks: 0
Missing blocks (with replication factor 1): 0
```

Irgendwo müssten sich die 3 TB aber jetzt auch finden lassen.
Das FS kann natürlich Auskunft darüber geben, wo was für Files welche Größe
haben.

```
$ hdfs dfs -du -h /
1.7 G    3.4 G    /data
73.8 M   147.5 M  /jars
545.6 M  5.1 G    /tmp
27.2 G   54.4 G   /user
62.3 M   124.5 M  /var
```

Der aufmerksame Leser merkt schnell, dass das so garnicht hinhaut. Was
heissen die beiden Zahlen überhaupt. In der Manpage findet man die
Spaltenbeschriftung mit `size`, `disk space consumed`, `name`.

Für den Unterschied zwischen Size und Diskspace muss man wissen, dass es
auch bei Hadoop Blocksizes gibt. Bei Hadoop nur um einiges größer. Dieses
Limit sollte man in der `hdfs-site.xml` finden.

``` xml
<property>
  <name>dfs.blocksize</name>
  <value>134217728</value>
</property>
```

Na gut, das sind halt jetzt ~134MB pro File. Die nächste Baustelle lauter
`replicationFactor`. Auch die kann in der `hdfs-site.xml` gesetzt werden.
Das muss aber nicht ausschliesslich der Fall sein. Spezielle
ReplicationFactors können pro File/Directory via Commandline gesetzt
werden.

``` xml
<property>
  <name>dfs.replication</name>
  <value>2</value>
</property>
```

Im Klartext heisst das jetzt: Ich lade eine 1,5KB grosse `foobar.txt` in
HDFS und verbrauche dafür 268MB (= 134MB * 2 Replication Factor).

Man kann sich außerdem noch einen kleinen Überblick verschaffen indem man
die `count` Option benutzt. Die Spaltenbeschriftung hab ich mir von
Stackoverflow ausgeliehen. Normalerweise ist die nicht dabei. Und überhaupt
Hadoop, Spaltenbenamsung!

```
$ hdfs dfs -count -q -h '/var/*'
QUOTA REMAINING_QUOTA SPACE_QUOTA REMAINING_SPACE_QUOTA DIR_COUNT FILE_COUNT CONTENT_SIZE FILE_NAME
none      inf            none             inf            187        367            61.7 M /var/log
```

Das haut aber immernoch nicht alles hin. Ein weiterer Befehl für die
Inspektion ist `hdfs fsck`. Indem man genau Anzahl der Blocks im HDFS
ausliesst, muss man sich nicht mit dem über den Daumen gepeilten ReplFactor
aus der `hdfs-site.xml` zufrieden geben.

```
$ hdfs fsck -blocks
Total size:    31726972623 B (Total open files size: 16694 B)
Total dirs:    705
Total files:   3545
Total symlinks:                0 (Files currently being written: 1)
Total blocks (validated):      3721 (avg. block size 8526464 B) (Total open file blocks (not validated): 1)
Minimally replicated blocks:   3721 (100.0 %)
Over-replicated blocks:        0 (0.0 %)
Under-replicated blocks:       0 (0.0 %)
Mis-replicated blocks:         0 (0.0 %)
Default replication factor:    2
```

Also gut, 3721 Blocks * 134MB = 496GB. Immernoch nicht genug. Irgendwas
geht falsch. Bei einem Blick in die Node details `hdfs dfsadmin -report`
fiel mir auf, dass auf ein paar Nodes 600GB used und auf anderen nur 2GB
used sind. Ein Fall für den Balancer. Dieser verteilt bei Ungleichgewicht
alle Daten neu und gerecht über alle Nodes.

```
$ hdfs balancer
INFO balancer.Balancer: 5 over-utilized:
[192.168.4.26:50010:DISK, 192.168.4.22:50010:DISK, 192.168.4.25:50010:DISK,
192.168.4.24:50010:DISK, 192.168.4.11:50010:DISK]
INFO balancer.Balancer: 7 underutilized:
[192.168.4.23:50010:DISK, 192.168.4.29:50010:DISK, 192.168.4.30:50010:DISK,
192.168.4.32:50010:DISK, 192.168.4.31:50010:DISK, 192.168.4.27:50010:DISK,
192.168.4.28:50010:DISK]
INFO balancer.Balancer: Need to move 1.33 TB to make the
cluster balanced.
INFO balancer.Balancer: Decided to move 10 GB bytes from
192.168.4.26:50010:DISK to 192.168.4.23:50010:DISK
INFO balancer.Balancer: Decided to move 10 GB bytes from
192.168.4.22:50010:DISK to 192.168.4.29:50010:DISK
INFO balancer.Balancer: Decided to move 10 GB bytes from
192.168.4.25:50010:DISK to 192.168.4.30:50010:DISK
```

Dann wurde es mir klar, dass beim letzten Upgrade des HDFS Clusters
vergessen wurde das Upgrade via `hdfs dfsadmin -finalizeUpgrade` abzuschliessen.
Alle Changes wurden in eine Art Snapshot vorgehalten, aber nicht wirklich
applied. So kam es zum Ungleichgewicht im Cluster. Das Upgrade hab ich im
Anschluss finalisiert und den Balancer in der Endlosschleife laufen lassen.

```
$ while true; do hdfs balancer ; done
INFO balancer.Dispatcher: Successfully moved
blk_1074285727_545003 with size=285553 from 192.168.4.24:50010:DISK to
192.168.4.29:50010:DISK through 192.168.4.24:50010
INFO balancer.Dispatcher: Successfully moved
blk_1074285725_545001 with size=285240 from 192.168.4.24:50010:DISK to
192.168.4.29:50010:DISK through 192.168.4.24:50010
```

Führte dann auch zum gewünschten Ergebnis. Der Balancer wird mit Sicherheit
noch längere Zeit laufen und unnötige Daten beim verteilen verwerfen. Aber
erste Veränderungen kann man bereits erkennen.

```
Configured Capacity: 10498624905216 (9.55 TB)
Present Capacity: 10496611639336 (9.55 TB)
DFS Remaining: 8326884075111 (7.57 TB)
DFS Used: 2169727564225 (1.97 TB)
DFS Used%: 20.67%
Under replicated blocks: 0
Blocks with corrupt replicas: 0
Missing blocks: 0
Missing blocks (with replication factor 1): 0
```

Außerdem hab ich gelernt, mir fürs Finaliseren des Upgrades einen
Kalendereintrag zu machen ;)