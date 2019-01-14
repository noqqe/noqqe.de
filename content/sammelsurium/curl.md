---
title: curl
date: 2013-11-05T20:44:06
tags: 
- Software
- curl
---

## Post JSON Data with CURL

    curl -v -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"author":"ADFASDF","email":"ASDFADSF","text":"SAFDASDF","parent":null}' http://localhost:6601/new?uri=%2Fblog%2F2013%2F09%2F08%2Fuseful-use-of-cat%2F

## Debug gzipped Connections

~~~
curl --compressed -s -D - http://cdn.example.com//lib/dashboardAPI/v2.13.42/library.min.js?t=1382096123 -o /dev/null
HTTP/1.1 200 OK
Date: Thu, 07 Nov 2013 12:05:08 GMT
Server: Apache
Last-Modified: Fri, 18 Oct 2013 11:35:23 GMT
ETag: "fba9b2-189b3d-4e9025474f90f"
Accept-Ranges: bytes
Content-Length: 1612605
Vary: User-Agent
Content-Type: text/x-js
~~~

normalerweise sollte das dann so aussehen

~~~
curl --compressed -s -D - http://cdn.example.com//lib/dashboardAPI/ -o /dev/null
HTTP/1.1 403 Forbidden
Date: Thu, 07 Nov 2013 12:14:47 GMT
Server: Apache
Accept-Ranges: bytes
Vary: Accept-Encoding,User-Agent
Content-Encoding: gzip
Content-Length: 1030
Content-Type: text/html
~~~