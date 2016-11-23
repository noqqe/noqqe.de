---
title: MySQL Binlog Recovery
date: 2013-04-02T14:38:32.000000
tags: 
- Databases
- MySQL
---


## Binlog point in Time Recovery

    mysqlbinlog --start-datetime="2013-03-20 06:26:00" --stop-datetime="2013-03-20
    10:19:00" /data/mysql/binlog/mysql-bin.[0-9]* | mysql -u root -p
