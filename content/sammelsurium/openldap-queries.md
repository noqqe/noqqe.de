---
title: OpenLDAP Queries
date: 2013-03-06T18:52:38
tags:
- Software
- OpenLDAP
---

### Suchen

Mit simple Bind (-x)

```
./ldapsearch -x -b dc=example,dc=com
./ldapsearch -x -b dc=example,dc=com '(objectClass=organizationalUnit)'

## Ohne Kommentare und nur das Attribut "dn"
./ldapsearch -LLL -x -b dc=example,dc=com '(objectClass=organizationalUnit)' dn
dn: ou=users,dc=example,dc=com

dn: ou=groups,dc=example,dc=com

```

Mit authenticated Bind

```
ldapsearch -x -H ldap://<hostname> -b dc=<company>,dc=com -D "CN=<user>,DC=com"
  -w <password> '(&(objectClass=user)(sAMAccountName=<usertofind>))'
```

### Wer bin ich eigentlich?

```
## ./ldapwhoami -xWD "cn=admin,dc=example,dc=com"
Enter LDAP Password:
dn:cn=admin,dc=example,dc=com

## ./ldapwhoami -x
anonymous
```

### Erste Entries adden in live

Solche Files sehen dann unter umständen so aus:

```
dn: uid=horst,ou=users,dc=example,dc=com
objectClass: top
objectClass: inetOrgPerson
objectClass: posixAccount
uid: horst
uidNumber: 10002
gidNumber: 10001
homeDirectory: /home/horst
loginShell: /bin/bash
cn: Horst
sn: Tappert
```

Als anonymous failed das natürlich(gut!)

``` bash
## ./ldapmodify -a -x -f /usr/local/etc/horst.ldif
adding new entry "uid=horst,ou=users,dc=example,dc=com"
ldap_add: Strong(er) authentication required (8)
        additional info: modifications require authentication
```

Zum Auth als Admin das folgende tun

``` bash
## ./ldapmodify -a -xWD "cn=admin,dc=example,dc=com" -f /usr/local/etc/bkodera.ldif
Enter LDAP Password:
adding new entry "uid=horst,ou=users,dc=example,dc=com"
```

### Entries löschen

``` bash
## ./ldapdelete -xWD "cn=admin,dc=example,dc=com" "uid=horst,ou=users,dc=n0q,dc=org"
Enter LDAP Password:
```

### Löschen mehrerer Einträge

``` bash
## ./ldapdelete -xWD "cn=admin,dc=example,dc=com" << EOF
> uid=bbergebunker,ou=zwerge,dc=example,dc=com
> uid=vbergebunker,ou=zwerge,dc=example,dc=com
> uid=iisenhaut,ou=zwerge,dc=example,dc=com
> uid=gisenhaut,ou=zwerge,dc=example,dc=com
> uid=ngraubart,ou=zwerge,dc=example,dc=com
> uid=fgraubart,ou=zwerge,dc=example,dc=com
> uid=cgoldlob,ou=zwerge,dc=example,dc=com
> uid=ngoldlob,ou=zwerge,dc=example,dc=com
> EOF
```

### Entries modifizieren

#### Von File

modify.ldif:

``` bash
dn: uid=horst,ou=users,dc=example,dc=com
changetype: modify
replace: loginShell
loginShell: /bin/sh
```

Einspielen via

``` bash
./ldapmodify -xWD "cn=admin,dc=example,dc=com" -f
/usr/local/etc/horst-modify.ldif
```

#### Interactive Mode

```
## ./ldapmodify -xWD "cn=admin,dc=example,dc=com" << EOF
> dn: uid=horst,ou=users,dc=example,dc=com
> changetype: modify
> replace: uidNumber
> uidNumber: 10042
> EOF
Enter LDAP Password:
modifying entry "uid=horst,ou=users,dc=example,dc=com"
```

### Relative disinguished  Name (RDN) anpassen

```
## ./ldapmodrdn -xWD "cn=admin,dc=example,dc=com" "uid=horsty,ou=users,dc=n0q,dc=org" uid=horst
Enter LDAP Password:
```

## Neue Gruppe anlegen

Neues .ldif File anlegen

```
dn: ou=zwerge,dc=example,dc=com
objectClass: top
objectClass: organizationalUnit
ou: zwerge
```

```
## ldapmodify -axWD "cn=admin,dc=example,dc=com" << EOF
dn: ou=oberzwerge,dc=example,dc=com
objectClass: top
objectClass: organizationalUnit
ou: oberzwerge
EOF
Enter LDAP Password:
adding new entry "ou=oberzwerge,dc=example,dc=com"
```

Modify mit -a für add

```
./ldapmodify -a -xWD "cn=admin,dc=example,dc=com" -f /usr/local/etc/ldif/zwerge-group.ldif
./ldapsearch -x
```

### LDAP passwd ändern

```
 ldappasswd -xWD "cn=admin,dc=example,dc=com" -s foo "uid=vvorschlaghammer,ou=zwerge,dc=n0q,dc=org"
```

oder interaktiv mit -S

### Ein Alias Objekt erstellen

```
## ldapmodify -axWD "cn=admin,dc=example,dc=com" << EOF
dn: uid=vvorschlaghammer,ou=oberzwerge,dc=example,dc=com
objectclass: alias
objectclass: extensibleObject
uid: vvorschlaghammer
aliasedobjectname: uid=vvorschlaghammer,ou=zwerge,dc=example,dc=com
EOF
```

### Erweiterte Konfiguration der Clienttools

ldap.conf

```
BASE    dc=example,dc=com
URI     ldap://localhost

## Super zum Testen:
#SIZELIMIT       3
SIZELIMIT       50
TIMELIMIT       15
DEREF           never
```

### Gruppen Admins einrichten

erstmal das password durch den Admin setzen für vvorschlanghammer.

```
## ldappasswd -xWD "uid=vvorschlaghammer,ou=zwerge,dc=example,dc=com"
Enter LDAP Password:
New password: xxx
```

Bisher sieht ein Modify so aus:

```
## ldapmodify -xWD "uid=vvorschlaghammer,ou=zwerge,dc=example,dc=com" << EOF
> dn: uid=jaxthieb,ou=zwerge,dc=example,dc=com
> changetype: modify
> replace: loginShell
> loginShell: /bin/zsh
> EOF
Enter LDAP Password:
modifying entry "uid=jaxthieb,ou=zwerge,dc=example,dc=com"
ldap_modify: Insufficient access (50)
```

Das heisst wir brauchen erstmal Admin Permissions für den User vvorschlaghammer.
Mit folgenden Access rules funktioniert das:

```
access to dn.children="ou=zwerge,dc=example,dc=com"
        by dn.exact="uid=vvorschlaghammer,ou=zwerge,dc=example,dc=com" write
        by self write
        by users read
        by * auth

access to *
        by self write
        by users read
        by * auth

access to *
        by dn.exact="cn=repl,dc=example,dc=com"
        by * break
```

Test mit Modify:

```
## ldapmodify -w foo -xD "uid=vvorschlaghammer,ou=zwerge,dc=example,dc=com" << EOF
dn: uid=jaxthieb,ou=zwerge,dc=example,dc=com
changetype: modify
replace: loginShell
loginShell: /bin/zsh
EOF
modifying entry "uid=jaxthieb,ou=zwerge,dc=example,dc=com"
```
