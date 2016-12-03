---
date: '2013-02-23T10:25:00'
tags:
- verzeichnisdienste
- openldap
- ldap
- debian
- bash
title: Orks und Zwerge im OpenLDAP
---

Ich setzte mich gerade wegen [LPIC 301](http://www.lpi.org/linux-certifications/programs/lpic-3/exam-301)
etwas mit [OpenLDAP](http://www.openldap.org/) auseinander. Um Dinge zu
testen wie Replikation über syncrepl, Index Aufbau oder Accesslists für
Gruppenadmins brauche ich Datensätze.

Um diese Daten nicht alle von Hand schreiben zu müssen, hatte ich mir überlegt
eine Art Automatismus zu basteln, mit der ich LDIF Files erzeugen kann.
Vor kurzem habe ich (als ich auf der Suche nach einem Hostnamen war) mal
[larisweb.de](http://www.larisweb.de/) gefunden. Dort gibt es tolle Generatoren
für Fantasy Namen die man mit ein bisschen `curl` schoen maschinell auslesen
kann.

``` bash
for x in {1..5} ; do
  curl --silent -d "neuerversuch=Start&menge_eingabe=30" http://www.larisweb.de/tools/namen_gen_zwerg.php \
  | grep "^<tr>" | sed -e 's#<tr><td>##g' -e 's#</td></tr>##g' -e 's#</td><td>#\n#'
done | sed -e 's#&szlig;#ss#g' -e 's#&uuml;#ue#g' -e 's#&auml;#ae#g' -e 's#&ouml;#oe#g' > /tmp/XPEFZKL.txt
```

Leider sind auf der Site keinerlei Informationen zur Lizenz oder ähnlichem
angegeben. .oO(Auch wenn ich somit keinerlei Recht hätte die Namen für
irgendwas zu benutzen sollte es für private Zwecke wohl ok sein, für was
gäbs die Seite sonst?).

Als erstes brauche ich aber eine Gruppe `ou=zwerge,dc=noqqe,dc=de` für die
neuen User in meinem Verzeichnis.

``` bash
./ldapmodify -a -xWD "cn=admin,dc=noqqe,dc=de" << EOF
dn: ou=zwerge,dc=noqqe,dc=de
objectClass: top
objectClass: organizationalUnit
ou: zwerge
EOF
```

Der Einfachheit halber habe ich mich bei den erzeugten Usern für die
`STRUCTURAL` Objektklasse `inetOrgPerson` gepaart mit der `AUXILIARY` Klasse
`posixAccount` entschieden.

``` bash
# GID & UID
LDIFUID=10000
LDIFGID=10000
NAMEFILE=/tmp/XPEFZKL.txt

function ldif () {

    set -- $*

    UNIQ="${1:0:1}$2"
    cat << EOF
dn: uid=$UNIQ,ou=zwerge,dc=noqqe,dc=de
objectClass: top
objectClass: inetOrgPerson
objectClass: posixAccount
uid: $2
uidNumber: $LDIFUID
gidNumber: $LDIFGID
homeDirectory: /home/$UNIQ
loginShell: /bin/bash
cn: $1
sn: $2

EOF

((LDIFUID++))
}

while read name ; do
   ldif "$name"
done < $NAMEFILE
```

Den Output des Skripts am Besten in ein File umleiten und dann mit
ldapmodify zum Directory hinzufügen.

``` bash
./ldapmodify -ac -xWD "cn=admin,dc=noqqe,dc=de" -f zwerge.ldif
```

Wichtig ist dabei das `-c` da ich innerhalb des Scripts keine Prüfung auf
duplicates durchführe. Im Continuous Operation Mode macht OpenLDAP bei
Fehlern einfach weiter mit dem LDIF File.