---
title: "AWS RDS"
date: 2020-09-01T15:37:47+02:00
tags:
- Databases
- AWS
---

AWS bietet Datenbanken an. MySQL/Oracle/Postgres. Alle davon in der Aurora
Version, angepasst aber Binärkompatibel. Ausserdem noch die Aurora Serverless
Varianten die sich bei jedem Query für eine bestimmte Zeit hochfahren.

<!--more-->

## Terraform RDS

Das schönste Modul für RDS Instanzen bringt der offizielle Terraform Account
gleich selbst mit.

```terraform
module "<name>-rds-instance" {
  source  = "terraform-aws-modules/rds/aws"
  version = "~> 2.0"

  identifier = "<name>"

  engine                = "postgres"
  engine_version        = "10.6"
  instance_class        = "db.t2.micro"
  allocated_storage     = 20
  family                = "postgres10"
  major_engine_version  = "10"
  deletion_protection   = true

  name     = "<database name>"
  username = jsondecode(data.aws_secretsmanager_secret_version.rds.secret_string)["username"]
  password = jsondecode(data.aws_secretsmanager_secret_version.rds.secret_string)["password"]
  port     = "5432"

  iam_database_authentication_enabled = true

  maintenance_window = "Mon:00:00-Mon:03:00"
  backup_window      = "03:00-06:00"

  subnet_ids = [ "..." ]

}
```

## PostgreSQL RDS via IAM

DB wird mit einem Master User und Master PW initialisiert. Dieses kann man
über `modify` wieder ändern. Jederzeit.

Verbindung:

```
psql -h <name>.<id>.eu-central-1.rds.amazonaws.com -p5432 --username lambda -d <db>
```

Muss dazu aber im selben Subnet und VPC sein.

Wenn ein User über IAM gebraucht wird um sich dort hin zu verbinden muss er
die Postgres `ROLE` `rds_iam` besitzen. Diese kommt direkt mit allen
Postgres Instanzen dazu.

```sql
GRANT rds_iam TO lambda;
```

Obacht! Dieses Role Assignment führt dazu, dass keine normalen PW Logins mehr
möglich sind. Die aufrufende Komponente braucht folgende `IAM` Komponente:

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "VisualEditor0",
            "Effect": "Allow",
            "Action": "rds-db:connect",
            "Resource": [
                "arn:aws:rds-db:eu-central-1:<acc id>:dbuser:<rds-id>/lambda"
            ]
        }
    ]
}
```

## MySQL Dump

MySQL aus RDS dumpen ist nicht so straight forward. Wegen Permissions usw.

```
mysql -u root -p<pw> -h <host> -e "show databases"
mysqldump --set-gtid-purged=OFF -u root -p<pw> -h <host> --databases <db1> <db2> <db3> > dump.sql
```
