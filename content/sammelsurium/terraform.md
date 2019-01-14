---
title: "Terraform"
date: 2018-05-02T08:49:56+02:00
tags:
- Software
- AWS
---

Bevor man irgendwas macht, muss die Codebase Initialisiert sein.

```
export TF_VAR_access_key=XXX
export TF_VAR_secret_key=XXX
terraform init
```

Änderungen übernehmen/löschen

```
terraform apply
terraform plan
terraform destroy
```

Alle Module anzeigen

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

Module updaten

```
terraform get -update
```

Remote State

Damit der State nicht nur lokal vorhanden ist sondern auch für andere
Kollegen, sollte man einen Terraform.

```
terraform {
  backend "s3" {
    bucket = "terraform-state"
    key    = "<myproject>.tfstate"
    region = "eu-central-1"
  }
}
```