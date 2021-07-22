---
title: "Google Cloud Cli"
date: 2021-07-22T09:43:47+02:00
tags:
- Software
- GCP
---

<!--more-->

## Auth

Genereller Login

```
gcloud auth login
gcloud config set project <project>
```

## Region

Regions/Locations können _nicht_ global spezifiziert werden, nur für einzelne
Produkte...

```
> gcloud config set functions/region europe-west3
Updated property [functions/region].

> gcloud config set compute/region europe-west3
```

Zum listen der jeweiligen Region:

```
gcloud functions regions list
NAME
projects/<project>/locations/us-central1
projects/<project>/locations/us-east1
projects/<project>/locations/us-east4
```

## Completion

Eine tolle Fish Completion gibts hier:

    fisher install lgathy/google-cloud-sdk-fish-completion

Oder auch einfach auf der GCP Webseite ausführen...
