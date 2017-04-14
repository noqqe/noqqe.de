---
date: 2017-04-14T11:34:46+01:00
tags:
- python
- date
- touch
title: Creation/Modification Dates
---

Solang du ein halbwegs normales Filesystem hat, speichert es dir
Access, Modification und Creation Dates als Meta Information ab. Das macht sogar
`HFS+` von Apple.

``` bash
$ stat foo
16777220 47263467 -rw-r--r-- 1 noqqe staff 0 16 "Mar  7 10:41:34 2017" "Mar  7 10:13:46 2017" "Mar  7 10:13:46 2017" "May  5 01:00:00 1970" 4096 8 0 foo
```

Man braucht kurz um die Infos geistig zu sortieren. Die Manpage von `stat` ist
garnicht mal so aussagekräftig. Zumindest unter macOS.

```
HISTORY

  The stat utility appeared in NetBSD 1.6 and FreeBSD 4.10.

BSD  May 8, 2003
```

Gut hat sich natürlich auch seit knapp 14 Jahren kein Arsch drum gekümmert.
Entweder ich bin von OpenBSD verwöhnt oder das ist wirklich so traurig.
Wenigstens kann die `ssh` Version mittlerweile `ed25519`.

Egal. `stat` ist jedenfalls eigentlich ganz schön cool! So gibt es mit `-s`
Shell kompatiblen Output

```
$ eval $(stat -s foo)
$ echo $st_mtime
1488878026
$ echo $st_ctime
1488878026
```

Außerdem kann ich das Datumsformat gleich mit spezifizieren.

```
$ stat -t %Y%m%d foo
16777220 47263467 -rw-r--r-- 1 noqqe staff 0 16 "20170307" "20170307" "20170307" "19700505" 4096 8 0 foo
```

Schön und gut. Kann ich diese Daten eigentlich auch setzen? Normalerweise machen das die Userland
Tools ja automatisch. Einige Tools können die Daten des Files aber angeben.

Das Modified und Access Date lässt sich mittels `touch` leicht selbst
definieren.

```
$ stat foo ; touch -macft 200101011400.00 foo ; stat foo
16777220 47410577 -rw-r--r-- 1 noqqe staff 0 0 "Mar 14 14:00:00 2017" "Mar 14 14:00:00 2017" "Mar 14 17:28:40 2017" "Mar  8 09:23:12 2017" 4096 0 0 foo
16777220 47410577 -rw-r--r-- 1 noqqe staff 0 0 "Jan  1 14:00:00 2001" "Jan  1 14:00:00 2001" "Mar 14 17:29:14 2017" "Jan  1 14:00:00 2001" 4096 0 0 foo
```

Oder wenn man es (weswegen ich den Post hier überhaupt schreibe)
programmatisch für `python` braucht:

``` python
import os, datetime

path = "/tmp/foo"
acc = datetime.datetime.utcnow().strftime("%s")
mod = datetime.datetime.utcnow().strftime("%s")

os.utime(path, (int(acc), int(mod)))
```

Das Creation Date ist so eine Sache. Eigentlich bedeutet `ctime` garnicht
creation time. Die
[sys/stat.h](http://pubs.opengroup.org/onlinepubs/9699919799/basedefs/sys_stat.h.html)
sagt hierzu:

> struct timespec st_atim Last data access timestamp.
> struct timespec st_mtim Last data modification timestamp.
> struct timespec st_ctim Last file status change timestamp.

Also zeigt der Timestamp lediglich wann sich das letzte mal die Metadaten
des Files geändert haben. `touch` bietet hier keinen Weg den Wert auf einen
speziellen Wert zu setzen. Aber man kann die `ctime` von einem anderen File
"mitkopieren".

```
$ touch -r foo foo2
$ stat foo*
10 805667 -rw-r--r-- 1 noqqe noqqe 3217063 6 "Jun 15 14:00:00 1989" "Jun 15 14:00:00 1989" "Mar  8 08:53:41 2017" 16384 4 0 foo
10 805666 -rw-r--r-- 1 noqqe noqqe 0 0 "Jun 15 14:00:00 1989" "Jun 15 14:00:00 1989" "Mar  8 09:21:50 2017" 16384 0 0 foo2
```

`-r` benutzt sozusagen ein anderes File als Vorlage.

Aufgefallen ist mir auch, dass unter OSX ein 4. Datetime Feld in der `stat`
Ausgabe ist. Das `birth` Feld. Dieses unterscheidet sich vom Creation
Field dadurch, dass es festhält wann der `inode` erstellt wurde.

Nachdem ich mich jetzt durch `stat` und File Flags von macOS, OpenBSD und
teilweise Linux gewurschtelt hab muss ich sagen das das alles ein
ziemlicher Wust ist bei dem jeder bisschen was anderes treibt. Je nach FS
natürlich.
