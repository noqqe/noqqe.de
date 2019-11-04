---
title: Debug a File with Tabs
date: 2014-01-21T14:40:59
tags:
- HowTos
---

Manchmal kann man auf Excel oder in Windows / Unix gemischten Dingen mit
Tabs.

```
$ file file.tab
file.tab:      ASCII text, with CRLF line terminators
```

```
$ cat file.tab
Friday  145720  user@aol.com       Chuck    Honk    PartnerLead     Returning       Male
Friday  153740  user@yahoo.co.uk     Horst    Tappert   User    Returning       Male
```

selbes mit -A (Tabs als ^I und ^M ist newline mit $ als limiter)

```
$ cat -A file.tab
Friday^I65584^user@hotmail.com^IChuck^IHonk^IUser^IReturning^IMale
Friday^I145720^user@aol.com^IHorst^ITappert^IUser^IReturning^IMale
```

Man kann dieses File auch mit Hexdump anschauen. die \\t sind hierbei die
Tabs. Logischerweise

```
$ hexdump -c head.txt
0000000   I   d  \t   A   c   t   i   o   n   D   a   t   e  \t   C   u
0000010   s   t   o   m   e   r   I   D  \t   E   m   a   i   l   A   d
0000020   d   r   e   s   s  \t   F   i   r   s   t   N   a   m   e  \t
0000030   S   u   r   N   a   m   e  \t   c   u   s   t   o   m   e   r
0000040   t   y   p   e  \t   T   y   p   e  \t   G   e   n   d   e   r
0000050  \t   O   c   c   S   t   a   t   u   s  \t   F   T   B   u   s
0000060   T   y   p   e  \t   F   T   O   c   c  \t   A   g   e   O   f
0000070   C   u   s   t   o   m   e   r  \t   P   r   o   d   u   c   t
```

Bei normalen Tabs in Vi entstehen (zumindest in meiner Konfiguration)
garkeine Tabs sonder spaces.

```
$ cat -A foo
1   2   8$
3   4   4$
3   4   4$
3   4   4$
3   4   4$
```

Wenn man das abschalten will kann man in Vi mit `CTRL`+`v`+`Tab` einen
solchen echten Tab einfügen

```
$ cat -A foo
1   2   8$
3   4   4$
3   4   4$
3   4   4$
^I3   4   4$
```

Das gleiche mit Vi erzeugt

```
$ hexdump -c foo
0000000  \t   1  \t   2  \t   3  \n  \t   1  \t   2  \t   3  \n  \t   1
0000010  \t   2  \t   3  \n  \t   1  \t   2  \t   3  \n
000001c
```

## Links

[Linux Non Printing Characters](http://www.if-not-true-then-false.com/2011/linux-display-show-file-contents-tabs-line-breaks-non-printing-characters/)
