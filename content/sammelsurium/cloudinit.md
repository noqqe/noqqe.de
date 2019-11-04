---
title: cloudinit
date: 2013-12-02T08:11:47
tags: 
- Software
- cloudconfig
---

## Running Command

    /usr/bin/python /usr/bin/cloud-init-cfg all config

## Demo Template

```
#cloud-config
apt_upgrade: false

runcmd:
{%- for index, role in server_role %}
- echo 'Facter.add("role{{index}}") do setcode do "{{ role }}" end end' >/usr/lib/ruby/1.8/facter/serverrole{{index}}.rb
{%- endfor %}
{% if zkid %}
- echo 'Facter.add("zookeeper_id") do setcode do "{{ zkid }}" end end' >/usr/lib/ruby/1.8/facter/zookeeperid.rb
{% endif %}

puppet:
  conf:
    agent:
      server: "{{puppet_server}}"
      environment: "staging"
```