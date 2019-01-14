---
comments:
- author: Anaximander
  content: "<p>Ich bin ebenfalls bei Uberspace. Mehr E-Mail-Spam kann ich bei mir
    nicht nachvollziehen. Im Bogofilter bleibt bei mir nicht mehr h\xE4ngen als sonst
    auch. Spamassassin nutze ich allerdings nicht.</p>"
  date: '2012-12-08T15:57:02'
- author: noqqe
  content: "<p>Bogofilter? Wie kann ich mir das jetzt vorstellen? L\xE4uft das auf
    den Servern schon? Selber installiert?</p>"
  date: '2012-12-08T16:07:35'
- author: Anaximander
  content: "<p>Nein, der l\xE4uft bei mir lokal und klinkt sich in mein E-Mail-Programm
    ein. Bin ehrlich gesagt zu faul das zu \xE4ndern, da die Trefferquote inzwischen
    bei 99,x% liegt und einige meiner anderen E-Mail-Adressen bei Anbietern sind,
    die oft keinen bzw. einen ziemlich schlechten Spamfilter anbieten.</p><p>Ich vermute
    aber, dass Bogofilter durchaus bei Uberspace laufen sollte, da er z. B. mit diversen
    MDA, wie den von <a href=\"http://uberspace.de\" rel=\"nofollow\">uberspace.de</a>
    verwendeten maildrop zusammenarbeitet. M\xFCsste man eigentlich mal ausprobieren...</p>"
  date: '2012-12-08T19:22:31'
- author: noqqe
  content: "<p>Das werd ich mir definitv mal durchschauen. Danke f\xFCr den Tipp!</p>"
  date: '2012-12-09T09:09:46'
- author: noqqe
  content: "<p>btw: Hab jetzt spamassassin durch bogofilter ersetzt. Seither keine
    einzige Spam Mail mehr in meiner Inbox. Alles super erkannt. Trainiere das teil
    jetzt im -u Mode und regelm\xE4\xDFig alle Stunde \xFCber nen Cronjob (wenn ich
    manuell was verschiebe). </p><p>Danke - Super Teil.</p>"
  date: '2012-12-22T11:20:05'
date: '2012-12-08T11:12:00'
tags:
- stats
- spam
- grep
- mail
- uberspace
- bash
- statistik
title: Spam
---

Ich wollte mich noch nie selber um meine Mails kümmern. Rumkonfigurieren,
Spam Listen und Erkennungschemen pflegen. Wuah. Außerdem sind mir meine
Mails viel zu wichtig, als das ich Sie auf meinen Frickelkisten hosten möchte.

{{< figure src="/uploads/2012/12/Nigerian-Prince-400x400.jpg" >}}

[Seit August](/blog/2012/08/17/ich-wechselte-zu-uberspace-dot-de/)
liegen meine Mails bei [Uberspace](https://uberspace.de)
rum. Das ist nun knapp 4 Monate her. Gefühlt bekomme ich in der letzten Zeit
viel mehr Spam. Sich aber von einem rein subjektivem empfinden leiten lassen war noch nie so
mein Ding. Deshalb hab ich mich etwas in der Statistik bemüht.

An was könnte das also liegen?

* Aktuelle Spam Welle?
* War ich unvorsichtig beim Erwähnen meiner Adresse im Netz?
* Am Hosting bei Überspace?

Ich habe mich dann in einer ruhigen Minute hingesetzt und ein paar `grep`s
später kam das folgende Liniendiagramm heraus.

Zugegebenermaßen habe ich kurz an dem Spamassassin gezweifelt, den
Uberspace pflegt. Klar Uberspace bekommt nie so viel Mails wie die Infrastruktur
von 1und1. Das die Erkennungsrate in Abhängigkeit zu der Menge an Spam steht
brauche ich hier sicher keinem erzählen.

Tatsächlich sieht es aber genau anders aus. Faktisch bekomme ich insgesamt mehr
Mails als früher und die Erkennungsrate &#40;in meinem Postfach markiert durch
&#91;SPAM&#93;&#41; ist höher als.. ever.

{{< figure src="/uploads/2012/12/yea-fuck-yea.gif" >}}

Das beantwortet natürlich nicht die Fragen, ob ich in letzter Zeit zu leichtfertig
mit meiner Adresse umgehe oder aktuell härter gespammed wird. Aber den
Spamassassin bzw. Uberspace liegt würde ich jetzt erstmal ausschliessen.
