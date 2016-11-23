---
title: find
date: 2014-02-20T17:12:57.000000
tags: 
- Software
- find
---


To find directories with wrong permissions:

~~~
$ find /data/share01/ \( ! -perm -o+w -and -type d \) -print
$ find /data/share01/ \( ! -perm -u+w -and -type d \) -print
~~~

Fix permissions

~~~
$ find /data/share01/ \( ! -perm -o+w -and -type d \) -exec chmod o+rwx {} \;
$ find /data/share01/ \( ! -perm -u+w -and -type d \) -exec chmod u+rwx {} \;
~~~
