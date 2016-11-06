---
title: "Securing Tomcat 7"
date: 2013-05-27T19:52:00+02:00
comments: true
categories:
- Web
- Linux
- Administration
tags:
- Tomcat
- Tomcat7
- Java
- HTTP
- Methods
- Secure
- Security
- ubuntuusers
- osbn
---

Zum Thema Application-Server absichern gibts unzählige Posts mit den
obskursten Vorschlägen. Hier meiner.

{{< figure src="/uploads/2013/05/beer.gif" >}}

## HTTP Methods

Die meisten Entwickler machen leider nur wenig Gebrauch von der vollen Breite
an HTTP Methoden. Das merkt man bei REST-APIs, aber auch anderswo. Überflüssige
Methoden können auch bereits vor der Applikation abgefangen werden.
Vorher aber mit Software/Kunden eurer Wahl abklären.

Per Default ist bei Tomcat nur `TRACE` deaktiviert.
In der `webapps/$APP/WEB-INF/web.xml` definieren:

``` xml
<?xml version="1.0" encoding="ISO-8859-1"?>
<web-app>
  [...]
  <security-constraint>
     <web-resource-collection>
          <web-resource-name>Forbidden</web-resource-name>
          <url-pattern>/*</url-pattern>
          <http-method>TRACE</http-method>
          <http-method>CONNECT</http-method>
          <http-method>PATCH</http-method>
          <http-method>OPTIONS</http-method>
     </web-resource-collection>
     <auth-constraint/>
  </security-constraint>
</web-app>
```

Ergebnis verfizieren:

``` bash
$ for x in PATCH CONNECT GET TRACE POST PUT OPTIONS HEAD ; do
>   echo -ne "$x:    \t"
>   curl -v -X $x http://localhost:8080/ 2>&1 | grep '< HTTP'
> done

PATCH:          < HTTP/1.1 403 Forbidden
CONNECT:        < HTTP/1.1 403 Forbidden
GET:            < HTTP/1.1 200 OK
TRACE:          < HTTP/1.1 405 Method Not Allowed
POST:           < HTTP/1.1 200 OK
PUT:            < HTTP/1.1 200 OK
OPTIONS:        < HTTP/1.1 403 Forbidden
HEAD:           < HTTP/1.1 200 OK
```

Ich kann übrigens jedem mal empfehlen sich [Know your HTTP](https://github.com/bigcompany/know-your-http)
anzuschauen.

## Sichere Ciphers

Gerade bei Auditoren ein beliebtes Thema. Man kennt das von Apache httpd mod_ssl Options. Den 40bit Schlüssel von
nebenan will man heute nicht mehr umbedingt Serverseitig offerieren.
Das Pendant zum Apache Tomcat, indem der Parameter `ciphers` im HTTP Connector her muss.

``` xml
<Connector port="8443"
  executor="Catalina-Threads"
  protocol="org.apache.coyote.http11.Http11Protocol"
  scheme="https"
  secure="true"
  sslProtocol="TLS"
  SSLEnabled="true"
  SSLCertificateFile="${catalina.base}/conf/host.crt"
  SSLCertificateKeyFile="${catalina.base}/conf/host.key"
  ciphers="SSL_RSA_WITH_RC4_128_MD5, SSL_RSA_WITH_RC4_128_SHA, TLS_RSA_WITH_AES_128_CBC_SHA, TLS_DHE_RSA_WITH_AES_128_CBC_SHA, TLS_DHE_DSS_WITH_AES_128_CBC_SHA, SSL_RSA_WITH_3DES_EDE_CBC_SHA, SSL_DHE_RSA_WITH_3DES_EDE_CBC_SHA, SSL_DHE_DSS_WITH_3DES_EDE_CBC_SHA"
/>
```

Gut überprüfen lässt sich das übrigens mit dem Tool `sslscan`

```
$ sslscan localhost:8443 | grep Accepted
    Accepted  SSLv3  128 bits  DHE-RSA-AES128-SHA
    Accepted  SSLv3  128 bits  AES128-SHA
    Accepted  SSLv3  168 bits  EDH-RSA-DES-CBC3-SHA
    Accepted  SSLv3  168 bits  DES-CBC3-SHA
    Accepted  SSLv3  128 bits  RC4-SHA
    Accepted  SSLv3  128 bits  RC4-MD5
    Accepted  TLSv1  128 bits  DHE-RSA-AES128-SHA
    Accepted  TLSv1  128 bits  AES128-SHA
    Accepted  TLSv1  168 bits  EDH-RSA-DES-CBC3-SHA
    Accepted  TLSv1  168 bits  DES-CBC3-SHA
    Accepted  TLSv1  128 bits  RC4-SHA
    Accepted  TLSv1  128 bits  RC4-MD5
```

Weitere Punkte sind zum Beispiel `Directory Indexing` und Plain
Passwords in `tomcat-users.xml` abschalten. Aber natürlich auch noch viel
Enterprise Vodoo den man im Internet findet.
