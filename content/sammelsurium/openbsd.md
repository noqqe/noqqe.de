---
title: OpenBSD
date: 2013-06-14T11:56:41
tags: 
- OS
- OpenBSD
---

## Networking

Interfaces (static)

``` bash
> cat /etc/hostname.em0
inet 213.95.21.200 255.255.255.0 NONE
inet6 alias 2001:780:3:5::122 64
inet6 alias 2001:780:132::1 48
```

oder DHCP

``` bash
$ cat /etc/hostname.em0
dhcp
```


Defaultgateway

    > cat /etc/mygate
    213.95.21.1
    2001:780:3:5::1

Neustart

    sh /etc/netstart em0

DNS 

```
> cat /etc/resolv.conf
search example.com
nameserver 125.2.3.4
nameserver 125.2.3.5
lookup file bind
```

Hostname

```
cat /etc/myname
hostname.local
```


## Timezone

    ln -fs /usr/share/zoneinfo/Europe/Berlin /etc/localtime

## Encoding & Charset

    cat /etc/wsconsctl.conf
    keyboard.encoding=de

und in .bash_profile und .profile

    export LC_CTYPE="en_US.UTF-8"
    export LC_MESSAGES="en_US.UTF-8"

## Users 

Switch to a user with disabled account.

``` bash
$ su _nrpe -s /bin/ksh
This account is currently not available.

$ sudo -u _nrpe bash -l
16:25 _nrpe:/usr/local/nagios/libexec$ /usr/local/nagios/libexec/check_mysql_health
```

## ktrace (strace)

``` bash
ktrace -t + python 1.py    ## <- Got error like before: "ImportError: Cannot load specified object"
kdump -f ktrace.out | grep -2 -i 'permission denied'
```

