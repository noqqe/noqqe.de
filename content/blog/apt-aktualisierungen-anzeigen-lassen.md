---
aliases:
- /archives/894
- /blog/2010/02/20/apt-aktualisierungen-anzeigen-lassen
comments:
- author: killermoehre
  content: "<p>Eventuell w\xE4re noch eine Notification per 'notify-send' und/oder
    ein Tray-Icon mit 'zenity --notification' m\xF6glich.</p><p>killermoehre</p>"
  date: '2010-02-21T00:59:24'
- author: Benjamin
  content: "<p>Mit python-apt sollte diese Abfrage auch m\xF6glich sein, wenn man
    sein Script in Python schreiben m\xF6chte.</p>"
  date: '2010-02-21T01:10:25'
- author: noqqe
  content: "<p>@killermoehre ja dar\xFCber hab ich auch schon nachgedacht. aber ob
    ich das bei 200 maschinen haben will ?</p><p>@benjamin: ja mit sicherheit. aber
    das w\xE4re bisschen overkill.</p>"
  date: '2010-02-21T01:51:57'
- author: Guest
  content: '<p>@Benjamin: Deine Aussage ist ca. so sinnvoll wie die folgende: "Wenn
    du das alles ganz anders machst, kannst du das alles auch ganz anders machen."</p>'
  date: '2010-02-21T01:55:00'
- author: killermoehre
  content: "<p>Wenn die Maschinen alle das gleiche laufen haben, dann braucht es doch
    nur auf einer \xFCberpr\xFCft zu werden, weil ja alle auf die selbe Paketquelle
    zugreifen.</p><p>killermoehre</p>"
  date: '2010-02-21T02:00:04'
- author: noqqe
  content: "<p>Im Grunde hast du recht. Aber ich m\xF6chte dich hier weder mit Firmeninterna
    Sachen aufkl\xE4ren, noch das restliche UpgradeSystem erl\xE4utern. Hauptgrund
    ist aber, dass mehrere Leute upgraden.</p><p>Das eigentliche Script, ist auch
    ganz anders als du dir das gerade vorstellst :)</p>"
  date: '2010-02-21T15:01:50'
- author: adun
  content: "<p>N\xF6, \"linux\" ist keine gute Wahl, gibt bestimmt hundert Paketnamen,
    die das auch enthalten und kein Kernel sind. Ich w\xFCrde \"aptitude search '~U
    ~m\"Ubuntu Kernel Team\" linux'\" empfehlen. Mit -F kannst du das Ausgabeformat
    \xE4ndern, nur scheint das unn\xF6tig zu sein, du willst ja nur wissen ob es \xFCberhaupt
    eine Ausgabe gab. Du musst also nur pr\xFCfen ob etwas nach StdOut geschrieben
    wird und brauchst die Pipe nicht auszupacken.</p>"
  date: '2010-02-22T19:20:09'
date: '2010-02-20T16:27:38'
tags:
- development
- kernel
- search
- aktualisierungen
- anzeigen
- show
- list
- update
- apt
- upgrade
- updates
- aptitude
- linux
- apt-get
- debian
- ubuntu
title: Apt | Aktualisierungen anzeigen lassen
---

Für ein kleines Script, welches ich auf der Arbeit verwende, habe ich
versucht einen Weg zu finden, zu prüfen ob Kernel-Updates verfügbar sind. Im
Netz und in der Man-Page von apt-get bzw aptitude wurde ich nicht eindeutig
fündig. Nach langem suchen ergaben sich allerdings folgende Möglichkeiten
Updates anzeigen zulassen:

```
apt-get --just-print upgrade
```

```
apt-get -s upgrade
```

```
aptitude search ~U
```


Ich fand allerdings die erste Möglichkeit am Einleuchtensten.  Die Ausgabe
ist zwar nicht zwingend Script geeignet, aber das lässt sich ja ändern:

    SUPDATEKERNEL=$(ssh root@$host "apt-get --just-print upgrade | grep linux | uniq | wc -l")

In der Variable $SUPDATEKERNEL steht logischerweise die Anzahl der
Verfügbaren Updates die "linux" enthalten. Linux fand ich persönlich jetzt
am einfachsten zur Identifikation von Kernel-Updates. Jemand bessere
Vorschläge?