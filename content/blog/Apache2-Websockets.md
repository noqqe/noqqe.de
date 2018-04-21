---
title: "Apache2 Websockets"
date: 2018-04-21T09:18:32+02:00
tags:
- Websockets
- Apache2
---

Immer diese Websockets.  Angenommen ich habe auf Port 8080 ein Stück Java
Software laufen, welches Websockets zwingend erfordert. So möchte ich
natürlich TLS Terminierung auf einem Reverse Proxy machen. In diesem Fall
Apache 2.4

```
<VirtualHost *:443>
  ServerName <hostname>

  ## Vhost docroot
  DocumentRoot "/var/www/redirect"

  ## Proxy rules
  ProxyRequests Off
  ProxyPreserveHost Off
  ProxyPass / ws://localhost:8080/ nocanon
  ProxyPassReverse / ws://localhost:8080/ nocanon
  ProxyPass / http://localhost:8080/ nocanon
  ProxyPassReverse / http://localhost:8080/ nocanon
</VirtualHost>
```

Wichtig ist dabei auch die Reihenfolge. `ws://` muss zwingend zuerst kommen,
sonst wird die Connection nicht weitergeleitet.

Außerdem geht das nicht mit Apache2.2 da dieser noch kein `mod_proxy_wstunnel`
hat, welches bei Apache2.4 natürlich vorher zu enablen gilt.
