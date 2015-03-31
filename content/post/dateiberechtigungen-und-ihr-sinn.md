---
date: 2008-05-23T11:55:05+02:00
type: post
slug: dateiberechtigungen-und-ihr-sinn
status: publish
comments: true
title: Dateiberechtigungen und ihr Sinn
aliases:
- /archives/63
categories:
- Hardware
- Linux
tags:
- shell
---

Neulich Abend hab ich mal wieder bisschen mit der Shell rumgespielt und mich gefragt was wohl im Detail die Ausgabe von

```
ls -l
```

bedeutet. Ausgabe lautet wie folgt:

```
-rwx------ 1 npx  npx  514572 2008-05-13 12:53 77551-home_sweet_home.png
-rw-r--r-- 1 root root    425 2008-05-20 10:41 aptcache
drwxr-xr-x 2 npx  npx    4096 2008-05-19 12:55 Bilder
drwx------ 2 npx  npx    4096 2008-05-19 13:15 Desktop
drwxr-xr-x 2 npx  npx    4096 2008-05-19 12:55 Dokumente
lrwxrwxrwx 1 npx  npx      26 2008-05-10 02:16 Examples -> /usr/share/example-content
drwxr-xr-x 2 npx  npx    4096 2008-05-19 12:55 Musik
drwxr-xr-x 2 npx  npx    4096 2008-05-19 12:55 Öffentlich
-r--r--r-- 1 npx  npx      87 2008-05-21 16:21 text1
-rwxr-xr-x 1 npx  npx      99 2008-05-15 22:06 path1
drwxr-xr-x 2 npx  npx    4096 2008-05-20 14:38 shell
drwxr-xr-x 2 npx  npx    4096 2008-05-19 12:55 Videos
drwxr-xr-x 2 npx  npx    4096 2008-05-19 12:55 Vorlagen
-rw-r--r-- 1 root root     84 2008-05-21 16:19 test
```

Springender Punkt bei der Sache sind erstes. Also ` drwxr-xr-x `
oder ähnliches. Mir war ja klar das es was mit Dateirechten zu tun hat.
Aber eigentlich gibt es ja nur 3 gänige Rechte!

  * r - read
  * w - write
  * x - execute

Wofür also die ganzen anderen Buchstaben und Striche?
War natürlich ein super Fall für mein "Linux Befehle Buch" :)

Demnach sind die ersten 3 Buchstaben zur Deklaration der Rechte des Eigentümers vorgesehen,
meistens volle Rechte:

```
rwx (lesen schreiben und ausführen)
7
```

Die zweiten 3 Bits sind für die Rechte der Gruppe:

```
r-x (nur lesen und ausführen)
5
```


und die letzten 3 stehen für andere Benutzer:

```
r-- (nur lesen)
4
```


Man hat sicher bemerkt das ich so lustige Zahlen unter die einzelnen Rechte geschrieben habe.
Auch das hat einen Grund. Und zwar kann man durch ein simples Oktalsystem die Rechte präzise mit einer Zahl benennen.

```
r = 4
w = 2
x = 1

```

Wenn man diese Werte beliebig miteinander kombiniert bekommt man immer nur einen
Wert pro kombination. Es werden also die Zahlen der einzelnen vergebenen Rechte
zusammengezählt und zum Wert einer Gruppe gemacht. Diese Werte setzt mann dann von
allen 3 Gruppen hintereinander und es ergibt sich wie unten eine Zahlen Kombination.
Da hat sich jemand mal richtig etwas dabei gedacht ;)

```
---rwxrw-r--
```

hätte jetzt genau den Wert 0754 und jeder weiß welche Rechte auf diese Datei wirken ;) Klasse Sache und ich hab wieder was gelernt.
