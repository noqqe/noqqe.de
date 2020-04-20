---
title: Hiera
date: 2015-01-29T15:38:39
tags:
- Puppet
- Hiera
- Software
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

* [Boolean in Hiera](https://groups.google.com/forum/#!topic/puppet-users/BLw91HEpP0I)
* [Puppet Bug #17474](https://projects.puppetlabs.com/issues/17474)

weil

    hiera("puppet::key: value", 'no')

immer true zurück gibt.

#### Hiera Query

    hiera -h -d users ::application=ldm ::applicationtier=ldm-na-prod ::environment=production

### Puppet Explain Query

    /opt/puppetlabs/bin/puppet lookup classes --merge unique --explain
