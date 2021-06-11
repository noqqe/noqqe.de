---
title: Linux Rename Host
date: 2012-02-24T10:08:20
tags:
- Linux
- OS
---

<!--more-->

## Files

Es gibt ein paar Stellen auf die man bei Debian achten muss.

    vim /etc/hostname
    vim /etc/hosts
    vim /etc/mailname

Und eben noch andere Applikationsspezifische

  * `/etc/puppetlabs/puppet/puppet.conf`
  * Icinga2


## Commands

Früher hat ein `hostname` gereicht

    hostname neuername

Seit systemd gibt es ein neues Tool


```
> hostnamectl
   Static hostname: dcex10stconfdoc.acme.com
         Icon name: computer-vm
           Chassis: vm
        Machine ID: 80e2ae319e7c4abcbdf0ee2840d6e8b2
           Boot ID: 5860e57dee9f4698af26b0f7b810dc7d
    Virtualization: vmware
  Operating System: Debian GNU/Linux 10 (buster)
            Kernel: Linux 4.19.0-16-amd64
      Architecture: x86-64
```

und neu setzen geht auch so.

```
> hostnamectl set-hostname dcex10stconfdoc.acme.com
```
