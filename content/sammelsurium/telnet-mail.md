---
title: Telnet Mail
date: 2013-07-26T09:16:50.000000
tags: 
- HowTos
---


Telnet whooop whooop

    ## telnet mx.example.com 25
    Trying 217.145.99.52...
    Connected to mx.example.com.
    Escape character is '^]'.
    220 mx02.example.com ESMTP Postfix (Debian/GNU)
    helo example.com
    250 mx02.example.com
    mail from: support@example.com
    250 2.1.0 Ok
    rcpt to info@example.com
    501 5.5.4 Syntax: RCPT TO:<address>
    rcpt to:info@example.com
    250 2.1.5 Ok
    data
    354 End data with <CR><LF>.<CR><LF>
    subject:test-nachricht
    Dies ist eine Test-Nachricht. Bitte ignorieren.
    .
    250 2.0.0 Ok: queued as 5914D1CAEB

    500 5.5.2 Error: bad syntax

### AUTH LOGIN

Vorbereitet sein mit BASE64 Username und Passwort

    perl -MMIME::Base64 -e 'print encode_base64("passwort");'
    perl -MMIME::Base64 -e 'print encode_base64("username");'

Verschlüsselte Verbindung aufbauen

    openssl s_client -connect 1.2.3.4:465
    ---
    220 s3.neomailbox.net ESMTP
    ehlo noc.example.com
    250-mx.example.com Hello noc.example.com [127.0.0.1]
    250-SIZE 52428800
    250-PIPELINING
    250-AUTH CRAM-MD5 PLAIN LOGIN
    250-STARTTLS
    250 HELP
    AUTH LOGIN
    334 LOLO
    TROLO
    334 NOPE
    LOLOTO
    235 Authentication succeeded
    MAIL FROM:wa1@noqqe.de
    250 OK
    RCPT TO:wa1@noqqe.de
    RENEGOTIATING
    depth=1 C = BE, O = GlobalSign nv-sa, CN = AlphaSSL CA - SHA256 - G2
    verify error:num=20:unable to get local issuer certificate
    verify return:0
