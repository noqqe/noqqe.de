---
title: "NixOS und 6 Major Upgrades zu 20.03"
date: 2020-06-23T13:48:19+02:00
tags:
- NixOS
---

Wenn du in deinem örtlichen [Hackspace](https://k4cg.org) eine Kiste findest und dir auffällt,
das du diese wahrscheinlich seit dem initialen Aufsetzen nicht mehr angefasst
hast. `nixos-version` sagt dir `17.03`. Uff. Das sind 4 Jahre! Es ist zwar
nur unser [Bastion](https://en.wikipedia.org/wiki/Bastion_host) Host, aber
trotzdem. Erstmal 6 Major Releases nachholen.

<!--more-->

{{< figure src="/uploads/2020/06/6upgrades.png" caption="Bash History von 6 NixOS upgrades" >}}

Das ist wörtlich meine Bash History beim Upgrade. 6 Major Upgrades am Stück
und ich musste eigentlich nichts tun! Ich feiere es sehr hart. Bei 18.03 hat
sich mal das Network Interface von `eth0` auf `ens18` umbenannt, aber ehrlich
gesagt weiss ich nichtmal welche Komponente das macht. Linux Kerneltreiber?
Systemd? NixOS? Egal.

Allerdings war es mit 500MB RAM ein bisschen knapp und so lief ich in ein
paar `OOM` Nachrichten im `dmesg`. Ich kann fix ein Swap File einbauen:

```nix
swapDevices = [
  {
    device = "/var/cache/private/tempswap";
  }
];
```

Dann liefen auch die OS Upgrades einfacher. Initialisieren (`mkswap` & `dd`)
muss man es allerdings selbst.

Auch unter Linux kann es seamless Upgrades geben :)
