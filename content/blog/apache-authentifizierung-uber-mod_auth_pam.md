---
aliases:
- /archives/1494
- /blog/2011/03/04/apache-authentifizierung-uber-mod_auth_pam
comments:
- author: Authentifizierung
  content: "<p>Danke f\xFCr die nette Zusammenfassung zum Thema Authentifizierung.
    Werde ich auf jeden Fall weiterempfehlen. Danke, Cassandra.</p>"
  date: '2011-03-08T15:58:29'
- author: Tobias
  content: "Sehr einfache und gutverst\xE4ndliche Anleitung zur Anbindung von pam.\nHat
    mir sehr geholfen.\n\nnoqqe.de werde ich gerne weiterempfehlen\nDanke, Tobias."
  date: '2015-03-16T10:20:21.504465'
- author: Anonymous
  content: "\"F\xFCgt man\ndiesen der Gruppe shadow hinzu, funktioniert die Authentifizierung\neinwandfrei.\"\n\nDas
    ist eine ganz schlechte Idee. Damit wird der isolierte Apache-Benutzer ad absurdum
    gef\xFChrt. Bei einer Sicherheitsl\xFCcke in Apache kann ein Angreifer die gesamten
    Passw\xF6rter manipulieren."
  date: '2016-02-04T12:14:55.948795'
date: '2011-03-04T17:53:19'
tags:
- development
- web
- ftp
- authpam_enabled
- unix_chkpwd
- linux
- vsftpd
- auth_pam
- auth
- apache2
- pam
- ubuntu
- authentication failure
- debian
- bash
title: "Apache | Authentifizierung \xFCber mod_auth_pam"
---

FTP-User, die Dateien auch über HTTP durchsuchen wollen, werden bei vsftpd
häufig über /etc/passwd als System User authentifiziert. Um nicht noch eine
zusätzliche htpasswd Datei pflegen zu müssen, bietet sich das Apache2 Modul
mod_auth_pam an. Allerdings nur wenn man weiss wie.

```
aptitude install libapache2-mod-auth-pam
```

```
$ vim /etc/apache2/sites-availabe/ftp.domain.com
<Location /dir>
AuthType Basic
AuthName "FTP-Auth"
AuthPAM_Enabled On
AuthBasicAuthoritative Off
Require user FTPUSER
</Location>
```

Sehr wichtig an dieser Stelle AuthBasicAuthoritative Off. Ansonsten
Internal Server Error. Die Schnittstelle mit der sich Apache2 gegen PAM
anmeldet, wird automatisch definiert.

```
# cat /etc/pam.d/apache2
@include common-auth
@include common-account
```

Nun ja, auch wenn eigentlich soweit alles klar sein sollte, schmeisst
Apache2 einen Internal Server Error.

```
Mar  1 12:37:59 host unix_chkpwd[2682]: password check failed for user (FTPUSER)
Mar  1 12:37:59 host apache2: pam_unix(apache2:auth): authentication  failure; logname= uid=xx euid=xx tty= ruser= rhost=123.123.123.123   user=FTPUSER
```

Was an den fehlenden Leserechten des Users "www-data" liegt. Fügt man
diesen der Gruppe shadow hinzu, funktioniert die Authentifizierung
einwandfrei.

```
$ vi /etc/group
shadow:x:42:www-data
```

Zusätzliche Hilfe:
[http://pam.sourceforge.net/mod_auth_pam/configure.html](http://pam.sourceforge.net/mod_auth_pam/configure.html)
[http://biblio.l0t3k.net/howto/en/user-authentication-howto/x302.html](http://biblio.l0t3k.net/howto/en/user-authentication-howto/x302.html)
