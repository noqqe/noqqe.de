---
title: pdf
date: 2017-08-27T12:15:30
tags:
- Software
- pdf
---


# PDF Passwort entfernen

    qpdf --password=passw0rd --decrypt foo.pdf unsecure-foo.pdf

falls das Passwort nicht bekannt sein sollte, kann man es mit Ghostscript
entfernen

    gs -q -dNOPAUSE -dBATCH -sDEVICE=pdfwrite -sOutputFile=unencrypted.pdf -c .setpdfwrite -f encrypted.pdf

# PDFs mergen

    pdftk novell-cert.pdf ripe.pdf lpic-1.pdf cisco-ccna1.pdf cat output Zertifikate.pdf

# PDFs aufteilen

    pdftk foo.pdf cat 1-12 output bar.pdf

