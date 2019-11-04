---
title: OpenLDAP Replikation
date: 2013-02-24T14:04:05
tags: 
- Software
- OpenLDAP
---

### Konfiguration am Master

In der Database Section den syncprov overlay definieren

    overlay syncprov
    syncprov-checkpoint     100     10
    syncprov-sessionlog     100

einen Replikationsuser anlegen

    ./ldapmodify -a -xWD "cn=admin,dc=example,dc=com" << EOF
    dn: cn=repl,dc=example,dc=com
    objectClass: organizationalRole
    objectClass: simpleSecurityObject
    cn: repl
    userPassword: {SSHA}hNQpve/j+JYBCLG8m29XTx4bJ0Qi7Kmq
    EOF

Entsprechende Access Regeln in der slapd.conf angelegen (general Section) und limit

    access to *
            by dn.exact="cn=repl,dc=example,dc=com"
            by * break

und in der Database Section

    limits dn.exact="cn=repl,dc=example,dc=com"
            size=unlimited
            time=unlimited

Um hier bei Replikation alles zu finden

### Konfiguration am Slave

In die Database Section:

    ## Replica Consumer
    syncrepl        rid=001
                    provider=ldaps://10.10.0.29
                    searchbase="ou=zwerge,dc=example,dc=com"
                    scope=sub
                    filter="objectClass=*"
                    attrs="*,+"
                    type=refreshOnly
                    interval="00:00:30:00"
                    retry="10 5 60 +"
                    binddn="cn=repl,dc=example,dc=com"
                    bindmethod=simple
                    credentials="foobar"
                    tls_reqcert=never
                    timelimit=unlimited
                    sizelimit=unlimited

### Überprüfung

Da nur die Gruppe Zwerge repliziert ist einfahc mal nach der ou=users suchen

$ ldapsearch -xWD "cn=admin,dc=example,dc=com" -H ldaps://localhost '(ou=users)'

Mal am Master was ändern:

    ldapmodify -xWD "cn=admin,dc=example,dc=com" << EOF
    dn: uid=vvorschlaghammer,ou=zwerge,dc=example,dc=com
    changetype: modify
    replace: loginShell
    loginShell: /bin/zsh
    EOF

die eingestellt Zeit des Intervall abwarten und kontrollieren:

    ldapsearch -x "(uid=vvorschlaghammer)" -WD "cn=admin,dc=example,dc=com" -LL

CSN verify

    ldapsearch  -LLL -xWD "cn=admin,dc=example,dc=com" -s base -b dc=example,dc=com contextCSN