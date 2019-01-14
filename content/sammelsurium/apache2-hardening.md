---
title: Apache2 Hardening
date: 2012-01-17T12:46:54
tags:
- Software
- Apache
---

## Informationen sinnvoll unterdrücken

Um einen Exploit für das Ziel zu finden ist schonmal die genaue Version der
Software wichtig. Um hier gleich von Anfang an so wenig wie möglich
preiszugeben empfiehlt sich die Apache Konfiguration:

    ServerTokens ProductOnly

Beispiele:

> HTTP/1.1 400 Bad Request
> Date: Thu, 12 Jan 2012 08:03:31 GMT
> Server: Apache
> ## Anstatt:
> HTTP/1.1 404 Not Found
> Date: Thu, 12 Jan 2012 07:53:05 GMT
> Server: Apache/2.2.16 (Debian) PHP/5.3.3-7+squeeze3 with Suhosin-Patch mod_ssl/2.2.16 OpenSSL/0.9.8o

Um im Browser bei der Fehlerseite zb. bei 404 keine "Signatur" anzeigen zu lassen:

    ServerSignature Off

## TRACE deaktivieren

    TraceEnable Off

Durch Traces kann man den Request inkl. Header zurückgeben. Dadurch sind XSS-Attacken möglich.

## PHP Informationen unterdrücken

Damit im Header auch keine PHP Informationen mehr auftauchen:

     PHP/5.3.3-7+squeeze3

in der php.ini setzen:

    expose_php = Off

## Directory Indexing

```
 Alias /doc/ "/usr/share/doc/"
 <Directory "/usr/share/doc/">
     Options '''Indexes''' MultiViews FollowSymLinks
     AllowOverride None
     Order deny,allow
     Deny from all
     Allow from 127.0.0.0/255.0.0.0 ::1/128
 </Directory>
```

Sollte da umbedingt rausgenommen werden. Generell ist zu überlegen ob der
/doc/ Alias überhaupt gebraucht werden kann.

## Apache Default Readme File

Diese Files will man evtl. löschen da sie 1. total unnütz sind und 2. auf
die Version des Apache2 schliessen lassen.

    /usr/share/apache2/icons/README
    /usr/share/apache2/icons/README.html

## ETag Headers

In Apache2

    Header unset ETag
    FileETag None

dazu will man auch mal den Blogpost
[Disable ETAG](http://www.lavluda.com/2008/10/20/website-optimization-01-disable-etag-in-apache-debianubuntu/)
lesen.

## Example Config

```
1. Apache Config:
FileETag MTime Size
=> ETag Information Leak

2. Das Apache Modul header enablen.

3. Clickjacking:
Header always append X-Frame-Options SAMEORIGIN
4. Prevent caching:
Header set Cache-Control "max-age=0, no-cache, no-store, must-revalidate"
Header set Pragma "no-cache"
Header set Expires "Wed, 11 Jan 1984 05:00:00 GMT"
Information Leakage:

5. Apache Server:
ServerTokens ProductOnly
ServerSignature Off

6. php.ini
expose_php = off
```