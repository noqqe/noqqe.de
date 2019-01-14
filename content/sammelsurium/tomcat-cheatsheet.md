---
title: Tomcat Cheatsheet
date: 2013-01-06T13:24:55
tags:
- Software
- Tomcat
---

* ApplicationServer zum Ausführen von Java Code
* Kompletter HTTP-Server
* 1995 angefangen

### Versionen

* Tomcat6 (old!)
* Tomcat7 (normal)
* Tomcat8 (Bleeding State of the Art Edge)

### Für was wird Tomcat benutzt

* Enterprise Umfeld
* Applikationen wie CMS Systeme, Wikis, CRM

## Tomcat installieren

Debian & Ubuntu paketiert

    apt-get install tomcat7
    /etc/init.d/tomcat7 start

### Directories

* bin/ <- Skripte
* conf/ <- Konfiguration
* lib/ <- Bibliotheken
* logs/ <- ...
* temp/
* webapps/ <- "DocumentRoots"
  * ROOT <- Default "vHost"
  * JamWiki <- Beispiel Applikation Wiki
  * liferay <- Beispiel Applikation CMS
* work/ <- Laufzeit Files

## Tomcat richtig installieren.

Wer seinen Kunden mag und nicht so auf Schmerzen steht kann aber auch
andere Dinge tun.

Eine Installation dieser Art sei wärmstens empfohlen.

### Vorteile

* Ihr entscheidet wann es upgedated wird ("omg, apt-get hat unser
  Tomcat-Setup zerstört")
* Kein überflüssiger Unsinn im Tomcat (Docs, Manageranwendung, default
  Passwörter, minimale Installation)
* Konfiguration lesbar und verständlich (Stichwort Kommentare)
* Update von Tomcat testweise auf neue Version
* Update von Java testweise auf neue Version

### Schritte

#### Tomcat herunterladen

    cd /usr/local/
    wget http://tomcat.apache.org/download-70.cgi
    tar xfvz apache-tomcat-XX-VV.tar.gz
    ln -s apache-tomcat-XX-VV tomcat

#### JDK herunterladen

Auf

[Oracle Downloads](http://www.oracle.com/technetwork/java/javase/downloads/)

License Foo auswählen und herunterladen. Am besten den Downloadlink kopieren danach.

    cd /usr/local
    wget http://download.oracle.com/otn-pub/java/jdk/7u11-b21/jdk-7u11-linux-x64.tar.gz?AuthParam=1358962018_0709088e98e2b0851395247274a50376
    tar xfvz jdk-7u11-linux-x64.tar.gz
    ln -s jdk-7u11-linux-x64 java

#### 1. Datenstruktur aufbauen

    cd /data
    mkdir tomcat/node01
    mkdir bin conf webapps logs work temp
    cp /usr/local/tomcat/conf/web.xml conf/

#### 2. conf/server.xml erstellen

Normale (von Debian/Ubuntu mitgelieferte) server.xml sehr überladen.
Hässlicher Kommentar Stil. Muss aber nicht sein.

Voll funktionsfähige minimale Konfiguration:

    <?xml version='1.0' encoding='utf-8'?>
    <Server port="8005" shutdown="SHUTDOWN">
      <Listener className="org.apache.catalina.core.JasperListener" />
      <Listener className="org.apache.catalina.core.JreMemoryLeakPreventionListener" />

      <Service name="Catalina">
        <Connector port="8080" protocol="HTTP/1.1" />
        <Connector port="8009" protocol="AJP/1.3" />
        <Engine name="Catalina" defaultHost="localhost" jvmRoute="node01">
          <Host name="localhost"  appBase="webapps"
                unpackWARs="true" autoDeploy="true" >
          </Host>
        </Engine>
      </Service>
    </Server>

#### 3. startup.sh & shutdown.sh

bin/startup.sh

    #!/bin/bash
    export CATALINA_HOME=/usr/local/tomcat
    export CATALINA_BASE=/data/tomcat/node01
    #export JAVA_OPTS=""
    export JAVA_HOME="/usr/local/java"
    export CATALINA_OPTS="-Xmx128m"
    export CATALINA_PID="${CATALINA_BASE}/logs/tomcat.pid"
    ${CATALINA_HOME}/bin/catalina.sh start $@

bin/shutdown.sh

    #!/bin/bash
    export CATALINA_HOME=/usr/local/tomcat
    export CATALINA_BASE=/data/tomcat/node01
    #export JAVA_OPTS=""
    export JAVA_HOME="/usr/local/java"
    export CATALINA_PID="${CATALINA_BASE}/logs/tomcat.pid"
    ${CATALINA_HOME}/bin/catalina.sh stop $@

/etc/init.d/ Skripte können nach Bedarf noch gebaut werden.
Lässt sich auch schön als bestimmter User ausführen via `su user -c ...`

#### 4. Glücklich sein

    /data/tomcat/node01/bin/startup.sh

## Logging

Einer der unschönen Teile des Tomcat

## Connectoren

### AJP/1.3 Connectoren

    <Connector port="8009" protocol="AJP/1.3" />
    <Connector port="8009" executor="Catalina-Threads" protocol="org.apache.coyote.ajp.AjpProtocol" />
    <Connector port="8010" executor="Catalina-Threads" protocol="org.apache.coyote.ajp.AjpNioProtocol" />

### HTTP/1.1 Connectoren

    <Connector port="8080" executor="Catalina-Threads" protocol="HTTP/1.1" />
    <Connector port="8080" executor="Catalina-Threads" protocol="org.apache.coyote.http11.Http11Protocol" />
    <Connector port="8090" executor="Catalina-Threads" protocol="org.apache.coyote.http11.Http11NioProtocol" />
    <Connector port="8070" executor="Catalina-Threads" protocol="org.apache.coyote.http11.Http11AprProtocol" />

### HTTP/1.1 Connector parametisiert (SSL)

~~~
<Connector port="8443"
    executor="Catalina-Threads"
    protocol="org.apache.coyote.http11.Http11AprProtocol"
    scheme="https"
    secure="true"
    sslProtocol="TLS"
  SSLEnabled="true"
  SSLCertificateFile="${catalina.base}/conf/tomcatcert.pem"
  SSLCertificateKeyFile="${catalina.base}/conf/tomcatkey.pem"
  SSLPassword="tomcat"
     />
~~~

## MySQL Ressource

    <GlobaleNamingResources>
    ...
      <Resource name="jdbc/DB" auth="Container"
              type="javax.sql.DataSource"
              driverClassName="com.mysql.jdbc.Driver"
              url="jdbc:mysql://127.0.0.1/tomcat"
              username="webappuser"
              password="tomcat"
              maxActive="20"
              maxIdle="10"
              maxWait="15000"
        removeAbandoned="true"
              validationQuery="SELECT 1"
              testOnBorrow="true"
              testWhileIdle="true"
              timeBetweenEvictionRunsMillis="10000"
              minEvictableIdleTimeMillis="180000"
              />
    ...
    </GlobalNamingResources>

$CATALINA_BASE/webapps/$webapp/META-INF/context.xml

    <Context>
       <ResourceLink name="jdbc/DB" global="jdbc/DB" type="javax.sql.Datasource" />
    </Context>

## VHosts im Tomcat

## Manager Anwendung

Es gibt eine Manager Anwendung die aus in der Tomcat
Distribution mitgeliefert wird.

## Tomcat und Apache

### Wann brauch ich das und wann nicht?

Faustregel: Wenn es keinen Statischen Content gibt
der separat gelagert wird einfach sein lassen.

Apache2 sonst nur Overhead.
Abstraktion kostet immer Performance ;)

### Anbindungsvarianten

Mehrere Wege das zu tun:

* mod_jk
* mod_proxy
* mod_proxy_ajp

## mod_jk

Es heisst _nicht_ mod_justkidding! :)

## Tomcat Monitoring

Bezieht sich größtenteils auf die Java VM (JVM).
Wie gehts dem Teil und was macht sie gerade?

* jvisualvm / jconsole
* jstat
* JMX

```
-Dcom.sun.management.jmxremote -Dcom.sun.management.jmxremote.port=10013 -Dcom.sun.management.jmxremote.authenticate=false -Dcom.sun.management.jmxremote.ssl=false
```

* MBeans

## Replikation?

Tomcat Breitenskalierung ist nicht gleich Tomcat Replikation.

    <distributable/>

Tomcat Tribes
