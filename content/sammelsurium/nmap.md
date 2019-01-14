---
title: nmap
date: 2012-01-19T13:59:25
tags: 
- Software
- nmap
- Pentesting
---

Net scan

    nmap -sv 10.20.30.0/24

Mit Script

    nmap -sV --script=/root/nmap/banner.nse <IP/Domain>