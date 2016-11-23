---
title: ntp
date: 2012-11-02T11:08:51.000000
tags: 
- Software
- ntp
---


Debian Quick and Dirty

~~~
aptitude install ntpdate && ntpdate pool.ntp.org && \
hwclock --systohc && \
aptitude purge ntpdate && aptitude install ntp && \
sed -i s/debian.pool/de.pool/g /etc/ntp.conf && /etc/init.d/ntp restart
~~~

Uhrenvergleich

~~~
$ ssh app01.example.com "date --rfc-3339=ns" ; ssh app2.example.com "date --rfc-3339=ns"
2012-11-02 10:43:45.608171723+01:00
2012-11-02 10:40:51.491816591+01:00
~~~

Useful Debugging Krempel

~~~
ntpq -p de.pool.ntp.org
ntpdate -d de.pool.ntp.org
~~~
