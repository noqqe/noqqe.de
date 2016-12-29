---
comments:
- author: gewalt-achim
  content: "<p>dein blog is nich valid: <a href=\"http://validator.w3.org/check?uri=http%3A%2F%2Fnoqqe.de\"
    rel=\"nofollow\">validator.w3.org/check?uri=htt...</a></p><p>und das is auch der
    einzige blog, der im planet auf der uu-page die \"zur\xFCck\"- und \"weiter\"-buttons
    nich sauber anzeigt (hab hier FF 14.0.1 @ xubuntu 12.04 lts).</p>"
  date: '2012-07-29T14:48:33'
- author: noqqe
  content: "<p>Danke f\xFCr den Hinweis. Der Hersteller meiner Octopress Blogging
    Engine wird das bestimmt bald richten. <br>Ansonsten... Ja. Schade, dass der w3
    Validator meinen Blog nicht toll findet. Tats\xE4chlich ist mir das ein bisschen
    egal. </p><p>Was hab ich mit den Zur\xFCck/Weiter Buttons vom Uu-Planeten zu tun?
    Die liefert doch nicht mein Octopress aus. <br></p>"
  date: '2012-07-29T15:16:50'
- author: gewalt-achim
  content: "<p>das k\xF6nnte in sofern relevant sein, als dass deine planet-beitr\xE4ge
    als einzige ein p-element in den \xFCberschriften des planets verursachen und
    das scheint zu anzeigeproblemen f\xFChren. im \xFCbrigen sind innerhalb von a-tags
    (ergo links) nur inline-elemente erlaubt und da z\xE4hlt p halt nicht dazu. von
    daher ist eine valide syntax im sinne der barrierefreiheit immer interessant,
    ganz unabh\xE4ngig davon, wie sehr dich das tangiert :)</p><p>der validator findet
    deinen blog sicher super, nur der code is halt der letzte dreck.</p><p>damit du
    nachvollziehen kannst was ich meine, hier ein kleiner auszug von der planet-website:</p><p>---------------------------------------------------------<br>grep
    -A1 '&lt;h3' 'noqqe.de=\"\" -a3=\"\" -b4=\"\" 06=\"\" 16'=\"\" 2012=\"\" &lt;h3=\"\"
    blog=\"\" foo=\"\" grep=\"\" |=\"\"&gt;<br>                  <a href=\"http://www.intux.de/?p=2506\"
    rel=\"nofollow\">Gravatar?</a><br>--<br>                &lt;h3&gt;<br>                  <a
    href=\"http://noqqe.de/blog/2012/06/16/stromverbrauch-fur-nerds-mit-google-charts/\"
    rel=\"nofollow\"></a></p><p><a href=\"http://noqqe.de/blog/2012/06/16/stromverbrauch-fur-nerds-mit-google-charts/\"
    rel=\"nofollow\">Stromverbrauch f\xFCr Nerds mit Google Charts</a></p><br>--<br>
    \               &lt;/h3&gt;&lt;h3&gt;<br>                  <a href=\"http://www.soeren-hentzschel.at/mozilla/firefox/2012/06/16/mozilla-veroffentlicht-firefox-13-0-1/\"
    rel=\"nofollow\">Mozilla ver\xF6ffentlicht Firefox 13.0.1</a><br>---------------------------------------------------------<p></p><p>das
    kriegst du von mir halt direkt an die backe, weil dein blog der einzige is, der
    das problem hat und du hier mit invalid code um die ecke kommst. solang das nich
    korrigiert ist, brauch man letztlich auch nich bei uu suchen, w\xFCrd ich mal
    meinen.<br>&lt;/h3&gt;&lt;/h3'&gt;</p>"
  date: '2012-07-29T20:13:23'
- author: gewalt-achim
  content: "<p>lol</p><p>na gro\xDFartig, dein disqus-plugin kommt mitm code nich
    klar ... eigentlich auch nich \xFCberraschend, wenn du einmal in die shice greifst
    ...</p>"
  date: '2012-07-29T20:14:36'
- author: noqqe
  content: "<p>Ich sags dir mal einfach so. Das ist die URL die Ubuntuusers bei mir
    abholt. <a href=\"http://noqqe.de/ubuntuusers.xml\" rel=\"nofollow\">http://noqqe.de/ubuntuusers.xm...</a></p><p>Nennt
    sich Atom Feed. </p><p>Wo der Uu-Planet  irgendwelche &lt; p&gt; Tags findet und
    warum du hier total ausrastet ist mir schleierhaft. </p><p>Was ich weiss ist,
    dass ich wohl nicht der einzige mit einem Octopress Blog im Uu-Planeten bin und
    es au\xDFerdem einen Bug im Inoyka Bugtracker dazu gibt. Aus dem Uu-Web-Team ist
    noch keiner an mich heran getreten, \xC4nderungen zu unternehmen. </p>"
  date: '2012-07-30T15:34:47'
- author: noqqe
  content: "<p>Ich sags dir mal einfach so. Das ist die URL die Ubuntuusers bei mir
    abholt. <a href=\"http://noqqe.de/ubuntuusers.xml\" rel=\"nofollow\">http://noqqe.de/ubuntuusers.xm...</a>
    Nennt sich Atom Feed. Wo der Uu-Planet irgendwelche &lt; p&gt; Tags findet und
    warum du hier total ausrastet ist mir schleierhaft. Was ich weiss ist, dass ich
    wohl nicht der einzige mit einem Octopress Blog im Uu-Planeten bin und es au\xDFerdem
    einen Bug im Inoyka Bugtracker dazu gibt. Aus dem Uu-Web-Team ist noch keiner
    an mich heran getreten, \xC4nderungen zu unternehmen....</p>"
  date: '2012-07-30T15:35:36'
- author: gewalt-achim
  content: "<p>na nu werden wir aber doch nich gleich ausfallend, nur weil mal einer
    helfen will oder? wenn das hier f\xFCr dich bereits \"total ausrasten\" ist, dann
    biste aber sehr d\xFCnnh\xE4utig. unterdessen wei\xDF ich worans liegt und w\xFCrd
    an deiner stelle mal das xml genauer angucken, aber scheint dich ja nich weiter
    zu interessieren. bleibts halt buggy...</p>"
  date: '2012-07-30T22:10:33'
- author: noqqe
  content: "<p>Okay. Wenn du m\xF6chtest schreib mir das doch mal per Mail. </p><p>Bist
    hier tats\xE4chlich der erste der wohl wei\xDF an was es liegt. </p><p>Findet
    sich hier immer wieder in den Kommentaren\"du bist doof weil du kein WP hast,
    das im UU Planeten f\xFCr aussieht.</p><p>Egal. Ich werd den Thread nachher hier
    entfernen weil er mit dem post so garnix zu tun hat. (wie die ganzen anderen Kommentare
    zu dem Thema auch. ) </p>"
  date: '2012-07-31T04:43:26'
date: '2012-07-02T19:00:00'
tags:
- web
- bars
- shell
- stats
- gnuplot
- plot
- barchart
- linechart
- sport
- runkeeper
- debian
title: Paying 5 bucks a month for stupid statistics?
---

Seit ich angefangen habe Ausflüge durch Wälder und anderem grobem Gelände mit
meinem Mountainbike zu machen ist [Runkeeper](http://runkeeper.com) mein Begleiter.

Mit der iPhone App kann ich allerhand Daten tracken wie

* Durchschnittliche Geschwindigkeit
* Höhenmeter
* Kalorien (wenn man deren Berechnungen glauben mag)
* Karten
* Länge der Fahrt

Natürlich will auch Runkeeper irgendwie Geld verdienen und
somit gibt es was Accounts angeht auch noch einen "Elite Pro Wahnsinns" Account
welcher irgendwie 5-8 Euro im Monat kostet und mir als Ausgleich bessere
Statistiken und sogenannte "Fitness Reports" liefert.

Ehrlichgesagt war ich davon nicht sehr angetan so viel Geld für ein paar extra
Skripte zu bezahlen. Deshalb hab ich mich entschlossen meine eigenen Graphen mit
[GnuPlot](http://gnuplot.org) zu malen und irgendwie zu einer Website zu machen.

## runkeeper-statistics

Im Grunde liefert mir Runkeeper alle Rohdaten die ich brauche. Sie weigern sich
nur mir die Statistiken daraus zu bauen die gerne hätte. Ich habe mir dann ein
Shell Script gebaut, welches gnuplot mit Daten füttert, die in einem CSV File
spezifiziert sind.

{{< figure src="/uploads/2012/07/calories-activity.png" >}}

Zuerst schreibe ich alle meine Daten in die activity.dat. Ziemlich im CSV Stil,
aber mir fiel nicht blöderes ein ohne SQL sprechen zu müssen.

``` csv
# Add your data here to generate awesome graphs ;)
# ID, Date,      Distance, Duration, Pace, Speed, Burned, Climb
  1,  2012-06-23, 14.07,    0.54,     3.54, 15.38, 458,    156
  2,  2012-06-26, 16.28,    1.09,     4.17, 14.03, 582,    292
  3,  2012-06-28, 17.65,    1.13,     4.11, 14.36, 618,    242
  4,  2012-06-30, 27.28,    1.47,     3.56, 15.24, 876,    379
```

Danach baut mir das Shellscript die GnuPlots und die HTML Site

```
$ ./runkeeper-statistics
```

Das alles gibts auf [github](http://github.com/noqqe/runkeeper-statistics/) und
die "Demo" hab ich mal hochgeladen und mit zufälligen Daten befüllt:<b>[Demo](/uploads/2012/07/runkeeper-statistics/html/index.html)</b>