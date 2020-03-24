---
title: "Java Keytool"
date: 2020-03-23T10:04:11+01:00
tags:
- Software
- Databases
- OS
---

"Kannst du mal eben ein Tomcat Zertifikat austauschen?". Gänsehaut.

Oracles Werkzeug zur Zertifikatsverwaltung macht keinen Spass.

## Anzeigen

Anzeigen der Inhalte eines Keystores.

```
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

```
openssl pkcs12 -export -in certs.pem -out certs.pkcs12
```

Es muss ein leerer Keystore angelegt werden. Dies geschieht mittels erzeugen
und leeren eines frisch generierten Keystores.

```
keytool -genkey -keyalg RSA -alias tomcat -keystore keystore.ks
keytool -delete -alias tomcat -keystore keystore.ks
```

Jetzt das generierte PKCS12 File in den (proprietären) Keystore

```
keytool -v -importkeystore -srckeystore certs.pkcs12 -srcstoretype PKCS12 -destkeystore keystore.ks -deststoretype JKS
```

Selbes kann auch für einen **Truststore** durchgeführt werden.

## Test

Den angelegten Keystore testen

```
java -Djavax.net.ssl.trustStore=ssl/keystore.jks -Djavax.net.ssl.trustStorePassword=<xxx> https://<nameofcert>
```
