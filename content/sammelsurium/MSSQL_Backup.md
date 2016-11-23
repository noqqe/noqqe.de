---
title: MSSQL Backup
date: 2012-04-18T17:17:35.000000
tags: 
- Databases
- MSSQL
---


## 1. Maintenance Set runterladen

http://ola.hallengren.com/

und mit dem SQL Manager auf dem Zielsystem öffnen und "Run" drücken


## 2. Extrahierte SQLCMD Commands für die Databases
~~~
sqlcmd -E -S .\ -d master -Q "EXECUTE [dbo].[DatabaseBackup] @Databases = 'SYSTEM_DATABASES', @Directory = N'C:\MSSQL_Backups', @BackupType = 'FULL', @Verify = 'Y', @CleanupTime = 48, @CheckSum = 'Y'" -b
sqlcmd -E -S .\ -d master -Q "EXECUTE [dbo].[DatabaseBackup] @Databases = 'USER_DATABASES', @Directory = N'C:\MSSQL_Backups', @BackupType = 'DIFF', @Verify = 'Y', @CleanupTime = 48, @CheckSum = 'Y'" -b
sqlcmd -E -S .\ -d master -Q "EXECUTE [dbo].[DatabaseBackup] @Databases = 'USER_DATABASES', @Directory = N'C:\MSSQL_Backups', @BackupType = 'FULL', @Verify = 'Y', @CleanupTime = 48, @CheckSum = 'Y'" -b
sqlcmd -E -S .\ -d master -Q "EXECUTE [dbo].[DatabaseBackup] @Databases = 'USER_DATABASES', @Directory = N'C:\MSSQL_Backups', @BackupType = 'LOG', @Verify = 'Y', @CleanupTime = 48, @CheckSum = 'Y'" -b
~~~

## 3. Extrahierte Database Integrity Checks und Optimize Jobs
~~~
sqlcmd -E -S .\ -d master -Q "EXECUTE [dbo].[DatabaseIntegrityCheck] @Databases = 'SYSTEM_DATABASES' " -b
sqlcmd -E -S .\ -d master -Q "EXECUTE [dbo].[DatabaseIntegrityCheck] @Databases = 'USER_DATABASES' " -b
sqlcmd -E -S .\ -d master -Q "EXECUTE [dbo].[IndexOptimize] @Databases = 'USER_DATABASES'" -b
~~~

## 4. Cleanup Jobs
~~~
sqlcmd -E -S .\ -d msdb -Q "DECLARE @CleanupDate datetime SET @CleanupDate = DATEADD(dd,-30,GETDATE()) EXECUTE dbo.sp_delete_backuphistory @oldest_date = @CleanupDate" -b
sqlcmd -E -S .\ -d msdb -Q "DECLARE @CleanupDate datetime SET @CleanupDate = DATEADD(dd,-30,GETDATE()) EXECUTE dbo.sp_purge_jobhistory @oldest_date = @CleanupDate" -b
~~~

## 5. Cleanup der Logfiles im FS
~~~
cmd /q /c "For /F "tokens=1 delims=" %v In (''ForFiles /P "' + @OutputFileDirectory + '" /m *_*_*_*.txt /d -30 2^>^&1'') do if EXIST "' + @OutputFileDirectory + '"\%v echo del "' + @OutputFileDirectory + '"\%v& del "' + @OutputFileDirectory + '"\%v"'
~~~


## 6. Scheduler Tasks für SQL Express 2008 R2

~~~
schtasks /create /tn "MSSQL - Backup UserDB FULL" /ru "sqlservice" /rp "xxx" /tr "C:\MSSQL_Backups\scripts\backup_userdatabaseFULL.cmd" /sc daily /st 01:00
schtasks /create /tn "MSSQL - Backup UserDB Log" /ru "sqlservice" /rp "xxx" /tr "C:\MSSQL_Backups\scripts\backup_userdatabaseLOG.cmd" /sc daily /st 03:00
schtasks /create /tn "MSSQL - Cleanup" /ru "sqlservice" /rp "xxx" /tr "C:\MSSQL_Backups\scripts\Cleanup.cmd" /sc daily /st 01:00
schtasks /create /tn "MSSQL - DatabaseIntegrity SystemDB" /ru "sqlservice" /rp "xxx" /tr "C:\MSSQL_Backups\scripts\DatabaseIntegrityCheck_SystemDB.cmd" /sc daily /st 03:00
schtasks /create /tn "MSSQL - DatabaseIntegrity UserDB" /ru "sqlservice" /rp "xxx" /tr "C:\MSSQL_Backups\scripts\DatabaseIntegrityCheck_UserDB.cmd" /sc daily /st 03:00
schtasks /create /tn "MSSQL - IndexOptimize UserDB" /ru "sqlservice" /rp "xxx" /tr "C:\MSSQL_Backups\scripts\IndexOptimize_UserDB.cmd" /sc daily /st 03:00
~~~
