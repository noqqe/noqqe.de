---
title: Semaphores
date: 2014-03-29T12:11:52.000000
tags: 
- OS
- OpenBSD
---


~~~
http://www.openbsd.org/cgi-bin/man.cgi?query=sem_open&apropos=0&sektion=0&manpath=OpenBSD+Current&arch=amd64&format=html

## uwsgi --ini /home/isso/uwsgi.cfg
[uWSGI] getting INI configuration from /home/isso/uwsgi.cfg
*** Starting uWSGI 2.0.3 (64bit) on [Sat Mar 29 12:07:19 2014] ***
compiled with version: 4.2.1 20070719  on 26 March 2014 19:46:22
os: OpenBSD-5.4 GENERIC.MP#41
nodename: o0.n0q.org
machine: amd64
clock source: unix
pcre jit disabled
detected number of CPU cores: 4
current working directory: /etc/rc.d
detected binary path: uwsgi
setuid() to 1003
your processes number limit is 128
your memory page size is 4096 bytes
detected max file descriptor number: 512
lock engine: ipcsem
uwsgi_lock_ipcsem_init()/semget(): No space left on device [core/lock.c line 507]
uwsgi_ipcsem_clear()/semctl(): Invalid argument [core/lock.c line 631]



12:07 root@o0:/etc/rc.d (master*) ## sysctl kern.seminfo
kern.seminfo.semmni=10
kern.seminfo.semmns=60
kern.seminfo.semmnu=30
kern.seminfo.semmsl=60
kern.seminfo.semopm=100
kern.seminfo.semume=10
kern.seminfo.semusz=112
kern.seminfo.semvmx=32767
kern.seminfo.semaem=16384


ipcs

2:08 root@o0:/etc/rc.d (master*) ## ipcs
Message Queues:
T       ID     KEY        MODE       OWNER    GROUP

Shared Memory:
T       ID     KEY        MODE       OWNER    GROUP

Semaphores:
T       ID     KEY        MODE       OWNER    GROUP
s   655360          0 --rw-rw-rw-     isso    wheel
s   720897          0 --rw-rw-rw-     isso    wheel
s   720898          0 --rw-rw-rw-     isso    wheel
s   655363          0 --rw-rw-rw-     isso    wheel
s   655364          0 --rw-rw-rw-     isso    wheel
s   655365          0 --rw-rw-rw-     isso    wheel
s   655366          0 --rw-rw-rw-     isso    wheel
s   655367          0 --rw-rw-rw-     isso    wheel
s   655368          0 --rw-rw-rw-     isso    wheel
s   655369          0 --rw-rw-rw-     isso    wheel


for x in $(ipcs | grep ^s | grep isso | awk '{print $2}') ; do echo $x ; done


ipcrm -s $id

12:16 root@o0:/etc/rc.d (master*) ## sysctl kern.seminfo.semmni=50
kern.seminfo.semmni: 10 -> 50



~~~



Sema