---
title: Puppet Loops
date: 2013-10-28T10:48:35
tags: 
- Puppet
---

If you need to simply loop through multiple iterations, simply use a definition.

With definition you are able to use "title" as the iteration variable. Here's a example

```
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

    flumeports { ["8800", "8801", "8802", "8803", "8804", "8805", "8806", "8807"]: }

## Parameterized Class

Definition

```
class echo_class ($to_echo = "default value") {
  notify {"What are we echoing? ${to_echo}.":}
}
```

Class Call

```
class {'echo_class':
  to_echo => 'Custom value',
}
```

## Array Iteration

```
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

## Debug a module

### extracting

## Check if everything works with the Future Parsers

    for x in $(find . -type f -iname '*.pp') ; do echo $x; puppet parser --parser future validate $x ; done