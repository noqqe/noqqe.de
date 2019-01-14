---
title: Tomcat Security
date: 2013-05-27T11:46:08
tags:
- Software
- Tomcat
---

~~~
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
~~~

Überprüfen mit:

~~~
sslscan frontend01.example.com | grep Accepted

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
~~~

### DENY HTTP Methods

~~~
<?xml version="1.0" encoding="ISO-8859-1"?>

<web-app xmlns="http://java.sun.com/xml/ns/javaee"
  xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
  xsi:schemaLocation="http://java.sun.com/xml/ns/javaee
                      http://java.sun.com/xml/ns/javaee/web-app_3_0.xsd"
  version="3.0"
  metadata-complete="true">

  <display-name>Welcome to Tomcat</display-name>
  <description>
     Welcome to Tomcat
  </description>

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
~~~

Überprüfung mit curl

~~~
for x in PATCH CONNECT GET TRACE POST PUT OPTIONS HEAD ; do echo -ne "$x:    \t" ; curl -v --insecure -X $x https://frontend01.example.com/ 2>&1 | grep '< HTTP' ; done
PATCH:          < HTTP/1.1 403 Forbidden
CONNECT:        < HTTP/1.1 403 Forbidden
GET:            < HTTP/1.1 200 OK
TRACE:          < HTTP/1.1 405 Method Not Allowed
POST:           < HTTP/1.1 200 OK
PUT:            < HTTP/1.1 200 OK
OPTIONS:        < HTTP/1.1 200 OK
~~~

Überprüfung mit w3af

~~~
$ w3af -s /root/w3af/w3af-full.script-a

http-settings
set userAgent "Mozilla/5.0 (X11; U; Linux i686; it; rv:1.9.2.13) Gecko/20101206 Ubuntu/10.10 (maverick) Firefox/3.6.13"
set timeout 5
back
plugins
output console,textFile,htmlFile
output config htmlFile
set verbosity 10
set fileName /root/w3af/report.html
back
output config textFile
set verbosity 10
set fileName /root/w3af/report.txt
output
back
output config console
set verbosity 5
back
discovery allowedMethods
back
target
set target https://frontend01.example.com
back
start
~~~

HTTP Code logging output:

~~~
## grep HTTP /root/w3af/report.txt | grep -v "200"
[ Mon May 27 14:45:41 2013 - debug ] OPTIONS https://frontend01.example.com/ returned HTTP code "403" - id: 2
[ Mon May 27 14:45:41 2013 - debug ] PATCH https://frontend01.example.com/ returned HTTP code "403" - id: 13
[ Mon May 27 14:45:41 2013 - debug ] OPTIONS https://frontend01.example.com/ returned HTTP code "403" - id: 16
[ Mon May 27 14:45:41 2013 - debug ] CONNECT https://frontend01.example.com/ returned HTTP code "403" - id: 24
[ Mon May 27 14:45:41 2013 - debug ] TRACE https://frontend01.example.com/ returned HTTP code "405" - id: 41
~~~

### Turn off Directory Listings

in WEB-INF/web.xml

~~~
<servlet>
  <servlet-name>default</servlet-name>
  <servlet-class>org.apache.catalina.servlets.DefaultServlet</servlet-class>
  <init-param>
    <param-name>debug</param-name>
    <param-value>0</param-value>
  </init-param>
  <init-param>
    <param-name>listings</param-name>
    <param-value>false</param-value>  <!-- make sure this is false -->
  </init-param>
  <load-on-startup>1</load-on-startup>
</servlet>
~~~

### Set Port 8005 SHUTDOWN

~~~
 <Server port="8005" shutdown="ReallyComplexWord">

or

 <Server port="-1" shutdown="ReallyComplexWord">
~~~

aber OBACHT!

### Crypted Passwords in tomcat-users.xml

The file ${tomcat_home}/conf/tomcat-users.xml stores user names and
passwords. By default the passwords are in clear text, e.g.:

~~~
<?xml version='1.0' encoding='utf-8'?>
<tomcat-users>
  <role rolename="tdsConfig"/>
  <role rolename="manager"/>
  <role rolename="admin"/>
  <user username="sysadmin" password="yrPassword" roles="manager,admin"/>
  <user username="cataloger" password="myPassword" roles="tdsConfig"/>
</tomcat-users>
~~~

Store them instead in digested form. First generate the digest (do this for each user):

~~~
  > cd  ${tomcat_home}/bin
  > ./digest.sh -a SHA yrPassword
  > yrPassword:ff01ea2afaae56c2b7da5e25ec18c505e58f12d7
~~~

If you are using DIGEST authentication, (only needed if you arent using
SSL), then use {username}:{realm}:{yrPassword} instead of {yrPassword} in
calculating the digest, for example sysadmin:THREDDS Data
Server:yrPassword. See this for more details.

Then cut and paste the digested passwords into the tomcat-users.xml:

~~~
<?xml version='1.0' encoding='utf-8'?>
<tomcat-users>
  <role rolename="tdsConfig"/>
  <role rolename="manager"/>
  <role rolename="admin"/>
  <user username="username1" password="ff01ea2afaae56c2b7da5e25ec18c505e58f12d7" roles="manager,admin"/>
  <user username="username2" password="5413ee24723bba2c5a6ba2d0196c78b3ee4628d1" roles="tdsConfig"/>
</tomcat-users>
~~~

Then change the server.xml file to tell it to use digested passwords, by
adding this `<Realm>` element to the `<Host>` element named localhost:

~~~
<Host name="localhost" debug="0" appBase="/opt/tomcat/webapps" unpackWARs="true" autoDeploy="true"
     xmlValidation="false" xmlNamespaceAware="false">
  <Realm className="org.apache.catalina.realm.MemoryRealm" digest="SHA" />
  ...
</Host>
~~~
