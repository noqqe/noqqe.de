---
title: "Terraform"
date: 2018-05-02T08:49:56+02:00
tags:
- Software
- AWS
---

## Bevor man irgendwas macht, muss die Codebase Initialisiert sein.

``` bash
export TF_VAR_access_key=XXX
export TF_VAR_secret_key=XXX
terraform init
```

## Änderungen übernehmen/löschen

``` bash
terraform apply
terraform plan
terraform destroy
```

## Alle Module anzeigen

```
terraform providers
.
├── provider.aws ~> 1.8.0
├── module.bastion_host
│   ├── provider.aws (inherited)
│   └── provider.template
└── module.vpc
    └── provider.aws (inherited)
```

## Module updaten

``` bash
terraform get -update
```

## Remote State

Damit der State nicht nur lokal vorhanden ist sondern auch für andere
Kollegen, sollte man einen Terraform.

``` terraform
terraform {
  backend "s3" {
    bucket = "terraform-state"
    key    = "<myproject>.tfstate"
    region = "eu-central-1"
  }
}
```

## Terraform State

Der State ist eine wichtige Komponente, mit der man viel mehr machen muss als
anfangs mal gedacht.

```
terraform state show
terraform state rm
```

Seit 0.13 gibt es eine neue Schreibweise für Terraform Provider
damit man im alten State einen neue Provider für bereits bestehende
Ressourcen umkonvertieren kann, gibt es folgenden Command

```
terraform state replace-provider "registry.terraform.io/-/google" \
"hashicorp/google"
```


## Eine Liste erweitern

Ein einfaches `concat` auf alle Elemente die in der neuen Liste sein sollen.

``` terraform
concat(data.aws_vpc.oozie_vpc.ids, ["vpc-xxx"])
```

## Data Provider Filters

Um eine Liste von allen Subnets in einem AWS VPC zu erhalten kann man
dynamisch auf Tag Namen filtern.

``` terraform
data "aws_subnet_ids" "xxx_subnets" {
  vpc_id = data.aws_vpc.xxx_vpc.id

  tags = {
    Name = "*-private-*"
  }
}
```

## If/Else Statement

Wenn man aufgrund einer Variable entscheiden will ob eine Ressource erstellt
werden soll oder nicht:

``` terraform
resource "aws_alb" "internal" {
  count           = var.create_internal_lb == true ? 1 : 0
  name            = "foo"
  subnets         = var.subnet_ids
}
```

## for_each mit Liste

```
 locals {
   mail_receipients = [
     "foo@example.net",
     "bar@example.net",
   ]
 }

module "sns_email" {
  source = "github.com/noqqe/terraform-modules-aws-sns-email-notification?ref=v0.3.0"

  for_each              = toset(local.mail_receipients)
  application_name      = "my-cloudwatch-alert"
  notification_endpoint = each.key
}

## Dynamic Blocks

Man kann dynamische blocks wie zum Beispiel

``` terraform
default_action {
  type             = "forward"
  target_group_arn = aws_alb_target_group.app.id
}
```

auch an eine Condition knüpfen. Dazu kann man mit dem Keyword `dynamic`
via `for_each` jene Bedingung abbilden.

``` terraform
dynamic "default_action" {
  for_each = var.http_redirect_to_https == false ? [var.http_redirect_to_https] : []
  content {
    type             = "forward"
    target_group_arn = aws_alb_target_group.app.id
  }
}
```

## Data Blocks für Policies

Anstelle von blöden File Includes (relative Pfade, Lesbarkeit, etc..) können
`json` Policy Objekte auch in Terraform nativ kodiert und dann als `.json`
Methode aus dem `data` Provider abgerufen werden

```terraform
data "aws_iam_policy_document" "this" {
  statement {
    actions = [
      "s3:ListBucket",
      "s3:ListBucketMultipartUploads",
      "s3:ListBucketVersions",
      "s3:ListMultipartUploadParts",
      "s3:PutObject",
      "s3:GetObject",
      "s3:DeleteObject",
      "s3:GetObjectVersion",
      "s3:DeleteObjectVersion",
      "s3:PutObjectAcl",
      "s3:GetObjectAcl",
    ]
    resources = [
      aws_s3_bucket.this.arn,
      "${aws_s3_bucket.this.arn}/*",
    ]
    effect = "Allow"
  }
}
```

Aufruf:

```terraform
resource "aws_iam_user_policy" "this" {
  name   = "Access${aws_iam_user.this.name}"
  user   = aws_iam_user.this.name
  policy = data.aws_iam_policy_document.this.json
}
```

## Terraform Secrets

Seit TF14 können secrets nur noch sensitiv als output ausgegeben werden.
Das ist vor allem eins: nervig.

```
> terraform state pull | jq '.resources[] | select(.type == "aws_iam_access_key") | .instances[0].attributes'
{
  "create_date": "2019-11-05T08:39:56Z",
  "encrypted_secret": null,
  "encrypted_ses_smtp_password_v4": null,
  "id": "AKIA.....",
  "key_fingerprint": null,
  "pgp_key": null,
  "secret": "...",
  "ses_smtp_password_v4": null,
  "status": "Active",
  "user": "<username>"
}
```

Interaktiv auf Resource zugreifen:

```
$ terraform console
> nonsensitive(aws_iam_access_key.<user>.encrypted_secret)
```

Workaround um trotzdem in normalen TF Run die Variable ausgeben zu können

```
resource "aws_iam_access_key" "brand_new_user" {
  user = aws_iam_user.brand_new_user.name
}

output "brand_new_user_id" {
  value = aws_iam_access_key.brand_new_user.id
}

data "template_file" "secret" {
  template = aws_iam_access_key.brand_new_user.encrypted_secret
}

output "brand_new_user_secret" {
  value     = data.template_file.secret.rendered
}
```
