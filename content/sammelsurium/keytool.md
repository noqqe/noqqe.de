---
title: "Java Keytool"
date: 2020-03-23T10:04:11+01:00
tags:
- Software
---

"Kannst du mal eben ein Tomcat Zertifikat austauschen?". Gänsehaut.

Oracles Werkzeug zur Zertifikatsverwaltung macht keinen Spass.

## Anzeigen

Anzeigen der Inhalte eines Keystores.

``` bash
keytool -list -keystore ssl/keystore.jks
keytool -list -v -keystore ssl/keystore.jks
```


## PEM zu JKS konvertieren

Ein herkömmliches `PEM` File kann zu einem `JKS` konvertiert werden.

Das `PEM` muss dazu wie folgt aussehen

* `private`
* `cert`
* `chain cert`
* `root cert`

Wenn das gegeben ist, PEM in `PKCS12` konvertieren.

``` bash
openssl pkcs12 -export -in certs.pem -out certs.pkcs12
```

Es muss ein leerer Keystore angelegt werden. Dies geschieht mittels erzeugen
und leeren eines frisch generierten Keystores.

``` bash
keytool -genkey -keyalg RSA -alias tomcat -keystore keystore.ks
keytool -delete -alias tomcat -keystore keystore.ks
```

Jetzt das generierte PKCS12 File in den (proprietären) Keystore

``` bash
keytool -v -importkeystore -srckeystore certs.pkcs12 -srcstoretype PKCS12 -destkeystore keystore.ks -deststoretype JKS
```

Selbes kann auch für einen **Truststore** durchgeführt werden. Der Befehl
dafür lautet:

``` bash
keytool -import -v -trustcacerts -alias tomcat-ca -file certs.pem -keystore truststore.ks
```

Beachten: Hier wird das `PEM` File benutzt.

## Storepasswort ändern

    keytool -storepasswd -new newpassword -keystore keystore.ks

## Key Passwort im Store ändern

Auch Keys in einem Store mit Passwort können nochmal Passwörter enthalten

    keytool -keypasswd -alias MyKeyAlias -new newpassword -keystore keystore.ks

## System Truststores

Die Trust Stores des System werden direkt mit dem JRE ausgeliefert.

```
<java_home>/jre/lib/security/cacerts

i.e.
/usr/lib/jvm/java-8-openjdk-amd64/jre/lib/security/cacerts
```

## Test

Den angelegten Keystore testen

```
java -Djavax.net.ssl.trustStore=ssl/keystore.jks -Djavax.net.ssl.trustStorePassword=<xxx> https://<nameofcert>
```

Oder mit dem Tool [SSLPoke](https://github.com/MichalHecko/SSLPoke)

```
$ /usr/lib/jvm/java-11-openjdk-amd64/bin/java \
  -Djavax.net.ssl.trustStore=/usr/lib/jvm/java-11-openjdk-amd64/jre/lib/security/cacerts \
  -Djavax.net.ssl.trustStorePassword=changeit \
  -jar SSLPoke-1.0.jar <host> 443

Successfully connected
```

