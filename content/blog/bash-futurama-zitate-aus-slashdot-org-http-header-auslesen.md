---
aliases:
- /archives/1255
- /blog/2010/09/26/bash-futurama-zitate-aus-slashdot-org-http-header-auslesen
comments:
- author: noqqe
  content: "<p>Update1: bei Commandlinefu gibt es ebenfalls eine Reihe von L\xF6sungen
    f\xFCr die HTTP-Header Abfrage von <a href=\"http://slashdot.org\" rel=\"nofollow\">slashdot.org</a>:
    <a href=\"http://www.commandlinefu.com/commands/tagged/234/futurama\" rel=\"nofollow\">http://www.commandlinefu.com/c...</a></p><p>Update2:
    Es scheint auch ein fortune-mod-paket zu geben, welches l\xE4ngere Zitate von
    Futurama enth\xE4lt: <a href=\"http://www.netmeister.org/apps/fortune-mod-futurama-0.2.tar.gz\"
    rel=\"nofollow\">http://www.netmeister.org/apps...</a></p>"
  date: '2010-09-26T17:08:49'
- author: The Compiler
  content: "<p>if [ $(grep \"$quote\" $target |wc -l) -lt 1 ]; then</p><p>Das ist
    b\xF6\xF6\xF6se :p Geht doch viel einfacher:</p><p>if ! grep -q \"$quote\" \"$target\";
    then</p>"
  date: '2010-09-26T18:28:42'
- author: noqqe
  content: "<p>@the compiler: <br>ah sehr nice. quiet und nur den r\xFCckgabewert
    benutzen. awesome.</p>"
  date: '2010-09-26T19:05:35'
- author: Christian D.
  content: "<p>@toadie.de:</p><p>Unschlagbar! Wer braucht schon Cowsay wenn das Bender
    \xFCbernehmen kann? ;-)))</p>"
  date: '2010-09-29T19:46:17'
date: '2010-09-26T12:56:42'
tags:
- development
- shell
- motd
- script
- welcome
- blog
- futurama
- slashdot
- oneline
- ubuntu
- message of the day
- debian
- bash
title: Bash | Futurama Zitate aus slashdot.org HTTP-Header auslesen
---

Bin heute morgen über ein Easter-Egg von Slashdot.org
[gestolpert](http://www.eastereggs.svensoltmann.de/content/view/686/26/).
HTTP-Header:

```
$ curl -Is slashdot.org
HTTP/1.1 200 OK
Server: Apache/1.3.41 (Unix) mod_perl/1.31-rc4
SLASH_LOG_DATA: shtml
X-Powered-By: Slash 2.005001
**X-Fry: You'll barely regret this.**
X-XRDS-Location: http://slashdot.org/slashdot.xrds
Cache-Control: no-cache
[...]
```

Bei so ziemlich jeder Anfrage steht an der Stelle ein neues Zitat. Da ich
sowieso total [auf Futurama stehe](/archives/995), dachte ich mir ich baue
die Zitate als Welcome-Message in meine Rechner ein:

```
$ curl -Is slashdot.org | sed -n '5p' | sed 's/^X-//'
Bender: OK, but I don't want anyone thinking we're robosexuals.
$ curl -Is slashdot.org | sed -n '5p' | sed 's/^X-//'
Fry: I can burp the alphabet. A, B, D ... no, wait ...
```

Wenn ich aber bei allen meinen Rechnern die Zeile einbinde, hat das
irgendwie ein bisschen was von DOS-Attacke. Muss ja nicht sein. Mit einem
Einzeiler hab ich mir die Quotes erstmal alle besorgt:


    target="/path/to/file/018" ; while true ; do quote="$(curl -Is slashdot.org |sed -n '5p' |sed 's/^X-//')" ; if [ $(grep "$quote" $target |wc -l) -lt 1 ]; then echo $quote >> $target ; echo $quote ; sleep 1 ; fi ; done

Hier gibt's alle Quotes die der Einzeiler bis jetzt gesammelt hat:
[/uploads/2009/09/018](/uploads/2009/09/018)

Futurama Quotes beim Login unter Ubuntu:

``` bash
$ wget /uploads/2009/09/018 -O  ~/.futurama
$ vi ~/.bashrc
quotes="$HOME/.futurama"
if [[ $- == *i* ]]; then
echo " "
rnd=$((RANDOM % $(cat $quotes | wc -l)+3)) ; sed -n "${rnd}p" $quotes
fi
```

Einloggen oder Shell starten sieht dann wie folgt aus:

```
$ ssh user@host
user@host password:
Linux host 2.6.32-22-generic #33-Ubuntu SMP Wed Apr 28 13:27:30 UTC 2010 i686 GNU/Linux
Ubuntu 10.04 LTS
[...]
You have new mail.
Bender: You can't count on God for jack! He pretty much told me so himself.
user@host:~$
```