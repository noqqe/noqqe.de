---
title: Windows Service Control
date: 2012-02-17T17:41:57.000000
tags: 
- OS
- Windows
---


Status, Starten, Stoppen von Prozessen

    sc query <kurzname>
    sc start <kurzname>
    sc stop <kurzname>

## Services die keiner kennt

* wuauserv - Windows Update Service

## Alle Updates in der Konsole listen

    wmic qfe get

    wmic qfe get | findstr <string>

## Ordner löschen --force

    rd c:\temp /s /q

## NFS unter Windows

http://blogs.msdn.com/b/sfu/archive/2008/04/14/all-well-almost-about-client-for-nfs-configuration-and-performance.aspx

## Proxy Einstellungen ins Terminal importieren

    netsh winhttp import proxy source=ie

## Administrator in CMD werden

    net user Administrator /active:yes

## Anzeigen aller Umgebungsvariablen

    SET
