---
title: Vagrant
date: 2013-11-21T14:04:21
tags: 
- Software
- Vagrant
---

multi machine example

```
## -*- mode: ruby -*-
## vi: set ft=ruby :

## Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|

  config.vm.box_url = "http://files.vagrantup.com/precise32.box"


  config.vm.define "puppetmaster" do |puppetmaster|
    puppetmaster.vm.box = "precise32"
    puppetmaster.vm.network "forwarded_port", guest: 22, host: 2200
    puppetmaster.vm.network "forwarded_port", guest: 8443, host: 4567
    puppetmaster.vm.network "forwarded_port", guest: 80, host: 4580
    puppetmaster.vm.network "forwarded_port", guest: 443, host: 4543
  end

  config.vm.define "puppetmaster2" do |puppetmaster2|
    puppetmaster2.vm.box = "precise32"
    puppetmaster2.vm.network "forwarded_port", guest: 8443, host: 4667
    puppetmaster2.vm.network "forwarded_port", guest: 80, host: 4680
    puppetmaster2.vm.network "forwarded_port", guest: 443, host: 4643
  end

  config.vm.define "puppetclient" do |puppetclient|
    puppetclient.vm.box = "precise32"
  end

end
```

