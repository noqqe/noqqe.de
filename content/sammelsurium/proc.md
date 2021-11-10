---
title: "Linux /proc"
date: 2021-10-06T09:35:34+02:00
tags:
- OS
---

Kernel Interface /proc

<!--more-->

## Prozess

Aktuelle Umgebungsvariablen eines Prozess sehen

    /proc/1123/environ

Aktuelles Working Dir anzeigen

    ls -lahd /proc/30022/cwd
    lrwxrwxrwx 1 root root 0 Oct  6 09:38 /proc/30022/cwd -> /usr/local/tomcat

Aktuelle Commandline des Prozess

    cat /proc/11988/cmdline
    /home/pwuser/chrome-linux/chrome --type=gpu-process --no-sandbox --disable-dev-shm-usage --disable-breakpad --headless

Prozessname

    cat /proc/11988/comm
    chrome

Konsumierter I/O Traffic

    cat /proc/11988/io


## Mounts

Anstelle von `mount` kann man auch als fstab formatiertes File ausgeben:

    cat /proc/mounts

Hilfreich:

    diff -y /proc/mounts /etc/fstab

## Peripherie

Cpus anzeigen

    cat /proc/cpuinfo
