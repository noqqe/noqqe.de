---
title: OpenLDAP PAM Auth
date: 2013-03-06T07:37:48.000000
tags: 
- Software
- OpenLDAP
---


Administrativer Account anlegen für Modify:

~~~
ldapmodify -a -xWD "cn=admin,dc=example,dc=com" << EOF
> dn: cn=auth,dc=example,dc=com
> objectClass: organizationalRole
> objectClass: simpleSecurityObject
> cn: auth
> userPassword: {SSHA}hNQpve/j+JYBCLG8m29XTx4bJ0Qi7Kmq
> EOF
Enter LDAP Password:
adding new entry "cn=auth,dc=example,dc=com"
~~~

Lesender Account anlegen für Auth und lesen:

~~~
ldapmodify -a -xWD "cn=admin,dc=example,dc=com" << EOF
dn: cn=rauth,dc=example,dc=com
objectClass: organizationalRole
objectClass: simpleSecurityObject
cn: rauth
userPassword: {SSHA}hNQpve/j+JYBCLG8m29XTx4bJ0Qi7Kmq
EOF

Enter LDAP Password:
adding new entry "cn=rauth,dc=example,dc=com"
~~~

Dann das PAM Module installieren

~~~
aptitude install libpam-ldap
~~~

Ich bin da nur dem Dialog gefolgt. Was ich bisher finden konnte:

unterhalb pam.d:

common-account

~~~
account	[success=2 new_authtok_reqd=done default=ignore]	pam_unix.so
account	[success=1 default=ignore]	pam_ldap.so
account	requisite			pam_deny.so
account	required			pam_permit.so
~~~

common-auth

~~~
auth	[success=2 default=ignore]	pam_unix.so nullok_secure
auth	[success=1 default=ignore]	pam_ldap.so use_first_pass
auth	requisite			pam_deny.so
auth	required			pam_permit.so
~~~

common-password

~~~
password	[success=2 default=ignore]	pam_unix.so obscure sha512
password	[success=1 user_unknown=ignore default=die]	pam_ldap.so use_authtok try_first_pass
password	requisite			pam_deny.so
password	required			pam_permit.so
~~~

common-session

~~~
session	[default=1]			pam_permit.so
session	requisite			pam_deny.so
session	required			pam_permit.so
session	required	pam_unix.so
session	optional        pam_ldap.so
~~~

common-session-noninteractive

~~~
session	[default=1]			pam_permit.so
session	requisite			pam_deny.so
session	required			pam_permit.so
session	required	pam_unix.so
session	optional			pam_ldap.so
~~~


In der /etc/pam_ldap.conf

~~~
base dc=example,dc=com
uri ldaps://10.10.0.29/
ldap_version 3
binddn cn=rauth,dc=example,dc=com
bindpw foobar
rootbinddn cn=auth,dc=example,dc=com
pam_password crypt
scope sub
~~~
