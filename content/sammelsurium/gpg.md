---
title: gpg
date: 2013-09-25T20:52:11.000000
tags: 
- Software
- gpg
---


#### Cheatsheet

This applies only to gpg version 1

encrypt

    gpg -c lala.tar.gz

decrypt

    gpg --output lulu.tar.gz --decrypt lala.tar.gz.gpg

encrypt with mail address as identifier

    gpg --encrypt --recipient 'myfriend@his.isp.net' foo.txt

get key from keyserver

    gpg --search-keys 'user@example.com'

search keys locally

    gpg --list-keys user

update all keys

    gpg --refresh-keys

delete keys

    gpg --delete-keys

#### Key Regeneration

* fuer entropy sorgen.
* backup vom alten zeug machen
* havedge installieren

		aptitude install havedge

* neuen key generieren

		gpg --gen-key

* signaturen updaten

		gpg --list-secret-keys


* neuen key mit altem key signieren um signaturen zu behalten

		gpg --default-key D2C909A2 --sign-key CDA4B775

* alten key revoken

		gpg --gen-revoke D2C909A2
		gpg --import revoke.txt

		sec  2048R/D2C909A2 2010-06-25 Florian Baumann <wa1@noqqe.de>

		Create a revocation certificate for this key? (y/N) y
		Please select the reason for the revocation:
		  0 = No reason specified
		  1 = Key has been compromised
		  2 = Key is superseded
		  3 = Key is no longer used
		  Q = Cancel
		(Probably you want to select 1 here)
		Your decision? 2
		Enter an optional description; end it with an empty line:
		> replaced by CDA4B775
		>
		Reason for revocation: Key is superseded
		replaced by CDA4B775
		Is this okay? (y/N) y

		You need a passphrase to unlock the secret key for
		user: "Florian Baumann <wa1@noqqe.de>"
		2048-bit RSA key, ID D2C909A2, created 2010-06-25

		ASCII armored output forced.
		Revocation certificate created.

		Please move it to a medium which you can hide away; if Mallory gets
		access to this certificate he can use it to make your key unusable.
		It is smart to print this certificate and store it away, just in case
		your media become unreadable.  But have some caution:  The print system of
		your machine might store the data and make it available to others!


* alten key auf keyserver laden

		gpg --list-key oldkeyid
		gpg --send-keys oldkeyid

* neuen key auf keyserver laden

		gpg --list-key newkeyid
		gpg --send-keys newkeyid

* neue pdfs fuer handout generieren

		aptitude install signing-party ghostscript
		gpg-key2ps -p a4 CDA4B775 > key.ps
		ps2pdf key.ps

#### Links

http://www.madboa.com/geek/gpg-quickstart/
