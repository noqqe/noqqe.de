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

## Network debugging

Ich nutze desöfteren dieses kleine Debugging Lambda um VPC Connections zu
entstören. Runtime Python3.9

```python
import json
import socket
from urllib.request import urlopen
import urllib.error

def check_http(url):
    try:
        return urlopen(url,timeout=1).code
    except urllib.error.HTTPError as e:
        return e.code
    except:
        return False

def check_tcp(uri):
    h = uri.split(":")[0]
    p = uri.split(":")[1]
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    result = sock.connect_ex((h,int(p)))
    sock.close()
    if result == 0:
        return h + ":" + p + " is open"
    else:
        return h + ":" + p + " is closed:" + result




def lambda_handler(event, context):

    ret = dict()
    ret["elk"] = check_tcp("10.11.70.26:5044")
    ret["heise"] = check_http("https://heise.de")
    ret["s3"] = check_http("https://s3.eu-central-1.amazonaws.com/")

    return {
        'statusCode': 200,
        'body': ret
    }
```
