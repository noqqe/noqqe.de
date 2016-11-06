---
categories:
- administration
- code
- databases
- devops
- linux
- osbn
- planetenblogger
- ubuntuusers
date: 2015-09-05T13:55:46+02:00
tags:
- DevOps
- NixOS
- Nix
- Distribution
- Paketmanager
- mlmmj
title: NixOS
---

Ja, tatsächlich mal wieder Linux. Der neue Server der [K4CG](https://k4cg.org)
läuft mit NixOS. Habe ich damit vorher schonmal was gemacht? Nö. Aber das
Konzept ist vielversprechend.

{{< figure src="/uploads/2015/09/nixos.jpg" >}}

NixOS ist nicht wie jede andere Distro. Sie besteht aus dem Paket-Manager `nix`
und einer einzigen deklarativen Config die mit Modulen benutzt wird.
Verschiedene Versionen von Libraries oder andere Software koexistieren auf dem
System. So ist es auch möglich das User ihre eigenen Programme installieren ohne
das System zu beeinflussen.

### NixOS Config

Um den `sshd` im OS zu konfigurieren, sieht die `/etc/nixos/configuration.nix`
unter anderem so aus..

```
# SSH
services.openssh = {
  enable = true;
  allowSFTP = false;
  forwardX11 = false;
  permitRootLogin = "no";
  passwordAuthentication = false;
  challengeResponseAuthentication = false;
};
```

oder einen neuen Benutzer anlegen

```
users.extraUsers.noqqe = {
  isNormalUser = true;
  extraGroups = [ "wheel" "networkmanager"];
  uid = 1000;
  openssh.authorizedKeys.keys = [ "ssh-rsa AAAAB..." ];
};
```

oder ein `Tor` Relay aufsetzen

```
services.tor.relay = {
  nickname = "nixe.k4cg.org";
  portSpec = 9001;
  exitPolicy = "reject *:*";
  isExit = false;
  bandwidthRate = "200000" ; # 200kb
  bandwidthBurst = "300000" ; # 300kb
  accountingStart = "day 00:00"; # reset daily
  accountingMax = "4 GB" ;
};
```

Alles funktioniert eigentlich ziemlich ähnlich. Ein bewährtes Vorgehen ist,
einfach auf [nixos.org Options](https://nixos.org/nixos/options.html) nachsehen
ob es den gewollten Service bereits als Modul gibt. Siehe mehr Beispiele aus
dem K4CG Server [hier](https://gist.github.com/noqqe/403167d30732d30cccd1)

Obacht: Nur weil es ein Paket in `nix` gibt, heisst das nicht immer, dass es
auch ein Modul gibt.

### NixOS System Commands

Um getätigte Changes am System live zu schalten

    vim /etc/nixos/configuration.nix
    nixos-rebuild switch

Einen Rollback zur letzten Version machen

    nixos-rebuild switch --rollback

### Nix Paket Manager

Alle Paket-Befehle beziehen sich nur auf den aktuellen Benutzer.  Installiere
ich also als `noqqe` ne historisch wertvolle Version von Firefox, interessiert
das (außer mich) keinen auf dem System. Dieses Konstrukt ist über ein
`nixos-profile` geregelt. Im Endeffekt Symlinks und ein bisschen im `PATH`
rumfummeln.

{{< figure src="/uploads/2015/09/profile.png" >}}

Ein kleines Cheatsheet:

Nach einem Paket suchen

    nix-env -qa | grep firefox
    nix-env -qas

Paket installieren/entfernen

    nix-env -i firefox
    nix-env -e firefox

Update aller Paket des aktuellen Users

    nix-env -u
    nix-env -u --dry-run

Na klar. Die ganzen Versionen und parallel Installationen verursachen ein
bisschen dramatisch angestiegene Diskspace Usage. Auch dafür soll gesorgt sein.

    nix-env --delete-generations old
    nix-collect-garbage -d
    nix-store --gc --print-dead

Mit diesen Commands werden alle unreferenzierten Versionen von zB. `netcat` entfernt.

### Probleme

Alle möglichen Sachen laufen auf dem K4CG Server. Die ein Webserver, Mediawiki,
MariaDB, IRC Bouncer und auch die Mailingliste.

Für die Mailingliste mit `mlmmj` war die Erfahrung bzgl. NixOS aber ziemlich enttäuschend.
Das Modul ist scheinbar seit ziemlich langer Zeit nicht mehr gepflegt worden. So
musste ich da erstmal etwas
[herumpatchen](https://github.com/NixOS/nixpkgs/commit/d43496300e0064afa940c8fe822762983051793c)
und `postfix` Config [updaten](https://github.com/NixOS/nixpkgs/pull/9661).

Um ein IPv6 Gateway zu hinterlegen, reicht die momentane stable Version
14.12 nicht aus. Dafür muss man dem Unstable Channel folgen.

Zusammenfassung: Ist schön, macht Spass, macht Sinn, ob es sich durchsetzt
keine Ahnung. Ob es auf einem produktiven System Sinn macht? Wir sind nun mal
ein Hackerspace.

