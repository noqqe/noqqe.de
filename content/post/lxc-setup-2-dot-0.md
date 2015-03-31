---
type: post
title: "LXC Setup 2.0"
date: 2012-10-07T20:00:00+02:00
comments: true
categories:
- ubuntuusers
- osbn
- Debian
- Linux
- Bash
- LXC
---

Es ist ja nun schon etwas her... da hab ich über mein initales Setup der LXC
Container (ja ich weiss, das ist wie HIV Virus sagen) [gebloggt](/blog/2012/05/14/lxc-not-ready-for-production-huh/).
Jedenfalls hat sich mein Setup, wie die Tools mittlerweile weiterentwickelt und
Dinge haben sich geändert.

{{< figure src="/uploads/2012/10/lxcsetup.png" >}}

## LXC Directory

Will man einen anderen Platz für seine LinuxContainer, dann kann man diesen
in `/etc/default/lxc` definieren

```
# Directory containing the container configurations
CONF_DIR=/etc/lxc
# Start /etc/lxc/example.conf, /etc/lxc/autostart.conf, etc.
#CONTAINERS="example autostart container"

LXC_DIRECTORY="/home/lxc"
```

FHS Konformität ist was anderes... aber zwecks meiner Partitionierung und
Verschlüsselungstechnisch macht das auf meinem Thinkpad einfach mehr Sinn.

## Meine Container

Ich nutze meine Container teils als Sandboxing, teils um der Dependency Hell zu
entkommen und Teils um Dingen nur bestimmte Ressourcen zuzuweisen. Das spiegelt
sich auch in meiner Übersicht wieder:

```
/home/lxc
├── vm10-core
├── vm11-stable
├── vm12-testing
├── vm13-sid
├── vm14-experimental
├── vm20-redis
├── vm21-matomat
├── vm22-bind
└── vm25-graylog2
```

vm10 ist eine Art Management VM. Hostet verschiedene Dinge.
vm{11..14} sind dabei meine Testcases für jeden Kanal von Debian.
Ab vm20 befinden sich dann meine Sandkästen zum spielen. Hier räume ich auch
regelmäßig immer mal wieder auf und schmeisse Dinge weg.

## Netze & DNS

Konvention ist alles. Die Ziffern nach vm kommen wie man sich vielleicht denken
konnte nicht von ungefähr. Sie beinhalten das letzte Oktett der jeweiligen der IP
des Containers

Innerhalb des 10.10.0.0/24 Netzes, dass über dnsmasq auch DHCP macht, regel ich
"DNS" über /etc/hosts Einträge. Hier gibt schon bestrebungen das automatisch zu
machen. Dauert aber noch.

```
# LXC Setup
10.10.0.10 vm10.lxc.local
10.10.0.11 vm11.lxc.local
10.10.0.12 vm12.lxc.local
10.10.0.13 vm13.lxc.local
10.10.0.14 vm14.lxc.local
10.10.0.20 vm20.lxc.local
10.10.0.21 vm21.lxc.local
10.10.0.22 vm22.lxc.local
10.10.0.25 vm25.lxc.local
```

und zur Erleichterung dann noch ein `search lxc.local` in die `/etc/resolv.conf`

## Cloning new Containers

Hier kommt wieder einmal raus, warum lxc m.E. noch nicht fertig ist. Es gibt
`lxc-clone`. Hört sich gut an. Nur nicht wenn man das default Home
verändert hat. Mit der Ausgangsconfig

```
# mounts point
lxc.mount.entry=proc /home/lxc/vm11-stable-template/rootfs/proc proc nodev,noexec,nosuid 0 0
lxc.mount.entry=devpts /home/lxc/vm11-stable-template/rootfs/dev/pts devpts defaults 0 0
lxc.mount.entry=sysfs /home/lxc/vm11-stable-template/rootfs/sys sysfs defaults 0 0
```

Wird aus dem geclonten Container (`lxc-clone -o vm11-stable-template -n
vm12-testing-template`)

```
# mounts point
lxc.mount.entry=proc /home/lxc/vm11-stable-template/rootfs/proc proc nodev,noexec,nosuid 0 0
lxc.mount.entry=devpts /home/lxc/vm11-stable-template/rootfs/dev/pts devpts defaults 0 0
lxc.mount.entry=sysfs /home/lxc/vm11-stable-template/rootfs/sys sysfs defaults 0 0
lxc.utsname = vm12-testing-template
lxc.rootfs = /var/lib/lxc/vm12-testing-template/rootfs
```

Das Default Config File scheint hier keine Rolle zu spielen. Deshalb hab ich
mir hier auch meine eigene Clone Komponente geschrieben, da kann ich auch gleich
das richtige Netzwerk konfigurieren.

``` bash 
function vm_clone () {
    if [ -z "$1" -o -z "$2" ] || [ ! -d "$LXC_DIRECTORY/$1" ]; then
        echo "ERROR: non-correct usage - see help"
        echo "ERROR: $LXC_DIRECTORY/$2 may exist or $LXC_DIRECTORY/$1 doesn't."
        exit 3
    fi

    # Copy configuration
    echo "INFO: Copying container..."
    cp -a $LXC_DIRECTORY/$1 $LXC_DIRECTORY/$2

    echo "INFO: Changing lxc config..."
    sed -i "s/$1/$2/g" $LXC_DIRECTORY/$2/config

    # ATTENTION!
    # This requries a convention of vm$IP-name-of-vm
    echo "INFO: Getting last octet from vmXX..."
    OCTET=$(echo $2 | sed 's#vm\(..\).*#\1#')

    echo "INFO: Making network config with 10.10.0.$OCTET..."
    sed -i "s/address.*/address 10.10.0.$OCTET/" $LXC_DIRECTORY/$2/rootfs/etc/network/interfaces
}
```

## mlxc 2.0

Außerdem hab ich viel an mlxc gebaut. Habe den
[gist](https://gist.github.com/2693967) aktualisiert.

Changelog

* VM stoppen
* Alle VMs stoppen
* Cloning
* Besseres Help
* Einheitlicher Output

