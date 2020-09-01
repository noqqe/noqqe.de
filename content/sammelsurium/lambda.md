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

