---
aliases:
- /archives/317
comments:
- author: Dominik
  content: "<p>Du hast es so gewollt... ;-)</p><p>Mir war nicht ganz klar, wie tief
    das Ganze gehen soll. Daher sind meine Kommentare an vielen Stellen wahrscheinlich
    schon zu detailliert.</p><p>1. Urspr\xFCnglich hie\xDF RAID \"Redundant Array
    of Inexpensive Disks\", daher machten Concat und RAID0 durchaus Sinn. Denn sie
    erlaubten preiswert gr\xF6\xDFere und schnellere Plattenspeicher zur Verf\xFCgung
    zu stellen. Erst im Laufe der Entwicklung wurde \"Inexpensive\" zu \"Independent\",
    da die meisten RAID-L\xF6sungen alles andere als billig waren.</p><p>2. Concats
    (NRAID o. Linear Mode) hast Du \xFCbrigens so gut wie gar nicht erw\xE4hnt. Dabei
    spielen die in der Praxis eine nicht unerhebliche Rolle. Es ist n\xE4mlich einfacher
    einen RAID-Verbund durch \"Concatenation\" zu vergr\xF6\xDFern, als durch Reorganisation.
    Letzteres geht meist nicht ohne Ausfallzeiten und das ist im Hochverf\xFCgbarkeitsbetrieb
    (und da wird immer mit RAID gearbeitet!)<br>nicht akzeptabel.</p><p>3. RAID5 kann
    durchaus auch effizient in Software implementiert werden. Das Hauptproblem bei
    allen RAID-Level 3-5 ist nicht die Berechnung der Parit\xE4t: Bei 5 Platten muss
    ich, auch wenn ich nur ein Byte schreiben will, jeweils einen vollen Block von
    4 Platten lesen und je einen Block auf zwei Platten schreiben! D.h. nicht das
    Berechnen der Parit\xE4t ist das Problem sondern das Lesen und Schreiben. Das
    kann ein Software-RAID genauso schnell wie ein Hardware-RAID, es sei denn Du f\xFChrst
    Caches ein. Das hat aber Auswirkungen auf die Datensicherheit.</p><p>4. Die Beschreibung
    von RAID3 ist nicht sehr hilfreich. Warum ist RAID3 verschwunden? Stichwort: Hotspot,
    Bottleneck. Das f\xFChrt automatisch zur Betrachtung RAID4 und RAID5. RAID4 wird
    \xFCbrigens von NetApp verwendet, ist also nicht unbedingt ein Nischenprodukt.</p><p>5.
    Was bei deiner Betrachtung zu RAID1 fehlt sind die m\xF6glichen Lesestrategien,
    man kennt da z.B. first (mirror), round-robin und geometric. Warum muss man beim
    Schreiben immer auf das OK aller Subspiegel warten? Was hat das f\xFCr Auswirkungen
    auf die Performance? Was ist mit Mehrfach-Spiegeln?</p><p>6. Bei der Betrachtung
    von RAID0+1 gegen\xFCber RAID10 fehlt mir die Frage, was denn genau passiert,
    wenn zwei beliebige Platten ausfallen. Du hast das f\xFCr RAID10 gut erkl\xE4rt.
    Wenn Du es f\xFCr RAID0+1 genauso erkl\xE4rst, dann wird relativ schnell klar,
    dass RAID0+1 nicht sonderlich sinnvoll ist.</p><p>7. Hier noch eine ganzer Schwung
    von weiterf\xFChrenden Fragen:<br>- Spare-Disks<br>- Recovery von RAIDs (Inkompatible
    On-Disk-Formate!)<br>- Der Unterschied zw. Spiegelung oder Replikation<br>- Wogegen
    sch\xFCtzt RAIDn (n&gt;0) und wogegen ein Backup?<br>- Datenintegrit\xE4t (\"silent
    data corruption\" -&gt; Sun ZFS)<br>- Anwendung des RAID-Gedanken in anderen Bereichen
    (RAIT, Mirrored RAM)</p>"
  date: '2008-11-10T03:30:08'
- author: noqqe
  content: "<p>Hallo Dominik und danke f\xFCr deinen ausf\xFChrlichen Kommentar!<br>Das
    hilft mir wirklich sehr! Nur wie du oben sch\xF6n erw\xE4hnt hast konntest du
    ja nicht wissen, wie tief das ganze geht.</p><p>Stand der Dinge ist, dass das
    Thema in 15 Minuten Vortrag vom Tisch sein soll und das Handout eine Art \xDCbersicht
    dazu darstellt. Also keine Doktor oder Facharbeit. Stell dir jetzt noch vor das
    ich (noch) in einer IT-System-KAUFMANNs Klasse referiere.</p><p>Trotz allem m\xF6chte
    ich dich deiner M\xFChe entsch\xE4digen und nehme kurz Stellung zu den, von dir
    genannten, Punkten.</p><p>1. Der Fakt ist mir durchaus bewusst, aber nicht relevant
    f\xFCr meine Zuh\xF6rer. Der Faktor des Speichervergr\xF6\xDFerns w\xE4re aber
    einbaubar.</p><p>2.Nicht Pr\xFCfungsrelevant. Insiderwissen und spezielle Anwendungsbereiche
    sind nicht wirklich f\xFCrs Handout geeignet. Wie gesagt. Kaufmannsklasse=Oberfl\xE4chenwissen.
    Oder willst du mir widersprechen das die verbreitetsten Arten von RAID 0 1 und
    5 sind? :)</p><p>3. Mh dar\xFCber l\xE4sst sich nun streiten. Sende und Lese Arbeit
    d\xFCrfte nicht missachtet werden. Jedenfalls bin ich der Meinung. (\xFCbrigens
    auch die aktuelle Ausgabe der C'T stimmt MIR hier zu :) )</p><p>4. Selbstverst\xE4ndlich.
    Muss aber troztdem gestehen das ich meinen Zuh\xF6rern die ableitung des nachteils
    an RAID3 durchaus zutraue. EVTL werde ich es aber einbauen.</p><p>5. Siehe 2.</p><p>6.
    Muss ich dir zustimmen.</p><p>7. Siehe 2.</p><p>Danke aufjedenfall! :)</p>"
  date: '2008-11-10T20:35:43'
date: '2008-11-09T18:54:14'
slug: projektarbeit-raid-handout
tags:
- hardware
- projektarbeit
- raid
- referat
- berufsschule
title: 'Projektarbeit: RAID-Handout'
---

Nächste Woche ist schon Abgabe: Hier mal der Release Candidate :P
Wenn Fehler oder ähnliches auffallen bitte Bescheid sagen :) Wirkt sich ja
alles auf meine Note aus :P

Anzuschauen unter:
[http://zwetschge.org/publications/RAID-Handout.pdf](http://zwetschge.org/publications/RAID-Handout.pdf)