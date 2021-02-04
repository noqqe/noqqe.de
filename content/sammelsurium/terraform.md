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
