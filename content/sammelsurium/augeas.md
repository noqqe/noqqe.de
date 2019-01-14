---
title: Augeas
date: 2014-01-07T13:40:15
tags: 
- Puppet
---

Interactive

~~~
$ aptitude install augeas-tools
$ augtool -b #make backup file
augtool> set KexAlgorithms diffie-hellman-group-exchange-sha256,diffie-hellman-group14-sha1,diffie-hellman-group-exchange-sha1
augtool> save
~~~

## Show errors

## Augeas within puppet

~~~
 ## permit root login no and switch off ubuntu
  augeas { "sshd_config":
    context => "/files/etc/ssh/sshd_config",
    changes => [
      "set PermitRootLogin no",
      "set DenyUsers/1 ubuntu",
      "set Ciphers aes256-gcm@openssh.com,aes128-gcm@openssh.com,aes256-ctr,aes128-ctr",
      "set KexAlgorithms diffie-hellman-group-exchange-sha256,diffie-hellman-group14-sha1,diffie-hellman-group-exchange-sha1",
      "set MACs/1    umac-128-etm@openssh.com",
      "set MACs/2    hmac-sha2-512",
      "set MACs/3    mac-sha2-256",
      "set MACs/4    hmac-ripemd160",
    ],
  }
~~~

## Strange usage syntax for sshd lense

[https://github.com/hercules-team/augeas/issues/69](https://github.com/hercules-team/augeas/issues/69)