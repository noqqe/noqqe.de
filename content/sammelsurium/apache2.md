---
title: Apache2
date: 2013-12-14T14:56:28.000000
tags: 
- Software
- Apache
---


## Print vHost Configuration

`apache2ctl -St` shows the following

    VirtualHost configuration:
    *:80     is a NameVirtualHost
             default server  (/etc/apache2/sites-enabled/000-default:4)
             port 80 namevhost  (/etc/apache2/sites-enabled/000-default:4)
             port 80 namevhost domain1.com (/etc/apache2/sites-enabled/www.domain1.com:4)
             port 80 namevhost www.domain2.com (/etc/apache2/sites-enabled/www.domain2.com:4)
    Syntax OK

## Apache2.4 Satisfy Any

Grant stuff changed a lot since the last version

    <Directory /var/www/html>
      AuthType basic
      AuthName "private"
      AuthUserFile /etc/htpasswd
      <RequireAny>
        Require valid-user
        Require ip 10.20.30.40
      </RequireAny>
    </Directory>
