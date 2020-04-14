---
title: "Terraform und ECS Deployments in CI/CD"
date: 2019-09-06T17:32:24+02:00
tags:
- Terraform
- AWS
- ECS
- Docker
---

Wie bereits im letzten Post, dreht es sich auch dieses Mal wieder um
Terraform und AWS. Allerdings genauer gesagt um den CI/CD Prozess.

## Terraform CI/CD

Mein Terraform Code wird in einem Repository aufbewahrt und regelmäßig auf
4 Umgebungen (dev, alpha, stage, prod) ausgerollt. Diesen Prozess habe ich
mittels Bamboo automatisiert.

Als "Test"-Stage läuft in allen Umgebungen ein `terraform init && terraform
plan`. Somit kann ich sehen ob einzelne Terraform runs überhaupt
funktionieren würden.

Als "Deploy"-Stage läuft das gleiche nochmal, nur mit `terraform apply`
anstelle von `plan`.

{{< figure src="/uploads/2019/09/tfbuild.png" >}}

Coole Sache. Zumindest soweit.

## Das Problem

Ja es gibt da ein Problem. Ein eher Konzeptionelles. Mit Terraform deklariere
ich einen definierten Zustand der kompletten Infrastruktur. Das ist der Grund
warum man es benutzt (neben Automatisierung).

Das Entwicklungsteam möchte jetzt selbstständig Versionen ihrer Software deployen. Sie bauen
auch ein Release in Bamboo was in einem Docker Container in unserer internen
Registry endet.

```
nexus.acme.com/library/software:0.0.3
```

In Terraform definiere ich aber welche Version in welcher der 4 Umgebungen
läuft.

``` terraform
module "ecs_software" {
  source = "../../modules/ecs"

  tags                = var.tags
  app_name            = "software"
  health_check_path   = "/api/actuator/health"
  app_image           = "nexus.acme.com/library/software:0.0.3"
  app_port            = 8080
  [...]
}
```

Der korrekteste Weg wäre gewesen, dass mein Dev-Team bei jedem Release einen
Pull-Request im Terraform stellt. Das ist aber umständlich und langwierig.

Was aber sonst tun?

## Lösungsmöglichkeiten

Ich habe hierzu wenige Ansätze im Netz gefunden, was ein Grund ist warum ich
diesen Blogpost schreibe.

Einige schreiben, dass sie ihre [ECS Task Definition](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definitions.html)
garnicht erst in Terraform anlegen und das die Entwickler direkt selbst
machen lassen.

Andere reichen die Versionen über Umgebungsvariablen in ihr Terraform und
Releases des Dev-Teams triggern dann das Terraform Deployment. An sich ist
diese Lösung ziemlich cool, aber wer schonmal mit
[Bamboo](https://www.atlassian.com/software/bamboo) gearbeitet hat weiss was
für eine Fummelei es sein kann Cross-Projekt Tigger zu bauen. Grüße gehen
raus an die Jungs von Atlassian ;)

## Unsere Lösung

Amazon bietet die Möglichkeit Container-Tags zu redeployen. In der
[Dokumentation](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/update-service.html)
wird auf den `awscli` Befehl `ecs --force-new-deployment` verwiesen.

Was also passiert ist, das pro Umgebung ein Docker Tag (dev, alpha, stage,
prod) in unserer Registry gepflegt wird, welches in Terraform als solches
hinterlegt ist.

``` terraform
module "ecs_software" {
  source = "../../modules/ecs"

  app_image           = "nexus.acme.com/library/software:dev"
  app_port            = 8080
  [...]
}
```

Diese Environment-Tags (nenn wir sie mal so) werden natürlich nicht per Hand
gesetzt. Sie sind Mittel zum Zweck. Zusammen mit dem Dev-Team habe ich dann folgenden Deployment Job eingerichtet
(grob skizziert):

* Release in Jira wird erstellt
* Release in Bamboo wird getriggert
* `docker pull nexus.acme.com/library/software:$version` (je nach Release
  Version z.B. `0.0.3`
* `docker tag $id nexus.acme.com/library/software:$environment`
* `docker push nexus.acme.com/library/software:$environment`

Der Deploy Job läd also das akutelle Release herunter, taggt das Image mit
dem gewünschten Environment (hier `dev`) und pushed das **aktualisierte** Tag
wieder in die Docker Registry.

Als Letztes folgt dann der eingangs angesprochene API Call Richtung AWS, doch bitte den Service neu zu deployen.
Dazu loggt sich ein AWS IAM User mit `ecs:UpdateService` Permissions
über `awscli` ein und redeployed den ECS Service.

``` bash
aws ecs update-service \
  --service my-software-service \
  --cluster my-software-cluster \
  --force-new-deployment
```

Das natürlich dann 1 mal für jede Umgebung eingerichtet mit dem
entsprechenden Inhalt und separaten Zugangsdaten und IAM Usern.

{{< figure src="/uploads/2019/09/dockerdeployment.png" >}}

Für die Entwickler ist das Deployment eines neuen Release nun mit einem Klick
erledigt und dauert keine 2 Minuten. Yay.

AWS ECS spawned 2 neue Tasks und rotiert die aktuell laufenden Instanzen
raus.

{{< figure src="/uploads/2019/09/ecstasks.png" >}}

Das Ziel ist also erreicht. Der Deployment Prozess des Dev Teams kollidiert
nicht mit definierten State der Umgebung aus Terraform und trotzdem ist der
maximale Grad an Automatisierung gehalten (im Vergleich dazu den ECS Task
einfach nicht in Terraform zu pflegen). Außerdem können wir nun ständig
Deployen ohne dauernd Terraform anpassen zu müssen.
