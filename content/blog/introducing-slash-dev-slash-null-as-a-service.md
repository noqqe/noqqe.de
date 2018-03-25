---
aliases:
- /blog/2013/10/29/introducing-slash-dev-slash-null-as-a-service
comments:
- author: Georg
  content: "<p>Jetzt noch eine Anbindung an ifttt bitte, damit ich absichtlich bei
    anderen Cloud-Diensten abgelegte Dateien unabsichtlich l\xF6schen kann, und ich
    bin dabei! ;)</p>"
  date: '2013-10-29T21:02:59'
- author: Chris Travers
  content: <p>Hi.  I tried to email you over this but mail@devnull<a href="http://-as-a-service.com"
    rel="nofollow">-as-a-service.com</a> is apparently not working.  I found a bug
    in your setup.  I figure it's probably well known so I will just report it here.</p><p>I
    wanted to let you know that the sample code yields a 404 error, that <a href="http://denull-as-a-service.com/dev/null"
    rel="nofollow">http://denull-as-a-service.com/dev/null</a> is not found.</p><p>My recommendation
    is that you create a simple FCGI script to simply return an HTTP 202 ACCPETED
    code.  In general it looks bad if your demo returns a file not found.</p><p>Best
    of luck on your new venture.  Might I suggest however that /dev/zero as a service
    might be the next big thing?  /dev/random might have limited applicability.  With
    /dev/zero, on the other hand, it would make it easier to say erase hard drives
    by piping the output of wget onto your hard drive.  And moreover, you'd be able
    to likely charge quite a bit more for the same number of customers, since you'd
    need to charge to cover your bandwidth.</p>
  date: '2013-10-30T01:11:56'
- author: noqqe
  content: <p>I think you got a typo in there - you wrote "denull" :) I tried once
    again and with the correct url it works fine :) </p><p>Thanks for your wishes
    :)</p>
  date: '2013-10-30T05:10:57'
- author: exceptionznet
  content: "<p>Wie w\xE4re es, eine Gruppe an Projekten in die Unterst\xFCtzerliste
    der DNAAS mit aufzunehmen. Exceptionz Project w\xE4re da gerne mit dabei :). (So
    im Sinne von \"Investoren\", \"Partner\" etc. bei SAAS-Plattformen)...</p>"
  date: '2013-10-30T18:28:51'
- author: A concerned customer
  content: <p>The service seems great!<br>I have a question. I usually run out of
    RAM , but I don't want to buy more sticks. Also, I don't want to waste space on
    my precious disks to store the pages that don't fit anymore. Could you offer swap
    as a service? thanks!</p>
  date: '2013-10-31T12:34:08'
date: '2013-10-29T17:12:00'
tags:
- web
- stats
- webservice
- business
- devops
- enterprise
- devnull
- cloud
title: Introducing /dev/null as a Service
---

[devnull-as-a-service.com](http://devnull-as-a-service.com)
