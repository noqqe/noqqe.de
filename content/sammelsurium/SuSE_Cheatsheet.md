---
title: SuSE Cheatsheet
date: 2012-02-29T13:11:13.000000
tags: 
- OS
- SUSE
- sles
---


#### Quellen refreshen

repository neu bauen

    $ zypper refs

neues repo abholen

    $ zypper refresh

suse registern und dabei die paketquellen aktualisieren

    $ suse_register

## Updates

Updates anzeigen

    $ zypper lu

Updates einspielen

    $ zypper up

Auf nächsten Service Pack upgraden

    $ wagon

Requirements anzeigen

    $ zypper info --requires spacewalk-client-tools

#### RPM

Installation und relocate in ein anderes Verzeichnis als paketiert

    $ rpm -ihv /home/noqqe/sorted/sw/jdk-7u9-linux-x64.rpm --relocate="/usr/java"="/usr/local/"

## Apache2 Wechsel von prefork auf worker

zuerst die nötigen Abhängigkeiten installieren (können co-existieren!)

    zypper install apache2-mod_fcgid php5-fastcgi

dann die sysconfig bearbeiten

    vim /etc/sysconfig/apache2
    APACHE_MPM="worker"

und dann mal durchstarten...

aber php ist kaputt.

    [Tue Jan 07 14:04:09 2014] [crit] Apache is running a threaded MPM, but
    your PHP Module is not compiled to be threadsafe.  You need to recompile
    PHP.

und dann oh mein gott ich muss alle wrapper und scheiss selber schreiben.

## Spacewalk

    spacewalk-channel -L
    spacewalk-channel -a -c "channel_name"
    rhnreg_ks --serverUrl=http://walk.example.com/XMLRPC --activationkey=1-9xxxf --force
