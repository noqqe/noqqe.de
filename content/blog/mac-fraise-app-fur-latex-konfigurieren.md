---
aliases:
- /blog/2010/12/23/mac-fraise-app-fur-latex-konfigurieren
- /archives/1437
comments:
- author: sbo
  content: "<p>Netter Tipp.<br>Allerdings wurde Fraise eingestellt, und Smultron wird
    aktiv weiter entwickelt. Das ist f\xFCr 3,99 im AppStore zu haben.</p>"
  date: '2011-01-25T04:10:22'
- author: noqqe
  content: <p>WTF BBQ ? </p><p>Hat da jemand Kohle gerochen, oder wie darf ich das
    verstehen?</p>
  date: '2011-01-25T22:31:58'
- author: sbo
  content: "<p>Das wei\xDF ich nicht, aber es ist ja quasi Fraise, nur dass noch Updates
    kommen. Auf der Fraise Seite k\xFCndigt der Dev ja an, an einem komplett neuen
    Editor zu werkeln..und der Post ist von Juni 2010. Und ich find die ~4\u20AC f\xFCr
    Smultron eigentlich ganz okay, die Auswahl ist bei OSX ja leider eher sp\xE4rlich.
    TextMate ist mir mit 30\u20AC zu teuer, TextWrangler ist irgendwie \xFCberladen
    und jEdit ist langsam. Also lieber 4\u20AC f\xFCr Smultron und gut ist. Au\xDFer
    du nimmst halt vi oder emacs, aber..</p>"
  date: '2011-01-28T16:58:07'
- author: sbo
  content: "<p>Ist es eigentlich Absicht, dass die Ausgabe von pdflatex an die Stelle
    des Cursors im LaTeX Dokument eingef\xFCgt wird? :)</p>"
  date: '2011-01-30T17:57:50'
- author: noqqe
  content: "<p>Whupps. Der Haken f\xFCr 'INLINE' geh\xF6rt da nicht hin, der da auf
    dem Screenshot zu sehen ist. </p><p>Wenn der raus ist, passier das nicht :)</p>"
  date: '2011-01-30T18:30:36'
date: '2010-12-23T14:16:03'
tags:
- development
- osx
- smultron
- latex
- fraise
- apple
- web
- app
- "tastenk\xFCrzel"
- shortcut
- mac
- r
- build
- bash
title: "Mac | Fraise App f\xFCr LaTeX konfigurieren"
---

Nachdem meine 30 Tage von [TextMate App](http://macromates.com/)
ausgelaufen sind, habe ich mich irgendwie an das Apfel+R für LaTeX setzen,
durchbauen und PDF anzeigen gewohnt.

Allerdings hab ich keine 50$ auf der hohen Kante um dir TextMate zu kaufen.
Aber es geht ja auch anders. Was früher
[Smultron](http://en.wikipedia.org/wiki/Smultron) hieß und mittlerweile vom
eigentlichen Entwickler eingestellt wurde, gibt es heute als Fork unter dem
Namen Fraise.

[Fraise](https://github.com/jfmoy/Fraise/) ist OpenSource und eigentlich
ein schöner schmaler Editor für Mac OS X.  Meine vermisste Funktion des
Apfel+R lässt sich auch hier "nachbilden". Mit Apfel+B werden alle
Kommandos angezeigt die mit Apfel+R verfügbar sind. Standardmäßig gibt es
dort auch einen Satz von Dummies und sonstigem Nützlichem.

{{< figure src="/uploads/2010/12/Screen-shot-2010-12-23-at-15.09.171.png" >}}

Neue Collection namens LaTeX angelegt und das Command-Set "Build LaTeX"
eingerichtet. Alles weitere ist im Endeffekt nur ein Shell-Skript, welches
das PDF erstellt, die überflüssigen Dateien aufräumt und das generierte pdf
mittels "open" öffnet.

Extrem hilfreich und billiger als TextMate. Das Skript habe ich bei github
gepasted.