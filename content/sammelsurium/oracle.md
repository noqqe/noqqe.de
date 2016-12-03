---
title: Oracle
date: 2012-01-04T14:15:50
tags: 
- Databases
- Oracle
---

Connect:

    /usr/lib/nagios/plugins ##  su - oracle
    /home/oracle > sqlplus / as user

History:

    SYS@ruix AS SYSDBA> ed

Query:

    select owner,index_name from dba_indexes where index_name like 'VALUE%'
    ## % ist normale wildcard
