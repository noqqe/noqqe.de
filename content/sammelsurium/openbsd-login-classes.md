---
title: OpenBSD Login Classes
date: 2016-05-25T15:22:49
tags:
- OpenBSD
- ulimit
- class
- login.conf
---

Hiermit lassen sich ziemlich flexibel und schön
einige Limitierungen pro User erledigen.

## Eine Login Class

```
$ sudo vim /etc/login.conf
staff:\
      :umask=022:\
      :datasize-max=1024M:\
      :datasize-cur=1024M:\
      :maxproc-max=1024:\
      :maxproc-cur=1024:\
      :openfiles-cur=2048:\
      :openfiles-max=2048:\
      :stacksize-cur=4M:\
      :localcipher=blowfish,8:\
      :tc=auth-defaults:\
      :tc=auth-ftp-defaults:
```

Aber welche Klasse hat mein User eigentlich?

    $ userinfo noqqe
    login   noqqe
    passwd  *
    uid     1001
    groups  noqqe wheel staff
    change  NEVER
    class
    gecos   ,,,
    dir     /home/noqqe
    shell   /usr/local/bin/bash
    expire  NEVER

Aha garkeine also. Wie assigne ich die?

    $ chsh noqqe
    ...
    class: staff
    ...

Gut. Jetzt hab ich die. Nachdem mir die Werte immernoch nicht gefallen, muss ich die
login.conf doch noch anpassen.

    $ vim /etc/login.conf
    ## build new db
    $ sudo cap_mkdb /etc/login.conf

Now.. Wie kann ich das jetzt querien?

    $ getcap -f /etc/login.conf -s openfiles-max staff
    2048

Oooolright.

    $ ulimit -a
    core file size          (blocks, -c) unlimited
    data seg size           (kbytes, -d) 1572864
    file size               (blocks, -f) unlimited
    max locked memory       (kbytes, -l) 331942
    max memory size         (kbytes, -m) 993656
    open files                      (-n) 2048
    pipe size            (512 bytes, -p) 1
    stack size              (kbytes, -s) 4096
    cpu time               (seconds, -t) unlimited
    max user processes              (-u) 256
    virtual memory          (kbytes, -v) 1576960