---
title: MongoDB Kernel Tuning
date: 2014-03-11T12:45:48
tags:
- Databases
- MongoDB
---

#### overcommit

2 bedeutet nur so viel overcommit wieder Swapspace hergibt.
80 ist die Ratio der Swap benutzt werden darf.

~~~
echo 2 > /proc/sys/vm/overcommit_memory
echo 80 > /proc/sys/vm/overcommit_ratio
vm.overcommit_memory=2
sysctl -p
~~~

Eventuell will man auch echo 1 machen wegen dem Journaling

~~~
"Thu Mar 20 14:16:55.481 [initandlisten] ** WARNING: /proc/sys/vm/overcommit_memory is 2",
"Thu Mar 20 14:16:55.481 [initandlisten] **          Journaling works best with it set to 0 or 1",
~~~

#### no atime

atime auf partition ausschalten

#### readahead bei partition

Einmalig:

    blockdev --setra 4096 /dev/sdb

Readahead warnings ausschalten, persistent

~~~
## anzeigen der readahead einstellungen
blockdev --report |grep sdb

## informationen anzeigen
udevadm info -a -p /sys/block/sdb/

## udev loggt in syslog, also dahin schauen
tail -f /var/log/messages

vi /etc/udev/rules.d/99-readahead.rules
> SUBSYSTEM=="block", ATTRS{model}=="Virtual disk*", ACTION=="add|change", KERNEL=="sdb", ATTR{bdi/read_ahead_kb}="512"

## testen
udevadm test --action=add /sys/block/sdb 2>&1 | grep 512

## wenn erfolgreichen durchführen
udevadm trigger

## verifizieren
blockdev --report |grep sdb
~~~

weiterführende links

[unix.stackexchange.com/questions/71364/persistent-blockdev-setra-read-ahead-setting](http://unix.stackexchange.com/questions/71364/persistent-blockdev-setra-read-ahead-setting)

[stackoverflow.com/questions/16715961/mongodb-readahead-warning](https://stackoverflow.com/questions/16715961/mongodb-readahead-warning)

[www.cyberciti.biz/tips/rhel-linux-optimize-3ware-raid-read-performance.html](http://www.cyberciti.biz/tips/rhel-linux-optimize-3ware-raid-read-performance.html#comment-159033)

#### ulimit

Fehlerbild

~~~
ERROR: mmap private failed with out of memory. (64 bit build)
Assertion: 13636:file /docdata/mongodb/data/xxx_letters.5 open/create failed in createPrivateMap (look in log for more information)
~~~

Folgende Werte sollten eingestellt sein.

~~~
-f (file size): unlimited
-t (cpu time): unlimited
-v (virtual memory): unlimited [1]
-n (open files): 64000
-m (memory size): unlimited [1]
-u (processes/threads): 32000
~~~

zum verifizieren kann folgende Bash funktion benutzt werden

~~~ { .bash }
return-limits(){

     for process in $@; do
          process_pids=`ps -C $process -o pid --no-headers | cut -d " " -f 2`

          if [ -z $@ ]; then
             echo "[no $process running]"
          else
             for pid in $process_pids; do
                   echo "[$process #$pid -- limits]"
                   cat /proc/$pid/limits
             done
          fi

     done

}
~~~

~~~
/etc/security/limits.conf
* soft nofile 64000
* hard nofile 64000
* soft rss unlimited
* hard rss unlimited
* soft as unlimited
* hard as unlimited
* soft nproc 32000
* hard nproc 32000
~~~
