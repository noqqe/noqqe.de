---
title: tar + ssh
date: 2014-02-13T13:28:25.000000
tags: 
- Software
- tar
---


3 times faster than scp

    ssh user@host tar cf - -C /dir/dir2/host . | tar xf -
