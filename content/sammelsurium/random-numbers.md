---
title: Random Numbers
date: 2011-06-17T11:25:44
tags: 
- Programming
- Bash
---

    for x in $(seq 1 13); do  echo -n "$((RANDOM % 9 +1 ))" ; done
