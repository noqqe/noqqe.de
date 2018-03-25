---
comments:
- author: Alex
  content: "Die klassichen Terminal Befehle die ich immer aktiviere:\n\n\nRemove the
    Dock delay:\ndefaults write com.apple.Dock autohide-delay -float 0\n\_ \_ killall
    Dock\n\_ \_\_\nSet the save dialogue window to expanded by default in all applications:\ndefaults
    write -g NSNavPanelExpandedStateForSaveMode -bool TRUE\n\nDisable the sound effects
    on boot:\nsudo nvram SystemAudioVolume=\" \"\n\nAutomatically quit printer app
    once the print jobs complete:\ndefaults write com.apple.print.PrintingPrefs \"Quit
    When Finished\" -bool true\n\nAvoid creating .DS_Store files on network volumes:\ndefaults
    write com.apple.desktopservices DSDontWriteNetworkStores -bool true\n\nShow the
    ~/Library folder:\nchflags nohidden ~/Library\n\nRemove the animation when hiding/showing
    the Dock:\ndefaults write com.apple.dock autohide-time-modifier -float 0"
  date: '2014-09-25T19:39:10.983594'
- author: "Rainer M\xFCller"
  content: "Hier ein paar n\xFCtzliche Helfer und Tools, die ich benutze. Gr\xF6\xDFtenteils
    kostenpflichtig, ich glaube aber, man kann sie zumindest alle kostenlos testen.
    Und teilweise gibt es f\xFCr dieselbe Funktionalit\xE4t durchaus auch andere (kostenlose)
    Tools, aber diese funktionieren halt f\xFCr mich ganz gut.\n\n\_* <a href=\"http://www.controlplaneapp.com/\">ControlPlane</a>:
    Automatische Aktionen ausf\xFChren bei Kontextwechsel. Zum Beispiel basierend
    auf WLAN-Netzwerk, Seriennummer des angeschlossenen Netzteils usw. kann man den
    Ton ausschalten oder Netzwerk-Einstellungen statisch/DHCP \xE4ndern.\n\n\_* <a
    href=\" http://hyperdock.bahoom.com\">HyperDock</a> und <a href=\"http://bahoom.com/hyperswitch\">HyperSwitch</a>:
    Window Management (Manual Tiling), Vorschaubildchen von Fenstern in Cmd+Tab und
    bei Mouse Over im Dock\n\n\_* <a href=\"http://www.alfredapp.com/\">Alfred</a>:
    Apps/Dokumente schnell finden und starten/\xF6ffnen\n\n\_* <a href=\"https://pqrs.org/osx/karabiner/index.html.en\">Karabiner/Seil</a>:
    Tastenbelegung \xE4ndern, unter anderem Caps Lock \xE4ndern und einen <a href=\"http://lolengine.net/blog/2012/06/17/compose-key-on-os-x\">echten
    Compose Key</a> einrichten\n\n\_* <a href=\"http://bjango.com/mac/istatmenus/\">iStat
    Menus</a>: Mehr Informationen in der Men\xFCleiste anzeigen"
  date: '2014-09-25T19:59:52.010134'
- author: Matze
  content: "* caps lock schalt ich grunds\xE4tzlich aus (systemeinstellungen > tastatur
    > sondertasten)\n* secrets prefpane ist sehr praktisch:\_http://secrets.blacktree.com/\n*
    kennst schon\_http://vimr.org/ ?\n* fenster mit tastatur bewegen: http://manytricks.com/moom/\n*\_quarant\xE4ne
    f\xFCr neue dateien abschalten: defaults write com.apple.LaunchServices LSQuarantine
    -bool NO"
  date: '2014-09-25T20:13:10.326839'
- author: F30
  content: "Puhh \u2013 Konfigurationstipps h\xE4tte ich bestimmt auch ein paar, aber
    da f\xE4llt mir grade nichts Konkretes ein.\n\nDeshalb nur kurz zwei Utilities,
    die ich \xE4u\xDFerst n\xFCtzlich finde: Zum Coden ist auf jeden Fall [Dash](http://kapeli.com/dash)
    sehr praktisch und um den \xDCberblick \xFCber die System-Auslastung zu behalten,
    kann ich nicht ohne das von Raim schon erw\xE4hnte [iStat Menus](http://bjango.com/mac/istatmenus/)
    leben."
  date: '2014-09-26T01:12:24.112297'
- author: spion
  content: "Eine viel wichtigere Frage, die sich mir stellt:\nBekomme ich irgendwo
    Key-Caps f\xFCr die Tasten, um die Dreifach-Belegung auf der Tastatur zu haben.
    Ich suche mir immer den Wolf f\xFCr z.B. '~' und '\\'."
  date: '2014-09-26T13:16:03.757029'
- author: spion
  content: "AccessMenuBarApps, um an die \xFCberdeckten Men\xFCApps zu kommen: http://www.ortisoft.de/accessmenubarapps/\n\nSmartReporter
    (Lite):\_http://www.corecode.at/smartreporter_lite/index.html \_\n\n\nClipMenu,
    praktische Zwischenablage:\_http://www.clipmenu.com/"
  date: '2014-09-26T13:35:31.600886'
- author: Anonymous
  content: Ich kann mir os x nicht ohne http://qsapp.com vorstellen.
  date: '2014-10-17T10:11:02.837883'
- author: Matze
  content: 'Da ich mich gerade mit Yosemite rumschlage:

    Systemeinstellungen > Tastatur > Kurzbefehle > Alle Steuerungen (unten links)'
  date: '2014-10-19T16:33:37.246922'
- author: noqqe
  content: "http://lifehacker.com/quickly-change-os-xs-default-screenshot-format-and-loc-1489014578
    - das ist auch n\xFCtzlich"
  date: '2014-10-23T14:15:25.232628'
- author: noqqe
  content: und ausserdem ein riesen Danke an alle Tipps in den Kommentaren! Hab sehr
    viel davon genutzt! <3
  date: '2014-10-23T14:20:48.961478'
date: '2014-09-25T16:34:00'
tags:
- locate
- osx
- openbsd
- apple
- german
- keyboard
- layout
- macbook
title: MacBook
---

Nachdem ich die letzte Zeit viele Betriebssysteme auf dem Thinkpad ausprobiert
haben (FreeBSD, ArchLinux, ElementaryOS, OpenBSD) und nichts mich wirklich
befriedigt hat, bin ich wieder zu Debian zurück.

{{< figure src="/uploads/2014/09/moe.gif" >}}

Einige Zeit später bekam ich in der Arbeit in MacBook. Fand wieder gefallen
daran. Kaufte MacBook Pro (14. Generation).

Dinge, dich ich bei beiden MacBooks als erstes verändert habe:

* Dock auf die linke Seite konfiguriert (in der Breite hab ich mehr Platz als in
  der Länge)
* `Key Repeat` und `Delay until Repeat` Rate für das Keyboard bis zum Anschlag
  verkürzt. Aber darüber hab ich bereits geflamed.
* FileVault aktivieren.
* Natural Scrolling deaktivieren (srsly, was soll das?)
* `German - Eliminate Dead Keys` Keyboard [Layout installieren](https://github.com/sebroeder/osx-keyboard-layout-german-no-deadkeys)
* Zwischen Fenstern der selben Applikation wechseln mit `Apfel` + `>`
* [caffeeine](http://lightheadsw.com/caffeine/) installieren (`caffeinate` für
  die Kommandozeile)
* Homebrew/MacPorts installieren
* [xquartz](http://xquartz.macosforge.org/landing/) installieren
* Hässlichen Terminal-Beep in iTerm2 oder TerminalApp [deaktivieren](http://superuser.com/questions/163994/how-can-i-get-the-mac-terminal-to-not-beep)
* `locate` [aktivieren](http://osxdaily.com/2011/11/02/enable-and-use-the-locate-command-in-the-mac-os-x-terminal/)

Noch weitere Tipps, anyone?
