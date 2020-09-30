---
title: "AWS Lambda"
date: 2020-09-01T15:27:32+02:00
tags:
- AWS
---

<!--more-->

## Invoking

Um ein Lambda von der Commandline auszuführen:

    aws lambda invoke --function-name <name> -

Allerdings ohne Parameter und auch keine Ausgabe von Logs.

## Update Function in Place

Für Automatisierung in CI/CD

```
aws lambda update-function-code --s3-bucket <bucketname> \
  --s3-key path/lambda.zip  --function-name <name>
```

## Terraform

```terraform
resource "aws_lambda_function" "acme-lambda" {
  function_name = "${var.environment}-acme-lambda"
  role          = aws_iam_role.lua-lambda-role.arn
  handler       = var.handler
  memory_size   = 2048
  timeout       = 75
  s3_bucket     = var.s3_bucket
  s3_key        = var.s3_filename
  runtime       = "java11"

  environment {
    variables = {
      PATH = "/usr/local/bin:/usr/bin/:/bin:/opt/bin:/var/task/"
    }
  }
}
```

## Logging

Siehe [Details](https://docs.aws.amazon.com/lambda/latest/dg/python-logging.html)

```python
import os

def lambda_handler(event, context):
    print('## ENVIRONMENT VARIABLES')
    print(os.environ)
    print('## EVENT')
    print(event)
```

