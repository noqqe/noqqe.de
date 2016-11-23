---
title: OpenLDAP Setup
date: 2013-02-15T09:37:33.000000
tags: 
- Software
- OpenLDAP
---


### Debian

~~~
$ aptitude install openldap
~~~

### Download und Compile per Hand

~~~
$ wget ftp://gd.tuwien.ac.at/infosys/network/OpenLDAP/openldap-release/openldap-2.4.33.tgz
$ aptitude install libdb-dev libdb++-dev gcc make libsasl2-dev groff-base time ## wegen soelim bei make
$ ./configure --enable-dynamic --enable-slapd --with-cyrus-sasl --with-tls=openssl --enable-bdb --enable-crypt --enable-syncprov
$ make depend
$ make
$ make test
$ sudo make install
~~~

Für leichtes starten und stoppen

~~~
#!/bin/bash

case $1 in
        start) /usr/local/libexec/slapd -f /usr/local/etc/openldap/slapd.conf  -h "ldap:// ldaps://" ;;
        stop) pkill slapd ;;
        restart) pkill slapd ; /usr/local/libexec/slapd -f /usr/local/etc/openldap/slapd.conf  -h "ldap:// ldaps://"  ;;
esac
~~~

## Initale Konfiguration

Passwort generieren (la)

~~~
$slappasswd -h {SSHA}
New password:
Re-enter new password:
{SSHA}njcljmYoBGW9wk+nj/GUMsZcqheYVYvF
~~~

~~~
$ vim slapd.conf

### General
include         /usr/local/etc/openldap/schema/core.schema
include         /usr/local/etc/openldap/schema/cosine.schema
include         /usr/local/etc/openldap/schema/inetorgperson.schema
defaultsearchbase "dc=example,dc=com"

### Database
database        bdb
suffix          "dc=example,dc=com"
rootdn          "cn=admin,dc=example,dc=com"
rootpw          {SSHA}njcljmYoBBW9wk+nj/GUMsZcqheYVYvF
~~~

### LDAPS SSL/TLS Konfigurieren

Zertifikate erstellen usw.

~~~
$ openssl genrsa -out host.key 2048
$ openssl req -new -nodes -key host.key -out host.csr
$ openssl x509 -req -days 365 -in host.csr -signkey host.key -out host.crt
~~~

In der General Section die Zertifikate einbinden:

~~~
TLSCertificateFile /usr/local/etc/openldap/certs/host.crt
TLSCertificatekeyFile /usr/local/etc/openldap/certs/host.key
~~~

im Client dann noch (ldap.conf):

~~~
TLS_REQCERT never
~~~

SSL

    ldapsearch -xWD "cn=admin,dc=example,dc=com" -H ldaps://localhost

TLS mit Fallback

    ldapsearch -xWD "cn=admin,dc=example,dc=com" -H ldap(s)://localhost -Z

TLS Force

    ldapsearch -xWD "cn=admin,dc=example,dc=com" -H ldap(s)://localhost -ZZ

### Access Rules

~~~
access to *
        by self write
        by users read
        by anonymous auth

access to *
        by dn.exact="cn=repl,dc=example,dc=com"
        by dn.exact="cn=admin,dc=example,dc=com"
        by * break
~~~
