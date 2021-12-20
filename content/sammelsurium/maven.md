---
title: "Maven"
date: 2021-12-20T10:57:19+01:00
tags:
- Software
- maven
- mvn
- java
---

Maven.

<!--more-->

> Maven's primary goal is to allow a developer to comprehend the complete
> state of a development effort in the shortest period of time. In order to
> attain this goal, Maven deals with several areas of concern:
>
> * Making the build process easy
> * Providing a uniform build system
> * Providing quality project information
> * Encouraging better development practices

## Commands

Upload eines jar in ein Maven Repo mittels Kommandozeile

```
 mvn -B deploy:deploy-file \
  -DgroupId=com.acme.product \
  -DartifactId=viewer \
  -Dversion=1.31.12 \
  -Dpackaging=jar \
  -Dfile=viewer.jar \
  -DgeneratePom=true \
  -DrepositoryId=acme-nexus \
  -Durl=https://nexus.acme.com/repository/acme-maven-release/
```

Upload eines war in ein Maven Repo mittels Kommandozeile

```
 mvn -B deploy:deploy-file \
  -DgroupId=com.acme.product \
  -DartifactId=viewer \
  -Dversion=1.31.12 \
  -Dpackaging=war \
  -Dfile=viewer.war \
  -DgeneratePom=true \
  -DrepositoryId=acme-nexus \
  -Durl=https://nexus.acme.com/repository/acme-maven-release/
```

## Settings

in `.m2/settings.xml`

Mirror

```xml
<settings>
  <mirrors>
    <mirror>
      <id>acme-nexus</id>
      <url>https://nexus.acme.com/repository/acme-maven-group/</url>
      <mirrorOf>central</mirrorOf>
    </mirror>
  </mirrors>
</settings>
```

Upload Repo (i.e. Nexus)

```xml
<settings>
  <servers>
    <server>
      <id>acme-nexus</id>
      <username>username</username>
      <password>xxx</password>
    </server>
  </servers>
</settings>
```

Proxy

```xml
<settings>
  <proxies>
    <proxy>
      <active>true</active>
      <protocol>http</protocol>
      <host>proxy.acme.com</host>
      <port>3128</port>
      <nonProxyHosts>127.0.0.1,10.40.88.56</nonProxyHosts>
    </proxy>
  </proxies>
</settings>
```
