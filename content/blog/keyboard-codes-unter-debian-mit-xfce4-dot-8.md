---
aliases:
- /blog/2012/04/20/keyboard-codes-unter-debian-mit-xfce4-dot-8
comments:
- author: Silvio Knizek
  content: "<p>Aber aber\u2026 sollten nicht inzwischen alle wissen, dass es schlecht
    ist, die .xml-Dateien von Xfce direkt an zufassen? Besser ist an dieser Stelle
    xfconf-query: \xBBxfconf-query -c xfce4-keyboard-shortcuts -p /commands/custom/hpModelock1
    -n -t string -s \"/usr/bin/i3lock\"\xAB<br>Wer es unbedingt graphisch braucht,
    kann auch xfce4-settings-editor verwenden</p><p>killermoehre<br>Xfce Supporter
    IRC</p>"
  date: '2012-04-20T19:54:39'
- author: noqqe
  content: "<p>\_Echt? Ne, war mir nicht wirklich klar. Wo kann man sich dazu denn
    schlaulesen, was passiert wenn ich an XML Files direkt hinfasse? </p>"
  date: '2012-04-21T09:13:06'
- author: Silvio Knizek
  content: "<p>Im besten Fall passiert nichts. Im schlimmsten Fall werden die XML-Dateien
    korrumpiert. Es ist einfach nicht empfehlenswert. Xfconf funktioniert \xFCber
    dbus-calls, die bei einer direkten Bearbeitung nicht ausgel\xF6st werden, womit
    die Applikationen nicht mit bekommen, das etwas ver\xE4ndert wurde.</p><p>killermoehre<br><a
    href=\"irc://irc.freenode.net/#xfce-de\" rel=\"nofollow\">irc://irc.freenode.net/#xfce-d...</a></p>"
  date: '2012-04-21T16:51:50'
- author: noqqe
  content: "<p>\_Alles klar. Angekommen und registriert :)<br>Danke!</p>"
  date: '2012-04-22T08:51:53'
- author: Silvio Knizek
  content: "<p>Xfce bringt \xFCbrigens seinen eigenes Screen-Lock-Frontend namens
    xflock4 mit. Einfach mal ausprobieren.</p>"
  date: '2012-04-23T19:01:11'
date: '2012-04-20T06:52:00'
tags:
- keyboard
- hardware
- xorg
- linux
- xfce
- debian
title: Keyboard Codes unter Debian mit xfce4.8
---

Zuerstmal sei gesagt, dass ich sowieso schon sehr begeistert bin wie viele von
den Spezialtasten auf dem Thinkpad schon Out-of-the-Box unter Debian
funktionieren. Trotzdem gab es noch ein paar die das nicht Taten oder zumindest
nicht mit der Software funktioniert hatte bei der ich es brauche.

## Keyboard Signale mit xev abfangen

Erstmal gehts darum rauszufinden was für lustige Tasten ich eigentlich gerade
drücke. Hier gibts verschiedene Spielereien mit dem Tool xev und sed:

    $ xev | sed -n 's/^.*keycode *\([0-9]\+\).*$/keycode \1 /p'
    keycode 36
    keycode 8
    keycode 160
    keycode 160

## XKeysymDB und benutzerdefiniertes Mapping

Nachdem man weiss, was für KeyCodes man gerade drückt kann man damit auch weiter
arbeiten. Diese KeyCodes müssen nun Keysyms zugeordnet werden.
Der X11 bringt hier automatisch eine kleine Übersicht mit die es sich
lohnt mal anzusehen:

    $ less /usr/share/X11/XKeysymDB

Obacht: Man kann keine neuen Keysyms definieren. Man muss Keysyms aus dem DB
File benutzen um diese dann auf den Keycode zu mappen. Dieses Mapping
findet folgendem File statt:

    $ vim ~/.Xmodmap
    keycode 160 = hpModelock1

Anschliessend kann man die Xmodmap neu laden:

    $ xmodmap ~/.Xmodmap

## xfce Keymapping

Der eigentlich spannende Teil ist, die Shortcuts jetzt mit xfce zu verknüpfen,
damit man bei Betätigung einen Befehl seiner Wahl ausführen kann. In xfce gibt
es eine kleine gui im Settings Manager (xfce4-settings-manager) mit der man
soetwas einrichten kann. Zumindest theoretisch. Leider können hier nur
Tastenkombinationen(!) hinterlegt werden. Keine Keysyms.

Aber das macht nichts. Wie so ziemlich alles gibts ein xml File in dem die Werte
stehen:

    $ vim .config/xfce4/xfconf/xfce-perchannel-xml/xfce4-keyboard-shortcuts.xml

``` xml
<property name="custom" type="empty">
  <property name="&lt;Alt&gt;F2" type="string" value="xfrun4"/>
  <property name="&lt;Control&gt;&lt;Alt&gt;Delete" type="string" value="xflock4"/>
  <property name="XF86Display" type="string" value="xrandr --auto"/>
  <property name="override" type="bool" value="true"/>
  <property name="hpModelock1" type="string" value="/usr/bin/i3lock"/>
</property>
```
Der letzte Eintrag wurde von mir per Hand eingetragen und funktioniert. Mit
fn+F2 kann ich jetzt (wie auf meinem Keyboard belabelt) meinen Bildschirm
sperren.

## Links

[http://en.gentoo-wiki.com/wiki/Multimedia_Keys](http://en.gentoo-wiki.com/wiki/Multimedia_Keys)

[http://wiki.ubuntuusers.de/Xmodmap](http://wiki.ubuntuusers.de/Xmodmap)

[http://wiki.xfce.org/faq#keyboard](http://wiki.xfce.org/faq#keyboard)