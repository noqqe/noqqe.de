---
title: maildrop
date: 2013-08-02T12:23:34.000000
tags: 
- Software
- maildrop
---


Bogofilter

   	xfilter "bogofilter -u -e -p -R -c /home/noqqe/.bogofilter.cf"

		## Filter with bogofilter Spam or Unsure
		if ( /^X-Bogosity: Spam, tests=bogofilter/:h ) {
		  to "$VUSERMAILDIR/.$SPAMDIR/"
		}
