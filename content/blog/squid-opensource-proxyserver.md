---
date: 2008-08-30T20:02:26+02:00
slug: squid-opensource-proxyserver
comments: true
title: Squid - OpenSource ProxyServer
aliases:
- /archives/207
- /blog/2008/08/30/squid-opensource-proxyserver/
tags:
- Hardware
- Linux
- EEE-PC
- Linux
- opensource
- proxy
- Squid
- ubuntu
- ubuntueee
---

Erste Übung in der Technik-Abteilung in der ich seit ca 1er Woche bin.

> Setz mal nen ProxyServer auf und Connecte mit deinem EeePc

also: **EeePC -> Arbeitsrechner (als Proxy missbraucht) -> Internet **
Um jetzt nicht 4500 Zeilen in den Blog zu posten hier kurz das Wichtigste:
Erstmal Minimum Config. Teils schon vorgearbeitet:

```
#Recommended minimum configuration:
http_access allow manager localhost
http_access deny manager
http_access allow purge localhost
http_access deny purge
http_access deny !Safe_ports
http_access deny CONNECT !SSL_ports
http_access allow all
http_access allow localhost
http_access deny all
http_access deny microsoft #Individuell eingestellt.
```

Standard-Proxy-Port einstellen der dann im Browser angegeben werden muss:

```
# Squid normally listens to port 3128
http_port 3128
```

AccessListen definieren:

```
#Recommended minimum configuration:
acl all src 0.0.0.0/0.0.0.0
acl manager proto cache_object
acl localhost src 127.0.0.1/255.255.255.255
acl to_localhost dst 127.0.0.0/8
acl microsoft dstdomain .microsoft.com # Verbotssite die oben eingebaut wird, hab ich hier definiert.
acl SSL_ports port 443		# https
acl SSL_ports port 563		# snews
acl SSL_ports port 873		# rsync
acl Safe_ports port 80		# http
acl Safe_ports port 21		# ftp
acl Safe_ports port 443		# https
acl Safe_ports port 70		# gopher
acl Safe_ports port 210		# wais
acl Safe_ports port 1025-65535	# unregistered ports
acl Safe_ports port 280		# http-mgmt
acl Safe_ports port 488		# gss-http
acl Safe_ports port 591		# filemaker
acl Safe_ports port 777		# multiling http
acl Safe_ports port 631		# cups
acl Safe_ports port 873		# rsync
acl Safe_ports port 901		# SWAT
acl purge method PURGE
acl CONNECT method CONNECT
```

So und jetzt mal eben sehen wo ich so war mit EeePc.

```
# Access-Log
access_log /var/log/squid/access.log squid
```

