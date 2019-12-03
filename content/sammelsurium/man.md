---
title: man
date: 2019-08-02T12:23:34
tags:
- Software
- man
---

Dieser Eintrag ist tatsächlich eine Art Mixtur aus `man`, `roff` und `mandoc`
von OpenBSD

Wie kann ich mir ein roff Dokument anzeigen?

```
man /tmp/motd.5
```

Easy as that. Okay.

Ich will aber noch ein bisschen mehr...

## Spickzettel

`.Sh` - Section
`.Xr` - verweis auf andere Sektion

Ein normales `roff` Dokument sieht in etwa so aus. Genau gesagt liegt hier
ein `mdoc` Format vor.

```
.\"	$OpenBSD: motd.5,v 1.10 2009/05/06 19:43:16 jmc Exp $
.\"	$NetBSD: motd.5,v 1.2 1994/12/28 18:58:53 glass Exp $
.\"
.\" This file is in the public domain.
.\"
.Dd $Mdocdate: May 6 2009 $
.Dt MOTD 5
.Os
.Sh NAME
.Nm motd
.Nd message of the day
.Sh DESCRIPTION
The
.Pa /etc/motd
file is normally displayed by
.Xr login 1
after a user has logged in but before the shell is run.
It is generally used for important system-wide announcements.
During system startup, a line containing the kernel version string
replaces all lines up to (but not including) the first blank line of
this file.
Customised messages can be added to this file after
that first blank line.
.Pp
Individual users may suppress the display of this file by creating a file named
.Pa .hushlogin
in their home directories.
.Sh FILES
.Bl -tag -width /etc/motd -compact
.It Pa /etc/motd
.El
.Sh EXAMPLES
.Bd -literal
OpenBSD 3.0 (GENERIC) #5: Sat Jan 26 20:13:07 MST 2002

Make sure you have a .forward file...

4/17	Machine will be down for backups all day Saturday.
.Ed
.Sh SEE ALSO
.Xr login 1
```

