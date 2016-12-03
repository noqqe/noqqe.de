---
title: w3af
date: 2012-01-19T13:57:15
tags: 
- Software
- w3af
- Pentesting
---

Normaler aufruf mit Script

    w3af_console -s /root/w3af/w3af.script

Normales Ablauf Skript mit konfigurierten reports in html und txt

~~~
http-settings
set userAgent "Mozilla/5.0 (X11; U; Linux i686; it; rv:1.9.2.13) Gecko/20101206 Ubuntu/10.10 (maverick) Firefox/3.6.13"
set timeout 5
back
plugins
output console,textFile,htmlFile
output config htmlFile
set verbosity 10
set fileName /root/w3af/report.html
back
output config textFile
set verbosity 10
set fileName /root/w3af/report.txt
output
back
output config console
set verbosity 5
back
discovery webSpider,allowedMethods,userDir,sitemapReader,robotsReader,pykto,hmap
audit xss,sqli,xpath,remoteFileInclude,osCommanding,localFileInclude,htaccessMethods,fileUpload,eval,blindSqli
back
target
set target http://hack.example.com/auth/login.php
back
start
~~~

Volle Discovery und Audit Einstellungen:

~~~
discovery MSNSpider,afd,allowedMethods,archiveDotOrg,content_negotiation,detectReverseProxy,detectTransparentProxy,digitSum,dir_bruter,dnsWildcard,domain_dot,dotNetErrors,favicon_identification,findBackdoor,
findCaptchas,findvhost,fingerGoogle,fingerMSN,fingerPKS,fingerprint_WAF,fingerprint_os,frontpage_version,ghdb,googleSpider,halberd,hmap,importResults,oracleDiscovery,phishtank,phpEggs,phpinfo,
ria_enumerator,robotsReader,serverHeader,serverStatus,sharedHosting,sitemapReader,slash,spiderMan,urlFuzzer,userDir,webDiff,webSpider,wordnet,wordpress_fingerprint,wsdlFinder,
xssedDotCom,yahooSiteExplorer,zone_h

audit LDAPi,blindSqli,buffOverflow,dav,eval,fileUpload,formatString,frontpage,generic,globalRedirect,htaccessMethods,localFileInclude,mxInjection, osCommanding,phishingVector,preg_replace,redos,remoteFileInclude,responseSplitting,sqli,ssi,sslCertificate,unSSL,xpath,xsrf,xss,xst
~~~
