---
title: Puppet
date: 2013-06-30T16:44:25
tags:
- Software
- Puppet
---

## Installation der Pakete

Repos von Puppetlabs laden

    wget http://apt.puppetlabs.com/puppetlabs-release-bionic.deb
    wget http://apt.puppetlabs.com/puppetlabs-release-buster.deb


    dpkg -i puppetlabs-release-wheezy.deb apt-get update && apt-get install puppet-agent

Zertifikat anfordern:

    puppet agent -t --server <puppetmaster> --environment <envname>

Nun auf dem entsprechenden puppet master den Fingerprint vergleichen und
auf dem puppet master den Certificate Request signieren:

    puppetserver ca sign <certname>

# Require

Um Abhängig von mehreren Resourcen zu sein:

``` puppet
require    => [ Group[users], Group[admin] ],
```

## defined()

Um zu checken ob eine Resource evtl noch an einer anderen Stelle definiert
worden ist, kann mit dieser Zeile gemacht werden.

```puppet
if ! defined(Package['apparmor']) {
  package { 'apparmor':
    ensure => purged,
  }
}
```

## Parameterized Class

Definition

``` puppet
class echo_class ($to_echo = "default value") {
  notify {"What are we echoing? ${to_echo}.":}
}
```

Class Call

``` puppet
class { 'echo_class':
  to_echo => 'Custom value',
}
```

## Array Iteration

``` puppet
$my_env => [ shared1, shared2, shared3 ]

define myResource {
  file { "/var/tmp/$name":
    ensure => directory,
    mode => 0600,
  }
  user { $name:
    ensure -> present,
  }
}

myResource { $my_env: }
```

## Default Ports um zum Master zu connecten

Port 8140 ist default

aber mit `--masterport 8888` kann das umspezifiziert werden.

## Puppet Nodes deaktivieren

Um Hosts aus Puppet/Puppetboard/PuppetDB zu werfen kann man folgende Zeilen
nutzen (Puppet 6)

    puppet node deactivate <certname>
    puppetserver ca list --all | grep fqdn
    puppetserver ca clean --certname <certname>

