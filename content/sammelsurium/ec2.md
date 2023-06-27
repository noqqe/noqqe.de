---
title: "AWS EC2"
date: 2023-06-27T15:11:29+02:00
tags:
- AWS
---

EC2 Cheatsheet

<!--more-->

## AWS CLI

Eine Resource über AWS CLI erstellen oder Terraform

```
aws ec2 run-instances \
  --image-id ami-030e490c34394591b \
  --instance-type t2.micro \
  --subnet-id subnet-xxx \
  --security-group-ids sg-xxx \
  --key-name xxx \
  --count 1 \
  --tag-specifications 'ResourceType=instance,Tags=[{Key=Name,Value=test}]'
```

Eine Resource über AWS CLI löschen

```
> aws ec2 terminate-instances --instance-ids i-094687axxx
{
    "TerminatingInstances": [
        {
            "CurrentState": {
                "Code": 32,
                "Name": "shutting-down"
            },
            "InstanceId": "i-094687a9axxx",
            "PreviousState": {
                "Code": 16,
                "Name": "running"
            }
        }
    ]
}
```

### Terraform 

Via offiziellem EC2 Modul (wahlweise auch cloudposse)

```
module "ec2_instance" {
  source  = "terraform-aws-modules/ec2-instance/aws"
  version = "~> 3.0"

  instance_type          = var.instance_type
  name                   = var.name
  ami                    = var.ami
  key_name               = var.key_name
  iam_instance_profile   = var.iam_instance_profile
  monitoring             = false
  vpc_security_group_ids = var.security_group_ids
  subnet_id              = var.subnet_id
  user_data              = data.template_file.user_data.rendered
  tags                   = var.tags
}

data "template_file" "user_data" {
  template = file(var.cloudinit_file)
  vars     = var.template_vars
}
```

## IAM Profile

```terraform
locals {
  role_policy_arns = [
    "arn:aws:iam::aws:policy/service-role/AmazonEC2RoleforSSM",
    "arn:aws:iam::aws:policy/CloudWatchAgentServerPolicy"
  ]
}

resource "aws_iam_instance_profile" "this" {
  name = var.name
  role = aws_iam_role.this.name
}

resource "aws_iam_role_policy_attachment" "this" {
  count      = length(local.role_policy_arns)
  role       = aws_iam_role.this.name
  policy_arn = element(local.role_policy_arns, count.index)
}


data "aws_iam_policy_document" "this" {
  statement {
    actions = [
      "secretsmanager:ListSecrets",
      "secretsmanager:DescribeSecret",
      "secretsmanager:GetResourcePolicy",
      "secretsmanager:GetSecretValue",
      "secretsmanager:ListSecretVersionIds"
    ]
    resources = [
      "*",
    ]
    effect = "Allow"
  }
}

resource "aws_iam_role_policy" "this" {
  name   = var.name
  role   = aws_iam_role.this.id
  policy = data.aws_iam_policy_document.this.json
}

resource "aws_iam_role" "this" {
  name = var.name
  path = "/"

  assume_role_policy = jsonencode(
    {
      "Version" : "2012-10-17",
      "Statement" : [
        {
          "Action" : "sts:AssumeRole",
          "Principal" : {
            "Service" : "ec2.amazonaws.com"
          },
          "Effect" : "Allow"
        }
      ]
    }
  )
}
```
