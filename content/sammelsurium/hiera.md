---
title: Hiera
date: 2015-01-29T15:38:39
tags: 
- Puppet
- Hiera
---

Hiera Cheatsheet

#### Commandline Queries

Query value from hiera

    hiera -c /etc/puppet/hiera.yaml testfact

Query mit FQDN

    hiera -c /etc/puppet/hiera.yaml ::fqdn=host.name.org testfact

    hiera -c /etc/puppet/hiera.yaml -h users::user ::clientcert=host.example.com ::environment=testing

#### Die Boolean Kontroverse

Eine Faustregel. Einfach niemals Hiera Booleans verwenden

* https://t.co/Osq7cJao4f
* https://t.co/lgFD5MlohT

weil

    hiera("puppet::key: value", 'no')

immer true zurück gibt.

#### Hiera Query

    hiera -h -d users ::application=ldm ::applicationtier=ldm-na-prod ::environment=production
