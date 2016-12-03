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


## Module Development

> noqqe | i can also use a local nixpkgs repo for testing module changes, right?
> noqqe | in wiki there is only described i can use that for package development
> kmicu | You can use a local nixpkgs or import a local module definition.
> noqqe | are there docs on that somewhere? im struggling getting that working
> kmicu | Almost always manual has a better info than wiki https://nixos.org/nixos/manual/sec-writing-modules.html
> noqqe | mh, when i do "nixos-rebuild -I /path/to/git/repo switch --upgrade" it does not apply the changes.. am i holding it wrong?
> kmicu | You don’t need to use ‘--upgrade’; that flag is for a channel update.
> kmicu | It’s hard to say why you cannot see changes w/o any details.
> kmicu | Maybe module is not enabled? Maybe is not registered in module-list.nix… etc.
> noqqe | okay. in work on the mlmmj module, because it needs some love. so i checked out the repo from master, edited
>       | https://github.com/NixOS/nixpkgs/blob/master/nixos/modules/services/mail/mlmmj.nix and did the comannd from above
> noqqe | i did nothing more
> noqqe | to test, i added a "touch /tmp/foobar" to the activation scripts, to see if it works
> kmicu | Ahh ‘-I nixpkgs=/path/to/your/fork’ and you should not use master.
> kmicu | (Unless you want to build a lot of things ;)
> noqqe | mh i dont want to build pkgs.. :( i just want to edit modules and test them before i make a pull request
> kmicu | So checkout your current nixpkgs revision https://github.com/NixOS/nixpkgs-channels
> noqqe | ah, f*ck. well, right.
>    -- | Notice(NixOS_GitHub): [nixpkgs] edolstra pushed 7 new commits to release-15.09: http://git.io/vZuy7
>    -- | Notice(NixOS_GitHub): nixpkgs/release-15.09 469b79b Eelco Dolstra: Remove openjdk namespace pollution...
>    -- | Notice(NixOS_GitHub): nixpkgs/release-15.09 7def439 Eelco Dolstra: Remove upower-old...
>    -- | Notice(NixOS_GitHub): nixpkgs/release-15.09 361d6cf Eelco Dolstra: upower: Remove unused dependencies...
> noqqe | kmicu: that works..thank you a lot
> kmicu | You could add your external module definition to ‘imports = [ … /my/module.nix … ]’ in configuration.nix. You don’t need to use fork, but if you
>       | want to contribute back then use fork, checkout channel revision, test with ‘-I nixpkgs=/path/to/fork’, more at
>       | http://hydra.nixos.org/build/25709458/download/1/nixpkgs/manual.html#chap-submitting-changes


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
