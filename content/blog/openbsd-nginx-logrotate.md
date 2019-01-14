---
comments:
- author: Martin
  content: 'FreeBSD hat auch weiterhin newsyslog: https://github.com/freebsd/freebsd/tree/master/usr.sbin/newsyslog

    Probier doch mal ob es unter OpenBSD kompiliert, wahrscheinlich erfordert es nur
    kleine Anpassungen der includes.

'
  date: '2014-03-02T16:46:15.344601'
- author: noqqe
  content: "Oh, hab auch bei FreeBSD in usr.bin/ gesucht. Thx!\n\nBei compilen bin
    ich mir nicht ganz sicher, das m\xFCsst ich quasi bei jedem Distupgrade erneut
    machen.. \xDCberleg ich mir mal"
  date: '2014-03-02T19:35:17.954852'
- author: waldner
  content: 'Das $(ls -1 ...) ist nicht notwendig, man kann einfach:

    for x in /var/www/logs/*.log; do ...

'
  date: '2014-03-06T22:20:32.983181'
- author: noqqe
  content: Jap, vollkommen richtig. Beim basteln entsteht sowas manchmal, man kennt
    das ja :)
  date: '2014-03-09T16:31:56.090328'
date: '2014-03-02T10:09:00'
tags:
- newsyslog
- openbsd
- log
- syslog
- nginx
- logrotate
- freebsd
- bsd
- rotation
- obsd
title: OpenBSD nginx Logrotate
---

Seit mittlerweile fast 3 Releases setze ich [OpenBSD](http://openbsd.org) auf
meinem Housing-Server ein. Ich mag das OS immernoch sehr gerne. Die WTF-Rate bei
Software ist hier denkbar niedrig. Ganz ohne gehts aber leider auch nicht.

{{< figure src="/uploads/2014/03/uhm.gif" >}}

Unter OpenBSD ist [newsyslog](http://www.weird.com/~woods/projects/newsyslog.html)
für die Logrotation verantwortlich. Das gut abgehangene Stück Software ist wie
immer toll dokumentiert und funktioniert. Nur leider ignoriert es Wildcards

```
# logfile_name      owner:group     mode count size when  flags
/var/cron/log       root:wheel  600  3     10   *     Z
/var/www/logs/*.log             644  7     *    *     Z "kill -s USR1 `cat /var/run/nginx.pid`"
/var/www/logs/error.log         644  7     *    24    Z "kill -s USR1 `cat /var/run/nginx.pid`"
```

Einen Testlauf mit dry-run lässt sich starten mit

```
$ newsyslog -nv
/var/cron/log <3Z>: size (KB): 2.28 [10] --> skipping
/var/www/logs/error.log <7Z>: age (hr): 1 [24] --> skipping
```

Unter [FreeBSD](http://www.freebsd.org/cgi/man.cgi?query=newsyslog.conf&sektion=5)
enthält die newsyslog Version das Flag `G`.

> G   indicates that the specified logfile_name is a shell pat-
>     tern, and that newsyslog(8) should archive all filenames
>     matching that pattern using the other options on this
>     line.  See glob(3) for details on syntax and matching
>     rules.

Unter oBSD ist dieser Modus leider nicht verfügbar. Was für Lösungen sind also möglich? Ich hab mich
vorerst dafür entschieden die Liste der Entries mit einem Einzeiler zu
generieren.

``` bash
$ for x in $(ls -1 /var/www/logs/*.log) ; do echo -e "$x\t\t" '644  7     *    24    Z' ; done
```

und beim letzten Eintrag den entsprechenden nginx reload Command anfügen.

``` bash
# logfile_name           owner:group mode count size when  flags
/var/cron/log            root:wheel  600  3     10   *     Z
/var/www/logs/vhost1_access.log      644  7     *    24    Z
/var/www/logs/vhost2_access.log      644  7     *    24    Z
/var/www/logs/vhost3_access.log      644  7     *    24    Z
/var/www/logs/vhost4_access.log      644  7     *    24    Z
/var/www/logs/vhost5_access.log      644  7     *    24    Z
/var/www/logs/vhost6_access.log      644  7     *    24    Z "kill -s USR1 `cat /var/run/nginx.pid`"
```

Ich kann mir nicht vorstellen, dass ich der Einzige mit diesem Problem bin,
daher bin ich auf alternative Vorschläge gespannt. Selber ein Skript schreiben
ist mir zu fehleranfällig und aufwändig, [logrotate](https://fedorahosted.org/logrotate/) ist
nicht paketiert. FreeBSD scheint newsyslog auch durch rsyslog ersetzt zu haben.
