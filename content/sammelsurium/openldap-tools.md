---
title: OpenLDAP Tools
date: 2013-03-06T18:53:14.000000
tags: 
- Software
- OpenLDAP
---


## Offline Client Tools

Sicherung wie mysqldump:

~~~
slapcat > HOLYSHIT.ldif
slapcat -I HOLYSHIT.ldif
~~~

Hinzufügen der initialen DB

~~~
slapadd -f /usr/local/etc/openldap/slapd.conf -l HOLYSHIT.ldif
~~~

### Indexing

In der DATABASE Section:

~~~
## Indices to maintain
index   objectClass     eq
index   default         eq,pres
index   uid             sub,eq,pres
index   sn,cn
index   departmentNumber        approx
~~~

und danach `slapindex`
