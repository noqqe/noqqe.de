---
title: "AWS API Gateway"
date: 2022-06-22T12:49:34+02:00
tags:
- AWS
---

<!--more-->

## Export API Definition

https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-export-api.html

in openapi 3

    aws apigateway get-export --parameters extensions='apigateway' --rest-api-id <id> --stage-name preprod --export-type oas30 export.json

in swagger

    aws apigateway get-export --parameters extensions='apigateway' --rest-api-id <id> --stage-name preprod --export-type swagger swagger.json

