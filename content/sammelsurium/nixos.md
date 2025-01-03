---
title: NixOS
date: 2021-01-29T13:40:58
tags:
- OS
- NixOS
---

Mein kleines aber feines Cheatsheet für [NixOS](https://nixos.org).

<!--more-->

## System Commands

Um das System zu verändern

    vim /etc/nixos/configuration.nix
    nixos-rebuild switch

Die Konfiguration erstmal testen

    nixos-rebuild test

Einen Rollback zur letzten Version machen

    nixos-rebuild switch --rollback

## Configuration

Was folgt ist eine Sammlung an coolen Schnippseln die ich immer wieder
verwende.

### Packages Installieren

Pakete am besten nicht für einzelne User installieren (mittels `nix`) sondern
in der Configuration von NixOS.

```nix
environment.systemPackages = with pkgs; [
  bzip2
  curl
  file
  fzf
  git
  gzip
  htop
];
```

### User Anlegen

```nix
users.extraUsers.noqqe = {
  isNormalUser = true;
  uid = 1000;
  extraGroups = [ "wheel" "networkmanager" ];
  openssh.authorizedKeys.keys = [
    "ssh-ed25519 xxxx
    "ssh-ed25519 xxx
  ];
};
```

### SSH Configuration

Im Grunde muss man da wenig machen.

```nix
# SSH
services.openssh = {
 enable = true;
 allowSFTP = true;
 forwardX11 = false;
 permitRootLogin = "yes";
 passwordAuthentication = true;
 challengeResponseAuthentication = false;
};
```

### Fish als default Shell

```nix
# Fish
programs.fish.enable = true;

users.extraUsers.noqqe = {
  isNormalUser = true;
  shell = pkgs.fish;
};
```

### Selbstverwaltung und Cleanup

u

Garbage Collector des Nix Paketstores automatisch triggern lassen
um mehr Plattenplatz freizugeben.

```nix
nix.gc = {
  automatic = true;
  options = "--delete-older-than 30d";
};
```

Mittels `nix-store --optimise` werden Pakete durch Hardlinks ersetzt und
dabei meist gut Plattenplatz frei. Das geht auch automatisch via systemd
Timer

```nix
nix.optimise.automatic = true;
```

Das ganze geht auch manuell

```shell
nix-env --delete-generations old
nix-collect-garbage -d
nix-store --gc --print-dead
```

### Autmatische Systemupdates

NixOS kann sich gut selbst verwalten, und wenn etwas schief geht kann man
jederzeit via `grub` die alte Konfiguration hochbooten.

```nix
system.autoUpgrade = {
  enable = true;
  allowReboot = true;
};
```

## systemd Unit

Eigene `systemd` Units

```nix
systemd.services.rezeptionistin = {
  enable = true;
  description = "Rezeptionistin IRC BOT";
  wants = [ "network.target" ];
  wantedBy = [ "multi-user.target"];
  environment = {
OPENSSL_CONF = "/etc/ssl/";
SSL_CERT_FILE = "/etc/ssl/certs/ca-certificates.crt";
LC_ALL="C";
  };
  serviceConfig = {
    ExecStart = "/usr/local/rezeptionistin/main.py -c /usr/local/rezeptionistin/config.ini";
    WorkingDirectory = "/usr/local/rezeptionistin/";
    User = "ircbot";
    Restart = "always";
    RestartSec = 30;
  };
};
```

## systemd Timer

Timer braucht natürlich eine Referenz auf einen Service vom Typ `oneshot`

```nix
systemd.services.offsitebackup = {
  enable = true;
  wants = [ "network.target" "multi-user.target"];
  path = [ "/run/current-system/sw/bin" pkgs.openssh ];
  serviceConfig = {
    Type = "oneshot";
    StandardOutput = "journal";
    ExecStart = "restic backup <path>";
    User = "root";
  };
};

# Trigger the oneshot service once a day
systemd.timers.offsitebackup = {
  enable = true;
  wantedBy = [ "multi-user.target"];
  timerConfig = {
    OnBootSec = "15min";
    OnUnitActiveSec = "1d";
  };
};
```

## Passwörter

Um die Files unter Versionsverwaltung stellen zu können lohnt es sich die
Passwörter in eine Art Passwordstore File auszulagern. In Klartext stehen sie
dort aber immernoch!

```nix
{ pkgs, ... }:

let
  passwords = import ./passwords.nix;
in
{
  systemd.services.snmpd = let
    snmpconfig = pkgs.writeTextFile {
      name = "snmpd.conf";
      text = ''
        rocommunity ${passwords.snmp}
        disk / 10000
      '';
  };
};
}
```

## Eigene Services / Extending NixOS

Auf Basis von [Extending NixOS](https://nixos.wiki/wiki/Extend_NixOS) ein
dynamisch gebauter Service

    services.ircClient.enable
    services.ircClient.user = "username";

also

    services.ircClient.<name>.enable
    services.ircClient.<name>.user = "username";

Um das zu erreichen braucht man ein bisschen `nix` Code.

```nix
{config, pkgs, lib, ...}:

with lib;
let
  cfg = config.services;

  # define variable for single service option for <name>
  msgOptions =
    { ... }: {
      options = {
          message = mkOption {
            default = "nix";
            type = with types; uniq string;
            description = ''
              Message to send.
            '';
          };
      };
    };

  # define set variable for systemd.service
  mkService = name: value: let
  in {
    wantedBy = [ "multi-user.target" ];
    after = [ "network.target" ];
    description = "Send my message.";
    serviceConfig = {
      Type = "oneshot";
      User = "root";
      StandardOutput = "journal";
      ExecStart = ''${pkgs.coreutils}/bin/echo "Hello ${value.message}"'';
    };
  };
in
{
  # build up 1 option to contain set of all services.msg.* from above
  options.services = {
    msg = mkOption {
      default = {};
      type = types.attrsOf (types.submodule msgOptions);
      description = "Top msg service gerne wieder";
    };
  };

  # merge all msgOptions to the systemd.services set from the let part above
  config = lib.mkIf (cfg.msg != {}) (lib.mkMerge [
    {
      systemd.services = mapAttrs' (n: v: nameValuePair "msg-${n}" (mkService n v)) cfg.msg;
    }
  ]);
}
```

und so kann ich dann meinen eigenen dynamischen Service einbauen
```nix
{
   imports = [
     ./my-service.nix
   ];

   services.msg.foo.message = "foo";
   services.msg.bar.message = "bar";
   services.msg.baz.message = "foo";

}
```
Also anstelle von 3 einzelnen systemd services kann ich das auch
dynamisieren.
