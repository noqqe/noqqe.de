---
title: LPIC 301 - Cheatsheet
date: 2013-03-06T19:21:04
tags:
- Software
- OpenLDAP
---

#### Entstehungsgeschichte LDAP

* Heilloses durcheinander in den 80ern
* ITU (vorher CCITT) beschliesst den X.500 Standard
* DSP -> Directory System Protocol
* DAP -> Directory Access Protocol
* DSA -> Directory Server Agent
* DUA -> Directory User Agent
* X.500 mit TCP/IP nicht möglich
* Middleware für TCP/IP Clients

#### Unterschiede DAP und LDAP

* Abgespeckte auf TCP IP angepasster Directory Dienst

* LDAP => TCP/IP
* DAP => OSI

#### Unterschiede relationale DB und LDAP

* LDAP für lesende Zugriffe optimiert
* LDAP hat keine Transaktionen
* LDAP ist für Verteilung gebaut
* RD meist auf einer Maschine

#### LDIF Dateien

Neues Personen Objekt

~~~
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
~~~

Neue Gruppe

~~~
dn: ou=zwerge,dc=example,dc=com
objectClass: top
objectClass: organizationalUnit
ou: zwerge
~~~

Modify Kodierung

~~~
dn: uid=horst,ou=users,dc=example,dc=com
changetype: modify
replace: uidNumber
uidNumber: 10042
~~~

#### Schema Dateien

wichtige Dateien:

~~~
/usr/local/etc/openldap/schema/core.schema
/usr/local/etc/openldap/schema/cosine.schema
/usr/local/etc/openldap/schema/inetorgperson.schema
~~~

#### lokale Struktur vs. Domain Struktur

* lokale Struktur zB -> cn=frank,ou=Mitarbeiter,o=Firma GmbH,l=Nuernberg,c=de
* Domain Struktur zB -> uid=frank,ou=users,dc=example,dc=com

#### Schemata

Grundsätzliche Schemata

* core.schema (grundlegendes)
* cosine.schema (COSINE und Internet-X.500 Schema)
* inetorgperson.schema (Personenbezogene Schemata)
* misc.schema (experimenteller Shit!)
* nis.schema (Network Information Service Schema)
* openldap.schema (LDAP spezifischer Kram)

Abhängigkeiten

~~~
core.schema <- cosine.schema <- nis.schema
core.schema <- cosine.schema <- inetorgperson.schema
core.schema <- cosine.schema <- inetorgperson.schema <- openldap.schema
~~~

#### Objektklassen

Beispiel Objektklassen

STRUCTURAL Classes (strukturelle)

* person
* organizationalPerson
* inetOrgPerson
* organizationalUnit
* groupOfNames
* alias

AUXILIARY Classes (ergänzende)

* posixAccount
* posixGroup
* extensibleObject

ABSTRACT Classes (nur aus Vererbungsgründen)

* top

Felder

* OID (2.16.840.1.113739.3.2.3)
* NAME
* DESC
* SUP organizationalPerson
* TYPE (STRUCTURAL/AUXILIARY)
* MUST ( 1 $ 2 $ )
* MAY ( 1 $ 2 $ 3 )

#### Attributstypen

Stichwort im Schema lautet `attributetype`

Felder

* NAME
* DESC
* EQUALITY
* ORDERING
* SUBSTR
* SYNTAX

#### configure Parameter

Beim bauen hab ich

~~~
./configure --enable-dynamic --enable-slapd --with-cyrus-sasl --with-tls=openssl --enable-bdb --enable-crypt --enable-syncprov
~~~

verwendet. Hier noch eine Übersicht

    --bindir=DIR            user executables [EPREFIX/bin]
    --sbindir=DIR           system admin executables [EPREFIX/sbin]
    --libexecdir=DIR        program executables [EPREFIX/libexec]
    --sysconfdir=DIR        read-only single-machine data [PREFIX/etc]
    --sharedstatedir=DIR    modifiable architecture-independent data [PREFIX/com]
    --localstatedir=DIR     modifiable single-machine data [PREFIX/var]
    --libdir=DIR            object code libraries [EPREFIX/lib]
    --includedir=DIR        C header files [PREFIX/include]
    --oldincludedir=DIR     C header files for non-gcc [/usr/include]
    --datarootdir=DIR       read-only arch.-independent data root [PREFIX/share]
    --datadir=DIR           read-only architecture-independent data [DATAROOTDIR]
    --infodir=DIR           info documentation [DATAROOTDIR/info]
    --localedir=DIR         locale-dependent data [DATAROOTDIR/locale]
    --mandir=DIR            man documentation [DATAROOTDIR/man]
    --docdir=DIR            documentation root [DATAROOTDIR/doc/PACKAGE]
    --htmldir=DIR           html documentation [DOCDIR]
    --dvidir=DIR            dvi documentation [DOCDIR]
    --pdfdir=DIR            pdf documentation [DOCDIR]
    --psdir=DIR             ps documentation [DOCDIR]
    --program-prefix=PREFIX            prepend PREFIX to installed program names
    --program-suffix=SUFFIX            append SUFFIX to installed program names
    --program-transform-name=PROGRAM   run sed PROGRAM on installed program names
    --build=BUILD     configure for building on BUILD [guessed]
    --host=HOST       cross-compile to build programs to run on HOST [BUILD]
    --target=TARGET   configure for building compilers for TARGET [HOST]
    --disable-option-checking  ignore unrecognized --enable/--with options
    --disable-FEATURE       do not include FEATURE (same as --enable-FEATURE=no)
    --enable-FEATURE[=ARG]  include FEATURE [ARG=yes]
    --enable-debug    enable debugging no|yes|traditional [yes]
    --enable-dynamic    enable linking built binaries with dynamic libs [no]
    --enable-syslog   enable syslog support [auto]
    --enable-proctitle    enable proctitle support [yes]
    --enable-ipv6     enable IPv6 support [auto]
    --enable-local    enable AF_LOCAL (AF_UNIX) socket support [auto]
    --enable-slapd    enable building slapd [yes]
    --enable-dynacl   enable run-time loadable ACL support (experimental) [no]
    --enable-aci    enable per-object ACIs (experimental) no|yes|mod [no]
    --enable-cleartext    enable cleartext passwords [yes]
    --enable-crypt    enable crypt(3) passwords [no]
    --enable-lmpasswd   enable LAN Manager passwords [no]
    --enable-spasswd    enable (Cyrus) SASL password verification [no]
    --enable-modules    enable dynamic module support [no]
    --enable-rewrite    enable DN rewriting in back-ldap and rwm overlay [auto]
    --enable-rlookups   enable reverse lookups of client hostnames [no]
    --enable-slapi        enable SLAPI support (experimental) [no]
    --enable-slp          enable SLPv2 support [no]
    --enable-wrappers   enable tcp wrapper support [no]
    --enable-backends   enable all available backends no|yes|mod
    --enable-bdb    enable Berkeley DB backend no|yes|mod [yes]
    --enable-dnssrv   enable dnssrv backend no|yes|mod [no]
    --enable-hdb    enable Hierarchical DB backend no|yes|mod [yes]
    --enable-ldap   enable ldap backend no|yes|mod [no]
    --enable-mdb    enable mdb database backend no|yes|mod [yes]
    --enable-meta   enable metadirectory backend no|yes|mod [no]
    --enable-monitor    enable monitor backend no|yes|mod [yes]
    --enable-ndb    enable MySQL NDB Cluster backend no|yes|mod [no]
    --enable-null   enable null backend no|yes|mod [no]
    --enable-passwd   enable passwd backend no|yes|mod [no]
    --enable-perl   enable perl backend no|yes|mod [no]
    --enable-relay      enable relay backend no|yes|mod [yes]
    --enable-shell    enable shell backend no|yes|mod [no]
    --enable-sock   enable sock backend no|yes|mod [no]
    --enable-sql    enable sql backend no|yes|mod [no]
    --enable-overlays   enable all available overlays no|yes|mod
    --enable-accesslog    In-Directory Access Logging overlay no|yes|mod [no]
    --enable-auditlog   Audit Logging overlay no|yes|mod [no]
    --enable-collect    Collect overlay no|yes|mod [no]
    --enable-constraint   Attribute Constraint overlay no|yes|mod [no]
    --enable-dds      Dynamic Directory Services overlay no|yes|mod [no]
    --enable-deref    Dereference overlay no|yes|mod [no] --enable-dyngroup   Dynamic Group overlay no|yes|mod [no] --enable-dynlist    Dynamic List overlay no|yes|mod [no] --enable-memberof   Reverse Group Membership overlay no|yes|mod [no]
    --enable-ppolicy    Password Policy overlay no|yes|mod [no]
    --enable-proxycache   Proxy Cache overlay no|yes|mod [no]
    --enable-refint   Referential Integrity overlay no|yes|mod [no]
    --enable-retcode    Return Code testing overlay no|yes|mod [no]
    --enable-rwm          Rewrite/Remap overlay no|yes|mod [no]
    --enable-seqmod   Sequential Modify overlay no|yes|mod [no]
    --enable-sssvlv   ServerSideSort/VLV overlay no|yes|mod [no]
    --enable-syncprov   Syncrepl Provider overlay no|yes|mod [yes]
    --enable-translucent  Translucent Proxy overlay no|yes|mod [no]
    --enable-unique       Attribute Uniqueness overlay no|yes|mod [no]
    --enable-valsort      Value Sorting overlay no|yes|mod [no]
    --enable-static[=PKGS]  build static libraries [default=yes]
    --enable-shared[=PKGS]  build shared libraries [default=yes]
    --enable-fast-install[=PKGS]
    --disable-dependency-tracking  speeds up one-time build
    --enable-dependency-tracking   do not reject slow dependency extractors
    --disable-libtool-lock  avoid locking (might break parallel builds)
    --with-PACKAGE[=ARG]    use PACKAGE [ARG=yes]
    --without-PACKAGE       do not use PACKAGE (same as --with-PACKAGE=no)
    --with-subdir=DIR       change default subdirectory used for installs
    --with-cyrus-sasl   with Cyrus SASL support [auto]
    --with-fetch      with fetch(3) URL support [auto]
    --with-threads    with threads [auto]
    --with-tls      with TLS/SSL support auto|openssl|gnutls|moznss [auto]
    --with-yielding-select  with implicitly yielding select [auto]
    --with-mp               with multiple precision statistics auto|longlong|long|bignum|gmp [auto]
    --with-odbc             with specific ODBC support iodbc|unixodbc|odbc32|auto [auto]
    --with-gnu-ld           assume the C compiler uses GNU ld [default=no]
    --with-pic              try to use only PIC/non-PIC objects [default=use
    --with-tags[=TAGS]      include additional configurations [automatic]

#### Backends

Datenspeicher Backends

* BDB - Berkley Database
* HDB - Derivat der BDB - Hierarchisch, viel RAM
* LDBM - Nicht mehr unterstützt
* LDIF - File Format Sicherung im LDIF Format
* SQL - Experimentell

Proxy Backends

* ldap - leitet an anderen LDAPServer weiter
* meta - leitet an mehrere LDAPServer weiter
* relay - Durch Overlay realisierte Weiterleitung von Teilbäumen

Misc Backends

* monitor - Zustand des DIT
* dnssrv - Auflösung von REFERALs durch SRV Records in DNS
* shell/perl - Scripting

#### Database

* Spezielle Einstellungen für kontretes Backend

#### BerkleyDB

* Übersicht aller Datenbanken:

~~~
db_stat -h /usr/local/var/openldap-data/ -m
~~~

* Konkrete Datenbanken auslesen

~~~
db_stat -h /usr/local/var/openldap-data/ -d dn2id.bdb
db_stat -h /usr/local/var/openldap-data/ -d id2entry.bdb
~~~

* Korrekte Cachesize errechnen

~~~
## db_stat -h /usr/local/var/openldap-data/ -d dn2id.bdb
Fri Mar  8 09:11:16 2013  Local time
53162 Btree magic number
9 Btree version number
Little-endian Byte order
duplicates, sorted duplicates Flags
2 Minimum keys per-page
4096  Underlying database page size
1007  Overflow key/data size
2 Number of levels in the tree
1149  Number of unique keys in the tree
2281  Number of data items in the tree
3 Number of tree internal pages
11444 Number of bytes free in tree internal pages (6% ff)
29  Number of tree leaf pages
48174 Number of bytes free in tree leaf pages (59% ff)
4 Number of tree duplicate pages
642 Number of bytes free in tree duplicate pages (96% ff)
0 Number of tree overflow pages
0 Number of bytes free in tree overflow pages (0% ff)
0 Number of empty pages
0 Number of pages on the free list
~~~

Die wichtigen Angaben nochmal in kurz

~~~
3 Number of tree internal pages
29  Number of tree leaf pages
~~~

Blockgröße des Dateisystems: 4KB

Formel:

    ( 1 root Page + 3 internal Pages + 29 leaf Pages) * 4KB Blocksize = 132 KB Cache Size

~~~
root@vm29-ldap:~## db_stat -h /usr/local/var/openldap-data/ -d id2entry.bdb
Fri Mar  8 09:30:37 2013  Local time
53162 Btree magic number
9 Btree version number
Little-endian Byte order
  Flags
2 Minimum keys per-page
16384 Underlying database page size
4079  Overflow key/data size
2 Number of levels in the tree
573 Number of unique keys in the tree
573 Number of data items in the tree
1 Number of tree internal pages
15728 Number of bytes free in tree internal pages (4% ff)
29  Number of tree leaf pages
45622 Number of bytes free in tree leaf pages (90% ff)
0 Number of tree duplicate pages
0 Number of bytes free in tree duplicate pages (0% ff)
0 Number of tree overflow pages
0 Number of bytes free in tree overflow pages (0% ff)
0 Number of empty pages
0 Number of pages on the free list
~~~

bei der

    ( 1 root Page + 3 internal Pages + 29 leaf Pages) * 4KB Blocksize = 132 KB Cache Size

Einbau

    DB_CONFIG
    set_cachesize 2 524288000 1
    dbconfig set_cachesize 2 524288000 1

#### net::ldap Perl Modul

     $ldap = Net::LDAP->new( 'ldap.bigfoot.com' ) or die "$@";

     $mesg = $ldap->bind ;    ## an anonymous bind

     $mesg = $ldap->search( ## perform a search
                            base   => "c=US",
                            filter => "(&(sn=Barr) (o=Texas Instruments))"
                          );

     $result = $ldap->add( 'cn=Barbara Jensen, o=University of Michigan, c=US',
                           attrs => [
                             'cn'   => ['Barbara Jensen', 'Barbs Jensen'],
                             'sn'   => 'Jensen',
                             'mail' => 'b.jensen@umich.edu',
                             'objectclass' => ['top', 'person',
                                               'organizationalPerson',
                                               'inetOrgPerson' ],
                           ]
                         );

     $mesg = $ldap->unbind;   ## take down session

#### Ablauf einer LDAP Session

BIND ->
<- BIND-RESPONSE
SEARCH REQUEST ->
<- SEARCH RESPONSE
UNBIND (eig BIND) ->
<- BIND RESPONSE

#### ACL

Reihenfolge matters, a lot!

Generell

```
access to <was>
       by <wer> <wie>
```

Beispiel

~~~
access to *
        by self write
        by dn.exact="cn=repl,dc=example,dc=com" read
        by dn.exact="cn=auth,dc=example,dc=com" read
        by users read
        by * auth
~~~

Access level und deren Privileges

* 0 - none (keine Permissions)
* d - disclose (Erlaubt Fehlermeldungen anzuzeigen)
* x - auth (Zugriff für Auth auf userPassword oder ähnliches)
* c - compare (vergleichen)
* s - search (NICHT! Suchergebnisse ausgeben, sondern Recht mit Suchfiltern zu suchen9
* r - read (ausgabe von Suchergebnissen)
* w - write (schreiben

Beispiel für nutzung von Privileges

~~~
access to dn.base=ou=manager usw
       by self +rw
       by * +x
~~~

#### DN Styles

* dn.exact
* dn.one
* dn.sub
* dn.children
* dn.regex

#### slurpd

* Alle Änderungen werden in ein Logfile geschrieben
* SLURPD ist ein separater Prozess
* Änderungen werden AN(!) den slave geschickt
* Nur Replikation des ganzen DIT möglich

##### Einrichtung

Am Master

~~~
replogfile $file
replica uri=""
        binddn=""
        bindmethod=simple
        credentials=secret
        starttls=yes
~~~

Am Slave

~~~
updatedn "cn=repl..." (User füer Schreibende Zugriffe)
updateref "ldap://" (Host an den schreibzugriffe weitergeleitet werden
~~~

Danach Inital dump erstellen und dann einspielen:

~~~
slapcat > master.ldif
slapadd -f slapd.conf -l master.ldif
~~~

danach

* start slapd master
* start slapd slave
* start slurpd master

#### Slurpd Verhalten im Fehlerfall

* Es wird eine .reject Datei erstellt und slurpd bleibt stehen

* Manuell die repliaktion wieder starten:

~~~
slurpd -r /usr/local/var/openldap-slurp/replica/slave.example.com\:389.rej -o
~~~

Auch one-Shot mode genannt.

#### syncrepl

* refreshOnly (Aktiver Modus)
  Slave ist normaler User der regelmäßig entsprechende Queries durch den DIT des Masters triggert

* refreshAndPersist (Passiver Modus)
  Slave baut einmal Verbindung auf und behält diese. Master Pushed dann automatisch Änderungen zum Slave

* Replikation nur bestimmter Teilbäume des DIT möglich

* Konfiguration am Master

~~~
overlay syncprov
syncprov-checkpoint     100     10
syncprov-sessionlog     100

limits dn.exact="cn=repl,dc=example,dc=com" size=unlimited time=unlimited
~~~

* Konfiguration am Slave

~~~
## Replica Consumer
syncrepl        rid=001
                provider=ldaps://10.10.0.29
                searchbase="ou=zwerge,dc=example,dc=com"
                scope=sub
                filter="objectClass=*"
                attrs="*,+"
                type=refreshOnly
                interval="00:00:02:00"
                retry="10 5 60 +"
                binddn="cn=repl,dc=example,dc=com"
                bindmethod=simple
                credentials="foobar"
                tls_reqcert=never
                timelimit=unlimited
                sizelimit=unlimited
~~~

#### Verschlüsselte Verbindung zum LDAP

* Serverseitig

~~~
TLSCertificateFile
TLSCertificateKeyFile
TLSCACertificateFile
TLSVerfiyClient never,allow,try,demand(hard)
~~~

* Clientseitig

~~~
TLS_CACERT /pfad/ (allgemein)
TLS_REQCERT never,allow,try,demand(hard)

## mit client auth
TLS_CERT /pfad/
TLS_KEY /pfad/
~~~

Verbindungsmethoden

* ldaps:// (SSL
* -Z (TLS, fallback auf unverschlüsselt)
* -ZZ (Ohne Fallback mit abbruch der conn)

#### security Methoden

* ssf - Mindestanforderung der Länge
* transport - Mindestlänge des Schlüssels über Netzwerk
* tls - Mindestanforderung über TLS
* sasl
* update_ssf - Min. für schreibende Zugriffe
* update_transport * Schreibend übers Netzwerk
* update_tls - Min für schreibende TLS connections
* update_sasl - Min für schreibende SASL Auth
* simple_bind - Min für simple Bind Verschlüsselung

#### SASL

* Mechs
    * PLAIN
    * CRAM-MD5
    * DIGEST-MD5
    * GSSAPI (Kerberos)
    * EXTERNAL

* Bind durch: -Y DIGEST-MD5 -U user@example.com

#### Wichtige Konfigurationsparameter slapd.conf

* access - Zugriffsbeschränkungen
* allow - aktiviert Funktionen in openLDAP
* authz-policy - Proxy Auth Mapping
* authz-regexp - Proxy Auth Mapping
* defaultsearchbase  - Standardmäßige Suche in Verzeichnisbaum
* disallow - deaktiviert Funktionen in OpenLDAP
* include - includiert Schemata oder auch andere Komponenten
* loglevel - definiert das loglevel in INT beim Start
* moduleload / modulepath
* pidfile
* referal - Wird an client übermittelt, wenn Query nicht teil des lokalen Trees
* require - z.B LDAPv3 erzwingen
* rootDSE - Infos über den LDAPServer
* security - Sicherheitsauflagen wie ssf
* sizelimit - Limit aller Results in einer Abfrage
* timelimit - Zeitliches Limit

#### Loglevel

~~~
0 - none
1 - trace (Systemaufrufe)
2 - packets
4 - args (zusätzliche Funktionsargumente)
8 - conns (Loggt verbindungen)
16 - BER (alle Pakete ausgeben)
32 - filter (Suchfilter debugging)
64 - config
128 - ACL Debugging
256 - stats
512 - stats2
1024 - shell (backend debugging)
2048 - parse (entry parsing)
4096 - cache
8192 - index
16384 - sync
-1 - all
~~~

#### Referals in LDAP

* Referenzieren auf andere Verzeichnisbäume
* Inhalt der nicht im lokalen DIT ist wird weitergeleitet

~~~
DEREF never
~~~

#### yp Dienste / white Pages

* Befehle

~~~
ypcat -x                 ## alle Maps ausgeben
ypcat passwd             ## passwd NIS ausgeben
yptest                   ## Testet Verbindung zum YP Server
/etc/yp.conf             ## enthält ypserver nis.example.com)
ypwhich -d exaple.com    ## YPServer finden
~~~

Typische Maps

* passwd
* group
* hosts
* networks
* protocols
* services
* fstab
* aliases
* rpc
* profile
* netgroup
* ethers
* netmasks
* bootparams

Schema Dateien

* nis.schema

ypldapd Konfiguration

* ypdomain - NIS DOMAIN
* ldaphost - hosts
* basedn
* binddn
* bindcred
* ldapversion

#### Pam ldap Modul

Grober umriss

#### nsswitch

/etc/nsswitch.conf

#### NIS

* Struktur
* wichtige NIS Kommandos ausprobieren
* NIS Migrationsskripte on und offline

#### Integration

##### ssh

Einfache Einbindung in die /etc/pam.d/sshd

##### FTP

Ebenfalls Einbindung über pam

##### httpd

auth-ldap (Zum Authentifizieren)
Beispiel:

~~~
<Directory /var/www>
 AuthType Basic
 AuthName Restricted
 AuthBasicProvider ldap
 AuthLDAPUrl ldap://127.0.0.1:389/ou=users,dc=example,dc=com?uid?one?objectClass=posixAccount
 require valid-user
</Directory>
~~~

mod-ldap (Cache tuning zum LDAP Server)

* LDAPCacheEntries 1024
* LDAPCacheTTL 600
* LDAPOpCacheEntries 1024 (bezieht sich auf Vergleichsoperationen)
* LDAPOpCacheTTL 600 (bezieht sich auf Vergleichsoperationen)

##### freeradius

Anbindung über Modules

~~~
modules {
 ldap {
  server = "ldap.example.com"
  basedn =
  filter =
  ldap_connections_number = 5
  timeout = 4

  tls {
   start_tls=no
   cacertfile =
   cacertdir =
   certfile =
   keyfile =
   require_cert = "demand"
  }
 }
}
~~~

##### Cups

den scheiss tu ich mir nicht an.

##### Postfix

Über Postmap Files gelöst

~~~
cat /etc/postfix/
server_host = ldap://
search_base = dc=...
query_filter = (mail=%u)
result_attribute = uid
version = 3
start_tls = yes
tls_ca_cert_file =
tls_cert_file =
tls_key_file =
tls_require_cert = yes
bind = sasl
sasl_mechs = EXTERNAL
~~~

##### Sendmail

ObjectClasses:

* person
* inetLocalMailRecipient
* qmailUser
* sendmailMTA
* sendmailMTAAlias
* sendmailMTAAliasObject
* sendmailMTAMap
* sendmailMTAMapObject

#### ldap zu Samba

Schemata samba3.schema

ObjectClasses

* sambaSamAccount
* sambaGroupMapping

smb.conf

~~~
passdb backend = ldapsam:ldap://slapd.example.com
ldap admin dn = "cn..."
ldap ssl = start tls
ldap user suffix = ou=users
ldap group suffix = ou=groups
ldap machine suffix = ou=computers
ldap suffix = dc=example=com
ldap filter = (&(uid\%u) (objectclass=sambaSamAccount)
~~~

pbedit

~~~
pbedit -Lv noqqe     ## user infos aus smb backend
pbedit -Lv rechner$  ## rechner infos
pbedit -Lw noqqe     ## altes smbpasswd format
pbedit -a noqqe      ## neuen benutzer anlegen

pbedit -i smbpasswd:/etc/samba/smbpasswd -e ldapsam:ldap://slapd.example.com   ## import smbpasswd und export nach ldap
~~~

#### Active Directory und LDAP

Funktion von AD (KERBEROS)

KINIT Command

-> TGT (Ticket Granting Ticket)
-> TGS (Ticket Granting Service)
-> TGS (Ticket Granting Service an Server übermitteln)

Anbindung in ldap.conf des PAM Moduls

~~~
pam_login_attribute sAMAccountName
pam_filter objectclass=User
pam_password ad
~~~

Anbindung von Linux an ad (was braucht man?)

pam_krb5 Modul:

~~~
auth sufficient pam_krb5.so
session required pam_krb5.so
account required pam_krb5.so
password sufficient pam_krb5.so
~~~

/etc/krb5.conf

~~~
[libdefaults]
default_realm = EXAMPLE.COM
dns_lookup_realm = true

[realms]
EXAMPLE.COM = {
 kdc = win.example.com:88
 admin_server = win.example.com:749
 default_domain = example.com
}

[domain_realm]
example.com = EXAMPLE.COM
.example.com = EXAMPLE.COM

[appdefaults]
pam = {
 debug = false
 ticket_lifetime = 36000
 renew_lifetime = 36000
 forewardable = true
 krb4_convert = false
}
~~~

#### Kapazitätsplanung

vmstat

~~~
vmstat 5
~~~

iostat

~~~
iostat -x 1
iostat -x
~~~

sar

~~~
sar -r    #Speicheraulsastung
sar -n DEV   ## Netzwerkauslastung
sar -d -p  #Festplattenauslastung
sar -f /path/ #bestimmtes file nutzen
sar -s 10:00:00 -e 12:00:00
sar -A     ## alles des aktuellen tages
~~~
