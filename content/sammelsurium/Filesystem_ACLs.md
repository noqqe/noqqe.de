---
title: Filesystem ACLs
date: 2011-11-28T14:46:45.000000
tags: 
- Filesystems
---


Filesystem muss mit der `acl` option gemountet sein.

Auslesen mit

    getfacl $ORDNER

Setzen (user axel darf rw auf die services):

    setfacl -m u:axel:rw- /etc/services

Setzen (Gruppe biertrinker darf rw auf die services):

    setfacl -m g:biertrinker:rw- /etc/services

Für Mehr einfach hier schaun: http://www.vanemery.com/Linux/ACL/linux-acl.html
Weitere Beispiele

    setfacl -R -m mensch:rwx ordner/
    setfacl -R -m www-data:rwx ordner/
    setfacl -R -d -m www-data:rwx ordner/
    setfacl -R -d -m g:www-data:rwx ordner/
