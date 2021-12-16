---
title: "AWS Secrets Manager Credentials in EC2 Instanzen verarbeiten"
date: 2021-12-16T10:42:05+01:00
tags:
- AWS
- EC2
- SecretsManager
- Terraform
- json
---

AWS hat ein Produkt, dass [AWS Secretsmanager](https://aws.amazon.com/secrets-manager/) heisst und einem unglaublich
wenig Speicherplatz für unglaublich teuer Geld anbietet.

Ab und an braucht man Geheimnisse auf seiner EC2 Instanz. Geheimnisse die ich
nicht in meinen Sourcecode einchecken will, aber irgendie maschinenlesbar
abgespeichert werden müssen.

Mittels oben genanntem SecretsManager und dem `aws` Commandline Tool lassen sich diese Secrets anzeigen.

<!--more-->

# Pricing?

{{< figure src="/uploads/2021/12/pricing.png" caption="AWS Secretsmanager Pricing" >}}

Ein Secret kostet $0.40. Außerdem darf ein Secret maximal 65536 bytes groß sein.
Wenn man das auf den Preis pro GB hochrechnet, kostet 1GB schlappe $6103,20.
Zum Vergleich, 1GB [S3](https://aws.amazon.com/s3/) kostet $0.023 pro GB.

Ich schweife ab.

# IAM Polciies

Der normale Weg mit `aws` Commandline Tool wäre, mich mittels Access Key und
Secret Key einzuloggen. Da ich aber nicht auf jeder meiner EC2 Instanzen
Credentials ablegen will, kann ich eine IAM Policy erstellen, die jeder
Instanz ohne Credentials den Zugang zu bestimmten Secrets erlaubt.

Beispielhaft ein Terraform IAM Policy Document

```terraform
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
      "arn:aws:secretsmanager:eu-central-1:xxx:secret:prod_rds-xxx",
    ]
    effect = "Allow"
  }
}
```

# `aws` Commandline Tool und `jq`

Wenn das Instanzprofil jetzt der EC2 Instanz zugewiesen ist, kann man `aws`
und `jq` benutzen, um die Values aus dem Secret auszulesen.

```json
$ aws secretsmanager get-secret-value --secret-id prod_rds
{
    "Name": "prod_rds",
    "SecretString": "{\"password\":\"secret1\",\"user\":\"username1\"}",
    "VersionStages": [
        "AWSCURRENT"
    ],
    "CreatedDate": 1637933085.518,
    "ARN": "arn:aws:secretsmanager:eu-central-1:xxx:secret:prod_rds-xxx"
}
```

Problem hierbei ist, dass der Output in Json und der SecretString erneut in
Json encodiert ist. Aber dafür gibt es eine Lösung. `jq`.

```json
$ aws secretsmanager get-secret-value --secret-id prod_rds \
  | jq -r '.SecretString'

{"password":"secret1","user":"username1"}
```

Der Trick ist jetzt, mittels `fromjson` die Value eines Attributs erneut zu
parsen. Danach kann man nämlich auf die eingebetteten Attribute wie
`password` zugreifen.


```
$ aws secretsmanager get-secret-value --secret-id prod_rds \
  | jq -r '.SecretString | fromjson | .password'

secret1
```

Das kann ich jetzt mittels Ansible oder CloudInit oder sonst was für meine
eigentliche Applikation einbetten und weiterverarbeiten.

Credentials im Repo?
Nein danke :)
