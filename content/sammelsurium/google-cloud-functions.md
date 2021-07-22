---
title: "Google Cloud Functions"
date: 2021-07-22T09:50:26+02:00
tags:
- GCP
---

Wie AWS Lambdas nur anders.

<!--more-->

## Execution

Eine Function kann im Webinterface via "Testing" ausgeführt werden.

Aber wenn der Timeout höher als 60 Sekunden ist, geht das nur über CMD

```
> gcloud functions call <fn_name> --data '{}' --region europe-west3
```

## Details

Details anzeigen via cmd.

```
> gcloud functions describe <fn_name>
availableMemoryMb: 256
buildId: ...
description: ...
entryPoint: main
environmentVariables:
  CLOUD_SQL_CONNECTION_NAME:
  DB_NAME:
httpsTrigger:
  securityLevel: SECURE_OPTIONAL
  url: ...
ingressSettings: ALLOW_ALL
labels:
  deployment-tool: terraform
name: projects/...
runtime: python37
status: ACTIVE
timeout: 300s
```

## Deploy

via `gcloud` Kommandline

```
gcloud functions deploy <fn_name> \
    --set-env-vars=DB_NAME=$DB_NAME \
    --trigger-http \
    --entry-point='main' \
    --runtime='python37' \
    --region=europe-west3 \
    --source="./cloud-function/" \
    --service-account=$service_account \
    --memory='128M'
```

via Terraform

``` terraform
# Compress source code
data "archive_file" "source" {
  type        = "zip"
  source_dir  = var.function_dir
  output_path = "/tmp/function-${var.name}.zip"
}

# Add source code zip to bucket
resource "google_storage_bucket_object" "zip" {
  name   = "${var.name}/source.zip"
  bucket = var.bucket
  source = data.archive_file.source.output_path
}

# Create the function resource resource
resource "google_cloudfunctions_function" "function" {
  name                  = var.name
  description           = var.description
  project               = var.project
  entry_point           = var.entry_point
  available_memory_mb   = var.memory
  trigger_http          = true
  environment_variables = var.environment_variables
  source_archive_bucket = var.bucket
  source_archive_object = google_storage_bucket_object.zip.name
  timeout               = var.timeout
  labels = {
    "deployment-tool" = "terraform"
  }
  region                = var.region
  runtime               = var.runtime
  service_account_email = var.service_account_email
  vpc_connector         = var.vpc_connector
}
```
