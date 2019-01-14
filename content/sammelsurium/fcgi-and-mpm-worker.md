---
title: fcgi and mpm-worker
date: 2012-03-06T13:54:01
tags:
- Software
- Apache
---

fcgi ist die Weiterentwicklung von fastcgi.

    aptitude install apache2-mpm-worker php5-fcgi

## Handler einfuegen

```
<Directory /var/www/>
        Options -Indexes FollowSymLinks -MultiViews +ExecCGI
        AddHandler fcgid-script .php
        FCGIWrapper "/usr/lib/cgi-bin/php5" .php
        AllowOverride None
        Order allow,deny
        allow from all
</Directory>
```

## TimeOuts oder Exceeds

Kann sein das die Max Request Len zu kurz ist für Files zum Uploaden
Dann will mann in `/etc/apache2/mods-available/fcgid.conf`

```
<IfModule mod_fcgid.c>
  AddHandler    fcgid-script .fcgi
  FcgidConnectTimeout 20
  MaxRequestLen 15728640
</IfModule>
```

## Links

[Apache FCGID Performace](http://2bits.com/articles/apache-fcgid-acceptable-performance-and-better-resource-utilization.html)

## puppet fastcgi & worker & php

``` { .puppet }
 class { 'apache':
        mpm_module => "worker",
        default_vhost => false,
 }
 apache::mod { 'fcgid': }

 package { 'php5-cgi':
        ensure => installed,
 }

 apache::vhost { 'localhost':
   port => '80',
   docroot => '/data',
   directories => [
        { path        => '/data',
          addhandlers => [{ handler => 'fcgid-script', extensions => ['.php']}],
          options    => ['Indexes','FollowSymLinks','MultiViews', 'ExecCGI'],
          custom_fragment => 'FCGIWrapper "/usr/lib/cgi-bin/php5" .php',
        }, ]
 }
```
