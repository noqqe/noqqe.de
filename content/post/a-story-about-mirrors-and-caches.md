---
layout: post
title: "A Story about Mirrors and Caches"
date: 2013-08-31T19:01:00+02:00
comments: true
categories:
- Debian
- apt
- apt-cacher-ng
- Updates
- Ubuntuusers
- puppet
- mirror
- cache
- lxc
- linuxcontaier
- linux
- kernel
- osbn
- ubuntuusers
---

[LinuxContainer](http://lxc.sourceforge.net/) sind für mich Wegwerfartikel. Benutze sie als Sandboxes um
neue Software auszuprobieren, bestimmte Software zu Betreiben und um der
[Dependency-Hölle](https://en.wikipedia.org/wiki/Dependency_hell) zu entkommen.
Auch dieser Blogpost wurde in so einer LXC-VM gebaut.


### Mirroring

Natürlich fällt bei (im Schnitt) 20-25 VMs zusätzlicher Verwaltungsaufwand an,
was ich bisher sehr gut mit Puppet erschlagen konnte. Ein anderes Problem sind
aber die OS Updates der Maschinen. Dazu nutze ich alle Monat
[clusterssh](https://github.com/duncs/clusterssh). Klar, wenn alle Maschinen
gleichzeitig Updates ziehen macht das weder meiner Internetanbindung noch dem
Mirror-Betreiber Spaß.

{{< figure src="/uploads/2013/08/w00t.jpg" >}}

Ich rechnete etwas hin und her, ob es sich wirklich lohnt einen eigenen lokalen Mirror zu betreiben.
Für die richtige Architektur von main/non-free/contrib braucht wohl ~150GB. Letztendlich rentiert sich das von
der Masse und Update-Synchronisierung erst hart spät. Ich tat es aber dann doch,
was mit dem [apt-mirror](http://apt-mirror.github.io/) einfach und schnell ging.

### Caching

Es gibt Caching Software für Debian Repos! Nach einem Schwatz mit einem Kollegen
kam die Erkenntnis das alles zu syncen garnicht nötig war.
Auf diese Idee war ich nicht gekommen. Nichtmal an die Möglichkeit gedacht. Konkret
möchte man [apt-cacher-ng](https://www.unix-ag.uni-kl.de/~bloch/acng/) nutzen.

`apt-cacher-ng` agiert jetzt quasi als Proxy zwischen den Containern und dem
konfigurierten Debian Mirror.

```
$ sudo aptitude install apt-cacher-ng
$ sudo /etc/init.d/apt-cacher-ng start
```

Möchte ein Client Updates herunterladen, bemüht sich apt-cacher-ng nach
Möglichkeit bereits heruntergeladene Pakete aus
`/var/cache/apt-cacher-ng/$mirror/debian` auszuliefern. Das spart Bandbreite und
Speicherplatz.

Clientseitig lässt sich die Nutzung des Proxys auf zweierlei Varianten
konfigurieren.

Entweder durch das `apt` HTTP Proxy Feature

    Acquire::http { Proxy "http://10.10.0.10:3142"; };

oder durch `sources.list` URLs

    deb http://10.10.0.10:3142/debian.cs.fau.de/debian stable main contrib non-free

Mir gefiel die erste Variante allerdings wesentlich besser. Ich kanns nicht
begründen, fühlt sich aber besser an.

### Puppet

Ausgerollen konnte ich die Konfiguration gegen apt-cacher-ng einfach über das
Puppetlabs [apt](https://forge.puppetlabs.com/puppetlabs/apt) Modul.

```
class { 'apt':
  always_apt_update    => false,
  proxy_host           => '10.10.0.10',
  proxy_port           => '3142',
}

apt::source { 'wheezy-mirror':
  location    => 'http://debian.cs.fau.de/debian/',
  release     => 'wheezy',
  include_src => false,
  repos       => 'main contrib non-free',
}
```

Außerdem war dies der erste Post den ich mit der [Cherry G80-3000LSCDE-2](http://www.cherry.de/cid/b2b_keyboards_G80-3000.htm?rdeLocaleAttr=en&cpssessionid=SID-837EAC29-341CE33E&WT.mc_id=)
verfasste. Es ist schöner als vorher.
