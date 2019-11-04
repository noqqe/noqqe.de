---
title: "Terraform Secrets in AWS Secrets Manager"
date: 2019-09-02T10:40:21+02:00
tags:
- Terraform
- AWS
---

Passwörter. Ein Quell der Freude, gerade beim Thema Infrastructure as Code
(in meinem Fall mittels [terraform](https://terraform.io))

Ich arbeite im Moment an einem neuen Projekt in der Arbeit in der ich
4 identische Umgebungen in [AWS](https://aws.amazon.com/) aufbauen darf und
natürlich werden dort diverse Systeme angebunden welche sich mittels
Authentifizierung unterhalten.

Vorher sah meine Codebase ungefähr so aus. Credentials wo man hinsah.

```
module "mysql_db" {
  source = "../../modules/rds"

  name                = "${var.environment}-mysql"

  subnet_ids          = module.dcc_vpc.private_subnet_ids
  tags                = var.tags
  vpc_id              = module.dcc_vpc.vpc_id
  password            = "kommtnieeinerdrauf42"
}

module "container" {
  source = "github.com/cloudposse/terraform-aws-ecs-container-definition"
  [...]

  repository_credentials = {
    username = "foo"
    password = "garantiertsicher23"
  }
}

module "ecs_task" {

  app_port = 8080
  app_env = [
    {
      name  = "DB_PASSWORD"
      value = "liebernichtinsrepocomitten1234"
    }
  ]
```

Natürlich bin ich nicht der Einzige mit solchen Problemen.
Es gab dann mehrere Optionen wie ich damit umgehe.

* Ignorieren (wie in anderen Projekten)
* Alles in Umgebungsvariablen auslagern
* [Hashicorp Vault](https://www.vaultproject.io/)
* [Secrets Provider](https://github.com/tweag/terraform-provider-secret)
* [AWS Secrets Manager](https://aws.amazon.com/secrets-manager/)

Ich habe mich dann für Letzteres entschieden. Vorrangig aus Neugier und
Bequemlichkeit. Mein Cloud Provider stellt mir ein Tool zur Verfügung und
mein Tooling unterstützt es. Also warum nicht.

Zuerst muss ich mich dann im Webinterface der AWS anmelden und die
Credentials hinterlegen.

{{< figure src="/uploads/2019/09/create.png" >}}

und danach aus der Übersicht die `ARN` heraussuchen. Diese wird danach wieder
bei der Anbindung über einen `Data Provider` in `terraform` benötigt.

{{< figure src="/uploads/2019/09/secretsmanager.png" >}}

Das fühlt sich am Anfang irgendwie falsch an, aber wenn man bedenkt das die
Credentials gerade deshalb nicht im Repo verwaltet werden, damit man sie eben
nicht im Repo hat gleicht das das dreckige Gefühl aus manuelle `ARN`s in den
Code zu pasten.

Im `data.tf` File der jeweiligen Umgebung kann ich mir die Resourcen jetzt
zum Weiterverarbeiten abholen:

```
data "aws_secretsmanager_secret" "db" {
  arn = "arn:aws:secretsmanager:eu-central-1:xxx"
}

data "aws_secretsmanager_secret_version" "db" {
  secret_id = data.aws_secretsmanager_secret.db.id
}
```

Wenn ich dann die Secrets im Code brauche, rufe ich sie über
`data.aws_secretsmanager_secret_version` ab.

```
module "mysql_db" {
  source = "../../modules/rds"

  name                = "${var.environment}-mysql"

  subnet_ids          = module.dcc_vpc.private_subnet_ids
  tags                = var.tags
  vpc_id              = module.dcc_vpc.vpc_id
  password            = jsondecode(data.aws_secretsmanager_secret_version.mysql.secret_string)["password"]
}

module "container" {
  source = "github.com/cloudposse/terraform-aws-ecs-container-definition"
  [...]

  repository_credentials = jsondecode(data.aws_secretsmanager_secret_version.repo.secret_string)
}
```

So nutze ich nun neben ALB, ECS, Fargate, RDS & Cloudwatch eben noch einen
weiteren AWS Service: Secrets Manager.

Alle anderen Optionen (siehe oben) gehen natürlich auch. Na gut, vielleicht
bis auf die Erste. Ich hatte allerdings keine Lust eine Vault Instanz selbst
zu betreiben, auch wenn Vault die am Besten unterstützte Lösung gewesen wäre.
Die Umgebungsvariablen fielen deswegen aus weil wir `terraform` mittels
Bamboo CI/CD im Prozess benutzen und dann dort die Credentials überall
verteilt wären.

Ich denke meine nächste Wahl wäre der Secrets Provider gewesen, da die
Credentials hier nur 1x in das Statefile importiert werden müssen und dann
für immer dort bestehen bleiben und auch auf Build Agents als auch auf den
Rechnern der Kollegen verfügbar wären. Mit dem AWS Service hatte ich mich
dann aber besser gefühlt.
