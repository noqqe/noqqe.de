---
title: cloudinit
date: 2013-12-02T08:11:47
tags:
- Software
- cloudconfig
---

## Running Command

    cloud-init init

## Scripts

Um ausgeführte runcmd

    cat /var/lib/cloud/instance/scripts/runcmd

Darf maximal 16kb sein, sonst:

    Error: expected length of user_data to be in the range (0 - 16384), got #cloud-config

## Befehle ausführen

```
bootcmd:
 - mkdir /reports

runcmd:
 - echo foo
```

## User anlegen

```
users:
  - name: foo
    gecos: foo
    sudo: ALL=(ALL) NOPASSWD:ALL
    groups: users, admin
    lock_passwd: true
    ssh_authorized_keys:
      - ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAA...
```

## Puppet Template

``` yaml
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
