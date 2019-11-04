---
comments:
- author: Vain
  content: "\"Ich blogge \xFCber Windows Themen. F\xFChle mich jetzt Enterprise-Ready\xA9.\"\n\nMade
    my day. :-)"
  date: '2013-12-05T21:13:01.647395'
date: '2013-12-05T18:28:00'
tags:
- windows
- rendering
- font
- aliasing
- gdipp
title: Fixing Windows Font Rendering
---

Seit ich auf der Arbeit Windows nutzen muss, störe ich mich am Font der systemweit
irgendwie "ausgefranzt" wirkt. Ausgefranzt ist natürlich ein total professionell
klingendes Wort, man könnte von irgendwas mit Anti-Aliasing sprechen aber dazu
hab ich mich genau nicht schlau gelesen.

Das erste Thema, auf das man stößt wenn man nach Font Rendering unter Microsofts
Betriebssystem schlauließt heisst [ClearType](https://en.wikipedia.org/wiki/ClearType),
doch egal wie viel ich damit herumgespielt hatte, kam dabei nicht heraus was ich
mir vorstellte. Nach wie vor brößliger Font.

{{< figure src="/uploads/2013/12/browser-ohne.png" >}}

Dann fand ich nach etwas längerem Herumgoogeln
[gdipp](https://code.google.com/p/gdipp/). Das Tool stellt, je nach Einstellung
Windows Service, eine DLL oder einen Registry Eintrag zur Verfügung und ersetzt
damit die vom OS bereitgestellte Font-Rendering-Engine. Mittel der Wahl ist die
Installation des Windows Services.

{{< figure src="/uploads/2013/12/browser-mit.png" >}}

Dann sieht auch gleich alles Systemweit schön aus. Outlook und sogar das Cygwin
Terminal.

{{< figure src="/uploads/2013/12/cmd-ohne.png" >}}
{{< figure src="/uploads/2013/12/cmd-mit.png" >}}

Ich blogge über Windows Themen. Fühle mich jetzt Enterprise-Ready&copy;.
