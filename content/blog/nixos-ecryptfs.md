---
title: "NixOS mit ecryptfs"
date: 2018-01-05T15:59:10+01:00
tags:
- NixOS
- ecryptfs
---

Dank [gpunktschmitz](https://gpunktschmitz.de) haben wir nun eine
Virtualisierungslösung in der [K4CG](https://k4cg.org) die mit Hyper-V umgesetzt wurde.
Ich habe damit noch nie was gemacht

Ich habe ein NixOS installiert und wollte ein verschlüsseltes
Volume via `ecryptfs` aufsetzen.

{{< figure src="/uploads/2018/01/ka.jpg" >}}

Also erstmal in die `/etc/nixos/configuration.nix` das `ecryptfs` hinzugefügt.

```
environment.systemPackages = with pkgs; [
  tmux
  ecryptfs
];
```

Via `nixos-rebuild switch` hab ich die neue Konfiguration dann angewendet.
Scheinbar reicht aber das reine Paket nicht aus um `ecryptfs` zu benutzen. Ich
muss das Kernel Modul noch laden!

```
boot.kernelModules = [
  "ecryptfs"
];
```

Ist das geil oder was? Natürlich hab ich mich zu früh gefreut. Ich hab auch
gleich den `mount` Block eingebaut, weswegen die Maschine nicht mehr bootete.

```
fileSystems."/data" = {
  neededForBoot = false;
  mountPoint = "/data";
  device = "/data";
  fsType = "ecryptfs";
  options = [
    "ecryptfs_cipher=blowfish"
    "ecryptfs_key_bytes=56"
    "ecryptfs_sig=xxx"
    "ecryptfs_passthrough=no"
    "ecryptfs_enable_filename_crypto=no"
  ];
};
```

Logisch, war ja auch irgendwie noch nichts initialisiert. Selbst wenn, wäre
keine Passphrase hinterlegt. Also Maschine kommt nicht mehr hoch, aber NixOS
legt für jede angewandte Konfiguration einen eigenen Grub Boot Entry an. Das
betrifft nicht nur den Kernel, sondern auch die Config des OS im Userspace. So
konnte ich einfach die letzte funktionierende Config booten und meinen Kram
fixen. Was übrigens nur ein `noauto` in den Optionen von `fstab` war.

{{< figure src="/uploads/2018/01/nixos-boot.png" >}}

Wenn meine VM nun bootet, mounte ich via

```
# mount /data
> Passphrase: xxx
```

die `ecryptfs` Partition und fertig.

Zum Abschluss muss ich noch sagen, dass es wahrscheinlich sinnvoll wäre das
ecryptfs im /home zu haben und bei Session Login via `pam` Modul automatisch
den `ecryptfs-passphrase-unwrap` zu triggern. Dazu war ich aber bisher zu faul
und werde es wahrscheinlich auch immer sein.
