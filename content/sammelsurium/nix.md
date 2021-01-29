---
title: "Nix"
date: 2021-01-29T13:33:10+01:00
tags:
- Software
- Nix
- NixOS
---

Der Nix Paket Manager.

In seltenen Fällen benutze / brauche ich den `Nix` Paketmanager.
Ich benutze ihn eigentlich nur unter `NixOS`, weswegen die meisten dieser
Befehle bessere Pendants in `configuration.nix` haben. Hierfür siehe
[NixOS](/sammelsurium/nixos.md)

<!--more-->

Nach einem Paket suchen

    nix-env -qa | grep firefox

oder (dauert ewig)

    nix-env search firefox

Pakete anzeigen und Status anzeigen lassen

    nix-env -qas

Paket installieren, aufpassen, nur für den lokalen User! Systemweit geht nur
unter NixOS configuration.nix

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
