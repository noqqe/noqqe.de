---
title: Exim
date: 2012-12-17T10:07:46.000000
tags: 
- Software
- exim
---


Mail Queue leeren

    exim -bp | exiqgrep -i | xargs exim -Mrm
