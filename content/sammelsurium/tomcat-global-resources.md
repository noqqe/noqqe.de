---
title: Tomcat Global Resources
date: 2014-03-19T13:58:21
tags: 
- Software
- Tomcat
---

## Globale Ressource definieren

``` xml
<GlobalNamingResources>

      <Resource name="jdbc/dbname" auth="Container" driverClassName="com.sybase.jdbc3.jdbc.SybDataSource"
              maxActive="20" maxIdle="10" maxWait="10000" type="javax.sql.DataSource"
              url="jdbc:sybase:Tds:10.20.30.40:5000/dbname" username="user" password="passwort"/>

      <Resource name="jdbc/dbname" auth="Container" driverClassName="com.mysql.jdbc.Driver"
              maxActive="20" maxIdle="10" maxWait="10000" type="javax.sql.DataSource"
              url="jdbc:mysql://localhost/dbname" username="user" password="passwort"/>

</GlobalNamingResources>
```

## Context

Context definieren und Ressource verlinken

Webapp liegt beispielsweise in /var/lib/tomcat7/webapps/appmobile-webapp.war

``` xml
<Host name="localhost"  appBase="webapps"
      unpackWARs="true" autoDeploy="true">

  <Valve className="org.apache.catalina.authenticator.SingleSignOn" />

  <Context path="/appmobile-webapp" docBase="appmobile-webapp.war" debug="0"
                 reloadable="true" crossContext="true">
   <ResourceLink
      name="jdbc/dbname"
      global="jdbc/dbname"
      type="javax.sql.DataSource"
    />

   <ResourceLink
      name="jdbc/dbname"
      global="jdbc/dbname"
      type="javax.sql.DataSource"
    />
  </Context>

</Host>
```