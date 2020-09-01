---
title: "AWS SAM"
date: 2020-09-01T15:36:22+02:00
tags:
- AWS
---

SAM muss losgelöst installiert werden auf der Dev Maschine, benutzt aber
`aws-cli` Profile zum binden

<!--more-->

```
sam package --template-file sam.yaml \
  --s3-bucket <bucketname> \
  --output-template-file packaged.yaml
```

Um das Package zu deployen

```
sam deploy --template-file packaged.yaml --stack-name <name> \
  --region eu-central-1 --s3-bucket <s3bucketname> --capabilities CAPABILITY_IAM
```

Dabei wird Cloudformation benutzt

Beispiel aus der `sam.yaml`

```yaml
AWSTemplateFormatVersion: '2010-09-09'
Transform: 'AWS::Serverless-2016-10-31'
Globals:
  Function:
    Runtime: python3.6
    CodeUri: .
    MemorySize: 128
    Timeout: 30
    VpcConfig:
      SecurityGroupIds:
        - sg-xxx
      SubnetIds:
        - subnet-xxx
        - subnet-xxx
        - subnet-xxx
    Environment:
      Variables:
        PGHOST: xxx
        PGUSER: lambda
        PGDATABASE: xxx
        PGPORT: 5432
        RDS_ROLE: arn:xxx
Resources:
  GetConfiguration:
    Type: 'AWS::Serverless::Function'
    Properties:
      Handler: lambda_function.get_configuration
      Role: 'arn:xxx'
      Events:
        Api:
          Type: Api
          Properties:
            RestApiId: !Ref ApiGatewayApi
            Path: '/configs/{index}/{configuration}'
            Method: GET
```
