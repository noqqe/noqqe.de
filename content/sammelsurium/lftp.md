---
title: lftp
date: 2013-09-12T12:45:49
tags: 
- Software
- lftp
---

## Upload a file

    lftp -e 'cd /somedir/ ; put test2-with-some-load.tar.gz ; bye' -u general,pass ftp.example.com
