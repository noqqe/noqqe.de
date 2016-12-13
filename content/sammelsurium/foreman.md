---
title: Foreman
date: 2013-11-21T13:59:09
tags: 
- Puppet
---

## Anleitung

http://theforeman.org/manuals/1.3/index.html#3.3.3DebianPackages

    aptitude install foreman foreman-mysql2 mysql-server mysql-client

    create database foreman
    GRANT ALL ON foreman.* TO 'foreman'@'localhost' IDENTIFIED BY 'XXX';

## Configuration

http://theforeman.org/manuals/1.3/index.html#3.5.1InitialSetup

~~~
/etc/foreman/database.yml
production:
  adapter: mysql2
  database: foreman
  username: foreman
  password: PASSWORD
  host: localhost
~~~

#### Pitfalls

If you are planning to setup a foreman server for reporting and totally
decoupled from your puppet master, pay attention to the following things

* Set puppet.conf on master to

```
[master]
reports = foreman
external_nodes = /etc/puppet/node.rb
node_terminus = exec
```

* Set correct SSL Certs in foreman hosts apache2 (copy over puppet masters)
* Update `nodes.rb` file from foreman homepage
* Adjust foreman.yaml in /etc/puppet

```
---
## Update for your Foreman and Puppet master hostname(s)
:url: "https://dcex1045mfman01.ext.gfk"
:ssl_ca: "/var/lib/puppet/ssl/certs/ca.pem"
:ssl_cert: "/var/lib/puppet/ssl/certs/dcex1045mpups01.ext.gfk.pem"
:ssl_key: "/var/lib/puppet/ssl/private_keys/dcex1045mpups01.ext.gfk.pem"
:user: "admin"
:password: "lol"
:puppetdir: "/var/lib/puppet"
:puppetuser: "puppet"
:facts: true
:timeout: 10
:threads: null
```

* Copy `/usr/lib/ruby/vendor_ruby/puppet/reports/foreman.rb` to your puppet
  master's `/usr/lib/ruby/vendor_ruby/puppet/reports/`
* Set `Puppet server` config item to correct master
* Update `Trusted puppetmaster hosts` to correct puppet master (as an
  array!)
* **IMPORTANT** Set `ENC environment` to `No`!

