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

## Apache2 Passenger

~~~
<VirtualHost *:443>
        SSLEngine on
        SSLProtocol -ALL +SSLv3 +TLSv1
        SSLCipherSuite ALL:!ADH:RC4+RSA:+HIGH:+MEDIUM:-LOW:-SSLv2:-EXP

        SSLCertificateFile      /var/lib/puppet/ssl/certs/ip-10-2-0-254.example.com.pem
        SSLCertificateKeyFile   /var/lib/puppet/ssl/private_keys/ip-10-2-0-254.example.com.pem
        SSLCertificateChainFile /var/lib/puppet/ssl/certs/ca.pem
        SSLCACertificateFile    /var/lib/puppet/ssl/certs/ca.pem
        ## If Apache complains about invalid signatures on the CRL, you can try disabling
        ## CRL checking by commenting the next line, but this is not recommended.
        SSLCARevocationFile     /var/lib/puppet/ssl/ca/ca_crl.pem
        SSLVerifyClient optional
        SSLVerifyDepth  1
        ## The `ExportCertData` option is needed for agent certificate expiration warnings
        SSLOptions +StdEnvVars +ExportCertData

        ## This header needs to be set if using a loadbalancer or proxy
        RequestHeader unset X-Forwarded-For

        RequestHeader set X-SSL-Subject %{SSL_CLIENT_S_DN}e
        RequestHeader set X-Client-DN %{SSL_CLIENT_S_DN}e
        RequestHeader set X-Client-Verify %{SSL_CLIENT_VERIFY}e

        ## you probably want to tune these settings
        PassengerHighPerformance on
        PassengerMaxPoolSize 12
        PassengerPoolIdleTime 1500
        ## PassengerMaxRequests 1000
        PassengerStatThrottleRate 120


        DocumentRoot /usr/share/foreman/public
        PassengerAppRoot /usr/share/foreman
        RackAutoDetect On
        RailsAutoDetect On
        AddDefaultCharset UTF-8

        <Directory /usr/share/foreman/public>
                <IfVersion < 2.4>
                 Allow from all
                </IfVersion>
                <IfVersion >= 2.4>
                 Require all granted
                </IfVersion>
        </Directory>

</VirtualHost>
~~~
