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

CPU Limits / Ulimit für den aktuellen Prozess

    cat /proc/12/limits
    Limit                     Soft Limit           Hard Limit           Units
    Max cpu time              unlimited            unlimited            seconds
    Max file size             unlimited            unlimited            bytes
    Max data size             unlimited            unlimited            bytes
    Max stack size            8388608              unlimited            bytes
    Max core file size        0                    unlimited            bytes
    Max resident set          unlimited            unlimited            bytes
    Max processes             1019662              1019662              processes
    Max open files            1024                 4096                 files
    Max locked memory         65536                65536                bytes
    Max address space         unlimited            unlimited            bytes
    Max file locks            unlimited            unlimited            locks
    Max pending signals       1019662              1019662              signals
    Max msgqueue size         819200               819200               bytes
    Max nice priority         0                    0
    Max realtime priority     0                    0
    Max realtime timeout      unlimited            unlimited            us

## Mounts

Anstelle von `mount` kann man auch als fstab formatiertes File ausgeben:

    cat /proc/mounts

Hilfreich:

    diff -y /proc/mounts /etc/fstab

## Peripherie

Cpus anzeigen

    cat /proc/cpuinfo

## Threads

Maximale Threads anzeigen

    cat /proc/sys/kernel/threads-max
    2039325
