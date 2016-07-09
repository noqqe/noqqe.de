---
categories:
- administration
- bsd
- openbsd
- osbn
date: 2016-07-09T11:09:55+02:00
tags:
- ulimit
- login
- maxfiles
- openbsd
title: OpenBSD Login Classes
---

Mit Login Classes lassen sich allerhand Limitierungen für Benutzer und
Gruppen in OpenBSD regeln. Wer also bei seinem Webserver auf das OpenFiles
Limit läuft oder beim Compilen von Software XY einen OutOfMemory Fehler
bekommt sollte sich das mal anschauen. Hier lassen sich auch Authentication
Methods und ganz viel anderer Kram definieren. Die
[Manpage](http://man.openbsd.org/OpenBSD-current/man5/login.conf.5) würde
ich jedem mal empfehlen.

	$ doas vim /etc/login.conf
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

Aber welche Klasse hat mein User eigentlich? Dafür gibt es eine extra
Spalte in der `/etc/passwd`.

		$ userinfo noqqe
		login   noqqe
		passwd  *
		uid     1001
		groups  noqqe wheel
		change  NEVER
		class
		gecos   ,,,
		dir     /home/noqqe
		shell   /usr/local/bin/bash
		expire  NEVER

Aha. Keine also. Wie assigne ich meinem User so eine Grouppe? Mit `chsh`.
Das geht bestimmt auch noch anders, aber in dem meinsten Fällen arbeite ich
mit dem Tool.

    $ chsh noqqe
    ...
    class: staff
    ...

Gut. Nachdem mir die Werte immernoch nicht gefallen, muss ich die
login.conf doch noch anpassen. Einfach die Gruppe an sich editieren und das
DB File neu bauen.

		$ vim /etc/login.conf
		# build new db
		$ sudo cap_mkdb /etc/login.conf

Erstmal stehts da jetzt. Ob das auch wirklich geklappt hat (zumindest die
Formatierung des Files) kann man überprüfen in dem man ein bestimmtes
Attribut einer Klasse überprüft.

		$ getcap -f /etc/login.conf -s openfiles-max staff
		4096

Alles klar, das hat geklappt. Das unangenehme bei der Sache ist, dass das
neue Limit erst beim nächsten Login greift. Also neue `ksh` starten reicht
nicht. Es muss ein neues TTY her. Bei Daemons wusste ich mir bisher nicht
anders zu helfen, als kurz zu rebooten. Irgendwelche Ideen?

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

`ulimit` ist hier (wie auch bei Linux) das mittel der Wahl um die neu
gesetzten Limits zu überprüfen.

Benutzereinschränkungen über die `login.conf` wirken irgendwie stabiler und
sicherer als `ulimit` Statements in irgendwelche `/etc/profile` wie man es
bei Linux so kennt. Auch die 1000 anderen Features die das System bietet
muss ich mir unbedingt noch anschauen.
