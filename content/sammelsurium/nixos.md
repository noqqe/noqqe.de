---
title: NixOS
date: 2015-08-20T08:43:56
tags:
- OS
- NixOS
---

#### System Commands

Um das System zu verändern

    vim /etc/nixos/configuration.nix
    nixos-rebuild switch

Die Konfiguration erstmal testen

    nixos-rebuild test

Einen Rollback zur letzten Version machen

    nixos-rebuild switch --rollback

#### Nix Paket Manager

Nach einem Paket suchen

    nix-env -qa | grep firefox

Pakete anzeigen und Status anzeigen lassen

    nix-env -qas

Paket installieren

    nix-env -i firefox

Ein Paket entfernen

    nix-env -e firefox

Channel updaten

    nix-channel --update nixpkgs

Upgrade aller Pakete machen

    nix-env -u '*'
    nix-env -u
    ## oder dry run
    nix-env -u --dry-run

Aktion von nix-env rückgängig machen

    nix-env --rollback

Alte unbenutzte Pakete aufräumen um Platz zu sparen

    nix-env --delete-generations old
    nix-collect-garbage -d
    nix-store --gc --print-dead

## Secure Passwords in configuration.nix

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
      in {
        description = "net-snmp daemon";
        wantedBy = [ "multi-user.target" ];
        serviceConfig = {
          ExecStart = "${pkgs.net_snmp}/bin/snmpd -f -c ${snmpconfig}";
          KillMode = "process";
          Restart = "always";
        };
      };
    }
