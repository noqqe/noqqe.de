---
title: Arachni
date: 2012-01-19T13:58:38
tags: 
- Software
- arachni
- Pentesting
---

Arachni ist ein sehr mächtiges wenn auch Zeitintensives Tool (gerade mit
Directory Scanning). Für einen vollen Test kann man:

    arachni -p --mods=.* --report=txt:outfile=output.txt http://domain.com

starten. Vielleicht möchte man aber weniger bis ganz wenig Module wie XSS testen:

    arachni -p --mods=xss* --report=txt:outfile=output.txt http://domain.com

Wenn dann auch noch bisschen SQL Injection und sonstige nicht zeitintensive
Checks dabei sein sollen wirds schnell etwas länger:

    arachni -p --mods=unencrypted_password_forms,interesting_responses,captcha,os_cmd_injection_timing,sqli_blind_rdiff,xss*,xss_script_tag,sqli,code_injection_timing, --report=txt:outfile=output.txt http://domain.com

Für mehr am besten die --help lesen.