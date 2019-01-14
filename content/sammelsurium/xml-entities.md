---
title: XML Entities
date: 2013-08-01T11:43:11
tags: 
- Programming
---

Inhalt des Files /data/x6/crawler/common/y.Timestamp.ent

    <!ENTITY timestamp "20130801112000">

File in das inkludiert wird

```
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE crawlerConfig SYSTEM "config.dtd" [
        <!ENTITY config SYSTEM "/data/x6/crawler/common/y.configKapitel.ent">
        <!ENTITY pi "PI191">
        <!ENTITY pathideskbuild "/data/x6/indices/y/build">
        <!ENTITY configidesk "/data/x6/crawler/y">
        <!ENTITY pathextracted "/data/x6/extracted/build">
        <!ENTITY xsltdir "/data/x6/misc/xslt">

        <!ENTITY % timestamp SYSTEM "/data/x6/crawler/common/y.Timestamp.ent">
        %timestamp;

]>
<crawler>
        &config;
</crawler>
```
