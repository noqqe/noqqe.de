---
title: OpenBSD nginx & php
date: 2012-11-04T19:40:48
tags: 
- OS
- OpenBSD
---

~~~
## export PKG_PATH=ftp://ftp.openbsd.org/pub/OpenBSD/5.2/packages/`machine -a`/
## pkg_add -i nginx
Ambiguous: choose package for nginx
 a       0: <None>
         1: nginx-1.0.15
         2: nginx-1.0.15-passenger
Your choice: 1

## pkg_add php-fpm php-mysql
~~~

## Configure

## PHP

Um MySQL nutzen zu können gibts diesen Symlink

~~~
## mkdir -p /var/www/var/run/mysql
## chown www:daemon /var/www/var/run/mysql
## ln -f /var/run/mysql/mysql.sock /var/www/var/run/mysql/mysql.sock
~~~

mysqli install

~~~
##  export PKG_PATH=ftp://ftp.openbsd.org/pub/OpenBSD/5.2/packages/amd64/
## echo $PKG_PATH
ftp://ftp.openbsd.org/pub/OpenBSD/5.2/packages/amd64/
## pkg_add php-mysqli
Ambiguous: choose package for php-mysqli
 a       0: <None>
         1: php-mysqli-5.2.17p6
         2: php-mysqli-5.3.14p0
Your choice: 2
php-mysqli-5.3.14p0: ok
--- +php-mysqli-5.3.14p0 -------------------
You can enable this module by creating a symbolic link from
/etc/php-5.3.sample/mysqli.ini to
/etc/php-5.3/mysqli.ini. As root:

    ln -sf /etc/php-5.3.sample/mysqli.ini /etc/php-5.3/mysqli.ini
## ln -sf /etc/php-5.3.sample/mysqli.ini /etc/php-5.3/mysqli.ini
~~~

php-gd is broken and could not be loaded
with php-5.3 -m

http://blog.endpoint.com/2011/02/debugging-php-extensions-with-dynamic.html
