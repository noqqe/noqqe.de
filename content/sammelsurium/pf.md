---
title: pf
date: 2012-11-04T13:15:42
tags: 
- pf
- OS
- OpenBSD
---

To enable basic pf on OpenBSD

    ## /etc/rc.conf.local
    pf_enable="YES"
    pf_rules="/etc/pf.conf"
    pflog_enable="YES"
    pflog_logfile="/var/log/pflog"

#### Quickies

Clean pf - deletes all rules

    /usr/local/bin/sudo /sbin/pfctl -d

Parse but do not load it

    sudo pfctl -nf /etc/pf.conf

Show Content of a table :

    pfctl -t $tablename -T show

Show Details for a table

    pfctl -vvsTables [$tablename]

Add IP to Table

    pfctl -t spammers -T add 1.2.3.0/24

#### pf show

    pfctl -f /etc/pf.conf     Load the pf.conf file
    pfctl -nf /etc/pf.conf    Parse the file, but don't load it
    pfctl -sr                 Show the current ruleset
    pfctl -ss                 Show the current state table
    pfctl -si                 Show filter stats and counters
    pfctl -sa                 Show EVERYTHING it can show

## pf labels (static)

show labels

    pfctl -sl

implement labels in pf.conf

    pass in on $extif proto tcp from any to any port 9 label discard

## pf labels (dynamic)

at some point, static labeling is not enough

    ## Allow outgoing
    pass out on $extif proto tcp to any port $tcpout label "tcp:out:$dstport"
    pass out on $extif proto udp to any port $udpout label "udp:out:$dstport"

    ## Incoming rule
    pass in on $extif proto tcp from any to any port $tcpin label "tcp:in:$dstport"
    pass in on $extif proto udp from any to any port $udpin label "udp:in:$dstport"

## pflog device

todo...