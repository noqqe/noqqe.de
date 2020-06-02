---
title: "jq"
date: 2020-05-29T13:05:12+02:00
tags:
- Software
- json
- jq
---

`jq` ist sozusagen das `sed` für `JSON`.

<!--more-->

## Generelle Syntax / Felder anzeigen

Wenn das Feld ausgelesen werden soll `def`

```json
[{
    "id": 123,
    "name": "John",
    "aux": [{
        "abc": "random",
        "def": "I want this"
    }],
    "blah": 23.11
}]
```

```bash
jq '.[].aux[].def' file.json
```

Dabei sind Arrays `[]` und Attribute davon `.aux` bzw. `.def`

## Elemente in Array filtern

```json
{
  "prefixes": [
    {
      "ip_prefix": "35.180.0.0/16",
      "region": "eu-west-3",
      "service": "AMAZON",
      "network_border_group": "eu-west-3"
    },
    {
      "ip_prefix": "52.94.76.0/22",
      "region": "us-west-2",
      "service": "AMAZON",
      "network_border_group": "us-west-2"
    }
  ]
}
```

Ich kann also nach allen Feldern in denen "region" "eu-central-1" ist, suchen
in dem ich den `select` Command benuzte

```bash
jq '.prefixes[] | select(.region == "eu-central-1")' f
```

## Multiple Filter

Filter (wie `select`) können mittelns `|` gechained werden.

```bash
jq '.prefixes[] | select(.region == "eu-central-1") | select(.ip_prefix | startswith("3."))' f`
```

Hier verbinde ich 2 Suchkritierien (json selbes wie oben)
