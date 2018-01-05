---
title: "Nixos mit ecryptfs"
date: 2018-01-05T13:28:10+01:00
tags:
- NixOS
- ecryptfs
---

Dank [gpunktschmitz](https://gpunktschmitz.de) haben
wir nun eine Virtualisierungslösung in der [K4CG](https://k4cg.org) auf der man
ein paar VMs hosten kann. Ich habe ein NixOS installiert und wollte ein verschlüsseltes Volume via `ecryptfs` aufsetzen.

Also erstmal in die `/etc/nixos/configuration.nix` das `ecryptfs` hinzugefügt.

```
environment.systemPackages = with pkgs; [
  tmux
  ecryptfs
];
```

Via `nixos-rebuild switch` hab ich die neue Konfiguration dann angewendet. Scheinbar reicht aber das reine Paket nicht aus um
`ecryptfs` zu benutzen. Ich muss das Kernel Modul noch laden!

```
boot.kernelModules = [
  "ecryptfs"
];
```

Ist das geil oder was? Natürlich

