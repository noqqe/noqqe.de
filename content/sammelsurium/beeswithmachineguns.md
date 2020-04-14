---
title: beeswithmachineguns
date: 2015-01-22T08:52:28
tags: 
- Software
- beeswithmachineguns
---

#### Usage

Basic usage, provide group, key, zones, amount of bees

    bees up -s 6 -g public -k ssh-user-key -z us-east-1a

The shitty part is with zones. Can easily end up in exceptions.

#### Config

Example config.
Zone specification does not work.. nevermind.

``` ini
[Credentials]
aws_access_key_id = AAABBBCCC
aws_secret_access_key = AAABBBCCC

[Boto]
debug = 0
num_retries = 10
ec2_region_name = us-west-1
cs_region_name = us-west-1
placement = us-west-1
```