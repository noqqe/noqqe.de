---
title: "AWS IAM"
date: 2020-09-30T14:38:31+02:00
tags:
- AWS
---

IAM verwaltet die ACLs innerhalb AWS Komponenten, Benutzern und Services

<!--more-->

## Terraform natives Policy Document

Je nach Möglichkeit sollte man einen Terraform `data` Provider benutzen um
die Policies abzubilden. Der eingebettete `JSON` Weg ist ziemlich hässlich
und wird mittels `EOF` Multiline String nicht schöner!

```terraform
data "aws_iam_policy_document" "acme-lambda-policy" {

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
      "${var.s3_lambda_bucket}/*",
      var.s3_lambda_bucket,
    ]

    effect = "Allow"
  }

  statement {
    actions   = ["kms:Encrypt", "kms:Decrypt", "kms:GenerateDataKey"]
    resources = [var.kms_key_arn]
    effect    = "Allow"
  }

  statement {
    effect    = "Allow"
    resources = ["*"]
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents"
    ]
  }

}
```

Dieses Policy Document enthält mehrere Statements und kann via einer Policy
Resource verwendet werden

```terraform
resource "aws_iam_policy" "acme-lambda-policy" {
  name        = "${var.environment}-acme-lambda-policy"
  path        = "/"
  policy      = data.aws_iam_policy_document.acme-lambda-policy.json
}
```

## Assume

... Todo...

## User Generierung mit Access Tokens

Nachdem ich immer wieder danach Suche, hier kurz dokumentiert.

Ein programmatischer User mit Access ID und Secret Key der Zugriff auf
bestimmte eingebettete Policies hat. Zum Beispiel Upload in einen S3 Bucket.

```terraform
data "aws_iam_policy_document" "acme-deployment" {
 statement {
   actions   = [...]
   resources = [...]
   effect    = "Allow"
 }
}

resource "aws_iam_user_policy" "deployment" {
 name   = "${aws_iam_user.deployment.name}-policy"
 user   = aws_iam_user.deployment.name
 policy = data.aws_iam_policy_document.acme-deployment.json
}

resource "aws_iam_user" "deployment" {
 name          = "${var.environment}-acme-deployment"
 force_destroy = "true"
}

resource "aws_iam_access_key" "deployment" {
 user = aws_iam_user.deployment.name
}
```
