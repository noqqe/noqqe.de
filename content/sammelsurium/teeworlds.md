---
title: Teeworlds
date: 2013-07-26T09:17:37
tags: 
- Software
- teeworlds
---

### Example Config

```
sv_name zwetschge.org X [GER][DM]
sv_map dm1
sv_maprotation dm1 dm6 dm8
sv_scorelimit 20
sv_timelimit 20
sv_gametype dm
sv_rcon_password ordnung2
sv_motd
sv_max_clients 10
sv_spectator_slots 2
sv_register
sv_port 830X

./teeworlds_srv -f dm.cfg >> /var/log/teeworldsX.log
```