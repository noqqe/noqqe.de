---
title: nginx
date: 2014-01-06T13:23:49
tags: 
- Software
- nginx
---

#### Security Configruation

Wahrscheinlich irgendwo im Internet gefunden

``` bash
ssl_prefer_server_ciphers on;
ssl_session_timeout  5m;
ssl_protocols TLSv1 TLSv1.1 TLSv1.2; ## not possible to do exclusive
ssl_ciphers 'EDH+CAMELLIA:EDH+aRSA:EECDH+aRSA+AESGCM:EECDH+aRSA+SHA384:EECDH+aRSA+SHA256:EECDH:+CAMELLIA256:+AES256:+CAMELLIA128:+AES128:+SSLv3:!aNULL:!eNULL:!LOW:!3DES:!MD5:!EXP:!PSK:!SRP:!DSS:!RC4:!SEED:!ECDSA:CAMELLIA256-SHA:AES256-SHA:CAMELLIA128-SHA:AES128-SHA';
add_header Strict-Transport-Security max-age=2592000;
ssl_ecdh_curve secp384r1;
```