---
title: "Systemd Vorzüge"
date: 2019-04-10T21:02:22+02:00
tags:
- systemd
- matomat
---

In letzter Zeit kam ich öfters in die Verlegenheit mit Systemd zu
interagieren. Der Matomat in der K4CG hatte so ein paar Spezialanforderungen
umzusetzen.

{{< figure src="/uploads/2019/04/matomat.png" >}}

Systemd hat den Ruf bloated und buggy zu sein, aber ich hatte in letzter Zeit
2 schöne Aspekte kennen gelernt.

## Auto Login

Wir bedienen den [Matomat](https://k4cg.org/index.php/Projekt:Heiko) über
TTY1. Beim Starten soll auf TTY1 einfach das Programm starten. Früher musste
dafür ein `/etc/inittab` Eintrag her. Heute gibt es dafür einen Systemd
Service

```
vim /etc/systemd/system/getty@tty1.service.d/override.conf
[Service]
Type=simple
ExecStart=-/sbin/agetty --autologin heiko --noclear %I 38400 linux
```

Es wird also automatisch der Benutzer `heiko` eingeloggt.

## Suspend

Aus irgend einem Grund hatte die alte Hardware immer das Gefühl sie müsse
sofort in den Suspend Modus wechseln.

Anfangs hatte ich einen
[Hack](https://unix.stackexchange.com/questions/188559/disable-suspend-sleep-mode-completely-on-fedora-21)
befolgt der einfach die beiden Services maskiert

```
sudo systemctl mask suspend.target
sudo systemctl mask sleep.target
```

Es hat tatsächlich mein Problem gelöst! Der Laptop blieb an. Ich dachte
“cool, so ein Service für den Zustand der Maschine”.

Nach einiger Zeit lief dann aber die Platte voll. Als ich nachsah warum, fand
ich ein 10G großes `auth.log`.

```
systemd-logind[528]: Suspending...
systemd-logind[528]: Failed to execute operation: Unit suspend.target is masked.
systemd-logind[528]: Suspending...
systemd-logind[528]: Failed to execute operation: Unit suspend.target is masked.
systemd-logind[528]: Suspending...
systemd-logind[528]: Failed to execute operation: Unit suspend.target is masked.
systemd-logind[528]: Suspending...
systemd-logind[528]: Failed to execute operation: Unit suspend.target is masked.
systemd-logind[528]: Suspending...
```

Ca. 30 mal pro Sekunde. Also versuchte irgendwas namens `systemd-logind`
immernoch den Laptop in den Suspend Modus zu befördern.

Ich fand dann
[logind.conf](https://www.freedesktop.org/software/systemd/man/logind.conf.html#)
und setzte dann folgende Settings.

```
$ vim /etc/systemd/logind.conf
[Login]
HandleSuspendKey=ignore
HandleHibernateKey=ignore
HandleLidSwitch=ignore
HandleLidSwitchDocked=ignore
```

Empfand ich mal unendlich schick, diese ganzen Hardware-nahen Dinge die
früher jede Distri über komische Shell Scripts und Custom Geklöppel lösen
musste an so einem Ort zu pflegen. Und noch dazu so schön und einfach.

"Es ist ja nicht alles schlecht mit Systemd…"
