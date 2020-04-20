---
title: Puppet
date: 2013-06-30T16:44:25
tags:
- Software
- Puppet
---

## Default Ports um zum Master zu connecten

Port 8140 ist default

aber mit `--masterport 8888` kann das umspezifiziert werden.

## Puppet Nodes deaktivieren

Um Hosts aus Puppet/Puppetboard/PuppetDB zu werfen kann man folgende Zeilen
nutzen (Puppet 6)

    puppet node deactivate <certname>
    puppetserver ca list --all | grep fqdn
    puppetserver ca clean --certname <certname>

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

## Puppet v2-v4 Loops

``` puppet
  define flumeports {
    @@nagios_service { "check_flume_port_${title}_${hostname}":
      use => "generic-service",
      service_description => "Flume Port ${title}",
      contact_groups => 'cg_ai_admins',
      check_command => "check_tcp!-p ${title}",
      notification_interval => "0",
      host_name => "$fqdn",
      target => "/etc/icinga/modules/service_check_flume_port_${title}_${::fqdn}.cfg",
    }
  }
```

Call the defined class with an array of objects:

``` puppet
flumeports { ["8800", "8801", "8802", "8803", "8804", "8805", "8806", "8807"]: }
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

## Check if everything works with the Future Parsers

Puppet kann mit dem Future Parser überprüfen ob Puppet 3 kompatibel

``` bash
for x in $(find . -type f -iname '*.pp') ; do
  echo $x;
  puppet parser --parser future validate $x
done
```
