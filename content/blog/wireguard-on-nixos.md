---
title: "WireGuard on NixOS"
date: 2019-06-03T12:35:55+02:00
tags:
- VPN
- WireGuard
- NixOS
---

[WireGuard](https://www.wireguard.com/) ist so ziemlich das was man als Admin immer wollte was VPN angeht.
Mir zog es immer die Gänsehaut auf wenn ich nur an dieses OpenVPN gegraddel
denken musste.

`wg` dagegen ist schön, einfach und unkomplizert. Über die Kernel-Modul
Implementation kann man streiten, aber das ist ja kein muss, da es auch
Userland Implementationen gibt wie unter
[iOS](https://www.wireguard.com/install/) oder letztens in
[OpenBSD](https://blog.jasper.la/wireguard-on-openbsd.html) als Paket

Ungefähr genau so schön ist die Implementierung in `NixOS`, welche mal wieder
wunderbare Option Bindings zur Verfügung stellt. Statt Kommandos und Befehle
ins Terminal zu tippern, funktioniert die Konfiguration wie gewohnt in
`configuration.nix`

``` nix
networking.wireguard.interfaces.wg0 = {
  ips = [ "192.168.42.1/24" ];
  privateKey = "..secret..";
  listenPort = 51666;
  peers = [
    {
      allowedIPs = [ "192.168.42.0/24" ];
      publicKey = "CiEyx82EHuibAc4AvB+BRbTVh9p1mDNIhBQ64mWUMA8=";
    }
  ];
};
```

Und fertig ist schon der WireGuard VPN Server.

{{< figure src="/uploads/2019/06/wireguard.png" >}}

Um jetzt den Client in meinem Beispiel eines anderen NixOS anzubinden, sieht
ähnlich aus.

``` nix
networking.wireguard.interfaces.wg0 = {
  ips = [ "192.168.42.2/24" ];
  privateKey = "..other secret..";
  peers = [
    {
      allowedIPs = [ "192.168.42.0/24" ];
      endpoint = "host.example.com:51666";
      publicKey = "fY0GiZHUVGSyuteTwrWTnBpURiGOwtJtP0p4gxBj8UY=";
    }
  ];
};
```

Der Client braucht beim Server dabei keinen `endpoint`, weil er sozusagem im
Roaming Modus läuft.

Ein bisschen Firewalling muss man halt noch betreiben, aber das funktioniert
im NixOS Stil sehr ähnlich

``` nix
networking.firewall = {
  enable = true;
  rejectPackets = true;
  allowedUDPPorts = [
    51666 # wireguard
  ];
  allowedTCPPorts = [
    22 # ssh
  ];
};
```

Ich hab hierbei einen festen Port gewählt, normalerweise wird dieser von
WireGuard dynamisch basierend auf dem Hostname gewählt.

Danke WireGuard das ihr VPN erträglich gemacht habt.
