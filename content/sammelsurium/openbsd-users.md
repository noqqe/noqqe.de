---
title: OpenBSD Users
date: 2014-03-28T16:27:12
tags: 
- OS
- OpenBSD
---

Switch to a user with disabled account.
~~~
$ su _nrpe -s /bin/ksh
This account is currently not available.

$ sudo -u _nrpe bash -l
16:25 _nrpe:/usr/local/nagios/libexec$ /usr/local/nagios/libexec/check_mysql_health
~~~