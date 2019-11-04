---
title: MSSQL Logfiles
date: 2012-02-01T13:33:43
tags: 
- Databases
- MSSQL
---

Typische Ablage fuer Log

    D:\Programme\MSSQL\MSSQL.1\MSSQL\LOG

Ablage der wirklichen Backups auf

    \\backup01\SQL\

## Recovery Modes

[http://msdn.microsoft.com/en-us/library/ms189275.aspx](http://msdn.microsoft.com/en-us/library/ms189275.aspx)

### Alle Recovery-Mode Datenbanken mit "Full" auflisten

Im SQL Server Studio auf "New Query" drücken, dass einfügen:

```
SELECT name FROM SYS.DATABASES WHERE recovery_model!=1 AND name NOT IN
('master','tempdb','msdb','ReportServerTempDB');
```

und dann execute. Unter anderem findet das Verwendung in dem NRPE Check:

```
C:\nrpe\plugins\check_mssql_nt.exe
/H:db.example.com /U:monitoring /P:xxx /CHK_QUERY:"SELECT
COUNT(name) FROM SYS.DATABASES WHERE recovery_model!=1 AND name NOT IN
('master','tempdb','msdb','ReportServerTempDB')" /FO: "%d Datenbanken ohne
Recovery-Model=Full" /Wv:1
```