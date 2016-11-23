---
title: supervisor
date: 2015-05-24T15:46:46.000000
tags: 
- Software
- supervisor
---


Automatisches Starten/Stoppen von Diensten
Beispielconfig sieht ungefähr so aus:

    [program:rezeptionistin]
    directory=/home/k4cg/Code/Rezeptionistin
    command=/home/k4cg/Code/Rezeptionistin/rezeptionistin.py
    autostart=true
    autorestart=true
    startsecs=10
    stdout_logfile=/home/k4cg/Code/Rezeptionistin/logs/rezeptionistin.log
    stdout_logfile_maxbytes=1MB
    stdout_logfile_backups=10
    stdout_capture_maxbytes=1MB
    stderr_logfile=/home/k4cg/Code/Rezeptionistin/logs/rezeptionistin.log
    stderr_logfile_maxbytes=1MB
    stderr_logfile_backups=10
    stderr_capture_maxbytes=1MB
    environment = HOME="/home/k4cg", USER="k4cg"

und dann einfach über

    service supervisor start
    service supervisor stop

bedienen

