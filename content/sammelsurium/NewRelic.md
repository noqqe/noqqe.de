---
title: NewRelic
date: 2014-04-03T15:20:17.000000
tags: 
- Software
- newrelic
---


#### SLES

folgende Schritte

    zypper install python
    zypper install python-setuptools
    cat /etc/ssl/certs/*.pem > /etc/ssl/certs/ca-certificates.crt
    easy_install pip
    pip --cert /etc/ssl/certs/ca-certificates.crt install newrelic_plugin_agent
    ## edit requests module for ssl verify off (broken SLES)
    vim /usr/local/lib64/python2.6/site-packages/newrelic_plugin_agent/agent.py
    ## create dirs
    mkdir /var/log/newrelic
    mkdir /var/run/newrelic
    mkdir /etc/newrelic/
    easy_install --upgrade requests
    newrelic_plugin_agent -c /etc/newrelic/newrelic_plugin_agent.cfg -f
    /etc/init.d/newrelic-memcached start
