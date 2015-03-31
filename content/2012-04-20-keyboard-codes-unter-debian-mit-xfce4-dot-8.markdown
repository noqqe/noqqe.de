---
layout: post
title: "Keyboard Codes unter Debian mit xfce4.8"
date: 2012-04-20T08:52:00+02:00
comments: true
categories:
- ubuntuusers
- Web
- Hardware
- Debian
- Linux
keywords: "keyboard, tastatur, shortcuts, Tastenkombinationen, fn, f2, xfce,
xev, sed, keycode, xmodmap"
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

``` xml xfce4-keyboard-shortcuts.xml
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
