---
title: OpenBSD MySQL
date: 2012-11-03T17:49:26
tags: 
- OS
- OpenBSD
---

~~~
## export PKG_PATH=ftp://ftp.openbsd.org/pub/OpenBSD/5.2/packages/`machine -a`/
## pkg_add mysql-server
~~~

## Initiales einrichten und starten

~~~
## /usr/local/bin/mysql_install_db
## mysqld_safe &
## /usr/local/bin/mysql_secure_installation
## mysql
~~~

## Separate Class fuer mysql definieren und automatisch starten

/etc/login.conf

~~~
#### Edited by noqqe
#
## Setting used by MySQL daemon
#
_mysql:\
        :datasize=infinity:\
        :maxproc=infinity:\
        :openfiles-cur=2048:\
        :openfiles-max=4096:\
        :stacksize-cur=8M:\
        :localcipher=blowfish,8:\
        :tc=default
~~~

/etc/rc.mysql

~~~
## Start MySQL server
if [ -x /usr/local/bin/mysqld_safe ] ; then
        su -c _mysql root -c '/usr/local/bin/mysqld_safe &' > /dev/null & echo -n ' mysql'
fi
~~~

/etc/rc.local

~~~
if [ -x /etc/rc.mysql ] ; then
        /etc/rc.mysql
fi
~~~

## Links

[http://www.openbsdsupport.org/mysql.htm](http://www.openbsdsupport.org/mysql.htm)
