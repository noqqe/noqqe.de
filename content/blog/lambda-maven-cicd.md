---
title: "Nexus und Lambda Deployments in CI/CD"
date: 2020-10-06T09:48:57+02:00
tags:
- fill
---

In letzter Zeit benutze ich gehäuft AWS Lambda für ein Projekt. Konkret rolle
ich dieses und zugehörige Permissions via [Terraform](https://terraform.io)
aus. In AWS stellen wir via [API Gateway](https://aws.amazon.com/api-gateway/) eine kleine
REST-Schnittstelle zur Verfügung die ein [Lambda (Java 11)](https://aws.amazon.com/lambda/) aufruft.

<!--more-->

{{< figure src="/uploads/2020/10/servers.png" >}}

Das Problem verhält sich ähnlich wie [beim letzten Mal](/blog/2019/09/06/terraform-und-ecs-deployments/).
Das Entwicklerteam möchte gerne losgelöst von Terraform neue Versionen des
Lambdas deployen. Um nicht jedesmal den Sysadmin anpingen zu müssen muss die
Lambda Funktion also losgelöst versioniert werden.

## Nexus .jar Download

Im Source Build Plan des Java Projekts endet eine `.jar` Datei im [Nexus Repository](https://www.sonatype.com/nexus/repository-pro)
Das Artifact selbst ist dann im Nexus mit folgenden Details hinterlegt:

```xml
<dependency>
  <groupId>com.acme.comp.compiler</groupId>
  <artifactId>compiler</artifactId>
  <version>0.0.10</version>
</dependency>
```

Anfangs habe ich versucht das Artifact über die Nexus REST API
herunterzuladen:

```bash
curl -vL -u <user>:<pass> -X GET \
  'https://nexus.acme.com/service/rest/v1/search/assets/download?sort=version&repository=<repo>&group=com.acme.comp.compiler&name=compiler&maven.baseVersion=0.0.10&maven.extension=jar' \
  --output compiler.jar
```

Aber das fand ich nicht so schön. Ein Kollege hat dann einen Command über
`mvn` gebaut, der schöner und kürzer ist.

```bash
mvn dependency:copy \
  -Dartifact=com.acme.comp.compiler:compiler:0.0.10 \
  -DoutputDirectory=. \
  -Dmdep.stripVersion=true
```

## S3 Upload

Nachdem das Artifact nun im aktuellen Deploy Plan verfügbar ist, kann es
hochgeladen werden in den S3 Bucket, aus dem die Lambda Funktion ihren Code
abholt.

```bash
aws configure set aws_access_key_id $bamboo_AWS_ACCESS_KEY
aws configure set aws_secret_access_key $bamboo_AWS_SECRET_ACCESS_KEY
aws configure set region eu-central-1
aws s3 cp compiler.jar s3://acme-lambda/compiler.jar
```

## Lambda Update

Das reicht aber noch nicht ganz. Jedes mal wenn sich das Objekt im S3 Bucket
ändert, muss auch die Lambda Funktion aktualisiert werden. Zum Glück bietet
das AWS CLI auch dafür eine Möglichkeit

```bash
aws lambda update-function-code --s3-bucket acme-lambda --s3-key compiler.jar  --function-name acme-lambda
```

Danach ist die Lambda Funktion mit der aktuellen Version ausgerollt ohne das
wir Terraform aktualisieren mussten.

Natürlich müssen alle Variablen in allen Beispielen hier mit dynamischen
Releases eurer CI/CD Platform befüllt werden. Ähnlich wie `$bamboo_AWS_SECRET_ACCESS_KEY`.
